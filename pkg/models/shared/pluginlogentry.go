package shared

import (
	"time"
)

type PluginLogEntrySourceEnum string

const (
	PluginLogEntrySourceEnumSystem  PluginLogEntrySourceEnum = "SYSTEM"
	PluginLogEntrySourceEnumPlugin  PluginLogEntrySourceEnum = "PLUGIN"
	PluginLogEntrySourceEnumConsole PluginLogEntrySourceEnum = "CONSOLE"
)

type PluginLogEntryTypeEnum string

const (
	PluginLogEntryTypeEnumDebug PluginLogEntryTypeEnum = "DEBUG"
	PluginLogEntryTypeEnumLog   PluginLogEntryTypeEnum = "LOG"
	PluginLogEntryTypeEnumInfo  PluginLogEntryTypeEnum = "INFO"
	PluginLogEntryTypeEnumWarn  PluginLogEntryTypeEnum = "WARN"
	PluginLogEntryTypeEnumError PluginLogEntryTypeEnum = "ERROR"
)

type PluginLogEntry struct {
	ID             string                   `json:"id"`
	InstanceID     string                   `json:"instance_id"`
	Message        string                   `json:"message"`
	PluginConfigID int64                    `json:"plugin_config_id"`
	PluginID       int64                    `json:"plugin_id"`
	Source         PluginLogEntrySourceEnum `json:"source"`
	TeamID         int64                    `json:"team_id"`
	Timestamp      time.Time                `json:"timestamp"`
	Type           PluginLogEntryTypeEnum   `json:"type"`
}
