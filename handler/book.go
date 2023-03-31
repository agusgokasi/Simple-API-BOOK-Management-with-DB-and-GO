package handler

import (
	"eight-learn/helper"
	"eight-learn/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h HttpServer) CreateBook(c *gin.Context) {
	in := model.Book{}

	err := c.BindJSON(&in)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	// call service
	res, err := h.app.CreateBook(in)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, res)
}

func (h HttpServer) GetBookById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		helper.BadRequest(c, "Invalid book ID")
		return
	}

	// call service
	res, err := h.app.GetBookById(id)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, res)
}

func (h HttpServer) GetBooks(c *gin.Context) {
	// call service
	res, err := h.app.GetBooks()
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, res)
}

func (h HttpServer) UpdateBook(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		helper.BadRequest(c, "Invalid book ID")
		return
	}

	in := model.Book{}
	err = c.BindJSON(&in)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	// call service
	res, err := h.app.UpdateBook(id, in)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, res)
}

func (h HttpServer) DeleteBook(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		helper.BadRequest(c, "Invalid book ID")
		return
	}

	// call service
	err = h.app.DeleteBook(id)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.OkWithMessage(c, "Book deleted successfully")
}
