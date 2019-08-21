/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { MonitoringMatchSelector, IMonitoringMatchSelector } from './monitoring-match-selector.model';
import { MonitoringAppProtoSelector, IMonitoringAppProtoSelector } from './monitoring-app-proto-selector.model';

export interface IMonitoringMatchRule {
    'source'?: IMonitoringMatchSelector;
    'destination'?: IMonitoringMatchSelector;
    'app-protocol-selectors'?: IMonitoringAppProtoSelector;
}


export class MonitoringMatchRule extends BaseModel implements IMonitoringMatchRule {
    'source': MonitoringMatchSelector = null;
    'destination': MonitoringMatchSelector = null;
    'app-protocol-selectors': MonitoringAppProtoSelector = null;
    public static propInfo: { [prop in keyof IMonitoringMatchRule]: PropInfoItem } = {
        'source': {
            required: false,
            type: 'object'
        },
        'destination': {
            required: false,
            type: 'object'
        },
        'app-protocol-selectors': {
            required: false,
            type: 'object'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return MonitoringMatchRule.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return MonitoringMatchRule.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (MonitoringMatchRule.propInfo[prop] != null &&
                        MonitoringMatchRule.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['source'] = new MonitoringMatchSelector();
        this['destination'] = new MonitoringMatchSelector();
        this['app-protocol-selectors'] = new MonitoringAppProtoSelector();
        this._inputValue = values;
        this.setValues(values, setDefaults);
    }

    /**
     * set the values for both the Model and the Form Group. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any, fillDefaults = true): void {
        if (values) {
            this['source'].setValues(values['source'], fillDefaults);
        } else {
            this['source'].setValues(null, fillDefaults);
        }
        if (values) {
            this['destination'].setValues(values['destination'], fillDefaults);
        } else {
            this['destination'].setValues(null, fillDefaults);
        }
        if (values) {
            this['app-protocol-selectors'].setValues(values['app-protocol-selectors'], fillDefaults);
        } else {
            this['app-protocol-selectors'].setValues(null, fillDefaults);
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'source': CustomFormGroup(this['source'].$formGroup, MonitoringMatchRule.propInfo['source'].required),
                'destination': CustomFormGroup(this['destination'].$formGroup, MonitoringMatchRule.propInfo['destination'].required),
                'app-protocol-selectors': CustomFormGroup(this['app-protocol-selectors'].$formGroup, MonitoringMatchRule.propInfo['app-protocol-selectors'].required),
            });
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('source') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('source').get(field);
                control.updateValueAndValidity();
            });
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('destination') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('destination').get(field);
                control.updateValueAndValidity();
            });
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('app-protocol-selectors') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('app-protocol-selectors').get(field);
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
            this['source'].setFormGroupValuesToBeModelValues();
            this['destination'].setFormGroupValuesToBeModelValues();
            this['app-protocol-selectors'].setFormGroupValuesToBeModelValues();
        }
    }
}

