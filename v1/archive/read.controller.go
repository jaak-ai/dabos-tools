package archive

import "io/ioutil"

func Read(file string) ([]byte, error) {
	return ioutil.ReadFile(file)
}
