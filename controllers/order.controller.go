package controllers

import (
	"encoding/csv"
	"github.com/PogunGun/golang-fiber-rest-api/database"
	"github.com/PogunGun/golang-fiber-rest-api/models"
	"github.com/gofiber/fiber/v2"
	"os"
	"strconv"
)

func AllOrders(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(database.DB, &models.Order{}, page))
}
func Export(c *fiber.Ctx) error {
	filePath := "./csv/orders.csv"
	if err := CreateFile(filePath); err != nil {
		return err
	}
	return c.Download(filePath)
}
func CreateFile(filePath string) error {
	file, err := os.Create(filePath)

	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	var orders []models.Order
	database.DB.Preload("OrderItems").Find(&orders)
	writer.Write([]string{
		"ID", "Name", "Email", "Product Title", "Price", "Quantity",
	})

	for _, order := range orders {
		data := []string{
			strconv.Itoa(int(order.Id)),
			order.FirstName + " " + order.LastName,
			order.Email,
			"",
			"",
			"",
		}
		if err := writer.Write(data); err != nil {
			return err
		}
		for _, orderItem := range order.OrderItems {

			data := []string{
				"",
				"",
				"",
				orderItem.ProductTitle,
				strconv.Itoa(int(orderItem.Price)),
				strconv.Itoa(int(orderItem.Quantity)),
			}
			if err := writer.Write(data); err != nil {
				return err
			}
		}
	}
	return nil
}
