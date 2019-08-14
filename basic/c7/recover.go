package main

import (
	"bytes"
	"fmt"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
)

// cd $GOPATH/src
// git clone --depth 1 https://github.com/golang/text.git golang.org/x/text

type charset struct {
	encoding.Encoding
	name string
}

var charsets = []charset{
	{charmap.Windows1252, "Windows1252"},
	{simplifiedchinese.GB18030, "GB18030"},
	{traditionalchinese.Big5, "Big5"},
	{unicode.UTF8, "UTF8"},
}

func convert(slice []byte, encoder *encoding.Encoder, decoder *encoding.Decoder) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(slice), encoder)
	slice2, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	reader2 := transform.NewReader(bytes.NewReader(slice2), decoder)
	slice3, err := ioutil.ReadAll(reader2)
	if err != nil {
		return nil, err
	}
	return slice3, nil
}

func recover(str string) {
	utf8bytes := []byte(str)
	for i := 0; i < len(charsets); i++ {
		for j := 0; j < len(charsets); j++ {
			if i != j {
				fmt.Printf("---- 原来编码(A)假设是: %s, 被错误解读为了(B): %s\n", charsets[j].name, charsets[i].name)
				out, err := convert(utf8bytes, charsets[i].NewEncoder(), charsets[j].NewDecoder())
				if err != nil {
					fmt.Errorf("error: %v\n", err)
					continue
				}
				fmt.Println(string(out))
				fmt.Println()
			}
		}
	}
}

func main() {
	recover("ÀÏÂí")
}
