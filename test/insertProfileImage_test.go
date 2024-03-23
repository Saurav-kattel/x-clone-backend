package test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/user"
)

func TestInsertProfileImage(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer mockDB.Close()

	db := sqlx.NewDb(mockDB, "sqlmock")
	parsedUuid := uuid.MustParse("96274110-f84c-46ab-80db-fe6fb4623ebb")
	newParsedUuid := uuid.MustParse("96274110-f84c-46ab-80db-fe6fb4623ebf")

	// Expecting an INSERT statement and returning a result indicating success
	mock.ExpectQuery(`INSERT INTO images\(image,userId\) VALUES\(\$1, \$2\) RETURNING id`).
		WithArgs([]byte{11, 22, 33, 44, 55, 66}, parsedUuid).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(newParsedUuid))

	// Function invocation
	uid, err := user.InsertProfileImage(db, parsedUuid, []byte{11, 22, 33, 44, 55, 66})
	t.Log(uid)
	if err != nil {
		t.Fatalf("Error inserting profile image: %v", err)
	}

	// Verify that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
