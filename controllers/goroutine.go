package controllers

import (
	"fmt"
	"github.com/eaciit/knot/knot.v1"
	"github.com/eaciit/toolkit"
	"time"
)

type GoroutineController struct {

}

func(w *GoroutineController) GoroutineAjax(r *knot.WebContext)interface{}{
	r.Config.OutputType = knot.OutputTemplate
	r.Config.ViewName = "goroutine/goroutine.html"

	return nil
}
func(w *GoroutineController) GoroutineKnockout(r *knot.WebContext)interface{}{
	r.Config.OutputType = knot.OutputTemplate
	r.Config.ViewName = "datetime/datetimeknockout.html"

	return nil
}
func(w *GoroutineController) GoroutineJson(r *knot.WebContext) interface{}{
	r.Config.OutputType = knot.OutputJson

	return time.Now().Format("2016-01-02T15:02Z")
}
func sample(s string){
	for i := 0; i < 5; i++ {
		time.Sleep(100 *time.Millisecond)
		fmt.Println(s)
	}
}
func (w *GoroutineController) Result(r *knot.WebContext) interface{}{
	r.Config.OutputType = knot.OutputJson

	payload := struct {
		x string
		y string
	}{}
	err := r.GetPayload(&payload)
	if err != nil{
		toolkit.Println(err.Error())
	}
	go sample(payload.x)
	sample(payload.y)

	return nil
}