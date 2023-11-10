package records

import (
	"torchbearer/pkg/models"
)

type Dependency = RecordsManager

type RecordsManager interface {
	Ready() bool
	skillRecordManager
	stockRecordManager
	traitRecordManager
}

type skillRecordManager interface {
	Skills() SkillRecords
	GetSkill(name string) (*models.SkillRecord, error)
}

type stockRecordManager interface {
	Stocks() StockRecords
	GetStock(name string) (*models.Stock, error)
}

type traitRecordManager interface {
	Traits() TraitRecords
	GetTraits(name string) (*models.TraitRecord, error)
}
