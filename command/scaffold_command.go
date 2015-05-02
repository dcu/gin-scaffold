package command

type ScaffoldCommand struct {
}

func (command *ScaffoldCommand) Execute(args []string) {
	modelCommand := &ModelCommand{}
	modelCommand.Execute(args)

	controllerCommand := &ControllerCommand{}
	controllerCommand.Execute([]string{modelCommand.ModelNamePlural})
}
