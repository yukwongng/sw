/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';


export interface INetworkTLSClientPolicySpec {
    'tls-client-certificates-selector'?: object;
    'tls-client-trust-roots'?: Array<string>;
    'tls-client-allowed-peer-id'?: Array<string>;
    '_ui'?: any;
}


export class NetworkTLSClientPolicySpec extends BaseModel implements INetworkTLSClientPolicySpec {
    /** Field for holding arbitrary ui state */
    '_ui': any = {};
    /** A map containing the certificate to use for a set of destinations. The key is a selector for workloads that exist either inside or outside the cluster. It can be based on labels, hostnames or "IP:port" pairs. The value is the name of the certificate to use for the selected destinations. The certificates "usage" field must contain "client". TODO: replace the first "string" type with proper selector type when available. A single "default" certificate which matches all destinations is allowed. If a destination matches multiple non-default map keys, an error is returned. If a destination does not match any map key (and there is no default), the outbound connection is initiated without TLS. */
    'tls-client-certificates-selector': object = null;
    /** The list of root certificates used to validate a trust chain presented by a server. If the list is empty, all roots certificates in the tenant scope are considered. */
    'tls-client-trust-roots': Array<string> = null;
    /** Valid DNS names or IP addresses that must appear in the server certificate SubjAltName or Common Name (if SAN is not specified). If not specified, client validates the IP address of the server. */
    'tls-client-allowed-peer-id': Array<string> = null;
    public static propInfo: { [prop in keyof INetworkTLSClientPolicySpec]: PropInfoItem } = {
        'tls-client-certificates-selector': {
            description:  `A map containing the certificate to use for a set of destinations. The key is a selector for workloads that exist either inside or outside the cluster. It can be based on labels, hostnames or "IP:port" pairs. The value is the name of the certificate to use for the selected destinations. The certificates "usage" field must contain "client". TODO: replace the first "string" type with proper selector type when available. A single "default" certificate which matches all destinations is allowed. If a destination matches multiple non-default map keys, an error is returned. If a destination does not match any map key (and there is no default), the outbound connection is initiated without TLS.`,
            required: false,
            type: 'object'
        },
        'tls-client-trust-roots': {
            description:  `The list of root certificates used to validate a trust chain presented by a server. If the list is empty, all roots certificates in the tenant scope are considered.`,
            required: false,
            type: 'Array<string>'
        },
        'tls-client-allowed-peer-id': {
            description:  `Valid DNS names or IP addresses that must appear in the server certificate SubjAltName or Common Name (if SAN is not specified). If not specified, client validates the IP address of the server.`,
            required: false,
            type: 'Array<string>'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return NetworkTLSClientPolicySpec.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return NetworkTLSClientPolicySpec.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (NetworkTLSClientPolicySpec.propInfo[prop] != null &&
                        NetworkTLSClientPolicySpec.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['tls-client-trust-roots'] = new Array<string>();
        this['tls-client-allowed-peer-id'] = new Array<string>();
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
        if (values && values['tls-client-certificates-selector'] != null) {
            this['tls-client-certificates-selector'] = values['tls-client-certificates-selector'];
        } else if (fillDefaults && NetworkTLSClientPolicySpec.hasDefaultValue('tls-client-certificates-selector')) {
            this['tls-client-certificates-selector'] = NetworkTLSClientPolicySpec.propInfo['tls-client-certificates-selector'].default;
        } else {
            this['tls-client-certificates-selector'] = null
        }
        if (values && values['tls-client-trust-roots'] != null) {
            this['tls-client-trust-roots'] = values['tls-client-trust-roots'];
        } else if (fillDefaults && NetworkTLSClientPolicySpec.hasDefaultValue('tls-client-trust-roots')) {
            this['tls-client-trust-roots'] = [ NetworkTLSClientPolicySpec.propInfo['tls-client-trust-roots'].default];
        } else {
            this['tls-client-trust-roots'] = [];
        }
        if (values && values['tls-client-allowed-peer-id'] != null) {
            this['tls-client-allowed-peer-id'] = values['tls-client-allowed-peer-id'];
        } else if (fillDefaults && NetworkTLSClientPolicySpec.hasDefaultValue('tls-client-allowed-peer-id')) {
            this['tls-client-allowed-peer-id'] = [ NetworkTLSClientPolicySpec.propInfo['tls-client-allowed-peer-id'].default];
        } else {
            this['tls-client-allowed-peer-id'] = [];
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'tls-client-certificates-selector': CustomFormControl(new FormControl(this['tls-client-certificates-selector']), NetworkTLSClientPolicySpec.propInfo['tls-client-certificates-selector']),
                'tls-client-trust-roots': CustomFormControl(new FormControl(this['tls-client-trust-roots']), NetworkTLSClientPolicySpec.propInfo['tls-client-trust-roots']),
                'tls-client-allowed-peer-id': CustomFormControl(new FormControl(this['tls-client-allowed-peer-id']), NetworkTLSClientPolicySpec.propInfo['tls-client-allowed-peer-id']),
            });
        }
        return this._formGroup;
    }

    setModelToBeFormGroupValues() {
        this.setValues(this.$formGroup.value, false);
    }

    setFormGroupValuesToBeModelValues() {
        if (this._formGroup) {
            this._formGroup.controls['tls-client-certificates-selector'].setValue(this['tls-client-certificates-selector']);
            this._formGroup.controls['tls-client-trust-roots'].setValue(this['tls-client-trust-roots']);
            this._formGroup.controls['tls-client-allowed-peer-id'].setValue(this['tls-client-allowed-peer-id']);
        }
    }
}

