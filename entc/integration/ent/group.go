// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/ent/group"
	"entgo.io/ent/entc/integration/ent/groupinfo"
)

// Group is the model entity for the Group schema.
type Group struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Active holds the value of the "active" field.
	Active bool `json:"active,omitempty"`
	// Expire holds the value of the "expire" field.
	Expire time.Time `json:"expire,omitempty"`
	// Type holds the value of the "type" field.
	Type *string `json:"type,omitempty"`
	// MaxUsers holds the value of the "max_users" field.
	MaxUsers int `json:"max_users,omitempty"`
	// Name holds the value of the "name" field.
	// field with multiple validators
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GroupQuery when eager-loading is set.
	Edges      GroupEdges `json:"edges"`
	group_info *int
}

// GroupEdges holds the relations/edges for other nodes in the graph.
type GroupEdges struct {
	// Files holds the value of the files edge.
	Files []*File `json:"files,omitempty"`
	// Blocked holds the value of the blocked edge.
	Blocked []*User `json:"blocked,omitempty"`
	// Users holds the value of the users edge.
	Users []*User `json:"users,omitempty"`
	// Info holds the value of the info edge.
	Info *GroupInfo `json:"info,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// FilesOrErr returns the Files value or an error if the edge
// was not loaded in eager-loading.
func (e GroupEdges) FilesOrErr() ([]*File, error) {
	if e.loadedTypes[0] {
		return e.Files, nil
	}
	return nil, &NotLoadedError{edge: "files"}
}

// BlockedOrErr returns the Blocked value or an error if the edge
// was not loaded in eager-loading.
func (e GroupEdges) BlockedOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.Blocked, nil
	}
	return nil, &NotLoadedError{edge: "blocked"}
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e GroupEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[2] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// InfoOrErr returns the Info value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GroupEdges) InfoOrErr() (*GroupInfo, error) {
	if e.loadedTypes[3] {
		if e.Info == nil {
			// The edge info was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: groupinfo.Label}
		}
		return e.Info, nil
	}
	return nil, &NotLoadedError{edge: "info"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Group) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case group.FieldActive:
			values[i] = new(sql.NullBool)
		case group.FieldID, group.FieldMaxUsers:
			values[i] = new(sql.NullInt64)
		case group.FieldType, group.FieldName:
			values[i] = new(sql.NullString)
		case group.FieldExpire:
			values[i] = new(sql.NullTime)
		case group.ForeignKeys[0]: // group_info
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Group", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Group fields.
func (gr *Group) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case group.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			gr.ID = int(value.Int64)
		case group.FieldActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field active", values[i])
			} else if value.Valid {
				gr.Active = value.Bool
			}
		case group.FieldExpire:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expire", values[i])
			} else if value.Valid {
				gr.Expire = value.Time
			}
		case group.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				gr.Type = new(string)
				*gr.Type = value.String
			}
		case group.FieldMaxUsers:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field max_users", values[i])
			} else if value.Valid {
				gr.MaxUsers = int(value.Int64)
			}
		case group.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				gr.Name = value.String
			}
		case group.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field group_info", value)
			} else if value.Valid {
				gr.group_info = new(int)
				*gr.group_info = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryFiles queries the "files" edge of the Group entity.
func (gr *Group) QueryFiles() *FileQuery {
	return (&GroupClient{config: gr.config}).QueryFiles(gr)
}

// QueryBlocked queries the "blocked" edge of the Group entity.
func (gr *Group) QueryBlocked() *UserQuery {
	return (&GroupClient{config: gr.config}).QueryBlocked(gr)
}

// QueryUsers queries the "users" edge of the Group entity.
func (gr *Group) QueryUsers() *UserQuery {
	return (&GroupClient{config: gr.config}).QueryUsers(gr)
}

// QueryInfo queries the "info" edge of the Group entity.
func (gr *Group) QueryInfo() *GroupInfoQuery {
	return (&GroupClient{config: gr.config}).QueryInfo(gr)
}

// Update returns a builder for updating this Group.
// Note that you need to call Group.Unwrap() before calling this method if this Group
// was returned from a transaction, and the transaction was committed or rolled back.
func (gr *Group) Update() *GroupUpdateOne {
	return (&GroupClient{config: gr.config}).UpdateOne(gr)
}

// Unwrap unwraps the Group entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gr *Group) Unwrap() *Group {
	tx, ok := gr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Group is not a transactional entity")
	}
	gr.config.driver = tx.drv
	return gr
}

// String implements the fmt.Stringer.
func (gr *Group) String() string {
	var builder strings.Builder
	builder.WriteString("Group(")
	builder.WriteString(fmt.Sprintf("id=%v, ", gr.ID))
	builder.WriteString("active=")
	builder.WriteString(fmt.Sprintf("%v", gr.Active))
	builder.WriteString(", ")
	builder.WriteString("expire=")
	builder.WriteString(gr.Expire.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := gr.Type; v != nil {
		builder.WriteString("type=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("max_users=")
	builder.WriteString(fmt.Sprintf("%v", gr.MaxUsers))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(gr.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Groups is a parsable slice of Group.
type Groups []*Group

func (gr Groups) config(cfg config) {
	for _i := range gr {
		gr[_i].config = cfg
	}
}
