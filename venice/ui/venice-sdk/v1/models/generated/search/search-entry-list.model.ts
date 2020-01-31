/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { SearchEntry, ISearchEntry } from './search-entry.model';

export interface ISearchEntryList {
    'entries'?: Array<ISearchEntry>;
    '_ui'?: any;
}


export class SearchEntryList extends BaseModel implements ISearchEntryList {
    /** Field for holding arbitrary ui state */
    '_ui': any = {};
    'entries': Array<SearchEntry> = null;
    public static propInfo: { [prop in keyof ISearchEntryList]: PropInfoItem } = {
        'entries': {
            required: false,
            type: 'object'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return SearchEntryList.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return SearchEntryList.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (SearchEntryList.propInfo[prop] != null &&
                        SearchEntryList.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['entries'] = new Array<SearchEntry>();
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
        if (values) {
            this.fillModelArray<SearchEntry>(this, 'entries', values['entries'], SearchEntry);
        } else {
            this['entries'] = [];
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'entries': new FormArray([]),
            });
            // generate FormArray control elements
            this.fillFormArray<SearchEntry>('entries', this['entries'], SearchEntry);
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('entries') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('entries').get(field);
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
            this.fillModelArray<SearchEntry>(this, 'entries', this['entries'], SearchEntry);
        }
    }
}

