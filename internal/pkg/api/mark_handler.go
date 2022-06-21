package api

import (
	"crud-go-server/internal/pkg/entity"
	"crud-go-server/internal/pkg/repo"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func NewMarkHandler(repo repo.MarkRepo) *MarkHandler {
	return &MarkHandler{markRepo: repo}
}

type MarkHandler struct {
	*BaseHandler
	markRepo repo.MarkRepo
}

func (h *MarkHandler) Route() chi.Router {
	router := chi.NewRouter()

	// with query student_no & module_code
	router.Get("/", h.handleGetMarks)
	router.Post("/create", h.handleCreate)
	router.Put("/", h.handleUpdateByID)
	router.Delete("/delete", h.handleDelete)

	return router
}

func (h *MarkHandler) handleGetMarks(w http.ResponseWriter, r *http.Request) {
	studentID := r.URL.Query().Get("student_no")
	if studentID == "" {
		studentID = ".*"
	}

	module_code := r.URL.Query().Get("module_code")
	if module_code == "" {
		module_code = ".*"
	}

	// Get items
	students, err := h.markRepo.GetMarks(studentID, module_code)
	if err != nil {
		_ = render.Render(w, r, entity.ErrNotFound())
		return
	}

	if students == nil {
		_ = render.Render(w, r, entity.ErrNotFound())
		return
	}

	h.Success(w, r, students)
}

func (h *MarkHandler) handleCreate(w http.ResponseWriter, r *http.Request) {
	var mark entity.Mark

	if err := json.NewDecoder(r.Body).Decode(&mark); err != nil {
		_ = render.Render(w, r, entity.ErrInvalidRequest(err))
		return
	}

	if err := h.markRepo.CreateMark(&mark); err != nil {
		_ = render.Render(w, r, entity.ErrInvalidRequest(err))
		return
	}
	h.Success(w, r, mark)
}

func (h *MarkHandler) handleUpdateByID(w http.ResponseWriter, r *http.Request) {
	var mark entity.Mark

	if err := json.NewDecoder(r.Body).Decode(&mark); err != nil {
		_ = render.Render(w, r, entity.ErrInvalidRequest(err))
		return
	}

	if err := h.markRepo.Update(&mark); err != nil {
		_ = render.Render(w, r, entity.ErrInvalidRequest(err))
		return
	}
	h.Success(w, r, mark)
}

func (h *MarkHandler) handleDelete(w http.ResponseWriter, r *http.Request) {
	var req DeleteMarkReq

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		_ = render.Render(w, r, entity.ErrInvalidRequest(err))
		return
	}

	if err := h.markRepo.Delete(req.StudentID, req.ModuleCode); err != nil {
		_ = render.Render(w, r, entity.ErrInvalidRequest(err))
		return
	}
	h.Success(w, r, nil)
}

type DeleteMarkReq struct {
	StudentID  string `json:"student_no"`
	ModuleCode string `json:"module_code"`
}
