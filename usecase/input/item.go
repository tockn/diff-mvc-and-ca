package input

import "errors"

type GetItem struct {
	ID string
}

type PostItem struct {
	Name string
}

func (p *PostItem) Validate() error {
	if len(p.Name) < 1 || 25 < len(p.Name) {
		return errors.New("validation error")
	}
	return nil
}
