// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/marmotedu/iam/internal/crmapiserver/model"
)

func newUctWasteCustomerFactory(db *gorm.DB, opts ...gen.DOOption) uctWasteCustomerFactory {
	_uctWasteCustomerFactory := uctWasteCustomerFactory{}

	_uctWasteCustomerFactory.uctWasteCustomerFactoryDo.UseDB(db, opts...)
	_uctWasteCustomerFactory.uctWasteCustomerFactoryDo.UseModel(&model.UctWasteCustomerFactory{})

	tableName := _uctWasteCustomerFactory.uctWasteCustomerFactoryDo.TableName()
	_uctWasteCustomerFactory.ALL = field.NewAsterisk(tableName)
	_uctWasteCustomerFactory.ID = field.NewInt64(tableName, "id")
	_uctWasteCustomerFactory.CustomerID = field.NewInt64(tableName, "customer_id")
	_uctWasteCustomerFactory.Province = field.NewString(tableName, "province")
	_uctWasteCustomerFactory.City = field.NewString(tableName, "city")
	_uctWasteCustomerFactory.Area = field.NewString(tableName, "area")
	_uctWasteCustomerFactory.DetailAddress = field.NewString(tableName, "detail_address")
	_uctWasteCustomerFactory.Linkman = field.NewString(tableName, "linkman")
	_uctWasteCustomerFactory.Job = field.NewString(tableName, "job")
	_uctWasteCustomerFactory.Mobile = field.NewString(tableName, "mobile")
	_uctWasteCustomerFactory.Email = field.NewString(tableName, "email")
	_uctWasteCustomerFactory.Adcode = field.NewInt64(tableName, "adcode")
	_uctWasteCustomerFactory.Position = field.NewString(tableName, "position")
	_uctWasteCustomerFactory.Createtime = field.NewInt64(tableName, "createtime")
	_uctWasteCustomerFactory.Updatetime = field.NewInt64(tableName, "updatetime")

	_uctWasteCustomerFactory.fillFieldMap()

	return _uctWasteCustomerFactory
}

type uctWasteCustomerFactory struct {
	uctWasteCustomerFactoryDo uctWasteCustomerFactoryDo

	ALL           field.Asterisk
	ID            field.Int64  // ID
	CustomerID    field.Int64  // 客户ID
	Province      field.String // 省份
	City          field.String // 城市
	Area          field.String // 区域
	DetailAddress field.String // 详细地址
	Linkman       field.String // 联系人
	Job           field.String // 职务
	Mobile        field.String // 联系电话
	Email         field.String // 邮箱
	Adcode        field.Int64  // 行政区域编码
	Position      field.String // 坐标点
	Createtime    field.Int64  // 创建时间
	Updatetime    field.Int64  // 更新时间

	fieldMap map[string]field.Expr
}

func (u uctWasteCustomerFactory) Table(newTableName string) *uctWasteCustomerFactory {
	u.uctWasteCustomerFactoryDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u uctWasteCustomerFactory) As(alias string) *uctWasteCustomerFactory {
	u.uctWasteCustomerFactoryDo.DO = *(u.uctWasteCustomerFactoryDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *uctWasteCustomerFactory) updateTableName(table string) *uctWasteCustomerFactory {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt64(table, "id")
	u.CustomerID = field.NewInt64(table, "customer_id")
	u.Province = field.NewString(table, "province")
	u.City = field.NewString(table, "city")
	u.Area = field.NewString(table, "area")
	u.DetailAddress = field.NewString(table, "detail_address")
	u.Linkman = field.NewString(table, "linkman")
	u.Job = field.NewString(table, "job")
	u.Mobile = field.NewString(table, "mobile")
	u.Email = field.NewString(table, "email")
	u.Adcode = field.NewInt64(table, "adcode")
	u.Position = field.NewString(table, "position")
	u.Createtime = field.NewInt64(table, "createtime")
	u.Updatetime = field.NewInt64(table, "updatetime")

	u.fillFieldMap()

	return u
}

func (u *uctWasteCustomerFactory) WithContext(ctx context.Context) IUctWasteCustomerFactoryDo {
	return u.uctWasteCustomerFactoryDo.WithContext(ctx)
}

func (u uctWasteCustomerFactory) TableName() string { return u.uctWasteCustomerFactoryDo.TableName() }

func (u uctWasteCustomerFactory) Alias() string { return u.uctWasteCustomerFactoryDo.Alias() }

func (u *uctWasteCustomerFactory) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *uctWasteCustomerFactory) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 14)
	u.fieldMap["id"] = u.ID
	u.fieldMap["customer_id"] = u.CustomerID
	u.fieldMap["province"] = u.Province
	u.fieldMap["city"] = u.City
	u.fieldMap["area"] = u.Area
	u.fieldMap["detail_address"] = u.DetailAddress
	u.fieldMap["linkman"] = u.Linkman
	u.fieldMap["job"] = u.Job
	u.fieldMap["mobile"] = u.Mobile
	u.fieldMap["email"] = u.Email
	u.fieldMap["adcode"] = u.Adcode
	u.fieldMap["position"] = u.Position
	u.fieldMap["createtime"] = u.Createtime
	u.fieldMap["updatetime"] = u.Updatetime
}

func (u uctWasteCustomerFactory) clone(db *gorm.DB) uctWasteCustomerFactory {
	u.uctWasteCustomerFactoryDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u uctWasteCustomerFactory) replaceDB(db *gorm.DB) uctWasteCustomerFactory {
	u.uctWasteCustomerFactoryDo.ReplaceDB(db)
	return u
}

type uctWasteCustomerFactoryDo struct{ gen.DO }

type IUctWasteCustomerFactoryDo interface {
	gen.SubQuery
	Debug() IUctWasteCustomerFactoryDo
	WithContext(ctx context.Context) IUctWasteCustomerFactoryDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUctWasteCustomerFactoryDo
	WriteDB() IUctWasteCustomerFactoryDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUctWasteCustomerFactoryDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUctWasteCustomerFactoryDo
	Not(conds ...gen.Condition) IUctWasteCustomerFactoryDo
	Or(conds ...gen.Condition) IUctWasteCustomerFactoryDo
	Select(conds ...field.Expr) IUctWasteCustomerFactoryDo
	Where(conds ...gen.Condition) IUctWasteCustomerFactoryDo
	Order(conds ...field.Expr) IUctWasteCustomerFactoryDo
	Distinct(cols ...field.Expr) IUctWasteCustomerFactoryDo
	Omit(cols ...field.Expr) IUctWasteCustomerFactoryDo
	Join(table schema.Tabler, on ...field.Expr) IUctWasteCustomerFactoryDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUctWasteCustomerFactoryDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUctWasteCustomerFactoryDo
	Group(cols ...field.Expr) IUctWasteCustomerFactoryDo
	Having(conds ...gen.Condition) IUctWasteCustomerFactoryDo
	Limit(limit int) IUctWasteCustomerFactoryDo
	Offset(offset int) IUctWasteCustomerFactoryDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUctWasteCustomerFactoryDo
	Unscoped() IUctWasteCustomerFactoryDo
	Create(values ...*model.UctWasteCustomerFactory) error
	CreateInBatches(values []*model.UctWasteCustomerFactory, batchSize int) error
	Save(values ...*model.UctWasteCustomerFactory) error
	First() (*model.UctWasteCustomerFactory, error)
	Take() (*model.UctWasteCustomerFactory, error)
	Last() (*model.UctWasteCustomerFactory, error)
	Find() ([]*model.UctWasteCustomerFactory, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UctWasteCustomerFactory, err error)
	FindInBatches(result *[]*model.UctWasteCustomerFactory, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UctWasteCustomerFactory) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUctWasteCustomerFactoryDo
	Assign(attrs ...field.AssignExpr) IUctWasteCustomerFactoryDo
	Joins(fields ...field.RelationField) IUctWasteCustomerFactoryDo
	Preload(fields ...field.RelationField) IUctWasteCustomerFactoryDo
	FirstOrInit() (*model.UctWasteCustomerFactory, error)
	FirstOrCreate() (*model.UctWasteCustomerFactory, error)
	FindByPage(offset int, limit int) (result []*model.UctWasteCustomerFactory, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUctWasteCustomerFactoryDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u uctWasteCustomerFactoryDo) Debug() IUctWasteCustomerFactoryDo {
	return u.withDO(u.DO.Debug())
}

func (u uctWasteCustomerFactoryDo) WithContext(ctx context.Context) IUctWasteCustomerFactoryDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u uctWasteCustomerFactoryDo) ReadDB() IUctWasteCustomerFactoryDo {
	return u.Clauses(dbresolver.Read)
}

func (u uctWasteCustomerFactoryDo) WriteDB() IUctWasteCustomerFactoryDo {
	return u.Clauses(dbresolver.Write)
}

func (u uctWasteCustomerFactoryDo) Session(config *gorm.Session) IUctWasteCustomerFactoryDo {
	return u.withDO(u.DO.Session(config))
}

func (u uctWasteCustomerFactoryDo) Clauses(conds ...clause.Expression) IUctWasteCustomerFactoryDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u uctWasteCustomerFactoryDo) Returning(value interface{}, columns ...string) IUctWasteCustomerFactoryDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u uctWasteCustomerFactoryDo) Not(conds ...gen.Condition) IUctWasteCustomerFactoryDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u uctWasteCustomerFactoryDo) Or(conds ...gen.Condition) IUctWasteCustomerFactoryDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u uctWasteCustomerFactoryDo) Select(conds ...field.Expr) IUctWasteCustomerFactoryDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u uctWasteCustomerFactoryDo) Where(conds ...gen.Condition) IUctWasteCustomerFactoryDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u uctWasteCustomerFactoryDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IUctWasteCustomerFactoryDo {
	return u.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (u uctWasteCustomerFactoryDo) Order(conds ...field.Expr) IUctWasteCustomerFactoryDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u uctWasteCustomerFactoryDo) Distinct(cols ...field.Expr) IUctWasteCustomerFactoryDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u uctWasteCustomerFactoryDo) Omit(cols ...field.Expr) IUctWasteCustomerFactoryDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u uctWasteCustomerFactoryDo) Join(table schema.Tabler, on ...field.Expr) IUctWasteCustomerFactoryDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u uctWasteCustomerFactoryDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUctWasteCustomerFactoryDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u uctWasteCustomerFactoryDo) RightJoin(table schema.Tabler, on ...field.Expr) IUctWasteCustomerFactoryDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u uctWasteCustomerFactoryDo) Group(cols ...field.Expr) IUctWasteCustomerFactoryDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u uctWasteCustomerFactoryDo) Having(conds ...gen.Condition) IUctWasteCustomerFactoryDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u uctWasteCustomerFactoryDo) Limit(limit int) IUctWasteCustomerFactoryDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u uctWasteCustomerFactoryDo) Offset(offset int) IUctWasteCustomerFactoryDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u uctWasteCustomerFactoryDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUctWasteCustomerFactoryDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u uctWasteCustomerFactoryDo) Unscoped() IUctWasteCustomerFactoryDo {
	return u.withDO(u.DO.Unscoped())
}

func (u uctWasteCustomerFactoryDo) Create(values ...*model.UctWasteCustomerFactory) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u uctWasteCustomerFactoryDo) CreateInBatches(values []*model.UctWasteCustomerFactory, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u uctWasteCustomerFactoryDo) Save(values ...*model.UctWasteCustomerFactory) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u uctWasteCustomerFactoryDo) First() (*model.UctWasteCustomerFactory, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UctWasteCustomerFactory), nil
	}
}

func (u uctWasteCustomerFactoryDo) Take() (*model.UctWasteCustomerFactory, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UctWasteCustomerFactory), nil
	}
}

func (u uctWasteCustomerFactoryDo) Last() (*model.UctWasteCustomerFactory, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UctWasteCustomerFactory), nil
	}
}

func (u uctWasteCustomerFactoryDo) Find() ([]*model.UctWasteCustomerFactory, error) {
	result, err := u.DO.Find()
	return result.([]*model.UctWasteCustomerFactory), err
}

func (u uctWasteCustomerFactoryDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UctWasteCustomerFactory, err error) {
	buf := make([]*model.UctWasteCustomerFactory, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u uctWasteCustomerFactoryDo) FindInBatches(result *[]*model.UctWasteCustomerFactory, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u uctWasteCustomerFactoryDo) Attrs(attrs ...field.AssignExpr) IUctWasteCustomerFactoryDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u uctWasteCustomerFactoryDo) Assign(attrs ...field.AssignExpr) IUctWasteCustomerFactoryDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u uctWasteCustomerFactoryDo) Joins(fields ...field.RelationField) IUctWasteCustomerFactoryDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u uctWasteCustomerFactoryDo) Preload(fields ...field.RelationField) IUctWasteCustomerFactoryDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u uctWasteCustomerFactoryDo) FirstOrInit() (*model.UctWasteCustomerFactory, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UctWasteCustomerFactory), nil
	}
}

func (u uctWasteCustomerFactoryDo) FirstOrCreate() (*model.UctWasteCustomerFactory, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UctWasteCustomerFactory), nil
	}
}

func (u uctWasteCustomerFactoryDo) FindByPage(offset int, limit int) (result []*model.UctWasteCustomerFactory, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u uctWasteCustomerFactoryDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u uctWasteCustomerFactoryDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u uctWasteCustomerFactoryDo) Delete(models ...*model.UctWasteCustomerFactory) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *uctWasteCustomerFactoryDo) withDO(do gen.Dao) *uctWasteCustomerFactoryDo {
	u.DO = *do.(*gen.DO)
	return u
}