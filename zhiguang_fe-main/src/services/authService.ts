import { apiFetch } from "./apiClient";
import type {
  AuthResponse,
  AuthenticatedUser,
  LoginRequest,
  LoginResponse,
  LogoutRequest,
  RefreshResponse,
  RegisterRequest,
  RegisterResponse,
  SendCodeRequest,
  SendCodeResponse
} from "@/types/auth";

const AUTH_PREFIX = "/api/v1/auth";

const isRecord = (value: unknown): value is Record<string, unknown> =>
  typeof value === "object" && value !== null;

const isProfileGender = (value: unknown) =>
  value === "" ||
  value === "MALE" ||
  value === "FEMALE" ||
  value === "OTHER" ||
  value === "UNKNOWN";

const parseAuthenticatedUser = (value: unknown): AuthenticatedUser => {
  if (
    !isRecord(value) ||
    typeof value.id !== "number" ||
    !Number.isFinite(value.id) ||
    typeof value.nickname !== "string" ||
    typeof value.avatar !== "string" ||
    typeof value.phone !== "string" ||
    (value.email !== undefined && typeof value.email !== "string") ||
    typeof value.zgId !== "string" ||
    typeof value.birthday !== "string" ||
    typeof value.school !== "string" ||
    typeof value.bio !== "string" ||
    !isProfileGender(value.gender) ||
    typeof value.tagsJson !== "string"
  ) {
    throw new Error("后端当前 user 响应缺少必要字段");
  }
  return value as AuthenticatedUser;
};

const parseSendCodeResponse = (value: unknown): SendCodeResponse => {
  if (
    !isRecord(value) ||
    typeof value.cooldownSeconds !== "number" ||
    !Number.isFinite(value.cooldownSeconds) ||
    value.cooldownSeconds <= 0 ||
    typeof value.expireSeconds !== "number" ||
    !Number.isFinite(value.expireSeconds) ||
    value.expireSeconds <= 0
  ) {
    throw new Error("后端验证码响应缺少有效的 cooldownSeconds 或 expireSeconds");
  }
  return value as SendCodeResponse;
};

const parseAuthResponse = (value: unknown): AuthResponse => {
  if (!isRecord(value) || !isRecord(value.user) || !isRecord(value.token)) {
    throw new Error("后端认证响应缺少 user 或 token");
  }
  const { user, token } = value;
  parseAuthenticatedUser(user);
  if (
    typeof token.accessToken !== "string" ||
    !token.accessToken ||
    typeof token.refreshToken !== "string" ||
    !token.refreshToken ||
    typeof token.accessTokenExpiresAt !== "number" ||
    !Number.isFinite(token.accessTokenExpiresAt) ||
    token.accessTokenExpiresAt <= 0 ||
    typeof token.refreshTokenExpiresAt !== "number" ||
    !Number.isFinite(token.refreshTokenExpiresAt) ||
    token.refreshTokenExpiresAt <= 0
  ) {
    throw new Error("后端认证响应包含无效的 token");
  }
  return value as AuthResponse;
};

export const authService = {
  sendCode: (payload: SendCodeRequest) =>
    apiFetch<unknown>(`${AUTH_PREFIX}/send-code`, {
      method: "POST",
      body: payload
    }).then(parseSendCodeResponse),

  register: (payload: RegisterRequest) =>
    apiFetch<unknown>(`${AUTH_PREFIX}/register`, {
      method: "POST",
      body: payload
    }).then(parseAuthResponse) as Promise<RegisterResponse>,

  login: (payload: LoginRequest) =>
    apiFetch<unknown>(`${AUTH_PREFIX}/login`, {
      method: "POST",
      body: payload
    }).then(parseAuthResponse) as Promise<LoginResponse>,

  logout: (payload: LogoutRequest, accessToken: string) =>
    apiFetch<void>(`${AUTH_PREFIX}/logout`, {
      method: "POST",
      body: payload,
      accessToken
    }),

  fetchCurrentUser: (accessToken: string) =>
    apiFetch<unknown>(`${AUTH_PREFIX}/me`, {
      accessToken
    }).then(parseAuthenticatedUser),

  refresh: (refreshToken: string) =>
    apiFetch<unknown>(`${AUTH_PREFIX}/token/refresh`, {
      method: "POST",
      body: { refreshToken },
      // 刷新接口不应携带（已过期的）access token
      accessToken: null
    }).then(parseAuthResponse) as Promise<RefreshResponse>
};
