/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, enumValidator } from './validators';
import { BaseModel, EnumDef } from './base-model';

import { MonitoringAlertDestination, IMonitoringAlertDestination } from './monitoring-alert-destination.model';

export interface IMonitoringAutoMsgAlertDestinationWatchHelperWatchEvent {
    'Type'?: string;
    'Object'?: IMonitoringAlertDestination;
}


export class MonitoringAutoMsgAlertDestinationWatchHelperWatchEvent extends BaseModel implements IMonitoringAutoMsgAlertDestinationWatchHelperWatchEvent {
    'Type': string;
    'Object': MonitoringAlertDestination;
    public static enumProperties: { [key: string] : EnumDef } = {
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any) {
        super();
        this['Object'] = new MonitoringAlertDestination();
        if (values) {
            this.setValues(values);
        }
    }

    /**
     * set the values.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any): void {
        if (values) {
            this['Type'] = values['Type'];
            this['Object'].setValues(values['Object']);
        }
    }

    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'Type': new FormControl(this['Type']),
                'Object': this['Object'].$formGroup,
            });
        }
        return this._formGroup;
    }

    setFormGroupValues() {
        if (this._formGroup) {
            this._formGroup.controls['Type'].setValue(this['Type']);
            this['Object'].setFormGroupValues();
        }
    }
}

