package task

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/clrajapaksha/to-do-list-app/entities"
	"github.com/go-chi/chi/v5"
)

func TestTaskServiceHandlers(t *testing.T) {
	taskRepository := &mockTaskRepository{}
	handler := NewHandler(taskRepository)

	payload := entities.Task{
		Titile:      "test title",
		Description: "test description",
	}
	marshalled, _ := json.Marshal(payload)
	t.Run("Should pass if the task payload is valid", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := chi.NewRouter()

		router.Post("/tasks", handler.CreateTask)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusCreated, rr.Code)
		}
	})
}

type mockTaskRepository struct{}

func (m *mockTaskRepository) Save(task *entities.Task) (*entities.Task, error) {
	return nil, nil
}

func (m *mockTaskRepository) FindAll() ([]entities.Task, error) {
	return nil, nil
}

func (m *mockTaskRepository) FindByID(id string) (*entities.Task, error) {
	return nil, nil
}

func (m *mockTaskRepository) Delete(task *entities.Task) error {
	return nil
}
