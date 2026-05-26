package ossx

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

// 与 Java 端 ext / objectKey 推导逻辑严格对齐。

const (
	SceneKnowPostContent = "knowpost_content"
	SceneKnowPostImage   = "knowpost_image"
	SceneAvatar          = "avatar"
)

// ExtFromContentType 推导扩展名。
//
//	若 explicitExt 非空，则优先使用（自动补 "." 前缀）；
//	否则按 scene + contentType 映射；映射不到时使用 default ".bin"/".img"。
func ExtFromContentType(contentType, scene, explicitExt string) string {
	if e := normalizeExt(explicitExt); e != "" {
		return e
	}
	ct := strings.ToLower(strings.TrimSpace(contentType))
	switch scene {
	case SceneKnowPostContent:
		switch ct {
		case "text/markdown":
			return ".md"
		case "text/html":
			return ".html"
		case "text/plain":
			return ".txt"
		case "application/json":
			return ".json"
		default:
			return ".bin"
		}
	case SceneKnowPostImage, SceneAvatar:
		switch ct {
		case "image/jpeg", "image/jpg":
			return ".jpg"
		case "image/png":
			return ".png"
		case "image/webp":
			return ".webp"
		default:
			return ".img"
		}
	default:
		return ".bin"
	}
}

// ObjectKeyFor 根据 scene 生成对象键。
//
//	knowpost_content → posts/{postId}/content{ext}
//	knowpost_image   → posts/{postId}/images/{yyyyMMdd UTC}/{rand8}{ext}
//	avatar           → {avatarFolder}/{userId}-{epochMilli}{ext}
func ObjectKeyFor(scene, postId string, userId int64, avatarFolder, ext string) string {
	switch scene {
	case SceneKnowPostContent:
		return fmt.Sprintf("posts/%s/content%s", postId, ext)
	case SceneKnowPostImage:
		date := time.Now().UTC().Format("20060102")
		return fmt.Sprintf("posts/%s/images/%s/%s%s", postId, date, randHex8(), ext)
	case SceneAvatar:
		folder := strings.TrimSuffix(avatarFolder, "/")
		if folder == "" {
			folder = "avatars"
		}
		return fmt.Sprintf("%s/%d-%d%s", folder, userId, time.Now().UnixMilli(), ext)
	default:
		return ""
	}
}

func normalizeExt(ext string) string {
	ext = strings.TrimSpace(ext)
	if ext == "" {
		return ""
	}
	if !strings.HasPrefix(ext, ".") {
		ext = "." + ext
	}
	return strings.ToLower(ext)
}

// randHex8 生成 8 个十六进制字符（与 Java UUID 取前 8 位等价的随机源）。
func randHex8() string {
	var b [4]byte
	_, _ = rand.Read(b[:])
	return hex.EncodeToString(b[:])
}
