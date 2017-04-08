package models

type PrefixXpathItem struct {
	Prefix string
	Arr    []*XpathItem
}

type PageTextItem struct {
	Regex []*RegexItem
	Xpath []*PrefixXpathItem
}

type PageItemModel struct {
	Route   string
	MainKey string
	Text    *PageTextItem
	Json    *JSONParam
}
