package internal

import (
	"context"

	"github.com/ssibrahimbas/ssi-we/pkg/db"
	"github.com/ssibrahimbas/ssi-we/pkg/helper"
	"github.com/ssibrahimbas/ssi-we/src/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repo struct {
	c   *mongo.Collection
	ctx context.Context
	db  *db.MongoDB
}

type RepoParams struct {
	Collection *mongo.Collection
	Ctx        context.Context
	DB         *db.MongoDB
}

func NewRepo(params *RepoParams) *Repo {
	return &Repo{
		c:   params.Collection,
		ctx: params.Ctx,
		db:  params.DB,
	}
}

func (r *Repo) Create(e *entity.Event) {
	_, err := r.c.InsertOne(r.ctx, e)
	helper.CheckErr(err)
}
