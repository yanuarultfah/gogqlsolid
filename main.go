package main

import (
	"gogqlgensolid/internal/config"
	dbpkg "gogqlgensolid/internal/db"
	graphql "gogqlgensolid/internal/graph"
	"gogqlgensolid/internal/metrics"
	repository "gogqlgensolid/internal/repository"
	service "gogqlgensolid/internal/usecase"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
)

func main() {

	metrics.InitPrometheus("9090")
	cfg := config.LoadConfig("internal/config/config.yaml")

	// Init Postgres DB with pooling
	dbpkg.Init(cfg.Postgres)

	userRepo := repository.NewUserRepository()
	userSvc := service.NewUserUseCase(userRepo)
	// authCtrl := controller.NewAuthController(userSvc)

	resolver := graphql.NewResolver(userSvc)
	schema, _ := graphql.NewSchema(resolver)

	gqlHandler := handler.New(&handler.Config{Schema: &schema, Pretty: true, GraphiQL: true})

	r := gin.Default()
	// r.POST("/login", authCtrl.Login)
	auth := r.Group("/")
	// auth.Use(middleware.JWTAuth(userSvc))
	auth.POST("/graphql", gin.WrapH(gqlHandler))
	auth.GET("/graphql", gin.WrapH(gqlHandler))

	log.Println("Running at :8180")
	r.Run(":8180")
}
