/*
* @Author: detailyang
* @Date:   2015-10-10 13:13:06
* @Last Modified by:   detailyang
* @Last Modified time: 2015-10-10 14:17:51
 */

package main

import (
	"logger"
)

func main() {
	l := logger.NewLogger("tcp://127.0.0.1:5140", "tcp://127.0.0.1:5141", "unix:///data/logs/tmp")
	l.Run()
}
