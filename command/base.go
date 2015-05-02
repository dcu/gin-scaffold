package command

type Base interface {
	Execute(args []string)
}
