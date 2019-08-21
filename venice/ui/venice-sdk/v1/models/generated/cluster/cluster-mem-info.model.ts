/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { ClusterMemInfo_type,  } from './enums';

export interface IClusterMemInfo {
    'type': ClusterMemInfo_type;
    'size'?: string;
}


export class ClusterMemInfo extends BaseModel implements IClusterMemInfo {
    'type': ClusterMemInfo_type = null;
    'size': string = null;
    public static propInfo: { [prop in keyof IClusterMemInfo]: PropInfoItem } = {
        'type': {
            enum: ClusterMemInfo_type,
            default: 'unknown',
            required: true,
            type: 'string'
        },
        'size': {
            required: false,
            type: 'string'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return ClusterMemInfo.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return ClusterMemInfo.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (ClusterMemInfo.propInfo[prop] != null &&
                        ClusterMemInfo.propInfo[prop].default != null);
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
        if (values && values['type'] != null) {
            this['type'] = values['type'];
        } else if (fillDefaults && ClusterMemInfo.hasDefaultValue('type')) {
            this['type'] = <ClusterMemInfo_type>  ClusterMemInfo.propInfo['type'].default;
        } else {
            this['type'] = null
        }
        if (values && values['size'] != null) {
            this['size'] = values['size'];
        } else if (fillDefaults && ClusterMemInfo.hasDefaultValue('size')) {
            this['size'] = ClusterMemInfo.propInfo['size'].default;
        } else {
            this['size'] = null
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'type': CustomFormControl(new FormControl(this['type'], [required, enumValidator(ClusterMemInfo_type), ]), ClusterMemInfo.propInfo['type']),
                'size': CustomFormControl(new FormControl(this['size']), ClusterMemInfo.propInfo['size']),
            });
        }
        return this._formGroup;
    }

    setModelToBeFormGroupValues() {
        this.setValues(this.$formGroup.value, false);
    }

    setFormGroupValuesToBeModelValues() {
        if (this._formGroup) {
            this._formGroup.controls['type'].setValue(this['type']);
            this._formGroup.controls['size'].setValue(this['size']);
        }
    }
}

