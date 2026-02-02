package tasks

import (
	"CA_Portal_backend/database"
	"CA_Portal_backend/models"
	"CA_Portal_backend/utils"
	"context"
	"log"
	"time"
	"strings"

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

	// Validate Drive link format
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
	username = strings.ToLower(username)
	log.Printf("Username (lowercase): %s", username)


	// Fill in submission details - ENSURE BOTH User AND Username are set
	taskSubmission.User = username
	taskSubmission.Username = username  // IMPORTANT: Set this too!
	taskSubmission.Timestamp = time.Now()
	taskSubmission.Verified = false
	taskSubmission.AdminComment = ""

	// Connect to database
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

	// If submission exists
	if err == nil {
		// Check if already verified
		if existingSubmission.Verified {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"message": "This task has already been verified. No further submissions allowed.",
			})
		}

		// Allow resubmission - update existing
		log.Printf("Updating submission for user: %s, task: %s", username, taskSubmission.Task)

		update := bson.M{
			"$set": bson.M{
				"drive_link":    taskSubmission.DriveLink,
				"username":      username, // IMPORTANT: Update username too
				"timestamp":     time.Now(),
				"verified":      false,
				"admin_comment": "",
				"image_url":     taskSubmission.ImageUrl,
			},
		}

		result, err := collection.UpdateOne(
			context.Background(),
			bson.M{"_id": existingSubmission.ID},
			update,
		)

		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error":   err.Error(),
				"message": "Failed to update submission",
			})
		}

		if result.MatchedCount == 0 {
			return c.Status(404).JSON(fiber.Map{
				"message": "Submission not found",
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"id":      existingSubmission.ID,
			"message": "Submission updated successfully",
		})
	} else if err != mongo.ErrNoDocuments {
		return c.Status(500).JSON(fiber.Map{
			"message": "Database error: " + err.Error(),
		})
	}

	// No existing submission - create new
	log.Printf("Creating new submission for user: %s, task: %s", username, taskSubmission.Task)

	res, err := collection.InsertOne(context.Background(), taskSubmission)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Failed to create task submission",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"id":      res.InsertedID,
		"message": "Submitted successfully",
	})
}