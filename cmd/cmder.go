package cmd

type Cmder interface {
	Handle(args []string) error
	Exec() error
}
