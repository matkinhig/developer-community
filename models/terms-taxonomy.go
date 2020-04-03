package models

type TermsTaxonomy struct {
	TermTaxonomyID uint32 `gorm:"primary_key;auto_increment" json:"term_taxonomy_id"`
	TermID         uint32 `json:"term_id"`
	Taxonomy       string `json:"taxonomy"`
	Description    string `json:"description"`
	Parent         uint32 `json:"parent"`
	Count          uint32 `json:"count"`
}

func (t *TermsTaxonomy) Validate() error {
	var err error
	return err
}
