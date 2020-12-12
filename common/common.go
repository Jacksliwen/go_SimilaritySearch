package common

//FeatureSet FeatureSet
type FeatureSet struct {
	SetName        string
	AllFeature     [][]float32
	IDInfo         []string
	MapFeatureInfo map[int32]string
}

//EngineInfo EngineInfo
type EngineInfo struct {
	SetName     string `json:"set_name"`
	FeatureSize int    `json:"featurn_sizze"`
	FeatureNum  int    `json:"feature_num"`
}

//SearchRet SearchRet
type SearchRet struct {
	SetName  string  `json:"set_name"`
	ID       string  `json:"id"`
	Distance float32 `json:"distance"`
}
