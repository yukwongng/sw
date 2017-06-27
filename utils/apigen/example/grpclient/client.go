// client is a example client using the generated client bindings.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"reflect"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/pensando/sw/api"
	"github.com/pensando/sw/api/generated/bookstore"
	bookstoreclient "github.com/pensando/sw/api/generated/bookstore/grpc/client"
	"github.com/pensando/sw/api/generated/network"
	netclient "github.com/pensando/sw/api/generated/network/grpc/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	var (
		grpcaddr = flag.String("grpc-server", "localhost:8082", "GRPC Port to listen on")
	)
	flag.Parse()

	ctx := context.Background()

	conn, err := grpc.Dial(*grpcaddr, grpc.WithInsecure(), grpc.WithTimeout(time.Second))
	if err != nil {
		fmt.Printf("Failed to connect to gRPC server [%s]\n", *grpcaddr)
		os.Exit(-1)
	}
	defer conn.Close()

	cl := bookstoreclient.NewBookstoreV1(conn)

	// Add a Publisher
	var pub bookstore.Publisher
	pub.Name = "SaharaPublishers"
	var pubspec bookstore.PublisherSpec
	pubspec.Id = "111"
	pubspec.Address = "#1 hilane, timbuktoo"
	pub.Spec = &pubspec

	fmt.Printf("Adding Publisher\n")
	ret, err := cl.AddPublisher(ctx, pub)
	if err != nil {
		fmt.Printf("failed to create publisher(%s)\n", err)
		os.Exit(-1)
	}
	if !reflect.DeepEqual(ret.Spec, pub.Spec) {
		fmt.Printf("updated object [Add] does not match \n\t[%+v]\n\t[%+v]\n", pub, ret)
		os.Exit(-1)
	}
	// Get the object.
	fmt.Printf("Retrieving Publisher\n")
	pub.Spec = nil
	ret, err = cl.GetPublisher(ctx, pub)
	if err != nil {
		fmt.Printf("failed to get publisher(%s)\n", err)
		os.Exit(-1)
	}
	//pub.Spec = &pubspec
	if !reflect.DeepEqual(ret.Spec, &pubspec) {
		fmt.Printf("updated object [GET] does not match \n\t[%+v]\n\t[%+v]\n", pub, ret)
		os.Exit(-1)
	}

	// Duplicate add of the object
	fmt.Printf("Try to Re-add Publisher\n")
	pub.Spec = &pubspec
	ret, err = cl.AddPublisher(ctx, pub)
	if err == nil {
		fmt.Printf("Was able to create duplicate publisher\n")
		os.Exit(-1)
	}

	// Update the Object
	fmt.Printf("Update Publisher\n")
	pubspec.Address = "#2 hilane, timbuktoo"
	pub.Spec = &pubspec
	ret, err = cl.UpdatePublisher(ctx, pub)
	if err != nil {
		fmt.Printf("failed to Update publisher(%s)\n", err)
		os.Exit(-1)
	}
	if !reflect.DeepEqual(ret.Spec, pub.Spec) {
		fmt.Printf("updated object [UPD] does not match \n\t[%+v]\n\t[%+v]\n", pub, ret)
		os.Exit(-1)
	}

	pub.Spec = nil
	ret, err = cl.GetPublisher(ctx, pub)
	if err != nil {
		fmt.Printf("failed to get publisher(%s)\n", err)
		os.Exit(-1)
	}
	if !reflect.DeepEqual(ret.Spec, &pubspec) {
		fmt.Printf("updated object[GET2] does not match \n\t[%+v]\n\t[%+v]\n", pub, ret)
		os.Exit(-1)
	}

	// Delete the Object
	fmt.Printf("Delete Publisher\n")
	pub.Spec = nil
	ret, err = cl.DeletePublisher(ctx, pub)
	if err != nil {
		fmt.Printf("failed to delete publisher(%s)\n", err)
		os.Exit(-1)
	}
	if !reflect.DeepEqual(ret.Spec, &pubspec) {
		fmt.Printf("Deleted object does not match \n\t[%+v]\n\t[%+v]\n", pubspec, ret.Spec)
		os.Exit(-1)
	}

	// tenant API tests
	tcl := netclient.NewTenantV1(conn)

	// Add a tenant
	tenant := network.Tenant{
		TypeMeta: api.TypeMeta{
			Kind:       "tenant",
			APIVersion: "v1",
		},
		ObjectMeta: api.ObjectMeta{
			Tenant: "tenant2",
			Name:   "tenant2",
		},
		Spec: network.TenantSpec{
			AdminUser: "admin",
		},
	}

	// create a new context with this metadata
	md := metadata.Pairs("req-method", "POST")
	ctx = metadata.NewContext(ctx, md)

	// perform a POST operation on tenant object
	tret, err := tcl.TenantOper(ctx, tenant)
	logrus.Infof("TenantOper returned: {%+v}, Err: %v", tret, err)
}
