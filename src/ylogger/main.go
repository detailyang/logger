/*
* @Author: detailyang
* @Date:   2015-10-10 13:13:06
* @Last Modified by:   detailyang
* @Last Modified time: 2015-10-11 21:24:17
 */

package main

import (
	"flag"
	"log"
	"log/syslog"
	"logger"
)

func main() {
	var config, topic, localServer, remoteServer, localFile string
	flag.StringVar(&config, "config", "", "config file")
	flag.StringVar(&topic, "topic", "logger", "syslog tag (topic)")
	flag.StringVar(&localServer, "localServer", "tcp://127.0.0.1:5140", "local server to log")
	flag.StringVar(&remoteServer, "remoteServer", "tcp://127.0.1:5140", "remote server to log")
	flag.StringVar(&localFile, "localFile", "unix:///data/logs/logger/failover.log", "local file to log")
	flag.Parse()

	log.SetOutput(os.Stdout)
	w, err := syslog.New(syslog.LOG_INFO, "ylogger")
	if err != nil {
		log.Println("[error] new syslog error ", err)
	}
	log.SetOutput(w)

	l := logger.NewLogger(config, topic, localServer, remoteServer, localFile)
	l.Run()
}
