package user

import (
	"context"

	"github.com/deltafi-trade/looopx-api/internal/svc"
	"github.com/deltafi-trade/looopx-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AccountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccountLogic {
	return &AccountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AccountLogic) Account(req *types.AccountReq) (resp *types.UserInfoReply, err error) {
	// todo: add your logic here and delete this line

	return
}
