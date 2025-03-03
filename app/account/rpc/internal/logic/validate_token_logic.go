package logic

import (
	"context"
	"github.com/gookit/goutil/strutil"
	"github.com/luxun9527/gex/common/errs"
	"github.com/luxun9527/gex/common/proto/define"
	logger "github.com/luxun9527/zlog"

	"github.com/luxun9527/gex/app/account/rpc/internal/svc"
	"github.com/luxun9527/gex/app/account/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ValidateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewValidateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ValidateTokenLogic {
	return &ValidateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 验证token是否有效。
func (l *ValidateTokenLogic) ValidateToken(in *pb.ValidateTokenReq) (*pb.ValidateTokenResp, error) {
	claims, err := l.svcCtx.JwtClient.ParseToken(in.Token)
	if err != nil {
		logx.Errorw("parse token failed", logger.ErrorField(err), logx.Field("token", in.Token))
		return nil, errs.Internal
	}

	tokenMd5 := strutil.Md5(in.Token)
	existed, err := l.svcCtx.RedisClient.ExistsCtx(l.ctx, define.AccountToken.WithParams(tokenMd5))
	if err != nil {
		logx.Errorw("get redis key failed", logger.ErrorField(err))
		return nil, errs.RedisErr
	}
	if !existed {
		return nil, errs.TokenValidateFailed
	}
	return &pb.ValidateTokenResp{
		Uid:      claims.UserID,
		Username: claims.Username,
	}, nil
}
