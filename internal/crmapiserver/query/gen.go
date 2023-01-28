// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q                         = new(Query)
	CrmApplet                 *crmApplet
	TestTable                 *testTable
	UctAdmin                  *uctAdmin
	UctCrmErpMap              *uctCrmErpMap
	UctHuke                   *uctHuke
	UctModifyOrderAudit       *uctModifyOrderAudit
	UctWasteCustomer          *uctWasteCustomer
	UctWasteCustomerFactory   *uctWasteCustomerFactory
	UctWasteCustomerGathering *uctWasteCustomerGathering
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	CrmApplet = &Q.CrmApplet
	TestTable = &Q.TestTable
	UctAdmin = &Q.UctAdmin
	UctCrmErpMap = &Q.UctCrmErpMap
	UctHuke = &Q.UctHuke
	UctModifyOrderAudit = &Q.UctModifyOrderAudit
	UctWasteCustomer = &Q.UctWasteCustomer
	UctWasteCustomerFactory = &Q.UctWasteCustomerFactory
	UctWasteCustomerGathering = &Q.UctWasteCustomerGathering
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                        db,
		CrmApplet:                 newCrmApplet(db, opts...),
		TestTable:                 newTestTable(db, opts...),
		UctAdmin:                  newUctAdmin(db, opts...),
		UctCrmErpMap:              newUctCrmErpMap(db, opts...),
		UctHuke:                   newUctHuke(db, opts...),
		UctModifyOrderAudit:       newUctModifyOrderAudit(db, opts...),
		UctWasteCustomer:          newUctWasteCustomer(db, opts...),
		UctWasteCustomerFactory:   newUctWasteCustomerFactory(db, opts...),
		UctWasteCustomerGathering: newUctWasteCustomerGathering(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	CrmApplet                 crmApplet
	TestTable                 testTable
	UctAdmin                  uctAdmin
	UctCrmErpMap              uctCrmErpMap
	UctHuke                   uctHuke
	UctModifyOrderAudit       uctModifyOrderAudit
	UctWasteCustomer          uctWasteCustomer
	UctWasteCustomerFactory   uctWasteCustomerFactory
	UctWasteCustomerGathering uctWasteCustomerGathering
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                        db,
		CrmApplet:                 q.CrmApplet.clone(db),
		TestTable:                 q.TestTable.clone(db),
		UctAdmin:                  q.UctAdmin.clone(db),
		UctCrmErpMap:              q.UctCrmErpMap.clone(db),
		UctHuke:                   q.UctHuke.clone(db),
		UctModifyOrderAudit:       q.UctModifyOrderAudit.clone(db),
		UctWasteCustomer:          q.UctWasteCustomer.clone(db),
		UctWasteCustomerFactory:   q.UctWasteCustomerFactory.clone(db),
		UctWasteCustomerGathering: q.UctWasteCustomerGathering.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.clone(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.clone(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:                        db,
		CrmApplet:                 q.CrmApplet.replaceDB(db),
		TestTable:                 q.TestTable.replaceDB(db),
		UctAdmin:                  q.UctAdmin.replaceDB(db),
		UctCrmErpMap:              q.UctCrmErpMap.replaceDB(db),
		UctHuke:                   q.UctHuke.replaceDB(db),
		UctModifyOrderAudit:       q.UctModifyOrderAudit.replaceDB(db),
		UctWasteCustomer:          q.UctWasteCustomer.replaceDB(db),
		UctWasteCustomerFactory:   q.UctWasteCustomerFactory.replaceDB(db),
		UctWasteCustomerGathering: q.UctWasteCustomerGathering.replaceDB(db),
	}
}

type queryCtx struct {
	CrmApplet                 ICrmAppletDo
	TestTable                 ITestTableDo
	UctAdmin                  IUctAdminDo
	UctCrmErpMap              IUctCrmErpMapDo
	UctHuke                   IUctHukeDo
	UctModifyOrderAudit       IUctModifyOrderAuditDo
	UctWasteCustomer          IUctWasteCustomerDo
	UctWasteCustomerFactory   IUctWasteCustomerFactoryDo
	UctWasteCustomerGathering IUctWasteCustomerGatheringDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		CrmApplet:                 q.CrmApplet.WithContext(ctx),
		TestTable:                 q.TestTable.WithContext(ctx),
		UctAdmin:                  q.UctAdmin.WithContext(ctx),
		UctCrmErpMap:              q.UctCrmErpMap.WithContext(ctx),
		UctHuke:                   q.UctHuke.WithContext(ctx),
		UctModifyOrderAudit:       q.UctModifyOrderAudit.WithContext(ctx),
		UctWasteCustomer:          q.UctWasteCustomer.WithContext(ctx),
		UctWasteCustomerFactory:   q.UctWasteCustomerFactory.WithContext(ctx),
		UctWasteCustomerGathering: q.UctWasteCustomerGathering.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	return &QueryTx{q.clone(q.db.Begin(opts...))}
}

type QueryTx struct{ *Query }

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
