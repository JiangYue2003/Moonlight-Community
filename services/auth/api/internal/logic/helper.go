package logic

import (
	"github.com/zhiguang/zhiguang-go/services/auth/api/internal/types"
	userpb "github.com/zhiguang/zhiguang-go/services/user/rpc/user"
)

// pbToAuthResp 将 RPC AuthResp 映射为 HTTP types.AuthResp。
func pbToAuthResp(r *userpb.AuthResp) *types.AuthResp {
	if r == nil {
		return nil
	}
	out := &types.AuthResp{}
	if r.User != nil {
		out.User = types.AuthUser{
			Id:       r.User.Id,
			Nickname: r.User.Nickname,
			Avatar:   r.User.Avatar,
			Phone:    r.User.Phone,
			ZgId:     r.User.ZgId,
			ZhId:     r.User.ZgId,
			Birthday: r.User.Birthday,
			School:   r.User.School,
			Bio:      r.User.Bio,
			Gender:   r.User.Gender,
			TagsJson: r.User.TagsJson,
			TagJson:  r.User.TagsJson,
		}
	}
	if r.Token != nil {
		out.Token = types.TokenPair{
			AccessToken:           r.Token.AccessToken,
			AccessExpiresAt:       r.Token.AccessExpiresAt,
			AccessTokenExpiresAt:  r.Token.AccessExpiresAt,
			RefreshToken:          r.Token.RefreshToken,
			RefreshExpiresAt:      r.Token.RefreshExpiresAt,
			RefreshTokenExpiresAt: r.Token.RefreshExpiresAt,
			RefreshTokenId:        r.Token.RefreshTokenId,
		}
	}
	return out
}
