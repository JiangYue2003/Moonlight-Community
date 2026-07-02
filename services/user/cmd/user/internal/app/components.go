package app

import (
	"context"

	storagerpcapp "github.com/zhiguang/zhiguang-go/services/storage/rpc/app"
	userrpcapp "github.com/zhiguang/zhiguang-go/services/user/rpc/app"
)

type userComponent struct {
	cfg userrpcapp.Config
}

func NewUserComponent(cfg userrpcapp.Config) Component {
	return &userComponent{cfg: cfg}
}

func (c *userComponent) Name() string { return "user-rpc" }

func (c *userComponent) Run(ctx context.Context) error {
	return userrpcapp.Run(ctx, c.cfg)
}

type storageComponent struct {
	cfg storagerpcapp.Config
}

func NewStorageComponent(cfg storagerpcapp.Config) Component {
	return &storageComponent{cfg: cfg}
}

func (c *storageComponent) Name() string { return "storage-rpc" }

func (c *storageComponent) Run(ctx context.Context) error {
	return storagerpcapp.Run(ctx, c.cfg)
}
