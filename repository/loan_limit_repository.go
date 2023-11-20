package repository

import (
	"xyz-multifinance/model/domain"

	"gorm.io/gorm"
)

type LoanLimitRepository interface {
	FindAll(db *gorm.DB, filters *map[string]string) domain.LoanLimits
	Create(db *gorm.DB, loanLimit *domain.LoanLimit) *domain.LoanLimit
	Update(db *gorm.DB, loanLimitID *string, loanLimit *domain.LoanLimit) *domain.LoanLimit
	Delete(db *gorm.DB, loanLimitID *string)
}
