package http

import (
	"SimilaritySearch/searchengine"
	"fmt"
	"io/ioutil"
	ComHttp "net/http"

	jsoniter "github.com/json-iterator/go"
)

//Request Request
type Request struct {
	SetName  string    `json:"set_name"`
	Feat     []float32 `json:"feat"`
	FeatInfo string    `json:"feat_info"`
	TopN     int32     `json:"top_n"`
}

//Reload Reload
func Reload(rep ComHttp.ResponseWriter, req *ComHttp.Request) {
	body, _ := ioutil.ReadAll(req.Body)
	var Body Request
	if err := jsoniter.Unmarshal(body, &Body); err == nil {
		if Body.SetName != "" {
			searchengine.InitEngine(Body.SetName)
			rep.Write([]byte("Reload OK!"))

		} else {
			rep.Write([]byte("SetName = nil, Reload Failed!"))
		}
	} else {
		rep.Write([]byte("Reload Unmarshal Failed err = " + fmt.Sprintf("%s", err)))
	}
}

//Addid Addid
func Addid(rep ComHttp.ResponseWriter, req *ComHttp.Request) {
	body, _ := ioutil.ReadAll(req.Body)
	var Body Request
	if err := jsoniter.Unmarshal(body, &Body); err == nil {
		if Body.SetName != "" && len(Body.Feat) != 0 && Body.FeatInfo != "" {
			searchengine.Addid(Body.SetName, Body.Feat, Body.FeatInfo)
			rep.Write([]byte("Addid OK!"))
		} else {
			rep.Write([]byte("SetName | Feat | FeatInfo = nil, Addid Failed!"))
		}
	} else {
		rep.Write([]byte("Addid Unmarshal Failed err = " + fmt.Sprintf("%s", err)))
	}
}

//Search Search
func Search(rep ComHttp.ResponseWriter, req *ComHttp.Request) {
	body, _ := ioutil.ReadAll(req.Body)
	var Body Request
	if err := jsoniter.Unmarshal(body, &Body); err == nil {
		if Body.SetName != "" && len(Body.Feat) != 0 && Body.TopN > 0 {
			ret, _ := searchengine.Search(Body.SetName, Body.Feat, 1, Body.TopN)
			if len(ret) < 0 {
				rep.Write([]byte("Search Failed num = nil"))
			}

			rett, _ := jsoniter.Marshal(ret)
			rep.Write(rett)
		} else {
			rep.Write([]byte("SetName | Feat | TopN = nil, Search Failed!"))
		}
	} else {
		rep.Write([]byte("Search Unmarshal Failed err = " + fmt.Sprintf("%s", err)))
	}
}

//Unload Unload
func Unload(rep ComHttp.ResponseWriter, req *ComHttp.Request) {
	body, _ := ioutil.ReadAll(req.Body)
	var Body Request
	if err := jsoniter.Unmarshal(body, &Body); err == nil {
		if Body.SetName != "" {
			searchengine.DeleteFaissEngine(Body.SetName)
			rep.Write([]byte("Unload OK!"))
		} else {
			rep.Write([]byte("SetName = nil, Unload Failed!"))
		}
	} else {
		rep.Write([]byte("Unload Unmarshal Failed err = " + fmt.Sprintf("%s", err)))
	}
}

func GetStatus(rep ComHttp.ResponseWriter, req *ComHttp.Request) {
	ret, _ := searchengine.GetAllEngineStatus()
	rett, _ := jsoniter.Marshal(ret)
	rep.Write(rett)
}
