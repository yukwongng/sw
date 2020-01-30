/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';


export interface IClusterHostStatus {
    'admitted-dscs'?: Array<string>;
}


export class ClusterHostStatus extends BaseModel implements IClusterHostStatus {
    /** AdmittedDSCs contains a list of admitted DistributedServiceCards that are on this host. */
    'admitted-dscs': Array<string> = null;
    public static propInfo: { [prop in keyof IClusterHostStatus]: PropInfoItem } = {
        'admitted-dscs': {
            description:  `AdmittedDSCs contains a list of admitted DistributedServiceCards that are on this host.`,
            required: false,
            type: 'Array<string>'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return ClusterHostStatus.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return ClusterHostStatus.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (ClusterHostStatus.propInfo[prop] != null &&
                        ClusterHostStatus.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['admitted-dscs'] = new Array<string>();
        this._inputValue = values;
        this.setValues(values, setDefaults);
    }

    /**
     * set the values for both the Model and the Form Group. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any, fillDefaults = true): void {
        if (values && values['admitted-dscs'] != null) {
            this['admitted-dscs'] = values['admitted-dscs'];
        } else if (fillDefaults && ClusterHostStatus.hasDefaultValue('admitted-dscs')) {
            this['admitted-dscs'] = [ ClusterHostStatus.propInfo['admitted-dscs'].default];
        } else {
            this['admitted-dscs'] = [];
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'admitted-dscs': CustomFormControl(new FormControl(this['admitted-dscs']), ClusterHostStatus.propInfo['admitted-dscs']),
            });
        }
        return this._formGroup;
    }

    setModelToBeFormGroupValues() {
        this.setValues(this.$formGroup.value, false);
    }

    setFormGroupValuesToBeModelValues() {
        if (this._formGroup) {
            this._formGroup.controls['admitted-dscs'].setValue(this['admitted-dscs']);
        }
    }
}

