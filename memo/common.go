package memo

import (
	"database/sql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

// grpc 오류코드에 대해서는 여기를 참고 https://github.com/grpc/grpc/blob/master/doc/statuscodes.md
func handleError(err error) error {
	// Return an error message based on the error type
	switch err {
	case nil:
		return nil
	case sql.ErrNoRows:
		slog.Info("", "error", err)
		return status.Errorf(codes.NotFound, "Not Found Data.")
	default:
		// Log the error
		slog.Error("", "error", err)

		return status.Errorf(codes.Internal, "Internal Server Error.")
	}
}
