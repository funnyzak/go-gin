package srv

import (
	"html/template"
	"io/fs"
	"net/http"

	"go-gin/cmd/srv/controller"
	"go-gin/cmd/srv/middleware"
	"go-gin/internal/config"
	"go-gin/internal/tmpl"
	"go-gin/resource"
	"go-gin/service/singleton"

	"github.com/gin-gonic/gin"
)

func NewRoute(config *config.Config) *gin.Engine {
	r := gin.Default()

	loadTemplates(r)

	serveStatic(r)
	// Serve uploaded files
	r.Static("/upload", config.Upload.Dir)

	routers(r)

	return r
}

// Serve static files
func serveStatic(r *gin.Engine) {
	staticFs, err := fs.Sub(resource.StaticFS, "static")
	if err != nil {
		singleton.Log.Fatal().Err(err).Msg("Error parsing static files")
		panic(err)
	}
	r.StaticFS("/static", http.FS(staticFs))
}

// Load templates
func loadTemplates(r *gin.Engine) {
	new_tmpl := template.New("").Funcs(tmpl.FuncMap)
	var err error
	new_tmpl, err = new_tmpl.ParseFS(resource.TemplateFS, "template/**/*.html", "template/*.html")
	if err != nil {
		singleton.Log.Fatal().Err(err).Msg("Error parsing templates")
		panic(err)
	}
	r.SetHTMLTemplate(new_tmpl)
}

func routers(r *gin.Engine) {
	// Logging middleware
	r.Use(middleware.LoggingHandler())
	// Rate limit middleware
	r.Use(middleware.RateLimiterHandler(singleton.Config.RateLimit.Max))

	r.POST("/login", controller.Login)
	userGroup := r.Group("/v1/user")
	{
		// Set auth middleware
		userGroup.Use(middleware.AuthHanlder())
		userGroup.GET("/refresh", controller.Refresh)
		userGroup.POST("/upload/creation", controller.UploadCreation)
	}

	shareGroup := r.Group("/share")
	{
		shareGroup.GET("/creation/:share_num", controller.GetCreation)
	}

	r.GET("/", controller.Home)
	r.GET("/health", controller.HealthCheck)
	r.NoRoute(controller.PageNotFound)
}
