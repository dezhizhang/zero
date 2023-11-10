package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/types/use"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *use.IdRequest) (*use.UserResponse, error) {
	// todo: add your logic here and delete this line

	return &use.UserResponse{
		Id:     "1001",
		Name:   "晓智科技",
		Gender: false,
	}, nil
}
