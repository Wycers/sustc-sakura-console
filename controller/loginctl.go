package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wycers/sustc-sakura-console/util"
	"github.com/wycers/sustc-sakura-console/model"
	"net/http"
	"encoding/json"
)

func LoginAction(c *gin.Context) {
	res := util.NewResult()

	login := &model.LoginRequest{}
	if err := c.BindJSON(login); err != nil {
		res.Code = -1
		res.Msg = "parses add article request failed"
		c.JSON(http.StatusBadRequest, res)
		return
	}


	login.Action = "login"
	JsonString, err := json.Marshal(login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	response := util.NewResult()
	exres := util.Exchange(string(JsonString))
	if exres == nil {
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	if err := json.Unmarshal(*exres, &response); err != nil {
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	if response.Code == -1 {
		res.Data = response.Msg
		c.JSON(http.StatusForbidden, res)
		return
	}
	if response.Code == 0 {
		session := util.GetSession(c)
		session.StudentID = login.Username
		session.JSESSIONID = response.Msg
		session.Save(c)
		res.Data = "success"
		c.JSON(http.StatusOK, res)
	}
}



func NewLoginRequest(username string, password string) *model.LoginRequest {
	return &model.LoginRequest{
		Action: "login",
		Username: username,
		Passwrod: password,
	}
}
