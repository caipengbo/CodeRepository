package internal

import (
	"fmt"
	"golang.org/x/net/html/charset"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func Test_determineEncoding(t *testing.T) {
	type args struct {
		r io.Reader
	}
	path := "./html/gbk_test.html"
	file, err := os.Open(path)
	t.Error(err)
	r, err := charset.NewReader(file, "UTF-8")

	if err != nil {
		fmt.Printf("Err: %s", err.Error())
	}
	all, err := ioutil.ReadAll(r)
	if err != nil {
		fmt.Printf("Err: %s", err.Error())
	}
	fmt.Print(string(all))
}

func TestEncoding(t *testing.T) {

}

func TestParseHtml(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
	}{
		{"Case1", args{url: "https://ssr1.scrape.center/"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parse(tt.args.url)
		})
	}
}
