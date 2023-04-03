package comment

import (
	"net/http"

	"github.com/deltafi-trade/looopx-api/internal/logic/comment"
	"github.com/deltafi-trade/looopx-api/internal/svc"
	"github.com/deltafi-trade/looopx-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ReadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CommentListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := comment.NewReadLogic(r.Context(), svcCtx)
		resp, err := l.Read(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
