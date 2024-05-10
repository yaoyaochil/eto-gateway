package utils

import (
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"strings"
)

func ConvertToGBK(str string) string {
	reader := transform.NewReader(strings.NewReader(str), simplifiedchinese.GBK.NewEncoder())
	d, _ := ioutil.ReadAll(reader)
	return string(d)
}

func ConvertToUnicode(str string) string {
	reader := transform.NewReader(strings.NewReader(str), simplifiedchinese.GBK.NewDecoder())
	d, _ := ioutil.ReadAll(reader)
	return string(d)
}
