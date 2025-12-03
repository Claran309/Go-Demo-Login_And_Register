package handlers

import (
	"GoGin/api/services"
	"GoGin/internal/model"
	"GoGin/internal/util"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	TodoService *services.TodoService
}

func NewTodoHandler(todoService *services.TodoService) *TodoHandler {
	return &TodoHandler{TodoService: todoService}
}

// Create 新增事项
func (h *TodoHandler) Create(c *gin.Context) {
	//捕获数据
	userID, _ := c.Get("user_id")
	var req model.CreateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.Error(c, 400, err.Error())
	}

	// 调用服务层
	todoTask, err := h.TodoService.CreateTodoTask(req, userID.(int))
	if err != nil {
		util.Error(c, 500, err.Error())
		return
	}

	// 返回响应
	util.Success(c, gin.H{
		"todo_task": todoTask,
	}, "todoTask created")
}

// Finish 完成事项
func (h *TodoHandler) Finish(c *gin.Context) {
	//捕获数据
	var req model.FinishTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.Error(c, 400, err.Error())
	}

	//调用服务层
	err := h.TodoService.FinishTodoTask(req.TodoID)
	if err != nil {
		util.Error(c, 500, err.Error())
		return
	}

	//返回响应
	util.Success(c, gin.H{
		"task_id": req.TodoID,
		"status":  "completed",
	}, "task finished")
}

// Delete 删除事项
func (h *TodoHandler) Delete(c *gin.Context) {
	//捕获数据
	var req model.DeleteTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.Error(c, 400, err.Error())
	}

	//调用服务层
	err := h.TodoService.DeleteTodoTask(req.TodoID)
	if err != nil {
		util.Error(c, 500, err.Error())
		return
	}

	//返回响应
	util.Success(c, gin.H{
		"task_id": req.TodoID,
	}, "task deleted")
}

// Info 获取事项列表
func (h *TodoHandler) Info(c *gin.Context) {
	//捕获数据
	userID, _ := c.Get("user_id")

	//调用服务层
	todos, dones, err := h.TodoService.GetInfo(userID.(int))
	if err != nil {
		util.Error(c, 500, err.Error())
		return
	}

	//返回响应
	util.Success(c, gin.H{
		"to_do_list": todos,
		"done_list":  dones,
	}, "task list information")
}
