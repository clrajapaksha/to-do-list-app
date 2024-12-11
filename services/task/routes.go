package task

import (
	"fmt"
	"net/http"

	"github.com/clrajapaksha/to-do-list-app/entities"
	"github.com/clrajapaksha/to-do-list-app/repository"
	"github.com/clrajapaksha/to-do-list-app/utils"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	repository repository.TaskRepository
}

func NewHandler(TaskRepository repository.TaskRepository) *Handler {
	return &Handler{repository: TaskRepository}
}

func (h *Handler) RegisterRoutes(router *chi.Mux) {
	router.Post("/tasks", h.CreateTask)
	router.Get("/tasks", h.GetAllTasks)
	router.Get("/tasks/{id}", h.GetTaskById)
}

// createTask responds with the task as JSON.
// @Tags Tasks
// @Summary create task
// @Description test description
// @Accept json
// @Produce json
// @Param data body entities.TaskCreate true "The input task struct"
// @Success 201 {object} entities.Task
// @Router /tasks [post]
func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	payload := entities.TaskCreate{}
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	task := entities.Task{
		Id:          utils.GetMD5Hash(payload.Titile),
		Titile:      payload.Titile,
		Description: payload.Description,
	}
	taskCreated, err := h.repository.Save(&task)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("task creation failed"))
		return
	}
	utils.WriteJSON(w, http.StatusCreated, taskCreated)

}

// getAllTasks responds with the list of all tasks as JSON.
// @Tags Tasks
// @Summary get all tasks
// @Description test description
// @Produce json
// @Success 200 {array} entities.Task
// @Router /tasks [get]
func (h *Handler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.repository.FindAll()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	utils.WriteJSON(w, http.StatusOK, map[string][]entities.Task{"tasks": tasks})
}

// getTaskByID responds with the task as JSON.
// @Tags Tasks
// @Summary get task by id
// @Description test description
// @Param id path string  true  "Task ID"
// @Produce json
// @Success 200 {object} entities.Task
// @Router /tasks/{id} [get]
func (h *Handler) GetTaskById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	task, err := h.repository.FindByID(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}
	if task == nil {
		utils.WriteJSON(w, http.StatusNotFound, map[string]string{"error": fmt.Sprintf("task not found for task_id: %s", id)})
		return
	}

	utils.WriteJSON(w, http.StatusOK, task)
}
