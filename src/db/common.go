package db

var TableSlice = make([]interface{}, 0)

type GeneralField struct {
	Id       int32 `gorm:"primaryKey;autoIncrement"`
	LastTime int64 `gorm:"autoCreateTime:milli;index"`
}

type ModelStruct interface {
	TableName() string
}
