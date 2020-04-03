package helper

type Pagination struct {
	Row   uint `form:"p"`
	Limit uint `form:"l"`
}

func (self *Pagination) GetRow() uint {
	if self.Row == 0 {
		return 1
	}
	if self.Row > 50 {
		return 50
	}
	return self.Row
}

func (self *Pagination) GetLimit() uint {
	if self.Limit == 0 || self.Limit > 20 {
		return 20
	}
	return self.Limit
}

func (self *Pagination) GetOffset() uint {
	row := self.GetRow()
	limit := self.GetLimit()
	offset := (row - 1) * limit
	return offset
}
