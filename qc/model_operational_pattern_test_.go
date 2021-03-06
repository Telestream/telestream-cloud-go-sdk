/*
 * Qc API
 *
 * Qc API
 *
 * API version: 3.0.0
 * Contact: cloudsupport@telestream.net
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package qc

type OperationalPatternTest struct {
	Op1a bool `json:"op1a,omitempty"`
	Op2a bool `json:"op2a,omitempty"`
	Op3a bool `json:"op3a,omitempty"`
	Op1b bool `json:"op1b,omitempty"`
	Op2b bool `json:"op2b,omitempty"`
	Op3b bool `json:"op3b,omitempty"`
	Op1c bool `json:"op1c,omitempty"`
	Op2c bool `json:"op2c,omitempty"`
	Op3c bool `json:"op3c,omitempty"`
	ExternalEssence OptionalFlag `json:"external_essence,omitempty"`
	NonStreamable OptionalFlag `json:"non_streamable,omitempty"`
	MultiTrack OptionalFlag `json:"multi_track,omitempty"`
	OpAtom bool `json:"op_atom,omitempty"`
	MultiSource OptionalFlag `json:"multi_source,omitempty"`
	MultiEssence OptionalFlag `json:"multi_essence,omitempty"`
	RejectOnError bool `json:"reject_on_error,omitempty"`
	Checked bool `json:"checked,omitempty"`
}
