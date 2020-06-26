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

type JobStatus struct {

	// The job ID.
	Id int64 `json:"id,omitempty"`

	// The status of the job.
	Status string `json:"status,omitempty"`

	// The repository handled by the job.
	Repository string `json:"repository,omitempty"`

	// The ID of the policy that triggered this job.
	PolicyId int64 `json:"policy_id,omitempty"`

	// The operation of the job.
	Operation string `json:"operation,omitempty"`

	// The repository's used tag list.
	Tags []Tags `json:"tags,omitempty"`

	// The creation time of the job.
	CreationTime string `json:"creation_time,omitempty"`

	// The update time of the job.
	UpdateTime string `json:"update_time,omitempty"`
}
