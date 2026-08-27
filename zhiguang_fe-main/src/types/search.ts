export type SearchHit = {
  contentId: string;
  contentType: string;
  title: string;
  description: string;
  snippet: string;
  tags: string[] | null;
  authorId: number;
  authorNickname: string;
  authorAvatar: string;
  likeCount: number;
  favoriteCount: number;
  viewCount: number;
  imgUrls: string[];
  isTop: boolean;
};

export type SearchResponse = {
  items: SearchHit[];
  nextAfter: string;
  hasMore: boolean;
};

export type SuggestResponse = {
  items: string[];
};
