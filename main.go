package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"gopkg.in/yaml.v2"
)

type conf struct {
	Endpoint        string `yaml:"endpoint"`
	AccessKeyID     string `yaml:"accessKeyID"`
	AccessKeySecret string `yaml:"accessKeySecret"`
	OssBucket       string `yaml:"OssBucket"`
	OssPath         string `yaml:"OssPath"`
	OssDomain       string `yaml:"OssDomain"`
}

var Setting = conf{}

var client *oss.Client

func UploadFile(filePath, fileName string) (string, error) {
	bucket, err := client.Bucket(Setting.OssBucket)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return Setting.OssDomain + filepath.Join(Setting.OssPath, fileName), bucket.PutObjectFromFile(filepath.Join(Setting.OssPath, fileName), filePath)
}

func main() {
	var filePath, fileUrl string
	flag.StringVar(&filePath, "f", "", "Specify the file to upload")
	flag.Parse()
	flag.Usage = func() {
		flag.PrintDefaults()
	}
	if flag.NFlag() != 1 {
		flag.Usage()
		return
	}
	if fileInfo, err := os.Stat(filePath); err != nil {
		fmt.Println(err)
		return
	} else if fileInfo.IsDir() {
		log.Fatal("can't upload directory!")
	} else {
		absFilepath, err := filepath.Abs(filePath)
		if err != nil {
			log.Fatal(err)
		}
		fileName := filepath.Base(absFilepath)
		fileUrl, err = UploadFile(absFilepath, fileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	fmt.Printf("Upload success!Download Link->%s\n", fileUrl)
}

func init() {
	var err error
	file, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	if err := yaml.Unmarshal(file, &Setting); err != nil {
		log.Fatal(err)
	}
	client, err = oss.New(Setting.Endpoint, Setting.AccessKeyID, Setting.AccessKeySecret, oss.Timeout(5, 10))
	if err != nil {
		log.Fatal(err)
	}

}
