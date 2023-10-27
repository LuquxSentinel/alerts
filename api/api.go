package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/luqus/s/middleware"
	"github.com/luqus/s/storage"
	"github.com/luqus/s/types"
	"log"
)

type APIServer struct {
	listenAddr   string
	router       *fiber.App
	authStorage  storage.AuthenticationStorage
	alertStorage storage.AlertStorage
	pubSubStorage storage.PubSub
}

func NewAPIServer(listenAddr string, authstore storage.AuthenticationStorage, alertStorage storage.AlertStorage, pubsubstorage storage.PubSub) *APIServer {
	return &APIServer{
		listenAddr:   listenAddr,
		router:       fiber.New(),
		authStorage:  authstore,
		alertStorage: alertStorage,
		pubSubStorage: pubsubstorage,
	}
}



func (api *APIServer) publishLocation(c *websocket.Conn) {
	log.Print("Upgraded")
	publishLocationInput  := new(types.PublishLocationInput)
	var (
//		mt int
		msg []byte
		err error
	)
	for {
		if _, msg, err =c.ReadMessage();err !=nil {
			log.Println(err)
			break
		}

		err := json.Unmarshal(msg, publishLocationInput)
		if err != nil {
			err = c.WriteJSON(fmt.Sprintf("error: %v",err.Error()))
			if err != nil {
				log.Println(err.Error())
				break
			}
		}


		err = api.pubSubStorage.Publish(context.Background(),publishLocationInput)
		if err !=nil {
			err = c.WriteJSON(fmt.Sprintf("error: %v",err.Error()))
			if err != nil {
				log.Println(err.Error())
				break
			}
		}
	}

}

func (api *APIServer) subcribeLocation(c *websocket.Conn)  {

	var (
		location *types.Location
		msg []byte
		err error
	)

	for {
		if _, msg,err = c.ReadMessage(); err != nil {
			log.Println(err.Error())
			break
		}

		location, err = api.pubSubStorage.Subscribe(context.Background(), string(msg))
		if err != nil {
			err = c.WriteJSON(fmt.Sprintf("error: %v",err.Error()))
			if err != nil {
				log.Println(err.Error())
				break
			}
		}

		if err := c.WriteJSON(location); err !=nil {
			log.Println(err.Error())
			break
		}


	}
}

func (api *APIServer) Run() error {
	api.router.Post("/register", api.registerUser)
	api.router.Post("/signin", api.login)
	api.router.Use("/ws", middleware.WebSocketUpgrade)
	api.router.Post("/ws/publish", websocket.New(api.publishLocation))

	return api.router.Listen(api.listenAddr)
}


