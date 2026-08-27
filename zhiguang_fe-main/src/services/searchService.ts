import { apiFetch } from "./apiClient";
import type { SearchResponse, SuggestResponse } from "@/types/search";

const SEARCH_PREFIX = "/api/v1/search";

const isRecord = (value: unknown): value is Record<string, unknown> =>
  typeof value === "object" && value !== null;

const isStringArray = (value: unknown): value is string[] =>
  Array.isArray(value) && value.every(item => typeof item === "string");

const isFiniteNumber = (value: unknown): value is number =>
  typeof value === "number" && Number.isFinite(value);

const parseSearchResponse = (value: unknown): SearchResponse => {
  if (
    !isRecord(value) ||
    !Array.isArray(value.items) ||
    typeof value.nextAfter !== "string" ||
    typeof value.hasMore !== "boolean"
  ) {
    throw new Error("后端搜索响应缺少 items 或分页字段");
  }
  const validItems = value.items.every(item =>
    isRecord(item) &&
    typeof item.contentId === "string" &&
    typeof item.contentType === "string" &&
    typeof item.title === "string" &&
    typeof item.description === "string" &&
    typeof item.snippet === "string" &&
    (item.tags === null || isStringArray(item.tags)) &&
    isFiniteNumber(item.authorId) &&
    typeof item.authorNickname === "string" &&
    typeof item.authorAvatar === "string" &&
    isFiniteNumber(item.likeCount) &&
    isFiniteNumber(item.favoriteCount) &&
    isFiniteNumber(item.viewCount) &&
    isStringArray(item.imgUrls) &&
    typeof item.isTop === "boolean"
  );
  if (!validItems) {
    throw new Error("后端搜索响应包含无效的 items");
  }
  return value as SearchResponse;
};

const parseSuggestResponse = (value: unknown): SuggestResponse => {
  if (!isRecord(value) || !Array.isArray(value.items) || !value.items.every(item => typeof item === "string")) {
    throw new Error("后端搜索联想响应缺少 items");
  }
  return value as SuggestResponse;
};

export const searchService = {
  query: (params: { q: string; size?: number; tags?: string; after?: string | null }) => {
    const { q, size = 20, tags, after } = params;
    const usp = new URLSearchParams();
    usp.set("q", q);
    if (size) usp.set("size", String(size));
    if (tags) usp.set("tags", tags);
    if (after) usp.set("after", after);
    return apiFetch<unknown>(`${SEARCH_PREFIX}/?${usp.toString()}`).then(parseSearchResponse);
  },

  suggest: (prefix: string, size = 10) => {
    const usp = new URLSearchParams();
    usp.set("prefix", prefix);
    if (size) usp.set("size", String(size));
    return apiFetch<unknown>(`${SEARCH_PREFIX}/suggest?${usp.toString()}`).then(parseSuggestResponse);
  }
};
