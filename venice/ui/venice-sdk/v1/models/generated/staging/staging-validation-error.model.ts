/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';


export interface IStagingValidationError {
    'uri'?: string;
    'method'?: string;
    'error'?: Array<string>;
    '_ui'?: any;
}


export class StagingValidationError extends BaseModel implements IStagingValidationError {
    /** Field for holding arbitrary ui state */
    '_ui': any = {};
    'uri': string = null;
    'method': string = null;
    'error': Array<string> = null;
    public static propInfo: { [prop in keyof IStagingValidationError]: PropInfoItem } = {
        'uri': {
            required: false,
            type: 'string'
        },
        'method': {
            required: false,
            type: 'string'
        },
        'error': {
            required: false,
            type: 'Array<string>'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return StagingValidationError.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return StagingValidationError.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (StagingValidationError.propInfo[prop] != null &&
                        StagingValidationError.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['error'] = new Array<string>();
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
        if (values && values['uri'] != null) {
            this['uri'] = values['uri'];
        } else if (fillDefaults && StagingValidationError.hasDefaultValue('uri')) {
            this['uri'] = StagingValidationError.propInfo['uri'].default;
        } else {
            this['uri'] = null
        }
        if (values && values['method'] != null) {
            this['method'] = values['method'];
        } else if (fillDefaults && StagingValidationError.hasDefaultValue('method')) {
            this['method'] = StagingValidationError.propInfo['method'].default;
        } else {
            this['method'] = null
        }
        if (values && values['error'] != null) {
            this['error'] = values['error'];
        } else if (fillDefaults && StagingValidationError.hasDefaultValue('error')) {
            this['error'] = [ StagingValidationError.propInfo['error'].default];
        } else {
            this['error'] = [];
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'uri': CustomFormControl(new FormControl(this['uri']), StagingValidationError.propInfo['uri']),
                'method': CustomFormControl(new FormControl(this['method']), StagingValidationError.propInfo['method']),
                'error': CustomFormControl(new FormControl(this['error']), StagingValidationError.propInfo['error']),
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
            this._formGroup.controls['error'].setValue(this['error']);
        }
    }
}

