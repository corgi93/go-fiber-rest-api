/*
	gorm같은 orm을 나중에 추가할 예정.
	우선 pure한 sql 쿼리로 작업하도록하는데 장점으로 쿼리 최적화하는 등의 장점도 있다.
*/

package queries

import (
	"fmt"

	"github.com/corgi93/go-fiber-rest-api/app/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// book model에 대한 쿼리 구조체
type BookQueries struct {
	*sqlx.DB
}

// 모든 책 조회
func (q *BookQueries) GetBooks() ([]models.Book, error) {
	// books변수 선언
	books := []models.Book{}

	query := `SELECT * FROM books`
	fmt.Println("query..", &q)

	//  query를 db로 날림 - q.Get -> q.Select로 수정.
	// err := q.Get(&books, query)
	err := q.Select(&books, query)
	if err != nil {
		fmt.Println("err..", err)
		// Return empty object and error.
		return books, err
	}
	fmt.Println("체크포인트")

	// query result
	return books, nil

}

// 책 한권 조회
func (q *BookQueries) GetBook(id uuid.UUID) (models.Book, error) {
	book := models.Book{}

	query := `SELECT * FROM books WHERE id = $1`

	err := q.Get(&book, query, id)
	if err != nil {
		return book, err
	}

	return book, nil
}

// 책 추가하기
func (q *BookQueries) CreateBook(b *models.Book) error {
	// query
	query := `INSERT INTO books VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	// fmt.Printf("Book: %+v\n", *b)
	fmt.Printf("Book: %+v\n", *b)
	fmt.Println("Attr::", b.BookAttrs)
	fmt.Println("Description::", b.BookAttrs.Description)

	//  db애 query 날리기
	_, err := q.Exec(query, b.ID, b.CreatedAt, b.UpdatedAt, b.UserID, b.Title, b.Author, b.BookStatus, b.BookAttrs)
	if err != nil {
		// Return only error.
		return err
	}

	// create는 아무것도 return하지 x
	return nil
}
