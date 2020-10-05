package controllers

import (
	"github.com/eaciit/knot/knot.v1"
	"github.com/eaciit/toolkit"
	"time"
)

type DatetimeController struct{
}

func(w *DatetimeController) DatetimeAjax(r *knot.WebContext) interface{}{
	r.Config.OutputType = knot.OutputTemplate
	r.Config.ViewName = "datetime/datetimeajax.html"

	return nil
}

func(w *DatetimeController) DatetimeKnockout(r *knot.WebContext) interface{}{
	r.Config.OutputType = knot.OutputTemplate
	r.Config.ViewName = "datetime/datetimeknockout.html"

	return nil
}

func(w *DatetimeController) DatetimeJson(r *knot.WebContext) interface{}{
	r.Config.OutputType = knot.OutputJson

	return time.Now().Format("2006-01-02T15:04:05Z")
}

func(w *DatetimeController) DatetimeResult(r *knot.WebContext) interface{}{
	r.Config.OutputType = knot.OutputJson

	payload := struct {
		Time string`json:"time"`
	}{}

	e := r.GetPayload(&payload)
	if e != nil{
		toolkit.Println(e.Error())
	}
	toolkit.Println(payload.Time)

	result := []toolkit.M{
		{"Now" : time.Now().Format("2006-01-02T15:04:05Z")},
		{"After" : time.Now().Add(12*time.Hour).Format("2006-01-02T15:04:05Z")},
	}
	return result
}
