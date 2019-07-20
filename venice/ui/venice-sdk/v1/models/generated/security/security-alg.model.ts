/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from './base-model';

import { SecurityALG_type,  } from './enums';
import { SecurityIcmp, ISecurityIcmp } from './security-icmp.model';
import { SecurityDns, ISecurityDns } from './security-dns.model';
import { SecurityFtp, ISecurityFtp } from './security-ftp.model';
import { SecuritySunrpc, ISecuritySunrpc } from './security-sunrpc.model';
import { SecurityMsrpc, ISecurityMsrpc } from './security-msrpc.model';

export interface ISecurityALG {
    'type': SecurityALG_type;
    'icmp'?: ISecurityIcmp;
    'dns'?: ISecurityDns;
    'ftp'?: ISecurityFtp;
    'sunrpc'?: Array<ISecuritySunrpc>;
    'msrpc'?: Array<ISecurityMsrpc>;
}


export class SecurityALG extends BaseModel implements ISecurityALG {
    'type': SecurityALG_type = null;
    'icmp': SecurityIcmp = null;
    'dns': SecurityDns = null;
    'ftp': SecurityFtp = null;
    'sunrpc': Array<SecuritySunrpc> = null;
    'msrpc': Array<SecurityMsrpc> = null;
    public static propInfo: { [prop: string]: PropInfoItem } = {
        'type': {
            enum: SecurityALG_type,
            default: 'icmp',
            required: true,
            type: 'string'
        },
        'icmp': {
            required: false,
            type: 'object'
        },
        'dns': {
            required: false,
            type: 'object'
        },
        'ftp': {
            required: false,
            type: 'object'
        },
        'sunrpc': {
            required: false,
            type: 'object'
        },
        'msrpc': {
            required: false,
            type: 'object'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return SecurityALG.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return SecurityALG.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (SecurityALG.propInfo[prop] != null &&
                        SecurityALG.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['icmp'] = new SecurityIcmp();
        this['dns'] = new SecurityDns();
        this['ftp'] = new SecurityFtp();
        this['sunrpc'] = new Array<SecuritySunrpc>();
        this['msrpc'] = new Array<SecurityMsrpc>();
        this.setValues(values, setDefaults);
    }

    /**
     * set the values for both the Model and the Form Group. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any, fillDefaults = true): void {
        if (values && values['type'] != null) {
            this['type'] = values['type'];
        } else if (fillDefaults && SecurityALG.hasDefaultValue('type')) {
            this['type'] = <SecurityALG_type>  SecurityALG.propInfo['type'].default;
        } else {
            this['type'] = null
        }
        if (values) {
            this['icmp'].setValues(values['icmp'], fillDefaults);
        } else {
            this['icmp'].setValues(null, fillDefaults);
        }
        if (values) {
            this['dns'].setValues(values['dns'], fillDefaults);
        } else {
            this['dns'].setValues(null, fillDefaults);
        }
        if (values) {
            this['ftp'].setValues(values['ftp'], fillDefaults);
        } else {
            this['ftp'].setValues(null, fillDefaults);
        }
        if (values) {
            this.fillModelArray<SecuritySunrpc>(this, 'sunrpc', values['sunrpc'], SecuritySunrpc);
        } else {
            this['sunrpc'] = [];
        }
        if (values) {
            this.fillModelArray<SecurityMsrpc>(this, 'msrpc', values['msrpc'], SecurityMsrpc);
        } else {
            this['msrpc'] = [];
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'type': CustomFormControl(new FormControl(this['type'], [required, enumValidator(SecurityALG_type), ]), SecurityALG.propInfo['type']),
                'icmp': CustomFormGroup(this['icmp'].$formGroup, SecurityALG.propInfo['icmp'].required),
                'dns': CustomFormGroup(this['dns'].$formGroup, SecurityALG.propInfo['dns'].required),
                'ftp': CustomFormGroup(this['ftp'].$formGroup, SecurityALG.propInfo['ftp'].required),
                'sunrpc': new FormArray([]),
                'msrpc': new FormArray([]),
            });
            // generate FormArray control elements
            this.fillFormArray<SecuritySunrpc>('sunrpc', this['sunrpc'], SecuritySunrpc);
            // generate FormArray control elements
            this.fillFormArray<SecurityMsrpc>('msrpc', this['msrpc'], SecurityMsrpc);
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('icmp') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('icmp').get(field);
                control.updateValueAndValidity();
            });
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('dns') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('dns').get(field);
                control.updateValueAndValidity();
            });
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('ftp') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('ftp').get(field);
                control.updateValueAndValidity();
            });
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('sunrpc') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('sunrpc').get(field);
                control.updateValueAndValidity();
            });
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('msrpc') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('msrpc').get(field);
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
            this['icmp'].setFormGroupValuesToBeModelValues();
            this['dns'].setFormGroupValuesToBeModelValues();
            this['ftp'].setFormGroupValuesToBeModelValues();
            this.fillModelArray<SecuritySunrpc>(this, 'sunrpc', this['sunrpc'], SecuritySunrpc);
            this.fillModelArray<SecurityMsrpc>(this, 'msrpc', this['msrpc'], SecurityMsrpc);
        }
    }
}

