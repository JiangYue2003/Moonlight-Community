import { apiFetch } from "./apiClient";
import type {
  CreateDraftResponse,
  PresignRequest,
  PresignResponse,
  ConfirmContentRequest,
  UpdateKnowPostRequest,
  FeedResponse,
  KnowpostDetailResponse,
  LikeActionResponse,
  FavActionResponse,
  CounterResponse,
  VisibleScope
} from "@/types/knowpost";

const KNOWPOST_PREFIX = "/api/v1/knowposts";
const STORAGE_PREFIX = "/api/v1/storage";

const isRecord = (value: unknown): value is Record<string, unknown> =>
  typeof value === "object" && value !== null;

const isFiniteNumber = (value: unknown): value is number =>
  typeof value === "number" && Number.isFinite(value);

const isStringArrayOrNull = (value: unknown) =>
  value === null || (Array.isArray(value) && value.every(item => typeof item === "string"));

const isVisibleScope = (value: unknown): value is VisibleScope =>
  value === "public" ||
  value === "followers" ||
  value === "school" ||
  value === "private" ||
  value === "unlisted";

const parseDraftResponse = (value: unknown): CreateDraftResponse => {
  if (!isRecord(value) || typeof value.id !== "string" || !value.id) {
    throw new Error("后端草稿响应缺少 id");
  }
  return value as CreateDraftResponse;
};

const parsePresignResponse = (value: unknown): PresignResponse => {
  if (
    !isRecord(value) ||
    typeof value.url !== "string" ||
    !value.url ||
    typeof value.contentUrl !== "string" ||
    typeof value.objectKey !== "string" ||
    !isRecord(value.headers) ||
    !Object.values(value.headers).every(header => typeof header === "string") ||
    !isFiniteNumber(value.expiresIn) ||
    value.expiresIn <= 0
  ) {
    throw new Error("后端预签名响应缺少 url、contentUrl 或上传元数据");
  }
  return value as PresignResponse;
};

const parseFeedResponse = (value: unknown): FeedResponse => {
  if (
    !isRecord(value) ||
    !Array.isArray(value.items) ||
    !isFiniteNumber(value.page) ||
    !isFiniteNumber(value.size) ||
    typeof value.hasMore !== "boolean"
  ) {
    throw new Error("后端 Feed 响应缺少 items 或分页字段");
  }
  const validItems = value.items.every(item =>
    isRecord(item) &&
    typeof item.id === "string" &&
    isFiniteNumber(item.creatorId) &&
    typeof item.title === "string" &&
    typeof item.description === "string" &&
    typeof item.contentUrl === "string" &&
    isStringArrayOrNull(item.tags) &&
    isStringArrayOrNull(item.imgUrls) &&
    isVisibleScope(item.visible) &&
    typeof item.isTop === "boolean" &&
    isFiniteNumber(item.publishTime)
  );
  if (!validItems) {
    throw new Error("后端 Feed 响应包含无效的 items");
  }
  return value as FeedResponse;
};

const parseDetailResponse = (value: unknown): KnowpostDetailResponse => {
  if (
    !isRecord(value) ||
    typeof value.id !== "string" ||
    isFiniteNumber(value.creatorId) === false ||
    typeof value.title !== "string" ||
    typeof value.description !== "string" ||
    isFiniteNumber(value.tagId) === false ||
    typeof value.contentUrl !== "string" ||
    typeof value.contentObjectKey !== "string" ||
    !isStringArrayOrNull(value.tags) ||
    !isStringArrayOrNull(value.imgUrls) ||
    typeof value.status !== "string" ||
    typeof value.type !== "string" ||
    typeof value.isTop !== "boolean" ||
    !isVisibleScope(value.visible) ||
    isFiniteNumber(value.createTime) === false ||
    isFiniteNumber(value.updateTime) === false ||
    isFiniteNumber(value.publishTime) === false
  ) {
    throw new Error("后端知文详情响应缺少必要字段");
  }
  return value as KnowpostDetailResponse;
};

const parseCounterResponse = (value: unknown): CounterResponse => {
  if (
    !isRecord(value) ||
    typeof value.entityType !== "string" ||
    typeof value.entityId !== "string" ||
    !isRecord(value.counts) ||
    !isFiniteNumber(value.counts.like) ||
    !isFiniteNumber(value.counts.fav)
  ) {
    throw new Error("后端计数响应缺少 counts");
  }
  return value as CounterResponse;
};

const parseLikeActionResponse = (value: unknown): LikeActionResponse => {
  if (!isRecord(value) || typeof value.changed !== "boolean" || typeof value.liked !== "boolean") {
    throw new Error("后端点赞响应缺少必要字段");
  }
  return value as LikeActionResponse;
};

const parseFavActionResponse = (value: unknown): FavActionResponse => {
  if (!isRecord(value) || typeof value.changed !== "boolean" || typeof value.faved !== "boolean") {
    throw new Error("后端收藏响应缺少必要字段");
  }
  return value as FavActionResponse;
};

export const knowpostService = {
  createDraft: () =>
    apiFetch<unknown>(`${KNOWPOST_PREFIX}/drafts`, { method: "POST" }).then(parseDraftResponse),

  presign: (payload: PresignRequest) =>
    apiFetch<unknown>(`${STORAGE_PREFIX}/presign`, { method: "POST", body: payload }).then(parsePresignResponse),

  confirmContent: (id: string, payload: ConfirmContentRequest) =>
    apiFetch<void>(`${KNOWPOST_PREFIX}/${id}/content/confirm`, { method: "POST", body: payload }),

  update: (id: string, payload: UpdateKnowPostRequest) => {
    const body: UpdateKnowPostRequest = {
      ...payload,
      ...(payload.tags !== undefined ? { tagsSet: true } : {}),
      ...(payload.imgUrls !== undefined ? { imgUrlsSet: true } : {})
    };
    return apiFetch<unknown>(`${KNOWPOST_PREFIX}/${id}`, { method: "PATCH", body }).then(parseDetailResponse);
  },

  publish: (id: string) =>
    apiFetch<unknown>(`${KNOWPOST_PREFIX}/${id}/publish`, { method: "POST" }).then(parseDetailResponse)
  ,
  
  // 设置置顶（需鉴权）
  setTop: (id: string, isTop: boolean, accessToken: string) =>
    apiFetch<void>(`${KNOWPOST_PREFIX}/${id}/top`, {
      method: "PATCH",
      body: { isTop },
      accessToken
    })
  ,

  // 设置可见性（需鉴权）
  setVisibility: (id: string, visible: VisibleScope, accessToken: string) =>
    apiFetch<void>(`${KNOWPOST_PREFIX}/${id}/visibility`, {
      method: "PATCH",
      body: { visible },
      accessToken
    })
  ,

  // 删除知文（需鉴权）
  remove: (id: string, accessToken: string) =>
    apiFetch<void>(`${KNOWPOST_PREFIX}/${id}`, {
      method: "DELETE",
      accessToken
    })
  ,

  // 获取首页 Feed 列表（公开内容）
  feed: (page = 1, size = 20) =>
    apiFetch<unknown>(`${KNOWPOST_PREFIX}/feed?page=${page}&size=${size}`).then(parseFeedResponse)
  ,

  // 获取我的知文（需鉴权）
  mine: (page = 1, size = 20, accessToken: string) =>
    apiFetch<unknown>(`${KNOWPOST_PREFIX}/mine?page=${page}&size=${size}`, {
      accessToken
    }).then(parseFeedResponse)
  ,

  // 获取知文详情（公开内容无需鉴权；非公开需要作者凭证）
  detail: (id: string, accessToken?: string) =>
    apiFetch<unknown>(`${KNOWPOST_PREFIX}/detail/${id}`, {
      accessToken: accessToken ?? null
    }).then(parseDetailResponse)
  ,

  // 生成知文摘要（需鉴权）
  suggestDescription: (content: string, accessToken: string) =>
    apiFetch<{ description: string }>(`${KNOWPOST_PREFIX}/description/suggest`, {
      method: "POST",
      body: { content },
      accessToken
    })
  ,

  // 点赞/取消点赞（需鉴权）
  like: (entityId: string, accessToken: string, entityType: string = "knowpost") =>
    apiFetch<unknown>(`/api/v1/action/like`, {
      method: "POST",
      body: { entityType, entityId },
      accessToken
    }).then(parseLikeActionResponse)
  ,
  unlike: (entityId: string, accessToken: string, entityType: string = "knowpost") =>
    apiFetch<unknown>(`/api/v1/action/unlike`, {
      method: "POST",
      body: { entityType, entityId },
      accessToken
    }).then(parseLikeActionResponse)
  ,

  // 收藏/取消收藏（需鉴权）
  fav: (entityId: string, accessToken: string, entityType: string = "knowpost") =>
    apiFetch<unknown>(`/api/v1/action/fav`, {
      method: "POST",
      body: { entityType, entityId },
      accessToken
    }).then(parseFavActionResponse)
  ,
  unfav: (entityId: string, accessToken: string, entityType: string = "knowpost") =>
    apiFetch<unknown>(`/api/v1/action/unfav`, {
      method: "POST",
      body: { entityType, entityId },
      accessToken
    }).then(parseFavActionResponse)
  ,

  // 获取计数（需鉴权）
  counters: (entityId: string, accessToken?: string | null, entityType: string = "knowpost") =>
    apiFetch<unknown>(`/api/v1/counter/${entityType}/${entityId}?metrics=like,fav`, {
      accessToken
    }).then(parseCounterResponse)
};

/**
 * 直传到预签名 URL。注意：S3/OSS 会在响应头返回 ETag。
 */
export async function uploadToPresigned(url: string, headers: Record<string, string>, file: File) {
  const resp = await fetch(url, {
    method: "PUT",
    headers,
    body: file,
    // 跨域上传通常不需要携带凭据
    credentials: "omit"
  });
  if (!resp.ok) {
    const text = await resp.text().catch(() => "");
    throw new Error(text || `上传失败：${resp.status}`);
  }
  // ETag 常带双引号
  const etag = resp.headers.get("ETag") || resp.headers.get("etag") || "";
  return { etag };
}

export async function computeSha256(file: File) {
  const buf = await file.arrayBuffer();
  const digest = await crypto.subtle.digest("SHA-256", buf);
  const hex = Array.from(new Uint8Array(digest)).map(b => b.toString(16).padStart(2, "0")).join("");
  return hex;
}
