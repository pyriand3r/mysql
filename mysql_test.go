package mysql_test

import (
	"platform/mock"
	"platform/mysql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestStatusCheck(t *testing.T) {
	db, mocked := mock.MySQL(t, "test")
	req := require.New(t)

	rows := sqlmock.NewRows([]string{"column"}).AddRow("true")
	mocked.ExpectQuery("SELECT true").WillReturnRows(rows)

	req.NoError(mysql.Check(db))
}

func TestStatusCheck_NoResponse(t *testing.T) {
	db, mocked := mock.MySQL(t, "test")
	req := require.New(t)

	rows := sqlmock.NewRows([]string{"column"})
	mocked.ExpectQuery("SELECT true").WillReturnRows(rows)

	err := mysql.Check(db)
	req.Error(err)
	req.Contains(err.Error(), "no rows in result")
}
