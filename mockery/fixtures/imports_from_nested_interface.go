package test

import (
	"github.com/DispatchMe/mockery/mockery/fixtures/http"
)

type HasConflictingNestedImports interface {
	RequesterNS
	Z() http.MyStruct
}
