package config

import (
	llmapiapp "github.com/zhiguang/zhiguang-go/services/llm/api/app"
	llmragindexerapp "github.com/zhiguang/zhiguang-go/services/llm/ragindexer/app"
)

type Config struct {
	DisableAPI bool `json:",default=false"`
	Api        llmapiapp.Config
	RagIndexer llmragindexerapp.Config
}
