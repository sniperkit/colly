package router

import (
	"github.com/feedlabs/feedify"
	"github.com/sniperkit/colly/plugins/data/aggregate/web/service/service/store/v1/controller"
)

func InitTokenRouters() {
	feedify.Router("/v1/org/:orgId:string/token", &controller.TokenController{}, "get:GetOrgList;post:PostOrg")
	feedify.Router("/v1/org/:orgId:string/token/:tokenId:string", &controller.TokenController{}, "get:GetOrg;delete:DeleteOrg")

	feedify.Router("/v1/admin/:adminId:string/token", &controller.TokenController{}, "get:GetAdminList;post:PostAdmin")
	feedify.Router("/v1/admin/:adminId:string/token/:tokenId:string", &controller.TokenController{}, "get:GetAdmin;delete:DeleteAdmin")
}
