CREATE TABLE IF NOT EXISTS feed_author_delivery_tier (
    author_id BIGINT UNSIGNED NOT NULL,
    tier TINYINT UNSIGNED NOT NULL,
    promotion_reason VARCHAR(32) NOT NULL,
    observed_followers BIGINT UNSIGNED NOT NULL,
    promoted_at DATETIME(3) NOT NULL,
    updated_at DATETIME(3) NOT NULL,
    PRIMARY KEY (author_id),
    CONSTRAINT chk_feed_author_delivery_tier_bigv CHECK (tier = 1),
    CONSTRAINT chk_feed_author_delivery_tier_evidence CHECK (observed_followers > 1000)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
