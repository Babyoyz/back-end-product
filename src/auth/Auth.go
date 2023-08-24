package auth

import (
	"back_product/model"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func LoginHandler(c *fiber.Ctx) error {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// JWT configuration
	secret := os.Getenv("JWT_SECRET")

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = "Thaweeporn.a@gmail.com"
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	st, err := token.SignedString([]byte(secret))
	if err != nil {
		modelResonseLogin := model.ResponseLogin{
			Message: "error sign jwt",
			Status:  500,
			Token:   "",
		}

		return c.JSON(modelResonseLogin)
	}

	modelResonseLogin := model.ResponseLogin{
		Message: "login successful",
		Status:  200,
		Token:   st,
	}

	return c.JSON(modelResonseLogin)
}

func AuthorizationMiddleware(c *fiber.Ctx) error {

	s := c.Get("Authorization")

	token := strings.TrimPrefix(s, "Bearer ")

	if err := validateToken(token); err != nil {

		modelResonseLogin := model.ResponseLogin{
			Message: "Authorization header extracted",
			Status:  500,
			Token:   "",
		}
		return c.JSON(modelResonseLogin)
	}

	return c.Next()

}

func validateToken(tokenString string) error {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	secret := os.Getenv("JWT_SECRET")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		fmt.Println("Error parsing token:", err)

		return err
	}

	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		fmt.Println("User:", claims["user"])
		fmt.Println("Expires:", claims["exp"])

		authorized := claims["authorized"].(bool)
		if authorized {
			fmt.Println("User is authorized.")
			return nil
		} else {

			return nil
		}
	} else {
		fmt.Println("Couldn't handle token:", err)
	}
	return err

}
