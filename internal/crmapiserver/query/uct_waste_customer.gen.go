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

func newUctWasteCustomer(db *gorm.DB, opts ...gen.DOOption) uctWasteCustomer {
	_uctWasteCustomer := uctWasteCustomer{}

	_uctWasteCustomer.uctWasteCustomerDo.UseDB(db, opts...)
	_uctWasteCustomer.uctWasteCustomerDo.UseModel(&model.UctWasteCustomer{})

	tableName := _uctWasteCustomer.uctWasteCustomerDo.TableName()
	_uctWasteCustomer.ALL = field.NewAsterisk(tableName)
	_uctWasteCustomer.ID = field.NewInt64(tableName, "id")
	_uctWasteCustomer.CusAccid = field.NewString(tableName, "cus_accid")
	_uctWasteCustomer.Crmid = field.NewString(tableName, "crmid")
	_uctWasteCustomer.AdminID = field.NewString(tableName, "admin_id")
	_uctWasteCustomer.BranchID = field.NewInt64(tableName, "branch_id")
	_uctWasteCustomer.Internal = field.NewInt64(tableName, "internal")
	_uctWasteCustomer.GreenCoin = field.NewInt64(tableName, "green_coin")
	_uctWasteCustomer.CustomerType = field.NewString(tableName, "customer_type")
	_uctWasteCustomer.CustomerCode = field.NewString(tableName, "customer_code")
	_uctWasteCustomer.Name = field.NewString(tableName, "name")
	_uctWasteCustomer.PrintTitleID = field.NewInt64(tableName, "print_title_id")
	_uctWasteCustomer.NameEn = field.NewString(tableName, "name_en")
	_uctWasteCustomer.AllotID = field.NewInt64(tableName, "allot_id")
	_uctWasteCustomer.ManagerID = field.NewInt64(tableName, "manager_id")
	_uctWasteCustomer.ServerID = field.NewInt64(tableName, "server_id")
	_uctWasteCustomer.SellerIncharge = field.NewInt64(tableName, "seller_incharge")
	_uctWasteCustomer.ServiceDepartment = field.NewInt64(tableName, "service_department")
	_uctWasteCustomer.Incharge = field.NewString(tableName, "incharge")
	_uctWasteCustomer.JuridicalPerson = field.NewString(tableName, "juridical_person")
	_uctWasteCustomer.ClassA = field.NewString(tableName, "class_a")
	_uctWasteCustomer.ClassB = field.NewString(tableName, "class_b")
	_uctWasteCustomer.TradeType = field.NewString(tableName, "trade_type")
	_uctWasteCustomer.Mobile = field.NewString(tableName, "mobile")
	_uctWasteCustomer.WasteType = field.NewString(tableName, "waste_type")
	_uctWasteCustomer.Level = field.NewInt64(tableName, "level")
	_uctWasteCustomer.CompanyAddress = field.NewString(tableName, "company_address")
	_uctWasteCustomer.CompanyPosition = field.NewString(tableName, "company_position")
	_uctWasteCustomer.Industry = field.NewInt64(tableName, "industry")
	_uctWasteCustomer.Website = field.NewString(tableName, "website")
	_uctWasteCustomer.SalesArea = field.NewString(tableName, "sales_area")
	_uctWasteCustomer.AnnualWaste = field.NewFloat32(tableName, "annual_waste")
	_uctWasteCustomer.Detail = field.NewString(tableName, "detail")
	_uctWasteCustomer.TaxCert = field.NewString(tableName, "tax_cert")
	_uctWasteCustomer.TradingCert = field.NewString(tableName, "trading_cert")
	_uctWasteCustomer.CompanyNature = field.NewInt64(tableName, "company_nature")
	_uctWasteCustomer.TaxType = field.NewInt64(tableName, "tax_type")
	_uctWasteCustomer.SettleWay = field.NewInt64(tableName, "settle_way")
	_uctWasteCustomer.BackPercent = field.NewInt64(tableName, "back_percent")
	_uctWasteCustomer.MonthReceivable = field.NewInt64(tableName, "month_receivable")
	_uctWasteCustomer.StatementsJSON = field.NewString(tableName, "statements_json")
	_uctWasteCustomer.IsHideZero = field.NewInt64(tableName, "is_hide_zero")
	_uctWasteCustomer.IsSeparate = field.NewInt64(tableName, "is_separate")
	_uctWasteCustomer.BlackState = field.NewInt64(tableName, "black_state")
	_uctWasteCustomer.State = field.NewString(tableName, "state")
	_uctWasteCustomer.Createtime = field.NewInt64(tableName, "createtime")
	_uctWasteCustomer.Updatetime = field.NewInt64(tableName, "updatetime")

	_uctWasteCustomer.fillFieldMap()

	return _uctWasteCustomer
}

type uctWasteCustomer struct {
	uctWasteCustomerDo uctWasteCustomerDo

	ALL               field.Asterisk
	ID                field.Int64   // ID
	CusAccid          field.String  // 互客系统id
	Crmid             field.String  // crmid
	AdminID           field.String  // 客户账号ID
	BranchID          field.Int64   // 分部ID
	Internal          field.Int64   // 内部客户 0否 1是
	GreenCoin         field.Int64   // 绿币
	CustomerType      field.String  // 客户类型:up=上游客户,down=下游客户
	CustomerCode      field.String  // 客户编码
	Name              field.String  // 企业全称
	PrintTitleID      field.Int64   // 客户单据打印id
	NameEn            field.String  // 企业全称(英文)
	AllotID           field.Int64   // 指定的分配人
	ManagerID         field.Int64   // 业务负责人ID
	ServerID          field.Int64   // 服务负责人ID
	SellerIncharge    field.Int64   // 销售负责人ID
	ServiceDepartment field.Int64   // 服务部门 1 客服 2企服
	Incharge          field.String  // 客户负责人
	JuridicalPerson   field.String  // 公司法人
	ClassA            field.String  // 客户类别a  A1 A2 A3 A4
	ClassB            field.String  // 客户类别b  B1 B2 B3 B4
	TradeType         field.String  // 行业类型   1模切、涂布类 = M 特钢类 = T 隔膜类 = G 车废类 = C 其他类 = Q
	Mobile            field.String  // 联系电话
	WasteType         field.String  // 废料种类
	Level             field.Int64   // 客户级别
	CompanyAddress    field.String  // 企业地址
	CompanyPosition   field.String  // 企业坐标
	Industry          field.Int64   // 行业类别
	Website           field.String  // 官网
	SalesArea         field.String  // 销售区域
	AnnualWaste       field.Float32 // 年度废料量
	Detail            field.String  // 详情描述
	TaxCert           field.String  // 税务登记证号
	TradingCert       field.String  // 营业执照号
	CompanyNature     field.Int64   // 企业性质
	TaxType           field.Int64   // 税票类型
	SettleWay         field.Int64   // 结算方式
	BackPercent       field.Int64   // 分成比例(%)
	MonthReceivable   field.Int64   // 月结算服务 1是 0不是
	StatementsJSON    field.String  // 客户的结算基础资料（json格式）
	IsHideZero        field.Int64   // 对账单是否屏蔽为零的项：0--不屏蔽     1--屏蔽
	IsSeparate        field.Int64   // 对账单是否收付款分离： 0--不分离    1--分离
	BlackState        field.Int64   // 重点关注客户
	State             field.String  // 状态:black=黑名单,consult=客户咨询,draft=草稿,enabled=已提交
	Createtime        field.Int64   // 创建时间
	Updatetime        field.Int64   // 更新时间

	fieldMap map[string]field.Expr
}

func (u uctWasteCustomer) Table(newTableName string) *uctWasteCustomer {
	u.uctWasteCustomerDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u uctWasteCustomer) As(alias string) *uctWasteCustomer {
	u.uctWasteCustomerDo.DO = *(u.uctWasteCustomerDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *uctWasteCustomer) updateTableName(table string) *uctWasteCustomer {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt64(table, "id")
	u.CusAccid = field.NewString(table, "cus_accid")
	u.Crmid = field.NewString(table, "crmid")
	u.AdminID = field.NewString(table, "admin_id")
	u.BranchID = field.NewInt64(table, "branch_id")
	u.Internal = field.NewInt64(table, "internal")
	u.GreenCoin = field.NewInt64(table, "green_coin")
	u.CustomerType = field.NewString(table, "customer_type")
	u.CustomerCode = field.NewString(table, "customer_code")
	u.Name = field.NewString(table, "name")
	u.PrintTitleID = field.NewInt64(table, "print_title_id")
	u.NameEn = field.NewString(table, "name_en")
	u.AllotID = field.NewInt64(table, "allot_id")
	u.ManagerID = field.NewInt64(table, "manager_id")
	u.ServerID = field.NewInt64(table, "server_id")
	u.SellerIncharge = field.NewInt64(table, "seller_incharge")
	u.ServiceDepartment = field.NewInt64(table, "service_department")
	u.Incharge = field.NewString(table, "incharge")
	u.JuridicalPerson = field.NewString(table, "juridical_person")
	u.ClassA = field.NewString(table, "class_a")
	u.ClassB = field.NewString(table, "class_b")
	u.TradeType = field.NewString(table, "trade_type")
	u.Mobile = field.NewString(table, "mobile")
	u.WasteType = field.NewString(table, "waste_type")
	u.Level = field.NewInt64(table, "level")
	u.CompanyAddress = field.NewString(table, "company_address")
	u.CompanyPosition = field.NewString(table, "company_position")
	u.Industry = field.NewInt64(table, "industry")
	u.Website = field.NewString(table, "website")
	u.SalesArea = field.NewString(table, "sales_area")
	u.AnnualWaste = field.NewFloat32(table, "annual_waste")
	u.Detail = field.NewString(table, "detail")
	u.TaxCert = field.NewString(table, "tax_cert")
	u.TradingCert = field.NewString(table, "trading_cert")
	u.CompanyNature = field.NewInt64(table, "company_nature")
	u.TaxType = field.NewInt64(table, "tax_type")
	u.SettleWay = field.NewInt64(table, "settle_way")
	u.BackPercent = field.NewInt64(table, "back_percent")
	u.MonthReceivable = field.NewInt64(table, "month_receivable")
	u.StatementsJSON = field.NewString(table, "statements_json")
	u.IsHideZero = field.NewInt64(table, "is_hide_zero")
	u.IsSeparate = field.NewInt64(table, "is_separate")
	u.BlackState = field.NewInt64(table, "black_state")
	u.State = field.NewString(table, "state")
	u.Createtime = field.NewInt64(table, "createtime")
	u.Updatetime = field.NewInt64(table, "updatetime")

	u.fillFieldMap()

	return u
}

func (u *uctWasteCustomer) WithContext(ctx context.Context) IUctWasteCustomerDo {
	return u.uctWasteCustomerDo.WithContext(ctx)
}

func (u uctWasteCustomer) TableName() string { return u.uctWasteCustomerDo.TableName() }

func (u uctWasteCustomer) Alias() string { return u.uctWasteCustomerDo.Alias() }

func (u *uctWasteCustomer) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *uctWasteCustomer) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 46)
	u.fieldMap["id"] = u.ID
	u.fieldMap["cus_accid"] = u.CusAccid
	u.fieldMap["crmid"] = u.Crmid
	u.fieldMap["admin_id"] = u.AdminID
	u.fieldMap["branch_id"] = u.BranchID
	u.fieldMap["internal"] = u.Internal
	u.fieldMap["green_coin"] = u.GreenCoin
	u.fieldMap["customer_type"] = u.CustomerType
	u.fieldMap["customer_code"] = u.CustomerCode
	u.fieldMap["name"] = u.Name
	u.fieldMap["print_title_id"] = u.PrintTitleID
	u.fieldMap["name_en"] = u.NameEn
	u.fieldMap["allot_id"] = u.AllotID
	u.fieldMap["manager_id"] = u.ManagerID
	u.fieldMap["server_id"] = u.ServerID
	u.fieldMap["seller_incharge"] = u.SellerIncharge
	u.fieldMap["service_department"] = u.ServiceDepartment
	u.fieldMap["incharge"] = u.Incharge
	u.fieldMap["juridical_person"] = u.JuridicalPerson
	u.fieldMap["class_a"] = u.ClassA
	u.fieldMap["class_b"] = u.ClassB
	u.fieldMap["trade_type"] = u.TradeType
	u.fieldMap["mobile"] = u.Mobile
	u.fieldMap["waste_type"] = u.WasteType
	u.fieldMap["level"] = u.Level
	u.fieldMap["company_address"] = u.CompanyAddress
	u.fieldMap["company_position"] = u.CompanyPosition
	u.fieldMap["industry"] = u.Industry
	u.fieldMap["website"] = u.Website
	u.fieldMap["sales_area"] = u.SalesArea
	u.fieldMap["annual_waste"] = u.AnnualWaste
	u.fieldMap["detail"] = u.Detail
	u.fieldMap["tax_cert"] = u.TaxCert
	u.fieldMap["trading_cert"] = u.TradingCert
	u.fieldMap["company_nature"] = u.CompanyNature
	u.fieldMap["tax_type"] = u.TaxType
	u.fieldMap["settle_way"] = u.SettleWay
	u.fieldMap["back_percent"] = u.BackPercent
	u.fieldMap["month_receivable"] = u.MonthReceivable
	u.fieldMap["statements_json"] = u.StatementsJSON
	u.fieldMap["is_hide_zero"] = u.IsHideZero
	u.fieldMap["is_separate"] = u.IsSeparate
	u.fieldMap["black_state"] = u.BlackState
	u.fieldMap["state"] = u.State
	u.fieldMap["createtime"] = u.Createtime
	u.fieldMap["updatetime"] = u.Updatetime
}

func (u uctWasteCustomer) clone(db *gorm.DB) uctWasteCustomer {
	u.uctWasteCustomerDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u uctWasteCustomer) replaceDB(db *gorm.DB) uctWasteCustomer {
	u.uctWasteCustomerDo.ReplaceDB(db)
	return u
}

type uctWasteCustomerDo struct{ gen.DO }

type IUctWasteCustomerDo interface {
	gen.SubQuery
	Debug() IUctWasteCustomerDo
	WithContext(ctx context.Context) IUctWasteCustomerDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUctWasteCustomerDo
	WriteDB() IUctWasteCustomerDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUctWasteCustomerDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUctWasteCustomerDo
	Not(conds ...gen.Condition) IUctWasteCustomerDo
	Or(conds ...gen.Condition) IUctWasteCustomerDo
	Select(conds ...field.Expr) IUctWasteCustomerDo
	Where(conds ...gen.Condition) IUctWasteCustomerDo
	Order(conds ...field.Expr) IUctWasteCustomerDo
	Distinct(cols ...field.Expr) IUctWasteCustomerDo
	Omit(cols ...field.Expr) IUctWasteCustomerDo
	Join(table schema.Tabler, on ...field.Expr) IUctWasteCustomerDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUctWasteCustomerDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUctWasteCustomerDo
	Group(cols ...field.Expr) IUctWasteCustomerDo
	Having(conds ...gen.Condition) IUctWasteCustomerDo
	Limit(limit int) IUctWasteCustomerDo
	Offset(offset int) IUctWasteCustomerDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUctWasteCustomerDo
	Unscoped() IUctWasteCustomerDo
	Create(values ...*model.UctWasteCustomer) error
	CreateInBatches(values []*model.UctWasteCustomer, batchSize int) error
	Save(values ...*model.UctWasteCustomer) error
	First() (*model.UctWasteCustomer, error)
	Take() (*model.UctWasteCustomer, error)
	Last() (*model.UctWasteCustomer, error)
	Find() ([]*model.UctWasteCustomer, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UctWasteCustomer, err error)
	FindInBatches(result *[]*model.UctWasteCustomer, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UctWasteCustomer) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUctWasteCustomerDo
	Assign(attrs ...field.AssignExpr) IUctWasteCustomerDo
	Joins(fields ...field.RelationField) IUctWasteCustomerDo
	Preload(fields ...field.RelationField) IUctWasteCustomerDo
	FirstOrInit() (*model.UctWasteCustomer, error)
	FirstOrCreate() (*model.UctWasteCustomer, error)
	FindByPage(offset int, limit int) (result []*model.UctWasteCustomer, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUctWasteCustomerDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u uctWasteCustomerDo) Debug() IUctWasteCustomerDo {
	return u.withDO(u.DO.Debug())
}

func (u uctWasteCustomerDo) WithContext(ctx context.Context) IUctWasteCustomerDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u uctWasteCustomerDo) ReadDB() IUctWasteCustomerDo {
	return u.Clauses(dbresolver.Read)
}

func (u uctWasteCustomerDo) WriteDB() IUctWasteCustomerDo {
	return u.Clauses(dbresolver.Write)
}

func (u uctWasteCustomerDo) Session(config *gorm.Session) IUctWasteCustomerDo {
	return u.withDO(u.DO.Session(config))
}

func (u uctWasteCustomerDo) Clauses(conds ...clause.Expression) IUctWasteCustomerDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u uctWasteCustomerDo) Returning(value interface{}, columns ...string) IUctWasteCustomerDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u uctWasteCustomerDo) Not(conds ...gen.Condition) IUctWasteCustomerDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u uctWasteCustomerDo) Or(conds ...gen.Condition) IUctWasteCustomerDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u uctWasteCustomerDo) Select(conds ...field.Expr) IUctWasteCustomerDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u uctWasteCustomerDo) Where(conds ...gen.Condition) IUctWasteCustomerDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u uctWasteCustomerDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IUctWasteCustomerDo {
	return u.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (u uctWasteCustomerDo) Order(conds ...field.Expr) IUctWasteCustomerDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u uctWasteCustomerDo) Distinct(cols ...field.Expr) IUctWasteCustomerDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u uctWasteCustomerDo) Omit(cols ...field.Expr) IUctWasteCustomerDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u uctWasteCustomerDo) Join(table schema.Tabler, on ...field.Expr) IUctWasteCustomerDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u uctWasteCustomerDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUctWasteCustomerDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u uctWasteCustomerDo) RightJoin(table schema.Tabler, on ...field.Expr) IUctWasteCustomerDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u uctWasteCustomerDo) Group(cols ...field.Expr) IUctWasteCustomerDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u uctWasteCustomerDo) Having(conds ...gen.Condition) IUctWasteCustomerDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u uctWasteCustomerDo) Limit(limit int) IUctWasteCustomerDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u uctWasteCustomerDo) Offset(offset int) IUctWasteCustomerDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u uctWasteCustomerDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUctWasteCustomerDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u uctWasteCustomerDo) Unscoped() IUctWasteCustomerDo {
	return u.withDO(u.DO.Unscoped())
}

func (u uctWasteCustomerDo) Create(values ...*model.UctWasteCustomer) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u uctWasteCustomerDo) CreateInBatches(values []*model.UctWasteCustomer, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u uctWasteCustomerDo) Save(values ...*model.UctWasteCustomer) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u uctWasteCustomerDo) First() (*model.UctWasteCustomer, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UctWasteCustomer), nil
	}
}

func (u uctWasteCustomerDo) Take() (*model.UctWasteCustomer, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UctWasteCustomer), nil
	}
}

func (u uctWasteCustomerDo) Last() (*model.UctWasteCustomer, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UctWasteCustomer), nil
	}
}

func (u uctWasteCustomerDo) Find() ([]*model.UctWasteCustomer, error) {
	result, err := u.DO.Find()
	return result.([]*model.UctWasteCustomer), err
}

func (u uctWasteCustomerDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UctWasteCustomer, err error) {
	buf := make([]*model.UctWasteCustomer, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u uctWasteCustomerDo) FindInBatches(result *[]*model.UctWasteCustomer, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u uctWasteCustomerDo) Attrs(attrs ...field.AssignExpr) IUctWasteCustomerDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u uctWasteCustomerDo) Assign(attrs ...field.AssignExpr) IUctWasteCustomerDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u uctWasteCustomerDo) Joins(fields ...field.RelationField) IUctWasteCustomerDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u uctWasteCustomerDo) Preload(fields ...field.RelationField) IUctWasteCustomerDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u uctWasteCustomerDo) FirstOrInit() (*model.UctWasteCustomer, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UctWasteCustomer), nil
	}
}

func (u uctWasteCustomerDo) FirstOrCreate() (*model.UctWasteCustomer, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UctWasteCustomer), nil
	}
}

func (u uctWasteCustomerDo) FindByPage(offset int, limit int) (result []*model.UctWasteCustomer, count int64, err error) {
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

func (u uctWasteCustomerDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u uctWasteCustomerDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u uctWasteCustomerDo) Delete(models ...*model.UctWasteCustomer) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *uctWasteCustomerDo) withDO(do gen.Dao) *uctWasteCustomerDo {
	u.DO = *do.(*gen.DO)
	return u
}
