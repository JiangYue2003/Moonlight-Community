import type { AuthUserResponse } from "./auth";
import type { ProfileResponse, ProfileUpdateRequest } from "./profile";
import type { RelationCountersResponse } from "./relation";
import { profileService } from "@/services/profileService";

type Assert<T extends true> = T;

type ProfileUpdateOmitsPhone = Assert<
  "phone" extends keyof ProfileUpdateRequest ? false : true
>;

type ProfileUpdateOmitsEmail = Assert<
  "email" extends keyof ProfileUpdateRequest ? false : true
>;

type ProfileUpdateUsesTagsJson = Assert<
  ProfileUpdateRequest extends { tagsJson?: string } ? true : false
>;

type ProfileUpdateOmitsLegacyTagJson = Assert<
  "tagJson" extends keyof ProfileUpdateRequest ? false : true
>;

type ProfileResponseUsesZgId = Assert<
  "zgId" extends keyof ProfileResponse ? true : false
>;

type ProfileResponseUsesTagsJson = Assert<
  "tagsJson" extends keyof ProfileResponse ? true : false
>;

type AuthUserUsesZgId = Assert<
  "zgId" extends keyof AuthUserResponse ? true : false
>;

type AuthUserUsesTagsJson = Assert<
  "tagsJson" extends keyof AuthUserResponse ? true : false
>;

type AuthUserOmitsLegacyZhId = Assert<
  "zhId" extends keyof AuthUserResponse ? false : true
>;

type AuthUserOmitsLegacyTagJson = Assert<
  "tagJson" extends keyof AuthUserResponse ? false : true
>;

type AvatarUploadUsesGatewayResponse = Assert<
  Awaited<ReturnType<typeof profileService.uploadAvatar>> extends {
    url: string;
    avatar: string;
    objectKey: string;
  } ? true : false
>;

type CountersExposeRealLikesReceived = Assert<
  RelationCountersResponse extends { likesReceived: number } ? true : false
>;

type CountersOmitHardcodedLikedPosts = Assert<
  "likedPosts" extends keyof RelationCountersResponse ? false : true
>;

type CountersOmitHardcodedFavedPosts = Assert<
  "favedPosts" extends keyof RelationCountersResponse ? false : true
>;
