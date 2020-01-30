/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';


export interface INetworkRoute {
    'prefix': string;
    'next-hop': string;
    'target-virtual-router'?: string;
}


export class NetworkRoute extends BaseModel implements INetworkRoute {
    /** Route Prefix for the route. Should be a valid v4 or v6 CIDR block. */
    'prefix': string = null;
    /** NextHop for the route. Should be a valid v4 or v6 CIDR block. */
    'next-hop': string = null;
    /** Target VirtualRouter instance. */
    'target-virtual-router': string = null;
    public static propInfo: { [prop in keyof INetworkRoute]: PropInfoItem } = {
        'prefix': {
            description:  `Route Prefix for the route. Should be a valid v4 or v6 CIDR block.`,
            hint:  '10.1.1.1/24, ff02::5/32 ',
            required: true,
            type: 'string'
        },
        'next-hop': {
            description:  `NextHop for the route. Should be a valid v4 or v6 CIDR block.`,
            hint:  '10.1.1.1/24, ff02::5/32 ',
            required: true,
            type: 'string'
        },
        'target-virtual-router': {
            description:  `Target VirtualRouter instance.`,
            required: false,
            type: 'string'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return NetworkRoute.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return NetworkRoute.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (NetworkRoute.propInfo[prop] != null &&
                        NetworkRoute.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this._inputValue = values;
        this.setValues(values, setDefaults);
    }

    /**
     * set the values for both the Model and the Form Group. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any, fillDefaults = true): void {
        if (values && values['prefix'] != null) {
            this['prefix'] = values['prefix'];
        } else if (fillDefaults && NetworkRoute.hasDefaultValue('prefix')) {
            this['prefix'] = NetworkRoute.propInfo['prefix'].default;
        } else {
            this['prefix'] = null
        }
        if (values && values['next-hop'] != null) {
            this['next-hop'] = values['next-hop'];
        } else if (fillDefaults && NetworkRoute.hasDefaultValue('next-hop')) {
            this['next-hop'] = NetworkRoute.propInfo['next-hop'].default;
        } else {
            this['next-hop'] = null
        }
        if (values && values['target-virtual-router'] != null) {
            this['target-virtual-router'] = values['target-virtual-router'];
        } else if (fillDefaults && NetworkRoute.hasDefaultValue('target-virtual-router')) {
            this['target-virtual-router'] = NetworkRoute.propInfo['target-virtual-router'].default;
        } else {
            this['target-virtual-router'] = null
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'prefix': CustomFormControl(new FormControl(this['prefix'], [required, ]), NetworkRoute.propInfo['prefix']),
                'next-hop': CustomFormControl(new FormControl(this['next-hop'], [required, ]), NetworkRoute.propInfo['next-hop']),
                'target-virtual-router': CustomFormControl(new FormControl(this['target-virtual-router']), NetworkRoute.propInfo['target-virtual-router']),
            });
        }
        return this._formGroup;
    }

    setModelToBeFormGroupValues() {
        this.setValues(this.$formGroup.value, false);
    }

    setFormGroupValuesToBeModelValues() {
        if (this._formGroup) {
            this._formGroup.controls['prefix'].setValue(this['prefix']);
            this._formGroup.controls['next-hop'].setValue(this['next-hop']);
            this._formGroup.controls['target-virtual-router'].setValue(this['target-virtual-router']);
        }
    }
}

