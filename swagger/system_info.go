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

type SystemInfo struct {

	// The storage of system.
	Storage []Storage `json:"storage,omitempty"`
}
