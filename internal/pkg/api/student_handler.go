package api

import (
	"crud-go-server/internal/pkg/entity"
	"crud-go-server/internal/pkg/repo"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func NewStudentHandler(repo repo.StudentRepo) *StudentHandler {
	return &StudentHandler{studentRepo: repo}
}

type StudentHandler struct {
	*BaseHandler
	studentRepo repo.StudentRepo
}

func (h *StudentHandler) Route() chi.Router {
	router := chi.NewRouter()

	router.Get("/", h.handleGetStudents)
	router.Get("/{id}", h.handleGetStudentByID)
	router.Post("/create", h.handleCreate)
	router.Put("/{id}", h.handleUpdateByID)
	router.Delete("/delete", h.handleDelete)

	return router
}

func (h *StudentHandler) handleGetStudents(w http.ResponseWriter, r *http.Request) {
	// Get items
	students, err := h.studentRepo.GetStudents()
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

func (h *StudentHandler) handleGetStudentByID(w http.ResponseWriter, r *http.Request) {
	var student *entity.FullStudent

	if id := chi.URLParam(r, "id"); id != "" {
		intID, err := strconv.Atoi(id)
		if err != nil {
			_ = render.Render(w, r, entity.ErrInvalidRequest(err))
			return
		}

		// Get item
		student, err = h.studentRepo.GetStudentByID(intID)
		if err != nil {
			_ = render.Render(w, r, entity.ErrNotFound())
			return
		}
	}

	if student == nil {
		_ = render.Render(w, r, entity.ErrNotFound())
		return
	}

	h.Success(w, r, student)
}

func (h *StudentHandler) handleCreate(w http.ResponseWriter, r *http.Request) {
	var student entity.Student

	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		_ = render.Render(w, r, entity.ErrInvalidRequest(err))
		return
	}

	if err := h.studentRepo.CreateStudent(&student); err != nil {
		_ = render.Render(w, r, entity.ErrInvalidRequest(err))
		return
	}
	h.Success(w, r, student)
}

func (h *StudentHandler) handleUpdateByID(w http.ResponseWriter, r *http.Request) {
	var id = chi.URLParam(r, "id")
	var student entity.Student

	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		_ = render.Render(w, r, entity.ErrInvalidRequest(err))
		return
	}

	student.StudentID = id
	if err := h.studentRepo.Update(&student); err != nil {
		_ = render.Render(w, r, entity.ErrInvalidRequest(err))
		return
	}
	h.Success(w, r, student)
}

func (h *StudentHandler) handleDelete(w http.ResponseWriter, r *http.Request) {
	var req DeleteStudentReq

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		_ = render.Render(w, r, entity.ErrInvalidRequest(err))
		return
	}

	if err := h.studentRepo.Delete(req.Id); err != nil {
		_ = render.Render(w, r, entity.ErrInvalidRequest(err))
		return
	}
	h.Success(w, r, nil)
}

type DeleteStudentReq struct {
	Id string `json:"id"`
}
