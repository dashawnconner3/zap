package zapcore

import (
	"sync/atomic"
	"time"
)

// ... existing code ...

func (s *sampler) Check(ent Entry, ce *CheckedEntry) *CheckedEntry {
	if !s.Enabled(ent.Level) {
		return ce
	}

	// ErrorLevel and higher should never be sampled.
	if ent.Level >= ErrorLevel {
		return ce.AddCore(ent, s)
	}

	// ... existing sampling logic ...
	return ce
}