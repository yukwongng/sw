/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';


export interface IStagingItemId {
    'uri'?: string;
    'method'?: string;
}


export class StagingItemId extends BaseModel implements IStagingItemId {
    'uri': string = null;
    'method': string = null;
    public static propInfo: { [prop in keyof IStagingItemId]: PropInfoItem } = {
        'uri': {
            required: false,
            type: 'string'
        },
        'method': {
            required: false,
            type: 'string'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return StagingItemId.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return StagingItemId.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (StagingItemId.propInfo[prop] != null &&
                        StagingItemId.propInfo[prop].default != null);
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
        if (values && values['uri'] != null) {
            this['uri'] = values['uri'];
        } else if (fillDefaults && StagingItemId.hasDefaultValue('uri')) {
            this['uri'] = StagingItemId.propInfo['uri'].default;
        } else {
            this['uri'] = null
        }
        if (values && values['method'] != null) {
            this['method'] = values['method'];
        } else if (fillDefaults && StagingItemId.hasDefaultValue('method')) {
            this['method'] = StagingItemId.propInfo['method'].default;
        } else {
            this['method'] = null
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'uri': CustomFormControl(new FormControl(this['uri']), StagingItemId.propInfo['uri']),
                'method': CustomFormControl(new FormControl(this['method']), StagingItemId.propInfo['method']),
            });
        }
        return this._formGroup;
    }

    setModelToBeFormGroupValues() {
        this.setValues(this.$formGroup.value, false);
    }

    setFormGroupValuesToBeModelValues() {
        if (this._formGroup) {
            this._formGroup.controls['uri'].setValue(this['uri']);
            this._formGroup.controls['method'].setValue(this['method']);
        }
    }
}

