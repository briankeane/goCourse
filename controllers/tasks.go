package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/briankeane/goCourse/common"
	"github.com/briankeane/goCourse/models"
	"github.com/gorilla/mux"
)

var Tasks = new(tasksController)

type tasksController struct{}

func (tc *tasksController) Create(w http.ResponseWriter, r *http.Request) {
	var t models.Task
	if err := json.NewDecoder(r.body).Decode(&t); err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	task, err := models.Tasks.Create(t.Name, t.Desc)
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(task)
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}
	common.JsonOk(w, res, http.StatusOK)
}

func (tc *tasksController) Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var t models.Task
	if err := json.NewDecoder(r.body).Decode(&t); err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	task, err := models.Tasks.Get(t.id)
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(task)
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}
	common.JsonOk(w, res, http.StatusOK)
}

func (tc *tasksController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var t models.Task
	if err := json.NewDecoder(r.body).Decode(&t); err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	task, err := models.Tasks.Delete(t.id)
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(task)
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}
	common.JsonOk(w, res, http.StatusOK)
}

func (tc *tasksController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var t models.Task
	if err := json.NewDecoder(r.body).Decode(&t); err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	task, err := models.Tasks.Update(t.id)
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(task)
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}
	common.JsonOk(w, res, http.StatusOK)
}
