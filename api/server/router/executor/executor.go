package executor

import (
	"github.com/akutz/gofig"

	"github.com/emccode/libstorage/api/registry"
	"github.com/emccode/libstorage/api/server/httputils"
	apihttp "github.com/emccode/libstorage/api/types/http"
)

func init() {
	registry.RegisterRouter(&router{})
}

type router struct {
	routes []apihttp.Route
}

func (r *router) Name() string {
	return "executor-router"
}

func (r *router) Init(config gofig.Config) {
	r.initRoutes()
}

// Routes returns the available routes.
func (r *router) Routes() []apihttp.Route {
	return r.routes
}

func (r *router) initRoutes() {

	r.routes = []apihttp.Route{

		// GET
		httputils.NewGetRoute(
			"executors",
			"/executors",
			r.executors),

		// GET
		httputils.NewGetRoute(
			"executorInspect",
			"/executors/{executor}",
			r.executorInspect),

		// HEAD
		httputils.NewHeadRoute(
			"executorHead",
			"/executors/{executor}",
			r.executorHead),
	}
}
