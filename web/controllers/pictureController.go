package controllers

import (
	"my-blog/models"
	"my-blog/service"
)

type PictureController struct {
	Server service.PictureService
}

func (c PictureController) GetTop() []models.BlogPicture {
	return c.Server.GetTop(6)
}
