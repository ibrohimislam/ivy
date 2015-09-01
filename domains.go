package ivy

type Data interface{}
type MetaData struct {
	Label string
	Type string
}

type Entity struct {
	dataSet []Data
	metaData []MetaData
}