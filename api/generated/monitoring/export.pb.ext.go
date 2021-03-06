// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package monitoring is a auto generated package.
Input file: export.proto
*/
package monitoring

import (
	"context"
	fmt "fmt"
	"strings"

	listerwatcher "github.com/pensando/sw/api/listerwatcher"
	"github.com/pensando/sw/venice/utils/kvstore"
	"github.com/pensando/sw/venice/utils/log"
	"github.com/pensando/sw/venice/utils/ref"

	"github.com/pensando/sw/api/interfaces"
	validators "github.com/pensando/sw/venice/utils/apigen/validators"
	"github.com/pensando/sw/venice/utils/runtime"
	"github.com/pensando/sw/venice/utils/transformers/storage"
)

// Dummy definitions to suppress nonused warnings
var _ kvstore.Interface
var _ log.Logger
var _ listerwatcher.WatcherClient

// AuthConfig_Algos_normal is a map of normalized values for the enum
var AuthConfig_Algos_normal = map[string]string{
	"md5":  "md5",
	"sha1": "sha1",
}

var AuthConfig_Algos_vname = map[int32]string{
	0: "md5",
	1: "sha1",
}

var AuthConfig_Algos_vvalue = map[string]int32{
	"md5":  0,
	"sha1": 1,
}

func (x AuthConfig_Algos) String() string {
	return AuthConfig_Algos_vname[int32(x)]
}

// PrivacyConfig_Algos_normal is a map of normalized values for the enum
var PrivacyConfig_Algos_normal = map[string]string{
	"aes128": "aes128",
	"des56":  "des56",
}

var PrivacyConfig_Algos_vname = map[int32]string{
	0: "des56",
	1: "aes128",
}

var PrivacyConfig_Algos_vvalue = map[string]int32{
	"des56":  0,
	"aes128": 1,
}

func (x PrivacyConfig_Algos) String() string {
	return PrivacyConfig_Algos_vname[int32(x)]
}

// SNMPTrapServer_SNMPVersions_normal is a map of normalized values for the enum
var SNMPTrapServer_SNMPVersions_normal = map[string]string{
	"v2c": "v2c",
	"v3":  "v3",
}

var SNMPTrapServer_SNMPVersions_vname = map[int32]string{
	0: "v2c",
	1: "v3",
}

var SNMPTrapServer_SNMPVersions_vvalue = map[string]int32{
	"v2c": 0,
	"v3":  1,
}

func (x SNMPTrapServer_SNMPVersions) String() string {
	return SNMPTrapServer_SNMPVersions_vname[int32(x)]
}

// ExportAuthType_normal is a map of normalized values for the enum
var ExportAuthType_normal = map[string]string{
	"certs":             "certs",
	"none":              "none",
	"token":             "token",
	"username-password": "username-password",
}

var ExportAuthType_vname = map[int32]string{
	0: "none",
	1: "username-password",
	2: "token",
	3: "certs",
}

var ExportAuthType_vvalue = map[string]int32{
	"none":              0,
	"username-password": 1,
	"token":             2,
	"certs":             3,
}

func (x ExportAuthType) String() string {
	return ExportAuthType_vname[int32(x)]
}

// SyslogFacility_normal is a map of normalized values for the enum
var SyslogFacility_normal = map[string]string{
	"auth":     "auth",
	"authpriv": "authpriv",
	"cron":     "cron",
	"daemon":   "daemon",
	"ftp":      "ftp",
	"kernel":   "kernel",
	"local0":   "local0",
	"local1":   "local1",
	"local2":   "local2",
	"local3":   "local3",
	"local4":   "local4",
	"local5":   "local5",
	"local6":   "local6",
	"local7":   "local7",
	"lpr":      "lpr",
	"mail":     "mail",
	"news":     "news",
	"syslog":   "syslog",
	"user":     "user",
	"uucp":     "uucp",
}

var SyslogFacility_vname = map[int32]string{
	0:   "kernel",
	8:   "user",
	16:  "mail",
	24:  "daemon",
	32:  "auth",
	40:  "syslog",
	48:  "lpr",
	56:  "news",
	64:  "uucp",
	72:  "cron",
	80:  "authpriv",
	88:  "ftp",
	128: "local0",
	136: "local1",
	144: "local2",
	152: "local3",
	160: "local4",
	168: "local5",
	176: "local6",
	184: "local7",
}

var SyslogFacility_vvalue = map[string]int32{
	"kernel":   0,
	"user":     8,
	"mail":     16,
	"daemon":   24,
	"auth":     32,
	"syslog":   40,
	"lpr":      48,
	"news":     56,
	"uucp":     64,
	"cron":     72,
	"authpriv": 80,
	"ftp":      88,
	"local0":   128,
	"local1":   136,
	"local2":   144,
	"local3":   152,
	"local4":   160,
	"local5":   168,
	"local6":   176,
	"local7":   184,
}

func (x SyslogFacility) String() string {
	return SyslogFacility_vname[int32(x)]
}

var _ validators.DummyVar
var validatorMapExport = make(map[string]map[string][]func(string, interface{}) error)

var storageTransformersMapExport = make(map[string][]func(ctx context.Context, i interface{}, toStorage bool) error)
var eraseSecretsMapExport = make(map[string]func(i interface{}))

// Clone clones the object into into or creates one of into is nil
func (m *AuthConfig) Clone(into interface{}) (interface{}, error) {
	var out *AuthConfig
	var ok bool
	if into == nil {
		out = &AuthConfig{}
	} else {
		out, ok = into.(*AuthConfig)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AuthConfig))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AuthConfig) Defaults(ver string) bool {
	var ret bool
	ret = true
	switch ver {
	default:
		m.Algo = "md5"
	}
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *ExportConfig) Clone(into interface{}) (interface{}, error) {
	var out *ExportConfig
	var ok bool
	if into == nil {
		out = &ExportConfig{}
	} else {
		out, ok = into.(*ExportConfig)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*ExportConfig))
	return out, nil
}

// Default sets up the defaults for the object
func (m *ExportConfig) Defaults(ver string) bool {
	var ret bool
	if m.Credentials != nil {
		ret = m.Credentials.Defaults(ver) || ret
	}
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *ExternalCred) Clone(into interface{}) (interface{}, error) {
	var out *ExternalCred
	var ok bool
	if into == nil {
		out = &ExternalCred{}
	} else {
		out, ok = into.(*ExternalCred)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*ExternalCred))
	return out, nil
}

// Default sets up the defaults for the object
func (m *ExternalCred) Defaults(ver string) bool {
	var ret bool
	ret = true
	switch ver {
	default:
		m.AuthType = "none"
	}
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *PSMExportTarget) Clone(into interface{}) (interface{}, error) {
	var out *PSMExportTarget
	var ok bool
	if into == nil {
		out = &PSMExportTarget{}
	} else {
		out, ok = into.(*PSMExportTarget)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*PSMExportTarget))
	return out, nil
}

// Default sets up the defaults for the object
func (m *PSMExportTarget) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *PrivacyConfig) Clone(into interface{}) (interface{}, error) {
	var out *PrivacyConfig
	var ok bool
	if into == nil {
		out = &PrivacyConfig{}
	} else {
		out, ok = into.(*PrivacyConfig)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*PrivacyConfig))
	return out, nil
}

// Default sets up the defaults for the object
func (m *PrivacyConfig) Defaults(ver string) bool {
	var ret bool
	ret = true
	switch ver {
	default:
		m.Algo = "des56"
	}
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *SNMPTrapServer) Clone(into interface{}) (interface{}, error) {
	var out *SNMPTrapServer
	var ok bool
	if into == nil {
		out = &SNMPTrapServer{}
	} else {
		out, ok = into.(*SNMPTrapServer)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*SNMPTrapServer))
	return out, nil
}

// Default sets up the defaults for the object
func (m *SNMPTrapServer) Defaults(ver string) bool {
	var ret bool
	if m.AuthConfig != nil {
		ret = m.AuthConfig.Defaults(ver) || ret
	}
	if m.PrivacyConfig != nil {
		ret = m.PrivacyConfig.Defaults(ver) || ret
	}
	ret = true
	switch ver {
	default:
		m.Port = "162"
		m.Version = "v2c"
	}
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *SyslogExportConfig) Clone(into interface{}) (interface{}, error) {
	var out *SyslogExportConfig
	var ok bool
	if into == nil {
		out = &SyslogExportConfig{}
	} else {
		out, ok = into.(*SyslogExportConfig)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*SyslogExportConfig))
	return out, nil
}

// Default sets up the defaults for the object
func (m *SyslogExportConfig) Defaults(ver string) bool {
	var ret bool
	ret = true
	switch ver {
	default:
		m.FacilityOverride = "user"
	}
	return ret
}

// Validators and Requirements

func (m *AuthConfig) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AuthConfig) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	if vs, ok := validatorMapExport["AuthConfig"][ver]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	} else if vs, ok := validatorMapExport["AuthConfig"]["all"]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	}
	return ret
}

func (m *AuthConfig) Normalize() {

	m.Algo = AuthConfig_Algos_normal[strings.ToLower(m.Algo)]

}

func (m *ExportConfig) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *ExportConfig) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error

	if m.Credentials != nil {
		{
			dlmtr := "."
			if path == "" {
				dlmtr = ""
			}
			npath := path + dlmtr + "Credentials"
			if errs := m.Credentials.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
				ret = append(ret, errs...)
			}
		}
	}
	if vs, ok := validatorMapExport["ExportConfig"][ver]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	} else if vs, ok := validatorMapExport["ExportConfig"]["all"]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	}
	return ret
}

func (m *ExportConfig) Normalize() {

	if m.Credentials != nil {
		m.Credentials.Normalize()
	}

}

func (m *ExternalCred) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *ExternalCred) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	if vs, ok := validatorMapExport["ExternalCred"][ver]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	} else if vs, ok := validatorMapExport["ExternalCred"]["all"]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	}
	return ret
}

func (m *ExternalCred) Normalize() {

	m.AuthType = ExportAuthType_normal[strings.ToLower(m.AuthType)]

}

func (m *PSMExportTarget) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *PSMExportTarget) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	return ret
}

func (m *PSMExportTarget) Normalize() {

}

func (m *PrivacyConfig) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *PrivacyConfig) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	if vs, ok := validatorMapExport["PrivacyConfig"][ver]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	} else if vs, ok := validatorMapExport["PrivacyConfig"]["all"]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	}
	return ret
}

func (m *PrivacyConfig) Normalize() {

	m.Algo = PrivacyConfig_Algos_normal[strings.ToLower(m.Algo)]

}

func (m *SNMPTrapServer) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *SNMPTrapServer) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error

	if m.AuthConfig != nil {
		{
			dlmtr := "."
			if path == "" {
				dlmtr = ""
			}
			npath := path + dlmtr + "AuthConfig"
			if errs := m.AuthConfig.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
				ret = append(ret, errs...)
			}
		}
	}

	if m.PrivacyConfig != nil {
		{
			dlmtr := "."
			if path == "" {
				dlmtr = ""
			}
			npath := path + dlmtr + "PrivacyConfig"
			if errs := m.PrivacyConfig.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
				ret = append(ret, errs...)
			}
		}
	}
	if vs, ok := validatorMapExport["SNMPTrapServer"][ver]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	} else if vs, ok := validatorMapExport["SNMPTrapServer"]["all"]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	}
	return ret
}

func (m *SNMPTrapServer) Normalize() {

	if m.AuthConfig != nil {
		m.AuthConfig.Normalize()
	}

	if m.PrivacyConfig != nil {
		m.PrivacyConfig.Normalize()
	}

	m.Version = SNMPTrapServer_SNMPVersions_normal[strings.ToLower(m.Version)]

}

func (m *SyslogExportConfig) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *SyslogExportConfig) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	if vs, ok := validatorMapExport["SyslogExportConfig"][ver]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	} else if vs, ok := validatorMapExport["SyslogExportConfig"]["all"]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	}
	return ret
}

func (m *SyslogExportConfig) Normalize() {

	m.FacilityOverride = SyslogFacility_normal[strings.ToLower(m.FacilityOverride)]

}

// Transformers

func (m *ExportConfig) ApplyStorageTransformer(ctx context.Context, toStorage bool) error {

	if m.Credentials != nil {
		if err := m.Credentials.ApplyStorageTransformer(ctx, toStorage); err != nil {
			return err
		}
	}

	return nil
}

func (m *ExportConfig) EraseSecrets() {

	if m.Credentials != nil {
		m.Credentials.EraseSecrets()
	}

	return
}

func (m *ExternalCred) ApplyStorageTransformer(ctx context.Context, toStorage bool) error {
	if vs, ok := storageTransformersMapExport["ExternalCred"]; ok {
		for _, v := range vs {
			if err := v(ctx, m, toStorage); err != nil {
				return err
			}
		}
	}
	return nil
}

func (m *ExternalCred) EraseSecrets() {
	if v, ok := eraseSecretsMapExport["ExternalCred"]; ok {
		v(m)
	}
	return
}

func init() {
	scheme := runtime.GetDefaultScheme()
	scheme.AddKnownTypes()

	validatorMapExport = make(map[string]map[string][]func(string, interface{}) error)

	validatorMapExport["AuthConfig"] = make(map[string][]func(string, interface{}) error)
	validatorMapExport["AuthConfig"]["all"] = append(validatorMapExport["AuthConfig"]["all"], func(path string, i interface{}) error {
		m := i.(*AuthConfig)

		if _, ok := AuthConfig_Algos_vvalue[m.Algo]; !ok {
			vals := []string{}
			for k1, _ := range AuthConfig_Algos_vvalue {
				vals = append(vals, k1)
			}
			return fmt.Errorf("%v did not match allowed strings %v", path+"."+"Algo", vals)
		}
		return nil
	})

	validatorMapExport["ExportConfig"] = make(map[string][]func(string, interface{}) error)
	validatorMapExport["ExportConfig"]["all"] = append(validatorMapExport["ExportConfig"]["all"], func(path string, i interface{}) error {
		m := i.(*ExportConfig)
		args := make([]string, 0)
		args = append(args, "1")
		args = append(args, "2048")

		if err := validators.StrLen(m.Destination, args); err != nil {
			return fmt.Errorf("%v failed validation: %s", path+"."+"Destination", err.Error())
		}
		return nil
	})

	validatorMapExport["ExportConfig"]["all"] = append(validatorMapExport["ExportConfig"]["all"], func(path string, i interface{}) error {
		m := i.(*ExportConfig)
		args := make([]string, 0)
		args = append(args, "0")
		args = append(args, "2048")

		if err := validators.StrLen(m.Gateway, args); err != nil {
			return fmt.Errorf("%v failed validation: %s", path+"."+"Gateway", err.Error())
		}
		return nil
	})

	validatorMapExport["ExportConfig"]["all"] = append(validatorMapExport["ExportConfig"]["all"], func(path string, i interface{}) error {
		m := i.(*ExportConfig)
		if err := validators.EmptyOr(validators.ProtoPort, m.Transport, nil); err != nil {
			return fmt.Errorf("%v failed validation: %s", path+"."+"Transport", err.Error())
		}
		return nil
	})

	validatorMapExport["ExternalCred"] = make(map[string][]func(string, interface{}) error)
	validatorMapExport["ExternalCred"]["all"] = append(validatorMapExport["ExternalCred"]["all"], func(path string, i interface{}) error {
		m := i.(*ExternalCred)

		if _, ok := ExportAuthType_vvalue[m.AuthType]; !ok {
			vals := []string{}
			for k1, _ := range ExportAuthType_vvalue {
				vals = append(vals, k1)
			}
			return fmt.Errorf("%v did not match allowed strings %v", path+"."+"AuthType", vals)
		}
		return nil
	})

	validatorMapExport["PrivacyConfig"] = make(map[string][]func(string, interface{}) error)
	validatorMapExport["PrivacyConfig"]["all"] = append(validatorMapExport["PrivacyConfig"]["all"], func(path string, i interface{}) error {
		m := i.(*PrivacyConfig)

		if _, ok := PrivacyConfig_Algos_vvalue[m.Algo]; !ok {
			vals := []string{}
			for k1, _ := range PrivacyConfig_Algos_vvalue {
				vals = append(vals, k1)
			}
			return fmt.Errorf("%v did not match allowed strings %v", path+"."+"Algo", vals)
		}
		return nil
	})

	validatorMapExport["SNMPTrapServer"] = make(map[string][]func(string, interface{}) error)
	validatorMapExport["SNMPTrapServer"]["all"] = append(validatorMapExport["SNMPTrapServer"]["all"], func(path string, i interface{}) error {
		m := i.(*SNMPTrapServer)

		if _, ok := SNMPTrapServer_SNMPVersions_vvalue[m.Version]; !ok {
			vals := []string{}
			for k1, _ := range SNMPTrapServer_SNMPVersions_vvalue {
				vals = append(vals, k1)
			}
			return fmt.Errorf("%v did not match allowed strings %v", path+"."+"Version", vals)
		}
		return nil
	})

	validatorMapExport["SyslogExportConfig"] = make(map[string][]func(string, interface{}) error)
	validatorMapExport["SyslogExportConfig"]["all"] = append(validatorMapExport["SyslogExportConfig"]["all"], func(path string, i interface{}) error {
		m := i.(*SyslogExportConfig)

		if _, ok := SyslogFacility_vvalue[m.FacilityOverride]; !ok {
			vals := []string{}
			for k1, _ := range SyslogFacility_vvalue {
				vals = append(vals, k1)
			}
			return fmt.Errorf("%v did not match allowed strings %v", path+"."+"FacilityOverride", vals)
		}
		return nil
	})

	{
		ExternalCredBearerTokenTx, err := storage.NewSecretValueTransformer()
		if err != nil {
			log.Fatalf("Error instantiating SecretStorageTransformer: %v", err)
		}
		storageTransformersMapExport["ExternalCred"] = append(storageTransformersMapExport["ExternalCred"],
			func(ctx context.Context, i interface{}, toStorage bool) error {
				var data []byte
				var err error
				m := i.(*ExternalCred)

				if toStorage {
					data, err = ExternalCredBearerTokenTx.TransformToStorage(ctx, []byte(m.BearerToken))
				} else {
					data, err = ExternalCredBearerTokenTx.TransformFromStorage(ctx, []byte(m.BearerToken))
				}
				m.BearerToken = string(data)

				return err
			})

		eraseSecretsMapExport["ExternalCred"] = func(i interface{}) {
			m := i.(*ExternalCred)

			var data []byte
			m.BearerToken = string(data)

			return
		}

	}

	{
		ExternalCredKeyDataTx, err := storage.NewSecretValueTransformer()
		if err != nil {
			log.Fatalf("Error instantiating SecretStorageTransformer: %v", err)
		}
		storageTransformersMapExport["ExternalCred"] = append(storageTransformersMapExport["ExternalCred"],
			func(ctx context.Context, i interface{}, toStorage bool) error {
				var data []byte
				var err error
				m := i.(*ExternalCred)

				if toStorage {
					data, err = ExternalCredKeyDataTx.TransformToStorage(ctx, []byte(m.KeyData))
				} else {
					data, err = ExternalCredKeyDataTx.TransformFromStorage(ctx, []byte(m.KeyData))
				}
				m.KeyData = string(data)

				return err
			})

		eraseSecretsMapExport["ExternalCred"] = func(i interface{}) {
			m := i.(*ExternalCred)

			var data []byte
			m.KeyData = string(data)

			return
		}

	}

	{
		ExternalCredPasswordTx, err := storage.NewSecretValueTransformer()
		if err != nil {
			log.Fatalf("Error instantiating SecretStorageTransformer: %v", err)
		}
		storageTransformersMapExport["ExternalCred"] = append(storageTransformersMapExport["ExternalCred"],
			func(ctx context.Context, i interface{}, toStorage bool) error {
				var data []byte
				var err error
				m := i.(*ExternalCred)

				if toStorage {
					data, err = ExternalCredPasswordTx.TransformToStorage(ctx, []byte(m.Password))
				} else {
					data, err = ExternalCredPasswordTx.TransformFromStorage(ctx, []byte(m.Password))
				}
				m.Password = string(data)

				return err
			})

		eraseSecretsMapExport["ExternalCred"] = func(i interface{}) {
			m := i.(*ExternalCred)

			var data []byte
			m.Password = string(data)

			return
		}

	}

}
