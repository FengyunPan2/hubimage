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

type PingTarget struct {

	// The target address URL string.
	Endpoint string `json:"endpoint,omitempty"`

	// The target server username.
	Username string `json:"username,omitempty"`

	// The target server password.
	Password string `json:"password,omitempty"`
}