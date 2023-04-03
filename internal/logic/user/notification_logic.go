package user

import (
	"context"

	"github.com/deltafi-trade/looopx-api/internal/svc"
	"github.com/deltafi-trade/looopx-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type NotificationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNotificationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NotificationLogic {
	return &NotificationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NotificationLogic) Notification() (resp *types.NotificationReply, err error) {
	// todo: add your logic here and delete this line

	return
}
