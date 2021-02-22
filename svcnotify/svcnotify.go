package svcnotify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/jlaffaye/ftp"
	_ "github.com/tealeg/xlsx/v3"
	_ "gopkg.in/gomail.v2"
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
	Config  Config
	FileBuf []byte
}

// LoadConfig 讀取設定
func (svc *SvcNotify) LoadConfig(filePath string) {
	var file, _ = os.OpenFile(filePath, os.O_RDONLY, 0755)
	b, _ := ioutil.ReadAll(file)

	json.Unmarshal(b, &svc.Config)

}

// FTPController ...
func (svc *SvcNotify) FTPController() {
	c, err := ftp.Dial(fmt.Sprintf("%v:%v", svc.Config.FTP.Host, svc.Config.FTP.Port), ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	err = c.Login(svc.Config.FTP.Account, svc.Config.FTP.Pwd)
	if err != nil {
		log.Fatal(err)
	}

	r, err := c.Retr(svc.Config.FTP.DocPath)
	if err != nil {
		panic(err)
	}
	defer r.Close()

	svc.FileBuf, err = ioutil.ReadAll(r)
}

// TriggerExec ...
func TriggerExec() {
	svc := SvcNotify{}
	svc.LoadConfig("config.json")
	svc.FTPController()

}
