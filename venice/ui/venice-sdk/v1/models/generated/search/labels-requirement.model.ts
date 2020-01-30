/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { LabelsRequirement_operator,  LabelsRequirement_operator_uihint  } from './enums';

export interface ILabelsRequirement {
    'key'?: string;
    'operator': LabelsRequirement_operator;
    'values'?: Array<string>;
}


export class LabelsRequirement extends BaseModel implements ILabelsRequirement {
    /** The label key that the condition applies to. */
    'key': string = null;
    /** Condition checked for the key. */
    'operator': LabelsRequirement_operator = null;
    /** Values contains one or more values corresponding to the label key. "equals" and "notEquals" operators need a single Value. "in" and "notIn" operators can have one or more values. */
    'values': Array<string> = null;
    public static propInfo: { [prop in keyof ILabelsRequirement]: PropInfoItem } = {
        'key': {
            description:  `The label key that the condition applies to.`,
            required: false,
            type: 'string'
        },
        'operator': {
            enum: LabelsRequirement_operator_uihint,
            default: 'equals',
            description:  `Condition checked for the key.`,
            required: true,
            type: 'string'
        },
        'values': {
            description:  `Values contains one or more values corresponding to the label key. "equals" and "notEquals" operators need a single Value. "in" and "notIn" operators can have one or more values.`,
            required: false,
            type: 'Array<string>'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return LabelsRequirement.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return LabelsRequirement.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (LabelsRequirement.propInfo[prop] != null &&
                        LabelsRequirement.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['values'] = new Array<string>();
        this._inputValue = values;
        this.setValues(values, setDefaults);
    }

    /**
     * set the values for both the Model and the Form Group. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any, fillDefaults = true): void {
        if (values && values['key'] != null) {
            this['key'] = values['key'];
        } else if (fillDefaults && LabelsRequirement.hasDefaultValue('key')) {
            this['key'] = LabelsRequirement.propInfo['key'].default;
        } else {
            this['key'] = null
        }
        if (values && values['operator'] != null) {
            this['operator'] = values['operator'];
        } else if (fillDefaults && LabelsRequirement.hasDefaultValue('operator')) {
            this['operator'] = <LabelsRequirement_operator>  LabelsRequirement.propInfo['operator'].default;
        } else {
            this['operator'] = null
        }
        if (values && values['values'] != null) {
            this['values'] = values['values'];
        } else if (fillDefaults && LabelsRequirement.hasDefaultValue('values')) {
            this['values'] = [ LabelsRequirement.propInfo['values'].default];
        } else {
            this['values'] = [];
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'key': CustomFormControl(new FormControl(this['key']), LabelsRequirement.propInfo['key']),
                'operator': CustomFormControl(new FormControl(this['operator'], [required, enumValidator(LabelsRequirement_operator), ]), LabelsRequirement.propInfo['operator']),
                'values': CustomFormControl(new FormControl(this['values']), LabelsRequirement.propInfo['values']),
            });
        }
        return this._formGroup;
    }

    setModelToBeFormGroupValues() {
        this.setValues(this.$formGroup.value, false);
    }

    setFormGroupValuesToBeModelValues() {
        if (this._formGroup) {
            this._formGroup.controls['key'].setValue(this['key']);
            this._formGroup.controls['operator'].setValue(this['operator']);
            this._formGroup.controls['values'].setValue(this['values']);
        }
    }
}

