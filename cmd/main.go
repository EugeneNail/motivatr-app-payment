package main

import (
	"fmt"
	"github.com/EugeneNail/motivatr-app-payment/internal/application/commands"
	"github.com/EugeneNail/motivatr-app-payment/internal/infrastructure/config"
	"github.com/EugeneNail/motivatr-app-payment/internal/infrastructure/repository/postgres"
	transport "github.com/EugeneNail/motivatr-app-payment/internal/transport/http"
	"github.com/EugeneNail/motivatr-lib-common/pkg/databases"
	middlewares "github.com/EugeneNail/motivatr-lib-common/pkg/middlewares/http"
	"net/http"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(fmt.Errorf("creating a config instance: %w", err))
	}

	db, err := databases.ConnectToPostgres(cfg.Db.Host, cfg.Db.Post, cfg.Db.Name, cfg.Db.User, cfg.Db.Password)
	if err != nil {
		panic(fmt.Errorf("connecting to the database: %w", err))
	}

	paymentRepository := postgres.NewPaymentRepository(db)

	createPaymentHandler := commands.NewCreatePaymentHandler(paymentRepository)

	httpHandler := transport.NewHandler(createPaymentHandler)

	router := http.NewServeMux()

	router.HandleFunc("POST /api/v1/payments", middlewares.Authenticate(cfg.Jwt.Salt)(middlewares.WriteJsonResponse(httpHandler.CreatePayment)))
	if err := http.ListenAndServe("0.0.0.0:10000", middlewares.DisableLocalCors(router)); err != nil {
		panic(err)
	}
}
