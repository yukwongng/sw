/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { MonitoringAlertStatus_severity,  } from './enums';
import { MonitoringAlertSource, IMonitoringAlertSource } from './monitoring-alert-source.model';
import { ApiObjectRef, IApiObjectRef } from './api-object-ref.model';
import { MonitoringAlertReason, IMonitoringAlertReason } from './monitoring-alert-reason.model';
import { MonitoringAuditInfo, IMonitoringAuditInfo } from './monitoring-audit-info.model';

export interface IMonitoringAlertStatus {
    'severity': MonitoringAlertStatus_severity;
    'source'?: IMonitoringAlertSource;
    'event-uri'?: string;
    'object-ref'?: IApiObjectRef;
    'message'?: string;
    'reason'?: IMonitoringAlertReason;
    'acknowledged'?: IMonitoringAuditInfo;
    'resolved'?: IMonitoringAuditInfo;
}


export class MonitoringAlertStatus extends BaseModel implements IMonitoringAlertStatus {
    /** Severity of an alert. */
    'severity': MonitoringAlertStatus_severity = null;
    /** Alert source or origin. */
    'source': MonitoringAlertSource = null;
    /** Event that triggered the alert. */
    'event-uri': string = null;
    /** Affected object. */
    'object-ref': ApiObjectRef = null;
    /** Message from the alert rule that triggered the alert. */
    'message': string = null;
    /** Captures all the requirements from the alert policy rule with matched value. All these requirements must be cleared to auto-resolve an alert. */
    'reason': MonitoringAlertReason = null;
    /** Username and time at which the alert was acknowledged. */
    'acknowledged': MonitoringAuditInfo = null;
    /** Username and time at which the alert was resolved. */
    'resolved': MonitoringAuditInfo = null;
    public static propInfo: { [prop in keyof IMonitoringAlertStatus]: PropInfoItem } = {
        'severity': {
            enum: MonitoringAlertStatus_severity,
            default: 'info',
            description:  `Severity of an alert.`,
            required: true,
            type: 'string'
        },
        'source': {
            description:  `Alert source or origin.`,
            required: false,
            type: 'object'
        },
        'event-uri': {
            description:  `Event that triggered the alert.`,
            required: false,
            type: 'string'
        },
        'object-ref': {
            description:  `Affected object.`,
            required: false,
            type: 'object'
        },
        'message': {
            description:  `Message from the alert rule that triggered the alert.`,
            required: false,
            type: 'string'
        },
        'reason': {
            description:  `Captures all the requirements from the alert policy rule with matched value. All these requirements must be cleared to auto-resolve an alert.`,
            required: false,
            type: 'object'
        },
        'acknowledged': {
            description:  `Username and time at which the alert was acknowledged.`,
            required: false,
            type: 'object'
        },
        'resolved': {
            description:  `Username and time at which the alert was resolved.`,
            required: false,
            type: 'object'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return MonitoringAlertStatus.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return MonitoringAlertStatus.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (MonitoringAlertStatus.propInfo[prop] != null &&
                        MonitoringAlertStatus.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['source'] = new MonitoringAlertSource();
        this['object-ref'] = new ApiObjectRef();
        this['reason'] = new MonitoringAlertReason();
        this['acknowledged'] = new MonitoringAuditInfo();
        this['resolved'] = new MonitoringAuditInfo();
        this._inputValue = values;
        this.setValues(values, setDefaults);
    }

    /**
     * set the values for both the Model and the Form Group. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any, fillDefaults = true): void {
        if (values && values['severity'] != null) {
            this['severity'] = values['severity'];
        } else if (fillDefaults && MonitoringAlertStatus.hasDefaultValue('severity')) {
            this['severity'] = <MonitoringAlertStatus_severity>  MonitoringAlertStatus.propInfo['severity'].default;
        } else {
            this['severity'] = null
        }
        if (values) {
            this['source'].setValues(values['source'], fillDefaults);
        } else {
            this['source'].setValues(null, fillDefaults);
        }
        if (values && values['event-uri'] != null) {
            this['event-uri'] = values['event-uri'];
        } else if (fillDefaults && MonitoringAlertStatus.hasDefaultValue('event-uri')) {
            this['event-uri'] = MonitoringAlertStatus.propInfo['event-uri'].default;
        } else {
            this['event-uri'] = null
        }
        if (values) {
            this['object-ref'].setValues(values['object-ref'], fillDefaults);
        } else {
            this['object-ref'].setValues(null, fillDefaults);
        }
        if (values && values['message'] != null) {
            this['message'] = values['message'];
        } else if (fillDefaults && MonitoringAlertStatus.hasDefaultValue('message')) {
            this['message'] = MonitoringAlertStatus.propInfo['message'].default;
        } else {
            this['message'] = null
        }
        if (values) {
            this['reason'].setValues(values['reason'], fillDefaults);
        } else {
            this['reason'].setValues(null, fillDefaults);
        }
        if (values) {
            this['acknowledged'].setValues(values['acknowledged'], fillDefaults);
        } else {
            this['acknowledged'].setValues(null, fillDefaults);
        }
        if (values) {
            this['resolved'].setValues(values['resolved'], fillDefaults);
        } else {
            this['resolved'].setValues(null, fillDefaults);
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'severity': CustomFormControl(new FormControl(this['severity'], [required, enumValidator(MonitoringAlertStatus_severity), ]), MonitoringAlertStatus.propInfo['severity']),
                'source': CustomFormGroup(this['source'].$formGroup, MonitoringAlertStatus.propInfo['source'].required),
                'event-uri': CustomFormControl(new FormControl(this['event-uri']), MonitoringAlertStatus.propInfo['event-uri']),
                'object-ref': CustomFormGroup(this['object-ref'].$formGroup, MonitoringAlertStatus.propInfo['object-ref'].required),
                'message': CustomFormControl(new FormControl(this['message']), MonitoringAlertStatus.propInfo['message']),
                'reason': CustomFormGroup(this['reason'].$formGroup, MonitoringAlertStatus.propInfo['reason'].required),
                'acknowledged': CustomFormGroup(this['acknowledged'].$formGroup, MonitoringAlertStatus.propInfo['acknowledged'].required),
                'resolved': CustomFormGroup(this['resolved'].$formGroup, MonitoringAlertStatus.propInfo['resolved'].required),
            });
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('source') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('source').get(field);
                control.updateValueAndValidity();
            });
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('object-ref') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('object-ref').get(field);
                control.updateValueAndValidity();
            });
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('reason') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('reason').get(field);
                control.updateValueAndValidity();
            });
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('acknowledged') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('acknowledged').get(field);
                control.updateValueAndValidity();
            });
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('resolved') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('resolved').get(field);
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
            this._formGroup.controls['severity'].setValue(this['severity']);
            this['source'].setFormGroupValuesToBeModelValues();
            this._formGroup.controls['event-uri'].setValue(this['event-uri']);
            this['object-ref'].setFormGroupValuesToBeModelValues();
            this._formGroup.controls['message'].setValue(this['message']);
            this['reason'].setFormGroupValuesToBeModelValues();
            this['acknowledged'].setFormGroupValuesToBeModelValues();
            this['resolved'].setFormGroupValuesToBeModelValues();
        }
    }
}

