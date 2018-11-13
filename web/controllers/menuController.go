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

func (c *MenuController) GetAll() {
	dataList := c.Service.GetAll()
	b, err := json.Marshal(dataList)
	if b != nil && err != nil {
		log.Fatal(err)
	}
	c.ctx.JSON(Xx{})
}

type Xx struct {
	id string
}

//根据ID查询
func (c *MenuController) GetById(id string) mvc.Result {
	return nil
}

//根据ID查询
func (c *MenuController) PostCreate(id string) error {
	return nil
}
