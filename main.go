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

	// Inicializa os repositórios
	userRepo := repository.NewUserRepo(db.DB)
	jwtRepo := &repository.JwtRepoImpl{}
	comicRepo := repository.NewComicRepoImpl(db.DB)
	exchangeRepo := repository.NewExchangeRepo(db.DB)
	reviewRepo := repository.NewReviewRepoImpl(db.DB)

	// Inicializa os serviços
	comicSvc := services.NewComicService(comicRepo)
	userSvc := services.NewUserService(userRepo, jwtRepo)
	exchangeSvc := services.NewExchangeService(exchangeRepo, comicRepo)
	reviewSvc := services.NewReviewService(reviewRepo)

	// Inicializa os handlers
	userHandler := handlers.NewUserHandler(userSvc)
	comicHandler := handlers.NewComicHandler(comicSvc)
	exchangeHandler := handlers.NewExchangeHandler(exchangeSvc)
	reviewHandler := handlers.NewReviewHandler(reviewSvc)

	// Rotas de usuário
	app.Post("/user/register", userHandler.Register)
	app.Post("/user/login", userHandler.Login)
	app.Get("/user/me", userHandler.JwtMiddleware, userHandler.Me)

	// Rotas de quadrinhos
	app.Put("/comic", userHandler.JwtMiddleware, comicHandler.CreateComic)
	app.Get("/comic", userHandler.JwtMiddleware, comicHandler.GetComics)

	// Rotas de troca de quadrinhos
	app.Post("/exchange", userHandler.JwtMiddleware, exchangeHandler.RequestExchange)
	app.Post("/exchange/:id/accept", userHandler.JwtMiddleware, exchangeHandler.AcceptExchange)
	app.Post("/exchange/:id/complete", userHandler.JwtMiddleware, exchangeHandler.CompleteExchange)

	// Rotas de avaliações
	app.Post("/comic/:id/review", userHandler.JwtMiddleware, reviewHandler.AddReview)
	app.Get("/comic/:id/reviews", reviewHandler.GetReviews)

	// Rota para Swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Inicia o servidor e adiciona tratamento de erro
	err := app.Listen(":3000")
	if err != nil {
		log.Fatal().Err(err).Msg("Erro ao iniciar o servidor")
	}
}

