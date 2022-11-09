package archive

import "os"

func Open(file string) (*os.File, error) {
	return os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
}
