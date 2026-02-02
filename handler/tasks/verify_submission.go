package tasks

import (
	"CA_Portal_backend/database"
	"CA_Portal_backend/models"
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func update_user_points(db *mongo.Database, ctx context.Context, username string, points int) error {
	// Normalize username
	username = strings.ToLower(strings.TrimSpace(username))
	log.Printf("Updating points for user: %s (adding %d points)", username, points)

	var result models.User
	if err := db.Collection("users").FindOne(ctx, bson.D{{Key: "username", Value: username}}).Decode(&result); err != nil {
		log.Printf("ERROR: User not found: %s, error: %v", username, err)
		return fmt.Errorf("user does not exist: %s", username)
	}

	log.Printf("Current user points: %d", result.Points)

	// Use $inc operator for atomic increment
	update := bson.M{
		"$inc": bson.M{
			"points": points,
		},
	}

	result_update, err := db.Collection("users").UpdateOne(ctx, bson.D{{Key: "username", Value: username}}, update)
	if err != nil {
		log.Printf("ERROR: Failed to update user points: %v", err)
		return fmt.Errorf("failed to update user points: %v", err)
	}

	log.Printf("✓ User points updated. Matched: %d, Modified: %d", result_update.MatchedCount, result_update.ModifiedCount)
	return nil
}

func get_task_points(db *mongo.Database, ctx context.Context, task_id string) (int, error) {
	var result models.Task
	objectId, err := primitive.ObjectIDFromHex(task_id)
	if err != nil {
		return 0, fmt.Errorf("invalid task ID format: %s", task_id)
	}

	if err := db.Collection("tasks").FindOne(ctx, bson.D{{Key: "_id", Value: objectId}}).Decode(&result); err != nil {
		log.Printf("ERROR: Task not found: %s, error: %v", task_id, err)
		return 0, fmt.Errorf("task does not exist: %s", task_id)
	}

	log.Printf("Task found: %s, points: %d", result.Title, result.Points)
	return result.Points, nil
}

func VerifySubmission(c *fiber.Ctx) error {
	log.Printf("=== VERIFY SUBMISSION REQUEST ===")

	ctx := context.Background()
	db, err := database.Connect()
	if err != nil {
		log.Printf("ERROR: Database connection failed: %v", err)
		return c.Status(500).JSON(fiber.Map{"message": "Database connection error: " + err.Error()})
	}

	collection := db.Collection("task_submissions")

	// Get submission ID from params
	submission_id := c.Params("submission_id")
	log.Printf("Submission ID: %s", submission_id)

	objectId, err := primitive.ObjectIDFromHex(submission_id)
	if err != nil {
		log.Printf("ERROR: Invalid submission ID format: %s", submission_id)
		return c.Status(400).JSON(fiber.Map{"message": "Invalid submission ID format"})
	}

	// Find the submission
	task_submission := new(models.TaskSubmission)
	if err := collection.FindOne(ctx, bson.D{{Key: "_id", Value: objectId}}).Decode(task_submission); err != nil {
		log.Printf("ERROR: Submission not found: %s, error: %v", submission_id, err)
		return c.Status(404).JSON(fiber.Map{"message": "Submission not found"})
	}

	log.Printf("Submission found - User: %s, Task: %s, Current Verified: %v",
		task_submission.User, task_submission.Task, task_submission.Verified)

	// Check if already verified
	if task_submission.Verified {
		log.Printf("WARNING: Submission already verified")
		return c.Status(400).JSON(fiber.Map{"message": "Submission already verified"})
	}

	// Get task points
	points := 0
	if points, err = get_task_points(db, ctx, task_submission.Task); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": err.Error()})
	}

	log.Printf("Task points to award: %d", points)

	// Update user points
	if err = update_user_points(db, ctx, task_submission.User, points); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	// **FIX: Update submission - set verified to true AND clear admin comment**
	update := bson.M{
		"$set": bson.M{
			"verified":      true,
			"admin_comment": "", // Clear comment when verified
			"last_reviewed_at":  time.Now(), // NEW: Track when admin last reviewed/commented
		},
	}

	updateResult, err := collection.UpdateOne(ctx, bson.D{{Key: "_id", Value: objectId}}, update)
	if err != nil {
		log.Printf("ERROR: Failed to update submission: %v", err)
		return c.Status(500).JSON(fiber.Map{"message": "Failed to verify submission: " + err.Error()})
	}

	log.Printf("✓ Submission update result - Matched: %d, Modified: %d", updateResult.MatchedCount, updateResult.ModifiedCount)

	if updateResult.MatchedCount == 0 {
		log.Printf("ERROR: Submission not found during update")
		return c.Status(404).JSON(fiber.Map{"message": "Submission not found"})
	}

	// Update local object
	task_submission.Verified = true
	task_submission.AdminComment = "" // Clear in response too

	log.Printf("✓✓✓ VERIFICATION SUCCESSFUL ✓✓✓")
	log.Printf("===================================\n")

	return c.Status(200).JSON(fiber.Map{
		"message":        "Task verified successfully",
		"submission":     task_submission,
		"points_awarded": points,
	})
}