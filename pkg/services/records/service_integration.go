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
	wiseRecordManager
}

type skillRecordManager interface {
	Skills() SkillRecords
	GetSkillByName(name string) (*models.Record, error)
}

type stockRecordManager interface {
	Stocks() StockRecords
	GetStockByName(name string) (*models.Stock, error)
}

type traitRecordManager interface {
	Traits() TraitRecords
	GetTraitByName(name string) (*models.TraitRecord, error)
}

type wiseRecordManager interface {
	Wises() WisesRecords
	GetWiseByName(name string) (*models.WiseRecord, error)
}
