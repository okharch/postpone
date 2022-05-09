package postpone

import (
	"github.com/stretchr/testify/require"
	"sync"
	"testing"
	"time"
	"yadex/config"
)

func TestPPExecute(t *testing.T) {
	config.SetLogger()
	var ppe PostponeExecutor
	executed := "no"
	var l sync.Mutex
	f := func(name string) func() {
		return func() {
			l.Lock()
			executed=name
			l.Unlock()
		}
	}
	ppe.Postpone(f("s/2"), time.Millisecond*100)

	time.Sleep(time.Millisecond*60)
	l.Lock()
	require.Equal(t, "no", executed)
	l.Unlock()

	time.Sleep(time.Millisecond*60)
	l.Lock()
	require.Equal(t, "s/2", executed)
	l.Unlock()

	ppe.Cancel()
	ppe.Postpone(f("ms"), time.Millisecond*60)
	time.Sleep(time.Millisecond*100)
	l.Lock()
	require.Equal(t, "ms", executed)
	ppe.Cancel()
	l.Unlock()
}
