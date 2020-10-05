package controllers

import (
	"github.com/eaciit/knot/knot.v1"
)

type FileuploadController struct {
}

func (w *FileuploadController) FileuploadAjax(r *knot.WebContext) interface{} {
	r.Config.OutputType = knot.OutputTemplate
	r.Config.ViewName = "fileupload/fileuploadajax.html"

	return nil
}
func (w *FileuploadController) FileuploadKnockout(r *knot.WebContext) interface{} {
	r.Config.OutputType = knot.OutputTemplate
	r.Config.ViewName = "fileupload/fileuploadknockout.html"

	return nil
}

func (w *FileuploadController) FileuploadJson(r *knot.WebContext) interface{} {
	r.Config.OutputType = knot.OutputJson
	return nil
}