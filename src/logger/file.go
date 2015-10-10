/*
* @Author: detailyang
* @Date:   2015-10-10 15:06:23
* @Last Modified by:   detailyang
* @Last Modified time: 2015-10-10 17:17:10
 */

package logger

import (
	"os"
)

type File struct {
	file *os.File
	Writter
}

func NewFile(file *os.File) *File {
	return &File{
		file: file,
	}
}

func (self *File) Write(b []byte) (int, error) {
	return self.file.Write(b)
}

func (self *File) Close() error {
	return self.file.Close()
}
