package gelbooru

type GelbooruSearch struct {
	Attributes GelbooruAttribute `json:"@attributes"`
	Posts      GelbooruPosts     `json:"post"`
}

type GelbooruAttribute struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Count  int `json:"count"`
}

type GelbooruPosts []GelbooruPost

type GelbooruPost struct {
	Identifier int    `json:"id"`
	Preview  string `json:"preview_url"`
}
