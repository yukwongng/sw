/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';


export interface ISecuritySecurityGroupStatus {
    'workloads'?: Array<string>;
    'Policies'?: Array<string>;
    '_ui'?: any;
}


export class SecuritySecurityGroupStatus extends BaseModel implements ISecuritySecurityGroupStatus {
    /** Field for holding arbitrary ui state */
    '_ui': any = {};
    /** List of workloads that are part of this security group. */
    'workloads': Array<string> = null;
    /** List of all policies attached to this security group. */
    'Policies': Array<string> = null;
    public static propInfo: { [prop in keyof ISecuritySecurityGroupStatus]: PropInfoItem } = {
        'workloads': {
            description:  `List of workloads that are part of this security group.`,
            required: false,
            type: 'Array<string>'
        },
        'Policies': {
            description:  `List of all policies attached to this security group.`,
            required: false,
            type: 'Array<string>'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return SecuritySecurityGroupStatus.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return SecuritySecurityGroupStatus.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (SecuritySecurityGroupStatus.propInfo[prop] != null &&
                        SecuritySecurityGroupStatus.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['workloads'] = new Array<string>();
        this['Policies'] = new Array<string>();
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
        if (values && values['workloads'] != null) {
            this['workloads'] = values['workloads'];
        } else if (fillDefaults && SecuritySecurityGroupStatus.hasDefaultValue('workloads')) {
            this['workloads'] = [ SecuritySecurityGroupStatus.propInfo['workloads'].default];
        } else {
            this['workloads'] = [];
        }
        if (values && values['Policies'] != null) {
            this['Policies'] = values['Policies'];
        } else if (fillDefaults && SecuritySecurityGroupStatus.hasDefaultValue('Policies')) {
            this['Policies'] = [ SecuritySecurityGroupStatus.propInfo['Policies'].default];
        } else {
            this['Policies'] = [];
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'workloads': CustomFormControl(new FormControl(this['workloads']), SecuritySecurityGroupStatus.propInfo['workloads']),
                'Policies': CustomFormControl(new FormControl(this['Policies']), SecuritySecurityGroupStatus.propInfo['Policies']),
            });
        }
        return this._formGroup;
    }

    setModelToBeFormGroupValues() {
        this.setValues(this.$formGroup.value, false);
    }

    setFormGroupValuesToBeModelValues() {
        if (this._formGroup) {
            this._formGroup.controls['workloads'].setValue(this['workloads']);
            this._formGroup.controls['Policies'].setValue(this['Policies']);
        }
    }
}

