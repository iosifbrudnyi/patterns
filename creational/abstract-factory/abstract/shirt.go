package abstract

type IShort interface {
	SetLogo(logo string)
	SetSize(size int)
	GetLogo() string
	GetSize() int
}

type Short struct {
	Logo string
	Size int
}

func (s *Short) SetLogo(logo string) {
	s.Logo = logo
}

func (s *Short) GetLogo() string {
	return s.Logo
}

func (s *Short) SetSize(size int) {
	s.Size = size
}

func (s *Short) GetSize() int {
	return s.Size
}
