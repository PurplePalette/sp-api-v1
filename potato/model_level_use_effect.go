/*
 * SweetPotato Server API
 *
 * Sonolusの基本APIを拡張する感じ。 ユーザー認証はFirebaseAuthorizationを通してやる。
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package potato

type LevelUseEffect struct {

	// If true (recommended), default resource specified by the engine will be used.
	UseDefault bool `json:"useDefault"`

	Item *Effect `json:"item,omitempty" validate:"omitempty"`
}
