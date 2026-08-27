import type { SendCodeResponse } from "@/types/auth";

export const getVerificationCodeCooldown = (response: SendCodeResponse) =>
  Math.max(1, Math.ceil(response.cooldownSeconds));
