/*
 * SweetPotato Server API
 *
 * Sonolusの基本APIを拡張する感じ。 ユーザー認証はFirebaseAuthorizationを通してやる。
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package potato

// GetBackgroundResponse - Response struct of getBackground
type GetBackgroundResponse struct {
	Item Background `json:"item"`

	Description string `json:"description"`

	Recommended []Background `json:"recommended"`
}
