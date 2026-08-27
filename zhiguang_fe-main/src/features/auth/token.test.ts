import { describe, expect, it } from "vitest";
import { getAccessTokenExpiry } from "./token";

describe("auth token response contract", () => {
  it("keeps the Gateway Unix-millisecond expiry unchanged", () => {
    expect(getAccessTokenExpiry(1_785_391_112_000)).toBe(1_785_391_112_000);
  });

  it("rejects an invalid expiry instead of inventing a local fallback", () => {
    expect(() => getAccessTokenExpiry(Number.NaN)).toThrow("accessTokenExpiresAt");
  });
});
