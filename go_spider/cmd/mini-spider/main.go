package main

import (
	"flag"
	"fmt"
	log "github.com/alecthomas/log4go"
	"minispider/internal"
)

const VERSION = "1.0.0"

func main() {
	// 解析命令行参数
	//-h(帮助)、-v(版本)、-c(配置文件路径）、-l（日志文件路径，2个日志：mini_spider.log和mini_spider.wf.log)
	var confPath = flag.String("c", "./spider.conf", "config file path")
	var logPath = flag.String("l", "./", "log file path")
	var version = flag.String("v", VERSION, "version")
	flag.Parse()
	fmt.Printf("config: %s, log: %s version: %s\n", *confPath, *logPath, *version)
	// TODO 配置 log conf
	// 解析配置文件
	config, err := internal.ParseConf(*confPath)
	if err != nil {
		log.Error("Parse config file failed: %s", err.Error())
		return
	}
	fmt.Printf("urlListFile: %s, threadCount: %d", config.Spider.UrlListFile, config.Spider.ThreadCount)
	// 启动爬虫

	// 接收退出信号

	// 检测超时，如果长时间没有新的任务添加，将退出

}
