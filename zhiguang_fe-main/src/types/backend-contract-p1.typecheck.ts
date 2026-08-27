import { knowpostService } from "@/services/knowpostService";
import { relationService } from "@/services/relationService";
import type {
  FeedItem,
  KnowpostDetailResponse,
  UpdateKnowPostRequest
} from "./knowpost";
import type { SearchResponse } from "./search";

type Assert<T extends true> = T;

type MetadataSupportsPresenceFlags = Assert<
  UpdateKnowPostRequest extends {
    tagsSet?: boolean;
    imgUrlsSet?: boolean;
  } ? true : false
>;

type FollowReturnsNoContent = Assert<
  ReturnType<typeof relationService.follow> extends Promise<void> ? true : false
>;

type FollowingReturnsPage = Assert<
  Awaited<ReturnType<typeof relationService.following>> extends {
    items: unknown[];
    nextCursor: number;
    hasMore: boolean;
  } ? true : false
>;

type PublicRelationStatusNeedsNoToken = Assert<
  [toUserId: number] extends Parameters<typeof relationService.status>
    ? true
    : false
>;

type PrivateRelationListRequiresToken = Assert<
  Parameters<typeof relationService.following>[1] extends string ? true : false
>;

type SearchUsesContentId = Assert<
  SearchResponse["items"][number] extends { contentId: string } ? true : false
>;

type SearchAllowsGatewayNullTags = Assert<
  null extends SearchResponse["items"][number]["tags"] ? true : false
>;

type SearchCursorUsesGatewayString = Assert<
  SearchResponse["nextAfter"] extends string ? true : false
>;

type FeedUsesGatewayFields = Assert<
  FeedItem extends {
    creatorId: number;
    imgUrls: string[] | null;
    contentUrl: string;
  } ? true : false
>;

type DetailUsesGatewayFields = Assert<
  KnowpostDetailResponse extends {
    creatorId: number;
    imgUrls: string[] | null;
    contentObjectKey: string;
  } ? true : false
>;

type MetadataUpdateReturnsDetail = Assert<
  Awaited<ReturnType<typeof knowpostService.update>> extends KnowpostDetailResponse
    ? true
    : false
>;

type PublishReturnsDetail = Assert<
  Awaited<ReturnType<typeof knowpostService.publish>> extends KnowpostDetailResponse
    ? true
    : false
>;

type PublicCounterNeedsNoToken = Assert<
  [entityId: string] extends Parameters<typeof knowpostService.counters> ? true : false
>;
