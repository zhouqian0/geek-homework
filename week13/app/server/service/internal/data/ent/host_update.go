// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"dipper/app/server/service/internal/data/ent/host"
	"dipper/app/server/service/internal/data/ent/predicate"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// HostUpdate is the builder for updating Host entities.
type HostUpdate struct {
	config
	hooks    []Hook
	mutation *HostMutation
}

// Where appends a list predicates to the HostUpdate builder.
func (hu *HostUpdate) Where(ps ...predicate.Host) *HostUpdate {
	hu.mutation.Where(ps...)
	return hu
}

// SetName sets the "name" field.
func (hu *HostUpdate) SetName(s string) *HostUpdate {
	hu.mutation.SetName(s)
	return hu
}

// SetManager sets the "manager" field.
func (hu *HostUpdate) SetManager(s string) *HostUpdate {
	hu.mutation.SetManager(s)
	return hu
}

// SetPhone sets the "phone" field.
func (hu *HostUpdate) SetPhone(s string) *HostUpdate {
	hu.mutation.SetPhone(s)
	return hu
}

// SetVerifyCode sets the "verify_code" field.
func (hu *HostUpdate) SetVerifyCode(s string) *HostUpdate {
	hu.mutation.SetVerifyCode(s)
	return hu
}

// SetCertNum sets the "cert_num" field.
func (hu *HostUpdate) SetCertNum(i int64) *HostUpdate {
	hu.mutation.ResetCertNum()
	hu.mutation.SetCertNum(i)
	return hu
}

// SetNillableCertNum sets the "cert_num" field if the given value is not nil.
func (hu *HostUpdate) SetNillableCertNum(i *int64) *HostUpdate {
	if i != nil {
		hu.SetCertNum(*i)
	}
	return hu
}

// AddCertNum adds i to the "cert_num" field.
func (hu *HostUpdate) AddCertNum(i int64) *HostUpdate {
	hu.mutation.AddCertNum(i)
	return hu
}

// SetIsDeleted sets the "is_deleted" field.
func (hu *HostUpdate) SetIsDeleted(u uint8) *HostUpdate {
	hu.mutation.ResetIsDeleted()
	hu.mutation.SetIsDeleted(u)
	return hu
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (hu *HostUpdate) SetNillableIsDeleted(u *uint8) *HostUpdate {
	if u != nil {
		hu.SetIsDeleted(*u)
	}
	return hu
}

// AddIsDeleted adds u to the "is_deleted" field.
func (hu *HostUpdate) AddIsDeleted(u int8) *HostUpdate {
	hu.mutation.AddIsDeleted(u)
	return hu
}

// SetCreatedAt sets the "created_at" field.
func (hu *HostUpdate) SetCreatedAt(t time.Time) *HostUpdate {
	hu.mutation.SetCreatedAt(t)
	return hu
}

// SetUpdatedAt sets the "updated_at" field.
func (hu *HostUpdate) SetUpdatedAt(t time.Time) *HostUpdate {
	hu.mutation.SetUpdatedAt(t)
	return hu
}

// Mutation returns the HostMutation object of the builder.
func (hu *HostUpdate) Mutation() *HostMutation {
	return hu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (hu *HostUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	hu.defaults()
	if len(hu.hooks) == 0 {
		affected, err = hu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HostMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			hu.mutation = mutation
			affected, err = hu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(hu.hooks) - 1; i >= 0; i-- {
			if hu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = hu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, hu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (hu *HostUpdate) SaveX(ctx context.Context) int {
	affected, err := hu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (hu *HostUpdate) Exec(ctx context.Context) error {
	_, err := hu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hu *HostUpdate) ExecX(ctx context.Context) {
	if err := hu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hu *HostUpdate) defaults() {
	if _, ok := hu.mutation.CreatedAt(); !ok {
		v := host.UpdateDefaultCreatedAt()
		hu.mutation.SetCreatedAt(v)
	}
	if _, ok := hu.mutation.UpdatedAt(); !ok {
		v := host.UpdateDefaultUpdatedAt()
		hu.mutation.SetUpdatedAt(v)
	}
}

func (hu *HostUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   host.Table,
			Columns: host.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: host.FieldID,
			},
		},
	}
	if ps := hu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: host.FieldName,
		})
	}
	if value, ok := hu.mutation.Manager(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: host.FieldManager,
		})
	}
	if value, ok := hu.mutation.Phone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: host.FieldPhone,
		})
	}
	if value, ok := hu.mutation.VerifyCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: host.FieldVerifyCode,
		})
	}
	if value, ok := hu.mutation.CertNum(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: host.FieldCertNum,
		})
	}
	if value, ok := hu.mutation.AddedCertNum(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: host.FieldCertNum,
		})
	}
	if value, ok := hu.mutation.IsDeleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: host.FieldIsDeleted,
		})
	}
	if value, ok := hu.mutation.AddedIsDeleted(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: host.FieldIsDeleted,
		})
	}
	if value, ok := hu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: host.FieldCreatedAt,
		})
	}
	if value, ok := hu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: host.FieldUpdatedAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, hu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{host.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// HostUpdateOne is the builder for updating a single Host entity.
type HostUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *HostMutation
}

// SetName sets the "name" field.
func (huo *HostUpdateOne) SetName(s string) *HostUpdateOne {
	huo.mutation.SetName(s)
	return huo
}

// SetManager sets the "manager" field.
func (huo *HostUpdateOne) SetManager(s string) *HostUpdateOne {
	huo.mutation.SetManager(s)
	return huo
}

// SetPhone sets the "phone" field.
func (huo *HostUpdateOne) SetPhone(s string) *HostUpdateOne {
	huo.mutation.SetPhone(s)
	return huo
}

// SetVerifyCode sets the "verify_code" field.
func (huo *HostUpdateOne) SetVerifyCode(s string) *HostUpdateOne {
	huo.mutation.SetVerifyCode(s)
	return huo
}

// SetCertNum sets the "cert_num" field.
func (huo *HostUpdateOne) SetCertNum(i int64) *HostUpdateOne {
	huo.mutation.ResetCertNum()
	huo.mutation.SetCertNum(i)
	return huo
}

// SetNillableCertNum sets the "cert_num" field if the given value is not nil.
func (huo *HostUpdateOne) SetNillableCertNum(i *int64) *HostUpdateOne {
	if i != nil {
		huo.SetCertNum(*i)
	}
	return huo
}

// AddCertNum adds i to the "cert_num" field.
func (huo *HostUpdateOne) AddCertNum(i int64) *HostUpdateOne {
	huo.mutation.AddCertNum(i)
	return huo
}

// SetIsDeleted sets the "is_deleted" field.
func (huo *HostUpdateOne) SetIsDeleted(u uint8) *HostUpdateOne {
	huo.mutation.ResetIsDeleted()
	huo.mutation.SetIsDeleted(u)
	return huo
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (huo *HostUpdateOne) SetNillableIsDeleted(u *uint8) *HostUpdateOne {
	if u != nil {
		huo.SetIsDeleted(*u)
	}
	return huo
}

// AddIsDeleted adds u to the "is_deleted" field.
func (huo *HostUpdateOne) AddIsDeleted(u int8) *HostUpdateOne {
	huo.mutation.AddIsDeleted(u)
	return huo
}

// SetCreatedAt sets the "created_at" field.
func (huo *HostUpdateOne) SetCreatedAt(t time.Time) *HostUpdateOne {
	huo.mutation.SetCreatedAt(t)
	return huo
}

// SetUpdatedAt sets the "updated_at" field.
func (huo *HostUpdateOne) SetUpdatedAt(t time.Time) *HostUpdateOne {
	huo.mutation.SetUpdatedAt(t)
	return huo
}

// Mutation returns the HostMutation object of the builder.
func (huo *HostUpdateOne) Mutation() *HostMutation {
	return huo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (huo *HostUpdateOne) Select(field string, fields ...string) *HostUpdateOne {
	huo.fields = append([]string{field}, fields...)
	return huo
}

// Save executes the query and returns the updated Host entity.
func (huo *HostUpdateOne) Save(ctx context.Context) (*Host, error) {
	var (
		err  error
		node *Host
	)
	huo.defaults()
	if len(huo.hooks) == 0 {
		node, err = huo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HostMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			huo.mutation = mutation
			node, err = huo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(huo.hooks) - 1; i >= 0; i-- {
			if huo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = huo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, huo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (huo *HostUpdateOne) SaveX(ctx context.Context) *Host {
	node, err := huo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (huo *HostUpdateOne) Exec(ctx context.Context) error {
	_, err := huo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (huo *HostUpdateOne) ExecX(ctx context.Context) {
	if err := huo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (huo *HostUpdateOne) defaults() {
	if _, ok := huo.mutation.CreatedAt(); !ok {
		v := host.UpdateDefaultCreatedAt()
		huo.mutation.SetCreatedAt(v)
	}
	if _, ok := huo.mutation.UpdatedAt(); !ok {
		v := host.UpdateDefaultUpdatedAt()
		huo.mutation.SetUpdatedAt(v)
	}
}

func (huo *HostUpdateOne) sqlSave(ctx context.Context) (_node *Host, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   host.Table,
			Columns: host.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: host.FieldID,
			},
		},
	}
	id, ok := huo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Host.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := huo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, host.FieldID)
		for _, f := range fields {
			if !host.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != host.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := huo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := huo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: host.FieldName,
		})
	}
	if value, ok := huo.mutation.Manager(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: host.FieldManager,
		})
	}
	if value, ok := huo.mutation.Phone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: host.FieldPhone,
		})
	}
	if value, ok := huo.mutation.VerifyCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: host.FieldVerifyCode,
		})
	}
	if value, ok := huo.mutation.CertNum(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: host.FieldCertNum,
		})
	}
	if value, ok := huo.mutation.AddedCertNum(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: host.FieldCertNum,
		})
	}
	if value, ok := huo.mutation.IsDeleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: host.FieldIsDeleted,
		})
	}
	if value, ok := huo.mutation.AddedIsDeleted(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: host.FieldIsDeleted,
		})
	}
	if value, ok := huo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: host.FieldCreatedAt,
		})
	}
	if value, ok := huo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: host.FieldUpdatedAt,
		})
	}
	_node = &Host{config: huo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, huo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{host.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}