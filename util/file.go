package util

import (
	"io/ioutil"
)

// WriteFile 写入文件
func WriteFile(path string, text string) error {
    err := ioutil.WriteFile(path, []byte(text), 0644)
    return err
}
