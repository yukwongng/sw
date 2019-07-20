/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */

// generate enum based on strings instead of numbers
// (see https://blog.rsuter.com/how-to-implement-an-enum-with-string-values-in-typescript/)
export enum ApiListWatchOptions_sort_order {
    'none' = "none",
    'by-name' = "by-name",
    'by-name-reverse' = "by-name-reverse",
    'by-version' = "by-version",
    'by-version-reverse' = "by-version-reverse",
    'by-creation-time' = "by-creation-time",
    'by-creation-time-reverse' = "by-creation-time-reverse",
    'by-mod-time' = "by-mod-time",
    'by-mod-time-reverse' = "by-mod-time-reverse",
}

export enum FieldsRequirement_operator {
    'equals' = "equals",
    'notequals' = "notequals",
    'in' = "in",
    'notin' = "notin",
    'gt' = "gt",
    'gte' = "gte",
    'lt' = "lt",
    'lte' = "lte",
}

export enum LabelsRequirement_operator {
    'equals' = "equals",
    'notequals' = "notequals",
    'in' = "in",
    'notin' = "notin",
}

export enum SearchPolicySearchResponse_status {
    'match' = "match",
    'miss' = "miss",
}

export enum SearchSearchRequest_sort_order {
    'ascending' = "ascending",
    'descending' = "descending",
}

export enum SearchSearchRequest_mode {
    'full' = "full",
    'preview' = "preview",
}

export enum SecuritySGRule_action {
    'permit' = "permit",
    'deny' = "deny",
    'reject' = "reject",
}


export enum ApiListWatchOptions_sort_order_uihint {
    'by-creation-time' = "By Creation Time",
    'by-creation-time-reverse' = "By Creation Time Reverse",
    'by-mod-time' = "By Modification Time",
    'by-mod-time-reverse' = "By Modification Time Reverse",
    'by-name' = "By Name",
    'by-name-reverse' = "By Name Reverse",
    'by-version' = "By Version",
    'by-version-reverse' = "By Version Reverse",
    'none' = "None",
}

export enum FieldsRequirement_operator_uihint {
    'gt' = "greater than",
    'gte' = "greater than or equals",
    'lt' = "less than",
    'lte' = "less than or equals",
    'notequals' = "not equals",
    'notin' = "not in",
}

export enum LabelsRequirement_operator_uihint {
    'notequals' = "not equals",
    'notin' = "not in",
}

export enum SecuritySGRule_action_uihint {
    'deny' = "Deny",
    'permit' = "Permit",
    'reject' = "Reject",
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
    SearchSearchRequest_sort_order = SearchSearchRequest_sort_order;
    SearchSearchRequest_mode = SearchSearchRequest_mode;
    SecuritySGRule_action = SecuritySGRule_action;

    ApiListWatchOptions_sort_order_uihint = ApiListWatchOptions_sort_order_uihint;
    FieldsRequirement_operator_uihint = FieldsRequirement_operator_uihint;
    LabelsRequirement_operator_uihint = LabelsRequirement_operator_uihint;
    SecuritySGRule_action_uihint = SecuritySGRule_action_uihint;
}
