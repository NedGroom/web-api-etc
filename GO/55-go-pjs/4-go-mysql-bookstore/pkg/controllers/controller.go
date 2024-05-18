package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/NedGroom/4-go-mysql-bookstore/pkg/utils"
	"github.com/NedGroom/4-go-mysql-bookstore/pkg/models"
	"github.com/NedGroom/4-go-mysql-bookstore/pkg/controllers"

)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request){
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById