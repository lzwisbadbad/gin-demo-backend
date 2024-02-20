package db

const DATAMANAGER_TABLE_NAME = "data_manager"

type DataManager struct {
	GeneralField
	DataId      string `gorm:"uniqueIndex"`
	DataContent string
	AddDataTime int64
}

func (dm *DataManager) TableName() string {
	return DATAMANAGER_TABLE_NAME
}

func init() {
	dataManager := new(DataManager)
	TableSlice = append(TableSlice, &dataManager)
}
