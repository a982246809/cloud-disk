package logic

import (
	"cloud-disk/core/helper"
	"cloud-disk/core/internal/middleware"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

type UserRepositorySaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	Auth   rest.Middleware
}

func NewUserRepositorySaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRepositorySaveLogic {
	return &UserRepositorySaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		Auth:   middleware.NewAuthMiddleware().Handle,
	}
}

func (l *UserRepositorySaveLogic) UserRepositorySave(req *types.UserRepositorySaveRequest, userIdentity string) (resp *types.UserRepositorySaveResponse, err error) {
	userRepository := &models.UserRepository{
		Identity:           helper.GetUUID(),
		UserIdentity:       userIdentity,
		Name:               req.Name,
		Ext:                req.Ext,
		ParentId:           req.ParentId,
		RepositoryIdentity: req.RepositoryIdentity,
	}
	_, err = l.svcCtx.Engine.Insert(userRepository)
	if err != nil {
		return
	}
	return
}
