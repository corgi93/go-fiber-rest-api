package database

import "github.com/corgi93/go-fiber-rest-api/app/queries"

// Queries struct for collect all app queries.
type Queries struct {
	// load queries from Book model
	*queries.BookQueries
}

func OpenDBConnection() (*Queries, error) {
	db, err := PostgreSQLConnection()

	if err != nil {
		return nil, err
	}

	return &Queries{
		// Set queries from models:
		BookQueries: &queries.BookQueries{DB: db}, // from Book model
	}, nil
}
