package config

import (
	storagerpcapp "github.com/zhiguang/zhiguang-go/services/storage/rpc/app"
	userrpcapp "github.com/zhiguang/zhiguang-go/services/user/rpc/app"
)

// Config merges user-rpc and storage-rpc configurations.
type Config struct {
	User    userrpcapp.Config
	Storage storagerpcapp.Config
}
