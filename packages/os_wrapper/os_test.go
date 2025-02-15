package os_wrapper_test

type OsMockImpl struct {
	Env map[string]string
}

func (mock *OsMockImpl) Getenv(key string) string {
	return mock.Env[key]
}

func NewMock(mp map[string]string) OsMockImpl {
	return OsMockImpl{mp}
}
