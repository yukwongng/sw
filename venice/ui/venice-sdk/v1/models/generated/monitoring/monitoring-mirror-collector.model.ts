/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, enumValidator } from './validators';
import { BaseModel, EnumDef } from './base-model';

import { MonitoringMirrorCollector_type,  MonitoringMirrorCollector_type_uihint  } from './enums';
import { ApiExportConfig, IApiExportConfig } from './api-export-config.model';

export interface IMonitoringMirrorCollector {
    'type'?: MonitoringMirrorCollector_type;
    'export-config'?: IApiExportConfig;
}


export class MonitoringMirrorCollector extends BaseModel implements IMonitoringMirrorCollector {
    'type': MonitoringMirrorCollector_type;
    'export-config': ApiExportConfig;
    public static enumProperties: { [key: string] : EnumDef } = {
        'type': {
            enum: MonitoringMirrorCollector_type_uihint,
            default: 'VENICE',
        },
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any) {
        super();
        this['export-config'] = new ApiExportConfig();
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
            this['type'] = values['type'];
            this['export-config'].setValues(values['export-config']);
        }
    }

    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'type': new FormControl(this['type'], [enumValidator(MonitoringMirrorCollector_type), ]),
                'export-config': this['export-config'].$formGroup,
            });
        }
        return this._formGroup;
    }

    setFormGroupValues() {
        if (this._formGroup) {
            this._formGroup.controls['type'].setValue(this['type']);
            this['export-config'].setFormGroupValues();
        }
    }
}

