/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { SearchPolicySearchResponse_status,  } from './enums';

export interface ISearchPolicySearchResponse {
    'status': SearchPolicySearchResponse_status;
    'results'?: object;
}


export class SearchPolicySearchResponse extends BaseModel implements ISearchPolicySearchResponse {
    /** Status of firewall policy search. */
    'status': SearchPolicySearchResponse_status = null;
    /** Result is Map of <NetworkSecurityPolicy object name, PolicyMatch Entry>. */
    'results': object = null;
    public static propInfo: { [prop in keyof ISearchPolicySearchResponse]: PropInfoItem } = {
        'status': {
            enum: SearchPolicySearchResponse_status,
            default: 'match',
            description:  `Status of firewall policy search.`,
            required: true,
            type: 'string'
        },
        'results': {
            description:  `Result is Map of <NetworkSecurityPolicy object name, PolicyMatch Entry>.`,
            required: false,
            type: 'object'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return SearchPolicySearchResponse.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return SearchPolicySearchResponse.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (SearchPolicySearchResponse.propInfo[prop] != null &&
                        SearchPolicySearchResponse.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this._inputValue = values;
        this.setValues(values, setDefaults);
    }

    /**
     * set the values for both the Model and the Form Group. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any, fillDefaults = true): void {
        if (values && values['status'] != null) {
            this['status'] = values['status'];
        } else if (fillDefaults && SearchPolicySearchResponse.hasDefaultValue('status')) {
            this['status'] = <SearchPolicySearchResponse_status>  SearchPolicySearchResponse.propInfo['status'].default;
        } else {
            this['status'] = null
        }
        if (values && values['results'] != null) {
            this['results'] = values['results'];
        } else if (fillDefaults && SearchPolicySearchResponse.hasDefaultValue('results')) {
            this['results'] = SearchPolicySearchResponse.propInfo['results'].default;
        } else {
            this['results'] = null
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'status': CustomFormControl(new FormControl(this['status'], [required, enumValidator(SearchPolicySearchResponse_status), ]), SearchPolicySearchResponse.propInfo['status']),
                'results': CustomFormControl(new FormControl(this['results']), SearchPolicySearchResponse.propInfo['results']),
            });
        }
        return this._formGroup;
    }

    setModelToBeFormGroupValues() {
        this.setValues(this.$formGroup.value, false);
    }

    setFormGroupValuesToBeModelValues() {
        if (this._formGroup) {
            this._formGroup.controls['status'].setValue(this['status']);
            this._formGroup.controls['results'].setValue(this['results']);
        }
    }
}

