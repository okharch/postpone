package postpone

import (
	"context"
	"sync"
	"time"
)

// implements object which can execute postponed operation
type PostponeExecutor struct {
	sync.Mutex
	cancel context.CancelFunc
}
func (ppe *PostponeExecutor) Postpone(f func(),timeout time.Duration) {
	go func() {
		ppe.Lock()
		if ppe.cancel != nil {
			ppe.cancel()
		}
		var ctx context.Context
		ctx, ppe.cancel = context.WithCancel(context.TODO())
		ppe.Unlock()
		select {
		case <-ctx.Done():
		case <-time.After(timeout):
			f()
		}
		ppe.Cancel()
	}()
}

func (ppe *PostponeExecutor) Cancel() {
	ppe.Lock()
	if ppe.cancel != nil {
		ppe.cancel()
	}
	ppe.cancel = nil
	ppe.Unlock()
}