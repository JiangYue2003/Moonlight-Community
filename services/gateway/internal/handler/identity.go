package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zhiguang/zhiguang-go/common/ctxdata"
	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	gh "github.com/zhiguang/zhiguang-go/services/gateway/internal/httpx"
	"github.com/zhiguang/zhiguang-go/services/gateway/internal/middleware"
	"github.com/zhiguang/zhiguang-go/services/gateway/internal/srv"
	authclient "github.com/zhiguang/zhiguang-go/services/user/rpc/client/auth"
	userclient "github.com/zhiguang/zhiguang-go/services/user/rpc/client/user"
	userpb "github.com/zhiguang/zhiguang-go/services/user/rpc/user"
)

func registerIdentityRoutes(r *gin.Engine, sc *srv.ServiceContext) {
	auth := r.Group("/api/v1/auth")
	{
		auth.POST("/send-code", sendCode(sc))
		auth.POST("/register", register(sc))
		auth.POST("/login", login(sc))
		auth.POST("/token/refresh", refresh(sc))
		auth.POST("/password/reset", passwordReset(sc))
		auth.POST("/logout", middleware.RequiredAuth(sc.AuthRpc), logout(sc))
		auth.GET("/me", me(sc))
	}

	profile := r.Group("/api/v1/profile", middleware.RequiredAuth(sc.AuthRpc))
	{
		profile.GET("/me", getProfileMe(sc))
		profile.PATCH("/", patchProfile(sc))
		profile.POST("/avatar", uploadAvatar(sc))
	}
}

func sendCode(sc *srv.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Scene      string `json:"scene"`
			Identifier string `json:"identifier"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			gh.WriteError(c, errorx.New(errorx.CodeBadRequest, err.Error()))
			return
		}
		resp, err := sc.AuthRpc.SendCode(c.Request.Context(), &authclient.SendCodeReq{
			Scene: req.Scene, Identifier: req.Identifier,
		})
		if err != nil {
			gh.WriteError(c, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"cooldownSeconds": resp.CooldownSeconds, "expireSeconds": resp.ExpireSeconds})
	}
}

func register(sc *srv.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Identifier string `json:"identifier"`
			Password   string `json:"password"`
			Code       string `json:"code"`
			Nickname   string `json:"nickname"`
			AgreeTerms bool   `json:"agreeTerms"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			gh.WriteError(c, errorx.New(errorx.CodeBadRequest, err.Error()))
			return
		}
		resp, err := sc.AuthRpc.Register(c.Request.Context(), &authclient.RegisterReq{
			Identifier: req.Identifier,
			Password:   req.Password,
			Code:       req.Code,
			Nickname:   req.Nickname,
			AgreeTerms: req.AgreeTerms,
			Ip:         c.ClientIP(),
			UserAgent:  c.Request.UserAgent(),
		})
		if err != nil {
			gh.WriteError(c, err)
			return
		}
		c.JSON(http.StatusOK, toAuthResp(resp))
	}
}

func login(sc *srv.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Identifier string `json:"identifier"`
			Password   string `json:"password"`
			Code       string `json:"code"`
			Channel    string `json:"channel"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			gh.WriteError(c, errorx.New(errorx.CodeBadRequest, err.Error()))
			return
		}
		resp, err := sc.AuthRpc.Login(c.Request.Context(), &authclient.LoginReq{
			Identifier: req.Identifier,
			Password:   req.Password,
			Code:       req.Code,
			Channel:    req.Channel,
			Ip:         c.ClientIP(),
			UserAgent:  c.Request.UserAgent(),
		})
		if err != nil {
			gh.WriteError(c, err)
			return
		}
		c.JSON(http.StatusOK, toAuthResp(resp))
	}
}

func refresh(sc *srv.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			RefreshToken string `json:"refreshToken"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			gh.WriteError(c, errorx.New(errorx.CodeBadRequest, err.Error()))
			return
		}
		resp, err := sc.AuthRpc.Refresh(c.Request.Context(), &authclient.RefreshReq{RefreshToken: req.RefreshToken})
		if err != nil {
			gh.WriteError(c, err)
			return
		}
		c.JSON(http.StatusOK, toAuthResp(resp))
	}
}

func passwordReset(sc *srv.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Identifier  string `json:"identifier"`
			Code        string `json:"code"`
			NewPassword string `json:"newPassword"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			gh.WriteError(c, errorx.New(errorx.CodeBadRequest, err.Error()))
			return
		}
		_, err := sc.AuthRpc.PasswordReset(c.Request.Context(), &authclient.PasswordResetReq{
			Identifier: req.Identifier, Code: req.Code, NewPassword: req.NewPassword,
		})
		if err != nil {
			gh.WriteError(c, err)
			return
		}
		c.Status(http.StatusNoContent)
	}
}

func logout(sc *srv.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			RefreshToken string `json:"refreshToken"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			gh.WriteError(c, errorx.New(errorx.CodeBadRequest, err.Error()))
			return
		}
		_, err := sc.AuthRpc.RevokeRefresh(c.Request.Context(), &authclient.RevokeRefreshReq{RefreshToken: req.RefreshToken})
		if err != nil {
			gh.WriteError(c, err)
			return
		}
		c.Status(http.StatusNoContent)
	}
}

func me(sc *srv.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := strings.TrimSpace(strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer "))
		resp, err := sc.AuthRpc.VerifyToken(c.Request.Context(), &authclient.VerifyTokenReq{AccessToken: token})
		if err != nil {
			gh.WriteError(c, err)
			return
		}
		if resp == nil || !resp.Valid {
			gh.WriteError(c, errorx.New(errorx.CodeUnauthorized, "invalid access token"))
			return
		}
		userResp, err := sc.UserRpc.GetById(c.Request.Context(), &userclient.GetByIdReq{Id: resp.UserId})
		if err != nil || userResp == nil || userResp.User == nil {
			c.JSON(http.StatusOK, gin.H{"id": resp.UserId, "nickname": resp.Nickname})
			return
		}
		u := userResp.User
		c.JSON(http.StatusOK, gin.H{
			"id": u.Id, "nickname": u.Nickname, "avatar": u.Avatar, "phone": u.Phone, "email": u.Email,
			"zgId": u.ZgId, "zhId": u.ZgId, "birthday": u.Birthday, "school": u.School,
			"bio": u.Bio, "gender": u.Gender, "tagsJson": u.TagsJson, "tagJson": u.TagsJson,
		})
	}
}

func getProfileMe(sc *srv.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, ok := ctxdata.GetUserId(c.Request.Context())
		if !ok {
			gh.WriteError(c, errorx.New(errorx.CodeUnauthorized, "missing user id"))
			return
		}
		resp, err := sc.UserRpc.GetById(c.Request.Context(), &userclient.GetByIdReq{Id: uid})
		if err != nil {
			gh.WriteError(c, err)
			return
		}
		c.JSON(http.StatusOK, toProfileResp(resp.User))
	}
}

func patchProfile(sc *srv.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, ok := ctxdata.GetUserId(c.Request.Context())
		if !ok {
			gh.WriteError(c, errorx.New(errorx.CodeUnauthorized, "missing user id"))
			return
		}
		var req struct {
			Nickname *string `json:"nickname"`
			Bio      *string `json:"bio"`
			Gender   *string `json:"gender"`
			Birthday *string `json:"birthday"`
			ZgId     *string `json:"zgId"`
			School   *string `json:"school"`
			TagsJson *string `json:"tagsJson"`
			TagJson  *string `json:"tagJson"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			gh.WriteError(c, errorx.New(errorx.CodeBadRequest, err.Error()))
			return
		}
		if req.ZgId != nil && *req.ZgId != "" {
			ex, err := sc.UserRpc.ExistsByZgIdExceptId(c.Request.Context(), &userclient.ExistsByZgIdExceptIdReq{ZgId: *req.ZgId, ExceptId: uid})
			if err != nil {
				gh.WriteError(c, err)
				return
			}
			if ex.Exists {
				gh.WriteError(c, errorx.New(errorx.CodeZgIdExists, "zgId already taken"))
				return
			}
		}
		in := &userclient.UpdateProfileReq{Id: uid}
		if req.Nickname != nil {
			in.Nickname, in.NicknameSet = *req.Nickname, true
		}
		if req.Bio != nil {
			in.Bio, in.BioSet = *req.Bio, true
		}
		if req.Gender != nil {
			in.Gender, in.GenderSet = *req.Gender, true
		}
		if req.Birthday != nil {
			in.Birthday, in.BirthdaySet = *req.Birthday, true
		}
		if req.ZgId != nil {
			in.ZgId, in.ZgIdSet = *req.ZgId, true
		}
		if req.School != nil {
			in.School, in.SchoolSet = *req.School, true
		}
		if req.TagsJson != nil {
			in.TagsJson, in.TagsJsonSet = *req.TagsJson, true
		} else if req.TagJson != nil {
			in.TagsJson, in.TagsJsonSet = *req.TagJson, true
		}
		resp, err := sc.UserRpc.UpdateProfile(c.Request.Context(), in)
		if err != nil {
			gh.WriteError(c, err)
			return
		}
		c.JSON(http.StatusOK, toProfileResp(resp.User))
	}
}

func uploadAvatar(sc *srv.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, ok := ctxdata.GetUserId(c.Request.Context())
		if !ok {
			gh.WriteError(c, errorx.New(errorx.CodeUnauthorized, "missing user id"))
			return
		}
		if err := c.Request.ParseMultipartForm(5 << 20); err != nil {
			gh.WriteError(c, errorx.Wrap(errorx.CodeBadRequest, "parse multipart failed", err))
			return
		}
		file, header, err := c.Request.FormFile("file")
		if err != nil {
			gh.WriteError(c, errorx.Wrap(errorx.CodeBadRequest, "missing form field 'file'", err))
			return
		}
		defer file.Close()
		if header.Size > 5<<20 {
			gh.WriteError(c, errorx.New(errorx.CodeBadRequest, "avatar exceeds 5MB"))
			return
		}
		contentType := header.Header.Get("Content-Type")
		if !strings.HasPrefix(contentType, "image/") {
			gh.WriteError(c, errorx.New(errorx.CodeBadRequest, "avatar must be an image"))
			return
		}
		ext := strings.TrimPrefix(sc.Oss.AvatarFolder(), sc.Oss.AvatarFolder())
		_ = ext
		key := fmt.Sprintf("%s/%d-%s", sc.Oss.AvatarFolder(), uid, header.Filename)
		if _, err := sc.Oss.PutObject(c.Request.Context(), key, file, contentType); err != nil {
			gh.WriteError(c, err)
			return
		}
		url := sc.Oss.BuildContentUrl(key)
		_, err = sc.UserRpc.UpdateProfile(c.Request.Context(), &userclient.UpdateProfileReq{
			Id: uid, Avatar: url, AvatarSet: true,
		})
		if err != nil {
			gh.WriteError(c, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"url": url, "avatar": url, "objectKey": key})
	}
}

func toAuthResp(resp *authclient.AuthResp) gin.H {
	u := resp.User
	t := resp.Token
	return gin.H{
		"user":  gin.H{"id": u.Id, "nickname": u.Nickname, "avatar": u.Avatar, "phone": u.Phone, "zgId": u.ZgId, "birthday": u.Birthday, "school": u.School, "bio": u.Bio, "gender": u.Gender, "tagsJson": u.TagsJson},
		"token": gin.H{"accessToken": t.AccessToken, "accessExpiresAt": t.AccessExpiresAt, "accessTokenExpiresAt": t.AccessExpiresAt, "refreshToken": t.RefreshToken, "refreshExpiresAt": t.RefreshExpiresAt, "refreshTokenExpiresAt": t.RefreshExpiresAt, "refreshTokenId": t.RefreshTokenId},
	}
}

func toProfileResp(u *userpb.UserInfo) gin.H {
	return gin.H{"id": u.Id, "nickname": u.Nickname, "avatar": u.Avatar, "bio": u.Bio, "zgId": u.ZgId, "gender": u.Gender, "birthday": u.Birthday, "school": u.School, "phone": u.Phone, "email": u.Email, "tagsJson": u.TagsJson, "tagJson": u.TagsJson}
}
