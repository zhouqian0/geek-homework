// Code generated by entc, DO NOT EDIT.

package host

import (
	"dipper/app/server/service/internal/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Manager applies equality check predicate on the "manager" field. It's identical to ManagerEQ.
func Manager(v string) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldManager), v))
	})
}

// Phone applies equality check predicate on the "phone" field. It's identical to PhoneEQ.
func Phone(v string) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPhone), v))
	})
}

// VerifyCode applies equality check predicate on the "verify_code" field. It's identical to VerifyCodeEQ.
func VerifyCode(v string) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVerifyCode), v))
	})
}

// CertNum applies equality check predicate on the "cert_num" field. It's identical to CertNumEQ.
func CertNum(v int64) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCertNum), v))
	})
}

// IsDeleted applies equality check predicate on the "is_deleted" field. It's identical to IsDeletedEQ.
func IsDeleted(v uint8) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsDeleted), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Host {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Host(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Host {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Host(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// ManagerEQ applies the EQ predicate on the "manager" field.
func ManagerEQ(v string) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldManager), v))
	})
}

// ManagerNEQ applies the NEQ predicate on the "manager" field.
func ManagerNEQ(v string) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldManager), v))
	})
}

// ManagerIn applies the In predicate on the "manager" field.
func ManagerIn(vs ...string) predicate.Host {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Host(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldManager), v...))
	})
}

// ManagerNotIn applies the NotIn predicate on the "manager" field.
func ManagerNotIn(vs ...string) predicate.Host {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Host(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldManager), v...))
	})
}

// ManagerGT applies the GT predicate on the "manager" field.
func ManagerGT(v string) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldManager), v))
	})
}

// ManagerGTE applies the GTE predicate on the "manager" field.
func ManagerGTE(v string) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldManager), v))
	})
}

// ManagerLT applies the LT predicate on the "manager" field.
func ManagerLT(v string) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldManager), v))
	})
}

// ManagerLTE applies the LTE predicate on the "manager" field.
func ManagerLTE(v string) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldManager), v))
	})
}

// ManagerContains applies the Contains predicate on the "manager" field.
func ManagerContains(v string) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldManager), v))
	})
}

// ManagerHasPrefix applies the HasPrefix predicate on the "manager" field.
func ManagerHasPrefix(v string) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldManager), v))
	})
}

// ManagerHasSuffix applies the HasSuffix predicate on the "manager" field.
func ManagerHasSuffix(v string) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldManager), v))
	})
}

// ManagerEqualFold applies the EqualFold predicate on the "manager" field.
func ManagerEqualFold(v string) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldManager), v))
	})
}

// ManagerContainsFold applies the ContainsFold predicate on the "manager" field.
func ManagerContainsFold(v string) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldManager), v))
	})
}

// PhoneEQ applies the EQ predicate on the "phone" field.
func PhoneEQ(v string) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPhone), v))
	})
}

// PhoneNEQ applies the NEQ predicate on the "phone" field.
func PhoneNEQ(v string) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPhone), v))
	})
}

// PhoneIn applies the In predicate on the "phone" field.
func PhoneIn(vs ...string) predicate.Host {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Host(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPhone), v...))
	})
}

// PhoneNotIn applies the NotIn predicate on the "phone" field.
func PhoneNotIn(vs ...string) predicate.Host {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Host(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPhone), v...))
	})
}

// PhoneGT applies the GT predicate on the "phone" field.
func PhoneGT(v string) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPhone), v))
	})
}

// PhoneGTE applies the GTE predicate on the "phone" field.
func PhoneGTE(v string) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPhone), v))
	})
}

// PhoneLT applies the LT predicate on the "phone" field.
func PhoneLT(v string) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPhone), v))
	})
}

// PhoneLTE applies the LTE predicate on the "phone" field.
func PhoneLTE(v string) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPhone), v))
	})
}

// PhoneContains applies the Contains predicate on the "phone" field.
func PhoneContains(v string) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPhone), v))
	})
}

// PhoneHasPrefix applies the HasPrefix predicate on the "phone" field.
func PhoneHasPrefix(v string) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPhone), v))
	})
}

// PhoneHasSuffix applies the HasSuffix predicate on the "phone" field.
func PhoneHasSuffix(v string) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPhone), v))
	})
}

// PhoneEqualFold applies the EqualFold predicate on the "phone" field.
func PhoneEqualFold(v string) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPhone), v))
	})
}

// PhoneContainsFold applies the ContainsFold predicate on the "phone" field.
func PhoneContainsFold(v string) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPhone), v))
	})
}

// VerifyCodeEQ applies the EQ predicate on the "verify_code" field.
func VerifyCodeEQ(v string) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVerifyCode), v))
	})
}

// VerifyCodeNEQ applies the NEQ predicate on the "verify_code" field.
func VerifyCodeNEQ(v string) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldVerifyCode), v))
	})
}

// VerifyCodeIn applies the In predicate on the "verify_code" field.
func VerifyCodeIn(vs ...string) predicate.Host {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Host(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldVerifyCode), v...))
	})
}

// VerifyCodeNotIn applies the NotIn predicate on the "verify_code" field.
func VerifyCodeNotIn(vs ...string) predicate.Host {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Host(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldVerifyCode), v...))
	})
}

// VerifyCodeGT applies the GT predicate on the "verify_code" field.
func VerifyCodeGT(v string) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldVerifyCode), v))
	})
}

// VerifyCodeGTE applies the GTE predicate on the "verify_code" field.
func VerifyCodeGTE(v string) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldVerifyCode), v))
	})
}

// VerifyCodeLT applies the LT predicate on the "verify_code" field.
func VerifyCodeLT(v string) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldVerifyCode), v))
	})
}

// VerifyCodeLTE applies the LTE predicate on the "verify_code" field.
func VerifyCodeLTE(v string) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldVerifyCode), v))
	})
}

// VerifyCodeContains applies the Contains predicate on the "verify_code" field.
func VerifyCodeContains(v string) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldVerifyCode), v))
	})
}

// VerifyCodeHasPrefix applies the HasPrefix predicate on the "verify_code" field.
func VerifyCodeHasPrefix(v string) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldVerifyCode), v))
	})
}

// VerifyCodeHasSuffix applies the HasSuffix predicate on the "verify_code" field.
func VerifyCodeHasSuffix(v string) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldVerifyCode), v))
	})
}

// VerifyCodeEqualFold applies the EqualFold predicate on the "verify_code" field.
func VerifyCodeEqualFold(v string) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldVerifyCode), v))
	})
}

// VerifyCodeContainsFold applies the ContainsFold predicate on the "verify_code" field.
func VerifyCodeContainsFold(v string) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldVerifyCode), v))
	})
}

// CertNumEQ applies the EQ predicate on the "cert_num" field.
func CertNumEQ(v int64) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCertNum), v))
	})
}

// CertNumNEQ applies the NEQ predicate on the "cert_num" field.
func CertNumNEQ(v int64) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCertNum), v))
	})
}

// CertNumIn applies the In predicate on the "cert_num" field.
func CertNumIn(vs ...int64) predicate.Host {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Host(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCertNum), v...))
	})
}

// CertNumNotIn applies the NotIn predicate on the "cert_num" field.
func CertNumNotIn(vs ...int64) predicate.Host {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Host(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCertNum), v...))
	})
}

// CertNumGT applies the GT predicate on the "cert_num" field.
func CertNumGT(v int64) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCertNum), v))
	})
}

// CertNumGTE applies the GTE predicate on the "cert_num" field.
func CertNumGTE(v int64) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCertNum), v))
	})
}

// CertNumLT applies the LT predicate on the "cert_num" field.
func CertNumLT(v int64) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCertNum), v))
	})
}

// CertNumLTE applies the LTE predicate on the "cert_num" field.
func CertNumLTE(v int64) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCertNum), v))
	})
}

// IsDeletedEQ applies the EQ predicate on the "is_deleted" field.
func IsDeletedEQ(v uint8) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsDeleted), v))
	})
}

// IsDeletedNEQ applies the NEQ predicate on the "is_deleted" field.
func IsDeletedNEQ(v uint8) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsDeleted), v))
	})
}

// IsDeletedIn applies the In predicate on the "is_deleted" field.
func IsDeletedIn(vs ...uint8) predicate.Host {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Host(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldIsDeleted), v...))
	})
}

// IsDeletedNotIn applies the NotIn predicate on the "is_deleted" field.
func IsDeletedNotIn(vs ...uint8) predicate.Host {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Host(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldIsDeleted), v...))
	})
}

// IsDeletedGT applies the GT predicate on the "is_deleted" field.
func IsDeletedGT(v uint8) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIsDeleted), v))
	})
}

// IsDeletedGTE applies the GTE predicate on the "is_deleted" field.
func IsDeletedGTE(v uint8) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIsDeleted), v))
	})
}

// IsDeletedLT applies the LT predicate on the "is_deleted" field.
func IsDeletedLT(v uint8) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIsDeleted), v))
	})
}

// IsDeletedLTE applies the LTE predicate on the "is_deleted" field.
func IsDeletedLTE(v uint8) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIsDeleted), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Host {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Host(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Host {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Host(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Host {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Host(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Host {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Host(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Host) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Host) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Host) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		p(s.Not())
	})
}