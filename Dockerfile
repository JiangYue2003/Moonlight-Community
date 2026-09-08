# 多阶段构建：编译阶段
FROM golang:1.23-alpine AS builder

# 设置工作目录
WORKDIR /build

# 安装必要的构建工具
RUN apk add --no-cache git

# 复制 go mod 文件
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 编译所有服务
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/gateway ./services/gateway
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/counter-merged ./services/counter/cmd/counter
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/knowpost-merged ./services/knowpost/cmd/knowpost
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/relation-merged ./services/relation/cmd/relation
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/search-merged ./services/search/cmd/search
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/llm-merged ./services/llm/cmd/llm
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/user-rpc ./services/user/cmd/user
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/storage ./services/storage/cmd/storage
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/agent ./services/agent/cmd/agent

# 运行阶段：最小化镜像
FROM alpine:latest

# 安装必要的运行时依赖
RUN apk --no-cache add ca-certificates tzdata

# 设置时区
ENV TZ=Asia/Shanghai

WORKDIR /app

# 从构建阶段复制二进制文件
COPY --from=builder /build/bin/ ./bin/

# 复制配置文件
COPY services/ ./services/
COPY certs/ ./certs/

# 默认命令（会被 docker-compose 覆盖）
CMD ["./bin/gateway"]
