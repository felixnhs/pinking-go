package db

import (
	"github.com/pocketbase/pocketbase/core"
)

var _ core.RecordProxy = (*Image)(nil)

type Image struct {
	core.BaseRecordProxy
}

const Image_Raw = "raw"
const Image_Active = "active"
const Image_CreatedBy = "createdby"
const Image_Created = "created"
const Image_UpdatedBy = "updatedby"
const Image_Updated = "updated"
const Image_Order = "order"

func (u *Image) GetRawFileName() string {
	return u.GetString(Image_Raw)
}

func (u *Image) GetOrder() int {
	return u.GetInt(Image_Order)
}

func (u *Image) SetOrder(order int) {
	u.Set(Image_Order, order)
}

func (u *Image) GetActive() bool {
	return u.GetBool(Image_Active)
}

func (u *Image) SetActive(active bool) {
	u.Set(Image_Active, active)
}

func (p *Image) GetCreatedBy() string {
	return p.GetString(Image_CreatedBy)
}

func (p *Image) SetCreatedBy(id string) {
	p.Set(Image_CreatedBy, id)
}

func (p *Image) GetUpdatedBy() string {
	return p.GetString(Image_UpdatedBy)
}

func (p *Image) SetUpdatedBy(id string) {
	p.Set(Image_UpdatedBy, id)
}
