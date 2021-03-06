/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';


export interface IApiObjectRef {
    'tenant'?: string;
    'namespace'?: string;
    'kind'?: string;
    'name'?: string;
    'uri'?: string;
    '_ui'?: any;
}


export class ApiObjectRef extends BaseModel implements IApiObjectRef {
    /** Field for holding arbitrary ui state */
    '_ui': any = {};
    /** Tenant of the object. */
    'tenant': string = null;
    /** Namespace of the object, for scoped objects. */
    'namespace': string = null;
    /** Kind represents the type of the API object. */
    'kind': string = null;
    /** Name of the object, unique within a Namespace for scoped objects. */
    'name': string = null;
    /** URI is a link to accessing the referenced object. */
    'uri': string = null;
    public static propInfo: { [prop in keyof IApiObjectRef]: PropInfoItem } = {
        'tenant': {
            description:  `Tenant of the object.`,
            required: false,
            type: 'string'
        },
        'namespace': {
            description:  `Namespace of the object, for scoped objects.`,
            required: false,
            type: 'string'
        },
        'kind': {
            description:  `Kind represents the type of the API object.`,
            required: false,
            type: 'string'
        },
        'name': {
            description:  `Name of the object, unique within a Namespace for scoped objects.`,
            required: false,
            type: 'string'
        },
        'uri': {
            description:  `URI is a link to accessing the referenced object.`,
            required: false,
            type: 'string'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return ApiObjectRef.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return ApiObjectRef.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (ApiObjectRef.propInfo[prop] != null &&
                        ApiObjectRef.propInfo[prop].default != null);
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
        if (values && values['tenant'] != null) {
            this['tenant'] = values['tenant'];
        } else if (fillDefaults && ApiObjectRef.hasDefaultValue('tenant')) {
            this['tenant'] = ApiObjectRef.propInfo['tenant'].default;
        } else {
            this['tenant'] = null
        }
        if (values && values['namespace'] != null) {
            this['namespace'] = values['namespace'];
        } else if (fillDefaults && ApiObjectRef.hasDefaultValue('namespace')) {
            this['namespace'] = ApiObjectRef.propInfo['namespace'].default;
        } else {
            this['namespace'] = null
        }
        if (values && values['kind'] != null) {
            this['kind'] = values['kind'];
        } else if (fillDefaults && ApiObjectRef.hasDefaultValue('kind')) {
            this['kind'] = ApiObjectRef.propInfo['kind'].default;
        } else {
            this['kind'] = null
        }
        if (values && values['name'] != null) {
            this['name'] = values['name'];
        } else if (fillDefaults && ApiObjectRef.hasDefaultValue('name')) {
            this['name'] = ApiObjectRef.propInfo['name'].default;
        } else {
            this['name'] = null
        }
        if (values && values['uri'] != null) {
            this['uri'] = values['uri'];
        } else if (fillDefaults && ApiObjectRef.hasDefaultValue('uri')) {
            this['uri'] = ApiObjectRef.propInfo['uri'].default;
        } else {
            this['uri'] = null
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'tenant': CustomFormControl(new FormControl(this['tenant']), ApiObjectRef.propInfo['tenant']),
                'namespace': CustomFormControl(new FormControl(this['namespace']), ApiObjectRef.propInfo['namespace']),
                'kind': CustomFormControl(new FormControl(this['kind']), ApiObjectRef.propInfo['kind']),
                'name': CustomFormControl(new FormControl(this['name']), ApiObjectRef.propInfo['name']),
                'uri': CustomFormControl(new FormControl(this['uri']), ApiObjectRef.propInfo['uri']),
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
            this._formGroup.controls['kind'].setValue(this['kind']);
            this._formGroup.controls['name'].setValue(this['name']);
            this._formGroup.controls['uri'].setValue(this['uri']);
        }
    }
}

