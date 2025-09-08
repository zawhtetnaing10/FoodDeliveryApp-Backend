package handlers

import (
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
)

// TODO: - Improve this. See how fmt, prec and bitSize work.
func convertFloatToPgtypeNumeric(num float64) (pgtype.Numeric, error) {
	var result pgtype.Numeric

	numString := strconv.FormatFloat(num, 'f', -1, 64)

	if err := result.Scan(numString); err != nil {
		return result, err
	}
	return result, nil
}
