// Package responsex 提供 HTTP 接口的统一响应封装。
package responsex

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/zhiguang/zhiguang-go/pkg/errorx"
)

// Body 统一响应结构 {code, message, data}。
type Body struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Ok 返回 200 + 成功 body。
func Ok(w http.ResponseWriter, data interface{}) {
	write(w, http.StatusOK, Body{Code: "OK", Message: "ok", Data: data})
}

// Fail 根据 error 写出 4xx/5xx；BizError 用其 Code，否则统一 INTERNAL_ERROR。
func Fail(w http.ResponseWriter, err error) {
	if err == nil {
		Ok(w, nil)
		return
	}
	if be := new(errorx.BizError); errors.As(err, &be) {
		write(w, mapStatus(be.Code), Body{Code: be.Code, Message: be.Message})
		return
	}
	write(w, http.StatusInternalServerError, Body{
		Code:    errorx.CodeInternalError,
		Message: err.Error(),
	})
}

func write(w http.ResponseWriter, status int, b Body) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(b)
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
