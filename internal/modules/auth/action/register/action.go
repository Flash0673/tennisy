package register

type Action struct{}

func New() *Action {
	return &Action{}
}

func (action *Action) Do() ([]byte, error) {
	return []byte("Hello world"), nil
}
