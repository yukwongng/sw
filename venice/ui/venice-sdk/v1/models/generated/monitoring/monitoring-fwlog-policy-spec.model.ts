/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, enumValidator } from './validators';
import { BaseModel, PropInfoItem } from './base-model';

import { MonitoringExportConfig, IMonitoringExportConfig } from './monitoring-export-config.model';
import { MonitoringFwlogPolicySpec_format,  MonitoringFwlogPolicySpec_format_uihint  } from './enums';
import { MonitoringFwlogPolicySpec_filter,  MonitoringFwlogPolicySpec_filter_uihint  } from './enums';
import { MonitoringSyslogExportConfig, IMonitoringSyslogExportConfig } from './monitoring-syslog-export-config.model';

export interface IMonitoringFwlogPolicySpec {
    'targets'?: Array<IMonitoringExportConfig>;
    'format'?: MonitoringFwlogPolicySpec_format;
    'filter'?: Array<MonitoringFwlogPolicySpec_filter>;
    'config'?: IMonitoringSyslogExportConfig;
}


export class MonitoringFwlogPolicySpec extends BaseModel implements IMonitoringFwlogPolicySpec {
    'targets': Array<MonitoringExportConfig> = null;
    'format': MonitoringFwlogPolicySpec_format = null;
    'filter': Array<MonitoringFwlogPolicySpec_filter> = null;
    'config': MonitoringSyslogExportConfig = null;
    public static propInfo: { [prop: string]: PropInfoItem } = {
        'targets': {
            type: 'object'
        },
        'format': {
            enum: MonitoringFwlogPolicySpec_format_uihint,
            default: 'SYSLOG_BSD',
            type: 'string'
        },
        'filter': {
            enum: MonitoringFwlogPolicySpec_filter_uihint,
            default: 'FIREWALL_ACTION_ALL',
            type: 'Array<string>'
        },
        'config': {
            type: 'object'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return MonitoringFwlogPolicySpec.propInfo[propName];
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (MonitoringFwlogPolicySpec.propInfo[prop] != null &&
                        MonitoringFwlogPolicySpec.propInfo[prop].default != null &&
                        MonitoringFwlogPolicySpec.propInfo[prop].default != '');
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any) {
        super();
        this['targets'] = new Array<MonitoringExportConfig>();
        this['filter'] = new Array<MonitoringFwlogPolicySpec_filter>();
        this['config'] = new MonitoringSyslogExportConfig();
        this.setValues(values);
    }

    /**
     * set the values for both the Model and the Form Group. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any, fillDefaults = true): void {
        if (values) {
            this.fillModelArray<MonitoringExportConfig>(this, 'targets', values['targets'], MonitoringExportConfig);
        }
        if (values && values['format'] != null) {
            this['format'] = values['format'];
        } else if (fillDefaults && MonitoringFwlogPolicySpec.hasDefaultValue('format')) {
            this['format'] = <MonitoringFwlogPolicySpec_format>  MonitoringFwlogPolicySpec.propInfo['format'].default;
        }
        if (values && values['filter'] != null) {
            this['filter'] = values['filter'];
        } else if (fillDefaults && MonitoringFwlogPolicySpec.hasDefaultValue('filter')) {
            this['filter'] = [ MonitoringFwlogPolicySpec.propInfo['filter'].default];
        }
        if (values) {
            this['config'].setValues(values['config']);
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'targets': new FormArray([]),
                'format': new FormControl(this['format'], [enumValidator(MonitoringFwlogPolicySpec_format), ]),
                'filter': new FormControl(this['filter']),
                'config': this['config'].$formGroup,
            });
            // generate FormArray control elements
            this.fillFormArray<MonitoringExportConfig>('targets', this['targets'], MonitoringExportConfig);
        }
        return this._formGroup;
    }

    setModelToBeFormGroupValues() {
        this.setValues(this.$formGroup.value, false);
    }

    setFormGroupValuesToBeModelValues() {
        if (this._formGroup) {
            this.fillModelArray<MonitoringExportConfig>(this, 'targets', this['targets'], MonitoringExportConfig);
            this._formGroup.controls['format'].setValue(this['format']);
            this._formGroup.controls['filter'].setValue(this['filter']);
            this['config'].setFormGroupValuesToBeModelValues();
        }
    }
}

