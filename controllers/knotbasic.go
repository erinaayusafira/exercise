package controllers

import (
	"github.com/eaciit/knot/knot.v1"
	"github.com/eaciit/toolkit"
	"time"
)
type KnotController struct{
	
}

func (w *KnotController) KnotAjax(r *knot.WebContext) interface{}{
	r.Config.OutputType = knot.OutputTemplate
	r.Config.ViewName = "knotbasic/knotajax.html"

	return nil
}
func(w *KnotController) KnotKnockout(r *knot.WebContext) interface{}{
	r.Config.OutputType = knot.OutputTemplate
	r.Config.ViewName = "knotbasic/knotknockout.html"

	return nil
}

func (w *KnotController) KnotJson(r *knot.WebContext) interface{}{
	r.Config.OutputType = knot.OutputJson
	sample := []toolkit.M{
		{"date": time.Now().Format("2006-01-02T15:04:05Z")},
	}
	return sample
}