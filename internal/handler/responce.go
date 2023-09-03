package handler

import "github.com/gin-gonic/gin"

func (h *Handler) ErrorPageHandler(c *gin.Context, httpCode int, msg interface{}) {
	c.HTML(httpCode, "error.html", msg)
}
