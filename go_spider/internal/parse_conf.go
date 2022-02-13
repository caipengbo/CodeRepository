package internal

import (
	"fmt"
	log "github.com/alecthomas/log4go"
	"gopkg.in/gcfg.v1"
	"os"
)

type ConfigSection struct {
	UrlListFile     string
	OutputDirectory string
	MaxDepth        int
	CrawlInterval   int
	CrawlTimeout    int
	TargetUrl       string
	ThreadCount     int
}

type Config struct {
	Spider ConfigSection
}

func isExist(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func checkConfig(config *Config) error {
	path := config.Spider.UrlListFile
	if path == "" {
		return fmt.Errorf("invalid urlListFile, is empty path")
	}
	if !isExist(path) {
		return fmt.Errorf("invalid urlListFile, %s is not exist", path)
	}
	return nil
}

func ParseConf(confPath string) (*Config, error) {
	cfg := Config{
		Spider: ConfigSection{
			UrlListFile:     "./url.data",
			OutputDirectory: "./output",
			MaxDepth:        1,
			CrawlInterval:   1,
			CrawlTimeout:    1,
			TargetUrl:       "",
			ThreadCount:     4,
		}}
	err := gcfg.FatalOnly(gcfg.ReadFileInto(&cfg, confPath))
	if err != nil {
		log.Error("Parse file error: %s\n", err.Error())
		return nil, err
	}
	err = checkConfig(&cfg)
	if err != nil {
		log.Error("Parse file error: %s\n", err.Error())
		return nil, err
	}
	log.Info("Parse config file success")
	return &cfg, nil
}
