import { useEffect, useMemo, useRef, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import AppLayout from "@/components/layout/AppLayout";
import MainHeader from "@/components/layout/MainHeader";
import Tag from "@/components/common/Tag";
import SectionHeader from "@/components/common/SectionHeader";
import { ArrowRightIcon, SparkIcon } from "@/components/icons/Icon";
import AuthStatus from "@/features/auth/AuthStatus";
import styles from "./CourseDetailPage.module.css";
import { knowpostService } from "@/services/knowpostService";
import { useAuth } from "@/context/AuthContext";
import type { KnowpostDetailResponse } from "@/types/knowpost";
import ReactMarkdown from "react-markdown";
import remarkGfm from "remark-gfm";
import LikeFavBar from "@/components/common/LikeFavBar";
import FollowButton from "@/components/common/FollowButton";

const EMPTY_IMAGES: string[] = [];

const QUICK_PROMPTS = [
  "总结本文核心要点",
  "有哪些关键技术或实践建议？",
  "这篇知文适合哪类受众阅读？",
];

const CourseDetailPage = () => {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const { tokens, user } = useAuth();
  const [detail, setDetail] = useState<KnowpostDetailResponse | null>(null);
  const [activeImage, setActiveImage] = useState(0);
  const [contentText, setContentText] = useState<string>("");
  const [error, setError] = useState<string | null>(null);
  const [previewOpen, setPreviewOpen] = useState(false);
  const [previewIndex, setPreviewIndex] = useState(0);
  const rowRef = useRef<HTMLDivElement | null>(null);
  const [visibleCount, setVisibleCount] = useState<number>(0);
  const [contentError, setContentError] = useState<string | null>(null);
  const previewBoxRef = useRef<HTMLDivElement | null>(null);
  const [showNavLeft, setShowNavLeft] = useState(false);
  const [showNavRight, setShowNavRight] = useState(false);
  const [isTouch, setIsTouch] = useState(false);

  // RAG 问答状态
  const [ragQuestion, setRagQuestion] = useState<string>("");
  const [ragAnswer, setRagAnswer] = useState<string>("");
  const [ragLoading, setRagLoading] = useState<boolean>(false);
  const [ragError, setRagError] = useState<string | null>(null);
  const ragESRef = useRef<EventSource | null>(null);
  const [ragTopK, setRagTopK] = useState<number>(5);
  const [ragMaxTokens, setRagMaxTokens] = useState<number>(1024);
  const images = detail?.imgUrls ?? EMPTY_IMAGES;

  useEffect(() => {
    let cancelled = false;
    const run = async () => {
      if (!id) return;
      setError(null);
      try {
        const resp = await knowpostService.detail(id, tokens?.accessToken ?? undefined);
        if (cancelled) return;
        setDetail(resp);
        setActiveImage(0);
        // 异步加载正文内容
        if (resp.contentUrl) {
          const allowAnonymous = resp.visible === "public";
          if (allowAnonymous || !!tokens?.accessToken) {
            try {
              const text = await fetch(resp.contentUrl, { credentials: "omit" }).then(r => {
                if (!r.ok) throw new Error(`HTTP ${r.status}`);
                return r.text();
              });
              if (!cancelled) {
                setContentText(text);
                setContentError(null);
              }
            } catch (e) {
              if (!cancelled) setContentError("正文暂不可读，可能为非公开或跨域受限");
            }
          } else {
            setContentError("该知文非公开，请登录后查看正文");
          }
        }
      } catch (err) {
        const msg = err instanceof Error ? err.message : "加载失败";
        if (!cancelled) setError(msg);
      }
    };
    run();
    return () => { cancelled = true; };
  }, [id, tokens?.accessToken]);

  // 计算一行可展示的图片数量
  useEffect(() => {
    const calc = () => {
      const el = rowRef.current;
      if (!el) return;
      const width = el.clientWidth;
      const itemW = 190;
      const gap = 14;
      const count = Math.max(1, Math.floor((width + gap) / (itemW + gap)));
      setVisibleCount(count);
    };
    calc();
    window.addEventListener("resize", calc);
    return () => window.removeEventListener("resize", calc);
  }, [detail?.imgUrls]);

  useEffect(() => {
    const touch = "ontouchstart" in window || navigator.maxTouchPoints > 0;
    setIsTouch(touch);
    if (touch) {
      setShowNavLeft(true);
      setShowNavRight(true);
    }
  }, []);

  const handlePreviewMouseMove = (e: React.MouseEvent<HTMLDivElement>) => {
    if (isTouch) return;
    const box = previewBoxRef.current;
    if (!box) return;
    const rect = box.getBoundingClientRect();
    const x = e.clientX - rect.left;
    const threshold = Math.max(60, Math.min(120, rect.width * 0.08));
    setShowNavLeft(x < threshold);
    setShowNavRight(x > rect.width - threshold);
  };

  const handlePreviewMouseLeave = () => {
    if (isTouch) return;
    setShowNavLeft(false);
    setShowNavRight(false);
  };

  const openPreview = (index: number) => {
    setPreviewIndex(index);
    setPreviewOpen(true);
  };

  const prevImage = () => {
    if (!images.length) return;
    setPreviewIndex((i) => (i - 1 + images.length) % images.length);
  };

  const nextImage = () => {
    if (!images.length) return;
    setPreviewIndex((i) => (i + 1) % images.length);
  };

  // 启动 RAG 流式问答
  const executeRag = (questionToAsk: string) => {
    if (!id) return;
    const q = questionToAsk.trim();
    if (!q) return;
    if (detail && detail.visible !== "public") {
      setRagError("仅公开知文支持问答");
      return;
    }
    setRagError(null);
    setRagAnswer("");
    if (ragESRef.current) {
      try { ragESRef.current.close(); } catch {}
      ragESRef.current = null;
    }
    const url = `/api/v1/knowposts/${id}/qa/stream?question=${encodeURIComponent(q)}&topK=${ragTopK}&maxTokens=${ragMaxTokens}`;
    const es = new EventSource(url);
    ragESRef.current = es;
    setRagLoading(true);
    es.onmessage = (e) => {
      setRagAnswer((prev) => prev + (e.data ?? ""));
    };
    es.onerror = () => {
      setRagLoading(false);
      try { es.close(); } catch {}
      ragESRef.current = null;
    };
  };

  const startRag = () => {
    executeRag(ragQuestion);
  };

  const handlePromptChipClick = (prompt: string) => {
    setRagQuestion(prompt);
    executeRag(prompt);
  };

  const stopRag = () => {
    if (ragESRef.current) {
      try { ragESRef.current.close(); } catch {}
      ragESRef.current = null;
    }
    setRagLoading(false);
  };

  useEffect(() => {
    return () => {
      if (ragESRef.current) {
        try { ragESRef.current.close(); } catch {}
        ragESRef.current = null;
      }
    };
  }, []);

  return (
    <AppLayout
      header={
        <MainHeader
          headline={detail?.title ?? "知文详情"}
          subtitle=""
          rightSlot={<AuthStatus />}
        />
      }
      variant="cardless"
    >
      <article className={styles.detailCard}>
        {error ? <div style={{ color: "var(--color-danger)" }}>{error}</div> : null}
        
        {images.length ? (
          <div ref={rowRef} className={styles.imageRow}>
            {(images.slice(0, visibleCount)).map((src, idx) => {
              const isLastVisible = idx === visibleCount - 1 && images.length > visibleCount;
              return (
                <div key={src + idx} className={styles.imageItem} onClick={() => openPreview(idx)}>
                  <img className={styles.image} src={src} alt={detail?.title ?? ""} />
                  {isLastVisible ? (
                    <div className={styles.moreBadge}>+{images.length - visibleCount}</div>
                  ) : null}
                </div>
              );
            })}
          </div>
        ) : null}

        <div className={styles.titleBlock}>
          <h1 className={styles.title}>{detail?.title || "无标题知文"}</h1>
          
          <div className={styles.metaHeader}>
            <div className={styles.authorBlock}>
              <div className={styles.authorAvatarFallback}>
                {detail ? String(detail.creatorId).slice(-2) : "U"}
              </div>
              <div className={styles.authorMeta}>
                <span className={styles.authorName}>{detail ? `用户 ${detail.creatorId}` : "创作者"}</span>
                {detail?.publishTime ? (
                  <span className={styles.publishTime}>发布于 {new Date(detail.publishTime).toLocaleDateString("zh-CN")}</span>
                ) : null}
              </div>
              {detail && user?.id !== detail.creatorId ? <FollowButton targetUserId={detail.creatorId} /> : null}
            </div>

            <div className={styles.bottomBar}>
              {detail ? (
                <LikeFavBar
                  entityId={detail.id}
                  fetchCounts
                />
              ) : null}
            </div>
          </div>

          {detail?.tags && detail.tags.length > 0 ? (
            <div className={styles.tagList}>
              {detail.tags.map(tag => (
                <Tag key={tag}>#{tag}</Tag>
              ))}
            </div>
          ) : null}
        </div>

        <div className={styles.contentRow}>
          <div className={styles.contentMain}>
            <SectionHeader title="知文正文" subtitle="" />
            <div className={`${styles.body} ${styles.markdown}`}>
              {contentText ? (
                <ReactMarkdown
                  remarkPlugins={[remarkGfm]}
                  components={{
                    a: ({ node, ...props }) => (
                      <a {...props} target="_blank" rel="noreferrer" />
                    ),
                    img: ({ node, ...props }) => (
                      <img {...props} style={{ maxWidth: "100%", borderRadius: 12 }} />
                    ),
                  }}
                >
                  {contentText}
                </ReactMarkdown>
              ) : (
                <div style={{ color: "var(--color-text-subtle)", padding: "20px 0" }}>暂无内容</div>
              )}
            </div>
            {contentError ? (
              <div style={{ color: "var(--color-danger)", marginTop: 12 }}>
                {contentError} {detail?.contentUrl ? (<a href={detail.contentUrl} target="_blank" rel="noreferrer">查看原文</a>) : null}
              </div>
            ) : null}
          </div>

          <aside className={styles.ragPanel}>
            <div className={styles.ragHeader}>
              <div className={styles.ragHeaderTitle}>
                <SparkIcon width={17} height={17} style={{ color: "var(--color-primary)" }} />
                <span>知光 AI 智能伴读</span>
              </div>
              <div className={styles.ragStatusPill}>
                <span className={styles.ragDot} />
                <span>实时索引就绪</span>
              </div>
            </div>

            <div className={styles.ragBody}>
              <div className={styles.ragChips}>
                {QUICK_PROMPTS.map((prompt) => (
                  <button
                    key={prompt}
                    type="button"
                    className={styles.ragChip}
                    onClick={() => handlePromptChipClick(prompt)}
                  >
                    {prompt}
                  </button>
                ))}
              </div>

              <textarea
                className={styles.ragTextarea}
                placeholder="围绕本知文提问，例如：这篇知文的核心观点是什么？"
                value={ragQuestion}
                onChange={(e) => setRagQuestion(e.target.value)}
                onKeyDown={(e) => {
                  if (e.key === "Enter" && (e.metaKey || e.ctrlKey)) {
                    e.preventDefault();
                    startRag();
                  }
                }}
              />
              <div className={styles.ragControls}>
                <button
                  type="button"
                  className={`${styles.ragBtn} ${styles.ragBtnPrimary}`}
                  onClick={startRag}
                  disabled={ragLoading || !ragQuestion.trim()}
                >
                  {ragLoading ? "思考生成中..." : "发送提问"}
                </button>
                <button
                  type="button"
                  className={`${styles.ragBtn} ${styles.ragBtnGhost}`}
                  onClick={stopRag}
                  disabled={!ragLoading}
                >
                  停止
                </button>
              </div>
              <div className={styles.ragHint}>
                基于知光 RAG 引擎实时检索，仅公开知文支持深度伴读回答。快捷键 ⌘+Enter 发送。
              </div>
              {ragError ? (
                <div style={{ color: "var(--color-danger)", fontSize: 13 }}>{ragError}</div>
              ) : null}
              <div className={styles.ragAnswer}>
                {ragAnswer ? (
                  <div className={styles.markdown}>
                    <ReactMarkdown
                      remarkPlugins={[remarkGfm]}
                      components={{
                        a: ({ node, ...props }) => (
                          <a {...props} target="_blank" rel="noreferrer" />
                        ),
                        img: ({ node, ...props }) => (
                          <img {...props} style={{ maxWidth: "100%", borderRadius: 12 }} />
                        ),
                      }}
                    >
                      {ragAnswer}
                    </ReactMarkdown>
                  </div>
                ) : (
                  <div className={styles.ragPlaceholder}>
                    {ragLoading ? "知光 AI 正在查阅索引并组织回答..." : "输入问题或点击上方快捷提示词，获取 AI 深度解答"}
                  </div>
                )}
              </div>
            </div>
          </aside>
        </div>

        {previewOpen && images.length ? (
          <div className={styles.previewOverlay} onClick={() => setPreviewOpen(false)}>
            <div
              className={styles.previewBox}
              ref={previewBoxRef}
              onMouseMove={handlePreviewMouseMove}
              onMouseLeave={handlePreviewMouseLeave}
              onClick={(e) => e.stopPropagation()}
            >
              <img className={styles.previewImage} src={images[previewIndex]} alt={detail?.title ?? ""} />
              <button
                type="button"
                className={`${styles.navButton} ${styles.navButtonLeft} ${showNavLeft ? styles.navButtonVisible : ""}`}
                onClick={(e) => { e.stopPropagation(); prevImage(); }}
                aria-label="上一张"
              >
                <ArrowRightIcon width={24} height={24} style={{ transform: "rotate(180deg)" }} />
              </button>
              <button
                type="button"
                className={`${styles.navButton} ${styles.navButtonRight} ${showNavRight ? styles.navButtonVisible : ""}`}
                onClick={(e) => { e.stopPropagation(); nextImage(); }}
                aria-label="下一张"
              >
                <ArrowRightIcon width={24} height={24} />
              </button>
              <button
                type="button"
                className={styles.closeButton}
                onClick={(e) => { e.stopPropagation(); setPreviewOpen(false); }}
                aria-label="关闭"
              >
                ×
              </button>
            </div>
          </div>
        ) : null}
      </article>
    </AppLayout>
  );
};

export default CourseDetailPage;
