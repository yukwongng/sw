/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, enumValidator } from './validators';
import { BaseModel, EnumDef } from './base-model';

import { ApiObjectMeta, IApiObjectMeta } from './api-object-meta.model';
import { AuthRoleBindingSpec, IAuthRoleBindingSpec } from './auth-role-binding-spec.model';
import { AuthRoleBindingStatus, IAuthRoleBindingStatus } from './auth-role-binding-status.model';

export interface IAuthRoleBinding {
    'kind'?: string;
    'api-version'?: string;
    'meta'?: IApiObjectMeta;
    'spec'?: IAuthRoleBindingSpec;
    'status'?: IAuthRoleBindingStatus;
}


export class AuthRoleBinding extends BaseModel implements IAuthRoleBinding {
    'kind': string;
    'api-version': string;
    'meta': ApiObjectMeta;
    /** Spec contains the configuration of the role binding. */
    'spec': AuthRoleBindingSpec;
    /** Status contains the current state of the role binding. */
    'status': AuthRoleBindingStatus;
    public static enumProperties: { [key: string] : EnumDef } = {
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any) {
        super();
        this['meta'] = new ApiObjectMeta();
        this['spec'] = new AuthRoleBindingSpec();
        this['status'] = new AuthRoleBindingStatus();
        if (values) {
            this.setValues(values);
        }
    }

    /**
     * set the values.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any): void {
        if (values) {
            this['kind'] = values['kind'];
            this['api-version'] = values['api-version'];
            this['meta'].setValues(values['meta']);
            this['spec'].setValues(values['spec']);
            this['status'].setValues(values['status']);
        }
    }

    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'kind': new FormControl(this['kind']),
                'api-version': new FormControl(this['api-version']),
                'meta': this['meta'].$formGroup,
                'spec': this['spec'].$formGroup,
                'status': this['status'].$formGroup,
            });
        }
        return this._formGroup;
    }

    setFormGroupValues() {
        if (this._formGroup) {
            this._formGroup.controls['kind'].setValue(this['kind']);
            this._formGroup.controls['api-version'].setValue(this['api-version']);
            this['meta'].setFormGroupValues();
            this['spec'].setFormGroupValues();
            this['status'].setFormGroupValues();
        }
    }
}

