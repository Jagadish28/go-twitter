package routes

import (
	"github.com/Jagadish28/go-twitter/pkg/controller"
	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("api/v1")

	route.Get("/tweets", controller.GetTweets)
	route.Post("/tweet", controller.PostTweet)

}
