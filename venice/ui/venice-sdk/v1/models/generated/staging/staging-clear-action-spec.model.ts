/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { StagingItemId, IStagingItemId } from './staging-item-id.model';

export interface IStagingClearActionSpec {
    'items'?: Array<IStagingItemId>;
    '_ui'?: any;
}


export class StagingClearActionSpec extends BaseModel implements IStagingClearActionSpec {
    /** Field for holding arbitrary ui state */
    '_ui': any = {};
    /** Empty Items indicates everyting in the buffer. */
    'items': Array<StagingItemId> = null;
    public static propInfo: { [prop in keyof IStagingClearActionSpec]: PropInfoItem } = {
        'items': {
            description:  `Empty Items indicates everyting in the buffer.`,
            required: false,
            type: 'object'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return StagingClearActionSpec.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return StagingClearActionSpec.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (StagingClearActionSpec.propInfo[prop] != null &&
                        StagingClearActionSpec.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['items'] = new Array<StagingItemId>();
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
            this.fillModelArray<StagingItemId>(this, 'items', values['items'], StagingItemId);
        } else {
            this['items'] = [];
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'items': new FormArray([]),
            });
            // generate FormArray control elements
            this.fillFormArray<StagingItemId>('items', this['items'], StagingItemId);
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('items') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('items').get(field);
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
            this.fillModelArray<StagingItemId>(this, 'items', this['items'], StagingItemId);
        }
    }
}

