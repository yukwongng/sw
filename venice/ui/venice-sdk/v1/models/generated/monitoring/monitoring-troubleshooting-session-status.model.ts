/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { MonitoringTroubleshootingSessionStatus_state,  MonitoringTroubleshootingSessionStatus_state_uihint  } from './enums';
import { MonitoringTsResult, IMonitoringTsResult } from './monitoring-ts-result.model';

export interface IMonitoringTroubleshootingSessionStatus {
    'state': MonitoringTroubleshootingSessionStatus_state;
    'troubleshooting-results'?: Array<IMonitoringTsResult>;
    '_ui'?: any;
}


export class MonitoringTroubleshootingSessionStatus extends BaseModel implements IMonitoringTroubleshootingSessionStatus {
    /** Field for holding arbitrary ui state */
    '_ui': any = {};
    'state': MonitoringTroubleshootingSessionStatus_state = null;
    /** Report is generated each time troubleshooting session is activated i.e time-window. */
    'troubleshooting-results': Array<MonitoringTsResult> = null;
    public static propInfo: { [prop in keyof IMonitoringTroubleshootingSessionStatus]: PropInfoItem } = {
        'state': {
            enum: MonitoringTroubleshootingSessionStatus_state_uihint,
            default: 'running',
            required: true,
            type: 'string'
        },
        'troubleshooting-results': {
            description:  `Report is generated each time troubleshooting session is activated i.e time-window.`,
            required: false,
            type: 'object'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return MonitoringTroubleshootingSessionStatus.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return MonitoringTroubleshootingSessionStatus.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (MonitoringTroubleshootingSessionStatus.propInfo[prop] != null &&
                        MonitoringTroubleshootingSessionStatus.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['troubleshooting-results'] = new Array<MonitoringTsResult>();
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
        if (values && values['state'] != null) {
            this['state'] = values['state'];
        } else if (fillDefaults && MonitoringTroubleshootingSessionStatus.hasDefaultValue('state')) {
            this['state'] = <MonitoringTroubleshootingSessionStatus_state>  MonitoringTroubleshootingSessionStatus.propInfo['state'].default;
        } else {
            this['state'] = null
        }
        if (values) {
            this.fillModelArray<MonitoringTsResult>(this, 'troubleshooting-results', values['troubleshooting-results'], MonitoringTsResult);
        } else {
            this['troubleshooting-results'] = [];
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'state': CustomFormControl(new FormControl(this['state'], [required, enumValidator(MonitoringTroubleshootingSessionStatus_state), ]), MonitoringTroubleshootingSessionStatus.propInfo['state']),
                'troubleshooting-results': new FormArray([]),
            });
            // generate FormArray control elements
            this.fillFormArray<MonitoringTsResult>('troubleshooting-results', this['troubleshooting-results'], MonitoringTsResult);
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('troubleshooting-results') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('troubleshooting-results').get(field);
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
            this._formGroup.controls['state'].setValue(this['state']);
            this.fillModelArray<MonitoringTsResult>(this, 'troubleshooting-results', this['troubleshooting-results'], MonitoringTsResult);
        }
    }
}

