/*
* @Author: detailyang
* @Date:   2015-10-10 10:27:18
* @Last Modified by:   detailyang
* @Last Modified time: 2015-10-10 11:26:04
 */

package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("./tmp", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	reader := bufio.NewReader(os.Stdin)
	for {
		line, isPrefix, err := reader.ReadLine()
		//TODO isprefix
		_ = isPrefix
		if err != nil {
			log.Println("get error ", err)
			break
		}
		log.Println(line)
		f.Write(line)
		f.Write([]byte{'\n'})
	}
}
