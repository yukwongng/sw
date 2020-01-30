/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { ApiAny, IApiAny } from './api-any.model';

export interface ISearchEntry {
    'object'?: IApiAny;
}


export class SearchEntry extends BaseModel implements ISearchEntry {
    /** TODO: Right now our api codegen does not support nested inline and hence this attribute cannot be be made embedded/inline. api.Any is already had embededed attribute Any. Once infra supports nested inline or an alternative, this attribute should be embedded and made inline to make json response user friendly for search. */
    'object': ApiAny = null;
    public static propInfo: { [prop in keyof ISearchEntry]: PropInfoItem } = {
        'object': {
            description:  `TODO: Right now our api codegen does not support nested inline and hence this attribute cannot be be made embedded/inline. api.Any is already had embededed attribute Any. Once infra supports nested inline or an alternative, this attribute should be embedded and made inline to make json response user friendly for search.`,
            required: false,
            type: 'object'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return SearchEntry.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return SearchEntry.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (SearchEntry.propInfo[prop] != null &&
                        SearchEntry.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['object'] = new ApiAny();
        this._inputValue = values;
        this.setValues(values, setDefaults);
    }

    /**
     * set the values for both the Model and the Form Group. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any, fillDefaults = true): void {
        if (values) {
            this['object'].setValues(values['object'], fillDefaults);
        } else {
            this['object'].setValues(null, fillDefaults);
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'object': CustomFormGroup(this['object'].$formGroup, SearchEntry.propInfo['object'].required),
            });
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('object') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('object').get(field);
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
            this['object'].setFormGroupValuesToBeModelValues();
        }
    }
}

