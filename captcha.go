package xcaptcha

import (
	"bytes"
	captcha "github.com/steambap/captcha"
	"image/color"
)

type Option struct {
	Width      int
	Height     int
	TextLength int
}

func New(opt Option) (code string, buffer *bytes.Buffer, err error) {
	if opt.Width == 0 {
		opt.Width = 150
		opt.Height = 50
	}
	if opt.TextLength == 0 {
		opt.TextLength = 4
	}
	data, err := captcha.New(opt.Width, opt.Height, func(options *captcha.Options) {
		options.BackgroundColor = color.RGBA{1, 1, 1, 1}
		options.CharPreset = "abcdefghijkmnpqrstuvwxyz23456789"
		options.TextLength = opt.TextLength
		options.FontScale = 1
		options.FontDPI = 90
		return
	})
	if err != nil {
		return
	}
	code = data.Text
	buffer = bytes.NewBuffer([]byte(nil))
	err = data.WriteImage(buffer)
	if err != nil {
		return
	}
	return
}
