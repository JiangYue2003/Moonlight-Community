# Elasticsearch IK Plugin

需要下载对应版本（与 docker-compose 中 ES 镜像版本严格一致：9.0.3）的 IK 分词插件，并解压到本目录：

```
deploy/compose/es-plugins/analysis-ik/
├── plugin-descriptor.properties
├── elasticsearch-analysis-ik-9.0.3.jar
├── ...其它依赖 jar
└── config/
    ├── IKAnalyzer.cfg.xml
    ├── main.dic
    └── ...
```

下载地址（任选其一）：
- 官方仓库 release：`https://github.com/infinilabs/analysis-ik/releases`
- ES 插件镜像市场（可用版本与 ES 一致）

放置完成后：

```bash
docker compose -f deploy/compose/docker-compose.dev.yml up -d elasticsearch
docker compose -f deploy/compose/docker-compose.dev.yml exec elasticsearch curl -s http://localhost:9200/_cat/plugins
# 期望看到 analysis-ik
```

如果暂时不需要中文分词（仅冒烟），可让此目录为空，ES 仍然能起，只是 mapping 中 `ik_max_word` / `ik_smart` 分析器会报错——此时不要建 `zhiguang_content_index`。
