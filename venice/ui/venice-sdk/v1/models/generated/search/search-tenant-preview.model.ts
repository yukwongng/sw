/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';


export interface ISearchTenantPreview {
    'tenants'?: object;
    '_ui'?: any;
}


export class SearchTenantPreview extends BaseModel implements ISearchTenantPreview {
    /** Field for holding arbitrary ui state */
    '_ui': any = {};
    'tenants': object = null;
    public static propInfo: { [prop in keyof ISearchTenantPreview]: PropInfoItem } = {
        'tenants': {
            required: false,
            type: 'object'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return SearchTenantPreview.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return SearchTenantPreview.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (SearchTenantPreview.propInfo[prop] != null &&
                        SearchTenantPreview.propInfo[prop].default != null);
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
        if (values && values['_ui']) {
            this['_ui'] = values['_ui']
        }
        if (values && values['tenants'] != null) {
            this['tenants'] = values['tenants'];
        } else if (fillDefaults && SearchTenantPreview.hasDefaultValue('tenants')) {
            this['tenants'] = SearchTenantPreview.propInfo['tenants'].default;
        } else {
            this['tenants'] = null
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'tenants': CustomFormControl(new FormControl(this['tenants']), SearchTenantPreview.propInfo['tenants']),
            });
        }
        return this._formGroup;
    }

    setModelToBeFormGroupValues() {
        this.setValues(this.$formGroup.value, false);
    }

    setFormGroupValuesToBeModelValues() {
        if (this._formGroup) {
            this._formGroup.controls['tenants'].setValue(this['tenants']);
        }
    }
}

