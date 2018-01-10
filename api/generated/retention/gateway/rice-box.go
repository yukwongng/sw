package retentionGwService

import (
	"github.com/GeertJohan/go.rice/embedded"
	"time"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "retention.swagger.json",
		FileModTime: time.Unix(1515619814, 0),
		Content:     string("{\n  \"swagger\": \"2.0\",\n  \"info\": {\n    \"title\": \"Service name\",\n    \"version\": \"version not set\"\n  },\n  \"schemes\": [\n    \"http\",\n    \"https\"\n  ],\n  \"consumes\": [\n    \"application/json\"\n  ],\n  \"produces\": [\n    \"application/json\"\n  ],\n  \"paths\": {\n    \"/{O.Tenant}/retentionPolicy\": {\n      \"post\": {\n        \"operationId\": \"AutoAddRetentionPolicy\",\n        \"responses\": {\n          \"200\": {\n            \"description\": \"\",\n            \"schema\": {\n              \"$ref\": \"#/definitions/retentionRetentionPolicy\"\n            }\n          }\n        },\n        \"parameters\": [\n          {\n            \"name\": \"O.Tenant\",\n            \"in\": \"path\",\n            \"required\": true,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"body\",\n            \"in\": \"body\",\n            \"required\": true,\n            \"schema\": {\n              \"$ref\": \"#/definitions/retentionRetentionPolicy\"\n            }\n          }\n        ],\n        \"tags\": [\n          \"RetentionPolicyV1\"\n        ]\n      }\n    },\n    \"/{O.Tenant}/retentionPolicy/{O.Name}\": {\n      \"get\": {\n        \"operationId\": \"AutoGetRetentionPolicy\",\n        \"responses\": {\n          \"200\": {\n            \"description\": \"\",\n            \"schema\": {\n              \"$ref\": \"#/definitions/retentionRetentionPolicy\"\n            }\n          }\n        },\n        \"parameters\": [\n          {\n            \"name\": \"O.Tenant\",\n            \"in\": \"path\",\n            \"required\": true,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.Name\",\n            \"in\": \"path\",\n            \"required\": true,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"T.Kind\",\n            \"description\": \"Kind represents the type of the API object.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"T.APIVersion\",\n            \"description\": \"APIVersion defines the version of the API object.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.Namespace\",\n            \"description\": \"Namespace of the object, for scoped objects.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.ResourceVersion\",\n            \"description\": \"Resource version in the object store. This can only be set by the server.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.UUID\",\n            \"description\": \"UUID is the unique identifier for the object. This can only be set by the server.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.CreationTime.time\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\",\n            \"format\": \"date-time\"\n          },\n          {\n            \"name\": \"O.ModTime.time\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\",\n            \"format\": \"date-time\"\n          },\n          {\n            \"name\": \"Spec.CompactionInterval\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"Spec.CompactionMethod\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"Spec.MaxRetentionTime\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"Status.MonitoringPolicies\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"array\",\n            \"items\": {\n              \"type\": \"string\"\n            }\n          },\n          {\n            \"name\": \"Status.EventPolicies\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"array\",\n            \"items\": {\n              \"type\": \"string\"\n            }\n          }\n        ],\n        \"tags\": [\n          \"RetentionPolicyV1\"\n        ]\n      },\n      \"delete\": {\n        \"operationId\": \"AutoDeleteRetentionPolicy\",\n        \"responses\": {\n          \"200\": {\n            \"description\": \"\",\n            \"schema\": {\n              \"$ref\": \"#/definitions/retentionRetentionPolicy\"\n            }\n          }\n        },\n        \"parameters\": [\n          {\n            \"name\": \"O.Tenant\",\n            \"in\": \"path\",\n            \"required\": true,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.Name\",\n            \"in\": \"path\",\n            \"required\": true,\n            \"type\": \"string\"\n          }\n        ],\n        \"tags\": [\n          \"RetentionPolicyV1\"\n        ]\n      },\n      \"put\": {\n        \"operationId\": \"AutoUpdateRetentionPolicy\",\n        \"responses\": {\n          \"200\": {\n            \"description\": \"\",\n            \"schema\": {\n              \"$ref\": \"#/definitions/retentionRetentionPolicy\"\n            }\n          }\n        },\n        \"parameters\": [\n          {\n            \"name\": \"O.Tenant\",\n            \"in\": \"path\",\n            \"required\": true,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.Name\",\n            \"in\": \"path\",\n            \"required\": true,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"body\",\n            \"in\": \"body\",\n            \"required\": true,\n            \"schema\": {\n              \"$ref\": \"#/definitions/retentionRetentionPolicy\"\n            }\n          }\n        ],\n        \"tags\": [\n          \"RetentionPolicyV1\"\n        ]\n      }\n    }\n  },\n  \"definitions\": {\n    \"apiListMeta\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"ResourceVersion\": {\n          \"type\": \"string\",\n          \"description\": \"Resource version of object store at the time of list generation.\"\n        }\n      },\n      \"description\": \"ListMeta contains the metadata for list of objects.\"\n    },\n    \"apiListWatchOptions\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"O\": {\n          \"$ref\": \"#/definitions/apiObjectMeta\"\n        },\n        \"LabelSelector\": {\n          \"type\": \"string\"\n        },\n        \"FieldSelector\": {\n          \"type\": \"string\"\n        },\n        \"PrefixWatch\": {\n          \"type\": \"boolean\",\n          \"format\": \"boolean\"\n        }\n      }\n    },\n    \"apiObjectMeta\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"Name\": {\n          \"type\": \"string\",\n          \"description\": \"Name of the object, unique within a Namespace for scoped objects.\"\n        },\n        \"Tenant\": {\n          \"type\": \"string\",\n          \"description\": \"Tenant is global namespace isolation for various objects. This can be automatically\\nfilled in many cases based on the tenant a user, who created the object, belongs go.\"\n        },\n        \"Namespace\": {\n          \"type\": \"string\",\n          \"description\": \"Namespace of the object, for scoped objects.\"\n        },\n        \"ResourceVersion\": {\n          \"type\": \"string\",\n          \"description\": \"Resource version in the object store. This can only be set by the server.\"\n        },\n        \"UUID\": {\n          \"type\": \"string\",\n          \"description\": \"UUID is the unique identifier for the object. This can only be set by the server.\"\n        },\n        \"Labels\": {\n          \"type\": \"object\",\n          \"additionalProperties\": {\n            \"type\": \"string\"\n          },\n          \"description\": \"Labels are arbitrary (key,value) pairs associated with any object.\"\n        },\n        \"CreationTime\": {\n          \"$ref\": \"#/definitions/apiTimestamp\",\n          \"title\": \"CreationTime is the creation time of Object\"\n        },\n        \"ModTime\": {\n          \"$ref\": \"#/definitions/apiTimestamp\",\n          \"title\": \"ModTime is the Last Modification time of Object\"\n        }\n      },\n      \"description\": \"ObjectMeta contains metadata that all objects stored in kvstore must have.\"\n    },\n    \"apiTimestamp\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"time\": {\n          \"type\": \"string\",\n          \"format\": \"date-time\"\n        }\n      }\n    },\n    \"apiTypeMeta\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"Kind\": {\n          \"type\": \"string\",\n          \"description\": \"Kind represents the type of the API object.\"\n        },\n        \"APIVersion\": {\n          \"type\": \"string\",\n          \"description\": \"APIVersion defines the version of the API object.\"\n        }\n      },\n      \"description\": \"TypeMeta contains the metadata about kind and version for all API objects.\"\n    },\n    \"retentionAutoMsgRetentionPolicyWatchHelper\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"Type\": {\n          \"type\": \"string\",\n          \"title\": \"Compaction Interval is the time in minutes, hours or days before compaction starts on the data\\nCompaction results into data granularity loss, therefore this period is kept to be as high as space permits\\nfor how much data is being collected\"\n        },\n        \"Object\": {\n          \"$ref\": \"#/definitions/retentionRetentionPolicy\",\n          \"title\": \"Compaction Method specifies the method to be used for aggregation i.e. 'linear', 'exponential'\\nWhen compaction method is is unspecified system defaults to exponential aggregation over the period of time\"\n        }\n      }\n    },\n    \"retentionRetentionPolicy\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"T\": {\n          \"$ref\": \"#/definitions/apiTypeMeta\",\n          \"title\": \"list of monitoring policies that refer to this collection policy\"\n        },\n        \"O\": {\n          \"$ref\": \"#/definitions/apiObjectMeta\",\n          \"title\": \"list of event policies that refer to this collection policy\"\n        },\n        \"Spec\": {\n          \"$ref\": \"#/definitions/retentionRetentionPolicySpec\"\n        },\n        \"Status\": {\n          \"$ref\": \"#/definitions/retentionRetentionPolicyStatus\"\n        }\n      }\n    },\n    \"retentionRetentionPolicyList\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"T\": {\n          \"$ref\": \"#/definitions/apiTypeMeta\"\n        },\n        \"ListMeta\": {\n          \"$ref\": \"#/definitions/apiListMeta\"\n        },\n        \"Items\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"$ref\": \"#/definitions/retentionRetentionPolicy\"\n          },\n          \"description\": \"Spec contains the configuration of the retention policy.\"\n        }\n      }\n    },\n    \"retentionRetentionPolicySpec\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"CompactionInterval\": {\n          \"type\": \"string\"\n        },\n        \"CompactionMethod\": {\n          \"type\": \"string\"\n        },\n        \"MaxRetentionTime\": {\n          \"type\": \"string\"\n        }\n      }\n    },\n    \"retentionRetentionPolicyStatus\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"MonitoringPolicies\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"string\"\n          }\n        },\n        \"EventPolicies\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"string\"\n          }\n        }\n      }\n    }\n  }\n}\n"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1510367067, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // "retention.swagger.json"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`../../../../../sw/api/generated/retention/swagger`, &embedded.EmbeddedBox{
		Name: `../../../../../sw/api/generated/retention/swagger`,
		Time: time.Unix(1510367067, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"retention.swagger.json": file2,
		},
	})
}
