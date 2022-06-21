package api

import (
	"crud-go-server/internal/pkg/entity"
	"crud-go-server/internal/pkg/repo"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func NewModuleHandler(repo repo.ModuleRepo) *ModuleHandler {
	return &ModuleHandler{moduleRepo: repo}
}

type ModuleHandler struct {
	*BaseHandler
	moduleRepo repo.ModuleRepo
}

func (h *ModuleHandler) Route() chi.Router {
	router := chi.NewRouter()

	router.Get("/", h.handleGetModules)
	router.Get("/{module_code}", h.handleGetModuleByID)
	router.Post("/create", h.handleCreate)
	router.Put("/{module_code}", h.handleUpdateByID)
	router.Delete("/delete", h.handleDelete)

	return router
}

func (h *ModuleHandler) handleGetModules(w http.ResponseWriter, r *http.Request) {
	// Get items
	students, err := h.moduleRepo.GetModules()
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

func (h *ModuleHandler) handleGetModuleByID(w http.ResponseWriter, r *http.Request) {
	var (
		module *entity.Module
		err    error
	)

	if id := chi.URLParam(r, "module_code"); id != "" {
		// Get item
		module, err = h.moduleRepo.GetModuleByID(id)
		if err != nil {
			_ = render.Render(w, r, entity.ErrNotFound())
			return
		}
	}

	if module == nil {
		_ = render.Render(w, r, entity.ErrNotFound())
		return
	}

	h.Success(w, r, module)
}

func (h *ModuleHandler) handleCreate(w http.ResponseWriter, r *http.Request) {
	var module entity.Module

	if err := json.NewDecoder(r.Body).Decode(&module); err != nil {
		_ = render.Render(w, r, entity.ErrInvalidRequest(err))
		return
	}

	if err := h.moduleRepo.CreateModule(&module); err != nil {
		_ = render.Render(w, r, entity.ErrInvalidRequest(err))
		return
	}
	h.Success(w, r, module)
}

func (h *ModuleHandler) handleUpdateByID(w http.ResponseWriter, r *http.Request) {
	var id = chi.URLParam(r, "module_code")
	var module entity.Module

	if err := json.NewDecoder(r.Body).Decode(&module); err != nil {
		_ = render.Render(w, r, entity.ErrInvalidRequest(err))
		return
	}

	module.ModuleCode = id
	if err := h.moduleRepo.Update(&module); err != nil {
		_ = render.Render(w, r, entity.ErrInvalidRequest(err))
		return
	}
	h.Success(w, r, module)
}

func (h *ModuleHandler) handleDelete(w http.ResponseWriter, r *http.Request) {
	var req DeleteModuleReq

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		_ = render.Render(w, r, entity.ErrInvalidRequest(err))
		return
	}

	if err := h.moduleRepo.Delete(req.Id); err != nil {
		_ = render.Render(w, r, entity.ErrInvalidRequest(err))
		return
	}
	h.Success(w, r, nil)
}

type DeleteModuleReq struct {
	Id string `json:"module_code"`
}
