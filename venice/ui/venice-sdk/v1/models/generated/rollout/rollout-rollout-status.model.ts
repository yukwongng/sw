/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { RolloutRolloutPhase, IRolloutRolloutPhase } from './rollout-rollout-phase.model';
import { RolloutRolloutStatus_state,  } from './enums';

export interface IRolloutRolloutStatus {
    'controller-nodes-status'?: Array<IRolloutRolloutPhase>;
    'controller-services-status'?: Array<IRolloutRolloutPhase>;
    'smartnics-status'?: Array<IRolloutRolloutPhase>;
    'state': RolloutRolloutStatus_state;
    'completion-percent'?: number;
    'start-time'?: Date;
    'end-time'?: Date;
    'prev-version'?: string;
}


export class RolloutRolloutStatus extends BaseModel implements IRolloutRolloutStatus {
    'controller-nodes-status': Array<RolloutRolloutPhase> = null;
    'controller-services-status': Array<RolloutRolloutPhase> = null;
    /** Rollout status of SmartNICs in the cluster. Has entries for SmartNICs on Controller nodes as well as workload nodes
    The entries are group by parallelism based on the order-constraints and max-parallel specified by the user. */
    'smartnics-status': Array<RolloutRolloutPhase> = null;
    'state': RolloutRolloutStatus_state = null;
    'completion-percent': number = null;
    'start-time': Date = null;
    'end-time': Date = null;
    'prev-version': string = null;
    public static propInfo: { [prop in keyof IRolloutRolloutStatus]: PropInfoItem } = {
        'controller-nodes-status': {
            required: false,
            type: 'object'
        },
        'controller-services-status': {
            required: false,
            type: 'object'
        },
        'smartnics-status': {
            description:  'Rollout status of SmartNICs in the cluster. Has entries for SmartNICs on Controller nodes as well as workload nodes The entries are group by parallelism based on the order-constraints and max-parallel specified by the user.',
            required: false,
            type: 'object'
        },
        'state': {
            enum: RolloutRolloutStatus_state,
            default: 'progressing',
            required: true,
            type: 'string'
        },
        'completion-percent': {
            required: false,
            type: 'number'
        },
        'start-time': {
            required: false,
            type: 'Date'
        },
        'end-time': {
            required: false,
            type: 'Date'
        },
        'prev-version': {
            required: false,
            type: 'string'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return RolloutRolloutStatus.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return RolloutRolloutStatus.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (RolloutRolloutStatus.propInfo[prop] != null &&
                        RolloutRolloutStatus.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['controller-nodes-status'] = new Array<RolloutRolloutPhase>();
        this['controller-services-status'] = new Array<RolloutRolloutPhase>();
        this['smartnics-status'] = new Array<RolloutRolloutPhase>();
        this._inputValue = values;
        this.setValues(values, setDefaults);
    }

    /**
     * set the values for both the Model and the Form Group. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any, fillDefaults = true): void {
        if (values) {
            this.fillModelArray<RolloutRolloutPhase>(this, 'controller-nodes-status', values['controller-nodes-status'], RolloutRolloutPhase);
        } else {
            this['controller-nodes-status'] = [];
        }
        if (values) {
            this.fillModelArray<RolloutRolloutPhase>(this, 'controller-services-status', values['controller-services-status'], RolloutRolloutPhase);
        } else {
            this['controller-services-status'] = [];
        }
        if (values) {
            this.fillModelArray<RolloutRolloutPhase>(this, 'smartnics-status', values['smartnics-status'], RolloutRolloutPhase);
        } else {
            this['smartnics-status'] = [];
        }
        if (values && values['state'] != null) {
            this['state'] = values['state'];
        } else if (fillDefaults && RolloutRolloutStatus.hasDefaultValue('state')) {
            this['state'] = <RolloutRolloutStatus_state>  RolloutRolloutStatus.propInfo['state'].default;
        } else {
            this['state'] = null
        }
        if (values && values['completion-percent'] != null) {
            this['completion-percent'] = values['completion-percent'];
        } else if (fillDefaults && RolloutRolloutStatus.hasDefaultValue('completion-percent')) {
            this['completion-percent'] = RolloutRolloutStatus.propInfo['completion-percent'].default;
        } else {
            this['completion-percent'] = null
        }
        if (values && values['start-time'] != null) {
            this['start-time'] = values['start-time'];
        } else if (fillDefaults && RolloutRolloutStatus.hasDefaultValue('start-time')) {
            this['start-time'] = RolloutRolloutStatus.propInfo['start-time'].default;
        } else {
            this['start-time'] = null
        }
        if (values && values['end-time'] != null) {
            this['end-time'] = values['end-time'];
        } else if (fillDefaults && RolloutRolloutStatus.hasDefaultValue('end-time')) {
            this['end-time'] = RolloutRolloutStatus.propInfo['end-time'].default;
        } else {
            this['end-time'] = null
        }
        if (values && values['prev-version'] != null) {
            this['prev-version'] = values['prev-version'];
        } else if (fillDefaults && RolloutRolloutStatus.hasDefaultValue('prev-version')) {
            this['prev-version'] = RolloutRolloutStatus.propInfo['prev-version'].default;
        } else {
            this['prev-version'] = null
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'controller-nodes-status': new FormArray([]),
                'controller-services-status': new FormArray([]),
                'smartnics-status': new FormArray([]),
                'state': CustomFormControl(new FormControl(this['state'], [required, enumValidator(RolloutRolloutStatus_state), ]), RolloutRolloutStatus.propInfo['state']),
                'completion-percent': CustomFormControl(new FormControl(this['completion-percent']), RolloutRolloutStatus.propInfo['completion-percent']),
                'start-time': CustomFormControl(new FormControl(this['start-time']), RolloutRolloutStatus.propInfo['start-time']),
                'end-time': CustomFormControl(new FormControl(this['end-time']), RolloutRolloutStatus.propInfo['end-time']),
                'prev-version': CustomFormControl(new FormControl(this['prev-version']), RolloutRolloutStatus.propInfo['prev-version']),
            });
            // generate FormArray control elements
            this.fillFormArray<RolloutRolloutPhase>('controller-nodes-status', this['controller-nodes-status'], RolloutRolloutPhase);
            // generate FormArray control elements
            this.fillFormArray<RolloutRolloutPhase>('controller-services-status', this['controller-services-status'], RolloutRolloutPhase);
            // generate FormArray control elements
            this.fillFormArray<RolloutRolloutPhase>('smartnics-status', this['smartnics-status'], RolloutRolloutPhase);
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('controller-nodes-status') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('controller-nodes-status').get(field);
                control.updateValueAndValidity();
            });
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('controller-services-status') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('controller-services-status').get(field);
                control.updateValueAndValidity();
            });
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('smartnics-status') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('smartnics-status').get(field);
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
            this.fillModelArray<RolloutRolloutPhase>(this, 'controller-nodes-status', this['controller-nodes-status'], RolloutRolloutPhase);
            this.fillModelArray<RolloutRolloutPhase>(this, 'controller-services-status', this['controller-services-status'], RolloutRolloutPhase);
            this.fillModelArray<RolloutRolloutPhase>(this, 'smartnics-status', this['smartnics-status'], RolloutRolloutPhase);
            this._formGroup.controls['state'].setValue(this['state']);
            this._formGroup.controls['completion-percent'].setValue(this['completion-percent']);
            this._formGroup.controls['start-time'].setValue(this['start-time']);
            this._formGroup.controls['end-time'].setValue(this['end-time']);
            this._formGroup.controls['prev-version'].setValue(this['prev-version']);
        }
    }
}

