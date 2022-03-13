// Code generated by entc, DO NOT EDIT.

package ent

import (
	"dipper/app/server/service/internal/data/ent/host"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Host is the model entity for the Host schema.
type Host struct {
	config `json:"-"`
	// ID of the ent.
	// 索引 id
	ID int64 `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	// 机房名称
	Name string `json:"name,omitempty"`
	// Manager holds the value of the "manager" field.
	// 管理员
	Manager string `json:"manager,omitempty"`
	// Phone holds the value of the "phone" field.
	// 联系号码
	Phone string `json:"phone,omitempty"`
	// VerifyCode holds the value of the "verify_code" field.
	// 验证码
	VerifyCode string `json:"verify_code,omitempty"`
	// CertNum holds the value of the "cert_num" field.
	// 证书编码
	CertNum int64 `json:"cert_num,omitempty"`
	// IsDeleted holds the value of the "is_deleted" field.
	// 是否被删除: 0否 1是
	IsDeleted uint8 `json:"is_deleted,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Host) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case host.FieldID, host.FieldCertNum, host.FieldIsDeleted:
			values[i] = new(sql.NullInt64)
		case host.FieldName, host.FieldManager, host.FieldPhone, host.FieldVerifyCode:
			values[i] = new(sql.NullString)
		case host.FieldCreatedAt, host.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Host", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Host fields.
func (h *Host) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case host.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			h.ID = int64(value.Int64)
		case host.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				h.Name = value.String
			}
		case host.FieldManager:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field manager", values[i])
			} else if value.Valid {
				h.Manager = value.String
			}
		case host.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				h.Phone = value.String
			}
		case host.FieldVerifyCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field verify_code", values[i])
			} else if value.Valid {
				h.VerifyCode = value.String
			}
		case host.FieldCertNum:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field cert_num", values[i])
			} else if value.Valid {
				h.CertNum = value.Int64
			}
		case host.FieldIsDeleted:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_deleted", values[i])
			} else if value.Valid {
				h.IsDeleted = uint8(value.Int64)
			}
		case host.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				h.CreatedAt = value.Time
			}
		case host.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				h.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Host.
// Note that you need to call Host.Unwrap() before calling this method if this Host
// was returned from a transaction, and the transaction was committed or rolled back.
func (h *Host) Update() *HostUpdateOne {
	return (&HostClient{config: h.config}).UpdateOne(h)
}

// Unwrap unwraps the Host entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (h *Host) Unwrap() *Host {
	tx, ok := h.config.driver.(*txDriver)
	if !ok {
		panic("ent: Host is not a transactional entity")
	}
	h.config.driver = tx.drv
	return h
}

// String implements the fmt.Stringer.
func (h *Host) String() string {
	var builder strings.Builder
	builder.WriteString("Host(")
	builder.WriteString(fmt.Sprintf("id=%v", h.ID))
	builder.WriteString(", name=")
	builder.WriteString(h.Name)
	builder.WriteString(", manager=")
	builder.WriteString(h.Manager)
	builder.WriteString(", phone=")
	builder.WriteString(h.Phone)
	builder.WriteString(", verify_code=")
	builder.WriteString(h.VerifyCode)
	builder.WriteString(", cert_num=")
	builder.WriteString(fmt.Sprintf("%v", h.CertNum))
	builder.WriteString(", is_deleted=")
	builder.WriteString(fmt.Sprintf("%v", h.IsDeleted))
	builder.WriteString(", created_at=")
	builder.WriteString(h.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(h.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Hosts is a parsable slice of Host.
type Hosts []*Host

func (h Hosts) config(cfg config) {
	for _i := range h {
		h[_i].config = cfg
	}
}