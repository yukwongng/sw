// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package bookstoreCliUtilsBackend is a auto generated package.
Input file: example.proto
*/
package cli

import (
	"context"
	"fmt"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/api/generated/apiclient"
	"github.com/pensando/sw/api/generated/bookstore"
	loginctx "github.com/pensando/sw/api/login/context"
	"github.com/pensando/sw/venice/cli/gen"
)

// CreateBookFlags specifies flags for Book create operation
var CreateBookFlags = []gen.CliFlag{
	{
		ID:     "ISBNId",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "IdProvider",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "Publisher",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "Terminate",
		Type:   "Bool",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "author",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "category",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "errata",
		Type:   "StringSlice",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "review",
		Type:   "StringSlice",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "year",
		Type:   "StringSlice",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
}

func removeBookOper(obj interface{}) error {
	if v, ok := obj.(*bookstore.Book); ok {
		v.UUID = ""
		v.ResourceVersion = ""
		v.CreationTime = api.Timestamp{}
		v.ModTime = api.Timestamp{}
		v.Status = bookstore.BookStatus{}
	}
	return nil
}

// CreateCouponFlags specifies flags for Coupon create operation
var CreateCouponFlags = []gen.CliFlag{
	{
		ID:     "DiscountCode",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
}

func removeCouponOper(obj interface{}) error {
	if v, ok := obj.(*bookstore.Coupon); ok {
		v.UUID = ""
		v.ResourceVersion = ""
		v.CreationTime = api.Timestamp{}
		v.ModTime = api.Timestamp{}
		v.Status = bookstore.CouponStatus{}
	}
	return nil
}

// CreateCustomerFlags specifies flags for Customer create operation
var CreateCustomerFlags = []gen.CliFlag{
	{
		ID:     "Address",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "CreditCardNumbers",
		Type:   "StringSlice",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "Password",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "date-of-birth",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "id",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "lucky-numbers",
		Type:   "StringSlice",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "mother-maiden-name",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "ssn",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
}

func removeCustomerOper(obj interface{}) error {
	if v, ok := obj.(*bookstore.Customer); ok {
		v.UUID = ""
		v.ResourceVersion = ""
		v.CreationTime = api.Timestamp{}
		v.ModTime = api.Timestamp{}
		v.Status = bookstore.CustomerStatus{}
	}
	return nil
}

// CreateOrderFlags specifies flags for Order create operation
var CreateOrderFlags = []gen.CliFlag{
	{
		ID:     "Id",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
}

func removeOrderOper(obj interface{}) error {
	if v, ok := obj.(*bookstore.Order); ok {
		v.UUID = ""
		v.ResourceVersion = ""
		v.CreationTime = api.Timestamp{}
		v.ModTime = api.Timestamp{}
		v.Status = bookstore.OrderStatus{}
	}
	return nil
}

// CreatePublisherFlags specifies flags for Publisher create operation
var CreatePublisherFlags = []gen.CliFlag{
	{
		ID:     "WebAddr",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "address",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "id",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
}

func removePublisherOper(obj interface{}) error {
	if v, ok := obj.(*bookstore.Publisher); ok {
		v.UUID = ""
		v.ResourceVersion = ""
		v.CreationTime = api.Timestamp{}
		v.ModTime = api.Timestamp{}
		v.Status = bookstore.PublisherStatus{}
	}
	return nil
}

// CreateStoreFlags specifies flags for Store create operation
var CreateStoreFlags = []gen.CliFlag{
	{
		ID:     "Contact",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
}

func removeStoreOper(obj interface{}) error {
	if v, ok := obj.(*bookstore.Store); ok {
		v.UUID = ""
		v.ResourceVersion = ""
		v.CreationTime = api.Timestamp{}
		v.ModTime = api.Timestamp{}
		v.Status = bookstore.StoreStatus{}
	}
	return nil
}

func init() {
	cl := gen.GetInfo()

	cl.AddCliInfo("bookstore.Book", "create", CreateBookFlags)
	cl.AddRemoveObjOperFunc("bookstore.Book", removeBookOper)

	cl.AddCliInfo("bookstore.Coupon", "create", CreateCouponFlags)
	cl.AddRemoveObjOperFunc("bookstore.Coupon", removeCouponOper)

	cl.AddCliInfo("bookstore.Customer", "create", CreateCustomerFlags)
	cl.AddRemoveObjOperFunc("bookstore.Customer", removeCustomerOper)

	cl.AddCliInfo("bookstore.Order", "create", CreateOrderFlags)
	cl.AddRemoveObjOperFunc("bookstore.Order", removeOrderOper)

	cl.AddCliInfo("bookstore.Publisher", "create", CreatePublisherFlags)
	cl.AddRemoveObjOperFunc("bookstore.Publisher", removePublisherOper)

	cl.AddCliInfo("bookstore.Store", "create", CreateStoreFlags)
	cl.AddRemoveObjOperFunc("bookstore.Store", removeStoreOper)

}

func restGetOrder(hostname, tenant, token string, obj interface{}) error {

	restcl, err := apiclient.NewRestAPIClient(hostname)
	if err != nil {
		return fmt.Errorf("cannot create REST client")
	}
	defer restcl.Close()
	loginCtx := loginctx.NewContextWithAuthzHeader(context.Background(), "Bearer "+token)

	if v, ok := obj.(*bookstore.Order); ok {
		nv, err := restcl.BookstoreV1().Order().Get(loginCtx, &v.ObjectMeta)
		if err != nil {
			return err
		}
		*v = *nv
	}

	if v, ok := obj.(*bookstore.OrderList); ok {
		opts := api.ListWatchOptions{ObjectMeta: api.ObjectMeta{Tenant: tenant}}
		nv, err := restcl.BookstoreV1().Order().List(loginCtx, &opts)
		if err != nil {
			return err
		}
		v.Items = nv
	}
	return nil

}

func restDeleteOrder(hostname, token string, obj interface{}) error {

	restcl, err := apiclient.NewRestAPIClient(hostname)
	if err != nil {
		return fmt.Errorf("cannot create REST client")
	}
	defer restcl.Close()
	loginCtx := loginctx.NewContextWithAuthzHeader(context.Background(), "Bearer "+token)

	if v, ok := obj.(*bookstore.Order); ok {
		nv, err := restcl.BookstoreV1().Order().Delete(loginCtx, &v.ObjectMeta)
		if err != nil {
			return err
		}
		*v = *nv
	}
	return nil

}

func restPostOrder(hostname, token string, obj interface{}) error {

	restcl, err := apiclient.NewRestAPIClient(hostname)
	if err != nil {
		return fmt.Errorf("cannot create REST client")
	}
	defer restcl.Close()
	loginCtx := loginctx.NewContextWithAuthzHeader(context.Background(), "Bearer "+token)

	if v, ok := obj.(*bookstore.Order); ok {
		nv, err := restcl.BookstoreV1().Order().Create(loginCtx, v)
		if err != nil {
			return err
		}
		*v = *nv
	}
	return nil

}

func restPutOrder(hostname, token string, obj interface{}) error {

	restcl, err := apiclient.NewRestAPIClient(hostname)
	if err != nil {
		return fmt.Errorf("cannot create REST client")
	}
	defer restcl.Close()
	loginCtx := loginctx.NewContextWithAuthzHeader(context.Background(), "Bearer "+token)

	if v, ok := obj.(*bookstore.Order); ok {
		nv, err := restcl.BookstoreV1().Order().Update(loginCtx, v)
		if err != nil {
			return err
		}
		*v = *nv
	}
	return nil

}

func restGetBook(hostname, tenant, token string, obj interface{}) error {

	restcl, err := apiclient.NewRestAPIClient(hostname)
	if err != nil {
		return fmt.Errorf("cannot create REST client")
	}
	defer restcl.Close()
	loginCtx := loginctx.NewContextWithAuthzHeader(context.Background(), "Bearer "+token)

	if v, ok := obj.(*bookstore.Book); ok {
		nv, err := restcl.BookstoreV1().Book().Get(loginCtx, &v.ObjectMeta)
		if err != nil {
			return err
		}
		*v = *nv
	}

	if v, ok := obj.(*bookstore.BookList); ok {
		opts := api.ListWatchOptions{ObjectMeta: api.ObjectMeta{Tenant: tenant}}
		nv, err := restcl.BookstoreV1().Book().List(loginCtx, &opts)
		if err != nil {
			return err
		}
		v.Items = nv
	}
	return nil

}

func restDeleteBook(hostname, token string, obj interface{}) error {
	return fmt.Errorf("delete operation not supported for Book object")
}

func restPostBook(hostname, token string, obj interface{}) error {
	return fmt.Errorf("create operation not supported for Book object")
}

func restPutBook(hostname, token string, obj interface{}) error {

	restcl, err := apiclient.NewRestAPIClient(hostname)
	if err != nil {
		return fmt.Errorf("cannot create REST client")
	}
	defer restcl.Close()
	loginCtx := loginctx.NewContextWithAuthzHeader(context.Background(), "Bearer "+token)

	if v, ok := obj.(*bookstore.Book); ok {
		nv, err := restcl.BookstoreV1().Book().Update(loginCtx, v)
		if err != nil {
			return err
		}
		*v = *nv
	}
	return nil

}

func restGetPublisher(hostname, tenant, token string, obj interface{}) error {
	return fmt.Errorf("get operation not supported for Publisher object")
}

func restDeletePublisher(hostname, token string, obj interface{}) error {
	return fmt.Errorf("delete operation not supported for Publisher object")
}

func restPostPublisher(hostname, token string, obj interface{}) error {
	return fmt.Errorf("create operation not supported for Publisher object")
}

func restPutPublisher(hostname, token string, obj interface{}) error {
	return fmt.Errorf("put operation not supported for Publisher object")
}

func restGetStore(hostname, tenant, token string, obj interface{}) error {

	restcl, err := apiclient.NewRestAPIClient(hostname)
	if err != nil {
		return fmt.Errorf("cannot create REST client")
	}
	defer restcl.Close()
	loginCtx := loginctx.NewContextWithAuthzHeader(context.Background(), "Bearer "+token)

	if v, ok := obj.(*bookstore.Store); ok {
		nv, err := restcl.BookstoreV1().Store().Get(loginCtx, &v.ObjectMeta)
		if err != nil {
			return err
		}
		*v = *nv
	}

	if v, ok := obj.(*bookstore.StoreList); ok {
		objMeta := api.ObjectMeta{}
		nv, err := restcl.BookstoreV1().Store().Get(loginCtx, &objMeta)
		if err != nil {
			return err
		}
		v.Items = append(v.Items, nv)
	}
	return nil

}

func restDeleteStore(hostname, token string, obj interface{}) error {

	restcl, err := apiclient.NewRestAPIClient(hostname)
	if err != nil {
		return fmt.Errorf("cannot create REST client")
	}
	defer restcl.Close()
	loginCtx := loginctx.NewContextWithAuthzHeader(context.Background(), "Bearer "+token)

	if v, ok := obj.(*bookstore.Store); ok {
		nv, err := restcl.BookstoreV1().Store().Delete(loginCtx, &v.ObjectMeta)
		if err != nil {
			return err
		}
		*v = *nv
	}
	return nil

}

func restPostStore(hostname, token string, obj interface{}) error {

	restcl, err := apiclient.NewRestAPIClient(hostname)
	if err != nil {
		return fmt.Errorf("cannot create REST client")
	}
	defer restcl.Close()
	loginCtx := loginctx.NewContextWithAuthzHeader(context.Background(), "Bearer "+token)

	if v, ok := obj.(*bookstore.Store); ok {
		nv, err := restcl.BookstoreV1().Store().Create(loginCtx, v)
		if err != nil {
			return err
		}
		*v = *nv
	}
	return nil

}

func restPutStore(hostname, token string, obj interface{}) error {

	restcl, err := apiclient.NewRestAPIClient(hostname)
	if err != nil {
		return fmt.Errorf("cannot create REST client")
	}
	defer restcl.Close()
	loginCtx := loginctx.NewContextWithAuthzHeader(context.Background(), "Bearer "+token)

	if v, ok := obj.(*bookstore.Store); ok {
		nv, err := restcl.BookstoreV1().Store().Update(loginCtx, v)
		if err != nil {
			return err
		}
		*v = *nv
	}
	return nil

}

func restGetCoupon(hostname, tenant, token string, obj interface{}) error {
	return fmt.Errorf("get operation not supported for Coupon object")
}

func restDeleteCoupon(hostname, token string, obj interface{}) error {
	return fmt.Errorf("delete operation not supported for Coupon object")
}

func restPostCoupon(hostname, token string, obj interface{}) error {
	return fmt.Errorf("create operation not supported for Coupon object")
}

func restPutCoupon(hostname, token string, obj interface{}) error {
	return fmt.Errorf("put operation not supported for Coupon object")
}

func restGetCustomer(hostname, tenant, token string, obj interface{}) error {

	restcl, err := apiclient.NewRestAPIClient(hostname)
	if err != nil {
		return fmt.Errorf("cannot create REST client")
	}
	defer restcl.Close()
	loginCtx := loginctx.NewContextWithAuthzHeader(context.Background(), "Bearer "+token)

	if v, ok := obj.(*bookstore.Customer); ok {
		nv, err := restcl.BookstoreV1().Customer().Get(loginCtx, &v.ObjectMeta)
		if err != nil {
			return err
		}
		*v = *nv
	}

	if v, ok := obj.(*bookstore.CustomerList); ok {
		opts := api.ListWatchOptions{ObjectMeta: api.ObjectMeta{Tenant: tenant}}
		nv, err := restcl.BookstoreV1().Customer().List(loginCtx, &opts)
		if err != nil {
			return err
		}
		v.Items = nv
	}
	return nil

}

func restDeleteCustomer(hostname, token string, obj interface{}) error {

	restcl, err := apiclient.NewRestAPIClient(hostname)
	if err != nil {
		return fmt.Errorf("cannot create REST client")
	}
	defer restcl.Close()
	loginCtx := loginctx.NewContextWithAuthzHeader(context.Background(), "Bearer "+token)

	if v, ok := obj.(*bookstore.Customer); ok {
		nv, err := restcl.BookstoreV1().Customer().Delete(loginCtx, &v.ObjectMeta)
		if err != nil {
			return err
		}
		*v = *nv
	}
	return nil

}

func restPostCustomer(hostname, token string, obj interface{}) error {

	restcl, err := apiclient.NewRestAPIClient(hostname)
	if err != nil {
		return fmt.Errorf("cannot create REST client")
	}
	defer restcl.Close()
	loginCtx := loginctx.NewContextWithAuthzHeader(context.Background(), "Bearer "+token)

	if v, ok := obj.(*bookstore.Customer); ok {
		nv, err := restcl.BookstoreV1().Customer().Create(loginCtx, v)
		if err != nil {
			return err
		}
		*v = *nv
	}
	return nil

}

func restPutCustomer(hostname, token string, obj interface{}) error {

	restcl, err := apiclient.NewRestAPIClient(hostname)
	if err != nil {
		return fmt.Errorf("cannot create REST client")
	}
	defer restcl.Close()
	loginCtx := loginctx.NewContextWithAuthzHeader(context.Background(), "Bearer "+token)

	if v, ok := obj.(*bookstore.Customer); ok {
		nv, err := restcl.BookstoreV1().Customer().Update(loginCtx, v)
		if err != nil {
			return err
		}
		*v = *nv
	}
	return nil

}

func init() {
	cl := gen.GetInfo()
	if cl == nil {
		return
	}

	cl.AddRestPostFunc("bookstore.Order", "v1", restPostOrder)
	cl.AddRestDeleteFunc("bookstore.Order", "v1", restDeleteOrder)
	cl.AddRestPutFunc("bookstore.Order", "v1", restPutOrder)
	cl.AddRestGetFunc("bookstore.Order", "v1", restGetOrder)

	cl.AddRestPutFunc("bookstore.Book", "v1", restPutBook)
	cl.AddRestGetFunc("bookstore.Book", "v1", restGetBook)

	cl.AddRestPostFunc("bookstore.Store", "v1", restPostStore)
	cl.AddRestDeleteFunc("bookstore.Store", "v1", restDeleteStore)
	cl.AddRestPutFunc("bookstore.Store", "v1", restPutStore)
	cl.AddRestGetFunc("bookstore.Store", "v1", restGetStore)

	cl.AddRestPostFunc("bookstore.Customer", "v1", restPostCustomer)
	cl.AddRestDeleteFunc("bookstore.Customer", "v1", restDeleteCustomer)
	cl.AddRestPutFunc("bookstore.Customer", "v1", restPutCustomer)
	cl.AddRestGetFunc("bookstore.Customer", "v1", restGetCustomer)

}
