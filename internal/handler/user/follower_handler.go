package user

import (
	"net/http"

	"github.com/deltafi-trade/looopx-api/internal/logic/user"
	"github.com/deltafi-trade/looopx-api/internal/svc"
	"github.com/deltafi-trade/looopx-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FollowerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserFollowerReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewFollowerLogic(r.Context(), svcCtx)
		resp, err := l.Follower(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
