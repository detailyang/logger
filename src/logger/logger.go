/*
* @Author: detailyang
* @Date:   2015-10-10 12:25:53
* @Last Modified by:   detailyang
* @Last Modified time: 2015-10-11 20:59:56
 */

package logger

import (
	"bufio"
	"io"
	"log"
	"os"
)

type Logger struct {
	lineChannel chan []byte
	writterList *WritterList
	config      Config
}

func NewLogger(configfile, topic, localServer, remoteServer, localFile string) *Logger {
	config := NewConfig(configfile)
	if config == nil {
		config = &Config{
			LocalServer:  localServer,
			RemoteServer: remoteServer,
			LocalFile:    localFile,
			Topic:        topic,
		}
	}
    config.Topic = topic
	if err := config.isValid(); err == false {
		config = &Config{
			LocalServer:  localServer,
			RemoteServer: remoteServer,
			LocalFile:    localFile,
			Topic:        topic,
		}
	}
	return &Logger{
		lineChannel: make(chan []byte),
		writterList: NewWritterList([]string{config.LocalServer, config.RemoteServer, config.LocalFile}),
	}
}

func (self *Logger) print() {
	for line := range self.lineChannel {
		_, err := self.writterList.Write(line)
		if err != nil {
			log.Println(err)
		}
	}
}

func (self *Logger) Run() {
	reader := bufio.NewReader(os.Stdin)
	go self.print()
	for {
		line, isPrefix, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				continue
			}
			log.Fatal(err)
		}
		if isPrefix == false {
			line = append(line, '\n')
		}
		self.lineChannel <- line
	}
}
