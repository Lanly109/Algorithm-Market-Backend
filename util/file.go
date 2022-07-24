package util

import (
	"io/ioutil"
)

// WriteFile 写入文件
func WriteFile(path string, text string) error {
    err := ioutil.WriteFile(path, []byte(text), 0644)
    return err
}

// ReadFile 读取文件
func ReadFile(path string) (string, error) {
    text, err := ioutil.ReadFile(path)
    return string(text), err
}
