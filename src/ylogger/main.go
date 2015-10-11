/*
* @Author: detailyang
* @Date:   2015-10-10 13:13:06
* @Last Modified by:   detailyang
* @Last Modified time: 2015-10-11 21:48:10
 */

package main

import (
	"flag"
	"logger"
)

func main() {
	var config, topic, localServer, remoteServer, localFile, logFile string
	flag.StringVar(&config, "config", "", "config file")
	flag.StringVar(&topic, "topic", "logger", "syslog tag (topic)")
	flag.StringVar(&localServer, "localServer", "tcp://127.0.0.1:5140", "local server to log")
	flag.StringVar(&remoteServer, "remoteServer", "tcp://127.0.1:5140", "remote server to log")
	flag.StringVar(&localFile, "localFile", "unix:///data/logs/logger/failover.log", "local file to log")
	flag.StringVar(&logFile, "logFile", "syslog", "default file (syslog) to log")
	flag.Parse()
	l := logger.NewLogger(config, logFile, topic, localServer, remoteServer, localFile)
	l.Run()
}
