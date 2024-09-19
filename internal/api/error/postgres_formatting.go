package error

import (
	"errors"
	"net/http"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jegj/linktly/internal/api/types"
)

func PostgresFormatting(dbError error) types.APIError {
	if pgErr, ok := dbError.(*pgconn.PgError); ok {
		if pgErr.Code == pgerrcode.UniqueViolation {
			return types.APIError{
				Msg:        pgErr.Error(),
				StatusCode: http.StatusConflict,
			}
		} else {
			return types.APIError{
				Msg:        pgErr.Error(),
				StatusCode: http.StatusInternalServerError,
			}
		}
	} else if errors.Is(dbError, pgx.ErrNoRows) {
		return types.APIError{
			Msg:        dbError.Error(),
			StatusCode: http.StatusNotFound,
		}
	} else {
		return types.APIError{
			Msg:        dbError.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
}
