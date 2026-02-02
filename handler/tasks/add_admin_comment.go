package tasks

import (
	"CA_Portal_backend/database"
	"CA_Portal_backend/models"
	"context"

	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// POST /api/submissions/comment/:submission_id
func AddAdminComment(c *fiber.Ctx) error {
	ctx := context.Background()

	// Get submission ID from params
	submissionID := c.Params("submission_id")
	objectID, err := primitive.ObjectIDFromHex(submissionID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid submission ID",
		})
	}

	// Parse request body
	var req struct {
		Comment string `json:"comment" binding:"required"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Failed to parse JSON body",
		})
	}

	if req.Comment == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Comment cannot be empty",
		})
	}

	// Connect to database
	db, err := database.Connect()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Database connection error: " + err.Error(),
		})
	}

	collection := db.Collection("task_submissions")

	// Check if submission exists
	var submission models.TaskSubmission
	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&submission)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Submission not found",
		})
	}

	// Update submission with new comment
	update := bson.M{
		"$set": bson.M{
			"admin_comment": req.Comment,
			"verified":      false, // Reset verification when admin comments
			"last_reviewed_at":  time.Now(), // Track when admin last reviewed/commented
		},
	}

	result, err := collection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to add comment: " + err.Error(),
		})
	}

	if result.MatchedCount == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Submission not found",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message":       "Comment added successfully",
		"submission_id": submissionID,
		"comment":       req.Comment,
	})
}