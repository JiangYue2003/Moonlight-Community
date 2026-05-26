// Package interceptor 提供 gRPC 服务端/客户端拦截器，
// 用于在 metadata 与 ctx 之间传递 userId。
package interceptor

import (
	"context"

	"github.com/zhiguang/zhiguang-go/common/ctxdata"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// UnaryUserIdInterceptor 服务端拦截器：从 metadata x-user-id 中提取 userId 并注入 ctx。
func UnaryUserIdInterceptor(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		if vals := md.Get(ctxdata.MetadataKeyUserId); len(vals) > 0 {
			if uid, err := ctxdata.ParseUserId(vals[0]); err == nil {
				ctx = ctxdata.WithUserId(ctx, uid)
			}
		}
	}
	return handler(ctx, req)
}

// UnaryUserIdClientInterceptor 客户端拦截器：把 ctx 中 userId 注入 metadata，
// 让被调端可以原样取到。HTTP→RPC 调用时挂到 zrpc.Client 上即可链路透传。
func UnaryUserIdClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	if uid, ok := ctxdata.GetUserId(ctx); ok {
		ctx = metadata.AppendToOutgoingContext(ctx, ctxdata.MetadataKeyUserId, ctxdata.FormatUserId(uid))
	}
	return invoker(ctx, method, req, reply, cc, opts...)
}
