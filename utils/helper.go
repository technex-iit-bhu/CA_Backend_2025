package utils

import (
	"github.com/gofiber/fiber/v2"
)

type resMessage struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// this function can be called as helpers.ResponseMsg(c, code, message, data)
func ResponseMsg(c *fiber.Ctx, code int, msg string, data interface{}) error {
	response := &resMessage{
		Code:    code,
		Message: msg,
		Data:    data,
	}
	return c.Status(code).JSON(response)
}
