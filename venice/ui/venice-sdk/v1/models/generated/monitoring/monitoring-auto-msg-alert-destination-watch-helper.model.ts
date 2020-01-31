/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { MonitoringAutoMsgAlertDestinationWatchHelperWatchEvent, IMonitoringAutoMsgAlertDestinationWatchHelperWatchEvent } from './monitoring-auto-msg-alert-destination-watch-helper-watch-event.model';

export interface IMonitoringAutoMsgAlertDestinationWatchHelper {
    'events'?: Array<IMonitoringAutoMsgAlertDestinationWatchHelperWatchEvent>;
    '_ui'?: any;
}


export class MonitoringAutoMsgAlertDestinationWatchHelper extends BaseModel implements IMonitoringAutoMsgAlertDestinationWatchHelper {
    /** Field for holding arbitrary ui state */
    '_ui': any = {};
    'events': Array<MonitoringAutoMsgAlertDestinationWatchHelperWatchEvent> = null;
    public static propInfo: { [prop in keyof IMonitoringAutoMsgAlertDestinationWatchHelper]: PropInfoItem } = {
        'events': {
            required: false,
            type: 'object'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return MonitoringAutoMsgAlertDestinationWatchHelper.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return MonitoringAutoMsgAlertDestinationWatchHelper.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (MonitoringAutoMsgAlertDestinationWatchHelper.propInfo[prop] != null &&
                        MonitoringAutoMsgAlertDestinationWatchHelper.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['events'] = new Array<MonitoringAutoMsgAlertDestinationWatchHelperWatchEvent>();
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
        if (values) {
            this.fillModelArray<MonitoringAutoMsgAlertDestinationWatchHelperWatchEvent>(this, 'events', values['events'], MonitoringAutoMsgAlertDestinationWatchHelperWatchEvent);
        } else {
            this['events'] = [];
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'events': new FormArray([]),
            });
            // generate FormArray control elements
            this.fillFormArray<MonitoringAutoMsgAlertDestinationWatchHelperWatchEvent>('events', this['events'], MonitoringAutoMsgAlertDestinationWatchHelperWatchEvent);
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('events') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('events').get(field);
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
            this.fillModelArray<MonitoringAutoMsgAlertDestinationWatchHelperWatchEvent>(this, 'events', this['events'], MonitoringAutoMsgAlertDestinationWatchHelperWatchEvent);
        }
    }
}

