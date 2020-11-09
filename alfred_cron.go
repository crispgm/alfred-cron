package main

import (
	"os"

	aw "github.com/deanishe/awgo"
	"github.com/lnquy/cron"
)

var wf *aw.Workflow

func run() {
	if len(os.Args) > 1 {
		arg := os.Args[1]

		exprDesc, _ := cron.NewDescriptor()
		desc, err := exprDesc.ToDescription(arg, cron.Locale_en)
		if err != nil {
			wf.Warn("Invalid expression", "Input something like `*/5 * * * *`")
			return
		}
		if err == nil {
			wf.NewItem(desc).
				Arg(desc).Valid(true)
			wf.SendFeedback()
			return
		}
	}
}

func main() {
	wf = aw.New()
	wf.Run(run)
}
