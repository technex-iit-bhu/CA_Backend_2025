package tasks

import (
	"CA_Backend/database"
	"CA_Backend/models"
	"CA_Backend/utils"
	"context"
	"log"
	"time"
	"github.com/gofiber/fiber/v2"
)

func SubmitTask(c *fiber.Ctx) error {
	task_submission := new(models.TaskSubmission)
	if err := c.BodyParser(task_submission); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Failed to parse JSON Body",
		})
	}

	if task_submission.DriveLink == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Drive link is required",
		})
	}

	if !utils.IsValidDriveLink(task_submission.DriveLink) {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid Drive link provided",
		})
	}

	tokenString := c.Get("Authorization")
	if len(tokenString) < 8 || tokenString[:7] != "Bearer " {
		return c.Status(400).JSON(fiber.Map{"message": "Authorization header missing or improperly formatted"})
	}
	token := tokenString[7:]
	username, _ := utils.DeserialiseUser(token)

	task_submission.User = username
	task_submission.Timestamp = time.Now()
	task_submission.Verified = false
	task_submission.AdminComment = ""
	ctx := context.Background()

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err.Error())
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}
	collection := db.Collection("task_submissions")
	if res, err := collection.InsertOne(ctx, task_submission); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Failed to create task submission!!",
		})
	} else {
		return c.Status(201).JSON(fiber.Map{
			"id":      res.InsertedID,
			"message": "Submitted successfully",
		})
	}
}
