package user

import (
	"context"

	"github.com/deltafi-trade/looopx-api/internal/svc"
	"github.com/deltafi-trade/looopx-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout() (resp *types.ReplyStatus, err error) {
	// todo: add your logic here and delete this line

	return
}
