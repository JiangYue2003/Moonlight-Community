package handler

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zhiguang/zhiguang-go/common/ctxdata"
	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	gh "github.com/zhiguang/zhiguang-go/services/gateway/internal/httpx"
	"github.com/zhiguang/zhiguang-go/services/gateway/internal/middleware"
	"github.com/zhiguang/zhiguang-go/services/gateway/internal/srv"
	counterclient "github.com/zhiguang/zhiguang-go/services/counter/rpc/client/counter"
	usercounterclient "github.com/zhiguang/zhiguang-go/services/counter/rpc/client/usercounter"
	knowpostclient "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/client/knowpost"
	llmclient "github.com/zhiguang/zhiguang-go/services/llm/rpc/client/llm"
	relationclient "github.com/zhiguang/zhiguang-go/services/relation/rpc/client/relation"
	searchclient "github.com/zhiguang/zhiguang-go/services/search/rpc/client/search"
	storageclient "github.com/zhiguang/zhiguang-go/services/storage/rpc/client/storage"
	authclient "github.com/zhiguang/zhiguang-go/services/user/rpc/client/auth"
	userclient "github.com/zhiguang/zhiguang-go/services/user/rpc/client/user"
	userpb "github.com/zhiguang/zhiguang-go/services/user/rpc/user"
)

func NewEngine(sc *srv.ServiceContext) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

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

	storage := r.Group("/api/v1/storage", middleware.RequiredAuth(sc.AuthRpc))
	{
		storage.POST("/presign", presign(sc))
	}

	publicKnowpost := r.Group("/api/v1/knowposts", middleware.OptionalAuth(sc.AuthRpc))
	{
		publicKnowpost.GET("/feed", getPublicFeed(sc))
		publicKnowpost.GET("/detail/:id", getDetail(sc))
		publicKnowpost.POST("/description/suggest", suggestDescription(sc))
		publicKnowpost.GET("/:id/qa/stream", qaCompatStream(sc))
	}

	privateKnowpost := r.Group("/api/v1/knowposts", middleware.RequiredAuth(sc.AuthRpc))
	{
		privateKnowpost.POST("/drafts", createDraft(sc))
		privateKnowpost.POST("/:id/content/confirm", confirmContent(sc))
		privateKnowpost.PATCH("/:id", patchMetadata(sc))
		privateKnowpost.POST("/:id/publish", publish(sc))
		privateKnowpost.PATCH("/:id/top", updateTop(sc))
		privateKnowpost.PATCH("/:id/visibility", updateVisibility(sc))
		privateKnowpost.DELETE("/:id", deleteKnowpost(sc))
		privateKnowpost.GET("/mine", getMyFeed(sc))
		privateKnowpost.POST("/:id/reindex", reindex(sc))
		privateKnowpost.POST("/:id/rag/reindex", reindex(sc))
	}

	relationPublic := r.Group("/api/v1/relation", middleware.OptionalAuth(sc.AuthRpc))
	{
		relationPublic.GET("/status", relationStatus(sc))
	}
	relationPrivate := r.Group("/api/v1/relation", middleware.RequiredAuth(sc.AuthRpc))
	{
		relationPrivate.POST("/follow", follow(sc))
		relationPrivate.POST("/unfollow", unfollow(sc))
		relationPrivate.GET("/following", listFollowing(sc))
		relationPrivate.GET("/followers", listFollowers(sc))
		relationPrivate.GET("/counter", relationCounter(sc))
	}

	action := r.Group("/api/v1/action", middleware.RequiredAuth(sc.AuthRpc))
	{
		action.POST("/like", toggleMetric(sc, "like", true))
		action.POST("/unlike", toggleMetric(sc, "like", false))
		action.POST("/fav", toggleMetric(sc, "fav", true))
		action.POST("/unfav", toggleMetric(sc, "fav", false))
	}
	r.GET("/api/v1/counter/:etype/:eid", getCounts(sc))

	search := r.Group("/api/v1/search", middleware.OptionalAuth(sc.AuthRpc))
	{
		search.GET("/", searchPosts(sc))
		search.GET("/suggest", suggestSearch(sc))
	}

	llmPrivate := r.Group("/api/v1/llm", middleware.RequiredAuth(sc.AuthRpc))
	{
		llmPrivate.POST("/describe", llmDescribe(sc))
	}
	llmCompat := r.Group("/api/v1/llm", middleware.OptionalAuth(sc.AuthRpc))
	{
		llmCompat.GET("/qa/stream", llmQaStream(sc))
	}

	return r
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
			Password: req.Password,
			Code: req.Code,
			Nickname: req.Nickname,
			AgreeTerms: req.AgreeTerms,
			Ip: c.ClientIP(),
			UserAgent: c.Request.UserAgent(),
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
			Password: req.Password,
			Code: req.Code,
			Channel: req.Channel,
			Ip: c.ClientIP(),
			UserAgent: c.Request.UserAgent(),
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
			Identifier string `json:"identifier"`
			Code string `json:"code"`
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
		var req struct{ RefreshToken string `json:"refreshToken"` }
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
			Bio *string `json:"bio"`
			Gender *string `json:"gender"`
			Birthday *string `json:"birthday"`
			ZgId *string `json:"zgId"`
			School *string `json:"school"`
			TagsJson *string `json:"tagsJson"`
			TagJson *string `json:"tagJson"`
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

func presign(sc *srv.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, ok := ctxdata.GetUserId(c.Request.Context())
		if !ok {
			gh.WriteError(c, errorx.New(errorx.CodeUnauthorized, "missing user id"))
			return
		}
		var req struct {
			Scene string `json:"scene"`
			PostId string `json:"postId"`
			ContentType string `json:"contentType"`
			Ext string `json:"ext"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			gh.WriteError(c, errorx.New(errorx.CodeBadRequest, err.Error()))
			return
		}
		resp, err := sc.StorageRpc.Presign(c.Request.Context(), &storageclient.PresignReq{
			UserId: uid, Scene: req.Scene, PostId: req.PostId, ContentType: req.ContentType, Ext: req.Ext,
		})
		if err != nil {
			gh.WriteError(c, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"url": resp.Url, "objectKey": resp.ObjectKey, "expiresIn": resp.ExpiresIn,
			"headers": resp.Headers, "contentUrl": resp.ContentUrl,
		})
	}
}

func createDraft(sc *srv.ServiceContext) gin.HandlerFunc { return func(c *gin.Context) { uid,_ := ctxdata.GetUserId(c.Request.Context()); resp, err := sc.KnowPostRpc.CreateDraft(c.Request.Context(), &knowpostclient.CreateDraftReq{CreatorId: uid}); if err != nil { gh.WriteError(c, err); return }; c.JSON(http.StatusOK, gin.H{"id": resp.Id}) } }
func confirmContent(sc *srv.ServiceContext) gin.HandlerFunc { return func(c *gin.Context) { uid,_ := ctxdata.GetUserId(c.Request.Context()); id, err := parsePathInt64(c, "id"); if err != nil { gh.WriteError(c, err); return }; var req struct{ ObjectKey string `json:"objectKey"`; Etag string `json:"etag"`; Size int64 `json:"size"`; Sha256 string `json:"sha256"` }; if err := c.ShouldBindJSON(&req); err != nil { gh.WriteError(c, errorx.New(errorx.CodeBadRequest, err.Error())); return }; _, err = sc.KnowPostRpc.ConfirmContent(c.Request.Context(), &knowpostclient.ConfirmContentReq{Id: id, CreatorId: uid, ObjectKey: req.ObjectKey, Etag: req.Etag, Size: req.Size, Sha256: req.Sha256}); if err != nil { gh.WriteError(c, err); return }; c.Status(http.StatusNoContent) } }
func patchMetadata(sc *srv.ServiceContext) gin.HandlerFunc { return func(c *gin.Context) { uid,_ := ctxdata.GetUserId(c.Request.Context()); id, err := parsePathInt64(c, "id"); if err != nil { gh.WriteError(c, err); return }; var req struct{ Title *string `json:"title"`; Description *string `json:"description"`; TagId *int64 `json:"tagId"`; Tags []string `json:"tags"`; TagsSet bool `json:"tagsSet"`; ImgUrls []string `json:"imgUrls"`; ImgUrlsSet bool `json:"imgUrlsSet"`; Visible *string `json:"visible"`; IsTop *bool `json:"isTop"` }; if err := c.ShouldBindJSON(&req); err != nil { gh.WriteError(c, errorx.New(errorx.CodeBadRequest, err.Error())); return }; in := &knowpostclient.PatchMetadataReq{Id: id, CreatorId: uid}; if req.Title != nil { in.Title, in.TitleSet = *req.Title, true }; if req.Description != nil { in.Description, in.DescriptionSet = *req.Description, true }; if req.TagId != nil { in.TagId, in.TagIdSet = *req.TagId, true }; in.Tags, in.TagsSet = req.Tags, req.TagsSet; in.ImgUrls, in.ImgUrlsSet = req.ImgUrls, req.ImgUrlsSet; if req.Visible != nil { in.Visible, in.VisibleSet = *req.Visible, true }; if req.IsTop != nil { in.IsTop, in.IsTopSet = *req.IsTop, true }; resp, err := sc.KnowPostRpc.PatchMetadata(c.Request.Context(), in); if err != nil { gh.WriteError(c, err); return }; c.JSON(http.StatusOK, toKnowPostDetail(resp)) } }
func publish(sc *srv.ServiceContext) gin.HandlerFunc { return func(c *gin.Context) { uid,_ := ctxdata.GetUserId(c.Request.Context()); id, err := parsePathInt64(c, "id"); if err != nil { gh.WriteError(c, err); return }; resp, err := sc.KnowPostRpc.Publish(c.Request.Context(), &knowpostclient.PublishReq{Id: id, CreatorId: uid}); if err != nil { gh.WriteError(c, err); return }; c.JSON(http.StatusOK, toKnowPostDetail(resp)) } }
func updateTop(sc *srv.ServiceContext) gin.HandlerFunc { return func(c *gin.Context) { uid,_ := ctxdata.GetUserId(c.Request.Context()); id, err := parsePathInt64(c, "id"); if err != nil { gh.WriteError(c, err); return }; var req struct{ IsTop bool `json:"isTop"` }; if err := c.ShouldBindJSON(&req); err != nil { gh.WriteError(c, errorx.New(errorx.CodeBadRequest, err.Error())); return }; _, err = sc.KnowPostRpc.UpdateTop(c.Request.Context(), &knowpostclient.UpdateTopReq{Id: id, CreatorId: uid, IsTop: req.IsTop}); if err != nil { gh.WriteError(c, err); return }; c.Status(http.StatusNoContent) } }
func updateVisibility(sc *srv.ServiceContext) gin.HandlerFunc { return func(c *gin.Context) { uid,_ := ctxdata.GetUserId(c.Request.Context()); id, err := parsePathInt64(c, "id"); if err != nil { gh.WriteError(c, err); return }; var req struct{ Visible string `json:"visible"` }; if err := c.ShouldBindJSON(&req); err != nil { gh.WriteError(c, errorx.New(errorx.CodeBadRequest, err.Error())); return }; _, err = sc.KnowPostRpc.UpdateVisibility(c.Request.Context(), &knowpostclient.UpdateVisibilityReq{Id: id, CreatorId: uid, Visible: req.Visible}); if err != nil { gh.WriteError(c, err); return }; c.Status(http.StatusNoContent) } }
func deleteKnowpost(sc *srv.ServiceContext) gin.HandlerFunc { return func(c *gin.Context) { uid,_ := ctxdata.GetUserId(c.Request.Context()); id, err := parsePathInt64(c, "id"); if err != nil { gh.WriteError(c, err); return }; _, err = sc.KnowPostRpc.Delete(c.Request.Context(), &knowpostclient.DeleteReq{Id: id, CreatorId: uid}); if err != nil { gh.WriteError(c, err); return }; c.Status(http.StatusNoContent) } }
func getPublicFeed(sc *srv.ServiceContext) gin.HandlerFunc { return func(c *gin.Context) { page := int32(queryInt(c, "page", 1)); size := int32(queryInt(c, "size", 20)); resp, err := sc.KnowPostRpc.GetPublicFeed(c.Request.Context(), &knowpostclient.GetPublicFeedReq{Page: page, Size: size}); if err != nil { gh.WriteError(c, err); return }; c.JSON(http.StatusOK, toFeedPage(resp)) } }
func getMyFeed(sc *srv.ServiceContext) gin.HandlerFunc { return func(c *gin.Context) { uid,_ := ctxdata.GetUserId(c.Request.Context()); page := int32(queryInt(c, "page", 1)); size := int32(queryInt(c, "size", 20)); resp, err := sc.KnowPostRpc.GetMyFeed(c.Request.Context(), &knowpostclient.GetMyFeedReq{CreatorId: uid, Page: page, Size: size}); if err != nil { gh.WriteError(c, err); return }; c.JSON(http.StatusOK, toFeedPage(resp)) } }
func getDetail(sc *srv.ServiceContext) gin.HandlerFunc { return func(c *gin.Context) { id, err := parsePathInt64(c, "id"); if err != nil { gh.WriteError(c, err); return }; viewer,_ := ctxdata.GetUserId(c.Request.Context()); resp, err := sc.KnowPostRpc.GetDetail(c.Request.Context(), &knowpostclient.GetDetailReq{Id: id, ViewerId: viewer}); if err != nil { gh.WriteError(c, err); return }; c.JSON(http.StatusOK, toKnowPostDetail(resp)) } }
func reindex(sc *srv.ServiceContext) gin.HandlerFunc { return func(c *gin.Context) { uid,_ := ctxdata.GetUserId(c.Request.Context()); id, err := parsePathInt64(c, "id"); if err != nil { gh.WriteError(c, err); return }; _, err = sc.KnowPostRpc.Reindex(c.Request.Context(), &knowpostclient.ReindexReq{Id: id, CreatorId: uid}); if err != nil { gh.WriteError(c, err); return }; c.Status(http.StatusNoContent) } }

func follow(sc *srv.ServiceContext) gin.HandlerFunc { return func(c *gin.Context) { uid,_ := ctxdata.GetUserId(c.Request.Context()); var req struct{ ToUserId int64 `json:"toUserId"` }; if err := c.ShouldBindJSON(&req); err != nil { gh.WriteError(c, errorx.New(errorx.CodeBadRequest, err.Error())); return }; _, err := sc.RelationRpc.Follow(c.Request.Context(), &relationclient.FollowReq{FromUserId: uid, ToUserId: req.ToUserId}); if err != nil { gh.WriteError(c, err); return }; c.Status(http.StatusNoContent) } }
func unfollow(sc *srv.ServiceContext) gin.HandlerFunc { return func(c *gin.Context) { uid,_ := ctxdata.GetUserId(c.Request.Context()); var req struct{ ToUserId int64 `json:"toUserId"` }; if err := c.ShouldBindJSON(&req); err != nil { gh.WriteError(c, errorx.New(errorx.CodeBadRequest, err.Error())); return }; _, err := sc.RelationRpc.Unfollow(c.Request.Context(), &relationclient.UnfollowReq{FromUserId: uid, ToUserId: req.ToUserId}); if err != nil { gh.WriteError(c, err); return }; c.Status(http.StatusNoContent) } }
func relationStatus(sc *srv.ServiceContext) gin.HandlerFunc { return func(c *gin.Context) { toUserId := int64(queryInt(c, "toUserId", 0)); fromUserId,_ := ctxdata.GetUserId(c.Request.Context()); resp, err := sc.RelationRpc.Status(c.Request.Context(), &relationclient.StatusReq{FromUserId: fromUserId, ToUserId: toUserId}); if err != nil { gh.WriteError(c, err); return }; c.JSON(http.StatusOK, gin.H{"following": resp.Following, "followedBy": resp.FollowedBy, "mutual": resp.Mutual}) } }
func listFollowing(sc *srv.ServiceContext) gin.HandlerFunc { return func(c *gin.Context) { userId := int64(queryInt(c, "userId", 0)); if userId == 0 { userId,_ = ctxdata.GetUserId(c.Request.Context()) }; limit := int32(queryInt(c, "limit", 20)); offset := int32(queryInt(c, "offset", 0)); cursor := int64(queryInt(c, "cursor", 0)); resp, err := sc.RelationRpc.ListFollowing(c.Request.Context(), &relationclient.ListReq{UserId: userId, Limit: limit, Offset: offset, Cursor: cursor}); if err != nil { gh.WriteError(c, err); return }; c.JSON(http.StatusOK, toRelationList(resp)) } }
func listFollowers(sc *srv.ServiceContext) gin.HandlerFunc { return func(c *gin.Context) { userId := int64(queryInt(c, "userId", 0)); if userId == 0 { userId,_ = ctxdata.GetUserId(c.Request.Context()) }; limit := int32(queryInt(c, "limit", 20)); offset := int32(queryInt(c, "offset", 0)); cursor := int64(queryInt(c, "cursor", 0)); resp, err := sc.RelationRpc.ListFollowers(c.Request.Context(), &relationclient.ListReq{UserId: userId, Limit: limit, Offset: offset, Cursor: cursor}); if err != nil { gh.WriteError(c, err); return }; c.JSON(http.StatusOK, toRelationList(resp)) } }
func relationCounter(sc *srv.ServiceContext) gin.HandlerFunc { return func(c *gin.Context) { userId := int64(queryInt(c, "userId", 0)); if userId == 0 { userId,_ = ctxdata.GetUserId(c.Request.Context()) }; resp, err := sc.UserCounterRpc.GetUserSnapshot(c.Request.Context(), &usercounterclient.GetUserSnapshotReq{UserId: userId}); if err != nil { gh.WriteError(c, err); return }; snap := resp.GetSnapshot(); c.JSON(http.StatusOK, gin.H{"followings": snap.GetFollowings(), "followers": snap.GetFollowers(), "posts": snap.GetPosts(), "likedPosts": int64(0), "favedPosts": int64(0), "likesReceived": snap.GetLikesReceived()}) } }

func toggleMetric(sc *srv.ServiceContext, metric string, add bool) gin.HandlerFunc { return func(c *gin.Context) { uid,_ := ctxdata.GetUserId(c.Request.Context()); var req struct{ EntityType string `json:"entityType"`; EntityId string `json:"entityId"` }; if err := c.ShouldBindJSON(&req); err != nil { gh.WriteError(c, errorx.New(errorx.CodeBadRequest, err.Error())); return }; resp, err := sc.CounterRpc.Toggle(c.Request.Context(), &counterclient.ToggleReq{EntityType: req.EntityType, EntityId: req.EntityId, Metric: metric, UserId: uid, Add: add}); if err != nil { gh.WriteError(c, err); return }; key := map[string]string{"like":"liked","fav":"faved"}[metric]; c.JSON(http.StatusOK, gin.H{"changed": resp.Changed, key: add}) } }
func getCounts(sc *srv.ServiceContext) gin.HandlerFunc { return func(c *gin.Context) { etype := c.Param("etype"); eid := c.Param("eid"); resp, err := sc.CounterRpc.GetCounts(c.Request.Context(), &counterclient.GetCountsReq{EntityType: etype, EntityId: eid}); if err != nil { gh.WriteError(c, err); return }; c.JSON(http.StatusOK, gin.H{"entityType": etype, "entityId": eid, "counts": resp.Counts}) } }

func searchPosts(sc *srv.ServiceContext) gin.HandlerFunc { return func(c *gin.Context) { viewer,_ := ctxdata.GetUserId(c.Request.Context()); size := int32(queryInt(c, "size", 20)); resp, err := sc.SearchRpc.Search(c.Request.Context(), &searchclient.SearchReq{Q: c.Query("q"), Size: size, Tags: c.Query("tags"), After: c.Query("after"), ViewerId: viewer}); if err != nil { gh.WriteError(c, err); return }; items := make([]gin.H, 0, len(resp.Items)); for _, item := range resp.Items { items = append(items, gin.H{"contentId": item.Id, "contentType": "knowpost", "title": item.Title, "description": item.Description, "snippet": item.Description, "tags": item.Tags, "authorId": item.AuthorId, "authorNickname": item.AuthorNickname, "authorAvatar": item.AuthorAvatar, "likeCount": item.LikeCount, "favoriteCount": item.FavoriteCount, "viewCount": 0, "imgUrls": []string{item.CoverImage}, "isTop": item.IsTop}) }; c.JSON(http.StatusOK, gin.H{"items": items, "nextAfter": resp.NextAfter, "hasMore": resp.HasMore}) } }
func suggestSearch(sc *srv.ServiceContext) gin.HandlerFunc { return func(c *gin.Context) { size := int32(queryInt(c, "size", 10)); resp, err := sc.SearchRpc.Suggest(c.Request.Context(), &searchclient.SuggestReq{Prefix: c.Query("prefix"), Size: size}); if err != nil { gh.WriteError(c, err); return }; c.JSON(http.StatusOK, gin.H{"items": resp.Items}) } }

func llmDescribe(sc *srv.ServiceContext) gin.HandlerFunc { return func(c *gin.Context) { uid,_ := ctxdata.GetUserId(c.Request.Context()); var req struct{ Body string `json:"body"`; Content string `json:"content"` }; if err := c.ShouldBindJSON(&req); err != nil { gh.WriteError(c, errorx.New(errorx.CodeBadRequest, err.Error())); return }; resp, err := sc.LlmRpc.Describe(c.Request.Context(), &llmclient.DescribeReq{UserId: uid, Body: req.Body, Content: req.Content}); if err != nil { gh.WriteError(c, err); return }; c.JSON(http.StatusOK, gin.H{"description": resp.Description}) } }
func suggestDescription(sc *srv.ServiceContext) gin.HandlerFunc { return llmDescribe(sc) }

func llmQaStream(sc *srv.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		postID := int64(queryInt(c, "postId", 0))
		if postID <= 0 {
			gh.WriteError(c, errorx.New(errorx.CodeBadRequest, "postId required"))
			return
		}
		streamQa(sc, c, postID)
	}
}

func qaCompatStream(sc *srv.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		postID, err := parsePathInt64(c, "id")
		if err != nil {
			gh.WriteError(c, err)
			return
		}
		streamQa(sc, c, postID)
	}
}

func streamQa(sc *srv.ServiceContext, c *gin.Context, postID int64) {
	userID, _ := ctxdata.GetUserId(c.Request.Context())
	topK := int32(queryInt(c, "topK", 5))
	maxTokens := int32(queryInt(c, "maxTokens", 1024))
	stream, err := sc.LlmRpc.QaStream(c.Request.Context(), &llmclient.QaStreamReq{
		UserId: userID, PostId: postID, Question: c.Query("question"), TopK: topK, MaxTokens: maxTokens,
	})
	if err != nil {
		gh.WriteError(c, err)
		return
	}
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Status(http.StatusOK)
	flusher, ok := c.Writer.(http.Flusher)
	if !ok {
		gh.WriteError(c, errorx.New(errorx.CodeInternalError, "streaming unsupported"))
		return
	}
	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			return
		}
		if _, err := fmt.Fprintf(c.Writer, "data: %s\n\n", chunk.Data); err != nil {
			return
		}
		flusher.Flush()
		if chunk.Done {
			return
		}
	}
}

func toAuthResp(resp *authclient.AuthResp) gin.H {
	u := resp.User
	t := resp.Token
	return gin.H{
		"user": gin.H{"id": u.Id, "nickname": u.Nickname, "avatar": u.Avatar, "phone": u.Phone, "zgId": u.ZgId, "birthday": u.Birthday, "school": u.School, "bio": u.Bio, "gender": u.Gender, "tagsJson": u.TagsJson},
		"token": gin.H{"accessToken": t.AccessToken, "accessExpiresAt": t.AccessExpiresAt, "accessTokenExpiresAt": t.AccessExpiresAt, "refreshToken": t.RefreshToken, "refreshExpiresAt": t.RefreshExpiresAt, "refreshTokenExpiresAt": t.RefreshExpiresAt, "refreshTokenId": t.RefreshTokenId},
	}
}

func toProfileResp(u *userpb.UserInfo) gin.H { return gin.H{"id": u.Id, "nickname": u.Nickname, "avatar": u.Avatar, "bio": u.Bio, "zgId": u.ZgId, "gender": u.Gender, "birthday": u.Birthday, "school": u.School, "phone": u.Phone, "email": u.Email, "tagsJson": u.TagsJson, "tagJson": u.TagsJson} }

func toKnowPostDetail(resp *knowpostclient.KnowPostDetail) gin.H {
	return gin.H{
		"id": resp.Id, "creatorId": resp.CreatorId, "title": resp.Title, "description": resp.Description, "tagId": resp.TagId,
		"tags": resp.Tags, "contentUrl": resp.ContentUrl, "contentObjectKey": resp.ContentObjectKey, "imgUrls": resp.ImgUrls,
		"visible": resp.Visible, "status": resp.Status, "type": resp.Type, "isTop": resp.IsTop,
		"createTime": resp.CreateTime, "updateTime": resp.UpdateTime, "publishTime": resp.PublishTime,
	}
}

func toFeedPage(resp *knowpostclient.FeedPage) gin.H {
	items := make([]gin.H, 0, len(resp.Items))
	for _, item := range resp.Items {
		items = append(items, gin.H{
			"id": item.Id, "creatorId": item.CreatorId, "title": item.Title, "description": item.Description,
			"contentUrl": item.ContentUrl, "tags": item.Tags, "imgUrls": item.ImgUrls, "visible": item.Visible,
			"isTop": item.IsTop, "publishTime": item.PublishTime,
		})
	}
	return gin.H{"items": items, "hasMore": resp.HasMore, "size": resp.Size, "page": resp.Page}
}

func toRelationList(resp *relationclient.ListResp) gin.H {
	items := make([]gin.H, 0, len(resp.Items))
	for _, item := range resp.Items {
		items = append(items, gin.H{"id": item.Id, "nickname": item.Nickname, "avatar": item.Avatar, "zgId": item.ZgId, "bio": item.Bio})
	}
	return gin.H{"items": items, "nextCursor": resp.NextCursor, "hasMore": resp.HasMore}
}

func parsePathInt64(c *gin.Context, name string) (int64, error) {
	v, err := strconv.ParseInt(c.Param(name), 10, 64)
	if err != nil {
		return 0, errorx.New(errorx.CodeBadRequest, "invalid "+name)
	}
	return v, nil
}

func queryInt(c *gin.Context, key string, def int) int {
	if s := c.Query(key); s != "" {
		if n, err := strconv.Atoi(s); err == nil {
			return n
		}
	}
	return def
}

var _ = context.Background
