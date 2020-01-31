/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';


export interface IApiTypeMeta {
    'kind'?: string;
    'api-version'?: string;
    '_ui'?: any;
}


export class ApiTypeMeta extends BaseModel implements IApiTypeMeta {
    /** Field for holding arbitrary ui state */
    '_ui': any = {};
    /** Kind represents the type of the API object. */
    'kind': string = null;
    /** APIVersion defines the version of the API object. This can only be set by the server. */
    'api-version': string = null;
    public static propInfo: { [prop in keyof IApiTypeMeta]: PropInfoItem } = {
        'kind': {
            description:  `Kind represents the type of the API object.`,
            required: false,
            type: 'string'
        },
        'api-version': {
            description:  `APIVersion defines the version of the API object. This can only be set by the server.`,
            required: false,
            type: 'string'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return ApiTypeMeta.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return ApiTypeMeta.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (ApiTypeMeta.propInfo[prop] != null &&
                        ApiTypeMeta.propInfo[prop].default != null);
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
        if (values && values['_ui']) {
            this['_ui'] = values['_ui']
        }
        if (values && values['kind'] != null) {
            this['kind'] = values['kind'];
        } else if (fillDefaults && ApiTypeMeta.hasDefaultValue('kind')) {
            this['kind'] = ApiTypeMeta.propInfo['kind'].default;
        } else {
            this['kind'] = null
        }
        if (values && values['api-version'] != null) {
            this['api-version'] = values['api-version'];
        } else if (fillDefaults && ApiTypeMeta.hasDefaultValue('api-version')) {
            this['api-version'] = ApiTypeMeta.propInfo['api-version'].default;
        } else {
            this['api-version'] = null
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'kind': CustomFormControl(new FormControl(this['kind']), ApiTypeMeta.propInfo['kind']),
                'api-version': CustomFormControl(new FormControl(this['api-version']), ApiTypeMeta.propInfo['api-version']),
            });
        }
        return this._formGroup;
    }

    setModelToBeFormGroupValues() {
        this.setValues(this.$formGroup.value, false);
    }

    setFormGroupValuesToBeModelValues() {
        if (this._formGroup) {
            this._formGroup.controls['kind'].setValue(this['kind']);
            this._formGroup.controls['api-version'].setValue(this['api-version']);
        }
    }
}

