import { useEffect, useMemo, useState } from "react";
import { useNavigate } from "react-router-dom";
import AppLayout from "@/components/layout/AppLayout";
import MainHeader from "@/components/layout/MainHeader";
import CourseCard from "@/components/cards/CourseCard";
import LikeFavBar from "@/components/common/LikeFavBar";
import SearchBar from "@/components/common/SearchBar";
import { knowpostService } from "@/services/knowpostService";
import AuthStatus from "@/features/auth/AuthStatus";
import type { FeedItem } from "@/types/knowpost";
import styles from "./HomePage.module.css";

const topics = ["全部", "学习方法", "科技", "设计", "生活方式", "阅读"];

const HomePage = () => {
  const [items, setItems] = useState<FeedItem[]>([]);
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  const [topic, setTopic] = useState("全部");
  const [query, setQuery] = useState("");
  const navigate = useNavigate();

  useEffect(() => {
    let cancelled = false;
    const run = async () => {
      setLoading(true);
      setError(null);
      try {
        const resp = await knowpostService.feed(1, 20);
        if (!cancelled) setItems(resp.items);
      } catch (err) {
        const msg = err instanceof Error ? err.message : "加载失败";
        if (!cancelled) setError(msg);
      } finally {
        if (!cancelled) setLoading(false);
      }
    };
    void run();
    return () => { cancelled = true; };
  }, []);

  const visibleItems = useMemo(() => {
    const keyword = query.trim().toLowerCase();
    return items.filter((item) => {
      const text = (item.title + " " + (item.description ?? "") + " " + (item.tags ?? []).join(" ")).toLowerCase();
      const matchesQuery = !keyword || text.includes(keyword);
      const matchesTopic = topic === "全部" || (item.tags ?? []).some((tag) => tag.includes(topic));
      return matchesQuery && matchesTopic;
    });
  }, [items, query, topic]);

  return (
    <AppLayout
      header={
        <MainHeader
          headline="发现"
          variant="discovery"
          rightSlot={<AuthStatus />}
          centerSlot={
            <SearchBar
              placeholder="搜索你感兴趣的知识"
              value={query}
              onChange={setQuery}
              onSubmit={() => navigate(query.trim() ? "/search?q=" + encodeURIComponent(query.trim()) : "/search")}
            />
          }
        >
          <div className={styles.topicRow} aria-label="内容主题">
            {topics.map((item) => (
              <button key={item} type="button" className={item === topic ? styles.topicActive : styles.topic} onClick={() => setTopic(item)}>
                {item}
              </button>
            ))}
          </div>
        </MainHeader>
      }
    >
      {error ? <div className={styles.stateError}>{error}</div> : null}
      <div className={styles.masonry}>
        {loading && items.length === 0 ? Array.from({ length: 10 }, (_, index) => (
          <div key={`skeleton-${index}`} className={styles.masonryItem}>
            <div className={styles.skeletonCard}>
              <div className={`${styles.skeletonCover} ${index % 3 === 1 ? styles.skeletonTall : ""}`} />
              <div className={styles.skeletonLine} />
              <div className={styles.skeletonMeta} />
            </div>
          </div>
        )) : null}
        {visibleItems.map(item => (
          <div key={item.id} className={styles.masonryItem}>
            <CourseCard id={item.id} title={item.title} summary={item.description ?? ""} tags={item.tags ?? []} teacher={{ name: "用户 " + item.creatorId }} coverImage={item.imgUrls?.find(Boolean)} to={"/post/" + item.id} footerExtra={<LikeFavBar entityId={item.id} compact fetchCounts />} />
          </div>
        ))}
        {loading && items.length > 0 ? <div className={styles.stateCard}>正在加载...</div> : null}
        {!loading && visibleItems.length === 0 ? <div className={styles.stateCard}>换一个主题或关键词试试</div> : null}
      </div>
    </AppLayout>
  );
};

export default HomePage;
