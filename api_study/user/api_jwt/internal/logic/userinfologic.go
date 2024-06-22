package logic

import (
	"context"
	"zeroStudy/api_study/user/api_jwt/internal/svc"
	"zeroStudy/api_study/user/api_jwt/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResponse, err error) {
	// todo: add your logic here and delete this line
	//userId := l.ctx.Value("user_id").(json.Number)
	//userName := l.ctx.Value("user_name").(string)
	//uid, _ := userId.Int64()
	//return &types.UserInfoResponse{
	//	UserId:   uint(uid),
	//	Username: userName,
	//}, nil
	return &types.UserInfoResponse{
		UserId:   1,
		Username: "aa",
	}, nil
}
