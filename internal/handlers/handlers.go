package internal

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/shegai01/LO_task/internal/logger"
	"github.com/shegai01/LO_task/internal/model"
	"github.com/shegai01/LO_task/internal/storage"
)

func initContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

type TaskHandler struct {
	someLog *logger.ChanLogs
	storage *storage.Storage
}

func NewTimerHandler(db *storage.Storage, someLog *logger.ChanLogs) *TaskHandler {
	h := &TaskHandler{
		storage: db,
		someLog: someLog,
	}

	return h
}

func Error(w http.ResponseWriter, statusCode int) {
	http.Error(w, http.StatusText(statusCode), statusCode)
}

func (h *TaskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.URL.Path == "/task" && r.Method == http.MethodGet:
		h.ShowList(w, r)

	case r.URL.Path == "/task" && (r.Method == http.MethodPost || r.Method == http.MethodGet):
		h.Create(w, r)

	case r.URL.Path == "/" && r.Method == http.MethodGet:
		h.ShowList(w, r)

	default:
		h.someLog.Info("")
		http.NotFoundHandler()
	}
}

func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	var task *model.Task

	title := r.URL.Query().Get("title")

	task, err := h.storage.Create(title)
	if err != nil {
		h.someLog.Info("h.storage.CreateTask:")
		Error(w, http.StatusInternalServerError)

		return
	}
	h.someLog.Info("Create task called: " + title)
	initContentType(w)

	if err = json.NewEncoder(w).Encode(task); err != nil {
		h.someLog.Info("json.NewEncoder.Encode:")
		Error(w, http.StatusInternalServerError)

		return
	}
}

func (h *TaskHandler) ShowList(w http.ResponseWriter, r *http.Request) {
	var allTasks []*model.Task

	allTasks, err := h.storage.List("")
	if err != nil {
		h.someLog.Info("h.storage.ShowList:")
		Error(w, http.StatusInternalServerError)

		return
	}

	initContentType(w)
	h.someLog.Info("ShowALL")

	if err = json.NewEncoder(w).Encode(allTasks); err != nil {
		h.someLog.Info("json.NewEncoder.Encode:")
		Error(w, http.StatusInternalServerError)

		return
	}
}

func (h *TaskHandler) GetbyID(w http.ResponseWriter, r *http.Request) {
	idTask := r.URL.Query().Get("id")
	if idTask == "" {
		h.someLog.Info("id query parameter is not set")
		Error(w, http.StatusBadRequest)

		return
	}

	strID, err := strconv.Atoi(idTask)
	if err != nil {
		h.someLog.Info("strconv.Atoi:")
		Error(w, http.StatusBadRequest)

		return
	}

	task, ok := h.storage.Get(int64(strID))
	if !ok {
		h.someLog.Info("h.storage.GetTaskByID:")
		Error(w, http.StatusInternalServerError)

		return
	}
	h.someLog.Info("Get by id: " + idTask)
	initContentType(w)

	if err = json.NewEncoder(w).Encode(task); err != nil {
		h.someLog.Info("json.NewEncoder.Encode:")
		Error(w, http.StatusInternalServerError)

		return
	}
}
