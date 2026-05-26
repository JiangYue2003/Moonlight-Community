package ossx

import (
	"errors"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// Client 持有一个 OSS bucket handle，方法集合上挂在它身上。
type Client struct {
	cfg    Config
	cli    *oss.Client
	bucket *oss.Bucket
}

// New 构造客户端；不会做实际的网络调用。Bucket 不存在的错误在调用时才暴露。
func New(cfg Config) (*Client, error) {
	if cfg.Endpoint == "" || cfg.Bucket == "" {
		return nil, errors.New("ossx: Endpoint and Bucket required")
	}
	if cfg.AccessKeyId == "" || cfg.AccessKeySecret == "" {
		return nil, errors.New("ossx: AccessKey credentials required")
	}
	cli, err := oss.New(cfg.Endpoint, cfg.AccessKeyId, cfg.AccessKeySecret)
	if err != nil {
		return nil, err
	}
	bucket, err := cli.Bucket(cfg.Bucket)
	if err != nil {
		return nil, err
	}
	return &Client{cfg: cfg, cli: cli, bucket: bucket}, nil
}

// AvatarFolder 暴露给上层用于 ObjectKeyFor。
func (c *Client) AvatarFolder() string {
	if c.cfg.AvatarFolder == "" {
		return "avatars"
	}
	return c.cfg.AvatarFolder
}
