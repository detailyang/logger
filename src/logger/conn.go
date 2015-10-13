/*
* @Author: detailyang
* @Date:   2015-10-10 15:01:22
* @Last Modified by:   detailyang
* @Last Modified time: 2015-10-13 12:11:03
 */

package logger

import (
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type Conn struct {
	conn  net.Conn
	Name  string
	Alive bool
	Stop  chan struct{}
}

func NewConn(conn net.Conn) *Conn {
	return &Conn{
		conn: conn,
		Stop: make(chan struct{}),
	}
}

func (self *Conn) Write(b []byte) (int, error) {
	return self.conn.Write(b)
}

func (self *Conn) Close() error {
	return self.conn.Close()
}

func (self *Conn) Connect() {
	var tmpConn net.Conn
	var err error

	urlSlice := strings.Split(self.Name, ":")
	switch urlSlice[0] {
	case "tcp":
		fallthrough
	case "udp":
		//neednt use tcp keepalive, we have proceess broken pipe and close it:)
		tmpConn, err = net.Dial(urlSlice[0], urlSlice[1][2:]+":"+urlSlice[2])
		if err != nil {
			log.Printf("[error] connect %s %s\r\n", self.Name, err.Error())
			return
		}
		if tmpConn, ok := tmpConn.(*net.TCPConn); ok {
			err = tmpConn.SetKeepAlive(true)
			if err != nil {
				log.Println("[error] tcp set keepalive ", err)
			}
			err = tmpConn.SetKeepAlivePeriod(60 * time.Second)
			if err != nil {
				log.Println("[error] tcp set keepalive Period", err)
			}
		}
	case "unix":
		localFile, err := os.OpenFile(urlSlice[1][2:], os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
		if err != nil {
			log.Printf("[error] connect %s %s\r\n", self.Name, err.Error())
			return
		}
		tmpConn = NewFileConn(localFile)
	default:
		break
	}
	if self.conn != nil {
		err = self.conn.Close()
		if err != nil {
			log.Printf("[error] close %s %s\r\n", self.Name, err.Error())
			return
		}
	}
	self.conn = tmpConn
	//should use mutex?
	self.Alive = true
}
