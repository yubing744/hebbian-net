package interfaces

type IModule interface {
	BindInput(name string, inputModule IModule, outputIndex int)
}
