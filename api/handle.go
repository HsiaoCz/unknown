package api

import (
	"go-hello/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *Server) handleGetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.EncodeJosn(w, http.StatusOK, utils.H{
			"message": "valid id",
		})
	}
	user := s.store.GetUserByID(id)
	err = utils.EncodeJosn(w, http.StatusOK, utils.H{
		"message": "获取成功",
		"data":    user,
	})
	if err != nil {
		log.Fatal(err)
	}
}
