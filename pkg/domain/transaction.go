package domain

import (
	"database/sql/driver"
)

// DealWithErrorInTransaction err 參數, 預期不等於 nil
func DealWithErrorInTransaction(tx driver.Tx, err error) error {
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return rollbackErr
		}
	}
	return nil
}
