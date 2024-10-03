package main

import (
	"github.com/MuhammadIbraAlfathar/app-tweet/config"
	"github.com/MuhammadIbraAlfathar/app-tweet/internal/domain/auth"
	"github.com/MuhammadIbraAlfathar/app-tweet/internal/domain/user"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	godotenv.Load()

	config.LoadEnv()

	db, err := config.NewPostgres()
	if err != nil {
		log.Fatal("ERROR CONNECT DB")
	}

	engine := config.NewGin()

	//USER
	userRepo := user.NewRepository(db)
	//userUseCase := user.NewUseCase(userRepo)

	//AUTH
	authUseCase := auth.NewUseCase(userRepo)
	auth.NewController(engine, authUseCase)

	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "hahah",
	//	})
	//})
	//r.Run()

	err = engine.Run()
	if err != nil {
		log.Fatal(err)
		//	test
	}
}
