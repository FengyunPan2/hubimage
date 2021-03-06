/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * OpenAPI spec version: 0.3.0
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type RepPolicyUpdate struct {

	// The target ID.
	TargetId int64 `json:"target_id,omitempty"`

	// The policy name.
	Name string `json:"name,omitempty"`

	// The policy's enabled status.
	Enabled int32 `json:"enabled,omitempty"`

	// The description of the policy.
	Description string `json:"description,omitempty"`

	// The cron string for schedule job.
	CronStr string `json:"cron_str,omitempty"`
}
