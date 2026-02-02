package tasks

import (
	"CA_Portal_backend/database"
	"CA_Portal_backend/models"
	"CA_Portal_backend/utils"
	"context"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetUserSubmissions(c *fiber.Ctx) error {
	log.Printf("=== GET USER SUBMISSIONS REQUEST ===")

	ctx := context.Background()
	db, err := database.Connect()
	if err != nil {
		log.Printf("ERROR: Database connection failed: %v", err)
		return c.Status(500).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Database connection error",
		})
	}
	log.Printf("✓ Database connected: %s", db.Name())

	tokenString := c.Get("Authorization")
	log.Printf("Authorization header: %s", tokenString[:50]+"...") // Log first 50 chars

	if len(tokenString) < 8 || tokenString[:7] != "Bearer " {
		log.Printf("ERROR: Invalid authorization header format")
		return c.Status(400).JSON(fiber.Map{"message": "Authorization header missing or improperly formatted"})
	}
	token := tokenString[7:]

	username, err := utils.DeserialiseUser(token)
	if err != nil {
		log.Printf("ERROR: Failed to deserialize token: %v", err)
		return c.Status(401).JSON(fiber.Map{"message": "Invalid token"})
	}

	log.Printf("Username from token: '%s'", username)

	// Convert to lowercase for consistent querying
	usernameLower := strings.ToLower(username)
	log.Printf("Username (lowercase): '%s'", usernameLower)

	var user_submissions []models.TaskSubmission

	// Query using lowercase username
	query := bson.M{"user": usernameLower}
	log.Printf("MongoDB query: %+v", query)

	cursor, err := db.Collection("task_submissions").Find(ctx, query, options.Find())
	if err != nil {
		log.Printf("ERROR: Database query failed: %v", err)
		return c.Status(500).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Failed to fetch submissions",
		})
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &user_submissions); err != nil {
		log.Printf("ERROR: Failed to decode results: %v", err)
		return c.Status(500).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Error decoding submissions",
		})
	}

	log.Printf("Query result: Found %d submissions", len(user_submissions))

	if len(user_submissions) > 0 {
		for i, sub := range user_submissions {
			log.Printf("  [%d] ID: %s, Task: %s, User: %s, Username: %s, AdminComment: '%s'",
				i+1, sub.ID.Hex(), sub.Task, sub.User, sub.Username, sub.AdminComment)
		}
	} else {
		log.Printf("  ⚠️  No submissions found for user: %s", usernameLower)

		// Debug: Check if any submissions exist
		count, _ := db.Collection("task_submissions").CountDocuments(ctx, bson.M{})
		log.Printf("  Total submissions in collection: %d", count)
	}

	// Return empty array instead of nil
	if user_submissions == nil {
		user_submissions = []models.TaskSubmission{}
	}

	log.Printf("Returning JSON with %d submissions", len(user_submissions))
	log.Printf("===================================\n")

	return c.Status(200).JSON(user_submissions)
}