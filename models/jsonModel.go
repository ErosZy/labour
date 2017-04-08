package models

type JSONItem struct{
	KeyValuePair
	TypeStr string
}

type JSONParam struct {
	Type int
	JSONs []JSONItem
}