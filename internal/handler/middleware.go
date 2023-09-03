package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ClientIdentity(c *gin.Context) {
	session, _ := h.CookieStore.Get(c.Request, "session")

	if session == nil {
		h.ErrorPageHandler(c, http.StatusUnauthorized, "Вы не авторизованы")
	}
}
