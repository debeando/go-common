package cli

type Arg struct {
	Name        string
	Description string
	Required    bool
	Type        uint8
	Default     any
	Function    func()
	Value       string
}

type Args []Arg

func (s *Args) Add(o Arg) {
	if !s.unique(o) {
		*s = append(*s, o)
	}
}

func (s *Args) unique(o Arg) bool {
	for _, i := range *s {
		if i.Name == o.Name {
			return true
		}
	}
	return false
}

func (s *Args) NamesLength() (l int) {
	for _, i := range *s {
		if len(i.Name) > l {
			l = len(i.Name)
		}
	}
	return l
}
