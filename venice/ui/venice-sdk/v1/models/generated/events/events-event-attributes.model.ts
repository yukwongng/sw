/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from './base-model';

import { EventsEventAttributes_severity,  EventsEventAttributes_severity_uihint  } from './enums';
import { EventsEventAttributes_category,  } from './enums';
import { ApiObjectRef, IApiObjectRef } from './api-object-ref.model';
import { EventsEventSource, IEventsEventSource } from './events-event-source.model';

export interface IEventsEventAttributes {
    'severity': EventsEventAttributes_severity;
    'type'?: string;
    'message'?: string;
    'category': EventsEventAttributes_category;
    'object-ref'?: IApiObjectRef;
    'source'?: IEventsEventSource;
    'count'?: number;
}


export class EventsEventAttributes extends BaseModel implements IEventsEventAttributes {
    'severity': EventsEventAttributes_severity = null;
    'type': string = null;
    'message': string = null;
    'category': EventsEventAttributes_category = null;
    'object-ref': ApiObjectRef = null;
    'source': EventsEventSource = null;
    'count': number = null;
    public static propInfo: { [prop: string]: PropInfoItem } = {
        'severity': {
            enum: EventsEventAttributes_severity_uihint,
            default: 'INFO',
            required: true,
            type: 'string'
        },
        'type': {
            required: false,
            type: 'string'
        },
        'message': {
            required: false,
            type: 'string'
        },
        'category': {
            enum: EventsEventAttributes_category,
            default: 'Cluster',
            required: true,
            type: 'string'
        },
        'object-ref': {
            required: false,
            type: 'object'
        },
        'source': {
            required: false,
            type: 'object'
        },
        'count': {
            required: false,
            type: 'number'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return EventsEventAttributes.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return EventsEventAttributes.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (EventsEventAttributes.propInfo[prop] != null &&
                        EventsEventAttributes.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['object-ref'] = new ApiObjectRef();
        this['source'] = new EventsEventSource();
        this.setValues(values, setDefaults);
    }

    /**
     * set the values for both the Model and the Form Group. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any, fillDefaults = true): void {
        if (values && values['severity'] != null) {
            this['severity'] = values['severity'];
        } else if (fillDefaults && EventsEventAttributes.hasDefaultValue('severity')) {
            this['severity'] = <EventsEventAttributes_severity>  EventsEventAttributes.propInfo['severity'].default;
        } else {
            this['severity'] = null
        }
        if (values && values['type'] != null) {
            this['type'] = values['type'];
        } else if (fillDefaults && EventsEventAttributes.hasDefaultValue('type')) {
            this['type'] = EventsEventAttributes.propInfo['type'].default;
        } else {
            this['type'] = null
        }
        if (values && values['message'] != null) {
            this['message'] = values['message'];
        } else if (fillDefaults && EventsEventAttributes.hasDefaultValue('message')) {
            this['message'] = EventsEventAttributes.propInfo['message'].default;
        } else {
            this['message'] = null
        }
        if (values && values['category'] != null) {
            this['category'] = values['category'];
        } else if (fillDefaults && EventsEventAttributes.hasDefaultValue('category')) {
            this['category'] = <EventsEventAttributes_category>  EventsEventAttributes.propInfo['category'].default;
        } else {
            this['category'] = null
        }
        if (values) {
            this['object-ref'].setValues(values['object-ref'], fillDefaults);
        } else {
            this['object-ref'].setValues(null, fillDefaults);
        }
        if (values) {
            this['source'].setValues(values['source'], fillDefaults);
        } else {
            this['source'].setValues(null, fillDefaults);
        }
        if (values && values['count'] != null) {
            this['count'] = values['count'];
        } else if (fillDefaults && EventsEventAttributes.hasDefaultValue('count')) {
            this['count'] = EventsEventAttributes.propInfo['count'].default;
        } else {
            this['count'] = null
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'severity': CustomFormControl(new FormControl(this['severity'], [required, enumValidator(EventsEventAttributes_severity), ]), EventsEventAttributes.propInfo['severity']),
                'type': CustomFormControl(new FormControl(this['type']), EventsEventAttributes.propInfo['type']),
                'message': CustomFormControl(new FormControl(this['message']), EventsEventAttributes.propInfo['message']),
                'category': CustomFormControl(new FormControl(this['category'], [required, enumValidator(EventsEventAttributes_category), ]), EventsEventAttributes.propInfo['category']),
                'object-ref': CustomFormGroup(this['object-ref'].$formGroup, EventsEventAttributes.propInfo['object-ref'].required),
                'source': CustomFormGroup(this['source'].$formGroup, EventsEventAttributes.propInfo['source'].required),
                'count': CustomFormControl(new FormControl(this['count']), EventsEventAttributes.propInfo['count']),
            });
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('object-ref') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('object-ref').get(field);
                control.updateValueAndValidity();
            });
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('source') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('source').get(field);
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
            this._formGroup.controls['type'].setValue(this['type']);
            this._formGroup.controls['message'].setValue(this['message']);
            this._formGroup.controls['category'].setValue(this['category']);
            this['object-ref'].setFormGroupValuesToBeModelValues();
            this['source'].setFormGroupValuesToBeModelValues();
            this._formGroup.controls['count'].setValue(this['count']);
        }
    }
}

