package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AuthPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "auth.html", nil)
}

func (h *Handler) NotFoundPageHandler(c *gin.Context) {

}
