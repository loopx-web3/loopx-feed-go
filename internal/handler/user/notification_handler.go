package user

import (
	"net/http"

	"github.com/deltafi-trade/looopx-api/internal/logic/user"
	"github.com/deltafi-trade/looopx-api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func NotificationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewNotificationLogic(r.Context(), svcCtx)
		resp, err := l.Notification()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
