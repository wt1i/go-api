package persistence

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// TxErrDefer commit or revert tx based on err, passing err to return
func TxErrDefer(tx *gorm.DB, err error) error {
	if r := recover(); r != nil {
		tx.Rollback()
		panic(r)
	}

	var commitErr error

	if err != nil {
		commitErr = tx.Rollback().Error
	} else {
		commitErr = tx.Commit().Error
	}

	if commitErr == nil {
		return errors.WithStack(err)
	}

	return errors.Wrap(errors.WithStack(err), "TxErrDefer err: "+commitErr.Error())
}

func IsDBError(err error) bool {
	return err != nil && !errors.Is(err, gorm.ErrRecordNotFound)
}

func IsDBRecordNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
