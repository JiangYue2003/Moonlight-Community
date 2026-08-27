import { afterEach, describe, expect, it, vi } from "vitest";
import { authService } from "./authService";
import { knowpostService, uploadToPresigned } from "./knowpostService";

const jsonResponse = (value: unknown) =>
  new Response(JSON.stringify(value), {
    status: 200,
    headers: { "Content-Type": "application/json" }
  });

const authResponse = {
  user: {
    id: 42,
    nickname: "知光",
    avatar: "",
    phone: "13800000000",
    zgId: "zg-42",
    birthday: "",
    school: "",
    bio: "",
    gender: "",
    tagsJson: "[]"
  },
  token: {
    accessToken: "access-token",
    accessTokenExpiresAt: 1_785_391_112_000,
    refreshToken: "refresh-token",
    refreshTokenExpiresAt: 1_785_477_512_000
  }
};

describe("P0 backend contracts", () => {
  afterEach(() => {
    vi.unstubAllGlobals();
  });

  it("sends the Gateway CODE login shape and preserves nested auth response", async () => {
    const fetchMock = vi.fn().mockResolvedValue(jsonResponse(authResponse));
    vi.stubGlobal("fetch", fetchMock);

    const result = await authService.login({
      identifier: "13800000000",
      code: "123456",
      channel: "CODE"
    });

    const [url, init] = fetchMock.mock.calls[0] as [string, RequestInit];
    expect(url).toBe("/api/v1/auth/login");
    expect(JSON.parse(String(init.body))).toEqual({
      identifier: "13800000000",
      code: "123456",
      channel: "CODE"
    });
    expect(result).toEqual(authResponse);
  });

  it("sends only scene and identifier for verification codes", async () => {
    const fetchMock = vi.fn().mockResolvedValue(jsonResponse({
      cooldownSeconds: 60,
      expireSeconds: 300
    }));
    vi.stubGlobal("fetch", fetchMock);

    const result = await authService.sendCode({
      scene: "LOGIN",
      identifier: "13800000000"
    });

    const [, init] = fetchMock.mock.calls[0] as [string, RequestInit];
    expect(JSON.parse(String(init.body))).toEqual({
      scene: "LOGIN",
      identifier: "13800000000"
    });
    expect(result.cooldownSeconds).toBe(60);
  });

  it("refreshes without Authorization and reads token from the nested response", async () => {
    const fetchMock = vi.fn().mockResolvedValue(jsonResponse(authResponse));
    vi.stubGlobal("fetch", fetchMock);

    const result = await authService.refresh("refresh-token");

    const [, init] = fetchMock.mock.calls[0] as [string, RequestInit];
    expect(JSON.parse(String(init.body))).toEqual({ refreshToken: "refresh-token" });
    expect(new Headers(init.headers).has("Authorization")).toBe(false);
    expect(result.token.accessToken).toBe("access-token");
  });

  it("keeps upload and public-content URLs separate", async () => {
    const presignResponse = {
      objectKey: "knowpost/42/content.md",
      url: "https://upload.test/signed?signature=secret",
      headers: { "x-upload": "required" },
      expiresIn: 600,
      contentUrl: "https://cdn.test/knowpost/42/content.md"
    };
    const fetchMock = vi.fn()
      .mockResolvedValueOnce(jsonResponse(presignResponse))
      .mockResolvedValueOnce(new Response(null, {
        status: 200,
        headers: { ETag: "\"etag-1\"" }
      }));
    vi.stubGlobal("fetch", fetchMock);

    const result = await knowpostService.presign({
      scene: "knowpost_content",
      postId: "42",
      contentType: "text/markdown",
      ext: ".md"
    });
    const file = new File(["hello"], "content.md", { type: "text/markdown" });
    await uploadToPresigned(result.url, result.headers, file);

    expect(result.contentUrl).toBe("https://cdn.test/knowpost/42/content.md");
    expect(fetchMock.mock.calls[1][0]).toBe("https://upload.test/signed?signature=secret");
  });

  it("rejects an auth response missing its required token node", async () => {
    const fetchMock = vi.fn().mockResolvedValue(jsonResponse({ user: authResponse.user }));
    vi.stubGlobal("fetch", fetchMock);

    await expect(authService.login({
      identifier: "13800000000",
      code: "123456",
      channel: "CODE"
    })).rejects.toThrow("token");
  });

  it("rejects malformed current-user and presign responses", async () => {
    const fetchMock = vi.fn()
      .mockResolvedValueOnce(jsonResponse({ nickname: "知光" }))
      .mockResolvedValueOnce(jsonResponse({
        objectKey: "knowpost/42/content.md",
        headers: {},
        expiresIn: 600,
        contentUrl: "https://cdn.test/content.md"
      }));
    vi.stubGlobal("fetch", fetchMock);

    await expect(authService.fetchCurrentUser("access-token")).rejects.toThrow("user");
    await expect(knowpostService.presign({
      scene: "knowpost_content",
      postId: "42",
      contentType: "text/markdown",
      ext: ".md"
    })).rejects.toThrow("url");
  });

  it("rejects an otherwise shaped user that omits required avatar and phone", async () => {
    const fetchMock = vi.fn().mockResolvedValue(jsonResponse({
      id: 42,
      nickname: "知光"
    }));
    vi.stubGlobal("fetch", fetchMock);

    await expect(authService.fetchCurrentUser("access-token")).rejects.toThrow("user");
  });
});
