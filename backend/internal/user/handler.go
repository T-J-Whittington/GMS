package user

type Handler struct {
	service Service
	logger *log.Logger
}

func NewHandler(service Service, logger *log.Logger) *Handler {
	return &Handler {
		user: user,
		logger: logger
	}
}

func (handler *Handler) RegisterRoutes(mux *http.ServeMux, prefix string) {
	mux.HandleFunc(prefix + "users", handler.handleUsers)
    mux.HandleFunc(prefix + "user/", handler.handleUserByID)
    mux.HandleFunc(prefix + "login", handler.handleLogin)
}

func (handler *Handler) handleUsers(response http.ResponseWriter, request *http.Request) {
    handler.logger.Printf("Getting users")
    
    users, err := handler.service.GetUsers()
    if err != nil {
        handler.logger.Printf("Error getting users: %v", err)
        http.Error(response, "Internal server error", 500)
        return
    }
    
    // Return JSON response
    json.NewEncoder(response).Encode(users)
}

func (handler *Handler) handleUserById(response http.ResponseWriter, request *http.Request) {
    handler.logger.Printf("Getting user by ID")
    
	// TODO: Just winged this request handling, make sure it works.
	userID := request.data.id;
    users, err := handler.service.GetUserById(userID)
    if err != nil {
        handler.logger.Printf("Error getting user for ID %s: %v.", userID, err)
        http.Error(response, "Internal server error", 500)
        return
    }
    
    // Return JSON response
    json.NewEncoder(response).Encode(users)
}

func (handler *Handler) handleLogin(response http.ResponseWriter, request *http.Request) {
    handler.logger.Printf("Getting user by ID")
    
	// TODO: Just winged this request handling, make sure it works.
	userID := request.data.id;
	userPassword := request.data.password;
    loggedIn, err := handler.service.Login(userID, userPassword)
    if err != nil {
        handler.logger.Printf("Error getting user for ID %s: %v.", userID, err)
        http.Error(response, "Internal server error", 500)
        return
    }
    
    // Return JSON response
    json.NewEncoder(response).Encode(loggedIn)
}