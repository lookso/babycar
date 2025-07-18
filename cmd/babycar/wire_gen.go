// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"babycare/internal/biz/baby"
	"babycare/internal/biz/car"
	"babycare/internal/biz/tree"
	"babycare/internal/conf"
	"babycare/internal/data"
	"babycare/internal/server"
	baby2 "babycare/internal/service/baby"
	car2 "babycare/internal/service/car"
	tree2 "babycare/internal/service/tree"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"

	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, errorHandle *conf.ErrorHandle, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	iCarRepo := data.NewCarData(dataData, logger)
	carBiz := car.NewCarBiz(iCarRepo, logger)
	carService := car2.NewCarService(carBiz, logger)
	treeBiz := tree.NewTreeBiz(dataData, logger)
	treeService := tree2.NewTreeService(treeBiz, logger)
	grpcServer := server.NewGRPCServer(confServer, logger, carService, treeService)
	iBabyRepo := data.NewBabyData(dataData, logger)
	babyBiz := baby.NewBabyBiz(iBabyRepo, logger)
	babyService := baby2.NewBabyService(babyBiz, logger)
	httpServer := server.NewHTTPServer(confServer, confData, errorHandle, carService, babyService, treeService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
