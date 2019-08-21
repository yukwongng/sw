/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';


export interface IAuthResource {
    'tenant'?: string;
    'group'?: string;
    'kind'?: string;
    'namespace'?: string;
    'name'?: string;
}


export class AuthResource extends BaseModel implements IAuthResource {
    'tenant': string = null;
    'group': string = null;
    'kind': string = null;
    'namespace': string = null;
    'name': string = null;
    public static propInfo: { [prop in keyof IAuthResource]: PropInfoItem } = {
        'tenant': {
            required: false,
            type: 'string'
        },
        'group': {
            required: false,
            type: 'string'
        },
        'kind': {
            required: false,
            type: 'string'
        },
        'namespace': {
            required: false,
            type: 'string'
        },
        'name': {
            required: false,
            type: 'string'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return AuthResource.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return AuthResource.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (AuthResource.propInfo[prop] != null &&
                        AuthResource.propInfo[prop].default != null);
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
        if (values && values['tenant'] != null) {
            this['tenant'] = values['tenant'];
        } else if (fillDefaults && AuthResource.hasDefaultValue('tenant')) {
            this['tenant'] = AuthResource.propInfo['tenant'].default;
        } else {
            this['tenant'] = null
        }
        if (values && values['group'] != null) {
            this['group'] = values['group'];
        } else if (fillDefaults && AuthResource.hasDefaultValue('group')) {
            this['group'] = AuthResource.propInfo['group'].default;
        } else {
            this['group'] = null
        }
        if (values && values['kind'] != null) {
            this['kind'] = values['kind'];
        } else if (fillDefaults && AuthResource.hasDefaultValue('kind')) {
            this['kind'] = AuthResource.propInfo['kind'].default;
        } else {
            this['kind'] = null
        }
        if (values && values['namespace'] != null) {
            this['namespace'] = values['namespace'];
        } else if (fillDefaults && AuthResource.hasDefaultValue('namespace')) {
            this['namespace'] = AuthResource.propInfo['namespace'].default;
        } else {
            this['namespace'] = null
        }
        if (values && values['name'] != null) {
            this['name'] = values['name'];
        } else if (fillDefaults && AuthResource.hasDefaultValue('name')) {
            this['name'] = AuthResource.propInfo['name'].default;
        } else {
            this['name'] = null
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'tenant': CustomFormControl(new FormControl(this['tenant']), AuthResource.propInfo['tenant']),
                'group': CustomFormControl(new FormControl(this['group']), AuthResource.propInfo['group']),
                'kind': CustomFormControl(new FormControl(this['kind']), AuthResource.propInfo['kind']),
                'namespace': CustomFormControl(new FormControl(this['namespace']), AuthResource.propInfo['namespace']),
                'name': CustomFormControl(new FormControl(this['name']), AuthResource.propInfo['name']),
            });
        }
        return this._formGroup;
    }

    setModelToBeFormGroupValues() {
        this.setValues(this.$formGroup.value, false);
    }

    setFormGroupValuesToBeModelValues() {
        if (this._formGroup) {
            this._formGroup.controls['tenant'].setValue(this['tenant']);
            this._formGroup.controls['group'].setValue(this['group']);
            this._formGroup.controls['kind'].setValue(this['kind']);
            this._formGroup.controls['namespace'].setValue(this['namespace']);
            this._formGroup.controls['name'].setValue(this['name']);
        }
    }
}

