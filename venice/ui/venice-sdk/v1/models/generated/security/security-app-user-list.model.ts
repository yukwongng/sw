/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, enumValidator } from './validators';
import { BaseModel, EnumDef } from './base-model';

import { SecurityAppUser, ISecurityAppUser } from './security-app-user.model';

export interface ISecurityAppUserList {
    'kind'?: string;
    'api-version'?: string;
    'resource-version'?: string;
    'Items'?: Array<ISecurityAppUser>;
}


export class SecurityAppUserList extends BaseModel implements ISecurityAppUserList {
    'kind': string;
    'api-version': string;
    'resource-version': string;
    'Items': Array<SecurityAppUser>;
    public static enumProperties: { [key: string] : EnumDef } = {
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any) {
        super();
        this['Items'] = new Array<SecurityAppUser>();
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
            this['resource-version'] = values['resource-version'];
            this.fillModelArray<SecurityAppUser>(this, 'Items', values['Items'], SecurityAppUser);
        }
    }

    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'kind': new FormControl(this['kind']),
                'api-version': new FormControl(this['api-version']),
                'resource-version': new FormControl(this['resource-version']),
                'Items': new FormArray([]),
            });
            // generate FormArray control elements
            this.fillFormArray<SecurityAppUser>('Items', this['Items'], SecurityAppUser);
        }
        return this._formGroup;
    }

    setFormGroupValues() {
        if (this._formGroup) {
            this._formGroup.controls['kind'].setValue(this['kind']);
            this._formGroup.controls['api-version'].setValue(this['api-version']);
            this._formGroup.controls['resource-version'].setValue(this['resource-version']);
            this.fillModelArray<SecurityAppUser>(this, 'Items', this['Items'], SecurityAppUser);
        }
    }
}

