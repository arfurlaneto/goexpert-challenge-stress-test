package internal

import (
	"fmt"
	"sync"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

type testReport struct {
	URL             string
	Runners         int64
	Duration        *time.Duration
	RequestCount    int64
	StatusCodeCount map[int]int64
	ErrorCount      int64
	mutex           sync.Mutex
}

func NewTestReport(url string, runners int64) *testReport {
	report := &testReport{}
	report.URL = url
	report.Runners = runners
	report.StatusCodeCount = map[int]int64{}
	return report
}

func (r *testReport) AddRequest(statusCode int) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.RequestCount++
	currentCount, ok := r.StatusCodeCount[statusCode]
	if !ok {
		r.StatusCodeCount[statusCode] = 1
	} else {
		r.StatusCodeCount[statusCode] = currentCount + 1
	}
}

func (r *testReport) AddError() {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.ErrorCount++
}

func (r *testReport) PrintReport(finished bool) {
	duration := "-"
	if r.Duration != nil {
		if r.Duration.Seconds() >= 1 {
			duration = fmt.Sprintf("%.1fs", r.Duration.Seconds())
		} else {
			duration = fmt.Sprintf("%dms", r.Duration.Milliseconds())
		}
	}

	total2XXRequests := int64(0)
	for key, value := range r.StatusCodeCount {
		if key >= 200 && key <= 299 {
			total2XXRequests = total2XXRequests + value
		}
	}

	twInner := table.NewWriter()
	twInner.SetStyle(table.StyleLight)
	twInner.Style().Options.SeparateRows = true
	twInner.Style().Options.SeparateColumns = true
	twInner.Style().Options.DrawBorder = true

	if len(r.StatusCodeCount) == 0 {
		twInner.AppendRows([]table.Row{{"-", "-"}})
	}

	for key, value := range r.StatusCodeCount {
		twInner.AppendRows([]table.Row{{key, value}})
	}

	tw := table.NewWriter()
	tw.SetStyle(table.StyleLight)
	tw.Style().Options.SeparateRows = true
	tw.Style().Options.SeparateColumns = false
	tw.Style().Options.DrawBorder = true
	if finished {
		tw.SetTitle(fmt.Sprintf("%s - TEST REPORT", APP_NAME))
	} else {
		tw.SetTitle(fmt.Sprintf("%s - TEST REPORT (INTERRUPTED)", APP_NAME))
	}

	tw.AppendRows([]table.Row{
		{"URL", r.URL},
		{"Runners", r.Runners},
		{"Duration", duration},
		{"Requests", r.RequestCount},
		{"2XX Requests", total2XXRequests},
		{"Requests per Status", twInner.Render()},
		{"Errors", r.ErrorCount},
	})

	fmt.Println(tw.Render())
}
