package api

import (


	"github.com/gofiber/fiber/v2"
	"github.com/luqus/s/storage"
)

type APIServer struct {
	listenAddr   string
	router       *fiber.App
	authStorage  storage.AuthenticationStorage
	alertStorage storage.AlertStorage
}

func NewAPIServer(listenAddr string, authstore storage.AuthenticationStorage, alertStorage storage.AlertStorage) *APIServer {
	return &APIServer{
		listenAddr:   listenAddr,
		router:       fiber.New(),
		authStorage:  authstore,
		alertStorage: alertStorage,
	}
}

func (api *APIServer) Run() error {



	return api.router.Listen(api.listenAddr)
}


