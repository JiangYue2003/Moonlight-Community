import { useCallback, useEffect, useRef, useState } from "react";
import { useSearchParams } from "react-router-dom";
import AppLayout from "@/components/layout/AppLayout";
import MainHeader from "@/components/layout/MainHeader";
import SectionHeader from "@/components/common/SectionHeader";
import SearchBar from "@/components/common/SearchBar";
import AuthStatus from "@/features/auth/AuthStatus";
import styles from "./SearchPage.module.css";
import { searchService } from "@/services/searchService";
import type { SearchHit } from "@/types/search";
import CourseCard from "@/components/cards/CourseCard";
import LikeFavBar from "@/components/common/LikeFavBar";
import feedStyles from "./HomePage.module.css";
import { useAuth } from "@/context/AuthContext";

const SearchPage = () => {
  const [searchParams] = useSearchParams();
  const queryFromUrl = searchParams.get("q")?.trim() ?? "";
  const [q, setQ] = useState(queryFromUrl);
  const [tags] = useState(""); // 逗号分隔
  const [size] = useState<number>(20);
  const [items, setItems] = useState<SearchHit[]>([]);
  const [after, setAfter] = useState<string | null>(null);
  const [hasMore, setHasMore] = useState<boolean>(false);
  const [loading, setLoading] = useState(false);
  const [suggestions, setSuggestions] = useState<string[]>([]);
  const [suggestLoading, setSuggestLoading] = useState(false);
  const debounceRef = useRef<number | null>(null);
  const { user } = useAuth();
  const [showLoginHint, setShowLoginHint] = useState(false);

  const executeSearch = useCallback(async (keyword: string) => {
    const text = keyword.trim();
    if (!text) return;
    if (!user) {
      setShowLoginHint(true);
    }
    setQ(text);
    setLoading(true);
    try {
      const resp = await searchService.query({ q: text, size, tags: tags.trim() || undefined });
      setItems(resp.items);
      setAfter(resp.nextAfter || null);
      setHasMore(!!resp.hasMore);
    } catch {
      setItems([]);
      setAfter(null);
      setHasMore(false);
    } finally {
      setLoading(false);
    }
  }, [size, tags, user]);

  useEffect(() => {
    if (!queryFromUrl) return;
    void executeSearch(queryFromUrl);
  }, [executeSearch, queryFromUrl]);

  return (
    <AppLayout
      header={
        <MainHeader
          headline="搜索"
          variant="discovery"
          rightSlot={<AuthStatus />}
          centerSlot={<SearchBar
            placeholder="搜索你想学习的知识..."
            value={q}
            suggestions={suggestions}
            suggestLoading={suggestLoading}
            onSuggestionClick={(s) => {
              executeSearch(s);
            }}
            onChange={(val) => {
              setQ(val);
              // 前缀联想：300ms 防抖
              if (debounceRef.current) window.clearTimeout(debounceRef.current);
              debounceRef.current = window.setTimeout(async () => {
                if (!val.trim()) { setSuggestions([]); return; }
                try {
                  setSuggestLoading(true);
                  const resp = await searchService.suggest(val.trim(), 10);
                  setSuggestions(resp.items);
                } catch {
                  setSuggestions([]);
                } finally {
                  setSuggestLoading(false);
                }
              }, 300);
            }}
            onSubmit={() => executeSearch(q)}
          />}
        />
      }
    >
      <>
        {showLoginHint && !user ? (
          <div className={styles.loginHint}>
            当前为未登录状态，登录后可获得更完整的推荐与学习记录。
          </div>
        ) : null}
        <SectionHeader title="搜索结果" subtitle={loading ? "加载中..." : items.length ? `共 ${items.length} 条（可能有更多）` : "请输入关键词后搜索"} />
        <div className={feedStyles.masonry}>
          {loading && items.length === 0 ? Array.from({ length: 10 }, (_, index) => (
            <div key={`skeleton-${index}`} className={feedStyles.masonryItem}>
              <div className={feedStyles.skeletonCard}>
                <div className={`${feedStyles.skeletonCover} ${index % 3 === 1 ? feedStyles.skeletonTall : ""}`} />
                <div className={feedStyles.skeletonLine} />
                <div className={feedStyles.skeletonMeta} />
              </div>
            </div>
          )) : null}
          {items.map(item => (
            <div key={item.contentId} className={feedStyles.masonryItem}>
              <CourseCard
                id={item.contentId}
                title={item.title}
                summary={item.description ?? ""}
                tags={item.tags ?? []}
                isTop={item.isTop}
                teacher={{ name: item.authorNickname || `用户 ${item.authorId}`, avatarUrl: item.authorAvatar || undefined }}
                coverImage={item.imgUrls?.find(Boolean)}
                to={`/post/${item.contentId}`}
                footerExtra={<LikeFavBar entityId={item.contentId} compact initialCounts={{ like: item.likeCount ?? 0, fav: item.favoriteCount ?? 0 }} />}
              />
            </div>
          ))}
        </div>
        {hasMore ? (
          <button
            className={styles.loadMoreBtn}
            type="button"
            onClick={async () => {
              if (!q.trim() || !after) return;
              setLoading(true);
              try {
                const resp = await searchService.query({ q: q.trim(), size, tags: tags.trim() || undefined, after });
                setItems(prev => [...prev, ...resp.items]);
                setAfter(resp.nextAfter || null);
                setHasMore(!!resp.hasMore);
              } catch {
                // 保持已有数据
              } finally {
                setLoading(false);
              }
            }}
          >加载更多</button>
        ) : null}
      </>
    </AppLayout>
  );
};

export default SearchPage;
