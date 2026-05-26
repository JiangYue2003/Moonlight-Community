# Phase 5.1 Fixes - 前后端协同适配清单

更新时间：2026-05-17  
适用范围：`zhiguang-go` 对接 `zhiguang_fe-main` 当前实现

## 1. 背景与目标

基于以下输入完成核对：
1. `zhiguang-go/docs/project-analysis.md`
2. `zhiguang_fe-main/docs/前后端协同说明.md`
3. `zhiguang-go` 实际路由、types、logic 代码

目标：让 Go 后端与现有前端在不改前端代码（或最小改前端）的前提下联调通过。

## 2. 必须修复（P0）

## 2.1 relation 服务

1. 修复 `POST /api/v1/relation/follow` 未实现问题  
现状：`FollowLogic` 仍为 TODO。  
要求：实现与 `UnfollowLogic` 对称的完整逻辑（取登录用户、调 relation-rpc、返回成功响应）。

2. follow/unfollow 入参改为兼容 query 参数  
现状：`FollowReq/UnfollowReq` 使用 `json:"toUserId"`，前端发送 `?toUserId=...`。  
要求：兼容 form/query 绑定（建议 `form:"toUserId"` 并保留 json 兼容）。

3. `following/followers` 响应改为前端期望数组结构  
现状：返回 `{items,nextCursor,hasMore}`。  
前端期望：`ProfileResponse[]`。  
要求：返回纯数组，或新增兼容接口并让前端当前路径得到数组。

4. `counter` 响应字段补齐  
现状：`CounterResp` 仅 `followings/followers/posts/likesReceived`。  
前端期望：`followings/followers/posts/likedPosts/favedPosts`。  
要求：补齐字段并统一语义（“我点赞/收藏过”或“我获赞/获藏”，二选一后前后端统一）。

5. `status` 鉴权策略与路由说明统一  
现状：路由标注可选鉴权，但 logic 实际强依赖登录用户。  
要求：二选一并统一：  
1) 改成强制鉴权；或  
2) 匿名时返回默认 false 三态，不报 401。

## 2.2 auth 服务

1. 登录请求去掉对 `channel` 的前端强依赖  
现状：后端要求 `channel=PASSWORD|CODE`，前端未传。  
要求：在 API 层自动推断：有 `code` 即 `CODE`，有 `password` 即 `PASSWORD`。

2. `/auth/send-code` 请求结构兼容前端  
现状：后端只用 `scene+identifier`；前端发送 `scene+identifierType+identifier`。  
要求：接受并忽略 `identifierType`（兼容字段），不要因多字段报错。

3. `/auth/register` 兼容前端无 nickname 场景  
现状：后端 `RegisterReq` 有 `nickname`。  
要求：支持 nickname 缺省并给默认值（例如“知光用户xxxx”）。

4. token 响应字段命名与前端对齐  
现状：`accessExpiresAt/refreshExpiresAt`。  
前端期望：`accessTokenExpiresAt/refreshTokenExpiresAt`。  
要求：返回前端字段名（或双字段兼容）。

5. `/auth/me` 返回完整用户信息  
现状：仅 `id + nickname`。  
前端需要：`avatar/phone/email/zgId(or zhId)/birthday/school/bio/gender/tagJson`。  
要求：按前端消费字段补齐。

6. `tagJson/tagsJson`、`zgId/zhId` 做兼容输出  
要求：短期双写双读，避免前端页面预填失败。

## 2.3 knowpost 服务

1. feed/detail 响应结构对齐前端
现状：返回 `creatorId/contentUrl/imgUrls...`，缺前端卡片/详情关键字段。  
前端需要至少包含：  
1) feed: `id,title,description,coverImage,tags,authorAvatar,authorNickname,tagJson,likeCount,favoriteCount,liked,faved,isTop`  
2) detail: `id,title,description,contentUrl,images,tags,authorAvatar,authorNickname,authorId,authorTagJson,likeCount,favoriteCount,liked,faved,isTop,visible,type,publishTime`

2. PATCH 元数据兼容 `tags/imgUrls` 直传
现状：依赖 `tagsSet/imgUrlsSet`，前端未传。  
要求：当收到 `tags`/`imgUrls` 字段时自动判定为 set=true。

3. 发布/编辑接口回包策略统一  
现状：部分接口返回 `KnowPostDetail`，前端按 `void` 使用。  
要求：保持前端不出错（返回体可有可无，但状态码与字段稳定）。

## 2.4 llm 与路由兼容层

1. 增加前端兼容路径（强烈建议）
现状：Go 路径为 `/api/v1/llm/describe`、`/api/v1/llm/qa/stream`。  
前端使用：`/api/v1/knowposts/description/suggest`、`/api/v1/knowposts/{id}/qa/stream`。  
要求：新增兼容路由转发到 llm logic，避免改前端。

2. SSE 鉴权策略与 EventSource 能力对齐
现状：`/api/v1/llm/qa/stream` 强制 Bearer；前端 EventSource 不带 Authorization。  
要求：二选一：  
1) 保持公开（按帖子可见性控制）；或  
2) 改为基于 Cookie 会话鉴权并确保 EventSource 可携带。  
当前前端更匹配方案 1。

3. 问答请求参数兼容
前端：path 中含 `postId`（`/knowposts/{id}/qa/stream`）+ query `question/topK/maxTokens`。  
要求：兼容层将 path id 映射到 llm 内部 `postId`。

## 2.5 search 服务

1. 搜索项字段名适配前端  
现状：`contentId/contentType/snippet...`。  
前端期望：`id,title,description,coverImage,tags,authorAvatar,authorNickname,tagJson,likeCount,favoriteCount,liked,faved,isTop`。  
要求：输出前端契约字段；不需要字段可额外保留但不能缺核心字段。

2. `suggest` 保持 `{items:string[]}`（已基本匹配）  
要求：确认空值时返回空数组而非 null。

## 2.6 profile 服务

1. PATCH 字段兼容 `tagJson`  
现状：后端字段为 `tagsJson`。  
要求：接受 `tagJson` 与 `tagsJson` 两种写法并归一。

2. PATCH 字段补齐 phone/email（若业务允许编辑）  
前端会提交 `phone`，当前后端 PatchReq 未定义。  
要求：若不允许编辑需明确忽略并返回稳定结果；允许则补实现。

3. 上传头像返回字段补齐  
现状：`{url, objectKey}`。  
前端期望：至少有 `avatar`（或完整 `ProfileResponse`）。  
要求：增加 `avatar` 字段（可与 `url` 同值）以兼容前端刷新逻辑。

## 3. 建议修复（P1）

1. 错误响应统一保证含 `message` 字段（前端直接展示）。  
2. 对前端已使用字段提供 1 个版本周期的双字段兼容（命名迁移缓冲）。  
3. 文档与代码同步：修复后同时更新 `project-analysis.md` 的“已知差异”章节。  
4. 增加 contract test：以 `zhiguang_fe-main` 请求样本做 API 兼容回归测试。

## 4. 执行顺序建议

1. P0-路由与鉴权兼容：`relation follow`、`auth login/send-code/me`、`llm 兼容路由 + SSE`。  
2. P0-响应结构适配：`knowpost feed/detail`、`relation list/counter`、`search items`。  
3. P0-参数写入兼容：`knowpost patch`、`profile tagJson/avatar`。  
4. P1-命名清理与文档收敛。

## 5. 验收清单（联调通过标准）

1. 登录/注册/刷新/me 全链路通过，前端不需要额外传 `channel`。  
2. 关注/取关成功，关系状态正常，关注/粉丝列表在前端弹窗正确显示。  
3. 首页 feed、搜索结果、详情页字段完整展示（作者、封面、计数、点赞收藏态）。  
4. 创建发布链路可用（草稿、上传、确认、更新、发布）。  
5. AI 摘要与 SSE 问答在当前前端页面直接可用。  
6. 资料页编辑与头像上传后，前端刷新用户信息正确。

