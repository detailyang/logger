/*
* @Author: detailyang
* @Date:   2015-10-10 13:36:16
* @Last Modified by:   detailyang
* @Last Modified time: 2015-10-10 14:15:40
 */

package logger

import (
	"errors"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

type Writter struct {
	io []io.Writer
}

func NewWritter(localServerString, remoteServerString, localFileString string) *Writter {
	//tcp://127.0.0.1:5140
	ls := strings.Split(localServerString, ":")
	localServerConn, err := net.Dial(ls[0], ls[1][2:]+":"+ls[2])
	if err != nil {
		log.Println("[error] connect local server ", err)
	}

	rs := strings.Split(remoteServerString, ":")
	remoteServerConn, err := net.Dial(rs[0], rs[1][2:]+":"+rs[2])
	if err != nil {
		log.Println("[error] connect remote server ", err)
	}

	lf := strings.Split(localFileString, ":")
	localFileConn, err := os.OpenFile(lf[1][2:], os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		log.Println("[error] open local file ", err)
	}

	return &Writter{
		io: []io.Writer{localServerConn, remoteServerConn, localFileConn},
	}
}

func (self *Writter) Write(msg []byte) (n int, err error) {
	for _, io := range self.io {
		n, err = io.Write(msg)
		if err == nil {
			return n, err
		}
	}

	return 0, errors.New("cannot write any server")
}
