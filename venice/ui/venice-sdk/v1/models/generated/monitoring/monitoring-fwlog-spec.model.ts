/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, enumValidator } from './validators';
import { BaseModel, EnumDef } from './base-model';

import { MonitoringFwlogSpec_filter,  MonitoringFwlogSpec_filter_uihint  } from './enums';
import { MonitoringFwlogExport, IMonitoringFwlogExport } from './monitoring-fwlog-export.model';

export interface IMonitoringFwlogSpec {
    'retention-time'?: string;
    'filter'?: Array<MonitoringFwlogSpec_filter>;
    'exports'?: Array<IMonitoringFwlogExport>;
}


export class MonitoringFwlogSpec extends BaseModel implements IMonitoringFwlogSpec {
    'retention-time': string;
    'filter': Array<MonitoringFwlogSpec_filter>;
    'exports': Array<MonitoringFwlogExport>;
    public static enumProperties: { [key: string] : EnumDef } = {
        'filter': {
            enum: MonitoringFwlogSpec_filter_uihint,
            default: 'FWLOG_ALL',
        },
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any) {
        super();
        this['filter'] = new Array<MonitoringFwlogSpec_filter>();
        this['exports'] = new Array<MonitoringFwlogExport>();
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
            this['retention-time'] = values['retention-time'];
            this.fillModelArray<MonitoringFwlogSpec_filter>(this, 'filter', values['filter']);
            this.fillModelArray<MonitoringFwlogExport>(this, 'exports', values['exports'], MonitoringFwlogExport);
        }
    }

    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'retention-time': new FormControl(this['retention-time']),
                'filter': new FormArray([]),
                'exports': new FormArray([]),
            });
            // generate FormArray control elements
            this.fillFormArray<MonitoringFwlogSpec_filter>('filter', this['filter']);
            // generate FormArray control elements
            this.fillFormArray<MonitoringFwlogExport>('exports', this['exports'], MonitoringFwlogExport);
        }
        return this._formGroup;
    }

    setFormGroupValues() {
        if (this._formGroup) {
            this._formGroup.controls['retention-time'].setValue(this['retention-time']);
            this.fillModelArray<MonitoringFwlogSpec_filter>(this, 'filter', this['filter']);
            this.fillModelArray<MonitoringFwlogExport>(this, 'exports', this['exports'], MonitoringFwlogExport);
        }
    }
}

