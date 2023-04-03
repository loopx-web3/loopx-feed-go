package user

import (
	"context"

	"github.com/deltafi-trade/looopx-api/internal/svc"
	"github.com/deltafi-trade/looopx-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowingLogic {
	return &FollowingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowingLogic) Following(req *types.UserFollowingReq) (resp *types.UserFollowingReply, err error) {
	// todo: add your logic here and delete this line

	return
}
