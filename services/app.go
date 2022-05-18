package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/memory"
	"go.uber.org/dig"
	"net/http"
	"runtime"
	"xd_working_trial/db"
	"xd_working_trial/dtos"
	"xd_working_trial/logger"
	"xd_working_trial/repositories"
)

type AppService interface {
	GetOSInfo(c *gin.Context) (*dtos.OSInfoResponse, error)
	GetUserInfo(c *gin.Context) (*dtos.UserInfoResponse, error)
	GetMetricInfo(c *gin.Context) (*dtos.MetricInfoResponse, error)
}

type appService struct {
	db             *db.DB
	userAccessRepo repositories.UserAccessRepository
}

type AppServiceArgs struct {
	dig.In
	DB             *db.DB `name:"xdDB"`
	UserAccessRepo repositories.UserAccessRepository
}

func NewAppService(args AppServiceArgs) AppService {
	return &appService{
		db:             args.DB,
		userAccessRepo: args.UserAccessRepo,
	}
}

func (a *appService) GetOSInfo(c *gin.Context) (*dtos.OSInfoResponse, error) {
	a.saveUserInfo(c)

	return &dtos.OSInfoResponse{
		Meta: dtos.Meta{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		},
		Data: dtos.OSInfoData{
			OSVersion: runtime.GOOS,
		},
	}, nil
}

func (a *appService) GetUserInfo(c *gin.Context) (*dtos.UserInfoResponse, error) {
	a.saveUserInfo(c)

	return &dtos.UserInfoResponse{
		Meta: dtos.Meta{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		},
		Data: dtos.UserInfoData{
			Info: "Xendit - Trial - Son Cao - 2022/05/17 - 2022/05/19",
		},
	}, nil
}

func (a *appService) GetMetricInfo(c *gin.Context) (*dtos.MetricInfoResponse, error) {
	a.saveUserInfo(c)

	memoryMetric, err := memory.Get()
	if err != nil {
		logger.NewLogger().Errorf(fmt.Sprintf("get memory metric error: %s", err.Error()))
		return nil, err
	}

	cpuMetric, er := cpu.Get()
	if er != nil {
		logger.NewLogger().Errorf(fmt.Sprintf("get cpu metric error: %s", err.Error()))
		return nil, err
	}

	return &dtos.MetricInfoResponse{
		Meta: dtos.Meta{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		},
		Data: dtos.MetricInfoData{
			CPUUsed:    fmt.Sprintf("%.2f %%", float64(cpuMetric.User)/float64(cpuMetric.Total)*100),
			MemoryUsed: fmt.Sprintf("%d bytes", memoryMetric.Used),
		},
	}, nil
}

func (a *appService) saveUserInfo(c *gin.Context) {
	fmt.Printf(c.GetHeader("User-Agent"))
	userAccess := &repositories.UserAccess{
		UserInfo: c.GetHeader("User-Agent"),
	}
	err := a.userAccessRepo.Set(a.db, userAccess)
	if err != nil {
		logger.NewLogger().Errorf(fmt.Sprintf("can not save user info with error: %s", err.Error()))
	}
}
