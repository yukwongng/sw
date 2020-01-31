/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { Telemetry_queryFwlog, ITelemetry_queryFwlog } from './telemetry-query-fwlog.model';

export interface ITelemetry_queryFwlogsQueryResult {
    'statement_id'?: number;
    'logs'?: Array<ITelemetry_queryFwlog>;
    '_ui'?: any;
}


export class Telemetry_queryFwlogsQueryResult extends BaseModel implements ITelemetry_queryFwlogsQueryResult {
    /** Field for holding arbitrary ui state */
    '_ui': any = {};
    'statement_id': number = null;
    'logs': Array<Telemetry_queryFwlog> = null;
    public static propInfo: { [prop in keyof ITelemetry_queryFwlogsQueryResult]: PropInfoItem } = {
        'statement_id': {
            required: false,
            type: 'number'
        },
        'logs': {
            required: false,
            type: 'object'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return Telemetry_queryFwlogsQueryResult.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return Telemetry_queryFwlogsQueryResult.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (Telemetry_queryFwlogsQueryResult.propInfo[prop] != null &&
                        Telemetry_queryFwlogsQueryResult.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['logs'] = new Array<Telemetry_queryFwlog>();
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
        if (values && values['statement_id'] != null) {
            this['statement_id'] = values['statement_id'];
        } else if (fillDefaults && Telemetry_queryFwlogsQueryResult.hasDefaultValue('statement_id')) {
            this['statement_id'] = Telemetry_queryFwlogsQueryResult.propInfo['statement_id'].default;
        } else {
            this['statement_id'] = null
        }
        if (values) {
            this.fillModelArray<Telemetry_queryFwlog>(this, 'logs', values['logs'], Telemetry_queryFwlog);
        } else {
            this['logs'] = [];
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'statement_id': CustomFormControl(new FormControl(this['statement_id']), Telemetry_queryFwlogsQueryResult.propInfo['statement_id']),
                'logs': new FormArray([]),
            });
            // generate FormArray control elements
            this.fillFormArray<Telemetry_queryFwlog>('logs', this['logs'], Telemetry_queryFwlog);
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('logs') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('logs').get(field);
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
            this._formGroup.controls['statement_id'].setValue(this['statement_id']);
            this.fillModelArray<Telemetry_queryFwlog>(this, 'logs', this['logs'], Telemetry_queryFwlog);
        }
    }
}

