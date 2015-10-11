/*
* @Author: detailyang
* @Date:   2015-10-10 15:01:22
* @Last Modified by:   detailyang
* @Last Modified time: 2015-10-11 11:33:48
 */

package logger

import (
	"net"
	"os"
)

type FileConn struct {
	net.Conn
	file *os.File
}

func NewFileConn(file *os.File) *FileConn {
	return &FileConn{
		file: file,
	}
}

func (self *FileConn) Write(b []byte) (int, error) {
	return self.file.Write(b)
}

func (self *FileConn) Close() error {
	return self.file.Close()
}
