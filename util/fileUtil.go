package utils

import (
	"io/ioutil"
	"strings"
)

func FileInt(filePath string, separator string) (values []int) {
	return ToIntArray(FileString(filePath, separator))
}

func FileString(filePath string, separator string) (values []string) {
	bytes, _ := ioutil.ReadFile(filePath)

	str := string(bytes)

	return strings.Split(str, separator)
}
