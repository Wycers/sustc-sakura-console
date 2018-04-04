package controller

import (
	"github.com/wycers/sustc-sakura-console/model"
	"github.com/gin-gonic/gin"
	"github.com/wycers/sustc-sakura-console/util"
	"encoding/json"
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
)

func DownloadAction(c *gin.Context)  {
	res := util.NewResult()
	defer c.JSON(http.StatusOK, res)
	session := util.GetSession(c)
	if session.JSESSIONID == "" {
		res.Code = -1
		res.Msg = "login first"
		return
	}

	request := &model.DownloadRequest{}
	if err := c.BindJSON(request); err != nil {
		res.Code = -5
		res.Msg = "parse failed"
		return
	}
	week := request.Week
	if week == 0 {
		res.Code = -2
		res.Msg = "week error"
		return
	}

	trans := NewTransRequest(session.JSESSIONID, week)
	JsonString, err := json.Marshal(trans)
	if err != nil {
		return
	}

	response := util.NewResult()

	exres := util.Exchange(string(JsonString))
	if exres == nil {
		return
	}
	if err := json.Unmarshal(*exres, &response); err != nil {
		return
	}
	if response.Code == -1 {
		res.Msg = response.Msg
		return
	}
	if response.Code == -3 {
		res.Msg = response.Msg
		return
	}
	if response.Code == 0 {
		res.Msg = fmt.Sprintf("%s-%d.ics", session.StudentID, week)
		data, err := ReadAll(response.Msg)
		if err != nil {
			res.Code = -6
			return
		}
		content := string(data[:])
		res.Data = content
	}
}
/*
c.Header("x-suggested-filename", fmt.Sprintf("%s-%d.ics",session.StudentID, week))

		data, err := ReadAll(response.Msg)
		if err != nil {
			res.Code = -6
			return
		}
		content := string(data[:])
		res.Data = content

*/
func ReadAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}


	return ioutil.ReadAll(f)
}


func NewTransRequest(jsessionid string, week int) *model.TransRequest {
	return &model.TransRequest{
		Action: "trans",
		JSESSIONID: jsessionid,
		Week: week,
	}
}