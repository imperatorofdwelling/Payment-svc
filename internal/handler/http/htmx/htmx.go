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
//go:embed payment.html
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
		r.Get("/payments", hdl.PaymentsPage)
	})

	return hdl
}

// TODO implement html page with redirect and userID in SaveCardPage. Figure out how to insert data into the page.

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

// TODO implement data with ConfirmationToken and ReturnUrl in PaymentsPage. Figure out how to insert data into the page.

func (hdl *htmxHandler) PaymentsPage(w http.ResponseWriter, r *http.Request) {
	h := hdl.htmx.NewHandler(w, r)

	data := map[string]any{
		"ConfirmationToken": "ct-2efb2d91-000f-5000-9000-1c4ddb928aac",
		"ReturnUrl":         "https://ya.ru/",
	}

	page := htmx.NewComponent("payment.html").FS(templates).SetData(data)

	_, err := h.Render(r.Context(), page)
	if err != nil {
		fmt.Printf("error rendering page: %v", err)
	}
}
