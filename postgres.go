/*
 * @author    Emmanuel Kofi Bessah
 * @email     bekinsoft@gmail.com
 */

package util

import (
	"fmt"
	"strings"

	"github.com/lib/pq"
)

// Constraints Constants
var (
	ConstraintDuplicateKey = "violates unique constraint \"primary\""
)

// IsConstraintError checks whether the given error is a unique constraint error or not
func IsConstraintError(err error, inerr, constraintName string) (ok bool, errr error) {
	if pqErr, ok := err.(*pq.Error); ok {
		if pqErr.Code == "23505" {
			if strings.Contains(pqErr.Message, ConstraintDuplicateKey) {
				return true, fmt.Errorf("Primary Key already exists")
			} else if pqErr.Constraint == constraintName || strings.Contains(pqErr.Message, constraintName) {
				return true, fmt.Errorf(inerr)
			}
		}
	}
	return false, nil
}
