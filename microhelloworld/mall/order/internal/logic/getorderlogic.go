package logic

import (
	"context"
	"user/types/user"

	"github.com/zeromicro/go-zero/core/logx"
	"order/internal/svc"
	"order/internal/types"
)

type GetOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderLogic {
	return &GetOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderLogic) GetOrder(req *types.IdRequest) (resp *types.OrderReply, err error) {
	// todo: add your logic here and delete this line

	rsp, err := l.svcCtx.UserRpc.GetUser(context.Background(), &user.IdRequest{
		Id: "123",
	})
	if err != nil {
		return nil, err
	}

	return &types.OrderReply{
		Id:     "123",
		Name:   "hello world",
		Gender: rsp.Gender,
	}, nil
}
