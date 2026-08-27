import type {
  AuthResponse,
  LoginRequest,
  RefreshResponse,
  RegisterRequest,
  SendCodeRequest,
  TokenResponse
} from "./auth";
import type { PresignResponse } from "./knowpost";

type Assert<T extends true> = T;
type Equal<A, B> =
  (<T>() => T extends A ? 1 : 2) extends
  (<T>() => T extends B ? 1 : 2) ? true : false;

type LoginRequiresChannel = Assert<
  LoginRequest extends { channel: "CODE" | "PASSWORD" } ? true : false
>;

type SendCodeOmitsIdentifierType = Assert<
  "identifierType" extends keyof SendCodeRequest ? false : true
>;

type RegisterRequiresNickname = Assert<
  RegisterRequest extends { nickname: string } ? true : false
>;

type RefreshReturnsAuthResponse = Assert<Equal<RefreshResponse, AuthResponse>>;

type TokenExpiriesUseGatewayNumbers = Assert<
  TokenResponse["accessTokenExpiresAt"] extends number
    ? TokenResponse["refreshTokenExpiresAt"] extends number
      ? true
      : false
    : false
>;

type PresignUsesBackendUrls = Assert<
  PresignResponse extends { url: string; contentUrl: string } ? true : false
>;

type PresignOmitsLegacyPutUrl = Assert<
  "putUrl" extends keyof PresignResponse ? false : true
>;
