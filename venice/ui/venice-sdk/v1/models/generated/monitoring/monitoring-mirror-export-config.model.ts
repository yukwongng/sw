/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';


export interface IMonitoringMirrorExportConfig {
    'destination': string;
}


export class MonitoringMirrorExportConfig extends BaseModel implements IMonitoringMirrorExportConfig {
    /** length of string should be between 1 and 2048 */
    'destination': string = null;
    public static propInfo: { [prop in keyof IMonitoringMirrorExportConfig]: PropInfoItem } = {
        'destination': {
            description:  'length of string should be between 1 and 2048',
            required: true,
            type: 'string'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return MonitoringMirrorExportConfig.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return MonitoringMirrorExportConfig.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (MonitoringMirrorExportConfig.propInfo[prop] != null &&
                        MonitoringMirrorExportConfig.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this._inputValue = values;
        this.setValues(values, setDefaults);
    }

    /**
     * set the values for both the Model and the Form Group. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any, fillDefaults = true): void {
        if (values && values['destination'] != null) {
            this['destination'] = values['destination'];
        } else if (fillDefaults && MonitoringMirrorExportConfig.hasDefaultValue('destination')) {
            this['destination'] = MonitoringMirrorExportConfig.propInfo['destination'].default;
        } else {
            this['destination'] = null
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'destination': CustomFormControl(new FormControl(this['destination'], [required, minLengthValidator(1), maxLengthValidator(2048), ]), MonitoringMirrorExportConfig.propInfo['destination']),
            });
        }
        return this._formGroup;
    }

    setModelToBeFormGroupValues() {
        this.setValues(this.$formGroup.value, false);
    }

    setFormGroupValuesToBeModelValues() {
        if (this._formGroup) {
            this._formGroup.controls['destination'].setValue(this['destination']);
        }
    }
}

