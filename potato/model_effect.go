/*
 * SweetPotato Server API
 *
 * Sonolusの基本APIを拡張する感じ。 ユーザー認証はFirebaseAuthorizationを通してやる。
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package potato

// Effect - An effect provides audio elements to levels / It defines audio effect used for specific level https://github.com/NonSpicyBurrito/sonolus-wiki/wiki/Effect
type Effect struct {

	// english and number only name for searching
	Name string `json:"name,omitempty" validate:"omitempty,alphanum,min=1,max=50"`

	// Reserved for future update. current default is 1.
	Version int32 `json:"version,omitempty" validate:"omitempty,gte=1,lte=1"`

	// base title of this content
	Title string `json:"title,omitempty" validate:"omitempty,alphanumunicode,min=1,max=100"`

	// something footer(ex. featuring xyz) for this content
	Subtitle string `json:"subtitle,omitempty" validate:"omitempty,alphanumunicode,min=1,max=100"`

	// author of this content
	Author string `json:"author,omitempty" validate:"omitempty,alphanumunicode,min=1,max=50"`

	Thumbnail SonolusResourceLocator `json:"thumbnail,omitempty" validate:"omitempty"`

	Data SonolusResourceLocator `json:"data,omitempty" validate:"omitempty"`

	// 独自要素: データを作成したエポックミリ秒(ソート用)
	CreatedTime int32 `json:"createdTime,omitempty" validate:"omitempty,gte=1"`

	// 独自要素: データを更新したエポックミリ秒(ソート用)
	UpdatedTime int32 `json:"updatedTime,omitempty" validate:"omitempty,gte=1"`

	// 独自要素: 譜面作成者のユーザーID
	UserId string `json:"userId,omitempty" validate:"omitempty,alphanum,min=1,max=50"`

	// 独自要素: サイト内および譜面情報欄に表示される説明文
	Description string `json:"description,omitempty" validate:"omitempty,alphanumunicode,min=1,max=3000"`
}
