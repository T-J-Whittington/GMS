package server

import (
	"net/http"
    "garage-management/internal/config"
)

type Server struct {
	mux		*http.ServeMux
	config	*config.Config
}

