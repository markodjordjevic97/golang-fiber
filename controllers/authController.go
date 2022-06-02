package controllers

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-example/database"
	"github.com/golang-example/models"
	"github.com/golang-example/util"
)

func Register(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["passwordConfirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Password do not match",
		})
	}

	user := models.User{
		FirstName: data["firstName"],
		LastName:  data["lastName"],
		Email:     data["email"],
		RoleId:    1,
	}

	user.SetPassword(data["password"])

	database.DB.Create(&user)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "email not found",
		})
	}

	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "incorect password",
		})
	}

	token, err := util.GenerateJWT(strconv.Itoa(int(user.Id)))

	if err != nil {
		c.SendStatus(fiber.StatusInternalServerError)
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Success",
		"token":   token,
	})
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	id, _ := util.ParseJWT(cookie)

	var user models.User

	database.DB.Where("id = ?", id).First(&user)

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Success Logout",
	})
}

func UpdateInfo(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	cookie := c.Cookies("jwt")

	id, _ := util.ParseJWT(cookie)

	userId, _ := strconv.Atoi(id)

	user := models.User{
		Id:        uint(userId),
		FirstName: data["firstName"],
		LastName:  data["lastName"],
		Email:     data["email"],
	}

	database.DB.Model(&user).Where("id = ?", id).Updates(user)

	return c.JSON(user)
}

func UpdatePassword(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["passwordConfirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Password do not match",
		})
	}

	cookie := c.Cookies("jwt")

	id, _ := util.ParseJWT(cookie)

	userId, _ := strconv.Atoi(id)

	user := models.User{
		Id: uint(userId),
	}

	user.SetPassword(data["password"])

	database.DB.Model(&user).Updates(user)

	return c.JSON(user)
}
