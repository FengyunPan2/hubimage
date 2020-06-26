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

type UserProfile struct {

	// The new email.
	Email string `json:"email,omitempty"`

	// The new realname.
	Realname string `json:"realname,omitempty"`

	// The new comment.
	Comment string `json:"comment,omitempty"`
}
