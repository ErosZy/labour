package models

type KeyValuePair struct{
	Key   string
	Value string
}

type RequestParam struct {
	KeyValuePair
}

type RegexItem struct {
	RegexStr string
	Match []KeyValuePair
}

type XpathItem struct {
	DomStr  string
	Key     string
	Type    int
	AttrKey string
}