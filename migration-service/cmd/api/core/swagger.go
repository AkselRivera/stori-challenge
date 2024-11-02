package core

import "github.com/gofiber/contrib/swagger"

var SwaggerConfig = swagger.Config{
	BasePath: "/",
	FilePath: "./docs/swagger.json",
	Path:     "docs",
	Title:    "Migration Service | API Docs",
}
