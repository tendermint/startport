package clispinner

import (
	"time"

	"github.com/briandowns/spinner"
)

var (
	refreshRate  = time.Millisecond * 200
	charset      = spinner.CharSets[4]
	colorSpinner = "blue"
)

type Spinner struct {
	sp *spinner.Spinner
}

// New creates a new spinner.
func New() *Spinner {
	sp := spinner.New(charset, refreshRate)
	sp.Color(colorSpinner)
	s := &Spinner{
		sp: sp,
	}
	s.SetText("Initializing...")
	return s
}

// SetText sets the text for spinner.
func (s *Spinner) SetText(text string) *Spinner {
	s.sp.Suffix = " " + text
	return s
}

// SetPreText sets the prefix for spinner.
func (s *Spinner) SetPreText(text string) *Spinner {
	s.sp.Prefix = text + " "
	return s
}

// SetPreText sets the prefix for spinner.
func (s *Spinner) SetCharset(charset []string) *Spinner {
	s.sp.UpdateCharSet(charset)
	return s
}

// SetPreText sets the prefix for spinner.
func (s *Spinner) SetColor(color string) *Spinner {
	s.sp.Color(color)
	return s
}

// Start starts spinning.
func (s *Spinner) Start() *Spinner {
	s.sp.Start()
	return s
}

// Stop stops spinning.
func (s *Spinner) Stop() *Spinner {
	s.sp.Stop()
	s.SetColor(colorSpinner)
	s.SetPreText("")
	s.SetCharset(charset)
	s.sp.Stop()
	return s
}
