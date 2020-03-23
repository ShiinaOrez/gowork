package models

type GetCompanyTags struct {
	TagList  struct{
		TagName string `json:"tag_name"`
		TagID   uint64 `json:"tag_list"`
	}  `json:"tag_list"`
}

type GetUserTags struct {
	TagList  struct{
		TagName string `json:"tag_name"`
		TagID   uint64 `json:"tag_list"`
	}  `json:"tag_list"`
}