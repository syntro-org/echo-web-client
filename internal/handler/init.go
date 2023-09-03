package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

const (
	SESSION_COOKIE_NAME = "session"
)

type Handler struct {
	BaseAPIUrl  string
	CookieStore sessions.CookieStore
}

func NewHandler(baseApiUrl, cookieSecrets string) *Handler {
	return &Handler{
		BaseAPIUrl:  baseApiUrl,
		CookieStore: *sessions.NewCookieStore([]byte(cookieSecrets)),
	}
}

func (h *Handler) Init() *gin.Engine {
	r := gin.New()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "static/")
	r.NoRoute(h.NotFoundPageHandler)

	r.GET("/", h.AuthPageHandler)

	identityRequired := r.Group("/me", h.ClientIdentity)
	{
		identityRequired.GET("/chat")
	}

	r.POST("/reg", h.RegHandler)
	r.GET("/auth", h.AuthHandler)

	return r
}
