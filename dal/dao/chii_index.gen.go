// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"gorm.io/plugin/soft_delete"
)

const TableNameIndex = "chii_index"

// Index mapped from table <chii_index>
type Index struct {
	ID           uint32                `gorm:"column:idx_id;type:mediumint(8);primaryKey;autoIncrement:true;comment:自动id"` // 自动id
	Type         uint8                 `gorm:"column:idx_type;type:tinyint(3) unsigned;not null"`
	Title        string                `gorm:"column:idx_title;type:varchar(80);not null;comment:标题"`                       // 标题
	Desc         string                `gorm:"column:idx_desc;type:mediumtext;not null;comment:简介"`                         // 简介
	ReplyCount   uint32                `gorm:"column:idx_replies;type:mediumint(8) unsigned;not null;comment:回复数"`          // 回复数
	SubjectCount uint32                `gorm:"column:idx_subject_total;type:mediumint(8) unsigned;not null;comment:内含条目总数"` // 内含条目总数
	CollectCount uint32                `gorm:"column:idx_collects;type:mediumint(8);not null;comment:收藏数"`                  // 收藏数
	Stats        string                `gorm:"column:idx_stats;type:mediumtext;not null"`
	CreatedTime  int32                 `gorm:"column:idx_dateline;type:int(10);not null;comment:创建时间"` // 创建时间
	UpdatedTime  uint32                `gorm:"column:idx_lasttouch;type:int(10) unsigned;not null"`
	CreatorID    uint32                `gorm:"column:idx_uid;type:mediumint(8);not null;comment:创建人UID"` // 创建人UID
	Deleted      soft_delete.DeletedAt `gorm:"column:idx_ban;type:tinyint(1) unsigned;not null;softDelete:flag"`
}

// TableName Index's table name
func (*Index) TableName() string {
	return TableNameIndex
}
