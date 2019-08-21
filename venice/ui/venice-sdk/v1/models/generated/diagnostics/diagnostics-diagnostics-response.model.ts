/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { ApiAny, IApiAny } from './api-any.model';

export interface IDiagnosticsDiagnosticsResponse {
    'diagnostics'?: IApiAny;
}


export class DiagnosticsDiagnosticsResponse extends BaseModel implements IDiagnosticsDiagnosticsResponse {
    'diagnostics': ApiAny = null;
    public static propInfo: { [prop in keyof IDiagnosticsDiagnosticsResponse]: PropInfoItem } = {
        'diagnostics': {
            required: false,
            type: 'object'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return DiagnosticsDiagnosticsResponse.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return DiagnosticsDiagnosticsResponse.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (DiagnosticsDiagnosticsResponse.propInfo[prop] != null &&
                        DiagnosticsDiagnosticsResponse.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['diagnostics'] = new ApiAny();
        this._inputValue = values;
        this.setValues(values, setDefaults);
    }

    /**
     * set the values for both the Model and the Form Group. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any, fillDefaults = true): void {
        if (values) {
            this['diagnostics'].setValues(values['diagnostics'], fillDefaults);
        } else {
            this['diagnostics'].setValues(null, fillDefaults);
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'diagnostics': CustomFormGroup(this['diagnostics'].$formGroup, DiagnosticsDiagnosticsResponse.propInfo['diagnostics'].required),
            });
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('diagnostics') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('diagnostics').get(field);
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
            this['diagnostics'].setFormGroupValuesToBeModelValues();
        }
    }
}

