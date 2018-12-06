package controllers

import (
	"github.com/my-blog/models"
	"github.com/my-blog/service"
)

type PictureController struct {
	Server service.PictureService
}

func (c PictureController) GetTop() []models.BlogPicture {
	return c.Server.GetTop(6)
}
