package config

import (
	"log"
	"github.com/joho/godotenv"
	"os"
)

func Config(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Warning: Error loading .env file")
	}
	return os.Getenv(key)
}

// Config for local tests

// func Config(key string) string {
// 	path,_:=os.Getwd()
// 	dir, err := filepath.Abs(filepath.Dir(path))
// 	if err != nil {
// 		fmt.Print(err)
// 	}
// 	environmentPath := filepath.Join(dir, ".env")
// 	err = godotenv.Load(environmentPath)
// 	if err != nil {
// 		fmt.Println("Error loading .env file")
// 	}
// 	return os.Getenv(key)
// }
