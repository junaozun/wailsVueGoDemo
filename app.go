package main

import (
	"changeme/internal/define"
	"changeme/internal/service"
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) ConnectionList() define.H {
	connList, err := service.ConnectionList()
	if err != nil {
		return define.M{
			"code": -1,
			"msg":  "ERROR:" + err.Error(),
		}
	}
	return define.M{
		"code": 200,
		"data": connList,
	}
}

func (a *App) ConnectionCreate(connection *define.Connection) define.H {
	err := service.ConnectionCreate(connection)
	if err != nil {
		return define.M{
			"code": -1,
			"msg":  "ERROR:" + err.Error(),
		}
	}
	return define.M{
		"code": 200,
		"msg":  "新建成功",
	}
}
