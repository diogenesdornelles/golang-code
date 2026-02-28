package data

import (
	"github.com/diogenesdornelles/pgormsqlc/generated"
	"github.com/jackc/pgx/v5/pgtype"
)

var Autor = generated.CreateAuthorParams{
	Name: "Brian Kernighan",
	Bio:  pgtype.Text{String: "Co-author of The C Programming Language and The Go Programming Language", Valid: true},
}
