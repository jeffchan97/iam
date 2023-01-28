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

func newCrmApplet(db *gorm.DB, opts ...gen.DOOption) crmApplet {
	_crmApplet := crmApplet{}

	_crmApplet.crmAppletDo.UseDB(db, opts...)
	_crmApplet.crmAppletDo.UseModel(&model.CrmApplet{})

	tableName := _crmApplet.crmAppletDo.TableName()
	_crmApplet.ALL = field.NewAsterisk(tableName)
	_crmApplet.Name = field.NewString(tableName, "name")
	_crmApplet.Type = field.NewString(tableName, "type")
	_crmApplet.CrmType = field.NewString(tableName, "crm_type")
	_crmApplet.Appkey = field.NewString(tableName, "appkey")
	_crmApplet.Appsecret = field.NewString(tableName, "appsecret")
	_crmApplet.Cid = field.NewString(tableName, "cid")
	_crmApplet.Status = field.NewString(tableName, "status")
	_crmApplet.Createtime = field.NewTime(tableName, "createtime")
	_crmApplet.Updatetime = field.NewTime(tableName, "updatetime")

	_crmApplet.fillFieldMap()

	return _crmApplet
}

type crmApplet struct {
	crmAppletDo crmAppletDo

	ALL        field.Asterisk
	Name       field.String
	Type       field.String
	CrmType    field.String
	Appkey     field.String
	Appsecret  field.String
	Cid        field.String
	Status     field.String // enable 启用  disable 禁用
	Createtime field.Time   // 创建时间
	Updatetime field.Time   // 修改时间

	fieldMap map[string]field.Expr
}

func (c crmApplet) Table(newTableName string) *crmApplet {
	c.crmAppletDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c crmApplet) As(alias string) *crmApplet {
	c.crmAppletDo.DO = *(c.crmAppletDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *crmApplet) updateTableName(table string) *crmApplet {
	c.ALL = field.NewAsterisk(table)
	c.Name = field.NewString(table, "name")
	c.Type = field.NewString(table, "type")
	c.CrmType = field.NewString(table, "crm_type")
	c.Appkey = field.NewString(table, "appkey")
	c.Appsecret = field.NewString(table, "appsecret")
	c.Cid = field.NewString(table, "cid")
	c.Status = field.NewString(table, "status")
	c.Createtime = field.NewTime(table, "createtime")
	c.Updatetime = field.NewTime(table, "updatetime")

	c.fillFieldMap()

	return c
}

func (c *crmApplet) WithContext(ctx context.Context) ICrmAppletDo {
	return c.crmAppletDo.WithContext(ctx)
}

func (c crmApplet) TableName() string { return c.crmAppletDo.TableName() }

func (c crmApplet) Alias() string { return c.crmAppletDo.Alias() }

func (c *crmApplet) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *crmApplet) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 9)
	c.fieldMap["name"] = c.Name
	c.fieldMap["type"] = c.Type
	c.fieldMap["crm_type"] = c.CrmType
	c.fieldMap["appkey"] = c.Appkey
	c.fieldMap["appsecret"] = c.Appsecret
	c.fieldMap["cid"] = c.Cid
	c.fieldMap["status"] = c.Status
	c.fieldMap["createtime"] = c.Createtime
	c.fieldMap["updatetime"] = c.Updatetime
}

func (c crmApplet) clone(db *gorm.DB) crmApplet {
	c.crmAppletDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c crmApplet) replaceDB(db *gorm.DB) crmApplet {
	c.crmAppletDo.ReplaceDB(db)
	return c
}

type crmAppletDo struct{ gen.DO }

type ICrmAppletDo interface {
	gen.SubQuery
	Debug() ICrmAppletDo
	WithContext(ctx context.Context) ICrmAppletDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICrmAppletDo
	WriteDB() ICrmAppletDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICrmAppletDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICrmAppletDo
	Not(conds ...gen.Condition) ICrmAppletDo
	Or(conds ...gen.Condition) ICrmAppletDo
	Select(conds ...field.Expr) ICrmAppletDo
	Where(conds ...gen.Condition) ICrmAppletDo
	Order(conds ...field.Expr) ICrmAppletDo
	Distinct(cols ...field.Expr) ICrmAppletDo
	Omit(cols ...field.Expr) ICrmAppletDo
	Join(table schema.Tabler, on ...field.Expr) ICrmAppletDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICrmAppletDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICrmAppletDo
	Group(cols ...field.Expr) ICrmAppletDo
	Having(conds ...gen.Condition) ICrmAppletDo
	Limit(limit int) ICrmAppletDo
	Offset(offset int) ICrmAppletDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICrmAppletDo
	Unscoped() ICrmAppletDo
	Create(values ...*model.CrmApplet) error
	CreateInBatches(values []*model.CrmApplet, batchSize int) error
	Save(values ...*model.CrmApplet) error
	First() (*model.CrmApplet, error)
	Take() (*model.CrmApplet, error)
	Last() (*model.CrmApplet, error)
	Find() ([]*model.CrmApplet, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CrmApplet, err error)
	FindInBatches(result *[]*model.CrmApplet, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.CrmApplet) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICrmAppletDo
	Assign(attrs ...field.AssignExpr) ICrmAppletDo
	Joins(fields ...field.RelationField) ICrmAppletDo
	Preload(fields ...field.RelationField) ICrmAppletDo
	FirstOrInit() (*model.CrmApplet, error)
	FirstOrCreate() (*model.CrmApplet, error)
	FindByPage(offset int, limit int) (result []*model.CrmApplet, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICrmAppletDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c crmAppletDo) Debug() ICrmAppletDo {
	return c.withDO(c.DO.Debug())
}

func (c crmAppletDo) WithContext(ctx context.Context) ICrmAppletDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c crmAppletDo) ReadDB() ICrmAppletDo {
	return c.Clauses(dbresolver.Read)
}

func (c crmAppletDo) WriteDB() ICrmAppletDo {
	return c.Clauses(dbresolver.Write)
}

func (c crmAppletDo) Session(config *gorm.Session) ICrmAppletDo {
	return c.withDO(c.DO.Session(config))
}

func (c crmAppletDo) Clauses(conds ...clause.Expression) ICrmAppletDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c crmAppletDo) Returning(value interface{}, columns ...string) ICrmAppletDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c crmAppletDo) Not(conds ...gen.Condition) ICrmAppletDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c crmAppletDo) Or(conds ...gen.Condition) ICrmAppletDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c crmAppletDo) Select(conds ...field.Expr) ICrmAppletDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c crmAppletDo) Where(conds ...gen.Condition) ICrmAppletDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c crmAppletDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ICrmAppletDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c crmAppletDo) Order(conds ...field.Expr) ICrmAppletDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c crmAppletDo) Distinct(cols ...field.Expr) ICrmAppletDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c crmAppletDo) Omit(cols ...field.Expr) ICrmAppletDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c crmAppletDo) Join(table schema.Tabler, on ...field.Expr) ICrmAppletDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c crmAppletDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICrmAppletDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c crmAppletDo) RightJoin(table schema.Tabler, on ...field.Expr) ICrmAppletDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c crmAppletDo) Group(cols ...field.Expr) ICrmAppletDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c crmAppletDo) Having(conds ...gen.Condition) ICrmAppletDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c crmAppletDo) Limit(limit int) ICrmAppletDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c crmAppletDo) Offset(offset int) ICrmAppletDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c crmAppletDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICrmAppletDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c crmAppletDo) Unscoped() ICrmAppletDo {
	return c.withDO(c.DO.Unscoped())
}

func (c crmAppletDo) Create(values ...*model.CrmApplet) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c crmAppletDo) CreateInBatches(values []*model.CrmApplet, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c crmAppletDo) Save(values ...*model.CrmApplet) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c crmAppletDo) First() (*model.CrmApplet, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CrmApplet), nil
	}
}

func (c crmAppletDo) Take() (*model.CrmApplet, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CrmApplet), nil
	}
}

func (c crmAppletDo) Last() (*model.CrmApplet, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CrmApplet), nil
	}
}

func (c crmAppletDo) Find() ([]*model.CrmApplet, error) {
	result, err := c.DO.Find()
	return result.([]*model.CrmApplet), err
}

func (c crmAppletDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CrmApplet, err error) {
	buf := make([]*model.CrmApplet, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c crmAppletDo) FindInBatches(result *[]*model.CrmApplet, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c crmAppletDo) Attrs(attrs ...field.AssignExpr) ICrmAppletDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c crmAppletDo) Assign(attrs ...field.AssignExpr) ICrmAppletDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c crmAppletDo) Joins(fields ...field.RelationField) ICrmAppletDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c crmAppletDo) Preload(fields ...field.RelationField) ICrmAppletDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c crmAppletDo) FirstOrInit() (*model.CrmApplet, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CrmApplet), nil
	}
}

func (c crmAppletDo) FirstOrCreate() (*model.CrmApplet, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CrmApplet), nil
	}
}

func (c crmAppletDo) FindByPage(offset int, limit int) (result []*model.CrmApplet, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c crmAppletDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c crmAppletDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c crmAppletDo) Delete(models ...*model.CrmApplet) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *crmAppletDo) withDO(do gen.Dao) *crmAppletDo {
	c.DO = *do.(*gen.DO)
	return c
}
