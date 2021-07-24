package xcaptcha

import (
	captcha "github.com/steambap/captcha"
	"image/color"
	"io"
)

type Option struct {
	Width int
	Height int
	TextLength int
}
func New(writer io.Writer, opt Option) (code string, err error) {
	if opt.Width == 0 {
		opt.Width = 150
		opt.Height = 50
	}
	if opt.TextLength == 0 {
		opt.TextLength = 4
	}
	data, err := captcha.New(opt.Width, opt.Height, func(options *captcha.Options) {
		options.BackgroundColor = color.RGBA{1,1,1,1}
		options.CharPreset = "abcdefghijkmnpqrstuvwxyz23456789"
		options.TextLength = opt.TextLength
		return
	})
	code = data.Text
	err = data.WriteImage(writer) ; if err != nil {
	    return
	}
	return
}