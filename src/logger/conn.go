/*
* @Author: detailyang
* @Date:   2015-10-10 15:01:22
* @Last Modified by:   detailyang
* @Last Modified time: 2015-10-11 10:41:33
 */

package logger

import (
	"net"
	"os"
)

type Conn struct {
	conn  net.Conn
	Name  string
	Alive bool
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

// type Conn interface {
// 	// Read reads data from the connection.
// 	// Read can be made to time out and return a Error with Timeout() == true
// 	// after a fixed time limit; see SetDeadline and SetReadDeadline.
// 	Read(b []byte) (n int, err error)

// 	// Write writes data to the connection.
// 	// Write can be made to time out and return a Error with Timeout() == true
// 	// after a fixed time limit; see SetDeadline and SetWriteDeadline.
// 	Write(b []byte) (n int, err error)

// 	// Close closes the connection.
// 	// Any blocked Read or Write operations will be unblocked and return errors.
// 	Close() error

// 	// LocalAddr returns the local network address.
// 	LocalAddr() Addr

// 	// RemoteAddr returns the remote network address.
// 	RemoteAddr() Addr

// 	// SetDeadline sets the read and write deadlines associated
// 	// with the connection. It is equivalent to calling both
// 	// SetReadDeadline and SetWriteDeadline.
// 	//
// 	// A deadline is an absolute time after which I/O operations
// 	// fail with a timeout (see type Error) instead of
// 	// blocking. The deadline applies to all future I/O, not just
// 	// the immediately following call to Read or Write.
// 	//
// 	// An idle timeout can be implemented by repeatedly extending
// 	// the deadline after successful Read or Write calls.
// 	//
// 	// A zero value for t means I/O operations will not time out.
// 	SetDeadline(t time.Time) error

// 	// SetReadDeadline sets the deadline for future Read calls.
// 	// A zero value for t means Read will not time out.
// 	SetReadDeadline(t time.Time) error

// 	// SetWriteDeadline sets the deadline for future Write calls.
// 	// Even if write times out, it may return n > 0, indicating that
// 	// some of the data was successfully written.
// 	// A zero value for t means Write will not time out.
// 	SetWriteDeadline(t time.Time) error
// }
