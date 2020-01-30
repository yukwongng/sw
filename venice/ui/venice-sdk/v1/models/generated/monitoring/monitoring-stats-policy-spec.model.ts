/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';


export interface IMonitoringStatsPolicySpec {
    'retention-time': string;
    'downsample-retention-time': string;
}


export class MonitoringStatsPolicySpec extends BaseModel implements IMonitoringStatsPolicySpec {
    /** RetentionTime defines for how long to keep the stats data before it is deleted The value is specified as a string format to be hours, days, or months etc. e.g. '24hrs', '72hours', '4days', '6d', '2months', '4mo', '1yr' Default is 48h. Should be a valid time duration. */
    'retention-time': string = null;
    /** DownSampleRetentionTime defines for how long to keep the down sampled data before it is deleted The value is specified as a string format to be hours, days, or months etc. e.g. '24hrs', '72hours', '4days', '6d', '2months', '4mo', '1yr' Default is 168h. Should be a valid time duration. */
    'downsample-retention-time': string = null;
    public static propInfo: { [prop in keyof IMonitoringStatsPolicySpec]: PropInfoItem } = {
        'retention-time': {
            default: '48h',
            description:  `RetentionTime defines for how long to keep the stats data before it is deleted The value is specified as a string format to be hours, days, or months etc. e.g. '24hrs', '72hours', '4days', '6d', '2months', '4mo', '1yr' Default is 48h. Should be a valid time duration.`,
            hint:  '2h',
            required: true,
            type: 'string'
        },
        'downsample-retention-time': {
            default: '168h',
            description:  `DownSampleRetentionTime defines for how long to keep the down sampled data before it is deleted The value is specified as a string format to be hours, days, or months etc. e.g. '24hrs', '72hours', '4days', '6d', '2months', '4mo', '1yr' Default is 168h. Should be a valid time duration.`,
            hint:  '2h',
            required: true,
            type: 'string'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return MonitoringStatsPolicySpec.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return MonitoringStatsPolicySpec.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (MonitoringStatsPolicySpec.propInfo[prop] != null &&
                        MonitoringStatsPolicySpec.propInfo[prop].default != null);
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
        if (values && values['retention-time'] != null) {
            this['retention-time'] = values['retention-time'];
        } else if (fillDefaults && MonitoringStatsPolicySpec.hasDefaultValue('retention-time')) {
            this['retention-time'] = MonitoringStatsPolicySpec.propInfo['retention-time'].default;
        } else {
            this['retention-time'] = null
        }
        if (values && values['downsample-retention-time'] != null) {
            this['downsample-retention-time'] = values['downsample-retention-time'];
        } else if (fillDefaults && MonitoringStatsPolicySpec.hasDefaultValue('downsample-retention-time')) {
            this['downsample-retention-time'] = MonitoringStatsPolicySpec.propInfo['downsample-retention-time'].default;
        } else {
            this['downsample-retention-time'] = null
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'retention-time': CustomFormControl(new FormControl(this['retention-time'], [required, ]), MonitoringStatsPolicySpec.propInfo['retention-time']),
                'downsample-retention-time': CustomFormControl(new FormControl(this['downsample-retention-time'], [required, ]), MonitoringStatsPolicySpec.propInfo['downsample-retention-time']),
            });
        }
        return this._formGroup;
    }

    setModelToBeFormGroupValues() {
        this.setValues(this.$formGroup.value, false);
    }

    setFormGroupValuesToBeModelValues() {
        if (this._formGroup) {
            this._formGroup.controls['retention-time'].setValue(this['retention-time']);
            this._formGroup.controls['downsample-retention-time'].setValue(this['downsample-retention-time']);
        }
    }
}

