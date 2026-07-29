#!/bin/bash

set -e

echo "========================================="
echo "  知光 All-in-One 容器启动"
echo "========================================="

# 初始化 MySQL 数据目录
if [ ! -d "/var/lib/mysql/mysql" ]; then
    echo "初始化 MySQL..."
    mysql_install_db --user=root --datadir=/var/lib/mysql
fi

# 启动 MySQL 并设置密码
echo "启动 MySQL..."
/usr/bin/mysqld --user=root --datadir=/var/lib/mysql --bind-address=0.0.0.0 &
MYSQL_PID=$!

# 等待 MySQL 启动
echo "等待 MySQL 就绪..."
for i in {1..30}; do
    if mysqladmin ping -h 127.0.0.1 --silent; then
        echo "MySQL 已就绪"
        break
    fi
    sleep 1
done

# 设置 MySQL 密码和创建数据库
echo "配置 MySQL..."
mysql -h 127.0.0.1 <<EOF
ALTER USER 'root'@'localhost' IDENTIFIED BY 'Zz123456';
CREATE USER IF NOT EXISTS 'root'@'%' IDENTIFIED BY 'Zz123456';
GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' WITH GRANT OPTION;
FLUSH PRIVILEGES;
CREATE DATABASE IF NOT EXISTS zhiguang CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE USER IF NOT EXISTS 'canal'@'%' IDENTIFIED WITH mysql_native_password BY 'canal';
GRANT SELECT, REPLICATION SLAVE, REPLICATION CLIENT ON *.* TO 'canal'@'%';
FLUSH PRIVILEGES;
EOF

# 执行数据库迁移（如果有）
if [ -d "/app/db/migrations" ]; then
    echo "执行数据库迁移..."
    for sql in /app/db/migrations/*.sql; do
        if [ -f "$sql" ]; then
            echo "执行 $sql"
            mysql -h 127.0.0.1 -uroot -pZz123456 zhiguang < "$sql" || true
        fi
    done
fi

# 停止 MySQL，让 supervisor 管理
kill $MYSQL_PID
wait $MYSQL_PID || true

echo "启动所有服务（通过 Supervisor）..."
exec /usr/bin/supervisord -c /etc/supervisord.conf
