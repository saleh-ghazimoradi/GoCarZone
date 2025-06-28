package handlers

import (
	"github.com/saleh-ghazimoradi/GoCarZone/internal/service"
	"net/http"
)

type EngineHandler struct {
	engineService service.EngineService
}

func (e *EngineHandler) GetEngineById(w http.ResponseWriter, r *http.Request) {}

func (e *EngineHandler) CreateEngine(w http.ResponseWriter, r *http.Request) {}

func (e *EngineHandler) UpdateEngine(w http.ResponseWriter, r *http.Request) {}

func (e *EngineHandler) DeleteEngine(w http.ResponseWriter, r *http.Request) {}

func NewEngineHandler(engineService service.EngineService) *EngineHandler {
	return &EngineHandler{
		engineService: engineService,
	}
}
