package loader

import (
	"fmt"
	"image"
	"io"
	"path"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"

	"github.com/mymmrac/hide-and-seek/assets"
)

func Image(filePath string) (*ebiten.Image, error) {
	imageFile, err := assets.FS.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("open %q: %w", filePath, err)
	}

	decodedImage, _, err := image.Decode(imageFile)
	if err != nil {
		return nil, fmt.Errorf("decode %q image: %w", filePath, err)
	}

	return ebiten.NewImageFromImage(decodedImage), nil
}

func Audio(ctx *audio.Context, filePath string, infinite bool) (*audio.Player, error) {
	audioFile, err := assets.FS.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("load audio: %w", err)
	}

	var audioStream io.Reader
	switch ext := path.Ext(filePath); ext {
	case ".ogg":
		var oggStream *vorbis.Stream
		oggStream, err = vorbis.DecodeWithoutResampling(audioFile)
		if err != nil {
			return nil, fmt.Errorf("decode ogg audio: %w", err)
		}

		if infinite {
			audioStream = audio.NewInfiniteLoop(oggStream, oggStream.Length())
		} else {
			audioStream = oggStream
		}
	default:
		return nil, fmt.Errorf("unsupported audio format: %s", ext)
	}

	audioPlayer, err := ctx.NewPlayer(audioStream)
	if err != nil {
		return nil, fmt.Errorf("new audio player: %w", err)
	}

	return audioPlayer, nil
}
