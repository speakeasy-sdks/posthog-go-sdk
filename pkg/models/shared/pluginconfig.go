package shared

import (
	"time"
)

type PluginConfigInput struct {
	Enabled *bool                  `json:"enabled,omitempty" form:"name=enabled" multipartForm:"name=enabled"`
	Error   map[string]interface{} `json:"error,omitempty" form:"name=error,json" multipartForm:"name=error,json"`
	Order   int64                  `json:"order" form:"name=order" multipartForm:"name=order"`
	Plugin  int64                  `json:"plugin" form:"name=plugin" multipartForm:"name=plugin"`
}

type PluginConfig struct {
	Config          string                 `json:"config"`
	CreatedAt       time.Time              `json:"created_at"`
	DeliveryRate24h string                 `json:"delivery_rate_24h"`
	Enabled         *bool                  `json:"enabled,omitempty"`
	Error           map[string]interface{} `json:"error,omitempty"`
	ID              int64                  `json:"id"`
	Order           int64                  `json:"order"`
	Plugin          int64                  `json:"plugin"`
	PluginInfo      string                 `json:"plugin_info"`
	TeamID          int64                  `json:"team_id"`
}
