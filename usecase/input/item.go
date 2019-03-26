package input

type GetItem struct {
	ID string
}

type PostItem struct {
	Name string `json:"name"`
}

func (p *PostItem) Validate() bool {
	if len(p.Name) < 1 || 25 < len(p.Name) {
		return false
	}
	return true
}
