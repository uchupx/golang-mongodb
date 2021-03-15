package reqres

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uchupx/golang-mongodb/model/user"
)

type UserRequest struct {
	UserRepo user.UserRepo
}

type PostUserRequest struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
	Email    *string `json:"email"`
	Nama     *string `json:"nama"`
}

func (r UserRequest) FindAll(c *gin.Context) {
	result, err := r.UserRepo.FindAll(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, result)
	return
}

func (r UserRequest) Insert(c *gin.Context) {
	var userReq PostUserRequest

	err := c.ShouldBind(&userReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	userSche := user.User{
		Username: *userReq.Username,
		Password: *userReq.Password,
		Email:    *userReq.Email,
		Nama:     *userReq.Nama,
	}

	err = r.UserRepo.Insert(c, userSche)
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
