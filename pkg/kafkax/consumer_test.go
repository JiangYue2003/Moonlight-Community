package kafkax

import (
	"context"
	"errors"
	"io"
	"sync"
	"testing"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/require"
)

type fakeConsumerReader struct {
	mu       sync.Mutex
	messages []kafka.Message
	fetched  int
	commits  []int64
	events   []string
}

func (r *fakeConsumerReader) FetchMessage(context.Context) (kafka.Message, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.fetched >= len(r.messages) {
		return kafka.Message{}, io.EOF
	}
	m := r.messages[r.fetched]
	r.fetched++
	r.events = append(r.events, "fetch:"+offsetString(m.Offset))
	return m, nil
}

func (r *fakeConsumerReader) CommitMessages(_ context.Context, messages ...kafka.Message) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, m := range messages {
		r.commits = append(r.commits, m.Offset)
		r.events = append(r.events, "commit:"+offsetString(m.Offset))
	}
	return nil
}

func (r *fakeConsumerReader) Close() error { return nil }

type fakeMessageWriter struct {
	messages []kafka.Message
	err      error
}

func (w *fakeMessageWriter) WriteMessages(_ context.Context, messages ...kafka.Message) error {
	w.messages = append(w.messages, messages...)
	return w.err
}

func (w *fakeMessageWriter) Close() error { return nil }

func TestConsumeMessagesRetriesSameMessageBeforeFetchingNext(t *testing.T) {
	reader := &fakeConsumerReader{messages: []kafka.Message{
		{Partition: 0, Offset: 41},
		{Partition: 0, Offset: 42},
	}}
	attempts := make(map[int64]int)
	var handled []int64

	err := consumeMessages(
		context.Background(),
		ConsumerConfig{},
		reader,
		nil,
		func(_ context.Context, message kafka.Message) error {
			handled = append(handled, message.Offset)
			reader.mu.Lock()
			reader.events = append(reader.events, "handle:"+offsetString(message.Offset))
			reader.mu.Unlock()
			attempts[message.Offset]++
			if message.Offset == 41 && attempts[message.Offset] < 3 {
				return errors.New("transient redis failure")
			}
			return nil
		},
		noWait,
	)

	require.NoError(t, err)
	require.Equal(t, []int64{41, 41, 41, 42}, handled)
	require.Equal(t, []int64{41, 42}, reader.commits)
	require.Equal(t, []string{
		"fetch:41",
		"handle:41",
		"handle:41",
		"handle:41",
		"commit:41",
		"fetch:42",
		"handle:42",
		"commit:42",
	}, reader.events)
}

func TestConsumeMessagesStopsRetryingWhenContextIsCanceled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	reader := &fakeConsumerReader{messages: []kafka.Message{
		{Partition: 0, Offset: 7},
		{Partition: 0, Offset: 8},
	}}

	err := consumeMessages(
		ctx,
		ConsumerConfig{},
		reader,
		nil,
		func(context.Context, kafka.Message) error {
			cancel()
			return errors.New("redis unavailable")
		},
		func(ctx context.Context, _ time.Duration) error {
			<-ctx.Done()
			return ctx.Err()
		},
	)

	require.NoError(t, err)
	require.Equal(t, 1, reader.fetched)
	require.Empty(t, reader.commits)
}

func TestConsumeMessagesProductionBackoffExitsOnCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	reader := &fakeConsumerReader{messages: []kafka.Message{{Partition: 0, Offset: 9}}}
	attempts := 0

	err := consumeMessages(
		ctx,
		ConsumerConfig{},
		reader,
		nil,
		func(context.Context, kafka.Message) error {
			attempts++
			cancel()
			return errors.New("redis unavailable")
		},
		waitForBackoff,
	)

	require.NoError(t, err)
	require.Equal(t, 1, attempts)
	require.Equal(t, 1, reader.fetched)
	require.Empty(t, reader.commits)
}

func TestConsumeMessagesWritesPoisonMessageBeforeCommittingIt(t *testing.T) {
	reader := &fakeConsumerReader{messages: []kafka.Message{
		{Topic: "events", Partition: 2, Offset: 21, Value: []byte("poison")},
		{Topic: "events", Partition: 2, Offset: 22, Value: []byte("valid")},
	}}
	writer := &fakeMessageWriter{}
	var handled []int64

	err := consumeMessages(
		context.Background(),
		ConsumerConfig{MaxRetries: 2, DlqTopic: "events-dlq"},
		reader,
		writer,
		func(_ context.Context, message kafka.Message) error {
			handled = append(handled, message.Offset)
			if message.Offset == 21 {
				return errors.New("poison")
			}
			return nil
		},
		noWait,
	)

	require.NoError(t, err)
	require.Equal(t, []int64{21, 21, 22}, handled)
	require.Equal(t, []int64{21, 22}, reader.commits)
	require.Len(t, writer.messages, 1)
	require.Equal(t, []byte("poison"), writer.messages[0].Value)
	require.Contains(t, writer.messages[0].Headers, kafka.Header{Key: "X-Retry-Count", Value: []byte("2")})
}

func TestConsumeMessagesDoesNotCommitWhenDLQWriteFails(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	reader := &fakeConsumerReader{messages: []kafka.Message{
		{Topic: "events", Partition: 0, Offset: 31},
		{Topic: "events", Partition: 0, Offset: 32},
	}}
	writer := &fakeMessageWriter{err: errors.New("dlq unavailable")}

	err := consumeMessages(
		ctx,
		ConsumerConfig{MaxRetries: 1, DlqTopic: "events-dlq"},
		reader,
		writer,
		func(context.Context, kafka.Message) error { return errors.New("poison") },
		func(context.Context, time.Duration) error {
			cancel()
			return context.Canceled
		},
	)

	require.NoError(t, err)
	require.Equal(t, 1, reader.fetched)
	require.Empty(t, reader.commits)
	require.Len(t, writer.messages, 1)
}

func noWait(context.Context, time.Duration) error { return nil }

func offsetString(offset int64) string {
	if offset == 0 {
		return "0"
	}
	var buf [20]byte
	i := len(buf)
	for offset > 0 {
		i--
		buf[i] = byte(offset%10) + '0'
		offset /= 10
	}
	return string(buf[i:])
}
