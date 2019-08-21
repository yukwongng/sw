/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { MonitoringAlertSpec_state,  MonitoringAlertSpec_state_uihint  } from './enums';

export interface IMonitoringAlertSpec {
    'state': MonitoringAlertSpec_state;
}


export class MonitoringAlertSpec extends BaseModel implements IMonitoringAlertSpec {
    'state': MonitoringAlertSpec_state = null;
    public static propInfo: { [prop in keyof IMonitoringAlertSpec]: PropInfoItem } = {
        'state': {
            enum: MonitoringAlertSpec_state_uihint,
            default: 'open',
            required: true,
            type: 'string'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return MonitoringAlertSpec.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return MonitoringAlertSpec.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (MonitoringAlertSpec.propInfo[prop] != null &&
                        MonitoringAlertSpec.propInfo[prop].default != null);
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
        if (values && values['state'] != null) {
            this['state'] = values['state'];
        } else if (fillDefaults && MonitoringAlertSpec.hasDefaultValue('state')) {
            this['state'] = <MonitoringAlertSpec_state>  MonitoringAlertSpec.propInfo['state'].default;
        } else {
            this['state'] = null
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'state': CustomFormControl(new FormControl(this['state'], [required, enumValidator(MonitoringAlertSpec_state), ]), MonitoringAlertSpec.propInfo['state']),
            });
        }
        return this._formGroup;
    }

    setModelToBeFormGroupValues() {
        this.setValues(this.$formGroup.value, false);
    }

    setFormGroupValuesToBeModelValues() {
        if (this._formGroup) {
            this._formGroup.controls['state'].setValue(this['state']);
        }
    }
}

