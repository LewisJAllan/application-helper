package asynchronous

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAsyncWaiter_Run(t *testing.T) {
	w := &AsyncWaiter{}
	wait := 300 * time.Millisecond
	go func() {
		time.Sleep(time.Millisecond * 100)
		w.Run(func() {
			time.Sleep(wait - time.Millisecond*100)
		})
		_ = w.Stop(context.Background())
	}()
	n := time.Now()
	err := w.Start(context.Background())
	assert.NoError(t, err)
	assert.InDelta(t, wait, time.Since(n), float64(time.Millisecond*10))
}

func TestNewAsyncWaiter(t *testing.T) {
	w := NewAsyncWaiter()
	w.Add(1)
	wait := 300 * time.Millisecond
	go func() {
		_ = w.Stop(context.Background())
		time.Sleep(wait)
		w.Done()
	}()
	n := time.Now()
	_ = w.Start(context.Background())
	assert.InDelta(t, wait, time.Since(n), float64(time.Millisecond*10))
}
