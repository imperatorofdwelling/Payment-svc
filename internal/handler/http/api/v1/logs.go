package v1

import (
	yoomodel "github.com/eclipsemode/go-yookassa-sdk/yookassa/model"
	"github.com/go-chi/chi/v5"
	"github.com/imperatorofdwelling/payment-svc/internal/service"
	"github.com/imperatorofdwelling/payment-svc/pkg/json"
	"go.uber.org/zap"
	"net/http"
)

type logsHandler struct {
	svc service.ILogsSvc
	log *zap.SugaredLogger
}

func NewLogsHandler(r chi.Router, svc service.ILogsSvc, log *zap.SugaredLogger) {
	handler := &logsHandler{svc: svc, log: log}

	r.Route("/logs", func(r chi.Router) {
		r.Post("/status", handler.changeStatus)
	})
}

// TODO implement changeStatus handler when server will be published
func (h *logsHandler) changeStatus(w http.ResponseWriter, r *http.Request) {
	const op = "handler.v1.lost.changeStatus"

	var notification yoomodel.Notification

	err := json.Read(r.Body, &notification)
	if err != nil {
		h.log.Error(op, zap.Error(err))
		json.WriteError(w, http.StatusBadRequest, err.Error(), json.DecodeBodyError)
		return
	}

	err = h.svc.UpdateLogTransactionStatus(r.Context(), notification.Object.ID, notification.Object.Status)
	if err != nil {
		h.log.Error(op, zap.Error(err))
		json.WriteError(w, http.StatusInternalServerError, err.Error(), json.DecodeBodyError)
		return
	}

	json.Write(w, http.StatusNoContent, nil)
}
