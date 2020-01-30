/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';


export interface IClusterClusterSpec {
    'quorum-nodes'?: Array<string>;
    'virtual-ip'?: string;
    'ntp-servers'?: Array<string>;
    'auto-admit-dscs'?: boolean;
    'certs'?: string;
    'key'?: string;
}


export class ClusterClusterSpec extends BaseModel implements IClusterClusterSpec {
    /** QuorumNodes contains the list of hostnames for nodes configured to be quorum nodes in the cluster. */
    'quorum-nodes': Array<string> = null;
    /** VirtualIP is the IP address for managing the cluster. It will be hosted by the winner of election between quorum nodes. */
    'virtual-ip': string = null;
    /** NTPServers contains the list of NTP servers for the cluster. */
    'ntp-servers': Array<string> = null;
    /** AutoAdmitDSCs when enabled auto-admits DSCs that are validated into Venice Cluster. When it is disabled, DSCs validated by CMD are set to Pending state and it requires Manual approval to be admitted into the cluster. */
    'auto-admit-dscs': boolean = null;
    /** Certs is the pem encoded certificate bundle used for API Gateway TLS. */
    'certs': string = null;
    /** Key is the pem encoded private key used for API Gateway TLS. We support RSA or ECDSA. */
    'key': string = null;
    public static propInfo: { [prop in keyof IClusterClusterSpec]: PropInfoItem } = {
        'quorum-nodes': {
            description:  `QuorumNodes contains the list of hostnames for nodes configured to be quorum nodes in the cluster.`,
            required: false,
            type: 'Array<string>'
        },
        'virtual-ip': {
            description:  `VirtualIP is the IP address for managing the cluster. It will be hosted by the winner of election between quorum nodes.`,
            required: false,
            type: 'string'
        },
        'ntp-servers': {
            description:  `NTPServers contains the list of NTP servers for the cluster.`,
            required: false,
            type: 'Array<string>'
        },
        'auto-admit-dscs': {
            description:  `AutoAdmitDSCs when enabled auto-admits DSCs that are validated into Venice Cluster. When it is disabled, DSCs validated by CMD are set to Pending state and it requires Manual approval to be admitted into the cluster.`,
            required: false,
            type: 'boolean'
        },
        'certs': {
            description:  `Certs is the pem encoded certificate bundle used for API Gateway TLS.`,
            required: false,
            type: 'string'
        },
        'key': {
            description:  `Key is the pem encoded private key used for API Gateway TLS. We support RSA or ECDSA.`,
            required: false,
            type: 'string'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return ClusterClusterSpec.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return ClusterClusterSpec.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (ClusterClusterSpec.propInfo[prop] != null &&
                        ClusterClusterSpec.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['quorum-nodes'] = new Array<string>();
        this['ntp-servers'] = new Array<string>();
        this._inputValue = values;
        this.setValues(values, setDefaults);
    }

    /**
     * set the values for both the Model and the Form Group. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any, fillDefaults = true): void {
        if (values && values['quorum-nodes'] != null) {
            this['quorum-nodes'] = values['quorum-nodes'];
        } else if (fillDefaults && ClusterClusterSpec.hasDefaultValue('quorum-nodes')) {
            this['quorum-nodes'] = [ ClusterClusterSpec.propInfo['quorum-nodes'].default];
        } else {
            this['quorum-nodes'] = [];
        }
        if (values && values['virtual-ip'] != null) {
            this['virtual-ip'] = values['virtual-ip'];
        } else if (fillDefaults && ClusterClusterSpec.hasDefaultValue('virtual-ip')) {
            this['virtual-ip'] = ClusterClusterSpec.propInfo['virtual-ip'].default;
        } else {
            this['virtual-ip'] = null
        }
        if (values && values['ntp-servers'] != null) {
            this['ntp-servers'] = values['ntp-servers'];
        } else if (fillDefaults && ClusterClusterSpec.hasDefaultValue('ntp-servers')) {
            this['ntp-servers'] = [ ClusterClusterSpec.propInfo['ntp-servers'].default];
        } else {
            this['ntp-servers'] = [];
        }
        if (values && values['auto-admit-dscs'] != null) {
            this['auto-admit-dscs'] = values['auto-admit-dscs'];
        } else if (fillDefaults && ClusterClusterSpec.hasDefaultValue('auto-admit-dscs')) {
            this['auto-admit-dscs'] = ClusterClusterSpec.propInfo['auto-admit-dscs'].default;
        } else {
            this['auto-admit-dscs'] = null
        }
        if (values && values['certs'] != null) {
            this['certs'] = values['certs'];
        } else if (fillDefaults && ClusterClusterSpec.hasDefaultValue('certs')) {
            this['certs'] = ClusterClusterSpec.propInfo['certs'].default;
        } else {
            this['certs'] = null
        }
        if (values && values['key'] != null) {
            this['key'] = values['key'];
        } else if (fillDefaults && ClusterClusterSpec.hasDefaultValue('key')) {
            this['key'] = ClusterClusterSpec.propInfo['key'].default;
        } else {
            this['key'] = null
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'quorum-nodes': CustomFormControl(new FormControl(this['quorum-nodes']), ClusterClusterSpec.propInfo['quorum-nodes']),
                'virtual-ip': CustomFormControl(new FormControl(this['virtual-ip']), ClusterClusterSpec.propInfo['virtual-ip']),
                'ntp-servers': CustomFormControl(new FormControl(this['ntp-servers']), ClusterClusterSpec.propInfo['ntp-servers']),
                'auto-admit-dscs': CustomFormControl(new FormControl(this['auto-admit-dscs']), ClusterClusterSpec.propInfo['auto-admit-dscs']),
                'certs': CustomFormControl(new FormControl(this['certs']), ClusterClusterSpec.propInfo['certs']),
                'key': CustomFormControl(new FormControl(this['key']), ClusterClusterSpec.propInfo['key']),
            });
        }
        return this._formGroup;
    }

    setModelToBeFormGroupValues() {
        this.setValues(this.$formGroup.value, false);
    }

    setFormGroupValuesToBeModelValues() {
        if (this._formGroup) {
            this._formGroup.controls['quorum-nodes'].setValue(this['quorum-nodes']);
            this._formGroup.controls['virtual-ip'].setValue(this['virtual-ip']);
            this._formGroup.controls['ntp-servers'].setValue(this['ntp-servers']);
            this._formGroup.controls['auto-admit-dscs'].setValue(this['auto-admit-dscs']);
            this._formGroup.controls['certs'].setValue(this['certs']);
            this._formGroup.controls['key'].setValue(this['key']);
        }
    }
}

