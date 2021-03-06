/*
* @Author: detailyang
* @Date:   2015-10-10 13:36:16
* @Last Modified by:   detailyang
* @Last Modified time: 2015-10-13 12:10:51
 */

package logger

import (
	"errors"
	// "io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type WritterList struct {
	Resources []*Conn
}

func NewWritterList(urls []string) *WritterList {
	var name string
	var alive bool
	var err error
	var tmpConn net.Conn

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
				log.Println("[error] connect ", url, " ", err)
			}
			if urlSlice[0] == "tcp" {
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
			}
		case "unix":
			localFile, err := os.OpenFile(urlSlice[1][2:], os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
			if err != nil {
				alive = false
				log.Println("[error] open local file ", err)
			}
			tmpConn = NewFileConn(localFile)
		default:
			continue
		}
		//for closure
		conn := NewConn(tmpConn)
		conn.Name = name
		conn.Alive = alive
		wl.Resources = append(wl.Resources, conn)
		go func(conn *Conn) {
			ticker := time.NewTicker(1 * time.Second)
			for {
				select {
				case <-ticker.C:
					if conn.Alive == false {
						conn.Connect()
					}
				case <-conn.Stop:
					ticker.Stop()
					return
				}
			}
		}(conn)
	}

	return wl
}

func (self *WritterList) Write(msg []byte) (n int, err error) {
	for _, resource := range self.Resources {
		if resource.Alive == false {
			continue
		}
		n, err = resource.Write(msg)
		if err != nil {
			resource.Alive = false
			log.Println("[error] write msg error", err)
			continue
		}
		//write empty message to detect broken pipe
		n, err = resource.Write([]byte{})
		if err == nil {
			return n, err
		}

		resource.Alive = false
		log.Println("[error] write empty msg error", err)
	}

	return 0, errors.New("cannot write any server")
}

func (self *WritterList) Close() {
	var done struct{}
	for _, resource := range self.Resources {
		resource.Stop <- done
		resource.Close()
	}
}
