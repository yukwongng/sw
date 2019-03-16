/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */

// generate enum based on strings instead of numbers
// (see https://blog.rsuter.com/how-to-implement-an-enum-with-string-values-in-typescript/)
export enum ApiListWatchOptions_sort_order {
    'None' = "None",
    'ByName' = "ByName",
    'ByNameReverse' = "ByNameReverse",
    'ByVersion' = "ByVersion",
    'ByVersionReverse' = "ByVersionReverse",
    'ByCreationTime' = "ByCreationTime",
    'ByCreationTimeReverse' = "ByCreationTimeReverse",
    'ByModTime' = "ByModTime",
    'ByModTimeReverse' = "ByModTimeReverse",
}

export enum FieldsRequirement_operator {
    'equals' = "equals",
    'notEquals' = "notEquals",
    'in' = "in",
    'notIn' = "notIn",
    'gt' = "gt",
    'gte' = "gte",
    'lt' = "lt",
    'lte' = "lte",
}

export enum LabelsRequirement_operator {
    'equals' = "equals",
    'notEquals' = "notEquals",
    'in' = "in",
    'notIn' = "notIn",
}

export enum SearchPolicySearchResponse_status {
    'MATCH' = "MATCH",
    'MISS' = "MISS",
}

export enum SearchSearchQuery_categories {
    'Cluster' = "Cluster",
    'Workload' = "Workload",
    'Security' = "Security",
    'Auth' = "Auth",
    'Network' = "Network",
    'Monitoring' = "Monitoring",
    'Telemetry' = "Telemetry",
    'Events' = "Events",
    'Alerts' = "Alerts",
    'AuditTrail' = "AuditTrail",
}

export enum SearchSearchQuery_kinds {
    'Cluster' = "Cluster",
    'Node' = "Node",
    'SmartNIC' = "SmartNIC",
    'Rollout' = "Rollout",
    'Tenant' = "Tenant",
    'Endpoint' = "Endpoint",
    'SecurityGroup' = "SecurityGroup",
    'SGPolicy' = "SGPolicy",
    'App' = "App",
    'AppUser' = "AppUser",
    'AppUserGrp' = "AppUserGrp",
    'Certificate' = "Certificate",
    'TrafficEncryptionPolicy' = "TrafficEncryptionPolicy",
    'User' = "User",
    'AuthenticationPolicy' = "AuthenticationPolicy",
    'Role' = "Role",
    'RoleBinding' = "RoleBinding",
    'Network' = "Network",
    'Service' = "Service",
    'LbPolicy' = "LbPolicy",
    'Alert' = "Alert",
    'AlertDestination' = "AlertDestination",
    'AlertPolicy' = "AlertPolicy",
    'Event' = "Event",
    'EventPolicy' = "EventPolicy",
    'StatsPolicy' = "StatsPolicy",
    'FlowExportPolicy' = "FlowExportPolicy",
    'FwlogPolicy' = "FwlogPolicy",
    'MirrorSession' = "MirrorSession",
    'Workload' = "Workload",
    'Host' = "Host",
    'AuditEvent' = "AuditEvent",
}

export enum SearchSearchRequest_sort_order {
    'Ascending' = "Ascending",
    'Descending' = "Descending",
}

export enum SearchSearchRequest_mode {
    'Full' = "Full",
    'Preview' = "Preview",
}

export enum SecuritySGRule_action {
    'PERMIT' = "PERMIT",
    'DENY' = "DENY",
    'REJECT' = "REJECT",
}


export enum ApiListWatchOptions_sort_order_uihint {
    'ByCreationTime' = "By Creation Time",
    'ByCreationTimeReverse' = "By Creation Time Reverse",
    'ByModTime' = "By Modification Time",
    'ByModTimeReverse' = "By Modification Time Reverse",
    'ByName' = "By Name",
    'ByNameReverse' = "By Name Reverse",
    'ByVersion' = "By Version",
    'ByVersionReverse' = "By Version Reverse",
    'None' = "None",
}

export enum FieldsRequirement_operator_uihint {
    'gt' = "greater than",
    'gte' = "greater than or equals",
    'lt' = "less than",
    'lte' = "less than or equals",
    'notEquals' = "not equals",
    'notIn' = "not in",
}

export enum LabelsRequirement_operator_uihint {
    'notEquals' = "not equals",
    'notIn' = "not in",
}




/**
 * bundle of all enums for databinding to options, radio-buttons etc.
 * usage in component:
 *   import { AllEnums, minValueValidator, maxValueValidator } from '../../models/webapi';
 *
 *   @Component({
 *       ...
 *   })
 *   export class xxxComponent implements OnInit {
 *       allEnums = AllEnums;
 *       ...
 *       ngOnInit() {
 *           this.allEnums = AllEnums.instance;
 *       }
 *   }
*/
export class AllEnums {
    private static _instance: AllEnums = new AllEnums();
    constructor() {
        if (AllEnums._instance) {
            throw new Error("Error: Instantiation failed: Use AllEnums.instance instead of new");
        }
        AllEnums._instance = this;
    }
    static get instance(): AllEnums {
        return AllEnums._instance;
    }

    ApiListWatchOptions_sort_order = ApiListWatchOptions_sort_order;
    FieldsRequirement_operator = FieldsRequirement_operator;
    LabelsRequirement_operator = LabelsRequirement_operator;
    SearchPolicySearchResponse_status = SearchPolicySearchResponse_status;
    SearchSearchQuery_categories = SearchSearchQuery_categories;
    SearchSearchQuery_kinds = SearchSearchQuery_kinds;
    SearchSearchRequest_sort_order = SearchSearchRequest_sort_order;
    SearchSearchRequest_mode = SearchSearchRequest_mode;
    SecuritySGRule_action = SecuritySGRule_action;

    ApiListWatchOptions_sort_order_uihint = ApiListWatchOptions_sort_order_uihint;
    FieldsRequirement_operator_uihint = FieldsRequirement_operator_uihint;
    LabelsRequirement_operator_uihint = LabelsRequirement_operator_uihint;
}
