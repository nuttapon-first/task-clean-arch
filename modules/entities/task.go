package entities

type NewTaskReq struct {
	Name string `json:"name" db:"name"  binding:"required"`
}

type NewTaskRes struct {
	ID     int    `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	Status int    `json:"status" db:"status"`
}

type GetTaskListReq struct{}

type GetTaskListRes struct {
	ID     int    `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	Status int    `json:"status" db:"status"`
}

type UpdateTaskListReq struct {
	ID     int    `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	Status int    `json:"status" db:"status"`
}

type UpdateTaskListRes struct {
	ID     int    `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	Status int    `json:"status" db:"status"`
}

type DeleteTaskListReq struct {
	ID int `json:"id" db:"id"`
}

type TaskUsecase interface {
	NewTask(req *NewTaskReq) (*NewTaskRes, error)
	GetTaskList(req *GetTaskListReq) (*[]GetTaskListRes, error)
	UpdateTaskById(req *UpdateTaskListReq) (*UpdateTaskListRes, error)
	DeleteTaskById(req *DeleteTaskListReq) error
}

type TaskRepository interface {
	NewTask(req *NewTaskReq) (*NewTaskRes, error)
	GetTaskList(req *GetTaskListReq) (*[]GetTaskListRes, error)
	UpdateTaskById(req *UpdateTaskListReq) (*UpdateTaskListRes, error)
	DeleteTaskById(req *DeleteTaskListReq) error
}
