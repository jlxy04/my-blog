package controllers

import (
	"encoding/json"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"log"
	"my-blog/service"
)

type MenuController struct {
	ctx     iris.Context
	Service service.MenuService
}

func (c *MenuController) GetAll() string {
	dataList := c.Service.GetAll()
	b, err := json.Marshal(dataList)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

//根据ID查询
func (c *MenuController) GetById(id string) mvc.Result {
	return nil
}
