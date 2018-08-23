/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, enumValidator } from './validators';
import { BaseModel, PropInfoItem } from './base-model';

import { SearchPolicySearchResponse_status,  } from './enums';

export interface ISearchPolicySearchResponse {
    'status'?: SearchPolicySearchResponse_status;
    'results'?: object;
}


export class SearchPolicySearchResponse extends BaseModel implements ISearchPolicySearchResponse {
    'status': SearchPolicySearchResponse_status = null;
    /** Result is Map of <SGPolicy object name, PolicyMatch Entry>. */
    'results': object = null;
    public static propInfo: { [prop: string]: PropInfoItem } = {
        'status': {
            enum: SearchPolicySearchResponse_status,
            default: 'MATCH',
            type: 'string'
        },
        'results': {
            description:  'Result is Map of &lt;SGPolicy object name, PolicyMatch Entry&gt;.',
            type: 'object'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return SearchPolicySearchResponse.propInfo[propName];
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (SearchPolicySearchResponse.propInfo[prop] != null &&
                        SearchPolicySearchResponse.propInfo[prop].default != null &&
                        SearchPolicySearchResponse.propInfo[prop].default != '');
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any) {
        super();
        this.setValues(values);
    }

    /**
     * set the values. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any): void {
        if (values && values['status'] != null) {
            this['status'] = values['status'];
        } else if (SearchPolicySearchResponse.hasDefaultValue('status')) {
            this['status'] = <SearchPolicySearchResponse_status>  SearchPolicySearchResponse.propInfo['status'].default;
        }
        if (values && values['results'] != null) {
            this['results'] = values['results'];
        } else if (SearchPolicySearchResponse.hasDefaultValue('results')) {
            this['results'] = SearchPolicySearchResponse.propInfo['results'].default;
        }
    }




    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'status': new FormControl(this['status'], [enumValidator(SearchPolicySearchResponse_status), ]),
                'results': new FormControl(this['results']),
            });
        }
        return this._formGroup;
    }

    setFormGroupValues() {
        if (this._formGroup) {
            this._formGroup.controls['status'].setValue(this['status']);
            this._formGroup.controls['results'].setValue(this['results']);
        }
    }
}

