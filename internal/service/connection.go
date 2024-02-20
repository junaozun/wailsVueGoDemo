package service

import (
	"changeme/internal/define"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"os"
)

func ConnectionList() ([]*define.Connection, error) {
	curPath, _ := os.Getwd()
	data, err := os.ReadFile(curPath + string(os.PathSeparator) + define.ConfigName)
	if errors.Is(err, os.ErrNotExist) {
		return nil, err
	}
	conf := new(define.Config)
	err = json.Unmarshal(data, conf)
	if err != nil {
		return nil, err
	}
	return conf.Connections, nil
}

func ConnectionCreate(conn *define.Connection) (err error) {
	if conn.Addr == "" {
		err = errors.New("链接地址不能为空")
		return
	}
	if conn.Name == "" {
		conn.Name = conn.Addr
	}
	if conn.Port == "" {
		conn.Port = "6379"
	}
	conn.Identity = uuid.New().String()
	conf := new(define.Config)
	curPath, _ := os.Getwd()
	data, err := os.ReadFile(curPath + string(os.PathSeparator) + define.ConfigName)
	if errors.Is(err, os.ErrNotExist) {
		// 配置文件内容初始化
		conf.Connections = []*define.Connection{conn}
		data, _ = json.Marshal(conf)
		// 创建写入配置文件
		os.MkdirAll(curPath, 0666)
		os.WriteFile(curPath+string(os.PathSeparator)+define.ConfigName, data, 0666)
		return nil
	}
	// 配置文件中追加
	json.Unmarshal(data, conf)
	conf.Connections = append(conf.Connections, conn)
	data, _ = json.Marshal(conf)
	os.WriteFile(curPath+string(os.PathSeparator)+define.ConfigName, data, 0666)
	return nil
}
