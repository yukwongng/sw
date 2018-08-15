// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package authApiServer is a auto generated package.
Input file: auth.proto
*/
package auth

import (
	"reflect"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/venice/utils/runtime"
)

var typesMapAuth = map[string]*api.Struct{

	"auth.AuthenticationPolicy": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(AuthenticationPolicy{}) },
		Fields: map[string]api.Field{
			"TypeMeta": api.Field{Name: "TypeMeta", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Map: false, Inline: true, FromInline: false, KeyType: "", Type: "api.TypeMeta"},

			"Kind": api.Field{Name: "Kind", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "kind", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"APIVersion": api.Field{Name: "APIVersion", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "api-version", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"O": api.Field{Name: "O", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "meta", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.ObjectMeta"},

			"Spec": api.Field{Name: "Spec", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "spec", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "auth.AuthenticationPolicySpec"},

			"Status": api.Field{Name: "Status", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "status", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "auth.AuthenticationPolicyStatus"},
		},

		CLITags: map[string]api.CLIInfo{
			"api-version":                   api.CLIInfo{Path: "APIVersion", Skip: false, Insert: "", Help: ""},
			"auth-method":                   api.CLIInfo{Path: "Spec.Authenticators.Radius.Servers[].AuthMethod", Skip: false, Insert: "", Help: ""},
			"authenticator-order":           api.CLIInfo{Path: "Spec.Authenticators.AuthenticatorOrder", Skip: false, Insert: "", Help: ""},
			"base-dn":                       api.CLIInfo{Path: "Spec.Authenticators.Ldap.BaseDN", Skip: false, Insert: "", Help: ""},
			"bind-dn":                       api.CLIInfo{Path: "Spec.Authenticators.Ldap.BindDN", Skip: false, Insert: "", Help: ""},
			"bind-password":                 api.CLIInfo{Path: "Spec.Authenticators.Ldap.BindPassword", Skip: false, Insert: "", Help: ""},
			"email":                         api.CLIInfo{Path: "Spec.Authenticators.Ldap.AttributeMapping.Email", Skip: false, Insert: "", Help: ""},
			"enabled":                       api.CLIInfo{Path: "Spec.Authenticators.Radius.Enabled", Skip: false, Insert: "", Help: ""},
			"fullname":                      api.CLIInfo{Path: "Spec.Authenticators.Ldap.AttributeMapping.Fullname", Skip: false, Insert: "", Help: ""},
			"group":                         api.CLIInfo{Path: "Spec.Authenticators.Ldap.AttributeMapping.Group", Skip: false, Insert: "", Help: ""},
			"group-object-class":            api.CLIInfo{Path: "Spec.Authenticators.Ldap.AttributeMapping.GroupObjectClass", Skip: false, Insert: "", Help: ""},
			"kind":                          api.CLIInfo{Path: "Kind", Skip: false, Insert: "", Help: ""},
			"nas-id":                        api.CLIInfo{Path: "Spec.Authenticators.Radius.NasID", Skip: false, Insert: "", Help: ""},
			"secret":                        api.CLIInfo{Path: "Spec.Secret", Skip: false, Insert: "", Help: ""},
			"server-name":                   api.CLIInfo{Path: "Spec.Authenticators.Ldap.Servers[].TLSOptions.ServerName", Skip: false, Insert: "", Help: ""},
			"skip-server-cert-verification": api.CLIInfo{Path: "Spec.Authenticators.Ldap.Servers[].TLSOptions.SkipServerCertVerification", Skip: false, Insert: "", Help: ""},
			"start-tls":                     api.CLIInfo{Path: "Spec.Authenticators.Ldap.Servers[].TLSOptions.StartTLS", Skip: false, Insert: "", Help: ""},
			"tenant":                        api.CLIInfo{Path: "Spec.Authenticators.Ldap.AttributeMapping.Tenant", Skip: false, Insert: "", Help: ""},
			"trusted-certs":                 api.CLIInfo{Path: "Spec.Authenticators.Radius.Servers[].TrustedCerts", Skip: false, Insert: "", Help: ""},
			"url":                           api.CLIInfo{Path: "Spec.Authenticators.Radius.Servers[].Url", Skip: false, Insert: "", Help: ""},
			"user":                          api.CLIInfo{Path: "Spec.Authenticators.Ldap.AttributeMapping.User", Skip: false, Insert: "", Help: ""},
			"user-object-class":             api.CLIInfo{Path: "Spec.Authenticators.Ldap.AttributeMapping.UserObjectClass", Skip: false, Insert: "", Help: ""},
		},
	},
	"auth.AuthenticationPolicySpec": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(AuthenticationPolicySpec{}) },
		Fields: map[string]api.Field{
			"Authenticators": api.Field{Name: "Authenticators", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "authenticators", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "auth.Authenticators"},

			"Secret": api.Field{Name: "Secret", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "secret", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_BYTES"},
		},
	},
	"auth.AuthenticationPolicyStatus": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(AuthenticationPolicyStatus{}) },
		Fields:    map[string]api.Field{},
	},
	"auth.Authenticators": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(Authenticators{}) },
		Fields: map[string]api.Field{
			"AuthenticatorOrder": api.Field{Name: "AuthenticatorOrder", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "authenticator-order", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Ldap": api.Field{Name: "Ldap", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "ldap", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "auth.Ldap"},

			"Local": api.Field{Name: "Local", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "local", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "auth.Local"},

			"Radius": api.Field{Name: "Radius", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "radius", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "auth.Radius"},
		},
	},
	"auth.Ldap": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(Ldap{}) },
		Fields: map[string]api.Field{
			"Enabled": api.Field{Name: "Enabled", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "enabled", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_BOOL"},

			"BaseDN": api.Field{Name: "BaseDN", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "base-dn", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"BindDN": api.Field{Name: "BindDN", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "bind-dn", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"BindPassword": api.Field{Name: "BindPassword", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "bind-password", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"AttributeMapping": api.Field{Name: "AttributeMapping", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "attribute-mapping", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "auth.LdapAttributeMapping"},

			"Servers": api.Field{Name: "Servers", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "servers", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "auth.LdapServer"},
		},
	},
	"auth.LdapAttributeMapping": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(LdapAttributeMapping{}) },
		Fields: map[string]api.Field{
			"User": api.Field{Name: "User", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "user", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"UserObjectClass": api.Field{Name: "UserObjectClass", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "user-object-class", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Tenant": api.Field{Name: "Tenant", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "tenant", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Group": api.Field{Name: "Group", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "group", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"GroupObjectClass": api.Field{Name: "GroupObjectClass", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "group-object-class", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Email": api.Field{Name: "Email", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "email", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Fullname": api.Field{Name: "Fullname", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "fullname", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},
		},
	},
	"auth.LdapServer": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(LdapServer{}) },
		Fields: map[string]api.Field{
			"Url": api.Field{Name: "Url", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "url", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"TLSOptions": api.Field{Name: "TLSOptions", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "tls-options", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "auth.TLSOptions"},
		},
	},
	"auth.Local": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(Local{}) },
		Fields: map[string]api.Field{
			"Enabled": api.Field{Name: "Enabled", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "enabled", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_BOOL"},
		},
	},
	"auth.PasswordCredential": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(PasswordCredential{}) },
		Fields: map[string]api.Field{
			"Username": api.Field{Name: "Username", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "username", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Password": api.Field{Name: "Password", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "password", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Tenant": api.Field{Name: "Tenant", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "tenant", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},
		},
	},
	"auth.Permission": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(Permission{}) },
		Fields: map[string]api.Field{
			"ResourceTenant": api.Field{Name: "ResourceTenant", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "resource-tenant", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"ResourceGroup": api.Field{Name: "ResourceGroup", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "resource-group", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"ResourceKind": api.Field{Name: "ResourceKind", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "resource-kind", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"ResourceNamespace": api.Field{Name: "ResourceNamespace", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "resource-namespace", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"ResourceNames": api.Field{Name: "ResourceNames", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "resource-names", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Actions": api.Field{Name: "Actions", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "actions", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},
		},
	},
	"auth.Radius": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(Radius{}) },
		Fields: map[string]api.Field{
			"Enabled": api.Field{Name: "Enabled", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "enabled", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_BOOL"},

			"NasID": api.Field{Name: "NasID", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "nas-id", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Servers": api.Field{Name: "Servers", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "servers", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "auth.RadiusServer"},
		},
	},
	"auth.RadiusServer": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(RadiusServer{}) },
		Fields: map[string]api.Field{
			"Url": api.Field{Name: "Url", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "url", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Secret": api.Field{Name: "Secret", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "secret", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"AuthMethod": api.Field{Name: "AuthMethod", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "auth-method", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"TrustedCerts": api.Field{Name: "TrustedCerts", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "trusted-certs", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},
		},
	},
	"auth.Role": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(Role{}) },
		Fields: map[string]api.Field{
			"TypeMeta": api.Field{Name: "TypeMeta", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Map: false, Inline: true, FromInline: false, KeyType: "", Type: "api.TypeMeta"},

			"Kind": api.Field{Name: "Kind", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "kind", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"APIVersion": api.Field{Name: "APIVersion", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "api-version", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"O": api.Field{Name: "O", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "meta", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.ObjectMeta"},

			"Spec": api.Field{Name: "Spec", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "spec", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "auth.RoleSpec"},

			"Status": api.Field{Name: "Status", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "status", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "auth.RoleStatus"},
		},

		CLITags: map[string]api.CLIInfo{
			"actions":            api.CLIInfo{Path: "Spec.Permissions[].Actions", Skip: false, Insert: "", Help: ""},
			"api-version":        api.CLIInfo{Path: "APIVersion", Skip: false, Insert: "", Help: ""},
			"kind":               api.CLIInfo{Path: "Kind", Skip: false, Insert: "", Help: ""},
			"resource-group":     api.CLIInfo{Path: "Spec.Permissions[].ResourceGroup", Skip: false, Insert: "", Help: ""},
			"resource-kind":      api.CLIInfo{Path: "Spec.Permissions[].ResourceKind", Skip: false, Insert: "", Help: ""},
			"resource-names":     api.CLIInfo{Path: "Spec.Permissions[].ResourceNames", Skip: false, Insert: "", Help: ""},
			"resource-namespace": api.CLIInfo{Path: "Spec.Permissions[].ResourceNamespace", Skip: false, Insert: "", Help: ""},
			"resource-tenant":    api.CLIInfo{Path: "Spec.Permissions[].ResourceTenant", Skip: false, Insert: "", Help: ""},
		},
	},
	"auth.RoleBinding": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(RoleBinding{}) },
		Fields: map[string]api.Field{
			"TypeMeta": api.Field{Name: "TypeMeta", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Map: false, Inline: true, FromInline: false, KeyType: "", Type: "api.TypeMeta"},

			"Kind": api.Field{Name: "Kind", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "kind", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"APIVersion": api.Field{Name: "APIVersion", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "api-version", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"O": api.Field{Name: "O", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "meta", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.ObjectMeta"},

			"Spec": api.Field{Name: "Spec", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "spec", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "auth.RoleBindingSpec"},

			"Status": api.Field{Name: "Status", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "status", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "auth.RoleBindingStatus"},
		},

		CLITags: map[string]api.CLIInfo{
			"api-version": api.CLIInfo{Path: "APIVersion", Skip: false, Insert: "", Help: ""},
			"kind":        api.CLIInfo{Path: "Kind", Skip: false, Insert: "", Help: ""},
			"role":        api.CLIInfo{Path: "Spec.Role", Skip: false, Insert: "", Help: ""},
			"user-groups": api.CLIInfo{Path: "Spec.UserGroups", Skip: false, Insert: "", Help: ""},
			"users":       api.CLIInfo{Path: "Spec.Users", Skip: false, Insert: "", Help: ""},
		},
	},
	"auth.RoleBindingSpec": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(RoleBindingSpec{}) },
		Fields: map[string]api.Field{
			"Users": api.Field{Name: "Users", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "users", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"UserGroups": api.Field{Name: "UserGroups", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "user-groups", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Role": api.Field{Name: "Role", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "role", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},
		},
	},
	"auth.RoleBindingStatus": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(RoleBindingStatus{}) },
		Fields:    map[string]api.Field{},
	},
	"auth.RoleSpec": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(RoleSpec{}) },
		Fields: map[string]api.Field{
			"Permissions": api.Field{Name: "Permissions", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "permissions", Pointer: false, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "auth.Permission"},
		},
	},
	"auth.RoleStatus": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(RoleStatus{}) },
		Fields:    map[string]api.Field{},
	},
	"auth.TLSOptions": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(TLSOptions{}) },
		Fields: map[string]api.Field{
			"StartTLS": api.Field{Name: "StartTLS", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "start-tls", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_BOOL"},

			"SkipServerCertVerification": api.Field{Name: "SkipServerCertVerification", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "skip-server-cert-verification", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_BOOL"},

			"ServerName": api.Field{Name: "ServerName", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "server-name", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"TrustedCerts": api.Field{Name: "TrustedCerts", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "trusted-certs", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},
		},
	},
	"auth.User": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(User{}) },
		Fields: map[string]api.Field{
			"TypeMeta": api.Field{Name: "TypeMeta", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Map: false, Inline: true, FromInline: false, KeyType: "", Type: "api.TypeMeta"},

			"Kind": api.Field{Name: "Kind", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "kind", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"APIVersion": api.Field{Name: "APIVersion", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "api-version", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"O": api.Field{Name: "O", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "meta", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.ObjectMeta"},

			"Spec": api.Field{Name: "Spec", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "spec", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "auth.UserSpec"},

			"Status": api.Field{Name: "Status", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "status", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "auth.UserStatus"},
		},

		CLITags: map[string]api.CLIInfo{
			"api-version": api.CLIInfo{Path: "APIVersion", Skip: false, Insert: "", Help: ""},
			"email":       api.CLIInfo{Path: "Spec.Email", Skip: false, Insert: "", Help: ""},
			"fullname":    api.CLIInfo{Path: "Spec.Fullname", Skip: false, Insert: "", Help: ""},
			"kind":        api.CLIInfo{Path: "Kind", Skip: false, Insert: "", Help: ""},
			"password":    api.CLIInfo{Path: "Spec.Password", Skip: false, Insert: "", Help: ""},
			"roles":       api.CLIInfo{Path: "Status.Roles", Skip: false, Insert: "", Help: ""},
			"type":        api.CLIInfo{Path: "Spec.Type", Skip: false, Insert: "", Help: ""},
			"user-groups": api.CLIInfo{Path: "Status.UserGroups", Skip: false, Insert: "", Help: ""},
		},
	},
	"auth.UserSpec": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(UserSpec{}) },
		Fields: map[string]api.Field{
			"Fullname": api.Field{Name: "Fullname", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "fullname", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Email": api.Field{Name: "Email", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "email", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Password": api.Field{Name: "Password", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "password", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Type": api.Field{Name: "Type", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "type", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},
		},
	},
	"auth.UserStatus": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(UserStatus{}) },
		Fields: map[string]api.Field{
			"Roles": api.Field{Name: "Roles", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "roles", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"UserGroups": api.Field{Name: "UserGroups", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "user-groups", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"LastSuccessfulLogin": api.Field{Name: "LastSuccessfulLogin", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "last-successful-login", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.Timestamp"},
		},
	},
}

func init() {
	schema := runtime.GetDefaultScheme()
	schema.AddSchema(typesMapAuth)
}
