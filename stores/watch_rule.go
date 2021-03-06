/*
 * API
 *
 * API
 *
 * API version: 1.0.0
 * Contact: you@your-company.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package stores

import (
	"time"
)

type WatchRule struct {

	Id string `json:"id,omitempty"`

	StoreId string `json:"store_id,omitempty"`

	ServiceId string `json:"service_id,omitempty"`

	ServiceName string `json:"service_name,omitempty"`

	Filters *interface{} `json:"filters,omitempty"`

	SyncScheduledAt time.Time `json:"sync_scheduled_at,omitempty"`

	Synced bool `json:"synced,omitempty"`

	Enabled bool `json:"enabled,omitempty"`
}
