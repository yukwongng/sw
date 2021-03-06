/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { Telemetry_queryFwlogsQuerySpec, ITelemetry_queryFwlogsQuerySpec } from './telemetry-query-fwlogs-query-spec.model';

export interface ITelemetry_queryFwlogsQueryList {
    'tenant'?: string;
    'namespace'?: string;
    'queries'?: Array<ITelemetry_queryFwlogsQuerySpec>;
    '_ui'?: any;
}


export class Telemetry_queryFwlogsQueryList extends BaseModel implements ITelemetry_queryFwlogsQueryList {
    /** Field for holding arbitrary ui state */
    '_ui': any = {};
    /** Tenant for the request. */
    'tenant': string = null;
    /** Namespace for the request. */
    'namespace': string = null;
    /** List of queries to execute. */
    'queries': Array<Telemetry_queryFwlogsQuerySpec> = null;
    public static propInfo: { [prop in keyof ITelemetry_queryFwlogsQueryList]: PropInfoItem } = {
        'tenant': {
            description:  `Tenant for the request.`,
            required: false,
            type: 'string'
        },
        'namespace': {
            description:  `Namespace for the request.`,
            required: false,
            type: 'string'
        },
        'queries': {
            description:  `List of queries to execute.`,
            required: false,
            type: 'object'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return Telemetry_queryFwlogsQueryList.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return Telemetry_queryFwlogsQueryList.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (Telemetry_queryFwlogsQueryList.propInfo[prop] != null &&
                        Telemetry_queryFwlogsQueryList.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['queries'] = new Array<Telemetry_queryFwlogsQuerySpec>();
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
        if (values && values['tenant'] != null) {
            this['tenant'] = values['tenant'];
        } else if (fillDefaults && Telemetry_queryFwlogsQueryList.hasDefaultValue('tenant')) {
            this['tenant'] = Telemetry_queryFwlogsQueryList.propInfo['tenant'].default;
        } else {
            this['tenant'] = null
        }
        if (values && values['namespace'] != null) {
            this['namespace'] = values['namespace'];
        } else if (fillDefaults && Telemetry_queryFwlogsQueryList.hasDefaultValue('namespace')) {
            this['namespace'] = Telemetry_queryFwlogsQueryList.propInfo['namespace'].default;
        } else {
            this['namespace'] = null
        }
        if (values) {
            this.fillModelArray<Telemetry_queryFwlogsQuerySpec>(this, 'queries', values['queries'], Telemetry_queryFwlogsQuerySpec);
        } else {
            this['queries'] = [];
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'tenant': CustomFormControl(new FormControl(this['tenant']), Telemetry_queryFwlogsQueryList.propInfo['tenant']),
                'namespace': CustomFormControl(new FormControl(this['namespace']), Telemetry_queryFwlogsQueryList.propInfo['namespace']),
                'queries': new FormArray([]),
            });
            // generate FormArray control elements
            this.fillFormArray<Telemetry_queryFwlogsQuerySpec>('queries', this['queries'], Telemetry_queryFwlogsQuerySpec);
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('queries') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('queries').get(field);
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
            this._formGroup.controls['tenant'].setValue(this['tenant']);
            this._formGroup.controls['namespace'].setValue(this['namespace']);
            this.fillModelArray<Telemetry_queryFwlogsQuerySpec>(this, 'queries', this['queries'], Telemetry_queryFwlogsQuerySpec);
        }
    }
}

