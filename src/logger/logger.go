/*
* @Author: detailyang
* @Date:   2015-10-10 12:25:53
* @Last Modified by:   detailyang
* @Last Modified time: 2015-10-14 11:32:50
 */

package logger

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"log/syslog"
	"os"
	"strconv"
)

type Logger struct {
	lineChannel chan []byte
	writterList *WritterList
	config      *Config
	header      []byte
	hostname    string
	pid         string
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
		f, err := os.OpenFile(config.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
		if err != nil {
			log.Println("[error] log to logfile ", err)
		} else {
			log.SetOutput(f)
		}
	}

	hostname, err := os.Hostname()
	if err != nil {
		log.Println("[error] get hostname ", err)
	}

	return &Logger{
		lineChannel: make(chan []byte),
		writterList: NewWritterList([]string{config.LocalServer, config.RemoteServer, config.LocalFile}),
		config:      config,
		hostname:    hostname,
		pid:         strconv.Itoa(os.Getpid()),
		header:      []byte("<30>1993-06-01 00:00:00 " + hostname + " " + config.Topic + " - - - - "),
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
	truncate := false
	//4096 is a page cache and stdin buffer size
	reader := bufio.NewReader(os.Stdin)
	go self.print()
	for {
		line, isPrefix, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				continue
			}
			log.Println("[error] read line ", err)
		}
		if truncate == false {
			var msg bytes.Buffer
			msg.Write(self.header)
			msg.Write(line)
			if isPrefix == false {
				msg.Write([]byte{'\n'})
			} else {
				truncate = true
			}
			self.lineChannel <- msg.Bytes()
		} else {
			if isPrefix == false {
				self.lineChannel <- append(line, '\n')
				truncate = false
			} else {
				self.lineChannel <- append(line)
			}
		}

	}
}

func (self *Logger) Stop() {
	self.writterList.Close()
}
