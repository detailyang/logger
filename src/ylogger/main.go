/*
* @Author: detailyang
* @Date:   2015-10-10 13:13:06
* @Last Modified by:   detailyang
* @Last Modified time: 2015-10-12 21:55:16
 */

package main

import (
	"flag"
	"log"
	"logger"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
)

var Buildstamp, Githash string

func main() {
	var pport, config, topic, localServer, remoteServer, localFile, logFile string
	var version, pprof bool

	flag.StringVar(&config, "config", "", "config file")
	flag.StringVar(&topic, "topic", "logger", "syslog tag (topic)")
	flag.StringVar(&localServer, "localServer", "tcp://127.0.0.1:5140", "local server to log")
	flag.StringVar(&remoteServer, "remoteServer", "tcp://127.0.1:5140", "remote server to log")
	flag.StringVar(&localFile, "localFile", "unix:///data/logs/logger/failover.log", "local file to log")
	flag.StringVar(&logFile, "logFile", "syslog", "default file (syslog) to log")
	flag.StringVar(&pport, "pport", "6666", "pprof port")
	flag.BoolVar(&version, "version", false, "version")
	flag.BoolVar(&pprof, "pprof", false, "pprof for logger")
	flag.Parse()

	if version == true {
		log.Println("Git Commit Hash: ", Githash)
		log.Println("UTC Build Time : ", Buildstamp)
		return
	}

	if pprof == true {
		go func() {
			log.Println(http.ListenAndServe("0.0.0.0:"+pport, nil))
		}()
	}

	l := logger.NewLogger(config, logFile, topic, localServer, remoteServer, localFile)

	sign := make(chan os.Signal, 1)
	signal.Notify(sign, syscall.SIGTERM)
	go func() {
		sig := <-sign
		log.Println("[info] receive signal ", sig)
		l.Stop()
		os.Exit(0)
	}()

	l.Run()
}
