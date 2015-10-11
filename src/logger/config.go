/*
* @Author: detailyang
* @Date:   2015-10-10 15:01:22
* @Last Modified by:   detailyang
* @Last Modified time: 2015-10-11 21:48:36
 */

package logger

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"regexp"
)

type Config struct {
	LocalServer  string `json:"local_server"`
	RemoteServer string `json:"remote_server"`
	LocalFile    string `json:"local_file"`
	LogFile      string `json:"log_file"`
	Topic        string
}

func NewConfig(file string) *Config {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Println("[error] open config file ", err)
		return nil
	}

	var config Config
	if err := json.Unmarshal(content, &config); err != nil {
		log.Println("[error] unmarshl config file ", err)
		return nil
	}
	return &config
}

func (self *Config) isValid() bool {
	pattern := regexp.MustCompile(`^(tcp|udp)://.*?:\d+$`)
	filePattern := regexp.MustCompile(`^unix://.*$`)
	if pattern.MatchString(self.LocalServer) == false {
		log.Println("[error] config local server is not the pattern ^(tcp|udp)://.*?:\\d+$")
		return false
	}
	if pattern.MatchString(self.RemoteServer) == false {
		log.Println("[error] config remote server is not the pattern ^(tcp|udp)://.*?:\\d+$")
		return false
	}
	if filePattern.MatchString(self.LocalFile) == false {
		log.Println("[error] config local file is not the pattern ^unix://.*$")
		return false
	}

	return true
}
