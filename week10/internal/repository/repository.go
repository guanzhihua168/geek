package repository

import (
	"github.com/jinzhu/gorm"
	"net/http"
	"parent-api-go/pkg/context"
	"parent-api-go/pkg/linkerd"
)

type repository struct {
	Ctx    *context.AppContext
	Remote *linkerd.Linkerd
	Db     *gorm.DB
}

type RemoteHandle func(ctx *context.AppContext) *linkerd.Linkerd
type DbHandle func(ctx *context.AppContext) *gorm.DB
type RepoOption func(*repository)

func linkerdOptions(c repository) []linkerd.LinkerdOptions {
	var options []linkerd.LinkerdOptions
	if c.Ctx != nil {
		options = append(options, linkerd.WithLogger(c.Ctx.Log))
		options = append(options, linkerd.WithContext(func(h *http.Request) {
			h.WithContext(c.Ctx.Request.Context())
		}))
	}

	return options
}
