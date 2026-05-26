// Package errorx 提供业务错误码与 BizError，用于跨服务统一错误传递。
package errorx

import (
	"errors"
	"fmt"
)

// 错误码与原 Java ErrorCode 枚举对齐。
const (
	CodeBadRequest              = "BAD_REQUEST"
	CodeInternalError           = "INTERNAL_ERROR"
	CodeIdentifierExists        = "IDENTIFIER_EXISTS"
	CodeIdentifierNotFound      = "IDENTIFIER_NOT_FOUND"
	CodeZgIdExists              = "ZGID_EXISTS"
	CodeVerificationNotFound    = "VERIFICATION_NOT_FOUND"
	CodeVerificationExpired     = "VERIFICATION_EXPIRED"
	CodeVerificationMismatch    = "VERIFICATION_MISMATCH"
	CodeVerificationTooMany     = "VERIFICATION_TOO_MANY_ATTEMPTS"
	CodeVerificationCooldown    = "VERIFICATION_COOLDOWN"
	CodeVerificationDailyLimit  = "VERIFICATION_DAILY_LIMIT"
	CodeInvalidCredentials      = "INVALID_CREDENTIALS"
	CodePasswordPolicyViolation = "PASSWORD_POLICY_VIOLATION"
	CodeTermsNotAccepted        = "TERMS_NOT_ACCEPTED"
	CodeRefreshTokenInvalid     = "REFRESH_TOKEN_INVALID"
	CodeUnauthorized            = "UNAUTHORIZED"
	CodeForbidden               = "FORBIDDEN"
	CodeNotFound                = "NOT_FOUND"
	CodeRateLimited             = "RATE_LIMITED"
)

// BizError 业务错误，包含可序列化的 Code 与人类可读的 Message。
type BizError struct {
	Code    string
	Message string
	Cause   error
}

func (e *BizError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %s: %v", e.Code, e.Message, e.Cause)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func (e *BizError) Unwrap() error { return e.Cause }

// New 构造无 cause 的业务错误。
func New(code, message string) *BizError {
	return &BizError{Code: code, Message: message}
}

// Wrap 在保留 cause 的同时构造业务错误。
func Wrap(code, message string, cause error) *BizError {
	return &BizError{Code: code, Message: message, Cause: cause}
}

// As 将任意 error 转为 *BizError；非业务错误返回 nil, false。
func As(err error) (*BizError, bool) {
	var be *BizError
	if errors.As(err, &be) {
		return be, true
	}
	return nil, false
}
