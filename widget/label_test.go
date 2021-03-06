package widget

import "testing"

import "github.com/stretchr/testify/assert"

import _ "github.com/fyne-io/fyne/test"
import "github.com/fyne-io/fyne/theme"

func TestLabelSize(t *testing.T) {
	label := NewLabel("Test")
	min := label.MinSize()

	assert.True(t, min.Width > theme.Padding()*2)

	label.SetText("Longer")
	assert.True(t, label.MinSize().Width > min.Width)
}
