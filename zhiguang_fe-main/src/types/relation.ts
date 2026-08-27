export type RelationStatusResponse = {
  following: boolean;
  followedBy: boolean;
  mutual: boolean;
};

export type RelationProfile = {
  id: number;
  nickname: string;
  avatar: string;
  zgId: string;
  bio: string;
};

export type RelationListResponse = {
  items: RelationProfile[];
  nextCursor: number;
  hasMore: boolean;
};

export type RelationCountersResponse = {
  followings: number;
  followers: number;
  posts: number;
  likesReceived: number;
};
