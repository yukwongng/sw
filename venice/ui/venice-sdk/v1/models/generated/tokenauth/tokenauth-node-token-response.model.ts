/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';


export interface ITokenauthNodeTokenResponse {
    'Token'?: string;
    '_ui'?: any;
}


export class TokenauthNodeTokenResponse extends BaseModel implements ITokenauthNodeTokenResponse {
    /** Field for holding arbitrary ui state */
    '_ui': any = {};
    'Token': string = null;
    public static propInfo: { [prop in keyof ITokenauthNodeTokenResponse]: PropInfoItem } = {
        'Token': {
            required: false,
            type: 'string'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return TokenauthNodeTokenResponse.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return TokenauthNodeTokenResponse.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (TokenauthNodeTokenResponse.propInfo[prop] != null &&
                        TokenauthNodeTokenResponse.propInfo[prop].default != null);
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
        if (values && values['Token'] != null) {
            this['Token'] = values['Token'];
        } else if (fillDefaults && TokenauthNodeTokenResponse.hasDefaultValue('Token')) {
            this['Token'] = TokenauthNodeTokenResponse.propInfo['Token'].default;
        } else {
            this['Token'] = null
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'Token': CustomFormControl(new FormControl(this['Token']), TokenauthNodeTokenResponse.propInfo['Token']),
            });
        }
        return this._formGroup;
    }

    setModelToBeFormGroupValues() {
        this.setValues(this.$formGroup.value, false);
    }

    setFormGroupValuesToBeModelValues() {
        if (this._formGroup) {
            this._formGroup.controls['Token'].setValue(this['Token']);
        }
    }
}

