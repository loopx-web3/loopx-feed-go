package user

import (
	"context"

	"github.com/deltafi-trade/looopx-api/internal/svc"
	"github.com/deltafi-trade/looopx-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SettingsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSettingsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SettingsLogic {
	return &SettingsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SettingsLogic) Settings(req *types.UserSettingReq) (resp *types.UserInfoReply, err error) {
	// todo: add your logic here and delete this line

	return
}
