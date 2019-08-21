/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { AuthAutoMsgUserPreferenceWatchHelperWatchEvent, IAuthAutoMsgUserPreferenceWatchHelperWatchEvent } from './auth-auto-msg-user-preference-watch-helper-watch-event.model';

export interface IAuthAutoMsgUserPreferenceWatchHelper {
    'events'?: Array<IAuthAutoMsgUserPreferenceWatchHelperWatchEvent>;
}


export class AuthAutoMsgUserPreferenceWatchHelper extends BaseModel implements IAuthAutoMsgUserPreferenceWatchHelper {
    'events': Array<AuthAutoMsgUserPreferenceWatchHelperWatchEvent> = null;
    public static propInfo: { [prop in keyof IAuthAutoMsgUserPreferenceWatchHelper]: PropInfoItem } = {
        'events': {
            required: false,
            type: 'object'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return AuthAutoMsgUserPreferenceWatchHelper.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return AuthAutoMsgUserPreferenceWatchHelper.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (AuthAutoMsgUserPreferenceWatchHelper.propInfo[prop] != null &&
                        AuthAutoMsgUserPreferenceWatchHelper.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['events'] = new Array<AuthAutoMsgUserPreferenceWatchHelperWatchEvent>();
        this._inputValue = values;
        this.setValues(values, setDefaults);
    }

    /**
     * set the values for both the Model and the Form Group. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any, fillDefaults = true): void {
        if (values) {
            this.fillModelArray<AuthAutoMsgUserPreferenceWatchHelperWatchEvent>(this, 'events', values['events'], AuthAutoMsgUserPreferenceWatchHelperWatchEvent);
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
            this.fillFormArray<AuthAutoMsgUserPreferenceWatchHelperWatchEvent>('events', this['events'], AuthAutoMsgUserPreferenceWatchHelperWatchEvent);
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
            this.fillModelArray<AuthAutoMsgUserPreferenceWatchHelperWatchEvent>(this, 'events', this['events'], AuthAutoMsgUserPreferenceWatchHelperWatchEvent);
        }
    }
}

