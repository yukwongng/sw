/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';


export interface IMonitoringMatchSelector {
    'ip-addresses'?: Array<string>;
    'mac-addresses'?: Array<string>;
}


export class MonitoringMatchSelector extends BaseModel implements IMonitoringMatchSelector {
    'ip-addresses': Array<string> = null;
    /** should be a valid MAC address */
    'mac-addresses': Array<string> = null;
    public static propInfo: { [prop in keyof IMonitoringMatchSelector]: PropInfoItem } = {
        'ip-addresses': {
            required: false,
            type: 'Array<string>'
        },
        'mac-addresses': {
            description:  'should be a valid MAC address',
            hint:  'aabb.ccdd.0000, aabb.ccdd.0000, aabb.ccdd.0000',
            required: false,
            type: 'Array<string>'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return MonitoringMatchSelector.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return MonitoringMatchSelector.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (MonitoringMatchSelector.propInfo[prop] != null &&
                        MonitoringMatchSelector.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['ip-addresses'] = new Array<string>();
        this['mac-addresses'] = new Array<string>();
        this._inputValue = values;
        this.setValues(values, setDefaults);
    }

    /**
     * set the values for both the Model and the Form Group. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any, fillDefaults = true): void {
        if (values && values['ip-addresses'] != null) {
            this['ip-addresses'] = values['ip-addresses'];
        } else if (fillDefaults && MonitoringMatchSelector.hasDefaultValue('ip-addresses')) {
            this['ip-addresses'] = [ MonitoringMatchSelector.propInfo['ip-addresses'].default];
        } else {
            this['ip-addresses'] = [];
        }
        if (values && values['mac-addresses'] != null) {
            this['mac-addresses'] = values['mac-addresses'];
        } else if (fillDefaults && MonitoringMatchSelector.hasDefaultValue('mac-addresses')) {
            this['mac-addresses'] = [ MonitoringMatchSelector.propInfo['mac-addresses'].default];
        } else {
            this['mac-addresses'] = [];
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'ip-addresses': CustomFormControl(new FormControl(this['ip-addresses']), MonitoringMatchSelector.propInfo['ip-addresses']),
                'mac-addresses': CustomFormControl(new FormControl(this['mac-addresses']), MonitoringMatchSelector.propInfo['mac-addresses']),
            });
        }
        return this._formGroup;
    }

    setModelToBeFormGroupValues() {
        this.setValues(this.$formGroup.value, false);
    }

    setFormGroupValuesToBeModelValues() {
        if (this._formGroup) {
            this._formGroup.controls['ip-addresses'].setValue(this['ip-addresses']);
            this._formGroup.controls['mac-addresses'].setValue(this['mac-addresses']);
        }
    }
}

