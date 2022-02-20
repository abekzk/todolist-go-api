// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kzuabe/todolist-go-api/internal/controller"
	"github.com/kzuabe/todolist-go-api/internal/repository"
	"github.com/kzuabe/todolist-go-api/internal/router"
	"github.com/kzuabe/todolist-go-api/internal/usecase"
	"github.com/kzuabe/todolist-go-api/pkg/middleware"
)

// Injectors from wire.go:

func initializeRouter() (*gin.Engine, error) {
	db, err := repository.NewDB()
	if err != nil {
		return nil, err
	}
	taskRepository := repository.NewTaskRepository(db)
	taskUseCase := usecase.NewTaskUseCase(taskRepository)
	taskController := controller.NewTaskController(taskUseCase)
	firebaseAuthMiddleware, err := middleware.NewFirebaseAuthMiddleware()
	if err != nil {
		return nil, err
	}
	engine := router.NewRouter(taskController, firebaseAuthMiddleware)
	return engine, nil
}
