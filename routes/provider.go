package routes

import (
	"github.com/totoval/framework/request"
	"github.com/totoval/framework/route"
	"totoval/routes/versions"
)

func Register(router *request.Engine) {
	defer route.Bind(router)

	versions.NewV1(router)
}
