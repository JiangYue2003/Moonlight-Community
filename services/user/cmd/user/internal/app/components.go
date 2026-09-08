package app

import (
	"context"

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
