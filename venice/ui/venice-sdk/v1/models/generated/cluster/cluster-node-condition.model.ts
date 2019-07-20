/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from './base-model';

import { ClusterNodeCondition_type,  ClusterNodeCondition_type_uihint  } from './enums';
import { ClusterNodeCondition_status,  ClusterNodeCondition_status_uihint  } from './enums';

export interface IClusterNodeCondition {
    'type': ClusterNodeCondition_type;
    'status': ClusterNodeCondition_status;
    'last-transition-time'?: string;
    'reason'?: string;
    'message'?: string;
}


export class ClusterNodeCondition extends BaseModel implements IClusterNodeCondition {
    'type': ClusterNodeCondition_type = null;
    'status': ClusterNodeCondition_status = null;
    'last-transition-time': string = null;
    'reason': string = null;
    /** A detailed message indicating details about the transition. */
    'message': string = null;
    public static propInfo: { [prop: string]: PropInfoItem } = {
        'type': {
            enum: ClusterNodeCondition_type_uihint,
            default: 'leader',
            required: true,
            type: 'string'
        },
        'status': {
            enum: ClusterNodeCondition_status_uihint,
            default: 'unknown',
            required: true,
            type: 'string'
        },
        'last-transition-time': {
            required: false,
            type: 'string'
        },
        'reason': {
            required: false,
            type: 'string'
        },
        'message': {
            description:  'A detailed message indicating details about the transition.',
            required: false,
            type: 'string'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return ClusterNodeCondition.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return ClusterNodeCondition.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (ClusterNodeCondition.propInfo[prop] != null &&
                        ClusterNodeCondition.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this.setValues(values, setDefaults);
    }

    /**
     * set the values for both the Model and the Form Group. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any, fillDefaults = true): void {
        if (values && values['type'] != null) {
            this['type'] = values['type'];
        } else if (fillDefaults && ClusterNodeCondition.hasDefaultValue('type')) {
            this['type'] = <ClusterNodeCondition_type>  ClusterNodeCondition.propInfo['type'].default;
        } else {
            this['type'] = null
        }
        if (values && values['status'] != null) {
            this['status'] = values['status'];
        } else if (fillDefaults && ClusterNodeCondition.hasDefaultValue('status')) {
            this['status'] = <ClusterNodeCondition_status>  ClusterNodeCondition.propInfo['status'].default;
        } else {
            this['status'] = null
        }
        if (values && values['last-transition-time'] != null) {
            this['last-transition-time'] = values['last-transition-time'];
        } else if (fillDefaults && ClusterNodeCondition.hasDefaultValue('last-transition-time')) {
            this['last-transition-time'] = ClusterNodeCondition.propInfo['last-transition-time'].default;
        } else {
            this['last-transition-time'] = null
        }
        if (values && values['reason'] != null) {
            this['reason'] = values['reason'];
        } else if (fillDefaults && ClusterNodeCondition.hasDefaultValue('reason')) {
            this['reason'] = ClusterNodeCondition.propInfo['reason'].default;
        } else {
            this['reason'] = null
        }
        if (values && values['message'] != null) {
            this['message'] = values['message'];
        } else if (fillDefaults && ClusterNodeCondition.hasDefaultValue('message')) {
            this['message'] = ClusterNodeCondition.propInfo['message'].default;
        } else {
            this['message'] = null
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'type': CustomFormControl(new FormControl(this['type'], [required, enumValidator(ClusterNodeCondition_type), ]), ClusterNodeCondition.propInfo['type']),
                'status': CustomFormControl(new FormControl(this['status'], [required, enumValidator(ClusterNodeCondition_status), ]), ClusterNodeCondition.propInfo['status']),
                'last-transition-time': CustomFormControl(new FormControl(this['last-transition-time']), ClusterNodeCondition.propInfo['last-transition-time']),
                'reason': CustomFormControl(new FormControl(this['reason']), ClusterNodeCondition.propInfo['reason']),
                'message': CustomFormControl(new FormControl(this['message']), ClusterNodeCondition.propInfo['message']),
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

