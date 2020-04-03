package models

type Terms struct {
	TermID    uint32 `gorm:"primary_key;auto_increment" json:"term_id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	TermGroup uint32 `josn:"term_group"`
}
