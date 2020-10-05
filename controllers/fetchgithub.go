package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/eaciit/knot/knot.v1"
	"github.com/eaciit/toolkit"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type FetchgithubController struct {
}
type Result struct {
	Login 		string			`json:"login"`
	Name    	string    		`json:"name"`
	Company 	string			`json:"company"`
	CreateAt	time.Time		`json:"create_at"`
	UpdateAt	time.Time		`json:"update_at"`
}
func (w *FetchgithubController) FetchgithubAjax(r *knot.WebContext) interface{} {
	r.Config.OutputType = knot.OutputTemplate
	r.Config.ViewName = "fetchgithub/fetchgithubajax.html"

	return nil
}
func (w *FetchgithubController) FetchgithubKnockout(r *knot.WebContext) interface{} {
	r.Config.OutputType = knot.OutputTemplate
	r.Config.ViewName = "fetchgithub/fetchgithubknockout.html"

	return nil
}

func (w *FetchgithubController) FetchgithubJson(r *knot.WebContext) interface{} {
	// set output type ke knot.OutputJson
	r.Config.OutputType = knot.OutputJson

	return time.Now().Format("2006-01-02T15:04Z")
}
func (w *FetchgithubController) Result(r *knot.WebContext) interface{} {
	// set output type ke knot.OutputJson
	r.Config.OutputType = knot.OutputJson
	payload := struct {
		Login   string
	}{}
	e := r.GetPayload(&payload)
	if e != nil {
		toolkit.Println(e.Error())
	}
	baseUrl := "https://api.github.com/users/"
	credentials := "?client_id=Iv1.b4f03e0ed50647ff&client_secrets=9f757a841dad62c18b7ff925db8d23be592e3081"
	toolkit.Sprintf("user %s : " , string(payload.Login) )
	url := baseUrl + payload.Login +credentials
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Result
	json.Unmarshal(responseData, &responseObject)

	var result = Result{
		Login: responseObject.Login,
		Name: responseObject.Name,
		Company: responseObject.Company,
		CreateAt: responseObject.CreateAt,
		UpdateAt: responseObject.UpdateAt,
	}
	toolkit.Println(result.Name)
	sample := []toolkit.M{
		{"login" : result.Login},
		{"name" : result.Name},
		{"company" : result.Company},
		{"created_at" : result.CreateAt},
		{"update_at" : result.UpdateAt},
	}
	fmt.Println(sample)
	return sample
}