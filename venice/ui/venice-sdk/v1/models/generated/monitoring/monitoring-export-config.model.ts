/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { MonitoringExternalCred, IMonitoringExternalCred } from './monitoring-external-cred.model';

export interface IMonitoringExportConfig {
    'destination': string;
    'transport'?: string;
    'credentials'?: IMonitoringExternalCred;
}


export class MonitoringExportConfig extends BaseModel implements IMonitoringExportConfig {
    /** length of string should be between 1 and 2048 */
    'destination': string = null;
    /** should be a valid layer3 or layer 4 protocol and port/type */
    'transport': string = null;
    'credentials': MonitoringExternalCred = null;
    public static propInfo: { [prop in keyof IMonitoringExportConfig]: PropInfoItem } = {
        'destination': {
            description:  'length of string should be between 1 and 2048',
            required: true,
            type: 'string'
        },
        'transport': {
            description:  'should be a valid layer3 or layer 4 protocol and port/type',
            hint:  'tcp/1234, arp',
            required: false,
            type: 'string'
        },
        'credentials': {
            required: false,
            type: 'object'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return MonitoringExportConfig.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return MonitoringExportConfig.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (MonitoringExportConfig.propInfo[prop] != null &&
                        MonitoringExportConfig.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['credentials'] = new MonitoringExternalCred();
        this._inputValue = values;
        this.setValues(values, setDefaults);
    }

    /**
     * set the values for both the Model and the Form Group. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any, fillDefaults = true): void {
        if (values && values['destination'] != null) {
            this['destination'] = values['destination'];
        } else if (fillDefaults && MonitoringExportConfig.hasDefaultValue('destination')) {
            this['destination'] = MonitoringExportConfig.propInfo['destination'].default;
        } else {
            this['destination'] = null
        }
        if (values && values['transport'] != null) {
            this['transport'] = values['transport'];
        } else if (fillDefaults && MonitoringExportConfig.hasDefaultValue('transport')) {
            this['transport'] = MonitoringExportConfig.propInfo['transport'].default;
        } else {
            this['transport'] = null
        }
        if (values) {
            this['credentials'].setValues(values['credentials'], fillDefaults);
        } else {
            this['credentials'].setValues(null, fillDefaults);
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'destination': CustomFormControl(new FormControl(this['destination'], [required, minLengthValidator(1), maxLengthValidator(2048), ]), MonitoringExportConfig.propInfo['destination']),
                'transport': CustomFormControl(new FormControl(this['transport']), MonitoringExportConfig.propInfo['transport']),
                'credentials': CustomFormGroup(this['credentials'].$formGroup, MonitoringExportConfig.propInfo['credentials'].required),
            });
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('credentials') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('credentials').get(field);
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
            this._formGroup.controls['destination'].setValue(this['destination']);
            this._formGroup.controls['transport'].setValue(this['transport']);
            this['credentials'].setFormGroupValuesToBeModelValues();
        }
    }
}

