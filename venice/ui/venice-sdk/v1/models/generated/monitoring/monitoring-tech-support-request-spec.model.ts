/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { LabelsSelector, ILabelsSelector } from './labels-selector.model';
import { TechSupportRequestSpecNodeSelectorSpec, ITechSupportRequestSpecNodeSelectorSpec } from './tech-support-request-spec-node-selector-spec.model';

export interface IMonitoringTechSupportRequestSpec {
    'collection-selector'?: ILabelsSelector;
    'node-selector'?: ITechSupportRequestSpecNodeSelectorSpec;
    'verbosity'?: number;
    '_ui'?: any;
}


export class MonitoringTechSupportRequestSpec extends BaseModel implements IMonitoringTechSupportRequestSpec {
    /** Field for holding arbitrary ui state */
    '_ui': any = {};
    /** CollectionSelector is a Label selector for modules to collect. */
    'collection-selector': LabelsSelector = null;
    /** NodeSelector is a label selector that selects nodes to collect tech support from. */
    'node-selector': TechSupportRequestSpecNodeSelectorSpec = null;
    /** Verbosity defines the verbosity level. */
    'verbosity': number = null;
    public static propInfo: { [prop in keyof IMonitoringTechSupportRequestSpec]: PropInfoItem } = {
        'collection-selector': {
            description:  `CollectionSelector is a Label selector for modules to collect.`,
            required: false,
            type: 'object'
        },
        'node-selector': {
            description:  `NodeSelector is a label selector that selects nodes to collect tech support from.`,
            required: false,
            type: 'object'
        },
        'verbosity': {
            description:  `Verbosity defines the verbosity level.`,
            required: false,
            type: 'number'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return MonitoringTechSupportRequestSpec.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return MonitoringTechSupportRequestSpec.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (MonitoringTechSupportRequestSpec.propInfo[prop] != null &&
                        MonitoringTechSupportRequestSpec.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['collection-selector'] = new LabelsSelector();
        this['node-selector'] = new TechSupportRequestSpecNodeSelectorSpec();
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
        if (values) {
            this['collection-selector'].setValues(values['collection-selector'], fillDefaults);
        } else {
            this['collection-selector'].setValues(null, fillDefaults);
        }
        if (values) {
            this['node-selector'].setValues(values['node-selector'], fillDefaults);
        } else {
            this['node-selector'].setValues(null, fillDefaults);
        }
        if (values && values['verbosity'] != null) {
            this['verbosity'] = values['verbosity'];
        } else if (fillDefaults && MonitoringTechSupportRequestSpec.hasDefaultValue('verbosity')) {
            this['verbosity'] = MonitoringTechSupportRequestSpec.propInfo['verbosity'].default;
        } else {
            this['verbosity'] = null
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'collection-selector': CustomFormGroup(this['collection-selector'].$formGroup, MonitoringTechSupportRequestSpec.propInfo['collection-selector'].required),
                'node-selector': CustomFormGroup(this['node-selector'].$formGroup, MonitoringTechSupportRequestSpec.propInfo['node-selector'].required),
                'verbosity': CustomFormControl(new FormControl(this['verbosity']), MonitoringTechSupportRequestSpec.propInfo['verbosity']),
            });
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('collection-selector') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('collection-selector').get(field);
                control.updateValueAndValidity();
            });
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('node-selector') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('node-selector').get(field);
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
            this['collection-selector'].setFormGroupValuesToBeModelValues();
            this['node-selector'].setFormGroupValuesToBeModelValues();
            this._formGroup.controls['verbosity'].setValue(this['verbosity']);
        }
    }
}

