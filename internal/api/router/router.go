package router

import (
	"log"
	"os"
	"school/internal/api/handler"
	"school/internal/build"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "school/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router interface {
	SetupRouter()
	Run()
}
type router struct {
	engine      *gin.Engine
	address     string
	releaseMode bool

	versionHandler   handler.Version
	studentHandler   handler.StudentHandler
	menuHandler      handler.MenuHandler
	subscribeHandler handler.SubscribeHandler
}

func NewRouter() Router {
	release := !build.Get().IsDevelopment()
	if release {
		gin.SetMode(gin.ReleaseMode)
	}

	STATION_ADRESS := ":8080"
	if os.Getenv("SCHOOL_ROUTER_ADRESS") != "" {
		STATION_ADRESS = os.Getenv("SCHOOL_ROUTER_ADRESS")
	}

	return &router{
		engine:      gin.Default(),
		address:     STATION_ADRESS,
		releaseMode: release,

		versionHandler:   handler.NewVersion(),
		studentHandler:   handler.NewStudentHandler(),
		menuHandler:      handler.NewMenuHandler(),
		subscribeHandler: handler.NewSubscribeHandler(),
	}
}

func (r *router) SetupRouter() {
	gin.EnableJsonDecoderDisallowUnknownFields()

	r.engine.SetTrustedProxies(nil)
	r.engine.Use(cors.New(corsConfig()))

	r.addRoutes()
}

func (r *router) Run() {
	log.Printf("start DFI-UI-Backend Server on %s", r.address)
	if err := r.engine.Run(r.address); err != nil {
		log.Fatalf("could not start DFI-UI-Backend Server: %v", err)
	}
}

func (r *router) addVersionRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/version", r.versionHandler.GetVersion)
}

func (r *router) addStudentRoutes(routerGroup *gin.RouterGroup) {
	student := routerGroup.Group("student")
	{
		student.GET("", r.studentHandler.GetAll)
		student.POST("", r.studentHandler.Create)
		student.PUT("", r.studentHandler.Update)
		student.DELETE("", r.studentHandler.Delete)
	}
}

func (r *router) addSubscribeRoutes(routerGroup *gin.RouterGroup) {
	student := routerGroup.Group("subscribe")
	{
		student.POST("", r.subscribeHandler.BulkCreate)
	}
}

func (r *router) addMenuRoutes(routerGroup *gin.RouterGroup) {
	student := routerGroup.Group("menu")
	{
		student.GET("", r.menuHandler.GetAll)
		student.POST("", r.menuHandler.Create)
		student.PUT("", r.menuHandler.Update)
		student.DELETE("", r.menuHandler.Delete)
	}
}

func (r *router) addRoutes() {
	public := r.engine.Group("/api")
	{
		public.Use(handler.ErrorMiddleware())
	}

	protected := public.Group("")
	{
		//protected.Use(handler.AuthMiddleware())
	}

	r.addVersionRoutes(public)
	r.addStudentRoutes(protected)
	r.addMenuRoutes(protected)
	r.addSubscribeRoutes(protected)

	if !r.releaseMode {
		r.engine.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		log.Printf("swagger ui: http://localhost%s/api/swagger/index.html", r.address)
	}
}

func corsConfig() cors.Config {
	config := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}

	if allowed := os.Getenv("SCHOOL_ALLOWED_ORIGINS_LIST"); allowed != "" {
		config.AllowOrigins = strings.Split(allowed, ";")
	}

	return config
}
