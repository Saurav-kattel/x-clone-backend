package test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/user"
)

func TestInsertImageId(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer mockDB.Close()

	db := sqlx.NewDb(mockDB, "sqlmock")
	parsedUuid := uuid.MustParse("96274110-f84c-46ab-80db-fe6fb4623ebb")
	newParsedUuid := uuid.MustParse("96274110-f84c-46ab-80db-fe6fb4623ebf")

	// Expecting an INSERT statement and returning a result indicating success
	mock.ExpectExec(`UPDATE users SET image_id = \$1 WHERE id = \$2`).
		WithArgs(newParsedUuid, parsedUuid).
		WillReturnResult(sqlmock.NewResult(1, 1))

	insertImgErr := user.InsertImageId(db, newParsedUuid.String(), parsedUuid.String())
	if insertImgErr != nil {
		t.Fatalf("error inserting image id %v", insertImgErr)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}
