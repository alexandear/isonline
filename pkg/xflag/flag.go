package xflag

type StringFlag struct {
	isSet bool
	value string
}

func (f *StringFlag) IsSet() bool {
	return f.isSet
}

func (f *StringFlag) Value() string {
	return f.value
}

func (f *StringFlag) Set(val string) error {
	f.value = val
	f.isSet = true
	return nil
}

func (f *StringFlag) String() string {
	return f.value
}
