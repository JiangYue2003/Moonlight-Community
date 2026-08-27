export type VerificationScene = "REGISTER" | "LOGIN" | "RESET_PASSWORD";

export type SendCodeRequest = {
  scene: VerificationScene;
  identifier: string;
};

export type SendCodeResponse = {
  cooldownSeconds: number;
  expireSeconds: number;
};

export type RegisterRequest = {
  identifier: string;
  code: string;
  password: string;
  nickname: string;
  agreeTerms: boolean;
};

// 与后端保持一致：登录/注册返回 AuthResponse，包含用户信息与令牌
import type { ProfileGender } from "@/types/profile";

export type AuthUserResponse = {
  id: number;
  nickname: string;
  avatar: string;
  phone: string;
  email?: string;
  zgId: string;
  birthday: string;
  school: string;
  bio: string;
  gender: ProfileGender;
  tagsJson: string;
};

export type TokenResponse = {
  accessToken: string;
  accessTokenExpiresAt: number; // Unix milliseconds
  refreshToken: string;
  refreshTokenExpiresAt: number; // Unix milliseconds
};

export type AuthResponse = {
  user: AuthUserResponse;
  token: TokenResponse;
};

// 别名以兼容现有代码引用
export type RegisterResponse = AuthResponse;

export type LoginRequest =
  | {
      identifier: string;
      password: string;
      channel: "PASSWORD";
      code?: never;
    }
  | {
      identifier: string;
      password?: never;
      code: string;
      channel: "CODE";
    };

export type LoginResponse = AuthResponse;

export type RefreshResponse = AuthResponse;

export type LogoutRequest = {
  refreshToken: string;
};

// 与 /auth/me 返回一致
export type AuthenticatedUser = AuthUserResponse;

export type ErrorResponse = {
  code: string;
  message: string;
  path?: string;
  timestamp?: string;
};
