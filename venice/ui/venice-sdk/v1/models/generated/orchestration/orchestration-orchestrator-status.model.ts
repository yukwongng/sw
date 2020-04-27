/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { OrchestrationOrchestratorStatus_connection_status,  OrchestrationOrchestratorStatus_connection_status_uihint  } from './enums';

export interface IOrchestrationOrchestratorStatus {
    'connection-status': OrchestrationOrchestratorStatus_connection_status;
    'last-transition-time'?: Date;
    'message'?: string;
    'orch-id'?: number;
    'discovered-namespaces'?: Array<string>;
    'incompatible-dscs'?: Array<string>;
    '_ui'?: any;
}


export class OrchestrationOrchestratorStatus extends BaseModel implements IOrchestrationOrchestratorStatus {
    /** Field for holding arbitrary ui state */
    '_ui': any = {};
    'connection-status': OrchestrationOrchestratorStatus_connection_status = null;
    'last-transition-time': Date = null;
    'message': string = null;
    'orch-id': number = null;
    'discovered-namespaces': Array<string> = null;
    'incompatible-dscs': Array<string> = null;
    public static propInfo: { [prop in keyof IOrchestrationOrchestratorStatus]: PropInfoItem } = {
        'connection-status': {
            enum: OrchestrationOrchestratorStatus_connection_status_uihint,
            default: 'unknown',
            required: true,
            type: 'string'
        },
        'last-transition-time': {
            required: false,
            type: 'Date'
        },
        'message': {
            required: false,
            type: 'string'
        },
        'orch-id': {
            required: false,
            type: 'number'
        },
        'discovered-namespaces': {
            required: false,
            type: 'Array<string>'
        },
        'incompatible-dscs': {
            required: false,
            type: 'Array<string>'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return OrchestrationOrchestratorStatus.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return OrchestrationOrchestratorStatus.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (OrchestrationOrchestratorStatus.propInfo[prop] != null &&
                        OrchestrationOrchestratorStatus.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['discovered-namespaces'] = new Array<string>();
        this['incompatible-dscs'] = new Array<string>();
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
        if (values && values['connection-status'] != null) {
            this['connection-status'] = values['connection-status'];
        } else if (fillDefaults && OrchestrationOrchestratorStatus.hasDefaultValue('connection-status')) {
            this['connection-status'] = <OrchestrationOrchestratorStatus_connection_status>  OrchestrationOrchestratorStatus.propInfo['connection-status'].default;
        } else {
            this['connection-status'] = null
        }
        if (values && values['last-transition-time'] != null) {
            this['last-transition-time'] = values['last-transition-time'];
        } else if (fillDefaults && OrchestrationOrchestratorStatus.hasDefaultValue('last-transition-time')) {
            this['last-transition-time'] = OrchestrationOrchestratorStatus.propInfo['last-transition-time'].default;
        } else {
            this['last-transition-time'] = null
        }
        if (values && values['message'] != null) {
            this['message'] = values['message'];
        } else if (fillDefaults && OrchestrationOrchestratorStatus.hasDefaultValue('message')) {
            this['message'] = OrchestrationOrchestratorStatus.propInfo['message'].default;
        } else {
            this['message'] = null
        }
        if (values && values['orch-id'] != null) {
            this['orch-id'] = values['orch-id'];
        } else if (fillDefaults && OrchestrationOrchestratorStatus.hasDefaultValue('orch-id')) {
            this['orch-id'] = OrchestrationOrchestratorStatus.propInfo['orch-id'].default;
        } else {
            this['orch-id'] = null
        }
        if (values && values['discovered-namespaces'] != null) {
            this['discovered-namespaces'] = values['discovered-namespaces'];
        } else if (fillDefaults && OrchestrationOrchestratorStatus.hasDefaultValue('discovered-namespaces')) {
            this['discovered-namespaces'] = [ OrchestrationOrchestratorStatus.propInfo['discovered-namespaces'].default];
        } else {
            this['discovered-namespaces'] = [];
        }
        if (values && values['incompatible-dscs'] != null) {
            this['incompatible-dscs'] = values['incompatible-dscs'];
        } else if (fillDefaults && OrchestrationOrchestratorStatus.hasDefaultValue('incompatible-dscs')) {
            this['incompatible-dscs'] = [ OrchestrationOrchestratorStatus.propInfo['incompatible-dscs'].default];
        } else {
            this['incompatible-dscs'] = [];
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'connection-status': CustomFormControl(new FormControl(this['connection-status'], [required, enumValidator(OrchestrationOrchestratorStatus_connection_status), ]), OrchestrationOrchestratorStatus.propInfo['connection-status']),
                'last-transition-time': CustomFormControl(new FormControl(this['last-transition-time']), OrchestrationOrchestratorStatus.propInfo['last-transition-time']),
                'message': CustomFormControl(new FormControl(this['message']), OrchestrationOrchestratorStatus.propInfo['message']),
                'orch-id': CustomFormControl(new FormControl(this['orch-id']), OrchestrationOrchestratorStatus.propInfo['orch-id']),
                'discovered-namespaces': CustomFormControl(new FormControl(this['discovered-namespaces']), OrchestrationOrchestratorStatus.propInfo['discovered-namespaces']),
                'incompatible-dscs': CustomFormControl(new FormControl(this['incompatible-dscs']), OrchestrationOrchestratorStatus.propInfo['incompatible-dscs']),
            });
        }
        return this._formGroup;
    }

    setModelToBeFormGroupValues() {
        this.setValues(this.$formGroup.value, false);
    }

    setFormGroupValuesToBeModelValues() {
        if (this._formGroup) {
            this._formGroup.controls['connection-status'].setValue(this['connection-status']);
            this._formGroup.controls['last-transition-time'].setValue(this['last-transition-time']);
            this._formGroup.controls['message'].setValue(this['message']);
            this._formGroup.controls['orch-id'].setValue(this['orch-id']);
            this._formGroup.controls['discovered-namespaces'].setValue(this['discovered-namespaces']);
            this._formGroup.controls['incompatible-dscs'].setValue(this['incompatible-dscs']);
        }
    }
}

