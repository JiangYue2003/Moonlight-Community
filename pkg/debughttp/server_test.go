package debughttp

import (
	"context"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	runtimepprof "runtime/pprof"
	"strings"
	"testing"
	"time"
)

func TestHandlerServesPprofEndpoints(t *testing.T) {
	server := httptest.NewServer(NewHandler())
	defer server.Close()

	for _, path := range []string{
		"/debug/pprof/",
		"/debug/pprof/goroutine?debug=1",
	} {
		resp, err := http.Get(server.URL + path)
		if err != nil {
			t.Fatalf("GET %s: %v", path, err)
		}
		body, readErr := io.ReadAll(resp.Body)
		resp.Body.Close()
		if readErr != nil {
			t.Fatalf("read %s: %v", path, readErr)
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("GET %s status=%d body=%s", path, resp.StatusCode, body)
		}
	}
}

func TestHandlerDoesNotUseDefaultServeMux(t *testing.T) {
	if NewHandler() == http.DefaultServeMux {
		t.Fatal("debug HTTP handler must not serve the process-wide default mux")
	}
}

func TestDisabledServerReturnsWithoutListening(t *testing.T) {
	err := New(Config{Enabled: false, ListenOn: "not-a-listen-address"}).Run(context.Background())
	if err != nil {
		t.Fatalf("disabled server returned error: %v", err)
	}
}

func TestServerReturnsListenError(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("reserve port: %v", err)
	}
	defer listener.Close()

	err = New(Config{Enabled: true, ListenOn: listener.Addr().String()}).Run(context.Background())
	if err == nil || !strings.Contains(err.Error(), "listen") {
		t.Fatalf("Run error=%v, want listen error", err)
	}
}

func TestServerStopsAndReleasesListenerWhenContextIsCanceled(t *testing.T) {
	reserved, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("reserve port: %v", err)
	}
	address := reserved.Addr().String()
	if err := reserved.Close(); err != nil {
		t.Fatalf("release reserved port: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan error, 1)
	go func() {
		done <- New(Config{Enabled: true, ListenOn: address}).Run(ctx)
	}()

	deadline := time.Now().Add(2 * time.Second)
	for {
		conn, dialErr := net.DialTimeout("tcp", address, 50*time.Millisecond)
		if dialErr == nil {
			conn.Close()
			break
		}
		if time.Now().After(deadline) {
			t.Fatalf("debug HTTP server did not listen on %s: %v", address, dialErr)
		}
		time.Sleep(10 * time.Millisecond)
	}

	cancel()
	select {
	case err := <-done:
		if err != nil {
			t.Fatalf("Run returned after cancellation: %v", err)
		}
	case <-time.After(3 * time.Second):
		t.Fatal("Run did not return after context cancellation")
	}

	reused, err := net.Listen("tcp", address)
	if err != nil {
		t.Fatalf("listener was not released after shutdown: %v", err)
	}
	reused.Close()
}

func TestServerCancelsActiveCPUProfileWhenContextIsCanceled(t *testing.T) {
	reserved, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("reserve port: %v", err)
	}
	address := reserved.Addr().String()
	if err := reserved.Close(); err != nil {
		t.Fatalf("release reserved port: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	serverDone := make(chan error, 1)
	go func() {
		serverDone <- New(Config{Enabled: true, ListenOn: address}).Run(ctx)
	}()
	waitForListener(t, address)

	profileDone := make(chan struct{})
	go func() {
		defer close(profileDone)
		resp, requestErr := http.Get("http://" + address + "/debug/pprof/profile?seconds=30")
		if requestErr == nil {
			io.Copy(io.Discard, resp.Body)
			resp.Body.Close()
		}
	}()
	waitForCPUProfile(t)

	cancel()
	select {
	case err := <-serverDone:
		if err != nil {
			t.Fatalf("Run returned after canceling an active profile: %v", err)
		}
	case <-time.After(time.Second):
		t.Fatal("Run did not stop an active profile promptly after context cancellation")
	}
	select {
	case <-profileDone:
	case <-time.After(time.Second):
		t.Fatal("active profile request did not stop after context cancellation")
	}
}

func waitForListener(t *testing.T, address string) {
	t.Helper()
	deadline := time.Now().Add(2 * time.Second)
	for {
		conn, err := net.DialTimeout("tcp", address, 50*time.Millisecond)
		if err == nil {
			conn.Close()
			return
		}
		if time.Now().After(deadline) {
			t.Fatalf("debug HTTP server did not listen on %s: %v", address, err)
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func waitForCPUProfile(t *testing.T) {
	t.Helper()
	deadline := time.Now().Add(2 * time.Second)
	for {
		if err := runtimepprof.StartCPUProfile(io.Discard); err != nil {
			return
		}
		runtimepprof.StopCPUProfile()
		if time.Now().After(deadline) {
			t.Fatal("CPU profile request did not start")
		}
		time.Sleep(10 * time.Millisecond)
	}
}
