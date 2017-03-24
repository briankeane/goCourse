package models

import (
	"testing"

	"github.com/briankeane/goCourse/common"
	"gopkg.in/mgo.v2/bson"
)

func TestFindOne(t *testing.T) {
	common.StartUp()
	var new_model = &Task{
		ID:   bson.NewObjectId(),
		Name: "Test Task",
		Desc: "Test Desc",
	}

	if err := common.DB.Tasks.Insert(new_model); err != nil {
		t.Errorf("error writing to DB")
	}

	task, err := Tasks.FindOne(new_model.ID.Hex())
	if err != nil {
		t.Errorf("error: " + err.Error())
	}

	if new_model.ID != task.ID {
		t.Errorf("Expected task ID to equal %s, but got %s", new_model.ID.Hex(), task.ID.Hex())
	}
	t.Logf("task.ID: %s", task.ID.Hex())

	if new_model.Desc != task.Desc {
		t.Errorf("Expected Desc to equal %s, but got %s", new_model.Desc, task.Desc)
	}
	t.Logf("task.Name: %s", task.Name)

	if new_model.Name != task.Name {
		t.Errorf("Expected task.Name to equal %s, but got %s", new_model.Name, task.Name)
	}
	t.Logf("task.Desc: %s", task.Desc)
}
