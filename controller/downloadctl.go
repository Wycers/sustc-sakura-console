package controller

import (
	"github.com/wycers/sustc-sakura-console/model"
	"github.com/gin-gonic/gin"
	"github.com/wycers/sustc-sakura-console/util"
	"encoding/json"
	"net/http"
	"fmt"
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

	request := &model.DownloadRequest{}
	if err := c.BindJSON(request); err != nil {
		res.Code = -5
		res.Msg = "parse failed"
		c.JSON(http.StatusBadRequest, res)
		return
	}
	week := request.Week
	if week == 0 {
		res.Code = -2
		res.Msg = "week error"
		c.JSON(http.StatusForbidden, res)
		return
	}

	trans := NewTransRequest(session.JSESSIONID, week)
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
		res.Msg = response.Msg
		c.JSON(http.StatusForbidden, res)
		return
	}
	if response.Code == -3 {
		res.Msg = response.Msg
		c.JSON(http.StatusUnauthorized, res)
		return
	}
	if response.Code == 0 {
		c.Header("content-disposition", "attachment; filename=" + fmt.Sprintf("%s-%d.ics",session.StudentID, week))
		c.Header("x-suggested-filename", fmt.Sprintf("%s-%d.ics",session.StudentID, week))
		c.File(response.Msg)
		res.Data = "success"
		c.JSON(http.StatusOK, res)
	}
}


func NewTransRequest(jsessionid string, week int) *model.TransRequest {
	return &model.TransRequest{
		Action: "trans",
		JSESSIONID: jsessionid,
		Week: week,
	}
}