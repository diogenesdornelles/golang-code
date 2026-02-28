package main

import (
	"context"
	"log"
	"reflect"

	"github.com/diogenesdornelles/pgormsqlc/generated"

	"github.com/diogenesdornelles/pgormsqlc/internal/platform"

	"github.com/diogenesdornelles/pgormsqlc/data"
)

func run() error {
	ctx := context.Background()

	config := platform.LoadConfig()
	conn, err := platform.ConnectDatabase(config, ctx)
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	queries := generated.New(conn)

	// list all authors
	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		return err
	}
	log.Println(authors)

	// create an author
	insertedAuthor, err := queries.CreateAuthor(ctx, data.Autor)
	if err != nil {
		return err
	}
	log.Println(insertedAuthor)

	// get the author we just inserted
	fetchedAuthor, err := queries.GetAuthor(ctx, insertedAuthor.ID)
	if err != nil {
		return err
	}

	// prints true
	log.Println(reflect.DeepEqual(insertedAuthor, fetchedAuthor))
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
