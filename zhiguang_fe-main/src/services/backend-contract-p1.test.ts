import { afterEach, describe, expect, it, vi } from "vitest";
import { knowpostService } from "./knowpostService";
import { relationService } from "./relationService";
import { searchService } from "./searchService";

const noContentResponse = () => new Response(null, { status: 204 });
const jsonResponse = (value: unknown) =>
  new Response(JSON.stringify(value), {
    status: 200,
    headers: { "Content-Type": "application/json" }
  });

const detailResponse = {
  id: "post-1",
  creatorId: 42,
  title: "标题",
  description: "摘要",
  tagId: 0,
  tags: null,
  contentUrl: "https://cdn.test/post-1.md",
  contentObjectKey: "knowpost/post-1/content.md",
  imgUrls: null,
  status: "published",
  type: "image_text",
  isTop: false,
  visible: "public",
  createTime: 1_785_391_100_000,
  updateTime: 1_785_391_110_000,
  publishTime: 1_785_391_112_000
};

describe("P1 backend request contracts", () => {
  afterEach(() => {
    vi.unstubAllGlobals();
  });

  it("marks even empty tag and image arrays as present and reads the 200 detail", async () => {
    const fetchMock = vi.fn().mockResolvedValue(jsonResponse(detailResponse));
    vi.stubGlobal("fetch", fetchMock);

    const result = await knowpostService.update("post-1", {
      tags: [],
      imgUrls: []
    });

    const [, init] = fetchMock.mock.calls[0] as [string, RequestInit];
    expect(JSON.parse(String(init.body))).toEqual({
      tags: [],
      tagsSet: true,
      imgUrls: [],
      imgUrlsSet: true
    });
    expect(result).toEqual(detailResponse);
  });

  it("sends follow target in a JSON body", async () => {
    const fetchMock = vi.fn().mockResolvedValue(noContentResponse());
    vi.stubGlobal("fetch", fetchMock);

    await relationService.follow(42, "access-token");

    const [url, init] = fetchMock.mock.calls[0] as [string, RequestInit];
    expect(url).toBe("/api/v1/relation/follow");
    expect(JSON.parse(String(init.body))).toEqual({ toUserId: 42 });
  });

  it("sends unfollow target in a JSON body", async () => {
    const fetchMock = vi.fn().mockResolvedValue(noContentResponse());
    vi.stubGlobal("fetch", fetchMock);

    await relationService.unfollow(42, "access-token");

    const [url, init] = fetchMock.mock.calls[0] as [string, RequestInit];
    expect(url).toBe("/api/v1/relation/unfollow");
    expect(JSON.parse(String(init.body))).toEqual({ toUserId: 42 });
  });

  it("calls the exact Gateway search route without a redirect", async () => {
    const response = {
      items: [{
        contentId: "post-1",
        contentType: "knowpost",
        title: "标题",
        description: "摘要",
        snippet: "摘要",
        tags: null,
        authorId: 42,
        authorNickname: "",
        authorAvatar: "",
        likeCount: 0,
        favoriteCount: 0,
        viewCount: 0,
        imgUrls: [""],
        isTop: false
      }],
      nextAfter: "",
      hasMore: false
    };
    const fetchMock = vi.fn().mockResolvedValue(jsonResponse(response));
    vi.stubGlobal("fetch", fetchMock);

    const result = await searchService.query({ q: "Load" });

    const [url] = fetchMock.mock.calls[0] as [string, RequestInit];
    expect(url).toBe("/api/v1/search/?q=Load&size=20");
    expect(result.items[0].contentId).toBe("post-1");
    expect(result.items[0].tags).toBeNull();
  });

  it("preserves the relation page envelope and sends required auth", async () => {
    const response = {
      items: [{ id: 7, nickname: "同学", avatar: "", zgId: "zg-7", bio: "" }],
      nextCursor: 123,
      hasMore: true
    };
    const fetchMock = vi.fn().mockResolvedValue(jsonResponse(response));
    vi.stubGlobal("fetch", fetchMock);

    const result = await relationService.following(42, "access-token");

    const [, init] = fetchMock.mock.calls[0] as [string, RequestInit];
    expect(new Headers(init.headers).get("Authorization")).toBe("Bearer access-token");
    expect(result).toEqual(response);
  });

  it("preserves Feed and detail Gateway fields", async () => {
    const feedResponse = {
      items: [{
        id: "post-1",
        creatorId: 42,
        title: "标题",
        description: "摘要",
        contentUrl: "https://cdn.test/post-1.md",
        tags: null,
        imgUrls: ["https://cdn.test/cover.png"],
        visible: "public",
        isTop: false,
        publishTime: 1_785_391_112_000
      }],
      page: 1,
      size: 20,
      hasMore: false
    };
    const fetchMock = vi.fn()
      .mockResolvedValueOnce(jsonResponse(feedResponse))
      .mockResolvedValueOnce(jsonResponse({
        ...detailResponse,
        imgUrls: ["https://cdn.test/cover.png"]
      }));
    vi.stubGlobal("fetch", fetchMock);

    const feed = await knowpostService.feed();
    const detail = await knowpostService.detail("post-1");

    expect(feed.items[0]).toMatchObject({
      id: "post-1",
      creatorId: 42,
      imgUrls: ["https://cdn.test/cover.png"]
    });
    expect(detail).toMatchObject({
      id: "post-1",
      creatorId: 42,
      contentObjectKey: "knowpost/post-1/content.md"
    });
  });

  it("loads public counters without Authorization", async () => {
    const fetchMock = vi.fn().mockResolvedValue(jsonResponse({
      entityType: "knowpost",
      entityId: "post-1",
      counts: { like: 3, fav: 2 }
    }));
    vi.stubGlobal("fetch", fetchMock);

    const result = await knowpostService.counters("post-1", null);

    const [, init] = fetchMock.mock.calls[0] as [string, RequestInit];
    expect(new Headers(init.headers).has("Authorization")).toBe(false);
    expect(result.counts).toEqual({ like: 3, fav: 2 });
  });

  it.each([
    ["Feed", () => knowpostService.feed()],
    ["search", () => searchService.query({ q: "Load" })],
    ["relation page", () => relationService.following(42, "access-token")]
  ])("rejects a malformed %s response instead of showing an empty list", async (_name, request) => {
    const fetchMock = vi.fn().mockResolvedValue(jsonResponse({ hasMore: false }));
    vi.stubGlobal("fetch", fetchMock);

    await expect(request()).rejects.toThrow("items");
  });

  it("rejects a Feed item missing required display fields", async () => {
    const fetchMock = vi.fn().mockResolvedValue(jsonResponse({
      items: [{
        id: "post-1",
        creatorId: 42,
        tags: null,
        imgUrls: null
      }],
      page: 1,
      size: 20,
      hasMore: false
    }));
    vi.stubGlobal("fetch", fetchMock);

    await expect(knowpostService.feed()).rejects.toThrow("items");
  });

  it("rejects a search item missing required author and count fields", async () => {
    const fetchMock = vi.fn().mockResolvedValue(jsonResponse({
      items: [{
        contentId: "post-1",
        tags: null,
        imgUrls: []
      }],
      nextAfter: "",
      hasMore: false
    }));
    vi.stubGlobal("fetch", fetchMock);

    await expect(searchService.query({ q: "Load" })).rejects.toThrow("items");
  });
});
