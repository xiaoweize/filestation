package example

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"gopkg.in/yaml.v2"
)

type conf struct {
	Endpoint        string `yaml:"endpoint"`
	AccessKeyID     string `yaml:"accessKeyID"`
	AccessKeySecret string `yaml:"accessKeySecret"`
}

var Setting = conf{}

var client *oss.Client

//测试阿里云oss Sdk Bucket接口
func TestListBuckets(t *testing.T) {
	bucket, err := client.Bucket("test-delicloud")
	if err != nil {
		t.Error(err)
	}
	if err := bucket.PutObjectFromFile("oss_test.go", "oss_test.go"); err != nil {
		t.Error(err)
	}
}

func init() {
	var err error
	client, err = oss.New(Setting.Endpoint, Setting.AccessKeyID, Setting.AccessKeySecret, oss.Timeout(5, 10))
	if err != nil {
		log.Fatal(err)
	}
	file, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := yaml.Unmarshal(file, &Setting); err != nil {
		log.Fatal(err)
	}
}
