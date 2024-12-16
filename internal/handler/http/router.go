package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/imperatorofdwelling/payment-svc/internal/config"
	v1 "github.com/imperatorofdwelling/payment-svc/internal/handler/http/api/v1"
	"github.com/imperatorofdwelling/payment-svc/internal/service"
	"github.com/imperatorofdwelling/payment-svc/internal/storage"
	"github.com/imperatorofdwelling/payment-svc/internal/storage/postgres"
	"github.com/imperatorofdwelling/payment-svc/internal/storage/redis"
	"github.com/rvinnie/yookassa-sdk-go/yookassa"
	"go.uber.org/zap"
	"strconv"
	"time"
)

type Router struct {
	Handler *chi.Mux
}

func NewRouter(s *storage.Storage, log *zap.SugaredLogger, cfg *config.Config) *Router {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/api/v1", func(r chi.Router) {
		yooclient := yookassa.NewClient(strconv.Itoa(cfg.ShopID), cfg.SecretKey)
		yookassaHdl := yookassa.NewPaymentHandler(yooclient)

		v1.NewStaticHandler(r, log)

		cardsRepo := postgres.NewCardsRepo(s.Psql, log.Named("cards_repo"))
		cardsSvc := service.NewCardsService(cardsRepo, log.Named("cards_service"))
		v1.NewCardsHandler(r, cardsSvc, log.Named("cards_handler"))

		logsRepo := postgres.NewLogsRepo(s.Psql, log.Named("logs_repo"))
		logsSvc := service.NewLogsService(logsRepo, log.Named("logs_service"))

		paymentRepo := postgres.NewPaymentRepo(s.Psql, log.Named("payment_repo"))
		paymentSvc := service.NewPaymentSvc(paymentRepo, logsSvc, log.Named("payment_service"))
		v1.NewPaymentsHandler(r, paymentSvc, yookassaHdl, log.Named("payment_handler"))

		_ = redis.NewTransactionRepo(s.Redis)
	})

	return &Router{
		Handler: r,
	}
}
