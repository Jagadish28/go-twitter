package controller

import (
	"strconv"
	"strings"

	"github.com/Jagadish28/go-twitter/pkg/database"
	"github.com/Jagadish28/go-twitter/pkg/model"

	"github.com/gofiber/fiber/v2"
)

func GetTweets(c *fiber.Ctx) error {

	userId := validateUserID(c.Get("cookie"), c)
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))

	// Calculate the offset
	offset := (page - 1) * pageSize

	var tweets []model.Tweet
	var total int64

	// Fetch products with pagination
	database.DB.Model(&model.Tweet{}).Where("user_id = ?", userId).Count(&total)
	database.DB.Offset(offset).Limit(pageSize).Where("user_id = ?", userId).Find(&tweets)

	response := fiber.Map{
		"data":     tweets,
		"page":     page,
		"pageSize": pageSize,
		"total":    total,
	}

	return c.JSON(response)
}

func PostTweet(c *fiber.Ctx) error {
	tweet := new(model.Tweet)
	if err := c.BodyParser(tweet); err != nil {
		return c.JSON(model.Response{Code: 500, Message: err.Error()})
	}
	database.DB.Create(&tweet)

	return c.JSON(tweet)
}

func validateUserID(requestCookie string, c *fiber.Ctx) int {
	if requestCookie == "" {
		c.JSON(model.Response{Code: 400, Message: "user Details not found"})
	}
	substrings := strings.Split(requestCookie, "=")
	var uId int
	if len(substrings) > 1 {
		userId, err := strconv.Atoi(substrings[1])
		if err != nil {
			c.JSON(model.Response{Code: 400, Message: "user Details not found"})
		}
		uId = userId
	} else {
		c.JSON(model.Response{Code: 400, Message: "user Details not found"})
	}

	return uId

}
