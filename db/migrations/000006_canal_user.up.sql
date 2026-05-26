-- Canal 复制账号（dev/集成环境用）。
-- 生产环境应该用人工流程在 DBA 控制下创建，不走 migration。
-- 阶段3 把它放进 migration 是为了让 `make migrate-up` 一把跑通。

CREATE USER IF NOT EXISTS 'canal'@'%' IDENTIFIED WITH mysql_native_password BY 'canal';
GRANT SELECT, REPLICATION SLAVE, REPLICATION CLIENT ON *.* TO 'canal'@'%';
FLUSH PRIVILEGES;
