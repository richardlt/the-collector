package files

import (
	"bytes"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"

	"github.com/nfnt/resize"
	errorsP "github.com/pkg/errors"
	"github.com/richardlt/the-collector/server/api/errors"
	"github.com/richardlt/the-collector/server/types"
)

// ResizeImageFile .
func ResizeImageFile(file *types.File, data []byte,
	size string) ([]byte, error) {
	var buf bytes.Buffer
	r := io.TeeReader(bytes.NewReader(data), &buf)

	i, _, err := image.Decode(r)
	if err != nil {
		return nil, errorsP.WithStack(err)
	}

	c, _, err := image.DecodeConfig(&buf)
	if err != nil {
		return nil, errorsP.WithStack(err)
	}

	w, h := calculateImageSize(c, types.SizeToPixels(size))
	ir := resize.Resize(w, h, i, resize.NearestNeighbor)

	var res bytes.Buffer
	switch file.ContentType {
	case types.ImageJpeg:
		err = jpeg.Encode(&res, ir, nil)
	case types.ImagePng:
		err = png.Encode(&res, ir)
	case types.ImageGif:
		err = gif.Encode(&res, ir, nil)
	default:
		return nil, errors.NewData("invalid file content type for encoding")
	}
	if err != nil {
		return nil, errorsP.WithStack(err)
	}

	return res.Bytes(), nil
}

// calculateImageSize .
func calculateImageSize(c image.Config, size int) (uint, uint) {
	if c.Width <= size && c.Height <= size {
		return uint(c.Width), uint(c.Height)
	}

	min := c.Width
	if c.Height < c.Width {
		min = c.Height
	}

	scale := float64(min) / float64(size)
	return uint(float64(c.Width) / scale),
		uint(float64(c.Height) / scale)
}
