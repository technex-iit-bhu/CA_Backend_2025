package tasks

import (
	"CA_Backend/database"
	"CA_Backend/models"
	"CA_Backend/utils"
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SubmitTask(c *fiber.Ctx) error {
	taskSubmission := new(models.TaskSubmission)
	if err := c.BodyParser(taskSubmission); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Failed to parse JSON Body",
		})
	}

	if taskSubmission.DriveLink == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Drive link is required",
		})
	}

	// If you want to validate the link format:
	if !utils.IsValidDriveLink(taskSubmission.DriveLink) {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid Drive link provided",
		})
	}

	// Extract token
	tokenString := c.Get("Authorization")
	if len(tokenString) < 8 || tokenString[:7] != "Bearer " {
		return c.Status(400).JSON(fiber.Map{
			"message": "Authorization header missing or improperly formatted",
		})
	}
	token := tokenString[7:]
	username, _ := utils.DeserialiseUser(token)

	// Fill in submission details
	taskSubmission.User = username
	taskSubmission.Timestamp = time.Now()
	taskSubmission.Verified = false
	taskSubmission.AdminComment = ""

	// Insert into DB
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err.Error())
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}
	collection := db.Collection("task_submissions")

	// Check if user already submitted for this task
	existingSubmission := new(models.TaskSubmission)
	err = collection.FindOne(context.Background(), bson.M{
		"user": username,
		"task": taskSubmission.Task,
	}).Decode(existingSubmission)

	// If err == nil => found doc => conflict
	if err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "You have already submitted this task.",
		})
	} else if err != mongo.ErrNoDocuments {
		// Some other error
		return c.Status(500).JSON(fiber.Map{
			"message": "Database error: " + err.Error(),
		})
	}

	// If we reach here => no existing submission => insert new
	res, err := collection.InsertOne(context.Background(), taskSubmission)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Failed to create task submission!!",
		})
	}
	return c.Status(201).JSON(fiber.Map{
		"id":      res.InsertedID,
		"message": "Submitted successfully",
	})
}
