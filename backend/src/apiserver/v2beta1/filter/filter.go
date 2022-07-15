// Copyright 2018 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package filter contains types and methods for parsing and applying filters to
// resources being queried by a ListXXX request.
package filter

import (
	"encoding/json"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes"
	"github.com/kubeflow/pipelines/backend/src/common/util"

	apiv2beta1 "github.com/kubeflow/pipelines/backend/api/v2beta1/go_client"
)

// Filter represents a filter that can be applied when querying an arbitrary API
// resource.
type Filter struct {
	filterProto *apiv2beta1.Filter

	eq  map[string][]interface{}
	neq map[string][]interface{}
	gt  map[string][]interface{}
	gte map[string][]interface{}
	lt  map[string][]interface{}
	lte map[string][]interface{}

	in map[string][]interface{}

	substring map[string][]interface{}
}

// filterForMarshaling is a helper struct for marshaling Filter into JSON. This
// is needed as we don't want to export the fields in Filter.
type filterForMarshaling struct {
	FilterProto string

	EQ  map[string][]interface{}
	NEQ map[string][]interface{}
	GT  map[string][]interface{}
	GTE map[string][]interface{}
	LT  map[string][]interface{}
	LTE map[string][]interface{}

	IN map[string][]interface{}

	SUBSTRING map[string][]interface{}
}

// MarshalJSON implements JSON Marshaler for Filter.
func (f *Filter) MarshalJSON() ([]byte, error) {
	m := &jsonpb.Marshaler{}
	s, err := m.MarshalToString(f.filterProto)
	if err != nil {
		return nil, err
	}
	return json.Marshal(&filterForMarshaling{
		FilterProto: s,
		EQ:          f.eq,
		NEQ:         f.neq,
		GT:          f.gt,
		GTE:         f.gte,
		LT:          f.lt,
		LTE:         f.lte,
		IN:          f.in,
		SUBSTRING:   f.substring,
	})
}

// UnmarshalJSON implements JSON Unmarshaler for Filter.
func (f *Filter) UnmarshalJSON(b []byte) error {
	ffm := filterForMarshaling{}
	err := json.Unmarshal(b, &ffm)
	if err != nil {
		return err
	}

	f.filterProto = &apiv2beta1.Filter{}
	err = jsonpb.UnmarshalString(ffm.FilterProto, f.filterProto)
	if err != nil {
		return err
	}

	f.eq = ffm.EQ
	f.neq = ffm.NEQ
	f.gt = ffm.GT
	f.gte = ffm.GTE
	f.lt = ffm.LT
	f.lte = ffm.LTE
	f.in = ffm.IN
	f.substring = ffm.SUBSTRING

	return nil
}

// New creates a new Filter from parsing the API filter protocol buffer.
func New(filterProto *apiv2beta1.Filter) (*Filter, error) {
	f := &Filter{
		filterProto: filterProto,
		eq:          make(map[string][]interface{}, 0),
		neq:         make(map[string][]interface{}, 0),
		gt:          make(map[string][]interface{}, 0),
		gte:         make(map[string][]interface{}, 0),
		lt:          make(map[string][]interface{}, 0),
		lte:         make(map[string][]interface{}, 0),
		in:          make(map[string][]interface{}, 0),
		substring:   make(map[string][]interface{}, 0),
	}

	if err := f.parseFilterProto(); err != nil {
		return nil, err
	}
	return f, nil
}

// NewWithKeyMap is like New, but takes an additional map and model name for mapping key names
// in the protocol buffer to an appropriate name for use when querying the
// model. For example, if the API name of a field is "name", the model name is "pipelines", and
// the equivalent column name is "Name", then filterProto with predicates against key "name"
// will be parsed as if the key value was "pipelines.Name".
func NewWithKeyMap(filterProto *apiv2beta1.Filter, keyMap map[string]string, modelName string) (*Filter, error) {
	// Fully qualify column name to avoid "ambiguous column name" error.
	var modelNamePrefix string
	if modelName != "" {
		modelNamePrefix = modelName + "."
	}

	for _, pred := range filterProto.Predicates {
		k, ok := keyMap[pred.Key]
		if !ok {
			return nil, util.NewInvalidInputError("no support for filtering on unrecognized field %q", pred.Key)
		}
		pred.Key = modelNamePrefix + k
	}
	return New(filterProto)
}

// AddToSelect builds a WHERE clause from the Filter f, adds it to the supplied
// SelectBuilder object and returns it for use in SQL queries.
func (f *Filter) AddToSelect(sb squirrel.SelectBuilder) squirrel.SelectBuilder {
	for k := range f.eq {
		for _, v := range f.eq[k] {
			m := map[string]interface{}{k: v}
			sb = sb.Where(squirrel.Eq(m))
		}
	}

	for k := range f.neq {
		for _, v := range f.neq[k] {
			m := map[string]interface{}{k: v}
			sb = sb.Where(squirrel.NotEq(m))
		}
	}

	for k := range f.gt {
		for _, v := range f.gt[k] {
			m := map[string]interface{}{k: v}
			sb = sb.Where(squirrel.Gt(m))
		}
	}

	for k := range f.gte {
		for _, v := range f.gte[k] {
			m := map[string]interface{}{k: v}
			sb = sb.Where(squirrel.GtOrEq(m))
		}
	}

	for k := range f.lt {
		for _, v := range f.lt[k] {
			m := map[string]interface{}{k: v}
			sb = sb.Where(squirrel.Lt(m))
		}
	}

	for k := range f.lte {
		for _, v := range f.lte[k] {
			m := map[string]interface{}{k: v}
			sb = sb.Where(squirrel.LtOrEq(m))
		}
	}

	// In
	for k := range f.in {
		for _, v := range f.in[k] {
			m := map[string]interface{}{k: v}
			sb = sb.Where(squirrel.Eq(m))
		}
	}

	for k := range f.substring {
		// Modify each string value v so it looks like %v% so we are doing a substring
		// match with the LIKE operator.
		for _, v := range f.substring[k] {
			like := make(squirrel.Like)
			like[k] = fmt.Sprintf("%%%s%%", v)
			sb = sb.Where(like)
		}
	}

	return sb
}

func checkPredicate(p *apiv2beta1.Predicate) error {
	switch p.Op {
	case apiv2beta1.Predicate_IN:
		switch t := p.Value.(type) {
		case *apiv2beta1.Predicate_IntValue, *apiv2beta1.Predicate_LongValue, *apiv2beta1.Predicate_StringValue, *apiv2beta1.Predicate_TimestampValue:
			return util.NewInvalidInputError("cannot use IN operator with scalar type %T", t)
		}

	case apiv2beta1.Predicate_EQUALS, apiv2beta1.Predicate_NOT_EQUALS, apiv2beta1.Predicate_GREATER_THAN, apiv2beta1.Predicate_GREATER_THAN_EQUALS, apiv2beta1.Predicate_LESS_THAN, apiv2beta1.Predicate_LESS_THAN_EQUALS:
		switch t := p.Value.(type) {
		case *apiv2beta1.Predicate_IntValues, *apiv2beta1.Predicate_LongValues, *apiv2beta1.Predicate_StringValues:
			return util.NewInvalidInputError("cannot use scalar operator %v on array type %T", p.Op, t)
		}

	case apiv2beta1.Predicate_IS_SUBSTRING:
		switch t := p.Value.(type) {
		case *apiv2beta1.Predicate_StringValue:
			return nil
		default:
			return util.NewInvalidInputError("cannot use non string value type %T with operator %v", p.Op, t)
		}

	default:
		return util.NewInvalidInputError("invalid predicate operation: %v", p.Op)
	}

	return nil
}

func (f *Filter) parseFilterProto() error {
	for _, pred := range f.filterProto.Predicates {
		if err := checkPredicate(pred); err != nil {
			return err
		}

		var m map[string][]interface{}
		switch pred.Op {
		case apiv2beta1.Predicate_EQUALS:
			m = f.eq
		case apiv2beta1.Predicate_NOT_EQUALS:
			m = f.neq
		case apiv2beta1.Predicate_GREATER_THAN:
			m = f.gt
		case apiv2beta1.Predicate_GREATER_THAN_EQUALS:
			m = f.gte
		case apiv2beta1.Predicate_LESS_THAN:
			m = f.lt
		case apiv2beta1.Predicate_LESS_THAN_EQUALS:
			m = f.lte
		case apiv2beta1.Predicate_IN:
			m = f.in
		case apiv2beta1.Predicate_IS_SUBSTRING:
			m = f.substring
		default:
			return util.NewInvalidInputError("invalid predicate operation: %v", pred.Op)
		}

		if err := addPredicateValue(m, pred); err != nil {
			return err
		}
	}

	return nil
}

func addPredicateValue(m map[string][]interface{}, p *apiv2beta1.Predicate) error {
	switch t := p.Value.(type) {
	case *apiv2beta1.Predicate_IntValue:
		m[p.Key] = append(m[p.Key], p.GetIntValue())
	case *apiv2beta1.Predicate_LongValue:
		m[p.Key] = append(m[p.Key], p.GetLongValue())
	case *apiv2beta1.Predicate_StringValue:
		m[p.Key] = append(m[p.Key], p.GetStringValue())
	case *apiv2beta1.Predicate_TimestampValue:
		ts, err := ptypes.Timestamp(p.GetTimestampValue())
		if err != nil {
			return util.NewInvalidInputError("invalid timestamp: %v", err)
		}
		m[p.Key] = append(m[p.Key], ts.Unix())

	case *apiv2beta1.Predicate_IntValues:
		var v []int32
		for _, i := range p.GetIntValues().GetValues() {
			v = append(v, i)
		}
		m[p.Key] = append(m[p.Key], v)

	case *apiv2beta1.Predicate_LongValues:
		var v []int64
		for _, i := range p.GetLongValues().GetValues() {
			v = append(v, i)
		}
		m[p.Key] = append(m[p.Key], v)

	case *apiv2beta1.Predicate_StringValues:
		var v []string
		for _, i := range p.GetStringValues().GetValues() {
			v = append(v, i)
		}
		m[p.Key] = append(m[p.Key], v)

	case nil:
		return util.NewInvalidInputError("no value set for predicate on key %q", p.Key)

	default:
		return util.NewInvalidInputError("unknown value type in Filter for predicate key %q: %T", p.Key, t)
	}

	return nil
}
