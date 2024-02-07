package controller

import (
	"fmt"

	"io/fs"
	"net/http"
	"text/template"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"

	"go-gin/cmd/srv/middleware"
	"go-gin/internal/config"
	"go-gin/internal/tmpl"
	"go-gin/model"
	"go-gin/resource"
	"go-gin/service/singleton"
)

func ServerWeb(port uint) *http.Server {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	loadTemplates(r)

	if config.Debug {
		gin.SetMode(gin.DebugMode)
		pprof.Register(r, model.DefaultPprofRoutePath)
	}
	return &http.Server{
		Addr:    ":" + port,
		Handler: nil,
	}
}

// func routers(r *gin.Engine) {
// 	r := gin.Default()

// 	loadTemplates(r)

// 	serveStatic(r)
// 	// Serve uploaded files
// 	r.Static("/upload", config.Upload.Dir)

// 	routers(r)

// 	return r
// }

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
	r.Use(middleware.LoggingHandler())
	r.Use(middleware.RateLimiterHandler(singleton.Conf.RateLimit.Max))

	// Serve common pages, e.g. home, ping
	cp := commonPage{r: r}
	cp.serve()

	// Serve share pages, e.g. post
	sp := sharePage{r: r}
	sp.serve()

	//API
	api := r.Group("api/v1")
	{
		uap := &userAuthAPI{api)
		uap.serve()
	}

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

	r.NoRoute(pageNotFound)
	r.NoMethod(pageNotFound)
}

func pageNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound,
		&model.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: fmt.Sprintf("Resource not found: %s", c.Request.RequestURI),
		})
}
