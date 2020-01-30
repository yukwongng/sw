/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { RolloutRolloutSpec_strategy,  RolloutRolloutSpec_strategy_uihint  } from './enums';
import { LabelsSelector, ILabelsSelector } from './labels-selector.model';
import { RolloutRolloutSpec_upgrade_type,  RolloutRolloutSpec_upgrade_type_uihint  } from './enums';

export interface IRolloutRolloutSpec {
    'version'?: string;
    'scheduled-start-time'?: Date;
    'scheduled-end-time'?: Date;
    'strategy': RolloutRolloutSpec_strategy;
    'max-parallel'?: number;
    'max-nic-failures-before-abort'?: number;
    'order-constraints'?: Array<ILabelsSelector>;
    'suspend'?: boolean;
    'dscs-only'?: boolean;
    'dsc-must-match-constraint'?: boolean;
    'upgrade-type': RolloutRolloutSpec_upgrade_type;
    'retry'?: boolean;
}


export class RolloutRolloutSpec extends BaseModel implements IRolloutRolloutSpec {
    /** New Version of the image to rollout to. */
    'version': string = null;
    /** Time, if specified, at which the rollout is supposed to start. (example:"2002-10-02T15:00:00.05Z"). */
    'scheduled-start-time': Date = null;
    /** ScheduledEndTime, if specified, after which the rollout is supposed to stop, if not completed by that time Typically represents the end of Maintenance window. (example:"2002-10-02T15:00:00.05Z"). */
    'scheduled-end-time': Date = null;
    'strategy': RolloutRolloutSpec_strategy = null;
    /** MaxParallel is the maximum number of nodes getting updated at any time This setting is applicable only to DistributedServiceCards. Controller nodes are always upgraded one after another. */
    'max-parallel': number = null;
    /** After these many failures are observed during DSC upgrade, the rollout process stops This setting applies to DSCs only. The controller nodes are rollout first and any failure there stops the rollout of DistributedServiceCards. */
    'max-nic-failures-before-abort': number = null;
    /** If specified, this is the sequence in which the DistributedServiceCards are upgraded based on their labels. if a DistributedServiceCard matches multiple constraints, the first one is used. Any DistributedServiceCard which does not match the constraints is upgraded at the end. This order is mainly for the DSCs on Hosts Controller nodes are always rollout one after other. */
    'order-constraints': Array<LabelsSelector> = null;
    /** When Set to true, the current rollout is suspended. Existing Nodes/Services/DistributedServiceCards in the middle of rollout continue rollout execution but any Nodes/Services/DistributedServiceCards which has not started Rollout will not be scheduled one. */
    'suspend': boolean = null;
    /** Dont upgrade Controller but only upgrade DistributedServiceCards. */
    'dscs-only': boolean = null;
    /** When DSCMustMatchConstraint is true, Any DSC which does not match OrderConstraints does not go through rollout. */
    'dsc-must-match-constraint': boolean = null;
    'upgrade-type': RolloutRolloutSpec_upgrade_type = null;
    /** If enabled, will retry rollout of failed naples within the maintenance window upto a max of 5 times. */
    'retry': boolean = null;
    public static propInfo: { [prop in keyof IRolloutRolloutSpec]: PropInfoItem } = {
        'version': {
            description:  `New Version of the image to rollout to.`,
            required: false,
            type: 'string'
        },
        'scheduled-start-time': {
            description:  `Time, if specified, at which the rollout is supposed to start. (example:"2002-10-02T15:00:00.05Z").`,
            required: false,
            type: 'Date'
        },
        'scheduled-end-time': {
            description:  `ScheduledEndTime, if specified, after which the rollout is supposed to stop, if not completed by that time Typically represents the end of Maintenance window. (example:"2002-10-02T15:00:00.05Z").`,
            required: false,
            type: 'Date'
        },
        'strategy': {
            enum: RolloutRolloutSpec_strategy_uihint,
            default: 'linear',
            required: true,
            type: 'string'
        },
        'max-parallel': {
            default: parseInt('2'),
            description:  `MaxParallel is the maximum number of nodes getting updated at any time This setting is applicable only to DistributedServiceCards. Controller nodes are always upgraded one after another.`,
            required: false,
            type: 'number'
        },
        'max-nic-failures-before-abort': {
            description:  `After these many failures are observed during DSC upgrade, the rollout process stops This setting applies to DSCs only. The controller nodes are rollout first and any failure there stops the rollout of DistributedServiceCards.`,
            required: false,
            type: 'number'
        },
        'order-constraints': {
            description:  `If specified, this is the sequence in which the DistributedServiceCards are upgraded based on their labels. if a DistributedServiceCard matches multiple constraints, the first one is used. Any DistributedServiceCard which does not match the constraints is upgraded at the end. This order is mainly for the DSCs on Hosts Controller nodes are always rollout one after other.`,
            required: false,
            type: 'object'
        },
        'suspend': {
            description:  `When Set to true, the current rollout is suspended. Existing Nodes/Services/DistributedServiceCards in the middle of rollout continue rollout execution but any Nodes/Services/DistributedServiceCards which has not started Rollout will not be scheduled one.`,
            required: false,
            type: 'boolean'
        },
        'dscs-only': {
            description:  `Dont upgrade Controller but only upgrade DistributedServiceCards.`,
            required: false,
            type: 'boolean'
        },
        'dsc-must-match-constraint': {
            description:  `When DSCMustMatchConstraint is true, Any DSC which does not match OrderConstraints does not go through rollout.`,
            required: false,
            type: 'boolean'
        },
        'upgrade-type': {
            enum: RolloutRolloutSpec_upgrade_type_uihint,
            default: 'disruptive',
            required: true,
            type: 'string'
        },
        'retry': {
            description:  `If enabled, will retry rollout of failed naples within the maintenance window upto a max of 5 times.`,
            required: false,
            type: 'boolean'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return RolloutRolloutSpec.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return RolloutRolloutSpec.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (RolloutRolloutSpec.propInfo[prop] != null &&
                        RolloutRolloutSpec.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['order-constraints'] = new Array<LabelsSelector>();
        this._inputValue = values;
        this.setValues(values, setDefaults);
    }

    /**
     * set the values for both the Model and the Form Group. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any, fillDefaults = true): void {
        if (values && values['version'] != null) {
            this['version'] = values['version'];
        } else if (fillDefaults && RolloutRolloutSpec.hasDefaultValue('version')) {
            this['version'] = RolloutRolloutSpec.propInfo['version'].default;
        } else {
            this['version'] = null
        }
        if (values && values['scheduled-start-time'] != null) {
            this['scheduled-start-time'] = values['scheduled-start-time'];
        } else if (fillDefaults && RolloutRolloutSpec.hasDefaultValue('scheduled-start-time')) {
            this['scheduled-start-time'] = RolloutRolloutSpec.propInfo['scheduled-start-time'].default;
        } else {
            this['scheduled-start-time'] = null
        }
        if (values && values['scheduled-end-time'] != null) {
            this['scheduled-end-time'] = values['scheduled-end-time'];
        } else if (fillDefaults && RolloutRolloutSpec.hasDefaultValue('scheduled-end-time')) {
            this['scheduled-end-time'] = RolloutRolloutSpec.propInfo['scheduled-end-time'].default;
        } else {
            this['scheduled-end-time'] = null
        }
        if (values && values['strategy'] != null) {
            this['strategy'] = values['strategy'];
        } else if (fillDefaults && RolloutRolloutSpec.hasDefaultValue('strategy')) {
            this['strategy'] = <RolloutRolloutSpec_strategy>  RolloutRolloutSpec.propInfo['strategy'].default;
        } else {
            this['strategy'] = null
        }
        if (values && values['max-parallel'] != null) {
            this['max-parallel'] = values['max-parallel'];
        } else if (fillDefaults && RolloutRolloutSpec.hasDefaultValue('max-parallel')) {
            this['max-parallel'] = RolloutRolloutSpec.propInfo['max-parallel'].default;
        } else {
            this['max-parallel'] = null
        }
        if (values && values['max-nic-failures-before-abort'] != null) {
            this['max-nic-failures-before-abort'] = values['max-nic-failures-before-abort'];
        } else if (fillDefaults && RolloutRolloutSpec.hasDefaultValue('max-nic-failures-before-abort')) {
            this['max-nic-failures-before-abort'] = RolloutRolloutSpec.propInfo['max-nic-failures-before-abort'].default;
        } else {
            this['max-nic-failures-before-abort'] = null
        }
        if (values) {
            this.fillModelArray<LabelsSelector>(this, 'order-constraints', values['order-constraints'], LabelsSelector);
        } else {
            this['order-constraints'] = [];
        }
        if (values && values['suspend'] != null) {
            this['suspend'] = values['suspend'];
        } else if (fillDefaults && RolloutRolloutSpec.hasDefaultValue('suspend')) {
            this['suspend'] = RolloutRolloutSpec.propInfo['suspend'].default;
        } else {
            this['suspend'] = null
        }
        if (values && values['dscs-only'] != null) {
            this['dscs-only'] = values['dscs-only'];
        } else if (fillDefaults && RolloutRolloutSpec.hasDefaultValue('dscs-only')) {
            this['dscs-only'] = RolloutRolloutSpec.propInfo['dscs-only'].default;
        } else {
            this['dscs-only'] = null
        }
        if (values && values['dsc-must-match-constraint'] != null) {
            this['dsc-must-match-constraint'] = values['dsc-must-match-constraint'];
        } else if (fillDefaults && RolloutRolloutSpec.hasDefaultValue('dsc-must-match-constraint')) {
            this['dsc-must-match-constraint'] = RolloutRolloutSpec.propInfo['dsc-must-match-constraint'].default;
        } else {
            this['dsc-must-match-constraint'] = null
        }
        if (values && values['upgrade-type'] != null) {
            this['upgrade-type'] = values['upgrade-type'];
        } else if (fillDefaults && RolloutRolloutSpec.hasDefaultValue('upgrade-type')) {
            this['upgrade-type'] = <RolloutRolloutSpec_upgrade_type>  RolloutRolloutSpec.propInfo['upgrade-type'].default;
        } else {
            this['upgrade-type'] = null
        }
        if (values && values['retry'] != null) {
            this['retry'] = values['retry'];
        } else if (fillDefaults && RolloutRolloutSpec.hasDefaultValue('retry')) {
            this['retry'] = RolloutRolloutSpec.propInfo['retry'].default;
        } else {
            this['retry'] = null
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'version': CustomFormControl(new FormControl(this['version']), RolloutRolloutSpec.propInfo['version']),
                'scheduled-start-time': CustomFormControl(new FormControl(this['scheduled-start-time']), RolloutRolloutSpec.propInfo['scheduled-start-time']),
                'scheduled-end-time': CustomFormControl(new FormControl(this['scheduled-end-time']), RolloutRolloutSpec.propInfo['scheduled-end-time']),
                'strategy': CustomFormControl(new FormControl(this['strategy'], [required, enumValidator(RolloutRolloutSpec_strategy), ]), RolloutRolloutSpec.propInfo['strategy']),
                'max-parallel': CustomFormControl(new FormControl(this['max-parallel']), RolloutRolloutSpec.propInfo['max-parallel']),
                'max-nic-failures-before-abort': CustomFormControl(new FormControl(this['max-nic-failures-before-abort']), RolloutRolloutSpec.propInfo['max-nic-failures-before-abort']),
                'order-constraints': new FormArray([]),
                'suspend': CustomFormControl(new FormControl(this['suspend']), RolloutRolloutSpec.propInfo['suspend']),
                'dscs-only': CustomFormControl(new FormControl(this['dscs-only']), RolloutRolloutSpec.propInfo['dscs-only']),
                'dsc-must-match-constraint': CustomFormControl(new FormControl(this['dsc-must-match-constraint']), RolloutRolloutSpec.propInfo['dsc-must-match-constraint']),
                'upgrade-type': CustomFormControl(new FormControl(this['upgrade-type'], [required, enumValidator(RolloutRolloutSpec_upgrade_type), ]), RolloutRolloutSpec.propInfo['upgrade-type']),
                'retry': CustomFormControl(new FormControl(this['retry']), RolloutRolloutSpec.propInfo['retry']),
            });
            // generate FormArray control elements
            this.fillFormArray<LabelsSelector>('order-constraints', this['order-constraints'], LabelsSelector);
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('order-constraints') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('order-constraints').get(field);
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
            this._formGroup.controls['version'].setValue(this['version']);
            this._formGroup.controls['scheduled-start-time'].setValue(this['scheduled-start-time']);
            this._formGroup.controls['scheduled-end-time'].setValue(this['scheduled-end-time']);
            this._formGroup.controls['strategy'].setValue(this['strategy']);
            this._formGroup.controls['max-parallel'].setValue(this['max-parallel']);
            this._formGroup.controls['max-nic-failures-before-abort'].setValue(this['max-nic-failures-before-abort']);
            this.fillModelArray<LabelsSelector>(this, 'order-constraints', this['order-constraints'], LabelsSelector);
            this._formGroup.controls['suspend'].setValue(this['suspend']);
            this._formGroup.controls['dscs-only'].setValue(this['dscs-only']);
            this._formGroup.controls['dsc-must-match-constraint'].setValue(this['dsc-must-match-constraint']);
            this._formGroup.controls['upgrade-type'].setValue(this['upgrade-type']);
            this._formGroup.controls['retry'].setValue(this['retry']);
        }
    }
}

