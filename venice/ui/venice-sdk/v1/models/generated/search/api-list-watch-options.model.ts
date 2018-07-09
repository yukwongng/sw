/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, enumValidator } from './validators';
import { BaseModel, EnumDef } from './base-model';


export interface IApiListWatchOptions {
    'name'?: string;
    'tenant'?: string;
    'namespace'?: string;
    'resource-version'?: string;
    'uuid'?: string;
    'labels'?: object;
    'creation-time'?: Date;
    'mod-time'?: Date;
    'self-link'?: string;
    'label-selector'?: string;
    'field-selector'?: string;
    'prefix-watch'?: boolean;
    'field-change-selector'?: Array<string>;
    'from'?: number;
    'max-results'?: number;
}


export class ApiListWatchOptions extends BaseModel implements IApiListWatchOptions {
    'name': string;
    'tenant': string;
    'namespace': string;
    'resource-version': string;
    'uuid': string;
    'labels': object;
    'creation-time': Date;
    'mod-time': Date;
    'self-link': string;
    'label-selector': string;
    'field-selector': string;
    'prefix-watch': boolean;
    'field-change-selector': Array<string>;
    'from': number;
    /** max. number of events to be fetched for the request. */
    'max-results': number;
    public static enumProperties: { [key: string] : EnumDef } = {
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any) {
        super();
        this['field-change-selector'] = new Array<string>();
        if (values) {
            this.setValues(values);
        }
    }

    /**
     * set the values.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any): void {
        if (values) {
            this['name'] = values['name'];
            this['tenant'] = values['tenant'];
            this['namespace'] = values['namespace'];
            this['resource-version'] = values['resource-version'];
            this['uuid'] = values['uuid'];
            this['labels'] = values['labels'];
            this['creation-time'] = values['creation-time'];
            this['mod-time'] = values['mod-time'];
            this['self-link'] = values['self-link'];
            this['label-selector'] = values['label-selector'];
            this['field-selector'] = values['field-selector'];
            this['prefix-watch'] = values['prefix-watch'];
            this.fillModelArray<string>(this, 'field-change-selector', values['field-change-selector']);
            this['from'] = values['from'];
            this['max-results'] = values['max-results'];
        }
    }

    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'name': new FormControl(this['name']),
                'tenant': new FormControl(this['tenant']),
                'namespace': new FormControl(this['namespace']),
                'resource-version': new FormControl(this['resource-version']),
                'uuid': new FormControl(this['uuid']),
                'labels': new FormControl(this['labels']),
                'creation-time': new FormControl(this['creation-time']),
                'mod-time': new FormControl(this['mod-time']),
                'self-link': new FormControl(this['self-link']),
                'label-selector': new FormControl(this['label-selector']),
                'field-selector': new FormControl(this['field-selector']),
                'prefix-watch': new FormControl(this['prefix-watch']),
                'field-change-selector': new FormArray([]),
                'from': new FormControl(this['from']),
                'max-results': new FormControl(this['max-results']),
            });
            // generate FormArray control elements
            this.fillFormArray<string>('field-change-selector', this['field-change-selector']);
        }
        return this._formGroup;
    }

    setFormGroupValues() {
        if (this._formGroup) {
            this._formGroup.controls['name'].setValue(this['name']);
            this._formGroup.controls['tenant'].setValue(this['tenant']);
            this._formGroup.controls['namespace'].setValue(this['namespace']);
            this._formGroup.controls['resource-version'].setValue(this['resource-version']);
            this._formGroup.controls['uuid'].setValue(this['uuid']);
            this._formGroup.controls['labels'].setValue(this['labels']);
            this._formGroup.controls['creation-time'].setValue(this['creation-time']);
            this._formGroup.controls['mod-time'].setValue(this['mod-time']);
            this._formGroup.controls['self-link'].setValue(this['self-link']);
            this._formGroup.controls['label-selector'].setValue(this['label-selector']);
            this._formGroup.controls['field-selector'].setValue(this['field-selector']);
            this._formGroup.controls['prefix-watch'].setValue(this['prefix-watch']);
            this.fillModelArray<string>(this, 'field-change-selector', this['field-change-selector']);
            this._formGroup.controls['from'].setValue(this['from']);
            this._formGroup.controls['max-results'].setValue(this['max-results']);
        }
    }
}

