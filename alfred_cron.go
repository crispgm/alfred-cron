package main

import (
	"os"

	"github.com/crispgm/alfred-cron/cron"
	"github.com/deanishe/awgo"
)

var wf *aw.Workflow

func run() {
	if len(os.Args) > 1 {
		text, err := cron.Call(os.Args[1])
		if err == nil {
			wf.NewItem(text)
			wf.SendFeedback()
		}
	}
	wf.WarnEmpty("No expression", "Input something like `1/* * * * *`")
}

func main() {
	wf = aw.New()
	wf.Run(run)
}
