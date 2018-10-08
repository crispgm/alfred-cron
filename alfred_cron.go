package main

import (
	"os"

	api "github.com/crispgm/alfred-cron/cron"
	"github.com/deanishe/awgo"
	"github.com/robfig/cron"
)

var wf *aw.Workflow

func run() {
	if len(os.Args) > 1 {
		arg := os.Args[1]
		parser := cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)
		_, err := parser.Parse(arg)
		if err != nil {
			wf.Warn("Invalid expression", "Input something like `*/5 * * * *`")
			return
		}
		text, err := api.Call(arg)
		if err == nil {
			wf.NewItem(text)
			wf.SendFeedback()
			return
		}
		wf.Warn("Call API failed", "Try it later")
	}
	wf.WarnEmpty("No expression", "Input something like `*/5 * * * *`")
}

func main() {
	wf = aw.New()
	wf.Run(run)
}
