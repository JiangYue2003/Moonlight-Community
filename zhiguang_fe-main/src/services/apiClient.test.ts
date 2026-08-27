import { afterEach, describe, expect, it, vi } from "vitest";
import { apiFetch } from "./apiClient";

describe("apiFetch", () => {
  afterEach(() => {
    vi.unstubAllGlobals();
  });

  it("returns undefined for a 204 response", async () => {
    const fetchMock = vi.fn().mockResolvedValue(new Response(null, { status: 204 }));
    vi.stubGlobal("fetch", fetchMock);

    await expect(apiFetch<void>("/api/v1/relation/follow")).resolves.toBeUndefined();
  });
});
