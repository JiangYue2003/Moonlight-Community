import { useEffect, useState } from "react";
import AppLayout from "@/components/layout/AppLayout";
import MainHeader from "@/components/layout/MainHeader";
import CourseCard from "@/components/cards/CourseCard";
import LikeFavBar from "@/components/common/LikeFavBar";
import { knowpostService } from "@/services/knowpostService";
import AuthStatus from "@/features/auth/AuthStatus";
import type { FeedItem } from "@/types/knowpost";
import styles from "./HomePage.module.css";

const HomePage = () => {
  const [items, setItems] = useState<FeedItem[]>([]);
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    let cancelled = false;
    const run = async () => {
      setLoading(true);
      setError(null);
      try {
        const resp = await knowpostService.feed(1, 20);
        if (!cancelled) {
          setItems(resp.items);
        }
      } catch (err) {
        const msg = err instanceof Error ? err.message : "加载失败";
        if (!cancelled) setError(msg);
      } finally {
        if (!cancelled) setLoading(false);
      }
    };
    run();
    return () => {
      cancelled = true;
    };
  }, []);

  return (
    <AppLayout
      header={
        <MainHeader
          headline="知光 · 让思想有温度，让知识会发光"
          rightSlot={<AuthStatus />}
        />
      }
    >
      {error ? <div>{error}</div> : null}
      <div className={styles.masonry}>
        {items.map(item => (
          <div key={item.id} className={styles.masonryItem}>
            <CourseCard
              id={item.id}
              title={item.title}
              summary={item.description ?? ""}
              tags={item.tags ?? []}
              teacher={{ name: `用户 ${item.creatorId}` }}
              coverImage={item.imgUrls?.find(Boolean)}
              to={`/post/${item.id}`}
              footerExtra={<LikeFavBar entityId={item.id} compact fetchCounts />}
            />
          </div>
        ))}
        {loading ? <div className={styles.masonryItem}><div>加载中…</div></div> : null}
        {!loading && items.length === 0 ? (
          <div className={styles.masonryItem}><div>暂无内容</div></div>
        ) : null}
      </div>
    </AppLayout>
  );
};

export default HomePage;
