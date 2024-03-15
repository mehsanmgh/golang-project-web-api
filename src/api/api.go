package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/mehsanmgh/golang-project-web-api/api/middlewares"
	"github.com/mehsanmgh/golang-project-web-api/api/routers"
	validation "github.com/mehsanmgh/golang-project-web-api/api/validations"
	"github.com/mehsanmgh/golang-project-web-api/config"
)

func InitServer(cfg *config.Config) {
	
	r := gin.New()
	// r := gin.Default()

	RegisterValidator()

	// middleware :
	r.Use(middlewares.Cors(cfg))
	r.Use(gin.Logger(), gin.Recovery(), middlewares.LimitByRequest() /*middlewares.OneTestMiddleware()*/)

	
	RegisterRoutes(r)

	r.Run(fmt.Sprintf(":%s", cfg.Server.Port)) // ":5005"

}

// use validator:
func RegisterValidator() {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("mobile", validation.IranianMobileNumberValidator, true)
		val.RegisterValidation("password", validation.PasswordValidator, true)
	}
}

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	v1 := api.Group("/v1")

	{
		health := v1.Group("/health")
		test_router := v1.Group("/test")
		routers.Health(health)
		routers.TestRouter(test_router)
	}

	// server := &http.Server{
	// 	Handler: r,
	// 	Addr: fmt.Sprintf(":%s", cfg.Sever.Port),
	// 	ReadTimeout: time.Second * 10,
	// }
	// server.ListenAndServe()

	v2 := api.Group("/v2")
	{
		health := v2.Group("/health")
		routers.Health(health)
	}
}
