package main

import (
    "garage-management/internal/server"
    "garage-management/internal/config"
)

func main() {
    //Start server.
    server.New(config.New())
}