/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { AuthRadiusServer, IAuthRadiusServer } from './auth-radius-server.model';

export interface IAuthRadiusDomain {
    'nas-id'?: string;
    'servers'?: Array<IAuthRadiusServer>;
    'tag'?: string;
}


export class AuthRadiusDomain extends BaseModel implements IAuthRadiusDomain {
    /** NasID is a string identifying the NAS(API Gw) originating the Access-Request. */
    'nas-id': string = null;
    'servers': Array<AuthRadiusServer> = null;
    /** Tag to group domains for authentication. */
    'tag': string = null;
    public static propInfo: { [prop in keyof IAuthRadiusDomain]: PropInfoItem } = {
        'nas-id': {
            description:  `NasID is a string identifying the NAS(API Gw) originating the Access-Request.`,
            required: false,
            type: 'string'
        },
        'servers': {
            required: false,
            type: 'object'
        },
        'tag': {
            description:  `Tag to group domains for authentication.`,
            required: false,
            type: 'string'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return AuthRadiusDomain.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return AuthRadiusDomain.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (AuthRadiusDomain.propInfo[prop] != null &&
                        AuthRadiusDomain.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['servers'] = new Array<AuthRadiusServer>();
        this._inputValue = values;
        this.setValues(values, setDefaults);
    }

    /**
     * set the values for both the Model and the Form Group. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any, fillDefaults = true): void {
        if (values && values['nas-id'] != null) {
            this['nas-id'] = values['nas-id'];
        } else if (fillDefaults && AuthRadiusDomain.hasDefaultValue('nas-id')) {
            this['nas-id'] = AuthRadiusDomain.propInfo['nas-id'].default;
        } else {
            this['nas-id'] = null
        }
        if (values) {
            this.fillModelArray<AuthRadiusServer>(this, 'servers', values['servers'], AuthRadiusServer);
        } else {
            this['servers'] = [];
        }
        if (values && values['tag'] != null) {
            this['tag'] = values['tag'];
        } else if (fillDefaults && AuthRadiusDomain.hasDefaultValue('tag')) {
            this['tag'] = AuthRadiusDomain.propInfo['tag'].default;
        } else {
            this['tag'] = null
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'nas-id': CustomFormControl(new FormControl(this['nas-id']), AuthRadiusDomain.propInfo['nas-id']),
                'servers': new FormArray([]),
                'tag': CustomFormControl(new FormControl(this['tag']), AuthRadiusDomain.propInfo['tag']),
            });
            // generate FormArray control elements
            this.fillFormArray<AuthRadiusServer>('servers', this['servers'], AuthRadiusServer);
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('servers') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('servers').get(field);
                control.updateValueAndValidity();
            });
        }
        return this._formGroup;
    }

    setModelToBeFormGroupValues() {
        this.setValues(this.$formGroup.value, false);
    }

    setFormGroupValuesToBeModelValues() {
        if (this._formGroup) {
            this._formGroup.controls['nas-id'].setValue(this['nas-id']);
            this.fillModelArray<AuthRadiusServer>(this, 'servers', this['servers'], AuthRadiusServer);
            this._formGroup.controls['tag'].setValue(this['tag']);
        }
    }
}

