package http

import (
	"github.com/eclipsemode/go-yookassa-sdk/yookassa"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/imperatorofdwelling/payment-svc/internal/config"
	v1 "github.com/imperatorofdwelling/payment-svc/internal/handler/http/api/v1"
	"github.com/imperatorofdwelling/payment-svc/internal/handler/http/htmx"
	kafka "github.com/imperatorofdwelling/payment-svc/internal/handler/kafka/consumer"
	consumer "github.com/imperatorofdwelling/payment-svc/internal/handler/kafka/consumer/payment"
	"github.com/imperatorofdwelling/payment-svc/internal/service"
	"github.com/imperatorofdwelling/payment-svc/internal/storage"
	"github.com/imperatorofdwelling/payment-svc/internal/storage/postgres"
	"github.com/imperatorofdwelling/payment-svc/internal/storage/redis"
	"go.uber.org/zap"
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

		rdbTransactionsRepo := redis.NewTransactionRepo(s.Redis)

		yooClient := yookassa.NewYookassaClient(cfg.PayApi.ShopID, cfg.PayApi.SecretKey, cfg.PayApi.PayoutAgentID, cfg.PayApi.PayoutSecretKey)
		yookassaPaymentsSvc := yookassa.NewPaymentsService(yooClient, log.Named("yookassa_handler"))
		yookassaPayoutsSvc := yookassa.NewPayoutsService(yooClient, log.Named("yookassa_handler"))

		cardsRepo := postgres.NewCardsRepo(s.Psql, log.Named("cards_repo"))
		cardsSvc := service.NewCardsService(cardsRepo, log.Named("cards_service"))

		logsRepo := postgres.NewLogsRepo(s.Psql, log.Named("logs_repo"))
		logsSvc := service.NewLogsService(logsRepo, log.Named("logs_service"))
		v1.NewLogsHandler(r, logsSvc, log.Named("logs_handler"))

		paymentRepo := postgres.NewPaymentRepo(s.Psql, log.Named("payment_repo"))
		paymentSvc := service.NewPaymentSvc(paymentRepo, logsSvc, log.Named("payment_service"))
		v1.NewPaymentsHandler(r, paymentSvc, yookassaPaymentsSvc, log.Named("payment_handler"))

		payoutSubscriber := service.NewPayoutSubscriber(rdbTransactionsRepo, logsSvc, yookassaPayoutsSvc)

		payoutsRepo := postgres.NewPayoutsRepo(s.Psql, log.Named("payouts_repo"))
		payoutsSvc := service.NewPayoutsService(payoutsRepo, payoutSubscriber, logsSvc, log.Named("payouts_service"))
		v1.NewPayoutsHandler(r, payoutsSvc, cardsSvc, yookassaPayoutsSvc, log.Named("payout_handler"))

		paymentConsumer := consumer.NewPaymentConsumer(log.Named("kafka_payment_consumer"), yookassaPaymentsSvc, paymentSvc)

		kafka.SetupKafkaConsumers(paymentConsumer)

		htmx.NewHTMXHandler(r, log.Named("htmx_handler"))

	})

	return &Router{
		Handler: r,
	}
}
