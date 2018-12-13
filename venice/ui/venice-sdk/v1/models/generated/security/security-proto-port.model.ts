/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, enumValidator } from './validators';
import { BaseModel, PropInfoItem } from './base-model';


export interface ISecurityProtoPort {
    'protocol'?: string;
    'ports'?: string;
}


export class SecurityProtoPort extends BaseModel implements ISecurityProtoPort {
    /** protocol is ip (v4/v6) protocol name/number; names can be: tcp, udp, igmp, icmp, gre, esp, etc. */
    'protocol': string = null;
    'ports': string = null;
    public static propInfo: { [prop: string]: PropInfoItem } = {
        'protocol': {
            description:  'protocol is ip (v4/v6) protocol name/number; names can be: tcp, udp, igmp, icmp, gre, esp, etc.',
            type: 'string'
        },
        'ports': {
            type: 'string'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return SecurityProtoPort.propInfo[propName];
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (SecurityProtoPort.propInfo[prop] != null &&
                        SecurityProtoPort.propInfo[prop].default != null &&
                        SecurityProtoPort.propInfo[prop].default != '');
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any) {
        super();
        this.setValues(values);
    }

    /**
     * set the values for both the Model and the Form Group. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any, fillDefaults = true): void {
        if (values && values['protocol'] != null) {
            this['protocol'] = values['protocol'];
        } else if (fillDefaults && SecurityProtoPort.hasDefaultValue('protocol')) {
            this['protocol'] = SecurityProtoPort.propInfo['protocol'].default;
        }
        if (values && values['ports'] != null) {
            this['ports'] = values['ports'];
        } else if (fillDefaults && SecurityProtoPort.hasDefaultValue('ports')) {
            this['ports'] = SecurityProtoPort.propInfo['ports'].default;
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'protocol': new FormControl(this['protocol']),
                'ports': new FormControl(this['ports']),
            });
        }
        return this._formGroup;
    }

    setModelToBeFormGroupValues() {
        this.setValues(this.$formGroup.value, false);
    }

    setFormGroupValuesToBeModelValues() {
        if (this._formGroup) {
            this._formGroup.controls['protocol'].setValue(this['protocol']);
            this._formGroup.controls['ports'].setValue(this['ports']);
        }
    }
}

