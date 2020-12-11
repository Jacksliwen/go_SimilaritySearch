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

//RepInfo RepInfo
type RepInfo struct {
	Index    int64   `json:"index"`
	Distance float32 `json:"distance"`
}

//Respones Respones
type Respones struct {
	SetName string    `json:"set_name"`
	Infos   []RepInfo `json:"resusts"`
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
		if Body.SetName != "" && Body.Feat != nil && Body.FeatInfo != "" {
			searchengine.LoadData(Body.SetName, (*float32)(&Body.Feat[0]), 1)
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
		if Body.SetName != "" && Body.Feat != nil && Body.TopN > 0 {
			var D []float32
			var I []int64
			num := searchengine.Search(Body.SetName, (*float32)(&Body.Feat[0]), 1, Body.TopN, (*int64)(&I[0]), (*float32)(&D[0]))
			if num < 0 {
				rep.Write([]byte("Search Failed num = " + fmt.Sprintf("%d", num)))
			}
			infoss := []RepInfo{}
			for i := 0; i < num; i++ {
				info := &RepInfo{
					Index:    I[i],
					Distance: D[i],
				}
				infoss = append(infoss, *info)
			}

			ret := &Respones{
				SetName: Body.SetName,
				Infos:   infoss,
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
