/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';


export interface IApiListMeta {
    'resource-version'?: string;
    '_ui'?: any;
}


export class ApiListMeta extends BaseModel implements IApiListMeta {
    /** Field for holding arbitrary ui state */
    '_ui': any = {};
    /** Resource version of object store at the time of list generation. */
    'resource-version': string = null;
    public static propInfo: { [prop in keyof IApiListMeta]: PropInfoItem } = {
        'resource-version': {
            description:  `Resource version of object store at the time of list generation.`,
            required: false,
            type: 'string'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return ApiListMeta.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return ApiListMeta.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (ApiListMeta.propInfo[prop] != null &&
                        ApiListMeta.propInfo[prop].default != null);
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
        if (values && values['resource-version'] != null) {
            this['resource-version'] = values['resource-version'];
        } else if (fillDefaults && ApiListMeta.hasDefaultValue('resource-version')) {
            this['resource-version'] = ApiListMeta.propInfo['resource-version'].default;
        } else {
            this['resource-version'] = null
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'resource-version': CustomFormControl(new FormControl(this['resource-version']), ApiListMeta.propInfo['resource-version']),
            });
        }
        return this._formGroup;
    }

    setModelToBeFormGroupValues() {
        this.setValues(this.$formGroup.value, false);
    }

    setFormGroupValuesToBeModelValues() {
        if (this._formGroup) {
            this._formGroup.controls['resource-version'].setValue(this['resource-version']);
        }
    }
}

