/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';


export interface IClusterPropagationStatus {
    'generation-id'?: string;
    'updated'?: number;
    'pending'?: number;
    'min-version'?: string;
    'status'?: string;
    'pending-dscs'?: Array<string>;
    '_ui'?: any;
}


export class ClusterPropagationStatus extends BaseModel implements IClusterPropagationStatus {
    /** Field for holding arbitrary ui state */
    '_ui': any = {};
    /** The Generation ID this status is for. */
    'generation-id': string = null;
    /** The number of Naples that this version has already been pushed to. */
    'updated': number = null;
    /** Number of Naples pending. If this is 0 it can be assumed that everything is up to date. */
    'pending': number = null;
    /** The Version running on the slowest Naples. */
    'min-version': string = null;
    /** Textual description of propagation status. */
    'status': string = null;
    /** List of DSCs where propagation did not complete. */
    'pending-dscs': Array<string> = null;
    public static propInfo: { [prop in keyof IClusterPropagationStatus]: PropInfoItem } = {
        'generation-id': {
            description:  `The Generation ID this status is for.`,
            required: false,
            type: 'string'
        },
        'updated': {
            description:  `The number of Naples that this version has already been pushed to.`,
            required: false,
            type: 'number'
        },
        'pending': {
            description:  `Number of Naples pending. If this is 0 it can be assumed that everything is up to date.`,
            required: false,
            type: 'number'
        },
        'min-version': {
            description:  `The Version running on the slowest Naples.`,
            required: false,
            type: 'string'
        },
        'status': {
            description:  `Textual description of propagation status.`,
            required: false,
            type: 'string'
        },
        'pending-dscs': {
            description:  `List of DSCs where propagation did not complete.`,
            required: false,
            type: 'Array<string>'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return ClusterPropagationStatus.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return ClusterPropagationStatus.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (ClusterPropagationStatus.propInfo[prop] != null &&
                        ClusterPropagationStatus.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['pending-dscs'] = new Array<string>();
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
        if (values && values['generation-id'] != null) {
            this['generation-id'] = values['generation-id'];
        } else if (fillDefaults && ClusterPropagationStatus.hasDefaultValue('generation-id')) {
            this['generation-id'] = ClusterPropagationStatus.propInfo['generation-id'].default;
        } else {
            this['generation-id'] = null
        }
        if (values && values['updated'] != null) {
            this['updated'] = values['updated'];
        } else if (fillDefaults && ClusterPropagationStatus.hasDefaultValue('updated')) {
            this['updated'] = ClusterPropagationStatus.propInfo['updated'].default;
        } else {
            this['updated'] = null
        }
        if (values && values['pending'] != null) {
            this['pending'] = values['pending'];
        } else if (fillDefaults && ClusterPropagationStatus.hasDefaultValue('pending')) {
            this['pending'] = ClusterPropagationStatus.propInfo['pending'].default;
        } else {
            this['pending'] = null
        }
        if (values && values['min-version'] != null) {
            this['min-version'] = values['min-version'];
        } else if (fillDefaults && ClusterPropagationStatus.hasDefaultValue('min-version')) {
            this['min-version'] = ClusterPropagationStatus.propInfo['min-version'].default;
        } else {
            this['min-version'] = null
        }
        if (values && values['status'] != null) {
            this['status'] = values['status'];
        } else if (fillDefaults && ClusterPropagationStatus.hasDefaultValue('status')) {
            this['status'] = ClusterPropagationStatus.propInfo['status'].default;
        } else {
            this['status'] = null
        }
        if (values && values['pending-dscs'] != null) {
            this['pending-dscs'] = values['pending-dscs'];
        } else if (fillDefaults && ClusterPropagationStatus.hasDefaultValue('pending-dscs')) {
            this['pending-dscs'] = [ ClusterPropagationStatus.propInfo['pending-dscs'].default];
        } else {
            this['pending-dscs'] = [];
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'generation-id': CustomFormControl(new FormControl(this['generation-id']), ClusterPropagationStatus.propInfo['generation-id']),
                'updated': CustomFormControl(new FormControl(this['updated']), ClusterPropagationStatus.propInfo['updated']),
                'pending': CustomFormControl(new FormControl(this['pending']), ClusterPropagationStatus.propInfo['pending']),
                'min-version': CustomFormControl(new FormControl(this['min-version']), ClusterPropagationStatus.propInfo['min-version']),
                'status': CustomFormControl(new FormControl(this['status']), ClusterPropagationStatus.propInfo['status']),
                'pending-dscs': CustomFormControl(new FormControl(this['pending-dscs']), ClusterPropagationStatus.propInfo['pending-dscs']),
            });
        }
        return this._formGroup;
    }

    setModelToBeFormGroupValues() {
        this.setValues(this.$formGroup.value, false);
    }

    setFormGroupValuesToBeModelValues() {
        if (this._formGroup) {
            this._formGroup.controls['generation-id'].setValue(this['generation-id']);
            this._formGroup.controls['updated'].setValue(this['updated']);
            this._formGroup.controls['pending'].setValue(this['pending']);
            this._formGroup.controls['min-version'].setValue(this['min-version']);
            this._formGroup.controls['status'].setValue(this['status']);
            this._formGroup.controls['pending-dscs'].setValue(this['pending-dscs']);
        }
    }
}

