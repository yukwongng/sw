/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { ClusterClusterCondition_type,  ClusterClusterCondition_type_uihint  } from './enums';
import { ClusterClusterCondition_status,  ClusterClusterCondition_status_uihint  } from './enums';

export interface IClusterClusterCondition {
    'type': ClusterClusterCondition_type;
    'status': ClusterClusterCondition_status;
    'last-transition-time'?: string;
    'reason'?: string;
    'message'?: string;
}


export class ClusterClusterCondition extends BaseModel implements IClusterClusterCondition {
    /** Type indicates a certain DSC condition. */
    'type': ClusterClusterCondition_type = null;
    /** Condition Status. */
    'status': ClusterClusterCondition_status = null;
    /** The last time the condition transitioned. */
    'last-transition-time': string = null;
    /** The reason for the condition's last transition. */
    'reason': string = null;
    /** A detailed message indicating details about the transition. */
    'message': string = null;
    public static propInfo: { [prop in keyof IClusterClusterCondition]: PropInfoItem } = {
        'type': {
            enum: ClusterClusterCondition_type_uihint,
            default: 'healthy',
            description:  `Type indicates a certain DSC condition.`,
            required: true,
            type: 'string'
        },
        'status': {
            enum: ClusterClusterCondition_status_uihint,
            default: 'unknown',
            description:  `Condition Status.`,
            required: true,
            type: 'string'
        },
        'last-transition-time': {
            description:  `The last time the condition transitioned.`,
            required: false,
            type: 'string'
        },
        'reason': {
            description:  `The reason for the condition's last transition.`,
            required: false,
            type: 'string'
        },
        'message': {
            description:  `A detailed message indicating details about the transition.`,
            required: false,
            type: 'string'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return ClusterClusterCondition.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return ClusterClusterCondition.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (ClusterClusterCondition.propInfo[prop] != null &&
                        ClusterClusterCondition.propInfo[prop].default != null);
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
        if (values && values['type'] != null) {
            this['type'] = values['type'];
        } else if (fillDefaults && ClusterClusterCondition.hasDefaultValue('type')) {
            this['type'] = <ClusterClusterCondition_type>  ClusterClusterCondition.propInfo['type'].default;
        } else {
            this['type'] = null
        }
        if (values && values['status'] != null) {
            this['status'] = values['status'];
        } else if (fillDefaults && ClusterClusterCondition.hasDefaultValue('status')) {
            this['status'] = <ClusterClusterCondition_status>  ClusterClusterCondition.propInfo['status'].default;
        } else {
            this['status'] = null
        }
        if (values && values['last-transition-time'] != null) {
            this['last-transition-time'] = values['last-transition-time'];
        } else if (fillDefaults && ClusterClusterCondition.hasDefaultValue('last-transition-time')) {
            this['last-transition-time'] = ClusterClusterCondition.propInfo['last-transition-time'].default;
        } else {
            this['last-transition-time'] = null
        }
        if (values && values['reason'] != null) {
            this['reason'] = values['reason'];
        } else if (fillDefaults && ClusterClusterCondition.hasDefaultValue('reason')) {
            this['reason'] = ClusterClusterCondition.propInfo['reason'].default;
        } else {
            this['reason'] = null
        }
        if (values && values['message'] != null) {
            this['message'] = values['message'];
        } else if (fillDefaults && ClusterClusterCondition.hasDefaultValue('message')) {
            this['message'] = ClusterClusterCondition.propInfo['message'].default;
        } else {
            this['message'] = null
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'type': CustomFormControl(new FormControl(this['type'], [required, enumValidator(ClusterClusterCondition_type), ]), ClusterClusterCondition.propInfo['type']),
                'status': CustomFormControl(new FormControl(this['status'], [required, enumValidator(ClusterClusterCondition_status), ]), ClusterClusterCondition.propInfo['status']),
                'last-transition-time': CustomFormControl(new FormControl(this['last-transition-time']), ClusterClusterCondition.propInfo['last-transition-time']),
                'reason': CustomFormControl(new FormControl(this['reason']), ClusterClusterCondition.propInfo['reason']),
                'message': CustomFormControl(new FormControl(this['message']), ClusterClusterCondition.propInfo['message']),
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
            this._formGroup.controls['status'].setValue(this['status']);
            this._formGroup.controls['last-transition-time'].setValue(this['last-transition-time']);
            this._formGroup.controls['reason'].setValue(this['reason']);
            this._formGroup.controls['message'].setValue(this['message']);
        }
    }
}

