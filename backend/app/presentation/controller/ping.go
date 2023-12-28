package controller

import (
	"context"

	"github.com/minomusigk/tech-blog-summary/backend/oapi"
)

type PingController struct{}

func NewPingController() *PingController {
	return &PingController{}
}

func (*PingController) GetPing(c context.Context, _ oapi.GetPingRequestObject) (oapi.GetPingResponseObject, error) {
	return oapi.GetPing200JSONResponse{Message: "pong"}, nil
}
