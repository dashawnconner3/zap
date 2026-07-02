func TestSamplerDynamicLevelTransition(t *testing.T) {
	lvl := NewAtomicLevelAt(InfoLevel)
	observed, logs := observer.New(lvl)
	sampler := NewSamplerWithOptions(
		observed,
		time.Second,
		1,
		0,
	)
	logger := New(sampler)

	logger.Info("info 1")
	logger.Info("info 2")

	lvl.SetLevel(ErrorLevel)

	logger.Error("error 1")
	logger.Error("error 2")

	entries := logs.All()
	errorCount := 0
	for _, entry := range entries {
		if entry.Level == ErrorLevel {
			errorCount++
		}
	}

	if errorCount != 2 {
		t.Errorf("Expected 2 error logs, got %d. Error logs were incorrectly dropped.", errorCount)
	}
}