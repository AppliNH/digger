// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models_generated

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/diggerhq/drift/model"
)

func newOrganisation(db *gorm.DB, opts ...gen.DOOption) organisation {
	_organisation := organisation{}

	_organisation.organisationDo.UseDB(db, opts...)
	_organisation.organisationDo.UseModel(&model.Organisation{})

	tableName := _organisation.organisationDo.TableName()
	_organisation.ALL = field.NewAsterisk(tableName)
	_organisation.ID = field.NewString(tableName, "id")
	_organisation.CreatedAt = field.NewTime(tableName, "created_at")
	_organisation.UpdatedAt = field.NewTime(tableName, "updated_at")
	_organisation.DeletedAt = field.NewField(tableName, "deleted_at")
	_organisation.Name = field.NewString(tableName, "name")
	_organisation.ExternalSource = field.NewString(tableName, "external_source")
	_organisation.ExternalID = field.NewString(tableName, "external_id")

	_organisation.fillFieldMap()

	return _organisation
}

type organisation struct {
	organisationDo

	ALL            field.Asterisk
	ID             field.String
	CreatedAt      field.Time
	UpdatedAt      field.Time
	DeletedAt      field.Field
	Name           field.String
	ExternalSource field.String
	ExternalID     field.String

	fieldMap map[string]field.Expr
}

func (o organisation) Table(newTableName string) *organisation {
	o.organisationDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o organisation) As(alias string) *organisation {
	o.organisationDo.DO = *(o.organisationDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *organisation) updateTableName(table string) *organisation {
	o.ALL = field.NewAsterisk(table)
	o.ID = field.NewString(table, "id")
	o.CreatedAt = field.NewTime(table, "created_at")
	o.UpdatedAt = field.NewTime(table, "updated_at")
	o.DeletedAt = field.NewField(table, "deleted_at")
	o.Name = field.NewString(table, "name")
	o.ExternalSource = field.NewString(table, "external_source")
	o.ExternalID = field.NewString(table, "external_id")

	o.fillFieldMap()

	return o
}

func (o *organisation) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *organisation) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 7)
	o.fieldMap["id"] = o.ID
	o.fieldMap["created_at"] = o.CreatedAt
	o.fieldMap["updated_at"] = o.UpdatedAt
	o.fieldMap["deleted_at"] = o.DeletedAt
	o.fieldMap["name"] = o.Name
	o.fieldMap["external_source"] = o.ExternalSource
	o.fieldMap["external_id"] = o.ExternalID
}

func (o organisation) clone(db *gorm.DB) organisation {
	o.organisationDo.ReplaceConnPool(db.Statement.ConnPool)
	return o
}

func (o organisation) replaceDB(db *gorm.DB) organisation {
	o.organisationDo.ReplaceDB(db)
	return o
}

type organisationDo struct{ gen.DO }

type IOrganisationDo interface {
	gen.SubQuery
	Debug() IOrganisationDo
	WithContext(ctx context.Context) IOrganisationDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IOrganisationDo
	WriteDB() IOrganisationDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IOrganisationDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IOrganisationDo
	Not(conds ...gen.Condition) IOrganisationDo
	Or(conds ...gen.Condition) IOrganisationDo
	Select(conds ...field.Expr) IOrganisationDo
	Where(conds ...gen.Condition) IOrganisationDo
	Order(conds ...field.Expr) IOrganisationDo
	Distinct(cols ...field.Expr) IOrganisationDo
	Omit(cols ...field.Expr) IOrganisationDo
	Join(table schema.Tabler, on ...field.Expr) IOrganisationDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IOrganisationDo
	RightJoin(table schema.Tabler, on ...field.Expr) IOrganisationDo
	Group(cols ...field.Expr) IOrganisationDo
	Having(conds ...gen.Condition) IOrganisationDo
	Limit(limit int) IOrganisationDo
	Offset(offset int) IOrganisationDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IOrganisationDo
	Unscoped() IOrganisationDo
	Create(values ...*model.Organisation) error
	CreateInBatches(values []*model.Organisation, batchSize int) error
	Save(values ...*model.Organisation) error
	First() (*model.Organisation, error)
	Take() (*model.Organisation, error)
	Last() (*model.Organisation, error)
	Find() ([]*model.Organisation, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Organisation, err error)
	FindInBatches(result *[]*model.Organisation, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Organisation) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IOrganisationDo
	Assign(attrs ...field.AssignExpr) IOrganisationDo
	Joins(fields ...field.RelationField) IOrganisationDo
	Preload(fields ...field.RelationField) IOrganisationDo
	FirstOrInit() (*model.Organisation, error)
	FirstOrCreate() (*model.Organisation, error)
	FindByPage(offset int, limit int) (result []*model.Organisation, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IOrganisationDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (o organisationDo) Debug() IOrganisationDo {
	return o.withDO(o.DO.Debug())
}

func (o organisationDo) WithContext(ctx context.Context) IOrganisationDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o organisationDo) ReadDB() IOrganisationDo {
	return o.Clauses(dbresolver.Read)
}

func (o organisationDo) WriteDB() IOrganisationDo {
	return o.Clauses(dbresolver.Write)
}

func (o organisationDo) Session(config *gorm.Session) IOrganisationDo {
	return o.withDO(o.DO.Session(config))
}

func (o organisationDo) Clauses(conds ...clause.Expression) IOrganisationDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o organisationDo) Returning(value interface{}, columns ...string) IOrganisationDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o organisationDo) Not(conds ...gen.Condition) IOrganisationDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o organisationDo) Or(conds ...gen.Condition) IOrganisationDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o organisationDo) Select(conds ...field.Expr) IOrganisationDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o organisationDo) Where(conds ...gen.Condition) IOrganisationDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o organisationDo) Order(conds ...field.Expr) IOrganisationDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o organisationDo) Distinct(cols ...field.Expr) IOrganisationDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o organisationDo) Omit(cols ...field.Expr) IOrganisationDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o organisationDo) Join(table schema.Tabler, on ...field.Expr) IOrganisationDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o organisationDo) LeftJoin(table schema.Tabler, on ...field.Expr) IOrganisationDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o organisationDo) RightJoin(table schema.Tabler, on ...field.Expr) IOrganisationDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o organisationDo) Group(cols ...field.Expr) IOrganisationDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o organisationDo) Having(conds ...gen.Condition) IOrganisationDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o organisationDo) Limit(limit int) IOrganisationDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o organisationDo) Offset(offset int) IOrganisationDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o organisationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IOrganisationDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o organisationDo) Unscoped() IOrganisationDo {
	return o.withDO(o.DO.Unscoped())
}

func (o organisationDo) Create(values ...*model.Organisation) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o organisationDo) CreateInBatches(values []*model.Organisation, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o organisationDo) Save(values ...*model.Organisation) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o organisationDo) First() (*model.Organisation, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Organisation), nil
	}
}

func (o organisationDo) Take() (*model.Organisation, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Organisation), nil
	}
}

func (o organisationDo) Last() (*model.Organisation, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Organisation), nil
	}
}

func (o organisationDo) Find() ([]*model.Organisation, error) {
	result, err := o.DO.Find()
	return result.([]*model.Organisation), err
}

func (o organisationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Organisation, err error) {
	buf := make([]*model.Organisation, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o organisationDo) FindInBatches(result *[]*model.Organisation, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o organisationDo) Attrs(attrs ...field.AssignExpr) IOrganisationDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o organisationDo) Assign(attrs ...field.AssignExpr) IOrganisationDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o organisationDo) Joins(fields ...field.RelationField) IOrganisationDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o organisationDo) Preload(fields ...field.RelationField) IOrganisationDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o organisationDo) FirstOrInit() (*model.Organisation, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Organisation), nil
	}
}

func (o organisationDo) FirstOrCreate() (*model.Organisation, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Organisation), nil
	}
}

func (o organisationDo) FindByPage(offset int, limit int) (result []*model.Organisation, count int64, err error) {
	result, err = o.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = o.Offset(-1).Limit(-1).Count()
	return
}

func (o organisationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o organisationDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o organisationDo) Delete(models ...*model.Organisation) (result gen.ResultInfo, err error) {
	return o.DO.Delete(models)
}

func (o *organisationDo) withDO(do gen.Dao) *organisationDo {
	o.DO = *do.(*gen.DO)
	return o
}
