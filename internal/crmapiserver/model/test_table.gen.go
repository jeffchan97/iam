// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameTestTable = "test_table"

// TestTable mapped from table <test_table>
type TestTable struct {
	ID   int64  `gorm:"column:id;type:int;not null" json:"id"`
	Name string `gorm:"column:name;type:varchar(50);not null" json:"name"`
}

// TableName TestTable's table name
func (*TestTable) TableName() string {
	return TableNameTestTable
}
