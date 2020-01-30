/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';


export interface IClusterBiosInfo {
    'vendor'?: string;
    'version'?: string;
    'fw-major-ver'?: string;
    'fw-minor-ver'?: string;
}


export class ClusterBiosInfo extends BaseModel implements IClusterBiosInfo {
    /** Vendor name. */
    'vendor': string = null;
    /** BIOS version. */
    'version': string = null;
    /** Firmware major release info. */
    'fw-major-ver': string = null;
    /** Firmware minor release info. */
    'fw-minor-ver': string = null;
    public static propInfo: { [prop in keyof IClusterBiosInfo]: PropInfoItem } = {
        'vendor': {
            description:  `Vendor name.`,
            required: false,
            type: 'string'
        },
        'version': {
            description:  `BIOS version.`,
            required: false,
            type: 'string'
        },
        'fw-major-ver': {
            description:  `Firmware major release info.`,
            required: false,
            type: 'string'
        },
        'fw-minor-ver': {
            description:  `Firmware minor release info.`,
            required: false,
            type: 'string'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return ClusterBiosInfo.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return ClusterBiosInfo.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (ClusterBiosInfo.propInfo[prop] != null &&
                        ClusterBiosInfo.propInfo[prop].default != null);
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
        if (values && values['vendor'] != null) {
            this['vendor'] = values['vendor'];
        } else if (fillDefaults && ClusterBiosInfo.hasDefaultValue('vendor')) {
            this['vendor'] = ClusterBiosInfo.propInfo['vendor'].default;
        } else {
            this['vendor'] = null
        }
        if (values && values['version'] != null) {
            this['version'] = values['version'];
        } else if (fillDefaults && ClusterBiosInfo.hasDefaultValue('version')) {
            this['version'] = ClusterBiosInfo.propInfo['version'].default;
        } else {
            this['version'] = null
        }
        if (values && values['fw-major-ver'] != null) {
            this['fw-major-ver'] = values['fw-major-ver'];
        } else if (fillDefaults && ClusterBiosInfo.hasDefaultValue('fw-major-ver')) {
            this['fw-major-ver'] = ClusterBiosInfo.propInfo['fw-major-ver'].default;
        } else {
            this['fw-major-ver'] = null
        }
        if (values && values['fw-minor-ver'] != null) {
            this['fw-minor-ver'] = values['fw-minor-ver'];
        } else if (fillDefaults && ClusterBiosInfo.hasDefaultValue('fw-minor-ver')) {
            this['fw-minor-ver'] = ClusterBiosInfo.propInfo['fw-minor-ver'].default;
        } else {
            this['fw-minor-ver'] = null
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'vendor': CustomFormControl(new FormControl(this['vendor']), ClusterBiosInfo.propInfo['vendor']),
                'version': CustomFormControl(new FormControl(this['version']), ClusterBiosInfo.propInfo['version']),
                'fw-major-ver': CustomFormControl(new FormControl(this['fw-major-ver']), ClusterBiosInfo.propInfo['fw-major-ver']),
                'fw-minor-ver': CustomFormControl(new FormControl(this['fw-minor-ver']), ClusterBiosInfo.propInfo['fw-minor-ver']),
            });
        }
        return this._formGroup;
    }

    setModelToBeFormGroupValues() {
        this.setValues(this.$formGroup.value, false);
    }

    setFormGroupValuesToBeModelValues() {
        if (this._formGroup) {
            this._formGroup.controls['vendor'].setValue(this['vendor']);
            this._formGroup.controls['version'].setValue(this['version']);
            this._formGroup.controls['fw-major-ver'].setValue(this['fw-major-ver']);
            this._formGroup.controls['fw-minor-ver'].setValue(this['fw-minor-ver']);
        }
    }
}

