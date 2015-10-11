/*
* @Author: detailyang
* @Date:   2015-10-10 12:25:53
* @Last Modified by:   detailyang
* @Last Modified time: 2015-10-11 21:52:00
 */

package logger

import (
	"bufio"
	"io"
	"log"
	"log/syslog"
	"os"
)

type Logger struct {
	lineChannel chan []byte
	writterList *WritterList
	config      Config
}

func NewLogger(configfile, logFile, topic, localServer, remoteServer, localFile string) *Logger {
	config := NewConfig(configfile)
	if config == nil {
		config = &Config{
			LocalServer:  localServer,
			RemoteServer: remoteServer,
			LocalFile:    localFile,
			Topic:        topic,
			LogFile:      logFile,
		}
	}
	config.Topic = topic
	if err := config.isValid(); err == false {
		config = &Config{
			LocalServer:  localServer,
			RemoteServer: remoteServer,
			LocalFile:    localFile,
			Topic:        topic,
			LogFile:      logFile,
		}
	}

	//default set syslog
	w, err := syslog.New(syslog.LOG_INFO, "ylogger")
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(w)

	if config.LogFile != "syslog" {
		f, err := os.OpenFile(config.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Println("[error] log to logfile ", err)
		} else {
            defer f.Close()
            log.SetOutput(f)
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
