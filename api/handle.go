package api

import (
	"encoding/json"
	"go-hello/models"
	"go-hello/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

func (s *Server) handleGetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.EncodeJSON(w, http.StatusOK, utils.H{
			"message": "valid id",
		})
	}
	user, err := s.store.GetUserByID(int64(id))
	if err != nil {
		utils.EncodeJSON(w, http.StatusOK, utils.H{
			"message": "获取失败,无效的number",
		})
		return
	}
	err = utils.EncodeJSON(w, http.StatusOK, utils.H{
		"message": "获取成功",
		"data":    user,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Server) handleUserRegister(w http.ResponseWriter, r *http.Request) {
	userRegister := &models.UserRegister{}
	err := json.NewDecoder(r.Body).Decode(userRegister)
	if err != nil {
		log.Fatal(err)
	}
	validate := validator.New()
	err = validate.Struct(userRegister)
	if err != nil {
		utils.EncodeJSON(w, http.StatusOK, utils.H{
			"Error": utils.ValidatorError(err),
		})
		return
	}
	effrow, err := s.store.GetUserByNameAndEmail(userRegister.Username, userRegister.Emial)
	if err != nil {
		log.Fatal(err)
	}
	if effrow != 0 {
		utils.EncodeJSON(w, http.StatusOK, utils.H{
			"message": "用户已经注册,请勿重复注册",
		})
		return
	}
	err = s.store.UserRegister(userRegister)
	if err != nil {
		log.Fatal(err)
	}
	utils.EncodeJSON(w, http.StatusOK, utils.H{
		"message": "注册成功!",
	})

}

func (s *Server) handleUserSignup(w http.ResponseWriter, r *http.Request) {
	userSign := &models.UserSign{}
	err := json.NewDecoder(r.Body).Decode(userSign)
	if err != nil {
		log.Fatal(err)
	}
	rowAffected := s.store.UserSignup(userSign.Username, userSign.Password)
	if rowAffected == 0 {
		utils.EncodeJSON(w, http.StatusOK, utils.H{
			"message": "用户名或密码错误",
		})
		return
	}
	token, err := GenJWT()
	if err != nil {
		log.Fatal(err)
	}
	utils.EncodeJSON(w, http.StatusOK, utils.H{
		"message": "登录成功",
		"token":   token,
	})
}
