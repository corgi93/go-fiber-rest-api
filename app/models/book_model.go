package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

// book 객체를 설명하는 Book struct
// 정의된 Book 구조체는 db의 books 테이블과 매핑되고 JSON 직렬화/역직렬화로 api요청과 응답에서 사용될 수 있고
// 또한 필드들에 태그를 사용해 DB와 JSON(field 값을 지정)간 매핑과 유효성 검사를 지정할 수 있습니다.
type Book struct {
	ID         uuid.UUID `db:"id" json:"id" validate:required,uuid`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
	UserID     uuid.UUID `db:"user_id" json:"user_id" validate:"required,uuid"`
	Title      string    `db:"title" json:"title" validate:"required,lte=255"`
	Author     string    `db:"author" json:"author" validate:"required,lte=255"`
	BookStatus int       `db:"book_status" json:"book_status" validate:"required,len=1"`
	BookAttrs  BookAttrs `db:"book_attrs" json:"book_attrs" validate:"required,dive"`
}

type BookAttrs struct {
	Picture     string `json:"picture"`
	Description string `json:"description"`
	Rating      int    `json:"rating" validate:"min=1,max=10"`
}

// 구조체의 JSON 인코딩 expression을 반환
func (b BookAttrs) Value() (driver.Value, error) {
	jsonData, err := json.Marshal(b)
	if err != nil {
		return nil, err // Marshal 오류 발생 시 오류 반환
	}

	// JSON 데이터를 문자열로 변환하여 반환
	return string(jsonData), nil
}

// Scan함수는 BookAttrs 구조체가 sql.Scanner 인터페이스를 구현하도록 함
// 이 방법은 JSON 인코딩된 값을 단순히 구조 field로 디코딩합니다.
func (b *BookAttrs) Scan(value interface{}) error {
	j, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []type failed")
	}

	return json.Unmarshal(j, &b)
}
