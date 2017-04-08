package models

type SchedulerTextItem struct {
	Regex []*RegexItem
	Xpath []*XpathItem
}

type SchedulerItemModel struct {
	Route string
	Text  *SchedulerTextItem
	Json  *JSONParam
}