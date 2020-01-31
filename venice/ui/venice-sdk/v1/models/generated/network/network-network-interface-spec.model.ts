/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { NetworkNetworkInterfaceSpec_admin_status,  } from './enums';
import { NetworkPauseSpec, INetworkPauseSpec } from './network-pause-spec.model';
import { NetworkNetworkInterfaceSpec_type,  } from './enums';
import { NetworkNetworkInterfaceSpec_ip_alloc_type,  } from './enums';
import { ClusterIPConfig, IClusterIPConfig } from './cluster-ip-config.model';

export interface INetworkNetworkInterfaceSpec {
    'admin-status': NetworkNetworkInterfaceSpec_admin_status;
    'speed'?: string;
    'mtu'?: number;
    'pause'?: INetworkPauseSpec;
    'type': NetworkNetworkInterfaceSpec_type;
    'attach-tenant'?: string;
    'attach-network'?: string;
    'ip-alloc-type': NetworkNetworkInterfaceSpec_ip_alloc_type;
    'ip-config'?: IClusterIPConfig;
    'mac-address'?: string;
    '_ui'?: any;
}


export class NetworkNetworkInterfaceSpec extends BaseModel implements INetworkNetworkInterfaceSpec {
    /** Field for holding arbitrary ui state */
    '_ui': any = {};
    /** Desired Admin state of the port. */
    'admin-status': NetworkNetworkInterfaceSpec_admin_status = null;
    /** Intefaae speed. */
    'speed': string = null;
    /** Mtu of the interface. */
    'mtu': number = null;
    /** Pause Spec. */
    'pause': NetworkPauseSpec = null;
    /** Type specifies the type of interface. */
    'type': NetworkNetworkInterfaceSpec_type = null;
    'attach-tenant': string = null;
    /** AttachNetwork associates the interface with a Network. This is only valid for HOST_PF type. */
    'attach-network': string = null;
    'ip-alloc-type': NetworkNetworkInterfaceSpec_ip_alloc_type = null;
    /** Interface IP Configuration if any. */
    'ip-config': ClusterIPConfig = null;
    /** Override system allocated MAC address. Should be a valid MAC address. */
    'mac-address': string = null;
    public static propInfo: { [prop in keyof INetworkNetworkInterfaceSpec]: PropInfoItem } = {
        'admin-status': {
            enum: NetworkNetworkInterfaceSpec_admin_status,
            default: 'up',
            description:  `Desired Admin state of the port.`,
            required: true,
            type: 'string'
        },
        'speed': {
            description:  `Intefaae speed.`,
            required: false,
            type: 'string'
        },
        'mtu': {
            description:  `Mtu of the interface.`,
            required: false,
            type: 'number'
        },
        'pause': {
            description:  `Pause Spec.`,
            required: false,
            type: 'object'
        },
        'type': {
            enum: NetworkNetworkInterfaceSpec_type,
            default: 'none',
            description:  `Type specifies the type of interface.`,
            required: true,
            type: 'string'
        },
        'attach-tenant': {
            required: false,
            type: 'string'
        },
        'attach-network': {
            description:  `AttachNetwork associates the interface with a Network. This is only valid for HOST_PF type.`,
            required: false,
            type: 'string'
        },
        'ip-alloc-type': {
            enum: NetworkNetworkInterfaceSpec_ip_alloc_type,
            default: 'none',
            required: true,
            type: 'string'
        },
        'ip-config': {
            description:  `Interface IP Configuration if any.`,
            required: false,
            type: 'object'
        },
        'mac-address': {
            description:  `Override system allocated MAC address. Should be a valid MAC address.`,
            hint:  'aabb.ccdd.0000, aabb.ccdd.0000, aabb.ccdd.0000',
            required: false,
            type: 'string'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return NetworkNetworkInterfaceSpec.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return NetworkNetworkInterfaceSpec.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (NetworkNetworkInterfaceSpec.propInfo[prop] != null &&
                        NetworkNetworkInterfaceSpec.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['pause'] = new NetworkPauseSpec();
        this['ip-config'] = new ClusterIPConfig();
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
        if (values && values['admin-status'] != null) {
            this['admin-status'] = values['admin-status'];
        } else if (fillDefaults && NetworkNetworkInterfaceSpec.hasDefaultValue('admin-status')) {
            this['admin-status'] = <NetworkNetworkInterfaceSpec_admin_status>  NetworkNetworkInterfaceSpec.propInfo['admin-status'].default;
        } else {
            this['admin-status'] = null
        }
        if (values && values['speed'] != null) {
            this['speed'] = values['speed'];
        } else if (fillDefaults && NetworkNetworkInterfaceSpec.hasDefaultValue('speed')) {
            this['speed'] = NetworkNetworkInterfaceSpec.propInfo['speed'].default;
        } else {
            this['speed'] = null
        }
        if (values && values['mtu'] != null) {
            this['mtu'] = values['mtu'];
        } else if (fillDefaults && NetworkNetworkInterfaceSpec.hasDefaultValue('mtu')) {
            this['mtu'] = NetworkNetworkInterfaceSpec.propInfo['mtu'].default;
        } else {
            this['mtu'] = null
        }
        if (values) {
            this['pause'].setValues(values['pause'], fillDefaults);
        } else {
            this['pause'].setValues(null, fillDefaults);
        }
        if (values && values['type'] != null) {
            this['type'] = values['type'];
        } else if (fillDefaults && NetworkNetworkInterfaceSpec.hasDefaultValue('type')) {
            this['type'] = <NetworkNetworkInterfaceSpec_type>  NetworkNetworkInterfaceSpec.propInfo['type'].default;
        } else {
            this['type'] = null
        }
        if (values && values['attach-tenant'] != null) {
            this['attach-tenant'] = values['attach-tenant'];
        } else if (fillDefaults && NetworkNetworkInterfaceSpec.hasDefaultValue('attach-tenant')) {
            this['attach-tenant'] = NetworkNetworkInterfaceSpec.propInfo['attach-tenant'].default;
        } else {
            this['attach-tenant'] = null
        }
        if (values && values['attach-network'] != null) {
            this['attach-network'] = values['attach-network'];
        } else if (fillDefaults && NetworkNetworkInterfaceSpec.hasDefaultValue('attach-network')) {
            this['attach-network'] = NetworkNetworkInterfaceSpec.propInfo['attach-network'].default;
        } else {
            this['attach-network'] = null
        }
        if (values && values['ip-alloc-type'] != null) {
            this['ip-alloc-type'] = values['ip-alloc-type'];
        } else if (fillDefaults && NetworkNetworkInterfaceSpec.hasDefaultValue('ip-alloc-type')) {
            this['ip-alloc-type'] = <NetworkNetworkInterfaceSpec_ip_alloc_type>  NetworkNetworkInterfaceSpec.propInfo['ip-alloc-type'].default;
        } else {
            this['ip-alloc-type'] = null
        }
        if (values) {
            this['ip-config'].setValues(values['ip-config'], fillDefaults);
        } else {
            this['ip-config'].setValues(null, fillDefaults);
        }
        if (values && values['mac-address'] != null) {
            this['mac-address'] = values['mac-address'];
        } else if (fillDefaults && NetworkNetworkInterfaceSpec.hasDefaultValue('mac-address')) {
            this['mac-address'] = NetworkNetworkInterfaceSpec.propInfo['mac-address'].default;
        } else {
            this['mac-address'] = null
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'admin-status': CustomFormControl(new FormControl(this['admin-status'], [required, enumValidator(NetworkNetworkInterfaceSpec_admin_status), ]), NetworkNetworkInterfaceSpec.propInfo['admin-status']),
                'speed': CustomFormControl(new FormControl(this['speed']), NetworkNetworkInterfaceSpec.propInfo['speed']),
                'mtu': CustomFormControl(new FormControl(this['mtu']), NetworkNetworkInterfaceSpec.propInfo['mtu']),
                'pause': CustomFormGroup(this['pause'].$formGroup, NetworkNetworkInterfaceSpec.propInfo['pause'].required),
                'type': CustomFormControl(new FormControl(this['type'], [required, enumValidator(NetworkNetworkInterfaceSpec_type), ]), NetworkNetworkInterfaceSpec.propInfo['type']),
                'attach-tenant': CustomFormControl(new FormControl(this['attach-tenant']), NetworkNetworkInterfaceSpec.propInfo['attach-tenant']),
                'attach-network': CustomFormControl(new FormControl(this['attach-network']), NetworkNetworkInterfaceSpec.propInfo['attach-network']),
                'ip-alloc-type': CustomFormControl(new FormControl(this['ip-alloc-type'], [required, enumValidator(NetworkNetworkInterfaceSpec_ip_alloc_type), ]), NetworkNetworkInterfaceSpec.propInfo['ip-alloc-type']),
                'ip-config': CustomFormGroup(this['ip-config'].$formGroup, NetworkNetworkInterfaceSpec.propInfo['ip-config'].required),
                'mac-address': CustomFormControl(new FormControl(this['mac-address']), NetworkNetworkInterfaceSpec.propInfo['mac-address']),
            });
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('pause') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('pause').get(field);
                control.updateValueAndValidity();
            });
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('ip-config') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('ip-config').get(field);
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
            this._formGroup.controls['admin-status'].setValue(this['admin-status']);
            this._formGroup.controls['speed'].setValue(this['speed']);
            this._formGroup.controls['mtu'].setValue(this['mtu']);
            this['pause'].setFormGroupValuesToBeModelValues();
            this._formGroup.controls['type'].setValue(this['type']);
            this._formGroup.controls['attach-tenant'].setValue(this['attach-tenant']);
            this._formGroup.controls['attach-network'].setValue(this['attach-network']);
            this._formGroup.controls['ip-alloc-type'].setValue(this['ip-alloc-type']);
            this['ip-config'].setFormGroupValuesToBeModelValues();
            this._formGroup.controls['mac-address'].setValue(this['mac-address']);
        }
    }
}

