package logging

type DummyWriter struct{}

func (dw *DummyWriter) Write(p []byte) (int, error) {
	return len(p), nil
}

func init() {

	LogSerive = logger{
		createSlogLogger(&DummyWriter{}),
	}
}
