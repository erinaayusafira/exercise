package main

import (
	"exercise/controllers"
	"github.com/eaciit/knot/knot.v1"
	"net/http"
	"os"
)

var (
	appViewPath = (func(dir string, _ error) string { return dir + "/views/" }(os.Getwd()))
)

func main(){
	app := knot.NewApp("")
	app.ViewsPath = appViewPath

	app.Register(new(controllers.KnotController))
	app.Register(new(controllers.DatetimeController))
	app.Register(new(controllers.FetchgithubController))
	app.Register(new(controllers.OrmdboxController))
	app.Register(new(controllers.FileuploadController))
	app.Register(new(controllers.PerlinController))
	app.Register(new(controllers.Perlin2Controller))
	app.Register(new(controllers.GoroutineController))

	app.LayoutTemplate = "template.html"
	app.DefaultOutputType = knot.OutputTemplate

	knot.RegisterApp(app)
	otherRoutes := map[string]knot.FnContent{
		"/": func(r *knot.WebContext) interface{} {
			http.Redirect(r.Writer, r.Request, "/fetchgithub/fetchgithubajax", http.StatusTemporaryRedirect)
			return true
	    },
	}
	knot.StartAppWithFn(app, "localhost:8999", otherRoutes)
}