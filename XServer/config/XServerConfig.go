package config

import (
	"encoding/json"
	"github.com/ddjjxx123/go/x-server/server"
	"io/ioutil"
)

type XServerObject struct {
	Server      server.IXServer //全局Server对象
	Host        string          //当前服务器主机IP
	Port        int             //当前服务器主机监听端口号
	HostName    string          //当前服务器名称
	Version     string          //当前版本号
	MaxDataSize int             //数据包的最大值
	MaxConn     int             //当前服务器主机允许的最大链接个数
}

var GlobalServerObject *XServerObject

func (x *XServerObject) Reload() {
	//读取自定义配置
	config, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		panic(err)
		return
	}
	//写入
	err = json.Unmarshal(config, &GlobalServerObject)
	if err != nil {
		panic(err)
		return
	}
}
func Init() {
	//默认配置
	GlobalServerObject = &XServerObject{
		HostName:    "X-Server",
		Version:     "V1.0",
		Port:        8888,
		Host:        "127.0.0.1",
		MaxConn:     12000,
		MaxDataSize: 4096,
	}
	GlobalServerObject.Reload()
}
