package main

import (
	"LogAgent/kafka"
	taillog "LogAgent/tailLog"
	"log"
	"time"

	"gopkg.in/ini.v1"
)

func run(topic string) {
	//读取日志,发送到kafka
	for {
		select {
		case line := <-taillog.ReadChan():
			kafka.SendToKafka(topic, line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
}

//logagent入口
func main() {
	//0、加载配置文件
	cfg, err := ini.Load("./conf/config.ini")
	if err != nil {
		log.Fatal("Read config file failed,err:", err)
	}
	addr := cfg.Section("kafka").Key("address").String()
	path := cfg.Section("taillog").Key("path").String()
	topic := cfg.Section("kafka").Key("topic").String()
	//1、初始化kafka连接
	err = kafka.Init([]string{addr})
	if err != nil {
		log.Fatal(err)
	}
	//2、打开日志文件，收集日志
	err = taillog.Init(path)
	if err != nil {
		log.Fatal(err)
	}
	run(topic)
}
