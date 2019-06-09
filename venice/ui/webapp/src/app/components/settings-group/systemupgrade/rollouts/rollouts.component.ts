import {ChangeDetectorRef, Component, Input, ViewEncapsulation} from '@angular/core';
import {Animations} from '@app/animations';
import {Utility} from '@app/common/Utility';
import { TablevieweditAbstract} from '@app/components/shared/tableviewedit/tableviewedit.component';
import {ControllerService} from '@app/services/controller.service';
import {ObjstoreService} from '@app/services/generated/objstore.service';
import {RolloutService} from '@app/services/generated/rollout.service';
import {IApiStatus, IRolloutRollout, RolloutRollout, RolloutRolloutStatus_state} from '@sdk/v1/models/generated/rollout';
import { IObjstoreObjectList  } from '@sdk/v1/models/generated/objstore';
import {Observable} from 'rxjs';
import {Router} from '@angular/router';
import {HttpEventUtility} from '@common/HttpEventUtility';
import { ToolbarData } from '@app/models/frontend/shared/toolbar.interface';
import { TableCol } from '@app/components/shared/tableviewedit';

/**
 * This component let user manage Venice Rollouts.
 * User can run CRUD opeations.  As well, user can click on one rollout to watch rollout status.
 * User can also go to rollout images management UI in toolbar [IMAGES] button
 */

@Component({
  selector: 'app-rollouts',
  templateUrl: './rollouts.component.html',
  styleUrls: ['./rollouts.component.scss'],
  animations: [Animations],
  encapsulation: ViewEncapsulation.None
})
export class RolloutsComponent extends TablevieweditAbstract <IRolloutRollout, RolloutRollout> {
  dataObjects: ReadonlyArray<RolloutRollout>;
  pendingRollouts: RolloutRollout[] = [];
  pastRollouts: RolloutRollout[] = [];

  isTabComponent: boolean = true;
  disableTableWhenRowExpanded: boolean = true;
  selectedRollOut: RolloutRollout;
  rolloutImages: IObjstoreObjectList ;

  rolloutsEventUtility: HttpEventUtility<RolloutRollout>;
  bodyicon: any = {
    margin: {
      top: '9px',
      left: '8px'
    },
    matIcon: 'update'
  };

  cols: TableCol[] = [
    {field: 'meta.name', header: 'Name', class: '', sortable: true, width: 10},
    {field: 'meta.mod-time', header: 'Updated Time', class: '', sortable: true, width: 10},
    {field: 'meta.creation-time', header: 'Creation Time', class: '', sortable: true, width: 10},
    {field: 'spec.version', header: 'Version', class: '', sortable: true, width: 5},
    {
      field: 'spec.scheduled-start-time',
      header: 'Scheduled Start Time',
      class: '',
      sortable: true,
      width: 10
    },
    {field: 'spec.strategy', header: 'Strategy', class: '', sortable: true, width: 10},
    {field: 'spec.upgrade-type', header: 'Upgrade Type', class: '', sortable: true, width: 10},
    {field: 'status.prev-version', header: 'Prev Version', class: '', sortable: true, width: 15},
    {field: 'status.state', header: 'State', class: '', sortable: true, width: 20},
  ];

  exportFilename: string = 'Venice-rollouts';

  tabIndex: number = 0;

  constructor(protected controllerService: ControllerService,
              protected cdr: ChangeDetectorRef,
              protected rolloutService: RolloutService,
              protected objstoreService: ObjstoreService,
              protected router: Router) {
    super(controllerService, cdr);
  }

  getClassName(): string {
    return this.constructor.name;
  }


  postNgInit() {
    this.setDefaultToolbar();
    this.getRollouts();
    this.getRolloutImages();
  }

  getRollouts() {
    this.rolloutsEventUtility = new HttpEventUtility<RolloutRollout>(RolloutRollout, true);
    this.dataObjects = this.rolloutsEventUtility.array as ReadonlyArray<RolloutRollout>;
    const subscription = this.rolloutService.WatchRollout().subscribe(
      response => {
        this.rolloutsEventUtility.processEvents(response);
        this.setDefaultToolbar();
        this.splitRollouts();
      },
      this.controllerService.restErrorHandler('Failed to get Rollouts info')
    );
    this.subscriptions.push(subscription);
  }

  splitRollouts() {
    this.pendingRollouts.length = 0;
    this.pastRollouts.length = 0;
    for  (let i = 0; i < this.dataObjects.length; i++) {
      if (this.dataObjects[i].status.state === RolloutRolloutStatus_state.PROGRESSING) {
        this.pendingRollouts.push(this.dataObjects[i]);
      } else {
        this.pastRollouts.push(this.dataObjects[i]);
      }
    }
  }

  getRolloutImages() {
    const sub = this.objstoreService.ListObject(Utility.ROLLOUT_IMGAGE_NAMESPACE).subscribe(
      (response) => {
        this.rolloutImages = response.body as IObjstoreObjectList;
      },
      (error) => {
        this.controllerService.invokeRESTErrorToaster('Failed to fetch rollout images.', error);
      }
    );
    this.subscriptions.push(sub);
  }

  setDefaultToolbar() {
    this.controllerService.setToolbarData(this.buildToolbarData());
  }

  buildToolbarData(): ToolbarData {
    const toolbarData: ToolbarData = {
      breadcrumb: [{label: 'System Upgrade', url: Utility.getBaseUIUrl() + 'settings/upgrade'}],
      buttons: []
    };
    if (this.isToBuildCreateButton()) {
      toolbarData.buttons.push(
        {
          cssClass: 'global-button-primary rollouts-button',
          text: 'CREATE ROLLOUT',
          computeClass: () => this.shouldEnableButtons && this.tabIndex < 1  ? '' : 'global-button-disabled',
          callback: () => {
            this.createNewObject();
          }
        }
      );
    }
    toolbarData.buttons.push(
      {
        cssClass: 'global-button-primary rollouts-button',
        text: 'ROLLOUT IMAGES',
        computeClass: () => this.shouldEnableButtons ? '' : 'global-button-disabled',
        callback: () => {
          this.invokeUploadImageUI();
        }
      }
    );
    return toolbarData;
  }

  isToBuildCreateButton(): boolean {
    const isBuild = true;
    for  (let i = 0; i < this.dataObjects.length; i++) {
      if (this.dataObjects[i].status.state === RolloutRolloutStatus_state.PROGRESSING) {
        return false;
      }
    }
    return isBuild;
  }

  deleteRecord(object: RolloutRollout): Observable<{ body: IRolloutRollout | IApiStatus | Error, statusCode: number }> {
     return this.rolloutService.RemoveRollout(object);
  }

  postDeleteRecord () {
    this.splitRollouts();
  }

  generateDeleteConfirmMsg(object: IRolloutRollout) {
    return 'Are you sure you want to delete rollout ' + object.meta.name;
  }

  generateDeleteSuccessMsg(object: IRolloutRollout) {
    return 'Deleted rollout ' + object.meta.name;
  }

  invokeUploadImageUI() {
    this.router.navigateByUrl('settings/upgrade/imageuload');
  }

  // Row expansion toggle
  onRowClick(event, rowData) {
    if (this.selectedRollOut) {
      this.closeRowExpand();
      this.selectedRollOut = null;
    } else {
      this.selectedRollOut = rowData;
      this.expandRowRequest(event, rowData);
    }
  }

  displayRollOut(): string {
    return JSON.stringify(this.selectedRollOut, null, 1);
  }

  displayColumn(exportData, col): any {
    const fields = col.field.split('.');
    const value = Utility.getObjectValueByPropertyPath(exportData, fields);
    const column = col.field;
    switch (column) {
      default:
        return value;
    }
  }

  selectedIndexChangeEvent(event) {
    // console.log('tab' + event);
    this.tabIndex = event;
  }

  dump(obj): string {
    return JSON.stringify(obj, null, 1);
  }

  isToDisablePastTab(): boolean {
    return (this.creatingMode);
  }

}
