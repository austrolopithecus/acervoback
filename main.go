package main

import (
	"acervoback/db"
	_ "acervoback/docs"
	"acervoback/handlers"
	"acervoback/repository"
	"acervoback/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

// @title Acervo Comics
// @version 1.0
// @description Projeto Acervo Comics
// @BasePath /
// @securityDefinitions.apikey TokenAuth
// @in header
// @name Authorization
func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	_ = godotenv.Load()
	app := fiber.New()
	db.InitDB()
	userRepo := repository.NewUserRepo(db.DB)
	jwtRepo := &repository.JwtRepoImpl{}
	comicRepo := repository.NewComicRepoImpl(db.DB)
	comicSvc := services.NewComicService(comicRepo)
	userSvc := services.NewUserService(userRepo, jwtRepo)
	userHandler := handlers.NewUserHandler(userSvc)
	comicHandler := handlers.NewComicHandler(comicSvc)
	app.Post("/user/register", userHandler.Register)
	app.Post("/user/login", userHandler.Login)
	app.Get("/user/me", userHandler.JwtMiddleware, userHandler.Me)
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Put("/comic", userHandler.JwtMiddleware, comicHandler.CreateComic)
	app.Get("/comic", userHandler.JwtMiddleware, comicHandler.GetComics)
	_ = app.Listen(":3000")
}
