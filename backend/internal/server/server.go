package server

import (
	"log"
	"net/http"

    "garage-management/internal/config"
    "garage-management/internal/user"
    // "garage-management/internal/inventory"
    // "garage-management/internal/customer"
    // "garage-management/internal/vehicle"
    // "garage-management/internal/job"
    // "garage-management/internal/payment"
)

const API_PREFIX string = '/api/v1/';

// Start up the server and set things up.
func New(config *config.Config) {
	//Try to begin server.
    log.Printf("Server starting on %s", config.Port);

	server := &Server{
        mux:    http.NewServeMux(),
        config: config,
    }

	log.Printf("Server started.");

	// Initialise database
	log.Prinf("Beginning database connection and migration.")
    db, err := database.New(cfg.DatabaseURL())
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    defer db.Close()

    // Run migrations
    if err := database.RunMigrations(cfg.DatabaseURL()); err != nil {
        log.Printf("Migration error: %v", err)
    }
    
    // Setup routes and serve.
    server.setupRoutes()
	log.Fatal(http.ListenAndServe(":" + config.Port, server.mux))
}

func (server *Server) setupRoutes() {
	log.Printf("Setting up routes.");
    server.mux.HandleFunc("/health", func(response http.ResponseWriter, request *http.Request) {
        response.WriteHeader(http.StatusOK)
        response.Write([]byte("OK"))
    });

    // API routes
	// @TODO: Is there some kind of variable variable so I can loop through this rather than
	// add one in manually each time?
    user.RegisterRoutes(server.mux, API_PREFIX)
    // vehicle.RegisterRoutes(server.mux, API_PREFIX)
	// inventory.RegisterRoutes(server.mux, API_PREFIX)
	// customer.RegisterRoutes(server.mux, API_PREFIX)
	// job.RegisterRoutes(server.mux, API_PREFIX)

	log.Printf("Routing successful.")
}