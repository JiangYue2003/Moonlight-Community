package userfeed

import (
	"context"
	"errors"
)

var errRefreshQueueFull = errors.New("user feed page refresh queue is full")

type refreshPromise struct {
	done    chan struct{}
	result  loadResult
	err     error
	started bool
}

type refreshJob struct {
	key     string
	loader  Loader
	promise *refreshPromise
}

func (c *pageCache) startRefreshWorkers() {
	for i := 0; i < c.cfg.RefreshWorkers; i++ {
		c.refreshWG.Add(1)
		go c.runRefreshWorker()
	}
}

func (c *pageCache) runRefreshWorker() {
	defer c.refreshWG.Done()
	for {
		// Give shutdown priority over an already-buffered job so Close does not
		// probabilistically drain the queue before workers exit.
		select {
		case <-c.ctx.Done():
			return
		default:
		}
		select {
		case <-c.ctx.Done():
			return
		case job := <-c.refreshQueue:
			if !c.beginRefresh(job) {
				c.completeRefresh(job, loadResult{}, context.Canceled)
				continue
			}
			c.refreshActive.Add(1)
			c.observeRefreshState()
			refreshCtx, cancel := context.WithTimeout(c.ctx, c.cfg.LoaderTimeout)
			result, err := c.loadAndStore(refreshCtx, job.key, job.loader)
			cancel()
			c.observe(OperationRefreshLoad, err)
			c.reportError(OperationRefreshLoad, err)
			c.refreshActive.Add(-1)
			c.completeRefresh(job, result, err)
		}
	}
}

func (c *pageCache) beginRefresh(job refreshJob) bool {
	c.loadMu.Lock()
	defer c.loadMu.Unlock()
	if c.closed || c.refreshPending[job.key] != job.promise || job.promise.started {
		return false
	}
	job.promise.started = true
	return true
}

// takeOverQueuedRefresh lets an expired foreground request bypass refresh
// queue head-of-line blocking. The worker and takeover race under loadMu: if
// the worker has started, the request waits for its bounded result; otherwise
// the queued job is canceled before any loader or write can begin.
func (c *pageCache) takeOverQueuedRefresh(key string, promise *refreshPromise) bool {
	c.loadMu.Lock()
	if c.refreshPending[key] != promise || promise.started {
		c.loadMu.Unlock()
		return false
	}
	promise.err = context.Canceled
	delete(c.refreshPending, key)
	c.refreshPendingCount.Add(-1)
	close(promise.done)
	c.loadMu.Unlock()
	c.observeRefreshState()
	return true
}

// scheduleRefresh reserves the full page key before enqueueing. A cold miss
// that crosses staleUntil while this job is queued or running waits on the
// same promise, so an older refresh cannot race and overwrite a newer load.
func (c *pageCache) scheduleRefresh(job refreshJob) bool {
	if job.loader == nil || c.ctx.Err() != nil {
		return false
	}
	c.loadMu.Lock()
	if c.closed || c.loads[job.key] != nil || c.refreshPending[job.key] != nil {
		c.loadMu.Unlock()
		return false
	}
	promise := &refreshPromise{done: make(chan struct{})}
	job.promise = promise
	c.refreshPending[job.key] = promise
	c.refreshPendingCount.Add(1)
	select {
	case c.refreshQueue <- job:
		c.loadMu.Unlock()
		c.observe(OperationRefreshEnqueue, nil)
		c.observeRefreshState()
		return true
	default:
		delete(c.refreshPending, job.key)
		c.refreshPendingCount.Add(-1)
		c.loadMu.Unlock()
		c.observe(OperationRefreshEnqueue, errRefreshQueueFull)
		c.observeRefreshState()
		c.reportError(OperationRefreshEnqueue, errRefreshQueueFull)
		return false
	}
}

func (c *pageCache) completeRefresh(job refreshJob, result loadResult, err error) {
	c.loadMu.Lock()
	if current := c.refreshPending[job.key]; current == job.promise {
		job.promise.result = result
		job.promise.err = err
		delete(c.refreshPending, job.key)
		c.refreshPendingCount.Add(-1)
		close(job.promise.done)
	}
	c.loadMu.Unlock()
	c.observeRefreshState()
}
