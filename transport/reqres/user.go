package reqres

import (
	"encoding/json"
	"net/http"

	"github.com/uchupx/golang-mongodb/helper"
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

func (h UserRequest) Users(w http.ResponseWriter, r *http.Request) {
	var res BodyRespone

	switch r.Method {
	case "GET":
		res = h.findAll(w, r)

	case "POST":
		res = h.insert(w, r)

	default:
		res = BodyRespone{
			StatusCode: http.StatusInternalServerError,
			Body:       "Sorry, only GET and POST methods are supported.",
		}
	}
	returnResponse(w, res)
	return
}
func (h UserRequest) findAll(w http.ResponseWriter, r *http.Request) BodyRespone {
	var response BodyRespone
	result, err := h.UserRepo.FindAll(r.Context())
	if err != nil {
		response = BodyRespone{
			StatusCode: http.StatusInternalServerError,
			Body:       err,
		}
	} else {
		response = BodyRespone{
			StatusCode: http.StatusOK,
			Body:       result,
		}
	}

	return response
}

func (h UserRequest) insert(w http.ResponseWriter, r *http.Request) BodyRespone {
	var response BodyRespone
	var userReq PostUserRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&userReq)

	if err != nil {
		er := ErrorMessage{
			StatusCode: http.StatusBadRequest,
			Error:      err.Error(),
		}
		response = BodyRespone{
			StatusCode: http.StatusInternalServerError,
			Body:       er,
		}

		return response
	}
	pwd, err := helper.HashPassword(*userReq.Password)

	if err != nil {
		er := ErrorMessage{
			StatusCode: http.StatusBadRequest,
			Error:      err.Error(),
		}

		response = BodyRespone{
			StatusCode: http.StatusInternalServerError,
			Body:       er,
		}

		return response
	}

	userSche := user.User{
		Username: *userReq.Username,
		Password: pwd,
		Email:    *userReq.Email,
		Nama:     *userReq.Nama,
	}

	err = h.UserRepo.Insert(r.Context(), userSche)
	if err != nil {
		er := ErrorMessage{
			StatusCode: http.StatusBadRequest,
			Error:      err.Error(),
		}
		response = BodyRespone{
			StatusCode: http.StatusInternalServerError,
			Body:       er,
		}

		return response
	}
	response = BodyRespone{
		StatusCode: http.StatusOK,
		Body:       "Succesfull insert user",
	}
	return response
}
