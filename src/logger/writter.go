/*
* @Author: detailyang
* @Date:   2015-10-10 13:36:16
* @Last Modified by:   detailyang
* @Last Modified time: 2015-10-10 17:21:13
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
	Name  string
	Alive bool
}

type WritterList struct {
	Resources []io.Writer
}

func NewWritterList(localServerString, remoteServerString, localFileString string) *WritterList {
	var name string
	var alive bool

	wl := &WritterList{
		Resources: make([]io.Writer, 0),
	}

	name = "local server"
	alive = true
	ls := strings.Split(localServerString, ":")
	localServerConn, err := net.Dial(ls[0], ls[1][2:]+":"+ls[2])
	if err != nil {
		alive = false
		log.Println("[error] connect local server ", err)
	}
	defer localServerConn.Close()
	conn := NewConn(localServerConn)
	conn.Name = name
	conn.Alive = alive
	wl.Resources = append(wl.Resources, conn)

	name = "local server"
	alive = true
	rs := strings.Split(remoteServerString, ":")
	remoteServerConn, err := net.Dial(rs[0], rs[1][2:]+":"+rs[2])
	if err != nil {
		alive = false
		log.Println("[error] connect remote server ", err)
	}
	conn = NewConn(remoteServerConn)
	conn.Name = name
	conn.Alive = alive
	wl.Resources = append(wl.Resources, conn)

	name = "local file"
	alive = true
	lf := strings.Split(localFileString, ":")
	localFileConn, err := os.OpenFile(lf[1][2:], os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		alive = false
		log.Println("[error] open local file ", err)
	}
	file := NewFile(localFileConn)
	file.Name = name
	file.Alive = alive
	wl.Resources = append(wl.Resources, file)

	return wl
}

func (self *WritterList) Write(msg []byte) (n int, err error) {
	for _, resource := range self.Resources {
		n, err = resource.Write(msg)
		if err != nil {
			continue
		}
		//write empty message to detect broken pipe
		n, err = resource.Write([]byte{})
		if err == nil {
			return n, err
		}
		log.Println(err)
	}

	return 0, errors.New("cannot write any server")
}
