package server

import (
	"github.com/gin-gonic/gin"
	_taskController "github.com/nuttapon-first/task-clean-arch/modules/tasks/controllers"
	_taskRepository "github.com/nuttapon-first/task-clean-arch/modules/tasks/repositories"
	_taskUseCase "github.com/nuttapon-first/task-clean-arch/modules/tasks/usecases"
)

func (s *Server) MapHandlers() error {

	taskPath := s.App.Group("/tasks")
	taskRepository := _taskRepository.NewTaskRepository(s.Db)
	taskUseCase := _taskUseCase.NewTaskUsecase(taskRepository)
	_taskController.NewTaskController(taskPath, taskUseCase)

	s.App.Use(func(c *gin.Context) {
		c.AbortWithStatusJSON(404, gin.H{
			"message": "end point not found",
		})
	})

	return nil
}
