package logic

import (
	"cloud-disk/core/helper"
	"cloud-disk/core/models"
	"context"
	"errors"
	"log"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.UserRegisterResponse, err error) {
	//判断code是否一致
	code1, err := l.svcCtx.RDB.Get(l.ctx, req.Email).Result()
	if err != nil {
		return nil, errors.New("未获取验证码")
	}
	code2 := req.Code
	if code1 != code2 {
		return nil, errors.New("验证码不一致")
	}
	count, err := l.svcCtx.Engine.Where("name = ?", req.Name).Count(models.UserBasic{})
	if err != nil {
		return nil, err
	}
	//判断用户名是否存在
	if count > 0 {
		return nil, errors.New("用户名已存在")
	}
	//数据入库
	user := models.UserBasic{
		Identity: helper.GetUUID(),
		Name:     req.Name,
		Password: helper.Md5(req.Password),
		Email:    req.Email,
	}

	insert, err := l.svcCtx.Engine.Insert(user)
	if err != nil {
		return nil, err
	}
	log.Println("insert user row", insert)

	return
}
