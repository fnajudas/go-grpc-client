package handlers

import (
	"base-project/constructs"
	"context"
	"net/http"
)

type studentSvc interface {
	GetStudentByNik(ctx context.Context, nik string) (*constructs.StudentResponse, error)
}

type Handler struct {
	studentSvc studentSvc
}

func NewHandler(studentSvc studentSvc) *Handler {
	return &Handler{
		studentSvc: studentSvc,
	}
}

func (h *Handler) GetStudent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	nik := r.URL.Query().Get("nik")
	if nik == "" {
		http.Error(w, "nik is required", http.StatusBadRequest)
		return
	}

	bRes, err := h.studentSvc.GetStudentByNik(ctx, nik)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"name": "` + bRes.Name + `"}`))
}
