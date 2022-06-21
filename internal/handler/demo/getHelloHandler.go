package demo

import (
	"net/http"

	"demoGo/internal/logic/demo"
	"demoGo/internal/svc"
	"demoGo/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetHelloHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HelloReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := demo.NewGetHelloLogic(r.Context(), svcCtx)
		resp, err := l.GetHello(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
