package httpx

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhiguang/zhiguang-go/pkg/errorx"
)

func WriteError(c *gin.Context, err error) {
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"code": "OK", "message": "ok"})
		return
	}
	var be *errorx.BizError
	if errors.As(err, &be) {
		c.JSON(mapStatus(be.Code), gin.H{"code": be.Code, "message": be.Message})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"code": errorx.CodeInternalError, "message": err.Error()})
}

func WriteJSON(c *gin.Context, status int, data any) {
	c.JSON(status, data)
}

func mapStatus(code string) int {
	switch code {
	case errorx.CodeBadRequest,
		errorx.CodeIdentifierExists,
		errorx.CodeZgIdExists,
		errorx.CodeVerificationNotFound,
		errorx.CodeVerificationExpired,
		errorx.CodeVerificationMismatch,
		errorx.CodeVerificationTooMany,
		errorx.CodeVerificationCooldown,
		errorx.CodeVerificationDailyLimit,
		errorx.CodePasswordPolicyViolation,
		errorx.CodeTermsNotAccepted:
		return http.StatusBadRequest
	case errorx.CodeUnauthorized,
		errorx.CodeRefreshTokenInvalid,
		errorx.CodeInvalidCredentials:
		return http.StatusUnauthorized
	case errorx.CodeForbidden:
		return http.StatusForbidden
	case errorx.CodeNotFound, errorx.CodeIdentifierNotFound:
		return http.StatusNotFound
	case errorx.CodeRateLimited:
		return http.StatusTooManyRequests
	default:
		return http.StatusInternalServerError
	}
}
