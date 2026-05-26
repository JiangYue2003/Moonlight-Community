// Package ossx 是阿里云 OSS Go SDK 的薄封装，提供：
//   - 预签名 PUT URL（前端直传）
//   - 头像后端中转上传
//   - 内容公开 URL 构造（PublicDomain 优先于 endpoint）
package ossx

// Config OSS 客户端参数；与原 Java OssProperties 字段对齐。
type Config struct {
	Endpoint          string `json:",default=oss-cn-shanghai.aliyuncs.com"`
	PublicDomain      string `json:",optional"` // 自定义 CDN 域名，可空
	Bucket            string
	AccessKeyId       string
	AccessKeySecret   string
	AvatarFolder      string `json:",default=avatars"`
	PresignExpiresSec int64  `json:",default=600"`
}
