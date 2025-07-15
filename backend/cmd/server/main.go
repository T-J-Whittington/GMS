package main

import (
    "log"
    "os"

    "garage-management/internal/config"
    "garage-management/internal/database"
    "garage-management/internal/handlers"
    "garage-management/internal/middleware"
    "garage-management/internal/repositories"
    "garage-management/internal/services"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
)

func main() {
    // Load environment variables
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }

    // Initialize configuration
    cfg := config.New()

    // Initialize database
    db, err := database.New(cfg.DatabaseURL())
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    defer db.Close()

    // Run migrations
    if err := database.RunMigrations(cfg.DatabaseURL()); err != nil {
        log.Printf("Migration error: %v", err)
    }

    // Initialize repositories
    userRepo := repositories.NewUserRepository(db)
    inventoryRepo := repositories.NewInventoryRepository(db)
    jobRepo := repositories.NewJobRepository(db)
    paymentRepo := repositories.NewPaymentRepository(db)

    // Initialize services
    userService := services.NewUserService(userRepo)
    inventoryService := services.NewInventoryService(inventoryRepo)
    jobService := services.NewJobService(jobRepo)
    paymentService := services.NewPaymentService(paymentRepo)

    // Initialize handlers
    userHandler := handlers.NewUserHandler(userService)
    inventoryHandler := handlers.NewInventoryHandler(inventoryService)
    jobHandler := handlers.NewJobHandler(jobService)
    paymentHandler := handlers.NewPaymentHandler(paymentService)

    // Setup router
    router := gin.Default()

    // CORS middleware
    router.Use(middleware.CORS())

    // Health check
    router.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "ok"})
    })

    // API routes
    api := router.Group("/api/v1")
    {
        // Auth routes
        auth := api.Group("/auth")
        {
            auth.POST("/register", userHandler.Register)
            auth.POST("/login", userHandler.Login)
        }

        // Protected routes
        protected := api.Group("/")
        protected.Use(middleware.AuthMiddleware(cfg.JWTSecret))
        {
            // User routes
            users := protected.Group("/users")
            {
                users.GET("/profile", userHandler.GetProfile)
                users.PUT("/profile", userHandler.UpdateProfile)
            }

            // Inventory routes
            inventory := protected.Group("/inventory")
            {
                inventory.GET("/", inventoryHandler.GetAll)
                inventory.POST("/", inventoryHandler.Create)
                inventory.GET("/:id", inventoryHandler.GetByID)
                inventory.PUT("/:id", inventoryHandler.Update)
                inventory.DELETE("/:id", inventoryHandler.Delete)
            }

            // Job routes
            jobs := protected.Group("/jobs")
            {
                jobs.GET("/", jobHandler.GetAll)
                jobs.POST("/", jobHandler.Create)
                jobs.GET("/:id", jobHandler.GetByID)
                jobs.PUT("/:id", jobHandler.Update)
                jobs.DELETE("/:id", jobHandler.Delete)
            }

            // Payment routes
            payments := protected.Group("/payments")
            {
                payments.GET("/", paymentHandler.GetAll)
                payments.POST("/", paymentHandler.Create)
                payments.GET("/:id", paymentHandler.GetByID)
                payments.PUT("/:id", paymentHandler.Update)
            }
        }
    }

    // Start server
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Server starting on port %s", port)
    if err := router.Run(":" + port); err != nil {
        log.Fatal("Failed to start server:", err)
    }
}