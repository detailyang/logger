/*
* @Author: detailyang
* @Date:   2015-10-10 12:25:53
* @Last Modified by:   detailyang
* @Last Modified time: 2015-10-10 14:17:20
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
	writter     *Writter
}

func NewLogger(localServerString, remoteServerString, localFileString string) *Logger {
	return &Logger{
		lineChannel: make(chan []byte),
		writter:     NewWritter(localServerString, remoteServerString, localFileString),
	}
}

func (self *Logger) print() {
	for line := range self.lineChannel {
        _, err := self.writter.Write(line)
        if err != nil {
            log.Println(err)
        }
		log.Println(line)
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
