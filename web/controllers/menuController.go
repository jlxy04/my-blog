package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"my-blog/service"
)

type MenuController struct {
	ctx     iris.Context
	Service service.MenuService
}

func (c *MenuController) GetAll() mvc.Result {
	return nil
}

//根据ID查询
func (c *MenuController) GetById(id string) mvc.Result {
	return nil
}
