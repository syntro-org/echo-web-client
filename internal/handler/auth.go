package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/syntro-org/echo-web-client/internal/model"
	"github.com/syntro-org/echo-web-client/pkg/util"
)

func (h *Handler) RegHandler(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		h.ErrorPageHandler(c, http.StatusBadRequest, err.Error())
		return
	}

	regClient := model.RegClientDTO{
		Login:      c.Request.PostFormValue("login"),
		UniqueName: c.Request.PostFormValue("unique_name"),
		Nickname:   c.Request.PostFormValue("nickname"),
		Hash:       c.Request.PostFormValue("hash"),
	}

	jsonBytes, err := json.Marshal(regClient)
	if err != nil {
		h.ErrorPageHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	reqSender := util.RequestSender{
		Url:         "http://localhost:8888/api/client/create",
		Method:      util.POST_METHOD,
		ContentType: util.JSON_CONTENT_TYPE,
		Content:     jsonBytes,
	}

	resp, result, err := reqSender.Send()
	if err != nil {
		h.ErrorPageHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	if resp.StatusCode != http.StatusOK {
		h.ErrorPageHandler(c, http.StatusBadRequest, result["err"])
		return
	}

	c.Redirect(http.StatusSeeOther, "/")
}

func (h *Handler) AuthHandler(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		h.ErrorPageHandler(c, http.StatusBadRequest, err.Error())
		return
	}

	authClient := model.AuthClientDTO{
		Login: c.Request.FormValue("login"),
		Hash:  c.Request.FormValue("hash"),
	}

	jsonBytes, err := json.Marshal(authClient)
	if err != nil {
		h.ErrorPageHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	reqSender := util.RequestSender{
		Url:         "http://localhost:8888/api/client/log-in",
		Method:      util.GET_METHOD,
		ContentType: util.JSON_CONTENT_TYPE,
		Content:     jsonBytes,
	}

	resp, result, err := reqSender.Send()
	if err != nil {
		h.ErrorPageHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	if resp.StatusCode != http.StatusOK {
		h.ErrorPageHandler(c, http.StatusBadRequest, result["err"])
		return
	}

	session, _ := h.CookieStore.Get(c.Request, SESSION_COOKIE_NAME)
	session.Values["jwt"] = result["msg"]
	session.Save(c.Request, c.Writer)

	c.Redirect(http.StatusSeeOther, "/me/chat")
}
