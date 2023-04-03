package user

import (
	"net/http"

	"github.com/deltafi-trade/looopx-api/internal/logic/user"
	"github.com/deltafi-trade/looopx-api/internal/svc"
	"github.com/deltafi-trade/looopx-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SettingsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserSettingReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewSettingsLogic(r.Context(), svcCtx)
		resp, err := l.Settings(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
