package resolvers

import (
	"fmt"
	"reflect"
)

// Repository stores all resolvers
type Repository map[string]resolver

// Add stores a new resolver
func (r Repository) Add(resolve string, handler interface{}) error {
	err := validators.run(reflect.TypeOf(handler))

	if err == nil {
		r[resolve] = resolver{handler}
	}

	return err
}

// Handle responds to the AppSync request
func (r Repository) Handle(ctx context) (interface{}, error) {
	handler, found := r[ctx.resolver()]

	if found {
		return handler.call(ctx.payload(), ctx.headers(), ctx.identity())
	}

	return nil, fmt.Errorf("no resolver found: %s", ctx.resolver())
}
