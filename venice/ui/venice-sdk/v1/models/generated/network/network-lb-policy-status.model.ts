/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';


export interface INetworkLbPolicyStatus {
    'type'?: Array<string>;
    '_ui'?: any;
}


export class NetworkLbPolicyStatus extends BaseModel implements INetworkLbPolicyStatus {
    /** Field for holding arbitrary ui state */
    '_ui': any = {};
    /** List of service objects referring this lb-policy. */
    'type': Array<string> = null;
    public static propInfo: { [prop in keyof INetworkLbPolicyStatus]: PropInfoItem } = {
        'type': {
            description:  `List of service objects referring this lb-policy.`,
            required: false,
            type: 'Array<string>'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return NetworkLbPolicyStatus.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return NetworkLbPolicyStatus.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (NetworkLbPolicyStatus.propInfo[prop] != null &&
                        NetworkLbPolicyStatus.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['type'] = new Array<string>();
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
        } else if (fillDefaults && NetworkLbPolicyStatus.hasDefaultValue('type')) {
            this['type'] = [ NetworkLbPolicyStatus.propInfo['type'].default];
        } else {
            this['type'] = [];
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'type': CustomFormControl(new FormControl(this['type']), NetworkLbPolicyStatus.propInfo['type']),
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
        }
    }
}

