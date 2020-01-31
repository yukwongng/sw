/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { ObjstoreAutoMsgBucketWatchHelperWatchEvent, IObjstoreAutoMsgBucketWatchHelperWatchEvent } from './objstore-auto-msg-bucket-watch-helper-watch-event.model';

export interface IObjstoreAutoMsgBucketWatchHelper {
    'events'?: Array<IObjstoreAutoMsgBucketWatchHelperWatchEvent>;
    '_ui'?: any;
}


export class ObjstoreAutoMsgBucketWatchHelper extends BaseModel implements IObjstoreAutoMsgBucketWatchHelper {
    /** Field for holding arbitrary ui state */
    '_ui': any = {};
    'events': Array<ObjstoreAutoMsgBucketWatchHelperWatchEvent> = null;
    public static propInfo: { [prop in keyof IObjstoreAutoMsgBucketWatchHelper]: PropInfoItem } = {
        'events': {
            required: false,
            type: 'object'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return ObjstoreAutoMsgBucketWatchHelper.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return ObjstoreAutoMsgBucketWatchHelper.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (ObjstoreAutoMsgBucketWatchHelper.propInfo[prop] != null &&
                        ObjstoreAutoMsgBucketWatchHelper.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['events'] = new Array<ObjstoreAutoMsgBucketWatchHelperWatchEvent>();
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
            this.fillModelArray<ObjstoreAutoMsgBucketWatchHelperWatchEvent>(this, 'events', values['events'], ObjstoreAutoMsgBucketWatchHelperWatchEvent);
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
            this.fillFormArray<ObjstoreAutoMsgBucketWatchHelperWatchEvent>('events', this['events'], ObjstoreAutoMsgBucketWatchHelperWatchEvent);
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
            this.fillModelArray<ObjstoreAutoMsgBucketWatchHelperWatchEvent>(this, 'events', this['events'], ObjstoreAutoMsgBucketWatchHelperWatchEvent);
        }
    }
}

