/*
* @Author: detailyang
* @Date:   2015-10-10 10:27:18
* @Last Modified by:   detailyang
* @Last Modified time: 2015-10-10 11:26:04
 */

package main

import (
	//"bufio"
    "log/syslog"
	"log"
	"os"
    "syscall"
)

func main() {
    w, err := syslog.New(syslog.LOG_INFO, "ylogger")
    if err != nil {
        log.Fatal(err)
    }
    log.SetOutput(w)

    log.Println("start up")
	f, err := os.OpenFile("/data/logs/tmp", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
    

    buf := make([]byte, 65536)
	for {
        rb, err := syscall.Read(0, buf)
        if err != nil {
            log.Println(err)
        }
		f.Write(buf[:rb])
	}
}
