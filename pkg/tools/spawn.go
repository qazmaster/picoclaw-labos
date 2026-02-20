package tools

import (
	"context"
	"fmt"
)

type SpawnTool struct {
	manager        *SubagentManager
	originChannel  string
	originChatID   string
	allowlistCheck func(targetAgentID string) bool
	callback       AsyncCallback // For async completion notification
}

func NewSpawnTool(manager *SubagentManager) *SpawnTool {
	return &SpawnTool{
		manager:       manager,
		originChannel: "cli",
		originChatID:  "direct",
	}
}

// SetCallback implements AsyncTool interface for async completion notification
func (t *SpawnTool) SetCallback(cb AsyncCallback) {
	t.callback = cb
}

func (t *SpawnTool) Name() string {
	return "spawn"
}

func (t *SpawnTool) Description() string {
	return "Spawn a subagent to handle a task in the background. Use this for complex or time-consuming tasks that can run independently. The subagent will complete the task and report back when done."
}

func (t *SpawnTool) Parameters() map[string]interface{} {
	return map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"task": map[string]interface{}{
				"type":        "string",
				"description": "The task for subagent to complete",
			},
			"label": map[string]interface{}{
				"type":        "string",
				"description": "Optional short label for the task (for display)",
			},
			"agent_id": map[string]interface{}{
				"type":        "string",
				"description": "Optional target agent ID to delegate the task to",
			},
			"ttl_minutes": map[string]interface{}{
				"type":        "number",
				"description": "Optional time-to-live in minutes. Subagent is terminated after this duration (default: 0 = no limit)",
			},
			"token_limit": map[string]interface{}{
				"type":        "number",
				"description": "Optional max token budget for the subagent. Limits total LLM tokens consumed (default: 0 = no limit)",
			},
		},
		"required": []string{"task"},
	}
}

func (t *SpawnTool) SetContext(channel, chatID string) {
	t.originChannel = channel
	t.originChatID = chatID
}

func (t *SpawnTool) SetAllowlistChecker(check func(targetAgentID string) bool) {
	t.allowlistCheck = check
}

func (t *SpawnTool) Execute(ctx context.Context, args map[string]interface{}) *ToolResult {
	task, ok := args["task"].(string)
	if !ok {
		return ErrorResult("task is required")
	}

	label, _ := args["label"].(string)
	agentID, _ := args["agent_id"].(string)

	// Parse TTL and token limit (SpawnStandard)
	var ttlMinutes int
	if v, ok := args["ttl_minutes"].(float64); ok && v > 0 {
		ttlMinutes = int(v)
	}
	var tokenLimit int
	if v, ok := args["token_limit"].(float64); ok && v > 0 {
		tokenLimit = int(v)
	}

	// Check allowlist if targeting a specific agent
	if agentID != "" && t.allowlistCheck != nil {
		if !t.allowlistCheck(agentID) {
			return ErrorResult(fmt.Sprintf("not allowed to spawn agent '%s'", agentID))
		}
	}

	if t.manager == nil {
		return ErrorResult("Subagent manager not configured")
	}

	// Pass callback to manager for async completion notification
	result, err := t.manager.Spawn(ctx, task, label, agentID, t.originChannel, t.originChatID, ttlMinutes, tokenLimit, t.callback)
	if err != nil {
		return ErrorResult(fmt.Sprintf("failed to spawn subagent: %v", err))
	}

	// Return AsyncResult since the task runs in background
	return AsyncResult(result)
}
