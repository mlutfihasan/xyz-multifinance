package repository

import (
	"strconv"

	"xyz-multifinance/helper"
	"xyz-multifinance/model/domain"

	"gorm.io/gorm"
)

type LoanLimitRepositoryImpl struct {
}

func NewLoanLimitRepository() LoanLimitRepository {
	return &LoanLimitRepositoryImpl{}
}

func (repository *LoanLimitRepositoryImpl) FindAll(db *gorm.DB, filters *map[string]string) domain.LoanLimits {
	loanLimits := domain.LoanLimits{}
	tx := db.Model(&domain.LoanLimit{})

	err := helper.ApplyFilter(tx, filters)
	helper.PanicIfError(err)

	err = tx.Find(&loanLimits).Error
	helper.PanicIfError(err)

	return loanLimits
}

func (repository *LoanLimitRepositoryImpl) Create(db *gorm.DB, loanLimit *domain.LoanLimit) *domain.LoanLimit {
	err := db.Create(&loanLimit).Error
	helper.PanicIfError(err)
	return loanLimit
}

func (repository *LoanLimitRepositoryImpl) Update(db *gorm.DB, loanLimitID *string, loanLimit *domain.LoanLimit) *domain.LoanLimit {
	idParsed, err := strconv.ParseUint(*loanLimitID, 10, 32)
	helper.PanicIfError(err)

	loanLimitIdParsed := uint(idParsed)

	err = db.Model(&domain.LoanLimit{}).
		Where(&domain.LoanLimit{
			LoanLimitID: loanLimitIdParsed,
		}).
		Updates(&loanLimit).Error
	helper.PanicIfError(err)

	err = db.First(&loanLimit, &domain.LoanLimit{LoanLimitID: loanLimitIdParsed}).Error
	helper.PanicIfError(err)

	return loanLimit
}

func (repository *LoanLimitRepositoryImpl) Delete(db *gorm.DB, loanLimitID *string) {
	idParsed, err := strconv.ParseUint(*loanLimitID, 10, 32)
	helper.PanicIfError(err)

	loanLimitIdParsed := uint(idParsed)

	loanLimit := &domain.LoanLimit{}
	tx := db.First(loanLimit, &domain.LoanLimit{LoanLimitID: loanLimitIdParsed}).Updates(&domain.LoanLimit{
		LoanLimitID: loanLimitIdParsed,
	})

	// Deleting the LoanLimit from the database.
	err = tx.Unscoped().Delete(loanLimit, &domain.LoanLimit{LoanLimitID: loanLimitIdParsed}).Error
	helper.PanicIfError(err)
}
