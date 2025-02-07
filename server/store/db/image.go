package db

import (
	"github.com/pocketbase/pocketbase/core"
)

var _ core.RecordProxy = (*Image)(nil)

type Image struct {
	core.BaseRecordProxy
}

func (u *Image) GetRawFileName() string {
	return u.GetString("raw")
}

func (u *Image) GetOrder() int {
	return u.GetInt("order")
}

func (u *Image) SetOrder(order int) {
	u.Set("order", order)
}

func (u *Image) GetActive() bool {
	return u.GetBool("active")
}

func (u *Image) SetActive(active bool) {
	u.Set("active", active)
}

func (p *Image) GetCreatedBy() string {
	return p.GetString("createdby")
}

func (p *Image) SetCreatedBy(id string) {
	p.Set("createdby", id)
}

func (p *Image) GetUpdatedBy() string {
	return p.GetString("updatedby")
}

func (p *Image) SetUpdatedBy(id string) {
	p.Set("updatedby", id)
}
