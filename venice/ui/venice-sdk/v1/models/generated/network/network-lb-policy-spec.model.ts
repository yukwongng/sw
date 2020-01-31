/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { NetworkHealthCheckSpec, INetworkHealthCheckSpec } from './network-health-check-spec.model';

export interface INetworkLbPolicySpec {
    'type'?: string;
    'algorithm'?: string;
    'session-affinity'?: string;
    'health-check'?: INetworkHealthCheckSpec;
    '_ui'?: any;
}


export class NetworkLbPolicySpec extends BaseModel implements INetworkLbPolicySpec {
    /** Field for holding arbitrary ui state */
    '_ui': any = {};
    /** Load balancing type. */
    'type': string = null;
    /** Load balancing algorithm. */
    'algorithm': string = null;
    /** Session affinity. */
    'session-affinity': string = null;
    /** Health check policy. */
    'health-check': NetworkHealthCheckSpec = null;
    public static propInfo: { [prop in keyof INetworkLbPolicySpec]: PropInfoItem } = {
        'type': {
            description:  `Load balancing type.`,
            required: false,
            type: 'string'
        },
        'algorithm': {
            description:  `Load balancing algorithm.`,
            required: false,
            type: 'string'
        },
        'session-affinity': {
            description:  `Session affinity.`,
            required: false,
            type: 'string'
        },
        'health-check': {
            description:  `Health check policy.`,
            required: false,
            type: 'object'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return NetworkLbPolicySpec.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return NetworkLbPolicySpec.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (NetworkLbPolicySpec.propInfo[prop] != null &&
                        NetworkLbPolicySpec.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['health-check'] = new NetworkHealthCheckSpec();
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
        if (values && values['type'] != null) {
            this['type'] = values['type'];
        } else if (fillDefaults && NetworkLbPolicySpec.hasDefaultValue('type')) {
            this['type'] = NetworkLbPolicySpec.propInfo['type'].default;
        } else {
            this['type'] = null
        }
        if (values && values['algorithm'] != null) {
            this['algorithm'] = values['algorithm'];
        } else if (fillDefaults && NetworkLbPolicySpec.hasDefaultValue('algorithm')) {
            this['algorithm'] = NetworkLbPolicySpec.propInfo['algorithm'].default;
        } else {
            this['algorithm'] = null
        }
        if (values && values['session-affinity'] != null) {
            this['session-affinity'] = values['session-affinity'];
        } else if (fillDefaults && NetworkLbPolicySpec.hasDefaultValue('session-affinity')) {
            this['session-affinity'] = NetworkLbPolicySpec.propInfo['session-affinity'].default;
        } else {
            this['session-affinity'] = null
        }
        if (values) {
            this['health-check'].setValues(values['health-check'], fillDefaults);
        } else {
            this['health-check'].setValues(null, fillDefaults);
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'type': CustomFormControl(new FormControl(this['type']), NetworkLbPolicySpec.propInfo['type']),
                'algorithm': CustomFormControl(new FormControl(this['algorithm']), NetworkLbPolicySpec.propInfo['algorithm']),
                'session-affinity': CustomFormControl(new FormControl(this['session-affinity']), NetworkLbPolicySpec.propInfo['session-affinity']),
                'health-check': CustomFormGroup(this['health-check'].$formGroup, NetworkLbPolicySpec.propInfo['health-check'].required),
            });
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('health-check') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('health-check').get(field);
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
            this._formGroup.controls['type'].setValue(this['type']);
            this._formGroup.controls['algorithm'].setValue(this['algorithm']);
            this._formGroup.controls['session-affinity'].setValue(this['session-affinity']);
            this['health-check'].setFormGroupValuesToBeModelValues();
        }
    }
}

