package appGwService

import (
	"github.com/GeertJohan/go.rice/embedded"
	"time"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "app.swagger.json",
		FileModTime: time.Unix(1505768515, 0),
		Content:     string("{\n  \"swagger\": \"2.0\",\n  \"info\": {\n    \"title\": \"Service name\",\n    \"version\": \"version not set\"\n  },\n  \"schemes\": [\n    \"http\",\n    \"https\"\n  ],\n  \"consumes\": [\n    \"application/json\"\n  ],\n  \"produces\": [\n    \"application/json\"\n  ],\n  \"paths\": {\n    \"/apps\": {\n      \"get\": {\n        \"operationId\": \"AutoListApp\",\n        \"responses\": {\n          \"200\": {\n            \"description\": \"\",\n            \"schema\": {\n              \"$ref\": \"#/definitions/appAppList\"\n            }\n          }\n        },\n        \"parameters\": [\n          {\n            \"name\": \"O.Name\",\n            \"description\": \"Name of the object, unique within a Namespace for scoped objects.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.Tenant\",\n            \"description\": \"Tenant is global namespace isolation for various objects. This can be automatically\\nfilled in many cases based on the tenant a user, who created the object, belongs go.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.Namespace\",\n            \"description\": \"Namespace of the object, for scoped objects.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.ResourceVersion\",\n            \"description\": \"Resource version in the object store. This can only be set by the server.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.UUID\",\n            \"description\": \"UUID is the unique identifier for the object. This can only be set by the server.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"LabelSelector\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"FieldSelector\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"PrefixWatch\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"boolean\",\n            \"format\": \"boolean\"\n          }\n        ],\n        \"tags\": [\n          \"AppV1\"\n        ]\n      }\n    },\n    \"/apps/{O.Name}\": {\n      \"get\": {\n        \"operationId\": \"AutoGetApp\",\n        \"responses\": {\n          \"200\": {\n            \"description\": \"\",\n            \"schema\": {\n              \"$ref\": \"#/definitions/appApp\"\n            }\n          }\n        },\n        \"parameters\": [\n          {\n            \"name\": \"O.Name\",\n            \"in\": \"path\",\n            \"required\": true,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"T.Kind\",\n            \"description\": \"Kind represents the type of the API object.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"T.APIVersion\",\n            \"description\": \"APIVersion defines the version of the API object.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.Tenant\",\n            \"description\": \"Tenant is global namespace isolation for various objects. This can be automatically\\nfilled in many cases based on the tenant a user, who created the object, belongs go.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.Namespace\",\n            \"description\": \"Namespace of the object, for scoped objects.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.ResourceVersion\",\n            \"description\": \"Resource version in the object store. This can only be set by the server.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.UUID\",\n            \"description\": \"UUID is the unique identifier for the object. This can only be set by the server.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"Status.Description\",\n            \"description\": \"Name is part of the meta\\nDescription of the AppUser - example \\\"cassandra user\\\".\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          }\n        ],\n        \"tags\": [\n          \"AppV1\"\n        ]\n      }\n    },\n    \"/{O.Tenant}/app-user-groups/{O.Name}\": {\n      \"get\": {\n        \"operationId\": \"AutoGetAppUserGrp\",\n        \"responses\": {\n          \"200\": {\n            \"description\": \"\",\n            \"schema\": {\n              \"$ref\": \"#/definitions/appAppUserGrp\"\n            }\n          }\n        },\n        \"parameters\": [\n          {\n            \"name\": \"O.Tenant\",\n            \"in\": \"path\",\n            \"required\": true,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.Name\",\n            \"in\": \"path\",\n            \"required\": true,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"T.Kind\",\n            \"description\": \"Kind represents the type of the API object.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"T.APIVersion\",\n            \"description\": \"APIVersion defines the version of the API object.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.Namespace\",\n            \"description\": \"Namespace of the object, for scoped objects.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.ResourceVersion\",\n            \"description\": \"Resource version in the object store. This can only be set by the server.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.UUID\",\n            \"description\": \"UUID is the unique identifier for the object. This can only be set by the server.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"Spec.AppUsers\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"array\",\n            \"items\": {\n              \"type\": \"string\"\n            }\n          },\n          {\n            \"name\": \"Spec.Description\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          }\n        ],\n        \"tags\": [\n          \"AppV1\"\n        ]\n      },\n      \"delete\": {\n        \"operationId\": \"AutoDeleteAppUserGrp\",\n        \"responses\": {\n          \"200\": {\n            \"description\": \"\",\n            \"schema\": {\n              \"$ref\": \"#/definitions/appAppUserGrp\"\n            }\n          }\n        },\n        \"parameters\": [\n          {\n            \"name\": \"O.Tenant\",\n            \"in\": \"path\",\n            \"required\": true,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.Name\",\n            \"in\": \"path\",\n            \"required\": true,\n            \"type\": \"string\"\n          }\n        ],\n        \"tags\": [\n          \"AppV1\"\n        ]\n      },\n      \"put\": {\n        \"operationId\": \"AutoUpdateAppUserGrp\",\n        \"responses\": {\n          \"200\": {\n            \"description\": \"\",\n            \"schema\": {\n              \"$ref\": \"#/definitions/appAppUserGrp\"\n            }\n          }\n        },\n        \"parameters\": [\n          {\n            \"name\": \"O.Tenant\",\n            \"in\": \"path\",\n            \"required\": true,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.Name\",\n            \"in\": \"path\",\n            \"required\": true,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"body\",\n            \"in\": \"body\",\n            \"required\": true,\n            \"schema\": {\n              \"$ref\": \"#/definitions/appAppUserGrp\"\n            }\n          }\n        ],\n        \"tags\": [\n          \"AppV1\"\n        ]\n      }\n    },\n    \"/{O.Tenant}/app-users\": {\n      \"get\": {\n        \"operationId\": \"AutoListAppUser\",\n        \"responses\": {\n          \"200\": {\n            \"description\": \"\",\n            \"schema\": {\n              \"$ref\": \"#/definitions/appAppUserList\"\n            }\n          }\n        },\n        \"parameters\": [\n          {\n            \"name\": \"O.Tenant\",\n            \"in\": \"path\",\n            \"required\": true,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.Name\",\n            \"description\": \"Name of the object, unique within a Namespace for scoped objects.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.Namespace\",\n            \"description\": \"Namespace of the object, for scoped objects.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.ResourceVersion\",\n            \"description\": \"Resource version in the object store. This can only be set by the server.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.UUID\",\n            \"description\": \"UUID is the unique identifier for the object. This can only be set by the server.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"LabelSelector\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"FieldSelector\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"PrefixWatch\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"boolean\",\n            \"format\": \"boolean\"\n          }\n        ],\n        \"tags\": [\n          \"AppV1\"\n        ]\n      },\n      \"post\": {\n        \"operationId\": \"AutoAddAppUser\",\n        \"responses\": {\n          \"200\": {\n            \"description\": \"\",\n            \"schema\": {\n              \"$ref\": \"#/definitions/appAppUser\"\n            }\n          }\n        },\n        \"parameters\": [\n          {\n            \"name\": \"O.Tenant\",\n            \"in\": \"path\",\n            \"required\": true,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"body\",\n            \"in\": \"body\",\n            \"required\": true,\n            \"schema\": {\n              \"$ref\": \"#/definitions/appAppUser\"\n            }\n          }\n        ],\n        \"tags\": [\n          \"AppV1\"\n        ]\n      }\n    },\n    \"/{O.Tenant}/app-users-groups\": {\n      \"get\": {\n        \"operationId\": \"AutoListAppUserGrp\",\n        \"responses\": {\n          \"200\": {\n            \"description\": \"\",\n            \"schema\": {\n              \"$ref\": \"#/definitions/appAppUserGrpList\"\n            }\n          }\n        },\n        \"parameters\": [\n          {\n            \"name\": \"O.Tenant\",\n            \"in\": \"path\",\n            \"required\": true,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.Name\",\n            \"description\": \"Name of the object, unique within a Namespace for scoped objects.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.Namespace\",\n            \"description\": \"Namespace of the object, for scoped objects.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.ResourceVersion\",\n            \"description\": \"Resource version in the object store. This can only be set by the server.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.UUID\",\n            \"description\": \"UUID is the unique identifier for the object. This can only be set by the server.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"LabelSelector\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"FieldSelector\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"PrefixWatch\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"boolean\",\n            \"format\": \"boolean\"\n          }\n        ],\n        \"tags\": [\n          \"AppV1\"\n        ]\n      },\n      \"post\": {\n        \"operationId\": \"AutoAddAppUserGrp\",\n        \"responses\": {\n          \"200\": {\n            \"description\": \"\",\n            \"schema\": {\n              \"$ref\": \"#/definitions/appAppUserGrp\"\n            }\n          }\n        },\n        \"parameters\": [\n          {\n            \"name\": \"O.Tenant\",\n            \"in\": \"path\",\n            \"required\": true,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"body\",\n            \"in\": \"body\",\n            \"required\": true,\n            \"schema\": {\n              \"$ref\": \"#/definitions/appAppUserGrp\"\n            }\n          }\n        ],\n        \"tags\": [\n          \"AppV1\"\n        ]\n      }\n    },\n    \"/{O.Tenant}/app-users/{O.Name}\": {\n      \"get\": {\n        \"operationId\": \"AutoGetAppUser\",\n        \"responses\": {\n          \"200\": {\n            \"description\": \"\",\n            \"schema\": {\n              \"$ref\": \"#/definitions/appAppUser\"\n            }\n          }\n        },\n        \"parameters\": [\n          {\n            \"name\": \"O.Tenant\",\n            \"in\": \"path\",\n            \"required\": true,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.Name\",\n            \"in\": \"path\",\n            \"required\": true,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"T.Kind\",\n            \"description\": \"Kind represents the type of the API object.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"T.APIVersion\",\n            \"description\": \"APIVersion defines the version of the API object.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.Namespace\",\n            \"description\": \"Namespace of the object, for scoped objects.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.ResourceVersion\",\n            \"description\": \"Resource version in the object store. This can only be set by the server.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.UUID\",\n            \"description\": \"UUID is the unique identifier for the object. This can only be set by the server.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"Spec.Description\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          }\n        ],\n        \"tags\": [\n          \"AppV1\"\n        ]\n      },\n      \"delete\": {\n        \"operationId\": \"AutoDeleteAppUser\",\n        \"responses\": {\n          \"200\": {\n            \"description\": \"\",\n            \"schema\": {\n              \"$ref\": \"#/definitions/appAppUser\"\n            }\n          }\n        },\n        \"parameters\": [\n          {\n            \"name\": \"O.Tenant\",\n            \"in\": \"path\",\n            \"required\": true,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.Name\",\n            \"in\": \"path\",\n            \"required\": true,\n            \"type\": \"string\"\n          }\n        ],\n        \"tags\": [\n          \"AppV1\"\n        ]\n      },\n      \"put\": {\n        \"operationId\": \"AutoUpdateAppUser\",\n        \"responses\": {\n          \"200\": {\n            \"description\": \"\",\n            \"schema\": {\n              \"$ref\": \"#/definitions/appAppUser\"\n            }\n          }\n        },\n        \"parameters\": [\n          {\n            \"name\": \"O.Tenant\",\n            \"in\": \"path\",\n            \"required\": true,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.Name\",\n            \"in\": \"path\",\n            \"required\": true,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"body\",\n            \"in\": \"body\",\n            \"required\": true,\n            \"schema\": {\n              \"$ref\": \"#/definitions/appAppUser\"\n            }\n          }\n        ],\n        \"tags\": [\n          \"AppV1\"\n        ]\n      }\n    }\n  },\n  \"definitions\": {\n    \"apiListMeta\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"ResourceVersion\": {\n          \"type\": \"string\",\n          \"description\": \"Resource version of object store at the time of list generation.\"\n        }\n      },\n      \"description\": \"ListMeta contains the metadata for list of objects.\"\n    },\n    \"apiListWatchOptions\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"O\": {\n          \"$ref\": \"#/definitions/apiObjectMeta\"\n        },\n        \"LabelSelector\": {\n          \"type\": \"string\"\n        },\n        \"FieldSelector\": {\n          \"type\": \"string\"\n        },\n        \"PrefixWatch\": {\n          \"type\": \"boolean\",\n          \"format\": \"boolean\"\n        }\n      }\n    },\n    \"apiObjectMeta\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"Name\": {\n          \"type\": \"string\",\n          \"description\": \"Name of the object, unique within a Namespace for scoped objects.\"\n        },\n        \"Tenant\": {\n          \"type\": \"string\",\n          \"description\": \"Tenant is global namespace isolation for various objects. This can be automatically\\nfilled in many cases based on the tenant a user, who created the object, belongs go.\"\n        },\n        \"Namespace\": {\n          \"type\": \"string\",\n          \"description\": \"Namespace of the object, for scoped objects.\"\n        },\n        \"ResourceVersion\": {\n          \"type\": \"string\",\n          \"description\": \"Resource version in the object store. This can only be set by the server.\"\n        },\n        \"UUID\": {\n          \"type\": \"string\",\n          \"description\": \"UUID is the unique identifier for the object. This can only be set by the server.\"\n        },\n        \"Labels\": {\n          \"type\": \"object\",\n          \"additionalProperties\": {\n            \"type\": \"string\"\n          },\n          \"description\": \"Labels are arbitrary (key,value) pairs associated with any object.\"\n        }\n      },\n      \"description\": \"ObjectMeta contains metadata that all objects stored in kvstore must have.\"\n    },\n    \"apiTypeMeta\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"Kind\": {\n          \"type\": \"string\",\n          \"description\": \"Kind represents the type of the API object.\"\n        },\n        \"APIVersion\": {\n          \"type\": \"string\",\n          \"description\": \"APIVersion defines the version of the API object.\"\n        }\n      },\n      \"description\": \"TypeMeta contains the metadata about kind and version for all API objects.\"\n    },\n    \"appApp\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"T\": {\n          \"$ref\": \"#/definitions/apiTypeMeta\"\n        },\n        \"O\": {\n          \"$ref\": \"#/definitions/apiObjectMeta\"\n        },\n        \"Spec\": {\n          \"$ref\": \"#/definitions/appAppSpec\"\n        },\n        \"Status\": {\n          \"$ref\": \"#/definitions/appAppStatus\"\n        }\n      },\n      \"title\": \"AppSpec - spec part of App object\"\n    },\n    \"appAppList\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"T\": {\n          \"$ref\": \"#/definitions/apiTypeMeta\",\n          \"title\": \"Name is already part of meta\\nDescription of app - example “mongo https://www.mongodb.com/”\"\n        },\n        \"ListMeta\": {\n          \"$ref\": \"#/definitions/apiListMeta\"\n        },\n        \"Items\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"$ref\": \"#/definitions/appApp\"\n          }\n        }\n      },\n      \"title\": \"AppStatus - status part of App object\"\n    },\n    \"appAppSpec\": {\n      \"type\": \"object\",\n      \"title\": \"App - Read-only objects auto-created by Venice\\nOne object per App that can be identified by Naples\"\n    },\n    \"appAppStatus\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"Description\": {\n          \"type\": \"string\",\n          \"title\": \"Name is part of the meta\\nDescription of the AppUser - example \\\"cassandra user\\\"\"\n        }\n      },\n      \"title\": \"AppUserSpec - spec part of AppUser object\"\n    },\n    \"appAppUser\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"T\": {\n          \"$ref\": \"#/definitions/apiTypeMeta\"\n        },\n        \"O\": {\n          \"$ref\": \"#/definitions/apiObjectMeta\"\n        },\n        \"Spec\": {\n          \"$ref\": \"#/definitions/appAppUserSpec\"\n        },\n        \"Status\": {\n          \"$ref\": \"#/definitions/appAppUserStatus\"\n        }\n      },\n      \"title\": \"AppUserStatus - status part of AppUser object\"\n    },\n    \"appAppUserGrp\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"T\": {\n          \"$ref\": \"#/definitions/apiTypeMeta\"\n        },\n        \"O\": {\n          \"$ref\": \"#/definitions/apiObjectMeta\"\n        },\n        \"Spec\": {\n          \"$ref\": \"#/definitions/appAppUserGrpSpec\"\n        },\n        \"Status\": {\n          \"$ref\": \"#/definitions/appAppUserGrpStatus\"\n        }\n      },\n      \"title\": \"AppUser is derived from application payload such as database login or\\nother application authentication mechanisms. \\nCreated by Venice Admin or Tenant Admin\"\n    },\n    \"appAppUserGrpList\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"T\": {\n          \"$ref\": \"#/definitions/apiTypeMeta\",\n          \"title\": \"List of appusers - example [“john”, “george”]\"\n        },\n        \"ListMeta\": {\n          \"$ref\": \"#/definitions/apiListMeta\",\n          \"title\": \"Description of the AppUserGrp - example \\\"SQL Users\\\"\"\n        },\n        \"Items\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"$ref\": \"#/definitions/appAppUserGrp\"\n          }\n        }\n      },\n      \"title\": \"AppUserGrpSpec - spec part of AppUserGrp object\"\n    },\n    \"appAppUserGrpSpec\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"AppUsers\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"string\"\n          }\n        },\n        \"Description\": {\n          \"type\": \"string\"\n        }\n      },\n      \"title\": \"AppUserGrpStatus - status part of AppUserGrp object\"\n    },\n    \"appAppUserGrpStatus\": {\n      \"type\": \"object\",\n      \"title\": \"AppUserGrp - Object representing group of AppUsers\\nCreated by Venice Admin or Tenant Admin\"\n    },\n    \"appAppUserList\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"T\": {\n          \"$ref\": \"#/definitions/apiTypeMeta\"\n        },\n        \"ListMeta\": {\n          \"$ref\": \"#/definitions/apiListMeta\"\n        },\n        \"Items\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"$ref\": \"#/definitions/appAppUser\"\n          }\n        }\n      }\n    },\n    \"appAppUserSpec\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"Description\": {\n          \"type\": \"string\"\n        }\n      }\n    },\n    \"appAppUserStatus\": {\n      \"type\": \"object\"\n    },\n    \"appAutoMsgAppUserGrpWatchHelper\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"Type\": {\n          \"type\": \"string\"\n        },\n        \"Object\": {\n          \"$ref\": \"#/definitions/appAppUserGrp\"\n        }\n      }\n    },\n    \"appAutoMsgAppUserWatchHelper\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"Type\": {\n          \"type\": \"string\"\n        },\n        \"Object\": {\n          \"$ref\": \"#/definitions/appAppUser\"\n        }\n      }\n    },\n    \"appAutoMsgAppWatchHelper\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"Type\": {\n          \"type\": \"string\"\n        },\n        \"Object\": {\n          \"$ref\": \"#/definitions/appApp\"\n        }\n      }\n    }\n  }\n}\n"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1505768100, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // "app.swagger.json"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`../../../../../sw/api/generated/app/swagger`, &embedded.EmbeddedBox{
		Name: `../../../../../sw/api/generated/app/swagger`,
		Time: time.Unix(1505768100, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"app.swagger.json": file2,
		},
	})
}
