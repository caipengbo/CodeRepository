package internal

import "testing"

func Test_crawl(t *testing.T) {
	cfg := Config{
		Spider: ConfigSection{
			MaxDepth:    2,
			ThreadCount: 8,
		}}
	url := Url{"https://ssr1.scrape.center/", 0}
	var urls []Url
	urls = append(urls, url)
	type args struct {
		urls   []Url
		config *Config
	}
	tests := []struct {
		name string
		args args
	}{
		{"Case1", args{urls, &cfg}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			crawl(tt.args.urls, tt.args.config)
		})
	}
}
