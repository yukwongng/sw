/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { ApiObjectMeta, IApiObjectMeta } from './api-object-meta.model';
import { NetworkNetworkInterfaceSpec, INetworkNetworkInterfaceSpec } from './network-network-interface-spec.model';
import { NetworkNetworkInterfaceStatus, INetworkNetworkInterfaceStatus } from './network-network-interface-status.model';

export interface INetworkNetworkInterface {
    'kind'?: string;
    'api-version'?: string;
    'meta'?: IApiObjectMeta;
    'spec'?: INetworkNetworkInterfaceSpec;
    'status'?: INetworkNetworkInterfaceStatus;
    '_ui'?: any;
}


export class NetworkNetworkInterface extends BaseModel implements INetworkNetworkInterface {
    /** Field for holding arbitrary ui state */
    '_ui': any = {};
    'kind': string = null;
    'api-version': string = null;
    /** Object name is Serial-Number of the SmartNIC. */
    'meta': ApiObjectMeta = null;
    /** NetworkInterfaceSpec contains the configuration of the network adapter. */
    'spec': NetworkNetworkInterfaceSpec = null;
    /** NetworkInterfaceStatus contains the current state of the network adapter. */
    'status': NetworkNetworkInterfaceStatus = null;
    public static propInfo: { [prop in keyof INetworkNetworkInterface]: PropInfoItem } = {
        'kind': {
            required: false,
            type: 'string'
        },
        'api-version': {
            required: false,
            type: 'string'
        },
        'meta': {
            description:  `Object name is Serial-Number of the SmartNIC.`,
            required: false,
            type: 'object'
        },
        'spec': {
            description:  `NetworkInterfaceSpec contains the configuration of the network adapter.`,
            required: false,
            type: 'object'
        },
        'status': {
            description:  `NetworkInterfaceStatus contains the current state of the network adapter.`,
            required: false,
            type: 'object'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return NetworkNetworkInterface.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return NetworkNetworkInterface.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (NetworkNetworkInterface.propInfo[prop] != null &&
                        NetworkNetworkInterface.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['meta'] = new ApiObjectMeta();
        this['spec'] = new NetworkNetworkInterfaceSpec();
        this['status'] = new NetworkNetworkInterfaceStatus();
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
        if (values && values['kind'] != null) {
            this['kind'] = values['kind'];
        } else if (fillDefaults && NetworkNetworkInterface.hasDefaultValue('kind')) {
            this['kind'] = NetworkNetworkInterface.propInfo['kind'].default;
        } else {
            this['kind'] = null
        }
        if (values && values['api-version'] != null) {
            this['api-version'] = values['api-version'];
        } else if (fillDefaults && NetworkNetworkInterface.hasDefaultValue('api-version')) {
            this['api-version'] = NetworkNetworkInterface.propInfo['api-version'].default;
        } else {
            this['api-version'] = null
        }
        if (values) {
            this['meta'].setValues(values['meta'], fillDefaults);
        } else {
            this['meta'].setValues(null, fillDefaults);
        }
        if (values) {
            this['spec'].setValues(values['spec'], fillDefaults);
        } else {
            this['spec'].setValues(null, fillDefaults);
        }
        if (values) {
            this['status'].setValues(values['status'], fillDefaults);
        } else {
            this['status'].setValues(null, fillDefaults);
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'kind': CustomFormControl(new FormControl(this['kind']), NetworkNetworkInterface.propInfo['kind']),
                'api-version': CustomFormControl(new FormControl(this['api-version']), NetworkNetworkInterface.propInfo['api-version']),
                'meta': CustomFormGroup(this['meta'].$formGroup, NetworkNetworkInterface.propInfo['meta'].required),
                'spec': CustomFormGroup(this['spec'].$formGroup, NetworkNetworkInterface.propInfo['spec'].required),
                'status': CustomFormGroup(this['status'].$formGroup, NetworkNetworkInterface.propInfo['status'].required),
            });
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('meta') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('meta').get(field);
                control.updateValueAndValidity();
            });
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('spec') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('spec').get(field);
                control.updateValueAndValidity();
            });
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('status') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('status').get(field);
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
            this._formGroup.controls['kind'].setValue(this['kind']);
            this._formGroup.controls['api-version'].setValue(this['api-version']);
            this['meta'].setFormGroupValuesToBeModelValues();
            this['spec'].setFormGroupValuesToBeModelValues();
            this['status'].setFormGroupValuesToBeModelValues();
        }
    }
}

