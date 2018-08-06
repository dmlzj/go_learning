package models

import (
		"github.com/revel/revel"
)

type Course struct {
	Id 			int64 `db:"id" json:"id"`
	Title 		string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
}

func (c *Course) Validate(v *revel.Validation){
	v.Check(c.Title,
		revel.ValidRequired(),
		)
	v.Check(c.Description,
		revel.ValidRequired(),
		)
}

