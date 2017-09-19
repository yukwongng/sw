// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package bookstore is a auto generated package.
Input file: protos/example.proto
*/
package bookstore

import (
	fmt "fmt"
	listerwatcher "github.com/pensando/sw/api/listerwatcher"
	"github.com/pensando/sw/venice/utils/kvstore"
	"github.com/pensando/sw/venice/utils/log"
)

// Dummy definitions to suppress nonused warnings
var _ kvstore.Interface
var _ log.Logger
var _ listerwatcher.WatcherClient

// MakeKey generates a KV store key for the object
func (m *Book) MakeKey(prefix string) string {
	return fmt.Sprint("/venice/", prefix, "/", "books/", m.Name)
}

// MakeKey generates a KV store key for the object
func (m *Order) MakeKey(prefix string) string {
	return fmt.Sprint("/venice/", prefix, "/", "orders/", m.Name)
}

// MakeKey generates a KV store key for the object
func (m *Publisher) MakeKey(prefix string) string {
	return fmt.Sprint("/venice/", prefix, "/", "publishers/", m.Name)
}

// MakeKey generates a KV store key for the object
func (m *BookList) MakeKey(prefix string) string {
	obj := Book{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *OrderList) MakeKey(prefix string) string {
	obj := Order{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *PublisherList) MakeKey(prefix string) string {
	obj := Publisher{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *AutoMsgBookWatchHelper) MakeKey(prefix string) string {
	obj := Book{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *AutoMsgOrderWatchHelper) MakeKey(prefix string) string {
	obj := Order{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *AutoMsgPublisherWatchHelper) MakeKey(prefix string) string {
	obj := Publisher{}
	return obj.MakeKey(prefix)
}
