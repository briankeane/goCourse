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

func (tc *tasksController) ShowAll(w http.ResponseWriter, r *http.Request) {
	tasks, err := models.Tasks.FindAll()
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
	}

	res, err := json.Marshal(tasks)
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}
	common.JsonOk(w, res, http.StatusOK)
}

func (tc *tasksController) Create(w http.ResponseWriter, r *http.Request) {
	var t models.Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
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
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	task, err := models.Tasks.FindOne(id)
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
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	err := models.Tasks.DeleteById(id)
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	common.JsonStatus(w, http.StatusOK)
}

func (tc *tasksController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var t models.Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	task, err := models.Tasks.Update(id, t.Name, t.Desc)
	res, err := json.Marshal(task)
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}
	common.JsonOk(w, res, http.StatusOK)
}
