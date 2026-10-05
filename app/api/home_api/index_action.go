package home_api

import (
	"github.com/gin-gonic/gin"

	"github.com/emo-lang/website-cn/app/views/features"
	"github.com/emo-lang/website-cn/app/views/home"
	"github.com/emo-lang/website-cn/app/views/quickstart"
	"github.com/emo-lang/website-cn/app/views/tour"
	"github.com/daqing/airway/lib/render"
)

func IndexAction(c *gin.Context) {
	render.HTML(c, home.Index())
}

func FeaturesAction(c *gin.Context) {
	render.HTML(c, features.Index())
}

func TourAction(c *gin.Context) {
	render.HTML(c, tour.Index())
}

func QuickstartAction(c *gin.Context) {
	render.HTML(c, quickstart.Index())
}
