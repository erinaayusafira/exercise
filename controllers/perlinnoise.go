package controllers

import (
	"github.com/eaciit/knot/knot.v1"
	"github.com/eaciit/toolkit"
	"strconv"
	"sync"
)

type PerlinController struct {
}

func (w *PerlinController) PerlinAjax(r *knot.WebContext) interface{} {
	r.Config.OutputType = knot.OutputTemplate
	r.Config.ViewName = "perlinnoise/perlinajax.html"
	return nil
}

func hasil(x, y float64, data *[]float64, wg *sync.WaitGroup) {
	hasil := x + y
	*data = append(*data, hasil)
	wg.Done()
}

func (w *PerlinController) PerlinJson(r *knot.WebContext) interface{} {
	r.Config.OutputType = knot.OutputJson

	payload := struct {
		X   string
		Y  string
	}{}
	e := r.GetPayload(&payload)
	if e != nil {
		toolkit.Println(e.Error())
	}
	toolkit.Println(payload.Y)
	toolkit.Println(payload.X)
	result := []float64{}

	startPoint, _ := strconv.ParseFloat(payload.X, 64)
	endPoint, _ := strconv.ParseFloat(payload.Y, 64)

	total := int(startPoint * endPoint)

	var wg sync.WaitGroup

	wg.Add(total)
	for i := 0.0; i < startPoint; i++ {
		for j := 0.0; j < endPoint; j++ {
			go hasil(i, j, &result, &wg)
		}
	}

	wg.Wait()
	return result
}