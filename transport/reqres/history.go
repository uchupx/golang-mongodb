package reqres

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uchupx/golang-mongodb/model/history"
)

type HistoryRequest struct {
	HistoryRepo history.HistoryRepo
}

type PostHistoryRequest struct {
	Deskripsi  string  `json:"deskripsi"`
	Jumlah     uint64  `json:"jumlah"`
	Keterangan *string `json:"keterangan"`
	UserId     string  `json:"user_id"`
}

func (r HistoryRequest) Insert(c *gin.Context) {
	var historyReq PostHistoryRequest

	err := c.ShouldBind(&historyReq)
	if err != nil {
		er := ErrorMessage{
			StatusCode: http.StatusBadRequest,
			Error:      err.Error(),
		}
		c.JSON(http.StatusBadRequest, er)
		return
	}

	historySche := history.History{
		Deskripsi:  historyReq.Deskripsi,
		Jumlah:     historyReq.Jumlah,
		Keterangan: historyReq.Keterangan,
		UserId:     historyReq.UserId,
	}

	err = r.HistoryRepo.Insert(c, historySche)
	if err != nil {
		er := ErrorMessage{
			StatusCode: http.StatusBadRequest,
			Error:      err.Error(),
		}
		c.JSON(http.StatusBadRequest, er)
		return
	}

	c.JSON(http.StatusCreated, "ok")
}
