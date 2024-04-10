package main

import (
	"net/http"

	"go-mail-sender/internal/domain/campaign"
	"go-mail-sender/internal/endpoints"
	database "go-mail-sender/internal/infra/db"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	campaignService := campaign.CampaignService{
		Repository: &database.CampaignRepository{},
	}

	handler := endpoints.Handler{
		CampaignService: campaignService,
	}

	router.Post("/campaign", endpoints.HandlerError(handler.CampaignPost))
	router.Get("/campaign", endpoints.HandlerError(handler.CampaignGet))

	http.ListenAndServe("127.0.0.1:3000", router)
}
