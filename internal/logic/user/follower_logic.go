package user

import (
	"context"

	"github.com/deltafi-trade/looopx-api/internal/svc"
	"github.com/deltafi-trade/looopx-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowerLogic {
	return &FollowerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowerLogic) Follower(req *types.UserFollowerReq) (resp *types.UserFollowerReply, err error) {
	// todo: add your logic here and delete this line

	return
}
