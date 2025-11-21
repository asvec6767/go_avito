package handlers

import (
	"encoding/json"
	"main/internal/usecase"
	"net/http"
	"strconv"
)

type UserHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler(userUseCase usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id_str := r.URL.Query().Get("id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.userUseCase.GetById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "applocation/json")
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) Createuser(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := h.userUseCase.Create(request.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
