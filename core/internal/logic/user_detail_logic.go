package logic

import (
	"cloud-disk/core/models"
	"context"
	"errors"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetailLogic) UserDetail(req *types.UserDetailRequest) (resp *types.UserDetailResponse, err error) {
	println(44)
	resp = &types.UserDetailResponse{}
	user := new(models.UserBasic)
	get, err := l.svcCtx.Engine.Where("identity=?", req.Identity).Get(user)
	if err != nil {
		return nil, err
	}
	if !get {
		return nil, errors.New("没有此用户")
	}

	resp.Name = user.Name
	resp.Email = user.Email
	return
}
