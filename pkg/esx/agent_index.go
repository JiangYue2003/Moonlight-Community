package esx

import (
	_ "embed"
	"encoding/json"
)

//go:embed mapping/agent_knowledge.json
var agentKnowledgeMappingJSON []byte

// AgentKnowledgeMapping 返回 AGENT 知识索引 mapping（用户隔离 + 向量检索）。
func AgentKnowledgeMapping() json.RawMessage { return agentKnowledgeMappingJSON }
