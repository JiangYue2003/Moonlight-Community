import { describe, expect, it } from "vitest";
import { getVerificationCodeCooldown } from "./verificationCode";

describe("verification-code response contract", () => {
  it("uses cooldownSeconds for the resend timer, not expireSeconds", () => {
    expect(getVerificationCodeCooldown({
      cooldownSeconds: 7,
      expireSeconds: 300
    })).toBe(7);
  });
});
