/*
* @Author: detailyang
* @Date:   2015-10-10 13:36:16
* @Last Modified by:   detailyang
* @Last Modified time: 2015-10-10 18:07:43
 */

package logger

import (
	"errors"
	// "io"
	"log"
	"net"
	"os"
	"strings"
)

type WritterList struct {
	Resources []*Conn
}

func NewWritterList(urls []string) *WritterList {
	var name string
	var alive bool
	var conn *Conn
	var tmpConn net.Conn
	var err error

	wl := &WritterList{
		Resources: make([]*Conn, 0),
	}

	for _, url := range urls {
		name = url
		alive = true
		urlSlice := strings.Split(url, ":")
		switch urlSlice[0] {
		case "tcp":
			fallthrough
		case "udp":
			tmpConn, err = net.Dial(urlSlice[0], urlSlice[1][2:]+":"+urlSlice[2])
			if err != nil {
				alive = false
				log.Println("[error] connect local server ", err)
			}
		case "unix":
			localFile, err := os.OpenFile(urlSlice[1][2:], os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
			if err != nil {
				alive = false
				log.Println("[error] open local file ", err)
			}
			defer localFile.Close()
			tmpConn, err = net.FileConn(localFile)
			if err != nil {
				alive = false
				log.Println("[error] copy local file conn ", err)
			}
		default:
			continue
		}
		conn = NewConn(tmpConn)
		conn.Name = name
		conn.Alive = alive
		wl.Resources = append(wl.Resources, conn)
	}

	return wl
}

func (self *WritterList) Write(msg []byte) (n int, err error) {
	for _, resource := range self.Resources {
		n, err = resource.Write(msg)
		log.Println(resource.Name)
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
