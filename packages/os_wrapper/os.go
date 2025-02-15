package os_wrapper

import "os"

type IOs interface {
	Getenv(string) string
}

type osServiceImpl struct{}

func (osServiceImpl) Getenv(key string) string {
	return os.Getenv(key)
}

var OsService = osServiceImpl{}
