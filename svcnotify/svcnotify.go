package svcnotify

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// FTPConfig ...
type FTPConfig struct {
	Host    string
	Port    int
	Account string
	Pwd     string
	DocPath string
}

// EmailConfig ...
type EmailConfig struct {
	Host    string
	Port    int
	Account string
	Pwd     string
}

// Config FTP Email 設定
type Config struct {
	FTP   FTPConfig
	Email EmailConfig
}

// SvcNotify ...
type SvcNotify struct {
	Config Config
}

// LoadConfig 讀取設定
func (svc *SvcNotify) LoadConfig(filePath string) {
	var file, _ = os.OpenFile(filePath, os.O_RDONLY, 0755)
	b, _ := ioutil.ReadAll(file)

	json.Unmarshal(b, &svc.Config)

}
