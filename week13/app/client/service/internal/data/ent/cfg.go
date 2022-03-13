// Code generated by entc, DO NOT EDIT.

package ent

import (
	"dipper/app/client/service/internal/data/ent/cfg"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Cfg is the model entity for the Cfg schema.
type Cfg struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Key holds the value of the "key" field.
	// 键
	Key string `json:"key,omitempty"`
	// Value holds the value of the "value" field.
	// 值
	Value string `json:"value,omitempty"`
	// IsDeleted holds the value of the "is_deleted" field.
	// 是否被删除: 0否 1是
	IsDeleted uint8 `json:"is_deleted,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Cfg) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case cfg.FieldID, cfg.FieldIsDeleted:
			values[i] = new(sql.NullInt64)
		case cfg.FieldKey, cfg.FieldValue:
			values[i] = new(sql.NullString)
		case cfg.FieldCreatedAt, cfg.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Cfg", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Cfg fields.
func (c *Cfg) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case cfg.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case cfg.FieldKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field key", values[i])
			} else if value.Valid {
				c.Key = value.String
			}
		case cfg.FieldValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				c.Value = value.String
			}
		case cfg.FieldIsDeleted:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_deleted", values[i])
			} else if value.Valid {
				c.IsDeleted = uint8(value.Int64)
			}
		case cfg.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case cfg.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Cfg.
// Note that you need to call Cfg.Unwrap() before calling this method if this Cfg
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Cfg) Update() *CfgUpdateOne {
	return (&CfgClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Cfg entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Cfg) Unwrap() *Cfg {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Cfg is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Cfg) String() string {
	var builder strings.Builder
	builder.WriteString("Cfg(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", key=")
	builder.WriteString(c.Key)
	builder.WriteString(", value=")
	builder.WriteString(c.Value)
	builder.WriteString(", is_deleted=")
	builder.WriteString(fmt.Sprintf("%v", c.IsDeleted))
	builder.WriteString(", created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Cfgs is a parsable slice of Cfg.
type Cfgs []*Cfg

func (c Cfgs) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
