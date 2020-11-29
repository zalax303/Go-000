package main

import (
	"database/sql"
	"github.com/pkg/errors"
	"log"
)

type People struct {
	Name string
}

func (p *People) NoDataError() error {
	return sql.ErrNoRows
}

func BizPeople(Name string) (*People, error) {
	p := &People{Name}
	if err := p.NoDataError(); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.WithMessage(err, "no data")
		} else {
			return nil, errors.Wrap(err, " people struct")
		}

	}

	return p, nil
}

func main() {
	name := "aaa"
	p, err := BizPeople(name)
	if err != nil {
		log.Println(err)

	}

	log.Printf("data info: %+v\n", p)
}
