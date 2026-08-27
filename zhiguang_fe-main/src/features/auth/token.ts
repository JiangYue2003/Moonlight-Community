export const getAccessTokenExpiry = (value: number) => {
  if (!Number.isFinite(value) || value <= 0) {
    throw new Error("后端返回了无效的 accessTokenExpiresAt");
  }
  return value;
};
