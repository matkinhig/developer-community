package models

import "time"

type Posts struct {
	ID                  uint32    `gorm:"primary_key;auto_increment" json:"id"`
	PostAuthor          string    `json:"post_author"`
	PostDate            time.Time `json:"post_date"`
	PostDateGmt         time.Time `json:"post_date_gmt"`
	PostContent         string    `json:"post_content"`
	PostTitle           string    `json:"post_title"`
	PostExcerpt         string    `json:"post_excerpt"`
	PostStatus          string    `json:"post_status"`
	CommentStatus       string    `json:"comment_status"`
	PingStatus          string    `josn:"ping_status"`
	PostPassword        string    `json:"post_password"`
	PostName            string    `json:"post_name"`
	ToPing              string    `json:"to_ping"`
	Pinged              string    `json:"pinged"`
	PostModify          time.Time `json:"post_modified"`
	PostModifyGmt       time.Time `json:"post_modified_gmt"`
	PostContentFiltered string    `json:"post_content_filtered"`
	PostParent          uint32    `json:"post_parent"`
	GuID                string    `json:"guid"`
	MenuOrder           uint32    `json:"menu_order"`
	PostType            string    `json:"post_type"`
	PostMineType        string    `json:"post_mime_type"`
	CommentCount        uint32    `json:"comment_count"`
}

func (p *Posts) Validate() error {
	var err error
	return err
}
