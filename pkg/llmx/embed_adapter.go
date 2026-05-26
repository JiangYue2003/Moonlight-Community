// Package llmx 提供 EINO Embedder 的 float64→float32 适配层。
//
// EINO DashScope embedder 返回 [][]float64；pkg/esx 的 KnnSearch 和 rag-indexer
// 的 bulk 写入使用 float32（与 ES dense_vector 存储精度一致）。
// 此文件提供转换函数，其余 LLM/Embedding 能力直接使用 EINO 原生接口。
package llmx

import (
	"context"

	"github.com/cloudwego/eino/components/embedding"
)

// EmbedFloat32 调用 EINO Embedder 并将结果从 float64 转换为 float32。
// 分批由 EINO DashScope embedder 内部处理。
func EmbedFloat32(ctx context.Context, emb embedding.Embedder, texts []string) ([][]float32, error) {
	if len(texts) == 0 {
		return nil, nil
	}
	vecs64, err := emb.EmbedStrings(ctx, texts)
	if err != nil {
		return nil, err
	}
	out := make([][]float32, len(vecs64))
	for i, v := range vecs64 {
		f32 := make([]float32, len(v))
		for j, x := range v {
			f32[j] = float32(x)
		}
		out[i] = f32
	}
	return out, nil
}
