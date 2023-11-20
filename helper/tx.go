package helper

import (
	"strings"

	"gorm.io/gorm"
)

func CommitOrRollback(tx *gorm.DB) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback().Error
		PanicIfError(errorRollback)
		panic(err)
	} else {
		errorCommit := tx.Commit().Error
		if errorCommit != nil && strings.Contains(errorCommit.Error(), "transaction has already been committed or rolled back") {
			return
		}
		PanicIfError(errorCommit)
	}
}
