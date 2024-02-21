package controller

import (
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"

	"go-gin/internal/gconfig"
	"go-gin/internal/gogin"
	"go-gin/pkg/mygin"
	"go-gin/resource"
	"go-gin/service/singleton"
)

func ServerWeb(port uint) *http.Server {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	loadTemplates(r)

	if singleton.Conf.Debug {
		gin.SetMode(gin.DebugMode)
		pprof.Register(r, gconfig.DefaultPprofRoutePath)
	}

	r.Use(mygin.RecordPath)

	serveStatic(r)

	routers(r)

	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", port),
		ReadHeaderTimeout: time.Second * 5,
		Handler:           r,
	}
	return srv
}

// Serve static files
func serveStatic(r *gin.Engine) {
	staticFs, err := fs.Sub(resource.StaticFS, "static")
	if err != nil {
		singleton.Log.Fatal().Err(err).Msg("Error parsing static files")
		panic(err)
	}
	r.StaticFS("/static", http.FS(staticFs))

	// Serve uploaded files
	r.Static("/upload", singleton.Conf.Upload.Dir)
}

// Load templates
func loadTemplates(r *gin.Engine) {
	new_tmpl := template.New("").Funcs(gogin.FuncMap)
	var err error
	new_tmpl, err = new_tmpl.ParseFS(resource.TemplateFS, "template/**/*.html", "template/*.html")
	if err != nil {
		singleton.Log.Fatal().Err(err).Msg("Error parsing templates")
		panic(err)
	}
	r.SetHTMLTemplate(new_tmpl)
}

func routers(r *gin.Engine) {

	r.Use(gogin.LoggingHandler())
	r.Use(mygin.GenerateContextIdHandler())

	if singleton.Conf.EnableCORS {
		r.Use(mygin.CORSHandler())
	}

	r.Use(gogin.RateLimiterHandler(singleton.Conf.RateLimit.Max))

	// Serve common pages, e.g. home, ping
	cp := commonPage{r: r}
	cp.serve()

	// Serve guest pages, e.g. register, login
	gp := guestPage{r: r}
	gp.serve()

	// Server show pages, e.g. post
	sp := showPage{r: r}
	sp.serve()

	// User pages, e.g. profile, setting
	up := userPage{r: r}
	up.serve()

	// Serve API
	api := r.Group("api")
	{
		ua := &userAPI{r: api}
		ua.serve()
	}

	page404 := func(c *gin.Context) {
		gogin.ShowErrorPage(c, mygin.ErrInfo{
			Code:  http.StatusNotFound,
			Title: "Page not found",
			Msg:   "The page you are looking for is not found",
			Link:  singleton.Conf.Site.BaseURL,
			Btn:   "Back to home",
		}, true)
	}
	r.NoRoute(page404)
	r.NoMethod(page404)
}
