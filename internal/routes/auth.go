package routes

import (
	"fmt"
	"os"
	"time"

	user "MarketPlaceBackEnd/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
)

var u user.UserData

func genJwtToket() (string, error) {
	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Minute * 12).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// The secret key will be stored in an environment variable (In the future...)
	t, err := token.SignedString([]byte(os.Getenv("JWT_SEED")))
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
		return c.Status(fiber.StatusBadRequest).JSON(RespStatus("1.0", fiber.StatusBadRequest, "Incorrect data", nil))
	}

	if err := user.CreateUser(email, pass); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(RespStatus("1.0", fiber.StatusUnauthorized, "A user with this Email has already registered", nil))
	} else {
		return c.Status(fiber.StatusOK).JSON(RespStatus("1.0", fiber.StatusOK, "Registration was successful!", nil))
	}
}

// WORK: Login user
func SignIn(c *fiber.Ctx) error {
	email, pass, err := beforeCreate(&u, c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(RespStatus("1.0", fiber.StatusBadRequest, "Incorrect data", nil))
	}

	if status, err := user.AuthUser(email, pass); !status {
		fmt.Println(status, err)
		return c.Status(fiber.StatusOK).JSON(RespStatus("1.0", fiber.StatusOK, "Incorrect email or password", nil))
	}

	// Issuing a token to a user
	if token, err := genJwtToket(); err != nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(RespStatus("1.0", fiber.StatusServiceUnavailable, "Technical failures", nil))
	} else {
		return c.Status(fiber.StatusOK).JSON(RespStatus("1.0", fiber.StatusOK, token, nil))
	}
}
