package list

import (
	"net/url"

	"github.com/golang/protobuf/jsonpb"

	api "github.com/kubeflow/pipelines/backend/api/go_client"
	"github.com/kubeflow/pipelines/backend/src/common/util"
)

// parseAPIFilter attempts to decode a url-encoded JSON-stringified api
// filter object. An empty string is considered valid input, and equivalent to
// the nil filter, which trivially does nothing.
func parseAPIFilter(encoded string) (*api.Filter, error) {
	if encoded == "" {
		return nil, nil
	}

	errorF := func(err error) (*api.Filter, error) {
		return nil, util.NewInvalidInputError("failed to parse valid filter from %q: %v", encoded, err)
	}

	decoded, err := url.QueryUnescape(encoded)
	if err != nil {
		return errorF(err)
	}

	f := &api.Filter{}
	if err := jsonpb.UnmarshalString(string(decoded), f); err != nil {
		return errorF(err)
	}
	return f, nil
}

func ValidatedListOptions(listable Listable, pageToken string, pageSize int, sortBy string, filterSpec string) (*Options, error) {
	defaultOpts := func() (*Options, error) {
		if listable == nil {
			return nil, util.NewInvalidInputError("Please specify a valid type to list. E.g., list runs or list jobs.")
		}

		f, err := parseAPIFilter(filterSpec)
		if err != nil {
			return nil, err
		}

		return NewOptions(listable, pageSize, sortBy, f)
	}

	if pageToken == "" {
		return defaultOpts()
	}

	opts, err := NewOptionsFromToken(pageToken, pageSize)
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
