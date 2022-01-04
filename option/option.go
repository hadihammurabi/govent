package option

type Option struct {
	Name string
	Args []interface{}
}

func NewOption(name string, args ...interface{}) *Option {
	return &Option{
		name,
		args,
	}
}
