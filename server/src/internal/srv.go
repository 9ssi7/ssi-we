package internal

import (
	"time"

	"github.com/ssibrahimbas/ssi-we/src/dto"
	"github.com/ssibrahimbas/ssi-we/src/entity"
)

type Srv struct {
	r *Repo
}

type SrvParams struct {
	Repo *Repo
}

func NewSrv(p *SrvParams) *Srv {
	return &Srv{
		r: p.Repo,
	}
}

func (s *Srv) Create(d *dto.EventCreateRequest) {
	e := &entity.Event{
		UUID:      "",
		Page:      d.Page,
		Element:   d.Element,
		Type:      d.Type,
		Params:    d.Params,
		User:      d.User,
		IP:        d.IP,
		CreatedAt: time.Now(),
	}
	s.r.Create(e)
}
