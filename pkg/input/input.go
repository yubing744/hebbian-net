package input

import "github.com/yubing744/hebbian-net/pkg/types"

type Input struct {
	Video   *types.Video
	Wave    *types.Wave
	Sensors []*types.Tensor
}
