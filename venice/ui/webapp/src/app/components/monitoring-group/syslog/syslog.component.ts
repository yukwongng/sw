import { Component, OnInit, Input, ViewEncapsulation, ViewChild } from '@angular/core';
import { Animations } from '@app/animations';
import { SelectItem } from 'primeng/primeng';
import { Utility } from '@app/common/Utility';
import { MonitoringSyslogExport, MonitoringExternalCred, IMonitoringSyslogExport, MonitoringExportConfig, MonitoringSyslogExportConfig } from '@sdk/v1/models/generated/monitoring';
import { FormArray, FormGroup } from '@angular/forms';

@Component({
  selector: 'app-syslog',
  templateUrl: './syslog.component.html',
  styleUrls: ['./syslog.component.scss'],
  animations: [Animations],
  encapsulation: ViewEncapsulation.None
})
export class SyslogComponent implements OnInit {
  @Input() syslogExport: IMonitoringSyslogExport;
  @Input() showSyslogOptions: boolean = true;
  @Input() formatOptions: SelectItem[] = Utility.convertEnumToSelectItem(MonitoringSyslogExport.propInfo['format'].enum);
  syslogServerForm: FormGroup;

  syslogOverrideOptions: SelectItem[] = Utility.convertEnumToSelectItem(MonitoringSyslogExportConfig.propInfo['facility-override'].enum);

  syslogCredentialOptions: SelectItem[] = Utility.convertEnumToSelectItem(MonitoringExternalCred.propInfo['auth-type'].enum);

  targets: any;

  constructor() { }

  ngOnInit() {
    let syslogExport;
    if (this.syslogExport != null) {
      syslogExport = new MonitoringSyslogExport(this.syslogExport);
    } else {
      syslogExport = new MonitoringSyslogExport();
    }
    this.syslogServerForm = syslogExport.$formGroup;
    const targets: any = this.syslogServerForm.get('targets');
    if (targets.controls.length === 0) {
      this.addTarget();
    }
    this.targets = (<any>this.syslogServerForm.get(['targets'])).controls;
  }

  getValues(): IMonitoringSyslogExport {
    // remove syslog configs if they are off
    if (this.showSyslogOptions) {
      return this.syslogServerForm.value;
    } else {
      const _ = Utility.getLodash();
      // Cloning so that we don't change the form group object
      const formValues = _.cloneDeep(this.syslogServerForm.value);
      delete formValues.config;
      return formValues;
    }
  }

  getSelectedCredentialMethod(index: number): string {
    return this.syslogServerForm.get(['targets', index, 'credentials', 'auth-type']).value;
  }

  addTarget() {
    const targets = this.syslogServerForm.get('targets') as FormArray;
    targets.insert(0, new MonitoringExportConfig().$formGroup);
  }

  removeTarget(index) {
    const targets = this.syslogServerForm.get('targets') as FormArray;
    if (targets.length > 1) {
      targets.removeAt(index);
    }
  }

}
