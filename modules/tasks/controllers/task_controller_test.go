package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nuttapon-first/task-clean-arch/modules/entities"
)

type fakeTaskUse struct{}

func (f *fakeTaskUse) NewTask(req *entities.NewTaskReq) (*entities.NewTaskRes, error) {
	newTask := &entities.NewTaskRes{
		ID:     1,
		Name:   "Test",
		Status: 0,
	}
	return newTask, nil
}

func (f *fakeTaskUse) GetTaskList(req *entities.GetTaskListReq) (*[]entities.GetTaskListRes, error) {
	taskList := &[]entities.GetTaskListRes{
		{
			ID:     2,
			Name:   "Test 2",
			Status: 0,
		},
	}
	return taskList, nil
}
func (f *fakeTaskUse) UpdateTaskById(req *entities.UpdateTaskListReq) (*entities.UpdateTaskListRes, error) {
	editTask := &entities.UpdateTaskListRes{
		ID:     3,
		Name:   "Test 3",
		Status: 1,
	}
	return editTask, nil
}
func (f *fakeTaskUse) DeleteTaskById(req *entities.DeleteTaskListReq) error {
	return nil
}

type newTaskResponse struct {
	Result entities.NewTaskRes `json:"result"`
}

func TestNewTask(t *testing.T) {
	fakeTaskUseCase := &fakeTaskUse{}
	fakeTaskController := &taskController{TaskUse: fakeTaskUseCase}

	body := map[string]interface{}{
		"name": "Test",
	}
	jsonStr, _ := json.Marshal(body)

	res := httptest.NewRecorder()
	c, r := gin.CreateTestContext(res)
	c.Request = httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(jsonStr))
	c.Request.Header.Set("Content-Type", "application/json; charset=utf-8")

	r.POST("/tasks", fakeTaskController.NewTask) // Call to a handler method
	r.ServeHTTP(res, c.Request)

	if status := res.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	response := &newTaskResponse{}
	json.Unmarshal(res.Body.Bytes(), response)
	want := "Test"
	if got := response.Result; got.Name != want {
		t.Errorf("Message: got %v want %v", got.Name, want)
	}
}

type getTaskResponse struct {
	Result []entities.GetTaskListRes `json:"result"`
}

func TestGetTaskList(t *testing.T) {
	fakeTaskUseCase := &fakeTaskUse{}
	fakeTaskController := &taskController{TaskUse: fakeTaskUseCase}

	res := httptest.NewRecorder()
	c, r := gin.CreateTestContext(res)
	c.Request = httptest.NewRequest(http.MethodGet, "/tasks", nil)
	c.Request.Header.Set("Content-Type", "application/json; charset=utf-8")

	r.GET("/tasks", fakeTaskController.GetTaskList) // Call to a handler method
	r.ServeHTTP(res, c.Request)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	response := &getTaskResponse{}
	json.Unmarshal(res.Body.Bytes(), response)
	want := map[string]interface{}{
		"id":     2,
		"name":   "Test 2",
		"status": 0,
	}
	if got := response.Result[0]; got.Name != want["name"] || got.Status != want["status"] {
		t.Errorf("Message:\n got %#v\n want %#v", got, want)
	}
}

type updateTaskResponse struct {
	Result entities.UpdateTaskListRes `json:"result"`
}

func TestEditTask(t *testing.T) {
	fakeTaskUseCase := &fakeTaskUse{}
	fakeTaskController := &taskController{TaskUse: fakeTaskUseCase}

	body := map[string]interface{}{
		"id":     3,
		"name":   "Test 3",
		"status": 1,
	}
	jsonStr, _ := json.Marshal(body)

	res := httptest.NewRecorder()
	c, r := gin.CreateTestContext(res)
	c.Request = httptest.NewRequest(http.MethodPut, "/tasks/3", bytes.NewBuffer(jsonStr))
	c.Request.Header.Set("Content-Type", "application/json; charset=utf-8")

	r.PUT("/tasks/:id", fakeTaskController.UpdateTaskById) // Call to a handler method
	r.ServeHTTP(res, c.Request)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	response := &updateTaskResponse{}
	json.Unmarshal(res.Body.Bytes(), response)

	if got := response.Result; got.Name != body["name"] || got.Status != body["status"] {
		t.Errorf("Message:\n got %#v\n want %#v", got, body)
	}
}

func TestDeleteTask(t *testing.T) {
	fakeTaskUseCase := &fakeTaskUse{}
	fakeTaskController := &taskController{TaskUse: fakeTaskUseCase}

	res := httptest.NewRecorder()
	c, r := gin.CreateTestContext(res)
	c.Request = httptest.NewRequest(http.MethodPut, "/tasks/4", nil)
	c.Request.Header.Set("Content-Type", "application/json; charset=utf-8")

	r.PUT("/tasks/:id", fakeTaskController.DeleteTaskById) // Call to a handler method
	r.ServeHTTP(res, c.Request)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
