/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { DiagnosticsModuleStatus_category,  } from './enums';
import { DiagnosticsServicePort, IDiagnosticsServicePort } from './diagnostics-service-port.model';

export interface IDiagnosticsModuleStatus {
    'node'?: string;
    'module'?: string;
    'category': DiagnosticsModuleStatus_category;
    'last-start'?: Date;
    'restart-count'?: number;
    'last-restart-reason'?: string;
    'service'?: string;
    'service-ports'?: Array<IDiagnosticsServicePort>;
    'mac-address'?: string;
    '_ui'?: any;
}


export class DiagnosticsModuleStatus extends BaseModel implements IDiagnosticsModuleStatus {
    /** Field for holding arbitrary ui state */
    '_ui': any = {};
    /** Node on which this process is running. */
    'node': string = null;
    /** Module is the name of the process/container. */
    'module': string = null;
    /** Category specifies whether process is part of Venice(controller) or Naples(io) subsystem. */
    'category': DiagnosticsModuleStatus_category = null;
    /** Last start time. */
    'last-start': Date = null;
    /** Number of times process got restarted. zero if never restarted. */
    'restart-count': number = null;
    /** Arbitrary string, json, backtrace, etc. offering clues for restart. */
    'last-restart-reason': string = null;
    /** Service is the name of the service/pod this process is part of. */
    'service': string = null;
    /** Ports on which this process is listening. */
    'service-ports': Array<DiagnosticsServicePort> = null;
    /** MACAddress of the smart nic on which this module runs. */
    'mac-address': string = null;
    public static propInfo: { [prop in keyof IDiagnosticsModuleStatus]: PropInfoItem } = {
        'node': {
            description:  `Node on which this process is running.`,
            required: false,
            type: 'string'
        },
        'module': {
            description:  `Module is the name of the process/container.`,
            required: false,
            type: 'string'
        },
        'category': {
            enum: DiagnosticsModuleStatus_category,
            default: 'venice',
            description:  `Category specifies whether process is part of Venice(controller) or Naples(io) subsystem.`,
            required: true,
            type: 'string'
        },
        'last-start': {
            description:  `Last start time.`,
            required: false,
            type: 'Date'
        },
        'restart-count': {
            description:  `Number of times process got restarted. zero if never restarted.`,
            required: false,
            type: 'number'
        },
        'last-restart-reason': {
            description:  `Arbitrary string, json, backtrace, etc. offering clues for restart.`,
            required: false,
            type: 'string'
        },
        'service': {
            description:  `Service is the name of the service/pod this process is part of.`,
            required: false,
            type: 'string'
        },
        'service-ports': {
            description:  `Ports on which this process is listening.`,
            required: false,
            type: 'object'
        },
        'mac-address': {
            description:  `MACAddress of the smart nic on which this module runs.`,
            required: false,
            type: 'string'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return DiagnosticsModuleStatus.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return DiagnosticsModuleStatus.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (DiagnosticsModuleStatus.propInfo[prop] != null &&
                        DiagnosticsModuleStatus.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['service-ports'] = new Array<DiagnosticsServicePort>();
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
        if (values && values['node'] != null) {
            this['node'] = values['node'];
        } else if (fillDefaults && DiagnosticsModuleStatus.hasDefaultValue('node')) {
            this['node'] = DiagnosticsModuleStatus.propInfo['node'].default;
        } else {
            this['node'] = null
        }
        if (values && values['module'] != null) {
            this['module'] = values['module'];
        } else if (fillDefaults && DiagnosticsModuleStatus.hasDefaultValue('module')) {
            this['module'] = DiagnosticsModuleStatus.propInfo['module'].default;
        } else {
            this['module'] = null
        }
        if (values && values['category'] != null) {
            this['category'] = values['category'];
        } else if (fillDefaults && DiagnosticsModuleStatus.hasDefaultValue('category')) {
            this['category'] = <DiagnosticsModuleStatus_category>  DiagnosticsModuleStatus.propInfo['category'].default;
        } else {
            this['category'] = null
        }
        if (values && values['last-start'] != null) {
            this['last-start'] = values['last-start'];
        } else if (fillDefaults && DiagnosticsModuleStatus.hasDefaultValue('last-start')) {
            this['last-start'] = DiagnosticsModuleStatus.propInfo['last-start'].default;
        } else {
            this['last-start'] = null
        }
        if (values && values['restart-count'] != null) {
            this['restart-count'] = values['restart-count'];
        } else if (fillDefaults && DiagnosticsModuleStatus.hasDefaultValue('restart-count')) {
            this['restart-count'] = DiagnosticsModuleStatus.propInfo['restart-count'].default;
        } else {
            this['restart-count'] = null
        }
        if (values && values['last-restart-reason'] != null) {
            this['last-restart-reason'] = values['last-restart-reason'];
        } else if (fillDefaults && DiagnosticsModuleStatus.hasDefaultValue('last-restart-reason')) {
            this['last-restart-reason'] = DiagnosticsModuleStatus.propInfo['last-restart-reason'].default;
        } else {
            this['last-restart-reason'] = null
        }
        if (values && values['service'] != null) {
            this['service'] = values['service'];
        } else if (fillDefaults && DiagnosticsModuleStatus.hasDefaultValue('service')) {
            this['service'] = DiagnosticsModuleStatus.propInfo['service'].default;
        } else {
            this['service'] = null
        }
        if (values) {
            this.fillModelArray<DiagnosticsServicePort>(this, 'service-ports', values['service-ports'], DiagnosticsServicePort);
        } else {
            this['service-ports'] = [];
        }
        if (values && values['mac-address'] != null) {
            this['mac-address'] = values['mac-address'];
        } else if (fillDefaults && DiagnosticsModuleStatus.hasDefaultValue('mac-address')) {
            this['mac-address'] = DiagnosticsModuleStatus.propInfo['mac-address'].default;
        } else {
            this['mac-address'] = null
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'node': CustomFormControl(new FormControl(this['node']), DiagnosticsModuleStatus.propInfo['node']),
                'module': CustomFormControl(new FormControl(this['module']), DiagnosticsModuleStatus.propInfo['module']),
                'category': CustomFormControl(new FormControl(this['category'], [required, enumValidator(DiagnosticsModuleStatus_category), ]), DiagnosticsModuleStatus.propInfo['category']),
                'last-start': CustomFormControl(new FormControl(this['last-start']), DiagnosticsModuleStatus.propInfo['last-start']),
                'restart-count': CustomFormControl(new FormControl(this['restart-count']), DiagnosticsModuleStatus.propInfo['restart-count']),
                'last-restart-reason': CustomFormControl(new FormControl(this['last-restart-reason']), DiagnosticsModuleStatus.propInfo['last-restart-reason']),
                'service': CustomFormControl(new FormControl(this['service']), DiagnosticsModuleStatus.propInfo['service']),
                'service-ports': new FormArray([]),
                'mac-address': CustomFormControl(new FormControl(this['mac-address']), DiagnosticsModuleStatus.propInfo['mac-address']),
            });
            // generate FormArray control elements
            this.fillFormArray<DiagnosticsServicePort>('service-ports', this['service-ports'], DiagnosticsServicePort);
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('service-ports') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('service-ports').get(field);
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
            this._formGroup.controls['node'].setValue(this['node']);
            this._formGroup.controls['module'].setValue(this['module']);
            this._formGroup.controls['category'].setValue(this['category']);
            this._formGroup.controls['last-start'].setValue(this['last-start']);
            this._formGroup.controls['restart-count'].setValue(this['restart-count']);
            this._formGroup.controls['last-restart-reason'].setValue(this['last-restart-reason']);
            this._formGroup.controls['service'].setValue(this['service']);
            this.fillModelArray<DiagnosticsServicePort>(this, 'service-ports', this['service-ports'], DiagnosticsServicePort);
            this._formGroup.controls['mac-address'].setValue(this['mac-address']);
        }
    }
}

