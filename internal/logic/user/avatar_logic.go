package user

import (
	"context"

	"github.com/deltafi-trade/looopx-api/internal/svc"
	"github.com/deltafi-trade/looopx-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AvatarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AvatarLogic {
	return &AvatarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AvatarLogic) Avatar() (resp *types.UserInfoReply, err error) {
	// todo: add your logic here and delete this line

	return
}
