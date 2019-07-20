/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from './base-model';

import { ApiObjectMeta, IApiObjectMeta } from './api-object-meta.model';
import { EventsEvent_severity,  } from './enums';
import { EventsEvent_type,  } from './enums';
import { EventsEvent_category,  } from './enums';
import { ApiObjectRef, IApiObjectRef } from './api-object-ref.model';
import { EventsEventSource, IEventsEventSource } from './events-event-source.model';

export interface IEventsEvent {
    'kind'?: string;
    'api-version'?: string;
    'meta'?: IApiObjectMeta;
    'severity'?: EventsEvent_severity;
    'type'?: EventsEvent_type;
    'message'?: string;
    'category'?: EventsEvent_category;
    'object-ref'?: IApiObjectRef;
    'source'?: IEventsEventSource;
    'count'?: number;
}


export class EventsEvent extends BaseModel implements IEventsEvent {
    'kind': string = null;
    'api-version': string = null;
    'meta': ApiObjectMeta = null;
    'severity': EventsEvent_severity = null;
    'type': EventsEvent_type = null;
    'message': string = null;
    'category': EventsEvent_category = null;
    'object-ref': ApiObjectRef = null;
    'source': EventsEventSource = null;
    'count': number = null;
    public static propInfo: { [prop: string]: PropInfoItem } = {
        'kind': {
            required: false,
            type: 'string'
        },
        'api-version': {
            required: false,
            type: 'string'
        },
        'meta': {
            required: false,
            type: 'object'
        },
        'severity': {
            enum: EventsEvent_severity,
            default: 'info',
            required: false,
            type: 'string'
        },
        'type': {
            enum: EventsEvent_type,
            required: false,
            type: 'string'
        },
        'message': {
            required: false,
            type: 'string'
        },
        'category': {
            enum: EventsEvent_category,
            default: 'cluster',
            required: false,
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
        return EventsEvent.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return EventsEvent.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (EventsEvent.propInfo[prop] != null &&
                        EventsEvent.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['meta'] = new ApiObjectMeta();
        this['object-ref'] = new ApiObjectRef();
        this['source'] = new EventsEventSource();
        this.setValues(values, setDefaults);
    }

    /**
     * set the values for both the Model and the Form Group. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any, fillDefaults = true): void {
        if (values && values['kind'] != null) {
            this['kind'] = values['kind'];
        } else if (fillDefaults && EventsEvent.hasDefaultValue('kind')) {
            this['kind'] = EventsEvent.propInfo['kind'].default;
        } else {
            this['kind'] = null
        }
        if (values && values['api-version'] != null) {
            this['api-version'] = values['api-version'];
        } else if (fillDefaults && EventsEvent.hasDefaultValue('api-version')) {
            this['api-version'] = EventsEvent.propInfo['api-version'].default;
        } else {
            this['api-version'] = null
        }
        if (values) {
            this['meta'].setValues(values['meta'], fillDefaults);
        } else {
            this['meta'].setValues(null, fillDefaults);
        }
        if (values && values['severity'] != null) {
            this['severity'] = values['severity'];
        } else if (fillDefaults && EventsEvent.hasDefaultValue('severity')) {
            this['severity'] = <EventsEvent_severity>  EventsEvent.propInfo['severity'].default;
        } else {
            this['severity'] = null
        }
        if (values && values['type'] != null) {
            this['type'] = values['type'];
        } else if (fillDefaults && EventsEvent.hasDefaultValue('type')) {
            this['type'] = <EventsEvent_type>  EventsEvent.propInfo['type'].default;
        } else {
            this['type'] = null
        }
        if (values && values['message'] != null) {
            this['message'] = values['message'];
        } else if (fillDefaults && EventsEvent.hasDefaultValue('message')) {
            this['message'] = EventsEvent.propInfo['message'].default;
        } else {
            this['message'] = null
        }
        if (values && values['category'] != null) {
            this['category'] = values['category'];
        } else if (fillDefaults && EventsEvent.hasDefaultValue('category')) {
            this['category'] = <EventsEvent_category>  EventsEvent.propInfo['category'].default;
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
        } else if (fillDefaults && EventsEvent.hasDefaultValue('count')) {
            this['count'] = EventsEvent.propInfo['count'].default;
        } else {
            this['count'] = null
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'kind': CustomFormControl(new FormControl(this['kind']), EventsEvent.propInfo['kind']),
                'api-version': CustomFormControl(new FormControl(this['api-version']), EventsEvent.propInfo['api-version']),
                'meta': CustomFormGroup(this['meta'].$formGroup, EventsEvent.propInfo['meta'].required),
                'severity': CustomFormControl(new FormControl(this['severity'], [enumValidator(EventsEvent_severity), ]), EventsEvent.propInfo['severity']),
                'type': CustomFormControl(new FormControl(this['type'], [enumValidator(EventsEvent_type), ]), EventsEvent.propInfo['type']),
                'message': CustomFormControl(new FormControl(this['message']), EventsEvent.propInfo['message']),
                'category': CustomFormControl(new FormControl(this['category'], [enumValidator(EventsEvent_category), ]), EventsEvent.propInfo['category']),
                'object-ref': CustomFormGroup(this['object-ref'].$formGroup, EventsEvent.propInfo['object-ref'].required),
                'source': CustomFormGroup(this['source'].$formGroup, EventsEvent.propInfo['source'].required),
                'count': CustomFormControl(new FormControl(this['count']), EventsEvent.propInfo['count']),
            });
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('meta') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('meta').get(field);
                control.updateValueAndValidity();
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
            this._formGroup.controls['kind'].setValue(this['kind']);
            this._formGroup.controls['api-version'].setValue(this['api-version']);
            this['meta'].setFormGroupValuesToBeModelValues();
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

