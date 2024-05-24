package models

type TblHotData struct {
	Id         int32  `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL;comment:'主键'"`
	OriId      int32  `gorm:"column:ori_id;NOT NULL;comment:'ori_id'"`
	OrderBy    int32  `gorm:"column:order_by;NOT NULL;comment:'顺序'"`
	Category   string `gorm:"column:category;NOT NULL;comment:'类别'"`
	CreateTime int32  `gorm:"column:create_time;NOT NULL;comment:'创建时间'"`
	UpdateTime int32  `gorm:"column:update_time;NOT NULL;comment:'更新时间'"`
}

func (*TblHotData) TableName() string {
	return "tblHotData"
}
