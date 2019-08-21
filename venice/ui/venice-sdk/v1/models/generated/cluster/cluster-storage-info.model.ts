/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { ClusterStorageDeviceInfo, IClusterStorageDeviceInfo } from './cluster-storage-device-info.model';

export interface IClusterStorageInfo {
    'devices'?: Array<IClusterStorageDeviceInfo>;
}


export class ClusterStorageInfo extends BaseModel implements IClusterStorageInfo {
    'devices': Array<ClusterStorageDeviceInfo> = null;
    public static propInfo: { [prop in keyof IClusterStorageInfo]: PropInfoItem } = {
        'devices': {
            required: false,
            type: 'object'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return ClusterStorageInfo.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return ClusterStorageInfo.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (ClusterStorageInfo.propInfo[prop] != null &&
                        ClusterStorageInfo.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['devices'] = new Array<ClusterStorageDeviceInfo>();
        this._inputValue = values;
        this.setValues(values, setDefaults);
    }

    /**
     * set the values for both the Model and the Form Group. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any, fillDefaults = true): void {
        if (values) {
            this.fillModelArray<ClusterStorageDeviceInfo>(this, 'devices', values['devices'], ClusterStorageDeviceInfo);
        } else {
            this['devices'] = [];
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'devices': new FormArray([]),
            });
            // generate FormArray control elements
            this.fillFormArray<ClusterStorageDeviceInfo>('devices', this['devices'], ClusterStorageDeviceInfo);
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('devices') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('devices').get(field);
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
            this.fillModelArray<ClusterStorageDeviceInfo>(this, 'devices', this['devices'], ClusterStorageDeviceInfo);
        }
    }
}

