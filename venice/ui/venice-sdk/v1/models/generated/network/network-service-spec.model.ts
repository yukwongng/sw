/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { NetworkTLSServerPolicySpec, INetworkTLSServerPolicySpec } from './network-tls-server-policy-spec.model';
import { NetworkTLSClientPolicySpec, INetworkTLSClientPolicySpec } from './network-tls-client-policy-spec.model';

export interface INetworkServiceSpec {
    'workload-labels'?: Array<string>;
    'virtual-ip'?: string;
    'ports'?: string;
    'lb-policy'?: string;
    'tls-server-policy'?: INetworkTLSServerPolicySpec;
    'tls-client-policy'?: INetworkTLSClientPolicySpec;
    '_ui'?: any;
}


export class NetworkServiceSpec extends BaseModel implements INetworkServiceSpec {
    /** Field for holding arbitrary ui state */
    '_ui': any = {};
    /** FIXME: maps are not working. change this after its fixed map<string, string> WorkloadSelector  = 1 [(gogoproto.nullable) = true, (gogoproto.jsontag) = "workload-labels,omitempty"]; workload selector for the service (list of labels to match). */
    'workload-labels': Array<string> = null;
    /** Virtual IP of the load balancer. */
    'virtual-ip': string = null;
    /** Load balancer port. */
    'ports': string = null;
    /** Load balancing policy (reference to LbPolicy object). */
    'lb-policy': string = null;
    /** TLS configuration for inbound connections. */
    'tls-server-policy': NetworkTLSServerPolicySpec = null;
    /** TLS configuration for outbound connections. */
    'tls-client-policy': NetworkTLSClientPolicySpec = null;
    public static propInfo: { [prop in keyof INetworkServiceSpec]: PropInfoItem } = {
        'workload-labels': {
            description:  `FIXME: maps are not working. change this after its fixed map<string, string> WorkloadSelector  = 1 [(gogoproto.nullable) = true, (gogoproto.jsontag) = "workload-labels,omitempty"]; workload selector for the service (list of labels to match).`,
            required: false,
            type: 'Array<string>'
        },
        'virtual-ip': {
            description:  `Virtual IP of the load balancer.`,
            required: false,
            type: 'string'
        },
        'ports': {
            description:  `Load balancer port.`,
            required: false,
            type: 'string'
        },
        'lb-policy': {
            description:  `Load balancing policy (reference to LbPolicy object).`,
            required: false,
            type: 'string'
        },
        'tls-server-policy': {
            description:  `TLS configuration for inbound connections.`,
            required: false,
            type: 'object'
        },
        'tls-client-policy': {
            description:  `TLS configuration for outbound connections.`,
            required: false,
            type: 'object'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return NetworkServiceSpec.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return NetworkServiceSpec.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (NetworkServiceSpec.propInfo[prop] != null &&
                        NetworkServiceSpec.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['workload-labels'] = new Array<string>();
        this['tls-server-policy'] = new NetworkTLSServerPolicySpec();
        this['tls-client-policy'] = new NetworkTLSClientPolicySpec();
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
        if (values && values['workload-labels'] != null) {
            this['workload-labels'] = values['workload-labels'];
        } else if (fillDefaults && NetworkServiceSpec.hasDefaultValue('workload-labels')) {
            this['workload-labels'] = [ NetworkServiceSpec.propInfo['workload-labels'].default];
        } else {
            this['workload-labels'] = [];
        }
        if (values && values['virtual-ip'] != null) {
            this['virtual-ip'] = values['virtual-ip'];
        } else if (fillDefaults && NetworkServiceSpec.hasDefaultValue('virtual-ip')) {
            this['virtual-ip'] = NetworkServiceSpec.propInfo['virtual-ip'].default;
        } else {
            this['virtual-ip'] = null
        }
        if (values && values['ports'] != null) {
            this['ports'] = values['ports'];
        } else if (fillDefaults && NetworkServiceSpec.hasDefaultValue('ports')) {
            this['ports'] = NetworkServiceSpec.propInfo['ports'].default;
        } else {
            this['ports'] = null
        }
        if (values && values['lb-policy'] != null) {
            this['lb-policy'] = values['lb-policy'];
        } else if (fillDefaults && NetworkServiceSpec.hasDefaultValue('lb-policy')) {
            this['lb-policy'] = NetworkServiceSpec.propInfo['lb-policy'].default;
        } else {
            this['lb-policy'] = null
        }
        if (values) {
            this['tls-server-policy'].setValues(values['tls-server-policy'], fillDefaults);
        } else {
            this['tls-server-policy'].setValues(null, fillDefaults);
        }
        if (values) {
            this['tls-client-policy'].setValues(values['tls-client-policy'], fillDefaults);
        } else {
            this['tls-client-policy'].setValues(null, fillDefaults);
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'workload-labels': CustomFormControl(new FormControl(this['workload-labels']), NetworkServiceSpec.propInfo['workload-labels']),
                'virtual-ip': CustomFormControl(new FormControl(this['virtual-ip']), NetworkServiceSpec.propInfo['virtual-ip']),
                'ports': CustomFormControl(new FormControl(this['ports']), NetworkServiceSpec.propInfo['ports']),
                'lb-policy': CustomFormControl(new FormControl(this['lb-policy']), NetworkServiceSpec.propInfo['lb-policy']),
                'tls-server-policy': CustomFormGroup(this['tls-server-policy'].$formGroup, NetworkServiceSpec.propInfo['tls-server-policy'].required),
                'tls-client-policy': CustomFormGroup(this['tls-client-policy'].$formGroup, NetworkServiceSpec.propInfo['tls-client-policy'].required),
            });
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('tls-server-policy') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('tls-server-policy').get(field);
                control.updateValueAndValidity();
            });
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('tls-client-policy') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('tls-client-policy').get(field);
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
            this._formGroup.controls['workload-labels'].setValue(this['workload-labels']);
            this._formGroup.controls['virtual-ip'].setValue(this['virtual-ip']);
            this._formGroup.controls['ports'].setValue(this['ports']);
            this._formGroup.controls['lb-policy'].setValue(this['lb-policy']);
            this['tls-server-policy'].setFormGroupValuesToBeModelValues();
            this['tls-client-policy'].setFormGroupValuesToBeModelValues();
        }
    }
}

