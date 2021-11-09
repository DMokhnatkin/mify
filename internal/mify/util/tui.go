package util

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"
)

func CreateWaitSpinner() *spinner.Spinner {
	return spinner.New(spinner.CharSets[9], 200*time.Millisecond)
}


type ProgressBar struct {
	bar *mpb.Bar
	progress *mpb.Progress
	statusFunc decor.DecorFunc
	total int64

	spinCycle int
	spinnerChars []string

}

func NewProgressBar(statusFunc decor.DecorFunc) *ProgressBar {
	return &ProgressBar{
		statusFunc: statusFunc,
		spinnerChars: []string{"/", "-", "\\", "|"},
	}
}

func (pb *ProgressBar) Create(total int64) {
	pb.total = total
	pb.progress = mpb.New(mpb.WithWidth(64))
	pb.bar = pb.progress.Add(pb.total,
		mpb.NewBarFiller(mpb.BarStyle().Lbound("[").Filler("=").Tip(">").Padding("-").Rbound("]")),
		mpb.PrependDecorators(
			decor.OnComplete(
				decor.Any(func(s decor.Statistics) string {
					return pb.statusFunc(s)
				}),
				"done",
			),
		),
		mpb.AppendDecorators(decor.Any(func(s decor.Statistics) string {
			return fmt.Sprintf("%d/%d", s.Current, s.Total)
		})),
	)
}

func (pb *ProgressBar) Abort() {
	if pb.bar == nil {
		return
	}
	pb.bar.Abort(true)
}

func (pb *ProgressBar) IncTotal() {
	if pb.bar == nil {
		return
	}
	pb.total += 1
	if pb.bar.Completed() {
		pb.Create(0)
	}
	pb.bar.SetTotal(pb.total, false)
}

func (pb *ProgressBar) Increment() {
	if pb.bar == nil {
		return
	}
	pb.bar.Increment()
}

func (pb *ProgressBar) Wait() {
	if pb.bar == nil {
		return
	}
	pb.bar.SetTotal(pb.total, true)
	pb.progress.Wait()
}

func (pb *ProgressBar) Spinner() string {
	char := pb.spinnerChars[pb.spinCycle]
	pb.spinCycle += 1
	if pb.spinCycle == len(pb.spinnerChars) {
		pb.spinCycle = 0
	}
	return char
}
