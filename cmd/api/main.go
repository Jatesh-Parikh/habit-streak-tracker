package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"habit-streak-tracker/internal/database"
	"habit-streak-tracker/internal/graph/generated"
	"habit-streak-tracker/internal/graph/resolvers"
	"habit-streak-tracker/internal/middleware"
	"habit-streak-tracker/internal/repository"
)

func main() {
	fmt.Println("🚀 Starting Habit Streak Tracker GraphQL API...")

	db, err := database.InitDB("./data/habits.db")

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	defer db.Close()

	fmt.Println("✅ Database connected successfully!")

	userRepo := repository.NewUserRepository(db)
	habitRepo := repository.NewHabitRepository(db)
	habitLogRepo := repository.NewHabitLogRepository(db)

	graphqlHandler := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &resolvers.Resolver{
					UserRepo:     userRepo,
					HabitRepo:    habitRepo,
					HabitLogRepo: habitLogRepo,
				},
			}))

	var router *gin.Engine = gin.Default()
	router.SetTrustedProxies(nil)

	router.GET("/playground", func(c *gin.Context) {
		playground.Handler("GraphQL Playground", "/graphql").ServeHTTP(c.Writer, c.Request)
	})

	router.POST("/graphql", middleware.AuthMiddleware(), func(c *gin.Context) {
		graphqlHandler.ServeHTTP(c.Writer, c.Request)
	})

	fmt.Println("GraphQL Playground: http://localhost:3000/playground")
	fmt.Println("GraphQL Endpoint:   http://localhost:3000/graphql")

	port := os.Getenv("PORT")

	router.Run(":" + port)
}
