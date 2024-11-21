package meta

import (
	"bytes"
	"image"
	"image/draw"
	"image/jpeg"
)

func (order Order) Convert(source []byte) ([]byte, error) {

	srcImage, err := jpeg.Decode(bytes.NewReader(source))
	if err != nil {
		return nil, err
	}
	// Create a new canvas
	canvas := image.NewRGBA(image.Rect(0, 0, order.Scramble.W, order.Scramble.H))

	// Apply crops
	for _, crop := range order.Scramble.Crops {
		// Define the source rectangle
		srcRect := image.Rect(crop.X2, crop.Y2, crop.X2+crop.W, crop.Y2+crop.H)

		// Define the destination rectangle
		destRect := image.Rect(crop.X, crop.Y, crop.X+crop.W, crop.Y+crop.H)

		// Draw the cropped image onto the canvas
		draw.Draw(canvas, destRect, srcImage, srcRect.Min, draw.Src)
	}
	var writer bytes.Buffer

	err = jpeg.Encode(&writer, canvas, nil)
	if err != nil {
		return nil, err
	}
	return writer.Bytes(), nil
}
