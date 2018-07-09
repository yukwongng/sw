/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, enumValidator } from './validators';
import { BaseModel, EnumDef } from './base-model';


export interface IProtobufAny {
    'type_url'?: string;
    'value'?: string;
}


export class ProtobufAny extends BaseModel implements IProtobufAny {
    /** A URL/resource name whose content describes the type of the
serialized protocol buffer message.

For URLs which use the scheme `http`, `https`, or no scheme, the
following restrictions and interpretations apply:

* If no scheme is provided, `https` is assumed.
* The last segment of the URL's path must represent the fully
  qualified name of the type (as in `path/google.protobuf.Duration`).
  The name should be in a canonical form (e.g., leading "." is
  not accepted).
* An HTTP GET on the URL must yield a [google.protobuf.Type][]
  value in binary format, or produce an error.
* Applications are allowed to cache lookup results based on the
  URL, or have them precompiled into a binary to avoid any
  lookup. Therefore, binary compatibility needs to be preserved
  on changes to types. (Use versioned type names to manage
  breaking changes.)

Schemes other than `http`, `https` (or the empty scheme) might be
used with implementation specific semantics. */
    'type_url': string;
    /** Must be a valid serialized protocol buffer of the above specified type. */
    'value': string;
    public static enumProperties: { [key: string] : EnumDef } = {
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any) {
        super();
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
            this['type_url'] = values['type_url'];
            this['value'] = values['value'];
        }
    }

    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'type_url': new FormControl(this['type_url']),
                'value': new FormControl(this['value']),
            });
        }
        return this._formGroup;
    }

    setFormGroupValues() {
        if (this._formGroup) {
            this._formGroup.controls['type_url'].setValue(this['type_url']);
            this._formGroup.controls['value'].setValue(this['value']);
        }
    }
}

