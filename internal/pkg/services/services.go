package services

type Service struct {
	Name        string
	ComposeFile string
	Init        Initialize
	Context     string
}

type Initialize func() error

func init() {

}
