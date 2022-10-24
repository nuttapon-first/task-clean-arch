package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nuttapon-first/task-clean-arch/modules/entities"
)

type taskController struct {
	TaskUse entities.TaskUsecase
}

func NewTaskController(r *gin.RouterGroup, taskUse entities.TaskUsecase) {
	controllers := &taskController{
		TaskUse: taskUse,
	}
	r.POST("", controllers.NewTask)
	r.GET("", controllers.GetTaskList)
	r.PUT("/:id", controllers.UpdateTaskById)
	r.DELETE("/:id", controllers.DeleteTaskById)
}

func (h *taskController) NewTask(c *gin.Context) {
	req := &entities.NewTaskReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	res, err := h.TaskUse.NewTask(req)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"result":  res,
	})

}

func (h *taskController) GetTaskList(c *gin.Context) {
	req := &entities.GetTaskListReq{}

	res, err := h.TaskUse.GetTaskList(req)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  res,
	})
}

func (h *taskController) UpdateTaskById(c *gin.Context) {
	req := &entities.UpdateTaskListReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	req.ID = id

	res, err := h.TaskUse.UpdateTaskById(req)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  res,
	})
}

func (h *taskController) DeleteTaskById(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	req := &entities.DeleteTaskListReq{ID: id}

	err = h.TaskUse.DeleteTaskById(req)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
