package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"google.golang.org/grpc/status"
)

// LoadingOverlay which displays a progress bar and a string.
// It covers all available space and blocks interactions with items behind it.
type LoadingOverlay struct {
	widget.BaseWidget

	jobQueue  chan Job
	retryChan chan struct{}

	loadingCard *widget.Card
	errBox      *fyne.Container
	errCard     *widget.Card
	errLabel    *widget.Label

	root *fyne.Container
}

// Job associates a function with a description.
type Job struct {
	// A short term, ideally one word, describing what the job does.
	// E.g. "Loading assets", "Querying endpoint", ...
	Desc string
	Run  func() error
}

// NewLoadingOverlay creates a new overlay with potentially long-running jobs.
//
// The overlay is initialized without any jobs.
// It will only become visible when and as long there are jobs to process.
// It will be hidden when no jobs are left.
func NewLoadingOverlay() *LoadingOverlay {
	// semi-transparent black background to obscure main view
	bgColor := color.RGBA{0, 0, 0, 200}
	bg := canvas.NewRectangle(bgColor)

	loadingCard := widget.NewCard(
		"Loading...",
		"Ready to start...",
		widget.NewProgressBarInfinite(),
	)

	// dedicated layout to switch to when error occurs
	retryChan := make(chan struct{}, 1)
	errLabel := widget.NewLabel("n/a")
	errLabel.Importance = widget.WarningImportance
	errLabel.Wrapping = fyne.TextWrapBreak
	errCard := widget.NewCard(
		"Error occured",
		"While n/a",
		errLabel,
	)
	errBox := container.NewVBox(
		errCard,
		widget.NewButton(
			"Retry",
			func() {
				retryChan <- struct{}{}
			},
		),
	)
	errBox.Hide()

	o := &LoadingOverlay{
		jobQueue:    make(chan Job, 3),
		retryChan:   retryChan,
		loadingCard: loadingCard,
		errCard:     errCard,
		errBox:      errBox,
		errLabel:    errLabel,
		root: container.NewStack(
			bg,
			container.NewCenter(container.NewStack(
				loadingCard,
				errBox,
			)),
		),
	}
	o.ExtendBaseWidget(o)
	// start processing in the background
	go o.run()
	return o
}

// CreateRenderer creates the renderer for o.
func (o *LoadingOverlay) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(o.root)
}

func (o *LoadingOverlay) run() {
	var job Job
	for {
		select {
		// check job queue
		case job = <-o.jobQueue:
			o.Show()
			o.loadingCard.SetSubTitle(job.Desc + "...")
			for {
				err := job.Run()
				if err == nil {
					break
				}
				o.onErr(job, err)
				<-o.retryChan
				o.onRetry()
			}
		default:
			// no jobs left
			if !o.Hidden {
				o.Hide()
			}
			// wait a little before checking for jobs to avoid overhead
			time.Sleep(250 * time.Millisecond)
		}
	}
}

func (o *LoadingOverlay) onErr(job Job, err error) {
	o.errCard.SetSubTitle("While " + job.Desc)
	if grpcErr, ok := status.FromError(err); ok {
		o.errLabel.SetText(grpcErr.Message())
	} else {
		o.errLabel.SetText(err.Error())
	}
	o.loadingCard.Hide()
	o.errBox.Show()
}

func (o *LoadingOverlay) onRetry() {
	o.errBox.Hide()
	o.loadingCard.Show()
}

// AddJobs adds an arbitrary amount of jobs to the job queue.
func (o *LoadingOverlay) AddJobs(jobs ...Job) {
	// run as goroutine so AddJobs does not block
	go func() {
		for _, job := range jobs {
			o.jobQueue <- job
		}
	}()
}
