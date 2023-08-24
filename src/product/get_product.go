package product

import (
	"back_product/model"

	"github.com/gofiber/fiber/v2"
)

func GetAllProduct(c *fiber.Ctx) error {

	modelResonseLogin := model.ResponseLogin{
		Message: "ok",
		Status:  200,
		Token:   "",
	}

	return c.JSON(modelResonseLogin)

}
