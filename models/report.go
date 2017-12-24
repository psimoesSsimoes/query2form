package models

type Report struct {
	id        int    `form:"id"`
	expected  int    `form:"expected"`
	measured  int    `form:"measured"`
	eventcode int    `form:"eventcode"`
	notes     string `form:"notes"`
	problem   string `form:"problem"`
	submit    string `form:"submit"`
}
