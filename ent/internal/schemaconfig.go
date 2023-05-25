// Code generated by ent, DO NOT EDIT.

package internal

import "context"

// SchemaConfig represents alternative schema names for all tables
// that can be passed at runtime.
type SchemaConfig struct {
	ChatHistories                        string // ChatHistories table.
	LogChatHistoriesRecap                string // LogChatHistoriesRecap table.
	LogSummarizations                    string // LogSummarizations table.
	MetricOpenAIChatCompletionTokenUsage string // MetricOpenAIChatCompletionTokenUsage table.
	SlackOAuthCredentials                string // SlackOAuthCredentials table.
	TelegramChatFeatureFlags             string // TelegramChatFeatureFlags table.
}

type schemaCtxKey struct{}

// SchemaConfigFromContext returns a SchemaConfig stored inside a context, or empty if there isn't one.
func SchemaConfigFromContext(ctx context.Context) SchemaConfig {
	config, _ := ctx.Value(schemaCtxKey{}).(SchemaConfig)
	return config
}

// NewSchemaConfigContext returns a new context with the given SchemaConfig attached.
func NewSchemaConfigContext(parent context.Context, config SchemaConfig) context.Context {
	return context.WithValue(parent, schemaCtxKey{}, config)
}
