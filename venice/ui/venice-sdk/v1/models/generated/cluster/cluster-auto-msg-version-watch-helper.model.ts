/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { ClusterAutoMsgVersionWatchHelperWatchEvent, IClusterAutoMsgVersionWatchHelperWatchEvent } from './cluster-auto-msg-version-watch-helper-watch-event.model';

export interface IClusterAutoMsgVersionWatchHelper {
    'events'?: Array<IClusterAutoMsgVersionWatchHelperWatchEvent>;
    '_ui'?: any;
}


export class ClusterAutoMsgVersionWatchHelper extends BaseModel implements IClusterAutoMsgVersionWatchHelper {
    /** Field for holding arbitrary ui state */
    '_ui': any = {};
    'events': Array<ClusterAutoMsgVersionWatchHelperWatchEvent> = null;
    public static propInfo: { [prop in keyof IClusterAutoMsgVersionWatchHelper]: PropInfoItem } = {
        'events': {
            required: false,
            type: 'object'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return ClusterAutoMsgVersionWatchHelper.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return ClusterAutoMsgVersionWatchHelper.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (ClusterAutoMsgVersionWatchHelper.propInfo[prop] != null &&
                        ClusterAutoMsgVersionWatchHelper.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['events'] = new Array<ClusterAutoMsgVersionWatchHelperWatchEvent>();
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
            this.fillModelArray<ClusterAutoMsgVersionWatchHelperWatchEvent>(this, 'events', values['events'], ClusterAutoMsgVersionWatchHelperWatchEvent);
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
            this.fillFormArray<ClusterAutoMsgVersionWatchHelperWatchEvent>('events', this['events'], ClusterAutoMsgVersionWatchHelperWatchEvent);
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
            this.fillModelArray<ClusterAutoMsgVersionWatchHelperWatchEvent>(this, 'events', this['events'], ClusterAutoMsgVersionWatchHelperWatchEvent);
        }
    }
}

