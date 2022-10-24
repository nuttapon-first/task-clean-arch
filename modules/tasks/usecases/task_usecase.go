package usecases

import (
	"log"

	"github.com/nuttapon-first/task-clean-arch/modules/entities"
)

type taskUses struct {
	TaskRepo entities.TaskRepository
}

func NewTaskUsecase(taskRepo entities.TaskRepository) entities.TaskUsecase {
	return &taskUses{
		TaskRepo: taskRepo,
	}
}

func (u *taskUses) NewTask(req *entities.NewTaskReq) (*entities.NewTaskRes, error) {
	res, err := u.TaskRepo.NewTask(req)
	if err != nil {
		log.Printf("new task error: %#v", err.Error())
		return res, err
	}
	return res, nil
}

func (u *taskUses) GetTaskList(req *entities.GetTaskListReq) (*[]entities.GetTaskListRes, error) {
	res, err := u.TaskRepo.GetTaskList(req)
	if err != nil {
		log.Printf("get task error: %#v", err.Error())
		return res, err
	}
	return res, nil
}

func (u *taskUses) UpdateTaskById(req *entities.UpdateTaskListReq) (*entities.UpdateTaskListRes, error) {
	res, err := u.TaskRepo.UpdateTaskById(req)
	if err != nil {
		log.Printf("update task error: %#v", err.Error())
		return res, err
	}
	return res, nil
}

func (u *taskUses) DeleteTaskById(req *entities.DeleteTaskListReq) error {
	err := u.TaskRepo.DeleteTaskById(req)
	if err != nil {
		log.Printf("delete task error: %#v", err.Error())
		return err
	}
	return nil
}
