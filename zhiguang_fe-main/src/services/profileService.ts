import { apiFetch } from "./apiClient";
import type {
  AvatarUploadResponse,
  ProfileResponse,
  ProfileUpdateRequest
} from "@/types/profile";

const PROFILE_PREFIX = "/api/v1/profile";

const isRecord = (value: unknown): value is Record<string, unknown> =>
  typeof value === "object" && value !== null;

const isProfileGender = (value: unknown) =>
  value === "" ||
  value === "MALE" ||
  value === "FEMALE" ||
  value === "OTHER" ||
  value === "UNKNOWN";

const parseProfileResponse = (value: unknown): ProfileResponse => {
  if (
    !isRecord(value) ||
    typeof value.id !== "number" ||
    !Number.isFinite(value.id) ||
    typeof value.nickname !== "string" ||
    typeof value.avatar !== "string" ||
    typeof value.bio !== "string" ||
    typeof value.zgId !== "string" ||
    !isProfileGender(value.gender) ||
    typeof value.birthday !== "string" ||
    typeof value.school !== "string" ||
    typeof value.email !== "string" ||
    typeof value.phone !== "string" ||
    typeof value.tagsJson !== "string"
  ) {
    throw new Error("后端 profile 响应缺少必要字段");
  }
  return value as ProfileResponse;
};

const parseAvatarUploadResponse = (value: unknown): AvatarUploadResponse => {
  if (
    !isRecord(value) ||
    typeof value.url !== "string" ||
    typeof value.avatar !== "string" ||
    typeof value.objectKey !== "string"
  ) {
    throw new Error("后端头像上传响应缺少必要字段");
  }
  return value as AvatarUploadResponse;
};

export const profileService = {
  update: (payload: ProfileUpdateRequest) =>
    apiFetch<unknown>(`${PROFILE_PREFIX}/`, {
      method: "PATCH",
      body: payload
    }).then(parseProfileResponse),

  uploadAvatar: (file: File) => {
    const form = new FormData();
    form.append("file", file);
    return apiFetch<unknown>(`${PROFILE_PREFIX}/avatar`, {
      method: "POST",
      body: form
    }).then(parseAvatarUploadResponse);
  }
};
