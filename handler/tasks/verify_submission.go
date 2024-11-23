package tasks

import (
	"CA_Backend/database"
	"CA_Backend/models"
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func update_user_points(db *mongo.Database, ctx context.Context, username string, points int) error {
	var result models.User
	if err := db.Collection("users").FindOne(ctx, bson.D{{Key: "username", Value: username}}).Decode(&result); err != nil {
		return fmt.Errorf("user does not exist")
	}
	result.Points += points

	if _, err := db.Collection("users").UpdateOne(ctx, bson.D{{Key: "username", Value: username}}, result); err != nil {
		return err
	} else {
		return nil
	}
}

func get_task_points(db *mongo.Database, ctx context.Context, task_id string) (int, error) {
	var result models.Task
	objectId, _ := primitive.ObjectIDFromHex(task_id)
	if err := db.Collection("tasks").FindOne(ctx, bson.D{{Key: "_id", Value: objectId}}).Decode(&result); err != nil {
		return 0, fmt.Errorf("task does not exist")
	}
	return result.Points, nil
}

func VerifySubmission(c *fiber.Ctx) error {
	ctx := context.Background()
	db, err := database.Connect()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	collection := db.Collection("task_submissions")

	task_submission := new(models.TaskSubmission)
	submssion_id := c.Params("submission_id")
	objectId, _ := primitive.ObjectIDFromHex(submssion_id)
	if err := collection.FindOne(ctx, bson.D{{Key: "_id", Value: objectId}}).Decode(task_submission); err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Submission not found"})
	}

	task_submission.Verified = true
	points := 0
	if points, err = get_task_points(db, ctx, task_submission.Task); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": err.Error()})
	}
	if err = update_user_points(db, ctx, task_submission.User, points); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}
	if _, err = collection.UpdateOne(ctx, bson.D{{Key: "_id", Value: objectId}}, task_submission); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	} else {
		return c.Status(200).JSON(fiber.Map{
			"message":    "Task updated Successfully",
			"submission": task_submission,
		})
	}
}
