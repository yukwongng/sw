import { Component, DoCheck, Input, IterableDiffer, IterableDiffers, OnChanges, OnDestroy, OnInit, Output, SimpleChanges, ViewChild, ViewEncapsulation, EventEmitter, ChangeDetectorRef } from '@angular/core';
import { Animations } from '@app/animations';
import { HttpEventUtility } from '@app/common/HttpEventUtility';
import { Utility } from '@app/common/Utility';
import { Icon } from '@app/models/frontend/shared/icon.interface';
import { ControllerService } from '@app/services/controller.service';
import { MonitoringService } from '@app/services/generated/monitoring.service';
import { IMonitoringAlertDestination } from '@sdk/v1/models/generated/monitoring';
import { Table } from 'primeng/table';
import { TabcontentComponent } from 'web-app-framework';


@Component({
  selector: 'app-destinations',
  templateUrl: './destinations.component.html',
  styleUrls: ['./destinations.component.scss'],
  animations: [Animations],
  encapsulation: ViewEncapsulation.None
})
export class DestinationpolicyComponent extends TabcontentComponent implements OnInit, OnChanges, OnDestroy, DoCheck {
  @ViewChild('destinationsTable') destinationsTurboTable: Table;

  headerIcon: Icon = {
    margin: {
      top: '0px',
      left: '0px'
    },
    matIcon: 'send'
  };
  globalFilterFields: string[] = ['meta.name', 'spec.email-list'];

  destinations: IMonitoringAlertDestination[];
  selectedDestinationPolicy: IMonitoringAlertDestination;
  count: number;
  arrayDiffers: IterableDiffer<IMonitoringAlertDestination>;
  expandedRowData: IMonitoringAlertDestination;

  cols: any[] = [
    { field: 'meta.name', header: 'Policy Name', class: 'destinationpolicy-column-name', sortable: true },
    { field: 'spec.email-list', header: 'Email List', class: 'destinationpolicy-column-email-list', sortable: true },
    { field: 'spec.snmp-trap-servers', header: 'SNMP TRAP Servers', class: 'destinationpolicy-column-snmp_trap_servers', sortable: false },
    { field: 'status.total-notifications-sent', header: 'Total Notication Sent', class: 'destinationpolicy-column-total-notifications-sent', sortable: false },
  ];

  creatingMode: boolean = false;
  editingMode: boolean = false;

  // If we receive new data, but the display is frozen (user editing),
  // this should be set to true so that when user exits editing, we can update the display
  hasNewData: boolean = true;

  // Whether the toolbar buttons should be enabled
  shouldEnableButtons: boolean = true;

  @Input() destinationData: IMonitoringAlertDestination[];
  @Output() tableRowExpandClick: EventEmitter<any> = new EventEmitter();

  constructor(protected _controllerService: ControllerService,
    protected _iterableDiffers: IterableDiffers,
    private cdr: ChangeDetectorRef,
    protected _monitoringService: MonitoringService) {
    super();
    this.arrayDiffers = _iterableDiffers.find([]).create(HttpEventUtility.trackBy);
  }

  ngOnInit() {
    if (this.isActiveTab) {
      this.setDefaultToolbar();
    }
    this.setRowData();
  }

  getClassName(): string {
    return this.constructor.name;
  }

  setDefaultToolbar() {
    const currToolbar = this._controllerService.getToolbarData();
    currToolbar.buttons = [
      {
        cssClass: 'global-button-primary destinations-button',
        text: 'ADD DESTINATION',
        computeClass: () => { return this.shouldEnableButtons ? '' : 'global-button-disabled' },
        callback: () => { this.createNewDestination() }
      },
    ];
    this._controllerService.setToolbarData(currToolbar);
  }

  createNewDestination() {
    // If a row is expanded, we shouldnt be able to open a create new policy form
    if (!this.editingMode) {
      this.creatingMode = true;
      this.editMode.emit(true);
    }
  }

  creationFormClose() {
    this.creatingMode = false;
    this.editMode.emit(false)
  }

  setRowData() {
    /**
     * We copy the data so that the table doesn't
     * automatically update when the input binding is updated
     * This allows us to freeze the table when a user is doing inline
     * editing on a row entry
     */
    const _ = Utility.getLodash();
    const items = _.cloneDeep(this.destinationData);
    this.destinations = items;
    if (items != null) {
      this.count = items.length;
    } else {
      this.count = 0;
    }
  }

  /**
   * We check if any of the objects in the array have changed
   * This kind of detection is not automatically done by angular
   * To improve efficiency, we check only the name and last mod time
   * (see trackBy function) instead of checking every object field.
   */
  ngDoCheck() {
    const changes = this.arrayDiffers.diff(this.destinationData);
    if (changes) {
      if (this.editingMode) {
        this.hasNewData = true;
      } else {
        this.setRowData();
      }
    }
  }

  ngOnChanges(changes: SimpleChanges) {
    // We only set the toolbar if we are becoming the active tab,
    if (changes.isActiveTab != null && this.isActiveTab) {
      this.setDefaultToolbar();
    }
  }

  ngOnDestroy() {
  }

  displayColumn(alerteventpolicies, col): any {
    const fields = col.field.split('.');
    const value = Utility.getObjectValueByPropertyPath(alerteventpolicies, fields);
    const column = col.field;
    switch (column) {
      case 'spec.email-list':
        return JSON.stringify(value, null, 2);
      case 'spec.snmp-trap-servers':
        return JSON.stringify(value, null, 2);
      default:
        return Array.isArray(value) ? JSON.stringify(value, null, 2) : value;
    }
  }

  onUpdateRecord(event, destinationpolicy) {
    // If in creation mode, don't allow row expansion
    if (this.creatingMode) {
      return;
    }
    if (!this.editingMode) {
      this.destinationsTurboTable.toggleRow(destinationpolicy, event);
      this.expandedRowData = destinationpolicy;
      this.editMode.emit(true);
      this.editingMode = true;
      this.shouldEnableButtons = false;
    } else {
      this.editingMode = false;
      this.editMode.emit(false);
      this.shouldEnableButtons = true;
      // We don't untoggle the row here, it will happen when rowExpandAnimationComplete
      // is called.
    }
  }

  /**
   * Called when a row expand animation finishes
   * The animation happens when the row expands, and when it collapses
   * If it is expanding, then we are in ediitng mode (set in onUpdateRecord).
   * If it is collapsing, then editingMode should be false, (set in onUpdateRecord).
   * When it is collapsing, we toggle the row on the turbo table
   * 
   * This is because we must wait for the animation to complete before toggling
   * the row on the turbo table for a smooth animation.
   * @param  $event Angular animation end event
   */
  rowExpandAnimationComplete($event) {
    if (!this.editingMode) {
      // we are exiting the row expand
      this.destinationsTurboTable.toggleRow(this.expandedRowData, event);
      this.expandedRowData = null;
      if (this.hasNewData) {
        this.setRowData();
      }
      // Needed to prevent "ExpressionChangedAfterItHasBeenCheckedError"
      // We force an additional change detection cycle
      this.cdr.detectChanges();
    }
  }

  onDeleteRecord(event, destination: IMonitoringAlertDestination) {
    this._monitoringService.DeleteAlertDestination(destination.meta.name);
  }

}
