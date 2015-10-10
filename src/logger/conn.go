/*
* @Author: detailyang
* @Date:   2015-10-10 15:01:22
* @Last Modified by:   detailyang
* @Last Modified time: 2015-10-10 17:21:04
 */

package logger

import (
	"net"
)

type Conn struct {
	conn net.Conn
	Writter
}

func NewConn(conn net.Conn) *Conn {
	return &Conn{
		conn: conn,
	}
}

func (self *Conn) Write(b []byte) (int, error) {
	return self.conn.Write(b)
}

func (self *Conn) Close() error {
	return self.conn.Close()
}