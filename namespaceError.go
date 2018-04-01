package errors

import (
	"fmt"
)

var NamespaceErrorsPrintNamespace bool = false

func NewNamespace(namespace string, err error) error {
	return &namespaceError{ns: namespace, err: err}
}

type namespaceError struct {
	ns  string
	err error
}

func (e *namespaceError) Error() string {
	if NamespaceErrorsPrintNamespace {
		return fmt.Sprintf("%s: %s", e.ns, e.err)
	}
	return e.err.Error()
}
