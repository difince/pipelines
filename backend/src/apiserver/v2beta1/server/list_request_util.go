// Copyright 2022 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"net/url"

	"github.com/golang/protobuf/jsonpb"
	apiv2beta1 "github.com/kubeflow/pipelines/backend/api/v2beta1/go_client"
	"github.com/kubeflow/pipelines/backend/src/apiserver/v2beta1/list"
	"github.com/kubeflow/pipelines/backend/src/common/util"
)

const (
	defaultPageSize = 20
	maxPageSize     = 200
)

// parseAPIFilter attempts to decode a url-encoded JSON-stringified api
// filter object. An empty string is considered valid input, and equivalent to
// the nil filter, which trivially does nothing.
func parseAPIFilter(encoded string) (*apiv2beta1.Filter, error) {
	if encoded == "" {
		return nil, nil
	}

	errorF := func(err error) (*apiv2beta1.Filter, error) {
		return nil, util.NewInvalidInputError("failed to parse valid filter from %q: %v", encoded, err)
	}

	decoded, err := url.QueryUnescape(encoded)
	if err != nil {
		return errorF(err)
	}

	f := &apiv2beta1.Filter{}
	if err := jsonpb.UnmarshalString(string(decoded), f); err != nil {
		return errorF(err)
	}
	return f, nil
}

func validatedListOptions(listable list.Listable, pageToken string, pageSize int, sortBy string, filterSpec string) (*list.Options, error) {
	defaultOpts := func() (*list.Options, error) {
		if listable == nil {
			return nil, util.NewInvalidInputError("Please specify a valid type to list. E.g., list runs or list jobs.")
		}

		f, err := parseAPIFilter(filterSpec)
		if err != nil {
			return nil, err
		}

		return list.NewOptions(listable, pageSize, sortBy, f)
	}

	if pageToken == "" {
		return defaultOpts()
	}

	opts, err := list.NewOptionsFromToken(pageToken, pageSize)
	if err != nil {
		return nil, err
	}

	if sortBy != "" || filterSpec != "" {
		// Sanity check that these match the page token.
		do, err := defaultOpts()
		if err != nil {
			return nil, err
		}

		if !opts.Matches(do) {
			return nil, util.NewInvalidInputError("page token does not match the supplied sort by and/or filtering criteria. Either specify the same criteria or leave the latter empty if page token is specified.")
		}
	}

	return opts, nil
}
