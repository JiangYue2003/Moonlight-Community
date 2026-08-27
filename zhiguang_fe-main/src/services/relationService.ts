import { apiFetch } from "./apiClient";
import type {
  RelationCountersResponse,
  RelationListResponse,
  RelationStatusResponse
} from "@/types/relation";

const RELATION_PREFIX = "/api/v1/relation";

const isRecord = (value: unknown): value is Record<string, unknown> =>
  typeof value === "object" && value !== null;

const parseRelationStatus = (value: unknown): RelationStatusResponse => {
  if (
    !isRecord(value) ||
    typeof value.following !== "boolean" ||
    typeof value.followedBy !== "boolean" ||
    typeof value.mutual !== "boolean"
  ) {
    throw new Error("后端关系状态响应缺少必要字段");
  }
  return value as RelationStatusResponse;
};

const parseRelationList = (value: unknown): RelationListResponse => {
  if (
    !isRecord(value) ||
    !Array.isArray(value.items) ||
    typeof value.nextCursor !== "number" ||
    !Number.isFinite(value.nextCursor) ||
    typeof value.hasMore !== "boolean"
  ) {
    throw new Error("后端关系列表响应缺少 items 或分页字段");
  }
  const validItems = value.items.every(item =>
    isRecord(item) &&
    typeof item.id === "number" &&
    Number.isFinite(item.id) &&
    typeof item.nickname === "string" &&
    typeof item.avatar === "string" &&
    typeof item.zgId === "string" &&
    typeof item.bio === "string"
  );
  if (!validItems) {
    throw new Error("后端关系列表响应包含无效的 items");
  }
  return value as RelationListResponse;
};

const parseRelationCounters = (value: unknown): RelationCountersResponse => {
  if (
    !isRecord(value) ||
    typeof value.followings !== "number" ||
    !Number.isFinite(value.followings) ||
    typeof value.followers !== "number" ||
    !Number.isFinite(value.followers) ||
    typeof value.posts !== "number" ||
    !Number.isFinite(value.posts) ||
    typeof value.likesReceived !== "number" ||
    !Number.isFinite(value.likesReceived)
  ) {
    throw new Error("后端个人统计响应缺少必要字段");
  }
  return value as RelationCountersResponse;
};

export const relationService = {
  follow: (toUserId: number, accessToken: string) =>
    apiFetch<void>(`${RELATION_PREFIX}/follow`, {
      method: "POST",
      body: { toUserId },
      accessToken
    }),

  unfollow: (toUserId: number, accessToken: string) =>
    apiFetch<void>(`${RELATION_PREFIX}/unfollow`, {
      method: "POST",
      body: { toUserId },
      accessToken
    }),

  status: (toUserId: number, accessToken?: string | null) =>
    apiFetch<unknown>(`${RELATION_PREFIX}/status?toUserId=${toUserId}`, {
      accessToken
    }).then(parseRelationStatus),

  following: (userId: number, accessToken: string, limit = 20, offset = 0, cursor?: number) => {
    const params = new URLSearchParams({ userId: String(userId), limit: String(limit), offset: String(offset) });
    if (typeof cursor === "number") params.set("cursor", String(cursor));
    return apiFetch<unknown>(`${RELATION_PREFIX}/following?${params.toString()}`, {
      accessToken
    }).then(parseRelationList);
  },

  followers: (userId: number, accessToken: string, limit = 20, offset = 0, cursor?: number) => {
    const params = new URLSearchParams({ userId: String(userId), limit: String(limit), offset: String(offset) });
    if (typeof cursor === "number") params.set("cursor", String(cursor));
    return apiFetch<unknown>(`${RELATION_PREFIX}/followers?${params.toString()}`, {
      accessToken
    }).then(parseRelationList);
  },

  counters: (userId: number, accessToken: string) =>
    apiFetch<unknown>(`${RELATION_PREFIX}/counter?userId=${userId}`, {
      accessToken
    }).then(parseRelationCounters)
};
