/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { NetworkRoute, INetworkRoute } from './network-route.model';

export interface INetworkRouteTableStatus {
    'routes'?: Array<INetworkRoute>;
}


export class NetworkRouteTableStatus extends BaseModel implements INetworkRouteTableStatus {
    'routes': Array<NetworkRoute> = null;
    public static propInfo: { [prop in keyof INetworkRouteTableStatus]: PropInfoItem } = {
        'routes': {
            required: false,
            type: 'object'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return NetworkRouteTableStatus.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return NetworkRouteTableStatus.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (NetworkRouteTableStatus.propInfo[prop] != null &&
                        NetworkRouteTableStatus.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['routes'] = new Array<NetworkRoute>();
        this._inputValue = values;
        this.setValues(values, setDefaults);
    }

    /**
     * set the values for both the Model and the Form Group. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any, fillDefaults = true): void {
        if (values) {
            this.fillModelArray<NetworkRoute>(this, 'routes', values['routes'], NetworkRoute);
        } else {
            this['routes'] = [];
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'routes': new FormArray([]),
            });
            // generate FormArray control elements
            this.fillFormArray<NetworkRoute>('routes', this['routes'], NetworkRoute);
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('routes') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('routes').get(field);
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
            this.fillModelArray<NetworkRoute>(this, 'routes', this['routes'], NetworkRoute);
        }
    }
}
