package controllers

import (
	"github.com/eaciit/knot/knot.v1"
)
type OrmdboxController struct {
}

func (w *OrmdboxController) OrmdboxAjax(r *knot.WebContext) interface{} {
	r.Config.OutputType = knot.OutputTemplate
	r.Config.ViewName = "ormdbox/ormdboxajax.html"

	return nil
}
