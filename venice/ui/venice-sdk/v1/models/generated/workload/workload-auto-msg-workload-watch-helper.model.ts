/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { WorkloadAutoMsgWorkloadWatchHelperWatchEvent, IWorkloadAutoMsgWorkloadWatchHelperWatchEvent } from './workload-auto-msg-workload-watch-helper-watch-event.model';

export interface IWorkloadAutoMsgWorkloadWatchHelper {
    'events'?: Array<IWorkloadAutoMsgWorkloadWatchHelperWatchEvent>;
}


export class WorkloadAutoMsgWorkloadWatchHelper extends BaseModel implements IWorkloadAutoMsgWorkloadWatchHelper {
    'events': Array<WorkloadAutoMsgWorkloadWatchHelperWatchEvent> = null;
    public static propInfo: { [prop in keyof IWorkloadAutoMsgWorkloadWatchHelper]: PropInfoItem } = {
        'events': {
            required: false,
            type: 'object'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return WorkloadAutoMsgWorkloadWatchHelper.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return WorkloadAutoMsgWorkloadWatchHelper.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (WorkloadAutoMsgWorkloadWatchHelper.propInfo[prop] != null &&
                        WorkloadAutoMsgWorkloadWatchHelper.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['events'] = new Array<WorkloadAutoMsgWorkloadWatchHelperWatchEvent>();
        this._inputValue = values;
        this.setValues(values, setDefaults);
    }

    /**
     * set the values for both the Model and the Form Group. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any, fillDefaults = true): void {
        if (values) {
            this.fillModelArray<WorkloadAutoMsgWorkloadWatchHelperWatchEvent>(this, 'events', values['events'], WorkloadAutoMsgWorkloadWatchHelperWatchEvent);
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
            this.fillFormArray<WorkloadAutoMsgWorkloadWatchHelperWatchEvent>('events', this['events'], WorkloadAutoMsgWorkloadWatchHelperWatchEvent);
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
            this.fillModelArray<WorkloadAutoMsgWorkloadWatchHelperWatchEvent>(this, 'events', this['events'], WorkloadAutoMsgWorkloadWatchHelperWatchEvent);
        }
    }
}

