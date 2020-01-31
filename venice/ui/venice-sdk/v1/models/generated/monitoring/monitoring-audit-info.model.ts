/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';


export interface IMonitoringAuditInfo {
    'user'?: string;
    'time'?: Date;
    '_ui'?: any;
}


export class MonitoringAuditInfo extends BaseModel implements IMonitoringAuditInfo {
    /** Field for holding arbitrary ui state */
    '_ui': any = {};
    /** Name of the user performed some action. */
    'user': string = null;
    /** Time at which the action was performed. */
    'time': Date = null;
    public static propInfo: { [prop in keyof IMonitoringAuditInfo]: PropInfoItem } = {
        'user': {
            description:  `Name of the user performed some action.`,
            required: false,
            type: 'string'
        },
        'time': {
            description:  `Time at which the action was performed.`,
            required: false,
            type: 'Date'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return MonitoringAuditInfo.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return MonitoringAuditInfo.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (MonitoringAuditInfo.propInfo[prop] != null &&
                        MonitoringAuditInfo.propInfo[prop].default != null);
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
        if (values && values['user'] != null) {
            this['user'] = values['user'];
        } else if (fillDefaults && MonitoringAuditInfo.hasDefaultValue('user')) {
            this['user'] = MonitoringAuditInfo.propInfo['user'].default;
        } else {
            this['user'] = null
        }
        if (values && values['time'] != null) {
            this['time'] = values['time'];
        } else if (fillDefaults && MonitoringAuditInfo.hasDefaultValue('time')) {
            this['time'] = MonitoringAuditInfo.propInfo['time'].default;
        } else {
            this['time'] = null
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'user': CustomFormControl(new FormControl(this['user']), MonitoringAuditInfo.propInfo['user']),
                'time': CustomFormControl(new FormControl(this['time']), MonitoringAuditInfo.propInfo['time']),
            });
        }
        return this._formGroup;
    }

    setModelToBeFormGroupValues() {
        this.setValues(this.$formGroup.value, false);
    }

    setFormGroupValuesToBeModelValues() {
        if (this._formGroup) {
            this._formGroup.controls['user'].setValue(this['user']);
            this._formGroup.controls['time'].setValue(this['time']);
        }
    }
}

