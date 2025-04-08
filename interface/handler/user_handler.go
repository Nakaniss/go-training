package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo-app/entity"
	"todo-app/interface/dto"
	"todo-app/usecase"
)

type UserHandler struct {
	Usecase *usecase.UserUsecase
}

func NewUserHandler(u *usecase.UserUsecase) *UserHandler {
	return &UserHandler{Usecase: u}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	createdUser, err := h.Usecase.CreateUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// レスポンスに機密情報を含めないためにDTOを使用(テストなのでEmailが表示されていますが許容しています)
	userResponse := dto.UserResponse{
		ID:    createdUser.ID,
		Name:  createdUser.Name,
		Email: createdUser.Email,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userResponse)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	user, err := h.Usecase.GetUser(int64(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var user entity.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	user.ID = int64(userID)
	if err := h.Usecase.UpdateUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := h.Usecase.DeleteUser(int64(userID)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
