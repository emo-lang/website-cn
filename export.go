package main

import (
	"github.com/daqing/airway/cmd"
	"github.com/daqing/airway/lib/static"

	"github.com/emo-lang/website-cn/app/views/features"
	"github.com/emo-lang/website-cn/app/views/home"
	"github.com/emo-lang/website-cn/app/views/quickstart"
	"github.com/emo-lang/website-cn/app/views/tour"
)

// Static export: `airway static:build` renders the marketing pages into
// dist/ together with the committed frontend bundle (see
// docs/static-export.md). Register more pages here, or run
// `generate scaffold` — each resource adds its own export_<name>.go.
func init() {
	cmd.SetStaticPages(
		static.Page{Slug: "/", Component: home.Index()},
		static.Page{Slug: "/features", Component: features.Index()},
		static.Page{Slug: "/tour", Component: tour.Index()},
		static.Page{Slug: "/quickstart", Component: quickstart.Index()},
	)
}
