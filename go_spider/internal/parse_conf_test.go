package internal

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

const configTemplate = `
[spider]
# 种子文件路径(默认 ./url.data)
urlListFile = %s
# 抓取结果存储目录(默认 ./output)
outputDirectory = ./output
# 最大抓取深度(种子为0级)
maxDepth = 1
# 抓取间隔. 单位: 秒
crawlInterval =  1
# 抓取超时. 单位: 秒
crawlTimeout = 1
# 需要存储的目标网页URL pattern(正则表达式)
targetUrl = .*.(htm|html)$
# 抓取routine数(默认 4)
threadCount = 8
`

func createTempConfigFile(urlListFile string) (string, error) {
	confString := fmt.Sprintf(configTemplate, urlListFile)
	file, err := ioutil.TempFile("", "test_parse_conf_")
	if err != nil {
		return "", err
	}
	if _, err = file.Write([]byte(confString)); err != nil {
		return "", err
	}
	if err = file.Close(); err != nil {
		return "", err
	}
	return file.Name(), nil
}

func TestParseConf(t *testing.T) {
	tmpUrlFile, err := ioutil.TempFile("", "url_data_file")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpUrlFile.Name())

	tmpConfFile1, err := createTempConfigFile("")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpConfFile1)

	tmpConfFile2, err := createTempConfigFile(tmpUrlFile.Name())
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpConfFile2)

	type args struct {
		confPath string
	}
	cfg := Config{
		Spider: ConfigSection{
			UrlListFile:     tmpUrlFile.Name(),
			OutputDirectory: "./output",
			MaxDepth:        1,
			CrawlInterval:   1,
			CrawlTimeout:    1,
			TargetUrl:       ".*.(htm|html)$",
			ThreadCount:     8,
		}}
	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
		{"ReadFileError", args{confPath: "./err.conf"}, nil, true},
		{"CheckConfigError", args{confPath: tmpConfFile1}, nil, true},
		{"ParseSuccess", args{confPath: tmpConfFile2}, &cfg, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseConf(tt.args.confPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseConf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseConf() got = %v, want %v", got, tt.want)
			}
		})
	}
}
