package output

import "github.com/yubing744/hebbian-net/pkg/types"

type Output struct {
	Video *types.Video
	Wave  *types.Wave
	Image *types.Image
	Text  string
}
