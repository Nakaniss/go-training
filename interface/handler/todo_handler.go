package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo-app/entity"
	"todo-app/usecase"
)

type TodoHandler struct {
	Usecase *usecase.TodoUsecase
}

func NewTodoHandler(u *usecase.TodoUsecase) *TodoHandler {
	return &TodoHandler{Usecase: u}
}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo entity.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if err := h.Usecase.CreateTodo(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *TodoHandler) GetTodo(w http.ResponseWriter, r *http.Request) {
	todoID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	todo, err := h.Usecase.GetTodo(int64(todoID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	todoID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var todo entity.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	todo.ID = int64(todoID)
	if err := h.Usecase.UpdateTodo(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	todoID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := h.Usecase.DeleteTodo(int64(todoID)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
