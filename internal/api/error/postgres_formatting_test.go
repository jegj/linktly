package error

import (
	"net/http"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/jegj/linktly/internal/api/types"
)

type pgxErrorBase struct {
	TableName        string
	Code             string
	Message          string
	Detail           string
	Hint             string
	Where            string
	SchemaName       string
	Severity         string
	ColumnName       string
	DataTypeName     string
	ConstraintName   string
	Position         int32
	InternalPosition int32
}

func TestPostgresFormattingCommonPostgresCode(t *testing.T) {
	tests := []struct {
		want     types.APIError
		testName string
		input    pgxErrorBase
	}{
		{
			testName: "Must return an APIError with the correct http status code based on the postgres error",
			input: pgxErrorBase{
				Severity:         "ERROR",
				Code:             "23505", // Example: unique violation code
				Message:          "duplicate key value violates unique constraint",
				Detail:           "Key (id)=(1) already exists.",
				Hint:             "Try to insert a different value.",
				Position:         15,
				InternalPosition: 10,
				Where:            "INSERT INTO users",
				SchemaName:       "public",
				TableName:        "users",
				ColumnName:       "id",
				DataTypeName:     "integer",
				ConstraintName:   "users_pkey",
			},
			want: types.APIError{
				Msg:        "ERROR: duplicate key value violates unique constraint (SQLSTATE 23505)",
				StatusCode: http.StatusConflict,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			pgErr := &pgconn.PgError{
				Code:             tt.input.Code,
				Message:          tt.input.Message,
				Detail:           tt.input.Detail,
				Hint:             tt.input.Hint,
				Position:         tt.input.Position,
				InternalPosition: tt.input.InternalPosition,
				Where:            tt.input.Where,
				SchemaName:       tt.input.SchemaName,
				TableName:        tt.input.TableName,
				ColumnName:       tt.input.ColumnName,
				DataTypeName:     tt.input.DataTypeName,
				ConstraintName:   tt.input.ConstraintName,
				Severity:         tt.input.Severity,
			}
			got := PostgresFormatting(pgErr)
			if got != tt.want {
				t.Errorf("PostgresFormatting( Code %q ) = %v, want %v", tt.input.Code, got, tt.want)
			}
		})
	}
}

func TestPostgresFormattingNoRowsError(t *testing.T) {
	tests := []struct {
		input    error
		want     types.APIError
		testName string
	}{
		{
			testName: "Must return an APIError with the http status code 404 when the query does not return rows",
			input:    pgx.ErrNoRows,
			want: types.APIError{
				Msg:        "no rows in result set",
				StatusCode: http.StatusNotFound,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			got := PostgresFormatting(tt.input)
			if got != tt.want {
				t.Errorf("PostgresFormatting( Code %q ) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
