package handlers

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/onsana/order_service/database"
	"github.com/onsana/order_service/model"
)

func CreateNewOrder(c fiber.Ctx) error {
	// order := new(model.Order)
	type input struct {
		Users    []model.User    `json:"users"`
		Products []model.Product `json:"products"`
	}
	p := new(input)

	if err := c.Bind().JSON(p); err != nil {
		return err
	}
	log.Println(p)
	// for _, user := range p.Users {
	// 	database.DB.Db.Create(&user)
	// }

	for _, product := range p.Products {
		log.Println(product)
		database.DB.Db.Create(&product)
	}

	return c.Status(200).JSON(p)
}
