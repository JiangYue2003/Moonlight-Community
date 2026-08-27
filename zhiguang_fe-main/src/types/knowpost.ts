export type CreateDraftResponse = {
  id: string;
};

export type PresignRequest = {
  scene: "knowpost_content" | "knowpost_image";
  postId: string;
  contentType: string;
  ext: string; // with dot, e.g. ".md" 
};

export type PresignResponse = {
  objectKey: string;
  url: string;
  headers: Record<string, string>;
  expiresIn: number;
  contentUrl: string;
};

export type ConfirmContentRequest = {
  objectKey: string;
  etag: string;
  size: number;
  sha256: string;
};

export type VisibleScope = "public" | "followers" | "school" | "private" | "unlisted";

export type UpdateKnowPostRequest = {
  title?: string;
  tagId?: number;
  tags?: string[];
  tagsSet?: boolean;
  imgUrls?: string[];
  imgUrlsSet?: boolean;
  visible?: VisibleScope;
  isTop?: boolean;
  description?: string;
};

export type FeedItem = {
  id: string;
  creatorId: number;
  title: string;
  description: string;
  contentUrl: string;
  tags: string[] | null;
  imgUrls: string[] | null;
  visible: VisibleScope;
  isTop: boolean;
  publishTime: number;
};

export type FeedResponse = {
  items: FeedItem[];
  page: number;
  size: number;
  hasMore: boolean;
};

export type KnowpostDetailResponse = {
  id: string;
  creatorId: number;
  title: string;
  description: string;
  tagId: number;
  tags: string[] | null;
  contentUrl: string;
  contentObjectKey: string;
  imgUrls: string[] | null;
  status: string;
  type: string;
  isTop: boolean;
  visible: VisibleScope;
  createTime: number;
  updateTime: number;
  publishTime: number;
};

// 点赞/取消点赞 响应
export type LikeActionResponse = {
  changed: boolean;
  liked: boolean;
};

// 收藏/取消收藏 响应
export type FavActionResponse = {
  changed: boolean;
  faved: boolean;
};

// 计数查询响应
export type CounterResponse = {
  entityType: string;
  entityId: string;
  counts: {
    like: number;
    fav: number;
  };
};
