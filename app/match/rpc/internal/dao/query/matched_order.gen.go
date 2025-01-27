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

	"github.com/luxun9527/gex/app/match/rpc/internal/dao/model"
)

func newMatchedOrder(db *gorm.DB, opts ...gen.DOOption) matchedOrder {
	_matchedOrder := matchedOrder{}

	_matchedOrder.matchedOrderDo.UseDB(db, opts...)
	_matchedOrder.matchedOrderDo.UseModel(&model.MatchedOrder{})

	tableName := _matchedOrder.matchedOrderDo.TableName()
	_matchedOrder.ALL = field.NewAsterisk(tableName)
	_matchedOrder.ID = field.NewInt64(tableName, "id")
	_matchedOrder.MatchID = field.NewString(tableName, "match_id")
	_matchedOrder.MatchSubID = field.NewString(tableName, "match_sub_id")
	_matchedOrder.SymbolID = field.NewInt32(tableName, "symbol_id")
	_matchedOrder.SymbolName = field.NewString(tableName, "symbol_name")
	_matchedOrder.TakerUserID = field.NewInt64(tableName, "taker_user_id")
	_matchedOrder.TakerOrderID = field.NewString(tableName, "taker_order_id")
	_matchedOrder.MakerOrderID = field.NewString(tableName, "maker_order_id")
	_matchedOrder.MakerUserID = field.NewInt64(tableName, "maker_user_id")
	_matchedOrder.TakerIsBuyer = field.NewInt32(tableName, "taker_is_buyer")
	_matchedOrder.Price = field.NewString(tableName, "price")
	_matchedOrder.Qty = field.NewString(tableName, "qty")
	_matchedOrder.Amount = field.NewString(tableName, "amount")
	_matchedOrder.MatchTime = field.NewInt64(tableName, "match_time")
	_matchedOrder.CreatedAt = field.NewInt64(tableName, "created_at")
	_matchedOrder.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_matchedOrder.fillFieldMap()

	return _matchedOrder
}

type matchedOrder struct {
	matchedOrderDo matchedOrderDo

	ALL          field.Asterisk
	ID           field.Int64  // 雪花算法id
	MatchID      field.String // 撮合id
	MatchSubID   field.String // 本次匹配的id，一次撮合会多次匹配
	SymbolID     field.Int32  // 交易对id
	SymbolName   field.String // 交易对名称
	TakerUserID  field.Int64  // taker用户id
	TakerOrderID field.String // taker订单id
	MakerOrderID field.String // maker订单id
	MakerUserID  field.Int64  // maker用户id
	TakerIsBuyer field.Int32  // taker是否是买单 1是 2否
	Price        field.String // 价格
	Qty          field.String // 数量(基础币)
	Amount       field.String // 金额（计价币）
	MatchTime    field.Int64  // 撮合时间
	CreatedAt    field.Int64  // 创建时间
	UpdatedAt    field.Int64  // 修改时间

	fieldMap map[string]field.Expr
}

func (m matchedOrder) Table(newTableName string) *matchedOrder {
	m.matchedOrderDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m matchedOrder) As(alias string) *matchedOrder {
	m.matchedOrderDo.DO = *(m.matchedOrderDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *matchedOrder) updateTableName(table string) *matchedOrder {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewInt64(table, "id")
	m.MatchID = field.NewString(table, "match_id")
	m.MatchSubID = field.NewString(table, "match_sub_id")
	m.SymbolID = field.NewInt32(table, "symbol_id")
	m.SymbolName = field.NewString(table, "symbol_name")
	m.TakerUserID = field.NewInt64(table, "taker_user_id")
	m.TakerOrderID = field.NewString(table, "taker_order_id")
	m.MakerOrderID = field.NewString(table, "maker_order_id")
	m.MakerUserID = field.NewInt64(table, "maker_user_id")
	m.TakerIsBuyer = field.NewInt32(table, "taker_is_buyer")
	m.Price = field.NewString(table, "price")
	m.Qty = field.NewString(table, "qty")
	m.Amount = field.NewString(table, "amount")
	m.MatchTime = field.NewInt64(table, "match_time")
	m.CreatedAt = field.NewInt64(table, "created_at")
	m.UpdatedAt = field.NewInt64(table, "updated_at")

	m.fillFieldMap()

	return m
}

func (m *matchedOrder) WithContext(ctx context.Context) *matchedOrderDo {
	return m.matchedOrderDo.WithContext(ctx)
}

func (m matchedOrder) TableName() string { return m.matchedOrderDo.TableName() }

func (m matchedOrder) Alias() string { return m.matchedOrderDo.Alias() }

func (m matchedOrder) Columns(cols ...field.Expr) gen.Columns {
	return m.matchedOrderDo.Columns(cols...)
}

func (m *matchedOrder) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *matchedOrder) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 16)
	m.fieldMap["id"] = m.ID
	m.fieldMap["match_id"] = m.MatchID
	m.fieldMap["match_sub_id"] = m.MatchSubID
	m.fieldMap["symbol_id"] = m.SymbolID
	m.fieldMap["symbol_name"] = m.SymbolName
	m.fieldMap["taker_user_id"] = m.TakerUserID
	m.fieldMap["taker_order_id"] = m.TakerOrderID
	m.fieldMap["maker_order_id"] = m.MakerOrderID
	m.fieldMap["maker_user_id"] = m.MakerUserID
	m.fieldMap["taker_is_buyer"] = m.TakerIsBuyer
	m.fieldMap["price"] = m.Price
	m.fieldMap["qty"] = m.Qty
	m.fieldMap["amount"] = m.Amount
	m.fieldMap["match_time"] = m.MatchTime
	m.fieldMap["created_at"] = m.CreatedAt
	m.fieldMap["updated_at"] = m.UpdatedAt
}

func (m matchedOrder) clone(db *gorm.DB) matchedOrder {
	m.matchedOrderDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m matchedOrder) replaceDB(db *gorm.DB) matchedOrder {
	m.matchedOrderDo.ReplaceDB(db)
	return m
}

type matchedOrderDo struct{ gen.DO }

func (m matchedOrderDo) Debug() *matchedOrderDo {
	return m.withDO(m.DO.Debug())
}

func (m matchedOrderDo) WithContext(ctx context.Context) *matchedOrderDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m matchedOrderDo) ReadDB() *matchedOrderDo {
	return m.Clauses(dbresolver.Read)
}

func (m matchedOrderDo) WriteDB() *matchedOrderDo {
	return m.Clauses(dbresolver.Write)
}

func (m matchedOrderDo) Session(config *gorm.Session) *matchedOrderDo {
	return m.withDO(m.DO.Session(config))
}

func (m matchedOrderDo) Clauses(conds ...clause.Expression) *matchedOrderDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m matchedOrderDo) Returning(value interface{}, columns ...string) *matchedOrderDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m matchedOrderDo) Not(conds ...gen.Condition) *matchedOrderDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m matchedOrderDo) Or(conds ...gen.Condition) *matchedOrderDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m matchedOrderDo) Select(conds ...field.Expr) *matchedOrderDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m matchedOrderDo) Where(conds ...gen.Condition) *matchedOrderDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m matchedOrderDo) Order(conds ...field.Expr) *matchedOrderDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m matchedOrderDo) Distinct(cols ...field.Expr) *matchedOrderDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m matchedOrderDo) Omit(cols ...field.Expr) *matchedOrderDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m matchedOrderDo) Join(table schema.Tabler, on ...field.Expr) *matchedOrderDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m matchedOrderDo) LeftJoin(table schema.Tabler, on ...field.Expr) *matchedOrderDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m matchedOrderDo) RightJoin(table schema.Tabler, on ...field.Expr) *matchedOrderDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m matchedOrderDo) Group(cols ...field.Expr) *matchedOrderDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m matchedOrderDo) Having(conds ...gen.Condition) *matchedOrderDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m matchedOrderDo) Limit(limit int) *matchedOrderDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m matchedOrderDo) Offset(offset int) *matchedOrderDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m matchedOrderDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *matchedOrderDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m matchedOrderDo) Unscoped() *matchedOrderDo {
	return m.withDO(m.DO.Unscoped())
}

func (m matchedOrderDo) Create(values ...*model.MatchedOrder) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m matchedOrderDo) CreateInBatches(values []*model.MatchedOrder, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m matchedOrderDo) Save(values ...*model.MatchedOrder) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m matchedOrderDo) First() (*model.MatchedOrder, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.MatchedOrder), nil
	}
}

func (m matchedOrderDo) Take() (*model.MatchedOrder, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.MatchedOrder), nil
	}
}

func (m matchedOrderDo) Last() (*model.MatchedOrder, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.MatchedOrder), nil
	}
}

func (m matchedOrderDo) Find() ([]*model.MatchedOrder, error) {
	result, err := m.DO.Find()
	return result.([]*model.MatchedOrder), err
}

func (m matchedOrderDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MatchedOrder, err error) {
	buf := make([]*model.MatchedOrder, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m matchedOrderDo) FindInBatches(result *[]*model.MatchedOrder, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m matchedOrderDo) Attrs(attrs ...field.AssignExpr) *matchedOrderDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m matchedOrderDo) Assign(attrs ...field.AssignExpr) *matchedOrderDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m matchedOrderDo) Joins(fields ...field.RelationField) *matchedOrderDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m matchedOrderDo) Preload(fields ...field.RelationField) *matchedOrderDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m matchedOrderDo) FirstOrInit() (*model.MatchedOrder, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.MatchedOrder), nil
	}
}

func (m matchedOrderDo) FirstOrCreate() (*model.MatchedOrder, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.MatchedOrder), nil
	}
}

func (m matchedOrderDo) FindByPage(offset int, limit int) (result []*model.MatchedOrder, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m matchedOrderDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m matchedOrderDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m matchedOrderDo) Delete(models ...*model.MatchedOrder) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *matchedOrderDo) withDO(do gen.Dao) *matchedOrderDo {
	m.DO = *do.(*gen.DO)
	return m
}
