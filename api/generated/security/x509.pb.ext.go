// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package security is a auto generated package.
Input file: x509.proto
*/
package security

import (
	"errors"
	fmt "fmt"

	listerwatcher "github.com/pensando/sw/api/listerwatcher"
	"github.com/pensando/sw/venice/utils/kvstore"
	"github.com/pensando/sw/venice/utils/log"
	"github.com/pensando/sw/venice/utils/runtime"

	validators "github.com/pensando/sw/venice/utils/apigen/validators"

	"github.com/pensando/sw/venice/globals"
)

// Dummy definitions to suppress nonused warnings
var _ kvstore.Interface
var _ log.Logger
var _ listerwatcher.WatcherClient

var _ validators.DummyVar
var validatorMapX509 = make(map[string]map[string][]func(string, interface{}) error)

// MakeKey generates a KV store key for the object
func (m *Certificate) MakeKey(prefix string) string {
	return fmt.Sprint(globals.RootPrefix, "/", prefix, "/", "certificates/", m.Tenant, "/", m.Name)
}

func (m *Certificate) MakeURI(cat, ver, prefix string) string {
	in := m
	return fmt.Sprint("/", cat, "/", prefix, "/", ver, "/tenant/", in.Tenant, "/certificates/", in.Name)
}

// Clone clones the object into into or creates one of into is nil
func (m *Certificate) Clone(into interface{}) (interface{}, error) {
	var out *Certificate
	var ok bool
	if into == nil {
		out = &Certificate{}
	} else {
		out, ok = into.(*Certificate)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *Certificate) Defaults(ver string) bool {
	m.Kind = "Certificate"
	var ret bool
	ret = m.Spec.Defaults(ver) || ret
	ret = m.Status.Defaults(ver) || ret
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *CertificateSpec) Clone(into interface{}) (interface{}, error) {
	var out *CertificateSpec
	var ok bool
	if into == nil {
		out = &CertificateSpec{}
	} else {
		out, ok = into.(*CertificateSpec)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *CertificateSpec) Defaults(ver string) bool {
	var ret bool
	ret = true
	switch ver {
	default:
		for k := range m.Usages {
			m.Usages[k] = "Server"
		}
	}
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *CertificateStatus) Clone(into interface{}) (interface{}, error) {
	var out *CertificateStatus
	var ok bool
	if into == nil {
		out = &CertificateStatus{}
	} else {
		out, ok = into.(*CertificateStatus)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *CertificateStatus) Defaults(ver string) bool {
	var ret bool
	ret = true
	switch ver {
	default:
		m.Validity = "Unknown"
	}
	return ret
}

// Validators

func (m *Certificate) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error

	dlmtr := "."
	if path == "" {
		dlmtr = ""
	}
	npath := path + dlmtr + "Spec"
	if errs := m.Spec.Validate(ver, npath, ignoreStatus); errs != nil {
		ret = append(ret, errs...)
	}
	if !ignoreStatus {

		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := path + dlmtr + "Status"
		if errs := m.Status.Validate(ver, npath, ignoreStatus); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *CertificateSpec) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	if vs, ok := validatorMapX509["CertificateSpec"][ver]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	} else if vs, ok := validatorMapX509["CertificateSpec"]["all"]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	}
	return ret
}

func (m *CertificateStatus) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	if vs, ok := validatorMapX509["CertificateStatus"][ver]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	} else if vs, ok := validatorMapX509["CertificateStatus"]["all"]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	}
	return ret
}

// Transformers

func init() {
	scheme := runtime.GetDefaultScheme()
	scheme.AddKnownTypes(
		&Certificate{},
	)

	validatorMapX509 = make(map[string]map[string][]func(string, interface{}) error)

	validatorMapX509["CertificateSpec"] = make(map[string][]func(string, interface{}) error)
	validatorMapX509["CertificateSpec"]["all"] = append(validatorMapX509["CertificateSpec"]["all"], func(path string, i interface{}) error {
		m := i.(*CertificateSpec)

		for k, v := range m.Usages {
			if _, ok := CertificateSpec_UsageValues_value[v]; !ok {
				return fmt.Errorf("%v[%v] did not match allowed strings", path+"."+"Usages", k)
			}
		}
		return nil
	})

	validatorMapX509["CertificateStatus"] = make(map[string][]func(string, interface{}) error)
	validatorMapX509["CertificateStatus"]["all"] = append(validatorMapX509["CertificateStatus"]["all"], func(path string, i interface{}) error {
		m := i.(*CertificateStatus)

		if _, ok := CertificateStatus_ValidityValues_value[m.Validity]; !ok {
			return errors.New("CertificateStatus.Validity did not match allowed strings")
		}
		return nil
	})

}
