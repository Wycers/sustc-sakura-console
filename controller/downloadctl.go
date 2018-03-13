package controller

import (
	"github.com/wycers/sustc-sakura-console/model"
	"github.com/gin-gonic/gin"
	"github.com/wycers/sustc-sakura-console/util"
	"encoding/json"
	"net/http"
)

func DownloadAction(c *gin.Context)  {
	res := util.NewResult()

	session := util.GetSession(c)
	if session.JSESSIONID == "" {
		res.Code = -1
		res.Msg = "login first"
		c.JSON(http.StatusForbidden, res)
		return
	}

	trans := NewTransRequest(session.JSESSIONID)
	JsonString, err := json.Marshal(trans)
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
		c.Header("content-disposition", "attachment; filename=" + session.StudentID + ".ics")
		c.File(response.Msg)
		res.Data = "success"
		c.JSON(http.StatusOK, res)
	}
}


func NewTransRequest(jsessionid string) *model.TransRequest {
	return &model.TransRequest{
		Action: "trans",
		JSESSIONID: jsessionid,
	}
}