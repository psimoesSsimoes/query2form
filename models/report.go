package models

type Report struct {
	id        int    `form:"id"`
	expected  int    `form:"expected"`
	measured  int    `form:"measured"`
	eventcode int    `form:"eventcode"`
	notes     string `form:"notes"`
	problem   string `form:"problem"`
	submit    string `form:"submit"`
	timestamp string `form:"timestamp"`
}

func (r *Report) FillReportInt(expected, measured, eventcode, id int, timestamp string) {
	r.expected = expected
	r.measured = measured
	r.eventcode = eventcode
	r.id = id
	r.timestamp = timestamp

}

func (r *Report) FillReportString(notes, problem string) {
	r.notes = notes
	r.problem = problem
}
