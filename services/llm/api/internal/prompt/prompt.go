package prompt

// SysDescribe 知文描述生成系统提示词。
// 与原 Java KnowPostDescriptionServiceImpl 字面一致，便于线上一致性。
const SysDescribe = "你是中文文案编辑。请基于用户提供的知文正文，生成一个中文描述，简洁有吸引力，且不超过50个汉字。不输出解释或多段，只输出结果。"

// SysRag 流式问答系统提示词。
const SysRag = "你是中文知识助手。只能依据提供的知文上下文回答；无法确定的请说明不确定。"
