/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { NetworkAutoMsgVirtualRouterWatchHelperWatchEvent, INetworkAutoMsgVirtualRouterWatchHelperWatchEvent } from './network-auto-msg-virtual-router-watch-helper-watch-event.model';

export interface INetworkAutoMsgVirtualRouterWatchHelper {
    'events'?: Array<INetworkAutoMsgVirtualRouterWatchHelperWatchEvent>;
}


export class NetworkAutoMsgVirtualRouterWatchHelper extends BaseModel implements INetworkAutoMsgVirtualRouterWatchHelper {
    'events': Array<NetworkAutoMsgVirtualRouterWatchHelperWatchEvent> = null;
    public static propInfo: { [prop in keyof INetworkAutoMsgVirtualRouterWatchHelper]: PropInfoItem } = {
        'events': {
            required: false,
            type: 'object'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return NetworkAutoMsgVirtualRouterWatchHelper.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return NetworkAutoMsgVirtualRouterWatchHelper.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (NetworkAutoMsgVirtualRouterWatchHelper.propInfo[prop] != null &&
                        NetworkAutoMsgVirtualRouterWatchHelper.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['events'] = new Array<NetworkAutoMsgVirtualRouterWatchHelperWatchEvent>();
        this._inputValue = values;
        this.setValues(values, setDefaults);
    }

    /**
     * set the values for both the Model and the Form Group. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any, fillDefaults = true): void {
        if (values) {
            this.fillModelArray<NetworkAutoMsgVirtualRouterWatchHelperWatchEvent>(this, 'events', values['events'], NetworkAutoMsgVirtualRouterWatchHelperWatchEvent);
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
            this.fillFormArray<NetworkAutoMsgVirtualRouterWatchHelperWatchEvent>('events', this['events'], NetworkAutoMsgVirtualRouterWatchHelperWatchEvent);
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
            this.fillModelArray<NetworkAutoMsgVirtualRouterWatchHelperWatchEvent>(this, 'events', this['events'], NetworkAutoMsgVirtualRouterWatchHelperWatchEvent);
        }
    }
}

