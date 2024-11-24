package main

import (
	"acervoback/db"
	"acervoback/handlers"
	"acervoback/repository"
	"acervoback/services"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func main() {
	// Carregar variáveis de ambiente
	if err := godotenv.Load(); err != nil {
		log.Warn().Msg("Arquivo .env não encontrado, usando variáveis padrão")
	}

	// Validar variáveis de ambiente
	if os.Getenv("DATABASE_URL") == "" {
		log.Fatal().Msg("DATABASE_URL não configurado")
	}
	if os.Getenv("GOOGLE_BOOKS_API_KEY") == "" {
		log.Warn().Msg("GOOGLE_BOOKS_API_KEY não configurado")
	}

	// Inicializar o banco de dados
	if err := db.InitDB(); err != nil {
		log.Fatal().Err(err).Msg("Erro ao conectar ao banco de dados")
	}

	// Inicializar Fiber
	app := fiber.New()

	// Inicializar repositórios
	userRepo := repository.NewUserRepo(db.DB)
	comicRepo := repository.NewComicRepoImpl(db.DB)
	exchangeRepo := repository.NewExchangeRepo(db.DB)
	reviewRepo := repository.NewReviewRepoImpl(db.DB)

	// Inicializar o JWT Repo
	jwtRepo := &repository.JwtRepoImpl{}

	// Chave de API da Google Books
	googleBooksAPIKey := os.Getenv("GOOGLE_BOOKS_API_KEY")

	// Inicializar serviços
	userSvc := services.NewUserService(userRepo, jwtRepo)
	comicSvc := services.NewComicService(comicRepo, googleBooksAPIKey)
	exchangeSvc := services.NewExchangeService(exchangeRepo, comicRepo)
	reviewSvc := services.NewReviewService(reviewRepo)

	// Inicializar handlers
	userHandler := handlers.NewUserHandler(userSvc)
	comicHandler := handlers.NewComicHandler(comicSvc)
	exchangeHandler := handlers.NewExchangeHandler(exchangeSvc)
	reviewHandler := handlers.NewReviewHandler(reviewSvc)

	// Rotas
	app.Post("/user/register", userHandler.Register)
	app.Post("/user/login", userHandler.Login)
	app.Get("/user/me", userHandler.JwtMiddleware, userHandler.Me)

	app.Put("/comic", userHandler.JwtMiddleware, comicHandler.CreateComic)
	app.Get("/comic", userHandler.JwtMiddleware, comicHandler.GetComics)
	app.Get("/exchange", userHandler.JwtMiddleware, exchangeHandler.ListExchanges)

	app.Post("/exchange", userHandler.JwtMiddleware, exchangeHandler.RequestExchange)
	app.Patch("/exchange/:id", userHandler.JwtMiddleware, exchangeHandler.AcceptExchange)

	app.Post("/comic/:id/review", userHandler.JwtMiddleware, reviewHandler.AddReview)
	app.Get("/comic/:id/reviews", reviewHandler.GetReviews)

	// Iniciar servidor
	if err := app.Listen(":3000"); err != nil {
		log.Fatal().Err(err).Msg("Erro ao iniciar o servidor")
	}
}
