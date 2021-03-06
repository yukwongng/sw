import { Component, OnInit, Input, ViewChild, ChangeDetectorRef, SimpleChanges, OnChanges, OnDestroy, ViewEncapsulation } from '@angular/core';
import { IMonitoringAlert, MonitoringAlert, MonitoringAlertSpec_state, MonitoringAlertStatus_severity, MonitoringAlertSpec_state_uihint } from '@sdk/v1/models/generated/monitoring';
import { EventsEvent_severity, EventsEventAttributes_severity, IApiListWatchOptions, IEventsEvent, ApiListWatchOptions_sort_order } from '@sdk/v1/models/generated/events';
import { TimeRange } from '@app/components/shared/timerange/utility';
import { TableCol, CustomExportMap } from '@app/components/shared/tableviewedit';
import { Icon } from '@app/models/frontend/shared/icon.interface';
import { Utility } from '@app/common/Utility';
import { HttpEventUtility } from '@app/common/HttpEventUtility';
import { AlerttableService } from '@app/services/alerttable.service';
import { ControllerService } from '@app/services/controller.service';
import { MonitoringService } from '@app/services/generated/monitoring.service';
import { SearchService } from '@app/services/generated/search.service';
import { UIConfigsService } from '@app/services/uiconfigs.service';
import { Observable, forkJoin, throwError, Subscription } from 'rxjs';
import { UIRolePermissions } from '@sdk/v1/models/generated/UI-permissions-enum';
import { AlertsEventsSelector } from '@app/components/shared/alertsevents/alertsevents.component';
import { IApiStatus } from '@sdk/v1/models/generated/search';
import { Animations } from '@app/animations';
import { PrettyDatePipe } from '@app/components/shared/Pipes/PrettyDate.pipe';
import { TimeRangeComponent } from '../../timerange/timerange.component';
import { ValidationErrors } from '@angular/forms';
import { IStagingBulkEditAction, IBulkeditBulkEditItem } from '@sdk/v1/models/generated/staging';
import { PentableComponent } from '@app/components/shared/pentable/pentable.component';
import { DataComponent } from '@app/components/shared/datacomponent/datacomponent.component';
import { Eventtypes } from '@app/enum/eventtypes.enum';

@Component({
  selector: 'app-alertstable',
  templateUrl: './alertstable.component.html',
  styleUrls: ['./alertstable.component.scss'],
  animations: [Animations],
  encapsulation: ViewEncapsulation.None
})
export class AlertstableComponent extends DataComponent implements OnInit, OnChanges, OnDestroy {
  @ViewChild('timeRangeComponent') timeRangeComponent: TimeRangeComponent;
  @ViewChild('alertsTable') alertsTable: PentableComponent;

  isTabComponent: boolean;
  disableTableWhenRowExpanded: boolean;
  exportFilename: string = 'PSM-alerts';

  // If provided, will only show alerts and events
  // where the source node matches
  @Input() selector: AlertsEventsSelector;
  @Input() searchedAlert: string;
  @Input() showAlertsAdvSearch: boolean = false;

  // this property indicate if user is authorized to update alerts
  alertUpdatable: boolean = true;

  // holds a subset (possibly all) of this.alerts
  // This are the alerts that will be displayed
  dataObjects: ReadonlyArray<MonitoringAlert> = [];

  subscriptions: Subscription[] = [];
  alertSubscription: Subscription;

  // Alert State filters
  selectedStateFilters = [MonitoringAlertSpec_state_uihint.open];
  possibleFilterStates = Object.values(MonitoringAlertSpec_state_uihint);

  // The current alert severity filter, set to null if it is on All.
  currentAlertSeverityFilter: string;

  severityEnum: any = EventsEventAttributes_severity;

  cols: TableCol[] = [
    { field: 'meta.mod-time', header: 'Modification Time', class: 'alertsevents-column-date', sortable: true, width: 15 },
    { field: 'meta.creation-time', header: 'Creation Time', class: 'alertsevents-column-date', sortable: true, width: 15 },
    { field: 'status.severity', header: 'Severity', class: 'alertsevents-column-severity', sortable: false, width: 8 },
    { field: 'status.message', header: 'Message', class: 'alertsevents-column-message', sortable: false, width: 25 },
    { field: 'spec.state', header: 'State', class: 'alertsevents-column-state', sortable: false, width: 5 },
    { field: 'status.source', header: 'Source Node & Component', class: 'alerts-column-source', sortable: false },
  ];

  alertsIcon: Icon = {
    margin: {
      top: '0px',
      left: '0px',
    },
    matIcon: 'notifications'
  };

  // TimeRange for alerts
  alertsSelectedTimeRange: TimeRange;
  alertsTimeConstraints: string = '';

  // Will hold mapping from severity types to counts
  alertNumbers: { [severity in MonitoringAlertStatus_severity]: number } = {
    'info': 0,
    'warn': 0,
    'critical': 0,
  };

  // Holds all alerts
  alerts: ReadonlyArray<MonitoringAlert> = [];

  // Query params to send for watch
  alertQuery = {};

  // Used for processing watch stream events
  alertsEventUtility: HttpEventUtility<MonitoringAlert>;

  exportMap: CustomExportMap = {};

  selectedAlert: MonitoringAlert = null;

  constructor(protected _controllerService: ControllerService,
    protected _alerttableService: AlerttableService,
    protected uiconfigsService: UIConfigsService,
    protected searchService: SearchService,
    protected monitoringService: MonitoringService,
  ) {
    super(_controllerService, uiconfigsService);
  }

  ngOnInit() {
    this._controllerService.publish(Eventtypes.COMPONENT_INIT, {
      'component': this.getClassName(), 'state': Eventtypes.COMPONENT_INIT
    });
    this.alertUpdatable = this.uiconfigsService.isAuthorized(UIRolePermissions.monitoringalert_update);
    this.genQueryBodies();
    // If get alerts/events wasn't triggered by on change
    if (!this.alertSubscription) {
      this.getAlerts();
    }
    if (this.searchedAlert) {
      this.getSearchedAlert();
    }
  }

  getSelectedDataObjects() {
    return this.alertsTable.selectedDataObjects;
  }

  clearSelectedDataObjects() {
    this.alertsTable.selectedDataObjects = [];
  }

  clearAlertNumbersObject() {
    this.alertNumbers = {
      'info': 0,
      'warn': 0,
      'critical': 0,
    };
  }

  getSearchedAlert() {
    const searchedAlertSubscription = this.monitoringService.GetAlert(this.searchedAlert).subscribe(
      response => {
        this.selectedAlert = response.body as MonitoringAlert;
      },
      (error) => {
        // User may tweak browser url and make invalid alert name in the url, we will catch and throw error.
        this.selectedAlert = null;
        this.controllerService.invokeRESTErrorToaster('Failed to fetch alert ' + this.searchedAlert, error);
      }
    );
    this.subscriptions.push(searchedAlertSubscription);
  }

  closeDetails() {
    this.selectedAlert = null;
  }

  getAlerts() {
    this.dataObjects = [];
    this.alertsEventUtility = new HttpEventUtility<MonitoringAlert>(MonitoringAlert);
    this.alerts = this.alertsEventUtility.array;
    if (this.alertSubscription) {
      this.alertSubscription.unsubscribe();
    }
    this.alertSubscription = this.monitoringService.WatchAlert(this.alertQuery).subscribe(
      response => {
        this.alertsEventUtility.processEvents(response);
        // Reset counters
        Object.keys(MonitoringAlertStatus_severity).forEach(severity => {
          this.alertNumbers[severity] = 0;
        });
        this.alerts.forEach(alert => {
          this.alertNumbers[alert.status.severity] += 1;
        });
        this.filterAlerts();
      },
      this._controllerService.webSocketErrorHandler('Failed to get Alert')
    );
    this.subscriptions.push(this.alertSubscription);
  }

  /**
 * Submits an HTTP request to mark the alert as resolved
 * @param alert Alert to resolve
 */
  resolveAlert(alert: MonitoringAlert, $event) {
    const summary = 'Alert Resolved';
    const msg = 'Marked alert as resolved';
    this.updateAlertState(alert, MonitoringAlertSpec_state.resolved, summary, msg);
    $event.stopPropagation();
  }

  acknowledgeAlert(alert: MonitoringAlert, $event) {
    const summary = 'Alert Acknowledged';
    const msg = 'Marked alert as acknowledged';
    this.updateAlertState(alert, MonitoringAlertSpec_state.acknowledged, summary, msg);
    $event.stopPropagation();
  }

  openAlert(alert: MonitoringAlert, $event) {
    const summary = 'Alert Opened';
    const msg = 'Marked alert as open';
    this.updateAlertState(alert, MonitoringAlertSpec_state.open, summary, msg);
    $event.stopPropagation();
  }

  filterAlerts() {
    // We put the filtering into a set timeout so that it gets pushed to the end of
    // the micro task queue.
    // Otherwise, the table rendering of the items happens before the user's action on the checkbox
    // becomes visible. This allows the checkbox animation to happen immediately, and then we render
    // the new table.
    setTimeout(() => {
      this.dataObjects = this.alerts.filter((item) => {
        // Checking severity filter
        if (this.currentAlertSeverityFilter != null && item.status.severity !== this.currentAlertSeverityFilter) {
          return false;
        }
        // checking state filter
        return this.selectedStateFilters.includes(MonitoringAlertSpec_state_uihint[item.spec.state]);
      });
    }, 0);
  }

  /**
   * This api serves html template
   */
  showBatchResolveIcon(): boolean {
    // we want selected alerts all NOT in RESOLVED state
    return this.showBatchIconHelper(MonitoringAlertSpec_state.resolved, true);
  }

  /**
 * This api serves html template
 */
  showBatchAcknowLedgeIcon(): boolean {
    // we want selected alerts all NOT in ACKNOWLEDGED state
    return this.showBatchIconHelper(MonitoringAlertSpec_state.acknowledged, true);
  }

  /**
   * This api serves html template
   */
  showBatchOpenIcon(): boolean {
    // we want selected alerts all NOT in open state
    return this.showBatchIconHelper(MonitoringAlertSpec_state.open, true);
  }


  invokeUpdateAlerts(newState: MonitoringAlertSpec_state, summary: string, successMsg: string) {
   // this.updateAlertOneByOne(newState, summary, msg);
   const failureMsg: string = 'Failed to update alerts';
   this.updateAlertBulkEdit (newState, successMsg, failureMsg );
  }

  updateAlertOneByOne(newState: MonitoringAlertSpec_state, summary: string, msg: string) {
    const observables = this.buildObservableList(newState);
    this.updateAlertListForkJoin(observables, summary, msg);
  }

  updateAlertBulkEdit(newState: MonitoringAlertSpec_state, summary: string, msg: string) {
    const alerts = this.getSelectedDataObjects();
    const successMsg: string = 'Updated ' + alerts.length + ' alerts';
    const failureMsg: string = 'Failed to update alerts';
    const stagingBulkEditAction = this.buildBulkEditUpdateAlertsPayload(alerts, newState);
    this.bulkEditHelper(alerts, stagingBulkEditAction, successMsg, failureMsg);
  }

  onBulkEditSuccess() {
    this.invokeTimeRangeValidator();
  }

  buildBulkEditUpdateAlertsPayload(updateAlerts: MonitoringAlert[], newState: MonitoringAlertSpec_state, buffername: string = ''): IStagingBulkEditAction {

    const stagingBulkEditAction: IStagingBulkEditAction = Utility.buildStagingBulkEditAction(buffername);
    stagingBulkEditAction.spec.items = [];

    for (const alert of updateAlerts) {
      alert.spec.state = newState;
      const obj = {
        uri: '',
        method: 'update',
        object: alert.getModelValues()
      };
      stagingBulkEditAction.spec.items.push(obj as IBulkeditBulkEditItem);
    }
    return (stagingBulkEditAction.spec.items.length > 0) ? stagingBulkEditAction : null;
  }

  /**
 * This api serves html template
 */
  resolveSelectedAlerts() {
    const summary = 'Alerts Resolved';
    const msg = 'Marked selected alerts as resolved';
    const newState = MonitoringAlertSpec_state.resolved;
    this.invokeUpdateAlerts(newState, summary, msg);
  }

  /**
   * This api serves html template
   */
  acknowledgeSelectedAlerts() {
    const summary = 'Alerts Acknowledged';
    const msg = 'Marked selected alerts as acknowledged';
    const newState = MonitoringAlertSpec_state.acknowledged;
    this.invokeUpdateAlerts(newState, summary, msg);
  }

  /**
   * This api serves html template
   */
  openSelectedAlerts() {
    const summary = 'Alerts Opened';
    const msg = 'Marked selected alerts as open';
    const newState = MonitoringAlertSpec_state.open;
    this.invokeUpdateAlerts(newState, summary, msg);
  }

  /**
 * This API serves html template.
 * It will filter events displayed in table
 */
  onAlertNumberClick(severityType: string) {
    if (this.currentAlertSeverityFilter === severityType) {
      this.currentAlertSeverityFilter = null;
    } else {
      this.currentAlertSeverityFilter = severityType;
    }
    this.filterAlerts();
  }

  getAlertSourceNameLink(rowData: MonitoringAlert): string {
    return Utility.getAlertSourceLink(rowData);
  }

  getAlertSourceNameTooltip(rowData: MonitoringAlert): string {
    const cat = rowData.status['object-ref'].kind;
    return 'Go to ' + (cat ? cat.toLowerCase() + ' ' : '') + 'details page';
  }

  setAlertsTimeRange(timeRange: TimeRange) {
    // update query and call getEvents
    setTimeout(() => {
      this.alertsSelectedTimeRange = timeRange;
      const start = this.alertsSelectedTimeRange.getTime().startTime.toISOString() as any;
      const end = this.alertsSelectedTimeRange.getTime().endTime.toISOString() as any;
      this.alertsTimeConstraints = 'meta.mod-time<' + end + ',' + 'meta.mod-time>' + start;
      this.clearAlertNumbersObject();
      this.genQueryBodies();
      this.getAlerts();
    }, 0);
  }

  displayColumn(data: MonitoringAlert, col: TableCol): any {
    const fields = col.field.split('.');
    const value = Utility.getObjectValueByPropertyPath(data, fields);
    return Array.isArray(value) ? JSON.stringify(value, null, 2) : value;
  }

  /**
 * Submits an HTTP request to update the state of the alert
 * @param alert Alert to resolve
 */
  updateAlertState(alert: MonitoringAlert, newState: MonitoringAlertSpec_state, summary: string, msg: string) {
    // Create copy so that when we modify it doesn't change the view
    const observable = this.buildUpdateAlertStateObservable(alert, newState);
    const subscription = observable.subscribe(
      response => {
        this._controllerService.invokeSuccessToaster(summary, msg);
        this.invokeTimeRangeValidator(); // Gets new time and update table accordingly. Avoids using Stale Time Query
      },
      this._controllerService.restErrorHandler(summary + ' Failed')
    );
    this.subscriptions.push(subscription);
  }

  invokeTimeRangeValidator() {
    const timeValidate = this.timeRangeComponent.timeRangeGroupValidator();
    const err: ValidationErrors = timeValidate(this.timeRangeComponent.timeFormGroup);
    if (err !== null) {
      this._controllerService.invokeErrorToaster('Time Range Error', err.message);
    }
  }

  buildUpdateAlertStateObservable(alert: MonitoringAlert, newState: MonitoringAlertSpec_state): Observable<any> {
    const payload = new MonitoringAlert(alert);
    payload.spec.state = newState;
    const observable = this.monitoringService.UpdateAlert(payload.meta.name, payload);
    return observable;
  }

  /**
   * This is a helper function
   * @param state
   * @param reversed
   * this.showBatchIconHelper(MonitoringAlertSpec_state.RESOLVED, true);
   *      means that we want selected alerts all NOT in RESOLVED state
   * this.showBatchIconHelper(MonitoringAlertSpec_state.open, false);
   *      means that we want selected alerts all in open state
   */
  showBatchIconHelper(state: MonitoringAlertSpec_state, reversed: boolean = true): boolean {
    const alerts = this.getSelectedDataObjects();
    if (!this.alertUpdatable || !alerts || alerts.length === 0) {
      return false;
    }
    for (let i = 0; i < alerts.length; i++) {
      const alert = alerts[i];
      if (!reversed) {
        if (alert.spec.state !== state) {
          return false;
        }
      } else {
        if (alert.spec.state === state) {
          return false;
        }
      }
    }
    return true;
  }

  updateAlertListForkJoin(observables: Observable<any>[], summary: string, msg: string) {
    if (observables.length <= 0) {
      return;
    }
    forkJoin(observables).subscribe(
      (results) => {
        const isAllOK = Utility.isForkjoinResultAllOK(results);
        if (isAllOK) {
          this._controllerService.invokeSuccessToaster(summary, msg);
          this.clearSelectedDataObjects();
          this.invokeTimeRangeValidator();
        } else {
          const error = Utility.joinErrors(results);
          this._controllerService.invokeRESTErrorToaster(summary, error);
        }
      },
      this._controllerService.restErrorHandler(summary + ' Failed')

    );
  }

  buildObservableList(newState: MonitoringAlertSpec_state): Observable<any>[] {
    const observables = [];
    const alerts = this.getSelectedDataObjects();
    for (let i = 0; alerts && i < alerts.length; i++) {
      const observable = this.buildUpdateAlertStateObservable(alerts[i], newState);
      observables.push(observable);
    }
    return observables;
  }

  genQueryBodies() {
    const fieldSelectorOptions = [];

    if (this.selector != null) {
      fieldSelectorOptions.push(this.selector.alertSelector.selector);
    }
    if (this.alertsTimeConstraints.length) {
      fieldSelectorOptions.push(this.alertsTimeConstraints);
    }

    this.alertQuery['field-selector'] = fieldSelectorOptions.join(',');
  }

  selectAlert($event) {
    if (this.selectedAlert && $event.rowData.meta.name === this.selectedAlert.meta.name) {
      this.selectedAlert = null;
    } else {
      this.selectedAlert = $event.rowData;
    }
  }

  dateToString(date) {
    const prettyDate = new PrettyDatePipe('en-US');
    return prettyDate.transform(date);
  }

  getEventNameFromURI(uri) {
    const split_uri = uri.split('/');
    return split_uri[split_uri.length - 1];
  }

  ngOnChanges(change: SimpleChanges) {
    if (change.selector) {
      this.genQueryBodies();
      this.getAlerts();
    }
  }

  ngOnDestroy() {
    this.subscriptions.forEach(
      subscription => {
        subscription.unsubscribe();
      }
    );

    if (this.alertSubscription != null) {
      this.alertSubscription.unsubscribe();
    }
  }
}
