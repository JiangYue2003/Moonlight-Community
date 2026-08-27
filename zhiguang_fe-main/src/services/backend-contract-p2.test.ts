import { afterEach, describe, expect, it, vi } from "vitest";
import { profileService } from "./profileService";

const jsonResponse = (value: unknown) =>
  new Response(JSON.stringify(value), {
    status: 200,
    headers: { "Content-Type": "application/json" }
  });

describe("P2 backend request contracts", () => {
  afterEach(() => {
    vi.unstubAllGlobals();
  });

  it("patches the exact Gateway profile route without a redirect", async () => {
    const profileResponse = {
      id: 42,
      nickname: "知光",
      avatar: "",
      bio: "",
      zgId: "zg-42",
      gender: "UNKNOWN",
      birthday: "",
      school: "",
      email: "",
      phone: "13800000000",
      tagsJson: "[\"Go\"]"
    };
    const fetchMock = vi.fn().mockResolvedValue(jsonResponse(profileResponse));
    vi.stubGlobal("fetch", fetchMock);

    const result = await profileService.update({ nickname: "知光", tagsJson: "[\"Go\"]" });

    const [url, init] = fetchMock.mock.calls[0] as [string, RequestInit];
    expect(url).toBe("/api/v1/profile/");
    expect(JSON.parse(String(init.body))).toEqual({
      nickname: "知光",
      tagsJson: "[\"Go\"]"
    });
    expect(result).toEqual(profileResponse);
  });

  it("reads the Gateway avatar-upload response shape", async () => {
    const response = {
      url: "https://cdn.test/avatar.png",
      avatar: "https://cdn.test/avatar.png",
      objectKey: "avatar/42.png"
    };
    const fetchMock = vi.fn().mockResolvedValue(jsonResponse(response));
    vi.stubGlobal("fetch", fetchMock);

    const result = await profileService.uploadAvatar(
      new File(["avatar"], "avatar.png", { type: "image/png" })
    );

    expect(result).toEqual(response);
  });

  it("rejects a malformed profile response", async () => {
    const fetchMock = vi.fn().mockResolvedValue(jsonResponse({ nickname: "知光" }));
    vi.stubGlobal("fetch", fetchMock);

    await expect(profileService.update({ nickname: "知光" })).rejects.toThrow("profile");
  });
});
