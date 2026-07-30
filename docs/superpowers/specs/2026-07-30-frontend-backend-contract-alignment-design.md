# 前后端契约对齐设计

## 目标

以后端 Gateway 当前实际路由和 JSON 契约为唯一事实源，修改
`zhiguang_fe-main` 的请求、响应类型和页面数据使用方式，消除已经确认的
前后端协议漂移。

本轮不为旧前端增加 Gateway 兼容字段，不维护两套协议，也不修改核心业务
后端的既有契约。

## 范围

### 纳入

按优先级处理以下前端问题：

1. P0：验证码登录缺少 `channel`。
2. P0：刷新令牌响应被错误地当作扁平 Token。
3. P0：预签名上传读取了不存在的 `putUrl`。
4. P1：知文元数据更新未发送 `tagsSet`、`imgUrlsSet`。
5. P1：关注和取消关注未发送 JSON body。
6. P1：关注/粉丝列表未读取分页响应对象。
7. P1：搜索结果把 `contentId` 当作 `id`。
8. P1：Feed 和详情页仍使用后端未返回的旧字段。
9. P2：资料编辑允许提交后端不支持更新的手机号。
10. P2：个人统计使用了没有真实数据来源的 `likedPosts`、`favedPosts`。
11. 与认证契约直接相关的注册和验证码请求字段同步修正。

### 不纳入

- LLM 和 Agent 的功能测试与代码修改。
- `cmd/loadtest` 的编译和测试问题。
- 跨域部署方案；当前开发环境继续通过 Vite `/api` 代理保持同源。
- Gateway 兼容别名、旧协议迁移层或共享代码生成方案。
- 后端 merged-service 配置测试、启动脚本、日志轮转。
- 前端目录是否纳入 Git 的仓库治理问题。

## 事实源

契约判定顺序如下：

1. `services/gateway/internal/handler/router.go` 的实际路由和请求/响应。
2. Gateway 调用的 RPC 类型。
3. 本地运行服务的真实 HTTP 响应。

旧 `.api` 文档、前端已有 TypeScript 类型和页面假设都不能覆盖上述事实源。

## 契约调整

### 认证

验证码登录请求改为：

```json
{
  "identifier": "手机号",
  "code": "验证码",
  "channel": "CODE"
}
```

密码登录类型保留后端支持的：

```json
{
  "identifier": "手机号、邮箱或知光号",
  "password": "密码",
  "channel": "PASSWORD"
}
```

前端不再发送 Gateway 未读取的 `identifierType`。

验证码请求只发送 `scene`、`identifier`。注册请求发送
`identifier`、`code`、`password`、`nickname`、`agreeTerms`，注册页增加昵称输入。

`POST /api/v1/auth/token/refresh` 按实际后端响应解析为完整
`AuthResponse`，从 `response.token` 更新 Token，同时使用 `response.user`
刷新当前用户。

### 上传与发布

预签名响应定义为：

```ts
type PresignResponse = {
  url: string;
  objectKey: string;
  expiresIn: number;
  headers: Record<string, string>;
  contentUrl: string;
};
```

上传 PUT 使用 `url`，业务数据中记录可访问地址时使用 `contentUrl`，不再从签名
URL 截断查询参数推导公开地址。

更新知文元数据时，只要请求包含 `tags` 或 `imgUrls`，同时发送对应的
`tagsSet: true` 或 `imgUrlsSet: true`。空数组也必须能够明确表示清空。

### 关系

关注和取消关注改为 JSON body：

```json
{ "toUserId": 123 }
```

响应按 `204 No Content` 处理，前端服务返回 `Promise<void>`。

关注和粉丝列表按分页对象解析：

```ts
type RelationListResponse = {
  items: RelationProfile[];
  nextCursor: number;
  hasMore: boolean;
};
```

弹窗直接使用 `items` 和 `hasMore`，加载下一页优先使用 `nextCursor`，不再用
`Array.isArray` 静默丢弃响应。

### 搜索

搜索结果定义独立的 `SearchHit`，字段与 Gateway 一致：

- `contentId`
- `contentType`
- `title`
- `description`
- `snippet`
- `tags`
- `authorId`
- `authorNickname`
- `authorAvatar`
- `likeCount`
- `favoriteCount`
- `viewCount`
- `imgUrls`
- `isTop`

列表 key、详情链接和点赞收藏实体 ID 均使用 `contentId`，封面使用
`imgUrls[0]`。

### Feed 与详情

前端类型直接使用 Gateway 当前返回字段，不再声明不存在的作者和计数字段。

Feed 使用：

- `id`
- `creatorId`
- `title`
- `description`
- `contentUrl`
- `tags`
- `imgUrls`
- `visible`
- `isTop`
- `publishTime`

详情在上述基础上使用 `contentObjectKey`、`tagId`、`status`、`type`、
`createTime`、`updateTime`。

封面和详情图片来自 `imgUrls`。由于 Gateway 没有返回作者昵称和头像，界面使用
稳定且真实的 `用户 {creatorId}` 作为作者显示，不再显示问号或伪造资料。

计数继续通过公开的 `/api/v1/counter/:entityType/:entityId` 获取。
计数请求不再要求必须存在 Access Token；点赞和收藏写操作仍要求登录。

### 资料与个人统计

`PATCH /api/v1/profile/` 不支持手机号和邮箱更新，因此：

- `ProfileUpdateRequest` 删除 `phone`、`email`。
- 编辑资料页不再提供可编辑手机号输入，已有手机号只作为只读账户信息展示。
- 页面不再对未提交成功的手机号显示“资料已保存”。

个人统计使用后端真实字段：

- `followings`
- `followers`
- `posts`
- `likesReceived`

前端“获赞”绑定 `likesReceived`。后端没有“获藏”数据源，因此移除“获藏”
展示，不再显示硬编码的零。

## 代码边界

- HTTP 契约集中在 `src/types` 和 `src/services`。
- 页面只消费已经按后端命名的类型，不在组件中维护字段兼容逻辑。
- 不创建通用映射框架；仅在确有展示模型需要时使用小型纯函数。
- 不修改 Gateway、RPC、数据库模型或生成代码。

## 错误处理

- 缺失必要响应字段时抛出明确错误，不再用空数组或 `undefined` 静默降级。
- 公开 Feed、搜索、详情和计数允许未登录访问。
- 401 仍交由现有 `apiFetch` 和认证上下文处理。
- 后端 204 响应不尝试解析 JSON。

## 测试策略

### 测试先行

前端目前没有测试运行器，因此增加与 Vite 配套的 Vitest 开发依赖和
`npm test` 脚本。生产代码修改前先增加失败测试，至少覆盖：

1. 验证码登录包含 `channel: "CODE"`。
2. 刷新响应从 `token` 节点解析。
3. 预签名上传使用 `url`，公开地址使用 `contentUrl`。
4. 元数据数组携带对应的 `*Set` 标志。
5. 关注请求发送 JSON body。
6. 关系列表保留 `items`、`nextCursor`、`hasMore`。
7. 搜索卡片 ID 使用 `contentId`。
8. Feed/详情使用 `imgUrls` 和 `creatorId`。

### 完成验证

- 新增 Vitest 用例全部通过。
- `npm run lint` 通过。
- `npm run build` 通过。
- 本地 Gateway 的 Feed、搜索、详情和计数接口返回 200。
- 浏览器首页不再显示问号作者。
- 搜索结果链接不再出现 `/post/undefined`，控制台不再出现重复 key 警告。
- 未登录时公开计数可以加载；写操作仍会要求登录。

## 实施顺序

1. 建立前端契约测试基础。
2. 修复认证和上传三个 P0 问题。
3. 修复元数据、关系、搜索、Feed 和详情 P1 问题。
4. 修复资料编辑和个人统计 P2 问题。
5. 执行静态检查、生产构建、HTTP 和浏览器回归。
