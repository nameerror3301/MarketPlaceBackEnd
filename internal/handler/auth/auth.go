package routes

import (
	"fmt"
	"time"

	responce "MarketPlaceBackEnd/internal/handler"

	user "MarketPlaceBackEnd/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
)

var u user.UserData

func genJwtToket() (string, error) {
	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 12).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// The secret key will be stored in an environment variable (In the future...)
	t, err := token.SignedString([]byte("SECRET"))
	if err != nil {
		logrus.Fatalf("Error in the process of generating a token for the user - %s", err)
	}
	return t, nil
}

func beforeCreate(u *user.UserData, c *fiber.Ctx) (string, string, error) {
	err := c.BodyParser(&u)
	if err != nil {
		return "", "", err
	}
	return u.Email, u.Password, nil
}

// WORK: Create user
func SignUp(c *fiber.Ctx) error {
	email, pass, err := beforeCreate(&u, c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responce.RespStatus("1.0", fiber.StatusBadRequest, "Incorrect data", nil))
	}

	if err := user.CreateUser(email, pass); err != nil {
		return c.Status(fiber.StatusOK).JSON(responce.RespStatus("1.0", fiber.StatusOK, "A user with this Email has already registered", nil))
	} else {
		return c.Status(fiber.StatusOK).JSON(responce.RespStatus("1.0", fiber.StatusUnauthorized, "Registration was successful !", nil))
	}
}

// WORK: Login user
func SignIn(c *fiber.Ctx) error {
	email, pass, err := beforeCreate(&u, c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responce.RespStatus("1.0", fiber.StatusBadRequest, "Incorrect data", nil))
	}

	if status, err := user.AuthUser(email, pass); !status {
		fmt.Println(status, err)
		return c.Status(fiber.StatusOK).JSON(responce.RespStatus("1.0", fiber.StatusOK, "Incorrect email or password", nil))
	}

	// Issuing a token to a user
	if token, err := genJwtToket(); err != nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(responce.RespStatus("1.0", fiber.StatusServiceUnavailable, "Technical failures", nil))
	} else {
		return c.Status(fiber.StatusOK).JSON(responce.RespStatus("1.0", fiber.StatusOK, token, nil))
	}
}
