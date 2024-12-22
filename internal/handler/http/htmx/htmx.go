package htmx

import (
	"embed"
	"fmt"
	"github.com/donseba/go-htmx"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
)

//go:embed create-card.html
var templates embed.FS

type htmxHandler struct {
	htmx *htmx.HTMX
	log  *zap.SugaredLogger
}

func NewHTMXHandler(r chi.Router, logger *zap.SugaredLogger) *htmxHandler {
	hdl := &htmxHandler{
		htmx: htmx.New(),
		log:  logger,
	}

	htmx.UseTemplateCache = false

	r.Route("/htmx", func(r chi.Router) {
		r.Get("/cards/{userId}", hdl.SaveCardPage)
	})

	return hdl
}

// TODO implement html page with redirect and userID in SaveCardPage

func (hdl *htmxHandler) SaveCardPage(w http.ResponseWriter, r *http.Request) {
	h := hdl.htmx.NewHandler(w, r)

	userID := chi.URLParam(r, "userId")

	data := map[string]any{
		"Text":     "The form for entering the bank card number.",
		"UserID":   userID,
		"Redirect": "https://ya.ru/",
	}

	page := htmx.NewComponent("create-card.html").FS(templates).SetData(data)

	_, err := h.Render(r.Context(), page)
	if err != nil {
		fmt.Printf("error rendering page: %v", err)
	}
}
