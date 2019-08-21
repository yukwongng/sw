/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';


export interface IAuthLdapAttributeMapping {
    'user'?: string;
    'user-object-class'?: string;
    'tenant'?: string;
    'group'?: string;
    'group-object-class'?: string;
    'email'?: string;
    'fullname'?: string;
}


export class AuthLdapAttributeMapping extends BaseModel implements IAuthLdapAttributeMapping {
    /** The name that the server uses for the UserID Attribute. */
    'user': string = null;
    'user-object-class': string = null;
    /** The tenant the server will use for authentication. */
    'tenant': string = null;
    /** The name that the server uses for the Group Member Attribute. By default, the attribute is set to member for standard schema, and sgMember for updated schema. */
    'group': string = null;
    'group-object-class': string = null;
    /** The name of the attribute for storing the users’ e-mail address. This attribute is primarily used for linked Authentication Server Users. It can also be used to identify users by their e-mail address in certificate authentication. */
    'email': string = null;
    /** The name that the server uses for the Name attribute. */
    'fullname': string = null;
    public static propInfo: { [prop in keyof IAuthLdapAttributeMapping]: PropInfoItem } = {
        'user': {
            description:  'The name that the server uses for the UserID Attribute.',
            required: false,
            type: 'string'
        },
        'user-object-class': {
            required: false,
            type: 'string'
        },
        'tenant': {
            description:  'The tenant the server will use for authentication.',
            required: false,
            type: 'string'
        },
        'group': {
            description:  'The name that the server uses for the Group Member Attribute. By default, the attribute is set to member for standard schema, and sgMember for updated schema.',
            required: false,
            type: 'string'
        },
        'group-object-class': {
            required: false,
            type: 'string'
        },
        'email': {
            description:  'The name of the attribute for storing the users’ e-mail address. This attribute is primarily used for linked Authentication Server Users. It can also be used to identify users by their e-mail address in certificate authentication.',
            required: false,
            type: 'string'
        },
        'fullname': {
            description:  'The name that the server uses for the Name attribute.',
            required: false,
            type: 'string'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return AuthLdapAttributeMapping.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return AuthLdapAttributeMapping.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (AuthLdapAttributeMapping.propInfo[prop] != null &&
                        AuthLdapAttributeMapping.propInfo[prop].default != null);
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
        if (values && values['user'] != null) {
            this['user'] = values['user'];
        } else if (fillDefaults && AuthLdapAttributeMapping.hasDefaultValue('user')) {
            this['user'] = AuthLdapAttributeMapping.propInfo['user'].default;
        } else {
            this['user'] = null
        }
        if (values && values['user-object-class'] != null) {
            this['user-object-class'] = values['user-object-class'];
        } else if (fillDefaults && AuthLdapAttributeMapping.hasDefaultValue('user-object-class')) {
            this['user-object-class'] = AuthLdapAttributeMapping.propInfo['user-object-class'].default;
        } else {
            this['user-object-class'] = null
        }
        if (values && values['tenant'] != null) {
            this['tenant'] = values['tenant'];
        } else if (fillDefaults && AuthLdapAttributeMapping.hasDefaultValue('tenant')) {
            this['tenant'] = AuthLdapAttributeMapping.propInfo['tenant'].default;
        } else {
            this['tenant'] = null
        }
        if (values && values['group'] != null) {
            this['group'] = values['group'];
        } else if (fillDefaults && AuthLdapAttributeMapping.hasDefaultValue('group')) {
            this['group'] = AuthLdapAttributeMapping.propInfo['group'].default;
        } else {
            this['group'] = null
        }
        if (values && values['group-object-class'] != null) {
            this['group-object-class'] = values['group-object-class'];
        } else if (fillDefaults && AuthLdapAttributeMapping.hasDefaultValue('group-object-class')) {
            this['group-object-class'] = AuthLdapAttributeMapping.propInfo['group-object-class'].default;
        } else {
            this['group-object-class'] = null
        }
        if (values && values['email'] != null) {
            this['email'] = values['email'];
        } else if (fillDefaults && AuthLdapAttributeMapping.hasDefaultValue('email')) {
            this['email'] = AuthLdapAttributeMapping.propInfo['email'].default;
        } else {
            this['email'] = null
        }
        if (values && values['fullname'] != null) {
            this['fullname'] = values['fullname'];
        } else if (fillDefaults && AuthLdapAttributeMapping.hasDefaultValue('fullname')) {
            this['fullname'] = AuthLdapAttributeMapping.propInfo['fullname'].default;
        } else {
            this['fullname'] = null
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'user': CustomFormControl(new FormControl(this['user']), AuthLdapAttributeMapping.propInfo['user']),
                'user-object-class': CustomFormControl(new FormControl(this['user-object-class']), AuthLdapAttributeMapping.propInfo['user-object-class']),
                'tenant': CustomFormControl(new FormControl(this['tenant']), AuthLdapAttributeMapping.propInfo['tenant']),
                'group': CustomFormControl(new FormControl(this['group']), AuthLdapAttributeMapping.propInfo['group']),
                'group-object-class': CustomFormControl(new FormControl(this['group-object-class']), AuthLdapAttributeMapping.propInfo['group-object-class']),
                'email': CustomFormControl(new FormControl(this['email']), AuthLdapAttributeMapping.propInfo['email']),
                'fullname': CustomFormControl(new FormControl(this['fullname']), AuthLdapAttributeMapping.propInfo['fullname']),
            });
        }
        return this._formGroup;
    }

    setModelToBeFormGroupValues() {
        this.setValues(this.$formGroup.value, false);
    }

    setFormGroupValuesToBeModelValues() {
        if (this._formGroup) {
            this._formGroup.controls['user'].setValue(this['user']);
            this._formGroup.controls['user-object-class'].setValue(this['user-object-class']);
            this._formGroup.controls['tenant'].setValue(this['tenant']);
            this._formGroup.controls['group'].setValue(this['group']);
            this._formGroup.controls['group-object-class'].setValue(this['group-object-class']);
            this._formGroup.controls['email'].setValue(this['email']);
            this._formGroup.controls['fullname'].setValue(this['fullname']);
        }
    }
}

