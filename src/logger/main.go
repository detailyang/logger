/*
* @Author: detailyang
* @Date:   2015-10-10 10:27:18
* @Last Modified by:   detailyang
* @Last Modified time: 2015-10-10 10:34:56
 */

package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Println("get error ", err)
			break
		}
		log.Println(line)
	}
}
