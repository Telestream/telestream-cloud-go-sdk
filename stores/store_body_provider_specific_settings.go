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

type StoreBodyProviderSpecificSettings struct {

	Host string `json:"host,omitempty"`

	Port string `json:"port,omitempty"`

	CacheControl string `json:"cache_control,omitempty"`

	Metadata string `json:"metadata,omitempty"`

	FaspPort string `json:"fasp_port,omitempty"`

	SshPort string `json:"ssh_port,omitempty"`

	MaxRate string `json:"max_rate,omitempty"`

	MinRate string `json:"min_rate,omitempty"`

	Policy string `json:"policy,omitempty"`
}
