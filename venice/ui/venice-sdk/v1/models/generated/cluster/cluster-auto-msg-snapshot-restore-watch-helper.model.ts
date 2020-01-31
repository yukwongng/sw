/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { ClusterAutoMsgSnapshotRestoreWatchHelperWatchEvent, IClusterAutoMsgSnapshotRestoreWatchHelperWatchEvent } from './cluster-auto-msg-snapshot-restore-watch-helper-watch-event.model';

export interface IClusterAutoMsgSnapshotRestoreWatchHelper {
    'events'?: Array<IClusterAutoMsgSnapshotRestoreWatchHelperWatchEvent>;
    '_ui'?: any;
}


export class ClusterAutoMsgSnapshotRestoreWatchHelper extends BaseModel implements IClusterAutoMsgSnapshotRestoreWatchHelper {
    /** Field for holding arbitrary ui state */
    '_ui': any = {};
    'events': Array<ClusterAutoMsgSnapshotRestoreWatchHelperWatchEvent> = null;
    public static propInfo: { [prop in keyof IClusterAutoMsgSnapshotRestoreWatchHelper]: PropInfoItem } = {
        'events': {
            required: false,
            type: 'object'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return ClusterAutoMsgSnapshotRestoreWatchHelper.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return ClusterAutoMsgSnapshotRestoreWatchHelper.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (ClusterAutoMsgSnapshotRestoreWatchHelper.propInfo[prop] != null &&
                        ClusterAutoMsgSnapshotRestoreWatchHelper.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['events'] = new Array<ClusterAutoMsgSnapshotRestoreWatchHelperWatchEvent>();
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
            this.fillModelArray<ClusterAutoMsgSnapshotRestoreWatchHelperWatchEvent>(this, 'events', values['events'], ClusterAutoMsgSnapshotRestoreWatchHelperWatchEvent);
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
            this.fillFormArray<ClusterAutoMsgSnapshotRestoreWatchHelperWatchEvent>('events', this['events'], ClusterAutoMsgSnapshotRestoreWatchHelperWatchEvent);
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
            this.fillModelArray<ClusterAutoMsgSnapshotRestoreWatchHelperWatchEvent>(this, 'events', this['events'], ClusterAutoMsgSnapshotRestoreWatchHelperWatchEvent);
        }
    }
}

