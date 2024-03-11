package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/mehsanmgh/golang-project-web-api/api/middlewares"
	"github.com/mehsanmgh/golang-project-web-api/api/routers"
	"github.com/mehsanmgh/golang-project-web-api/api/validations"
	"github.com/mehsanmgh/golang-project-web-api/config"
)

func InitServer() {
	cfg := config.GetConfig()
	r := gin.New()
	// r := gin.Default()

	// use validator:

	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("mobile", validations.IranianMobileNumberValidator, true)
	}

	// middleware :
	r.Use(gin.Logger(), gin.Recovery(), middlewares.LimitByRequest() /*middlewares.OneTestMiddleware()*/)

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

	r.Run(fmt.Sprintf(":%s", cfg.Server.Port)) // ":5005"

}
