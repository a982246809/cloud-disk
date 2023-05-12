package logic

import (
	"cloud-disk/core/helper"
	"cloud-disk/core/models"
	"context"
	"errors"
	"time"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MailCoedSendRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMailCoedSendRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MailCoedSendRegisterLogic {
	return &MailCoedSendRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MailCoedSendRegisterLogic) MailCoedSendRegister(req *types.MailCodeSendRequest) (resp *types.MailCodeSendRequest, err error) {
	//邮箱没注册
	count, err := l.svcCtx.Engine.Where("email = ?", req.Email).Count(new(models.UserBasic))
	if err != nil {
		return nil, err
	}
	if count > 0 {
		err = errors.New("邮箱已注册")
		return
	}
	code := helper.RandCode()
	//存储验证码
	l.svcCtx.RDB.Set(l.ctx, req.Email, code, time.Second*500)

	err = helper.MailSendCode(req.Email, code)
	if err != nil {
		return nil, err
	}
	return
}
