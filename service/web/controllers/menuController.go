package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/my-blog/models"
	"github.com/my-blog/service"
)

type MenuController struct {
	Service service.MenuService
}

func (c *MenuController) GetAll() []models.BlogMenu {
	return c.Service.GetAll()
}

//根据ID查询
func (c *MenuController) GetById(id string) mvc.Result {
	return nil
}

//根据ID查询
func (c *MenuController) PostCreate(ctx iris.Context) error {
	m := &models.BlogMenu{}
	ctx.ReadJSON(m)
	return nil
}
