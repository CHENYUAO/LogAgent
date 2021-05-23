package taillog

import (
	"github.com/hpcloud/tail"
)

//日志文件收集模块

var (
	tails   *tail.Tail
	logchan chan string
)

func Init(fileName string) (err error) {
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}
	tails, err = tail.TailFile(fileName, config)
	if err != nil {
		return err
	}
	return nil
}

func ReadChan() <-chan *tail.Line {
	return tails.Lines
}
