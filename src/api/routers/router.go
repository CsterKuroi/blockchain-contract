// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"unicontract/src/api/controllers"
	"unicontract/src/api/filters"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true}))

	auth_verify := beego.AppConfig.DefaultBool("auth_verify", false)
	auth_verify_rate_limit := beego.AppConfig.DefaultBool("auth_verify_rate_limit", false)

	beego.InsertFilter("/*", beego.BeforeRouter, filters.MonitorFilter, false)
	// filter shouldn`t use the api log!

	//TODO temp deal
	_ = auth_verify
	_ = auth_verify_rate_limit
	//if auth_verify {
	//	// auth request app_id and app_key, return token
	//	beego.InsertFilter("/*", beego.BeforeRouter, filters.APIBasicFilter, true)
	//	if auth_verify_rate_limit {
	//		beego.InsertFilter("/*", beego.BeforeRouter, filters.APIRateLimitFilter, true)
	//	}
	//
	//}

	//todo 1. auth_verify=true
	//false will ignore all the filters
	//todo 2. basic http method, content-type and others verify! auth_verify_http=true
	//default only verify the content-type
	//todo 3. basic filter timestamp filter 请求时间戳过滤功能, auth_verify_timestamp=true
	//todo 4. filter parameters, verify the input parameters if all in api.ALLOW_REQUEST_PARAMETERS_ALL, auth_verify_parameters=false
	//sort fields must in api.ALLOW_REQUEST_PARAMETERS_MODEL
	//todo 5. verify the basic parameter sign(except sign, encrypt-> sign只针对 parameters进行加密， 请求参数字典序进行hash) auth_verify_sign=false
	//maybe cost time
	//todo 6. verify the token parameter from redis, auth_verify_token=false
	//todo 7. rate Limit verify, depend the step 6! auth_verify_rate_limit=false

	ns := beego.NewNamespace("/v1/unicontract",
		beego.NSNamespace("/contract",
			beego.NSRouter("/create", &controllers.ContractController{}, "post:Create"),
			beego.NSRouter("/queryPublishContract", &controllers.ContractController{}, "post:QueryPublishContract"),
			beego.NSRouter("/queryContractContent", &controllers.ContractController{}, "post:QueryContractContent"),
			beego.NSRouter("/queryExecuteContractContent", &controllers.ContractController{}, "post:QueryExecuteContractContent"),
			beego.NSRouter("/query", &controllers.ContractController{}, "post:Query"),
			beego.NSRouter("/queryAll", &controllers.ContractController{}, "post:QueryAll"),
			beego.NSRouter("/queryLog", &controllers.ContractController{}, "post:QueryLog"),
			beego.NSRouter("/pressTest", &controllers.ContractController{}, "post:PressTest"),
			beego.NSRouter("/terminate", &controllers.ContractController{}, "post:Terminate"),
			//demo使用---------------------------------------------------------------------------------------------------
			beego.NSRouter("/queryOutput", &controllers.ContractController{}, "post:QueryOutput"),
			beego.NSRouter("/queryOutputNum", &controllers.ContractController{}, "post:QueryOutputNum"),
			beego.NSRouter("/queryOutputDuration", &controllers.ContractController{}, "post:QueryOutputDuration"),
			beego.NSRouter("/queryAccountBalance", &controllers.ContractController{}, "post:QueryAccountBalance"),
			beego.NSRouter("/queryAmmeterBalance", &controllers.ContractController{}, "post:QueryAmmeterBalance"),
			beego.NSRouter("/queryRecords", &controllers.ContractController{}, "post:QueryRecords"),
			//demo使用---------------------------------------------------------------------------------------------------

			// 部署检测使用
			beego.NSRouter("/detect", &controllers.ContractController{}, "get:DetectConfig"),
		),
	)
	beego.AddNamespace(ns)
}
