// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package serviceexporter

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-micro-saas/service-layout/app/testing-service/internal/biz/biz"
	"github.com/go-micro-saas/service-layout/app/testing-service/internal/data/data"
	"github.com/go-micro-saas/service-layout/app/testing-service/internal/service/service"
	"github.com/ikaiguang/go-srv-kit/service/cleanup"
	"github.com/ikaiguang/go-srv-kit/service/setup"
)

// Injectors from wire.go:

func exportServices(launcherManager setuputil.LauncherManager, hs *http.Server, gs *grpc.Server) (cleanuputil.CleanupManager, error) {
	logger, err := setuputil.GetLogger(launcherManager)
	if err != nil {
		return nil, err
	}
	testingDataRepo := data.NewTestingData(logger)
	testingBizRepo := biz.NewTestingBiz(logger, testingDataRepo)
	srvTestdataServer := service.NewTestingV1Service(logger, testingBizRepo)
	cleanupManager, err := service.RegisterServices(hs, gs, srvTestdataServer)
	if err != nil {
		return nil, err
	}
	return cleanupManager, nil
}
