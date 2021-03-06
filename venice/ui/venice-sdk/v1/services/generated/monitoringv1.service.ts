import { AbstractService, ServerEvent } from './abstract.service';
import { HttpClient } from '../../../../webapp/node_modules/@angular/common/http';
import { Observable } from '../../../../webapp/node_modules/rxjs';
import { Injectable } from '../../../../webapp/node_modules/@angular/core';
import { TrimDefaultsAndEmptyFields, TrimUIFields } from '../../../v1/utils/utility';

import { IMonitoringMirrorSessionList,MonitoringMirrorSessionList,IApiStatus,ApiStatus,IMonitoringMirrorSession,MonitoringMirrorSession,ApiLabel,IApiLabel,IMonitoringTroubleshootingSessionList,MonitoringTroubleshootingSessionList,IMonitoringTroubleshootingSession,MonitoringTroubleshootingSession,IMonitoringAlertDestinationList,MonitoringAlertDestinationList,IMonitoringAlertDestination,MonitoringAlertDestination,IMonitoringAlertPolicyList,MonitoringAlertPolicyList,IMonitoringAlertPolicy,MonitoringAlertPolicy,IMonitoringAlertList,MonitoringAlertList,IMonitoringAlert,MonitoringAlert,IMonitoringArchiveRequestList,MonitoringArchiveRequestList,IMonitoringArchiveRequest,MonitoringArchiveRequest,MonitoringCancelArchiveRequest,IMonitoringCancelArchiveRequest,IMonitoringAuditPolicy,MonitoringAuditPolicy,IMonitoringEventPolicyList,MonitoringEventPolicyList,IMonitoringEventPolicy,MonitoringEventPolicy,IMonitoringFlowExportPolicyList,MonitoringFlowExportPolicyList,IMonitoringFlowExportPolicy,MonitoringFlowExportPolicy,IMonitoringFwlogPolicyList,MonitoringFwlogPolicyList,IMonitoringFwlogPolicy,MonitoringFwlogPolicy,IMonitoringStatsAlertPolicyList,MonitoringStatsAlertPolicyList,IMonitoringStatsAlertPolicy,MonitoringStatsAlertPolicy,IMonitoringTechSupportRequestList,MonitoringTechSupportRequestList,IMonitoringTechSupportRequest,MonitoringTechSupportRequest,IMonitoringAutoMsgMirrorSessionWatchHelper,MonitoringAutoMsgMirrorSessionWatchHelper,IMonitoringAutoMsgAlertDestinationWatchHelper,MonitoringAutoMsgAlertDestinationWatchHelper,IMonitoringAutoMsgAlertPolicyWatchHelper,MonitoringAutoMsgAlertPolicyWatchHelper,IMonitoringAutoMsgAlertWatchHelper,MonitoringAutoMsgAlertWatchHelper,IMonitoringAutoMsgArchiveRequestWatchHelper,MonitoringAutoMsgArchiveRequestWatchHelper,IMonitoringAutoMsgAuditPolicyWatchHelper,MonitoringAutoMsgAuditPolicyWatchHelper,IMonitoringAutoMsgEventPolicyWatchHelper,MonitoringAutoMsgEventPolicyWatchHelper,IMonitoringAutoMsgFlowExportPolicyWatchHelper,MonitoringAutoMsgFlowExportPolicyWatchHelper,IMonitoringAutoMsgFwlogPolicyWatchHelper,MonitoringAutoMsgFwlogPolicyWatchHelper,IMonitoringAutoMsgStatsAlertPolicyWatchHelper,MonitoringAutoMsgStatsAlertPolicyWatchHelper,IMonitoringAutoMsgTechSupportRequestWatchHelper,MonitoringAutoMsgTechSupportRequestWatchHelper } from '../../models/generated/monitoring';

@Injectable()
export class Monitoringv1Service extends AbstractService {
  constructor(protected _http: HttpClient) {
    super(_http);
  }

  /**
   * Override super
   * Get the service class-name
  */
  getClassName(): string {
    return this.constructor.name;
  }

  /** List MirrorSession objects */
  public ListMirrorSession_1(queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringMirrorSessionList | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/MirrorSession';
    const opts = {
      eventID: 'ListMirrorSession_1',
      objType: 'MonitoringMirrorSessionList',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringMirrorSessionList | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Create MirrorSession object */
  public AddMirrorSession_1(body: IMonitoringMirrorSession, stagingID: string = "", trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringMirrorSession | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/MirrorSession';
    const opts = {
      eventID: 'AddMirrorSession_1',
      objType: 'MonitoringMirrorSession',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new MonitoringMirrorSession(body), null, trimDefaults)
    }
    return this.invokeAJAXPostCall(url, body, opts) as Observable<{body: IMonitoringMirrorSession | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Get MirrorSession object */
  public GetMirrorSession_1(O_Name, queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringMirrorSession | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/MirrorSession/{O.Name}';
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'GetMirrorSession_1',
      objType: 'MonitoringMirrorSession',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringMirrorSession | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Delete MirrorSession object */
  public DeleteMirrorSession_1(O_Name, stagingID: string = ""):Observable<{body: IMonitoringMirrorSession | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/MirrorSession/{O.Name}';
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'DeleteMirrorSession_1',
      objType: 'MonitoringMirrorSession',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXDeleteCall(url, opts) as Observable<{body: IMonitoringMirrorSession | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Update MirrorSession object */
  public UpdateMirrorSession_1(O_Name, body: IMonitoringMirrorSession, stagingID: string = "", previousVal: IMonitoringMirrorSession = null, trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringMirrorSession | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/MirrorSession/{O.Name}';
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'UpdateMirrorSession_1',
      objType: 'MonitoringMirrorSession',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new MonitoringMirrorSession(body), previousVal, trimDefaults)
    }
    return this.invokeAJAXPutCall(url, body, opts) as Observable<{body: IMonitoringMirrorSession | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Label MirrorSession object */
  public LabelMirrorSession_1(O_Name, body: IApiLabel, stagingID: string = "", trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringMirrorSession | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/MirrorSession/{O.Name}/label';
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'LabelMirrorSession_1',
      objType: 'MonitoringMirrorSession',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new ApiLabel(body), null, trimDefaults)
    }
    return this.invokeAJAXPostCall(url, body, opts) as Observable<{body: IMonitoringMirrorSession | IApiStatus | Error, statusCode: number}>;
  }
  
  /** List TroubleshootingSession objects */
  public ListTroubleshootingSession_1(queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringTroubleshootingSessionList | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/TroubleshootingSession';
    const opts = {
      eventID: 'ListTroubleshootingSession_1',
      objType: 'MonitoringTroubleshootingSessionList',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringTroubleshootingSessionList | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Create TroubleshootingSession object */
  public AddTroubleshootingSession_1(body: IMonitoringTroubleshootingSession, stagingID: string = "", trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringTroubleshootingSession | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/TroubleshootingSession';
    const opts = {
      eventID: 'AddTroubleshootingSession_1',
      objType: 'MonitoringTroubleshootingSession',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new MonitoringTroubleshootingSession(body), null, trimDefaults)
    }
    return this.invokeAJAXPostCall(url, body, opts) as Observable<{body: IMonitoringTroubleshootingSession | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Get TroubleshootingSession object */
  public GetTroubleshootingSession_1(O_Name, queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringTroubleshootingSession | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/TroubleshootingSession/{O.Name}';
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'GetTroubleshootingSession_1',
      objType: 'MonitoringTroubleshootingSession',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringTroubleshootingSession | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Delete TroubleshootingSession object */
  public DeleteTroubleshootingSession_1(O_Name, stagingID: string = ""):Observable<{body: IMonitoringTroubleshootingSession | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/TroubleshootingSession/{O.Name}';
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'DeleteTroubleshootingSession_1',
      objType: 'MonitoringTroubleshootingSession',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXDeleteCall(url, opts) as Observable<{body: IMonitoringTroubleshootingSession | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Update TroubleshootingSession object */
  public UpdateTroubleshootingSession_1(O_Name, body: IMonitoringTroubleshootingSession, stagingID: string = "", previousVal: IMonitoringTroubleshootingSession = null, trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringTroubleshootingSession | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/TroubleshootingSession/{O.Name}';
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'UpdateTroubleshootingSession_1',
      objType: 'MonitoringTroubleshootingSession',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new MonitoringTroubleshootingSession(body), previousVal, trimDefaults)
    }
    return this.invokeAJAXPutCall(url, body, opts) as Observable<{body: IMonitoringTroubleshootingSession | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Label TroubleshootingSession object */
  public LabelTroubleshootingSession_1(O_Name, body: IApiLabel, stagingID: string = "", trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringTroubleshootingSession | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/TroubleshootingSession/{O.Name}/label';
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'LabelTroubleshootingSession_1',
      objType: 'MonitoringTroubleshootingSession',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new ApiLabel(body), null, trimDefaults)
    }
    return this.invokeAJAXPostCall(url, body, opts) as Observable<{body: IMonitoringTroubleshootingSession | IApiStatus | Error, statusCode: number}>;
  }
  
  /** List AlertDestination objects */
  public ListAlertDestination_1(queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringAlertDestinationList | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/alertDestinations';
    const opts = {
      eventID: 'ListAlertDestination_1',
      objType: 'MonitoringAlertDestinationList',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringAlertDestinationList | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Create AlertDestination object */
  public AddAlertDestination_1(body: IMonitoringAlertDestination, stagingID: string = "", trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringAlertDestination | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/alertDestinations';
    const opts = {
      eventID: 'AddAlertDestination_1',
      objType: 'MonitoringAlertDestination',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new MonitoringAlertDestination(body), null, trimDefaults)
    }
    return this.invokeAJAXPostCall(url, body, opts) as Observable<{body: IMonitoringAlertDestination | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Get AlertDestination object */
  public GetAlertDestination_1(O_Name, queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringAlertDestination | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/alertDestinations/{O.Name}';
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'GetAlertDestination_1',
      objType: 'MonitoringAlertDestination',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringAlertDestination | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Delete AlertDestination object */
  public DeleteAlertDestination_1(O_Name, stagingID: string = ""):Observable<{body: IMonitoringAlertDestination | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/alertDestinations/{O.Name}';
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'DeleteAlertDestination_1',
      objType: 'MonitoringAlertDestination',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXDeleteCall(url, opts) as Observable<{body: IMonitoringAlertDestination | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Update AlertDestination object */
  public UpdateAlertDestination_1(O_Name, body: IMonitoringAlertDestination, stagingID: string = "", previousVal: IMonitoringAlertDestination = null, trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringAlertDestination | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/alertDestinations/{O.Name}';
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'UpdateAlertDestination_1',
      objType: 'MonitoringAlertDestination',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new MonitoringAlertDestination(body), previousVal, trimDefaults)
    }
    return this.invokeAJAXPutCall(url, body, opts) as Observable<{body: IMonitoringAlertDestination | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Label AlertDestination object */
  public LabelAlertDestination_1(O_Name, body: IApiLabel, stagingID: string = "", trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringAlertDestination | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/alertDestinations/{O.Name}/label';
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'LabelAlertDestination_1',
      objType: 'MonitoringAlertDestination',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new ApiLabel(body), null, trimDefaults)
    }
    return this.invokeAJAXPostCall(url, body, opts) as Observable<{body: IMonitoringAlertDestination | IApiStatus | Error, statusCode: number}>;
  }
  
  /** List AlertPolicy objects */
  public ListAlertPolicy_1(queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringAlertPolicyList | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/alertPolicies';
    const opts = {
      eventID: 'ListAlertPolicy_1',
      objType: 'MonitoringAlertPolicyList',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringAlertPolicyList | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Create AlertPolicy object */
  public AddAlertPolicy_1(body: IMonitoringAlertPolicy, stagingID: string = "", trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringAlertPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/alertPolicies';
    const opts = {
      eventID: 'AddAlertPolicy_1',
      objType: 'MonitoringAlertPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new MonitoringAlertPolicy(body), null, trimDefaults)
    }
    return this.invokeAJAXPostCall(url, body, opts) as Observable<{body: IMonitoringAlertPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Get AlertPolicy object */
  public GetAlertPolicy_1(O_Name, queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringAlertPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/alertPolicies/{O.Name}';
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'GetAlertPolicy_1',
      objType: 'MonitoringAlertPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringAlertPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Delete AlertPolicy object */
  public DeleteAlertPolicy_1(O_Name, stagingID: string = ""):Observable<{body: IMonitoringAlertPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/alertPolicies/{O.Name}';
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'DeleteAlertPolicy_1',
      objType: 'MonitoringAlertPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXDeleteCall(url, opts) as Observable<{body: IMonitoringAlertPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Update AlertPolicy object */
  public UpdateAlertPolicy_1(O_Name, body: IMonitoringAlertPolicy, stagingID: string = "", previousVal: IMonitoringAlertPolicy = null, trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringAlertPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/alertPolicies/{O.Name}';
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'UpdateAlertPolicy_1',
      objType: 'MonitoringAlertPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new MonitoringAlertPolicy(body), previousVal, trimDefaults)
    }
    return this.invokeAJAXPutCall(url, body, opts) as Observable<{body: IMonitoringAlertPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Label AlertPolicy object */
  public LabelAlertPolicy_1(O_Name, body: IApiLabel, stagingID: string = "", trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringAlertPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/alertPolicies/{O.Name}/label';
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'LabelAlertPolicy_1',
      objType: 'MonitoringAlertPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new ApiLabel(body), null, trimDefaults)
    }
    return this.invokeAJAXPostCall(url, body, opts) as Observable<{body: IMonitoringAlertPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** List Alert objects */
  public ListAlert_1(queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringAlertList | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/alerts';
    const opts = {
      eventID: 'ListAlert_1',
      objType: 'MonitoringAlertList',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringAlertList | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Get Alert object */
  public GetAlert_1(O_Name, queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringAlert | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/alerts/{O.Name}';
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'GetAlert_1',
      objType: 'MonitoringAlert',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringAlert | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Update Alert object */
  public UpdateAlert_1(O_Name, body: IMonitoringAlert, stagingID: string = "", previousVal: IMonitoringAlert = null, trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringAlert | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/alerts/{O.Name}';
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'UpdateAlert_1',
      objType: 'MonitoringAlert',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new MonitoringAlert(body), previousVal, trimDefaults)
    }
    return this.invokeAJAXPutCall(url, body, opts) as Observable<{body: IMonitoringAlert | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Label Alert object */
  public LabelAlert_1(O_Name, body: IApiLabel, stagingID: string = "", trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringAlert | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/alerts/{O.Name}/label';
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'LabelAlert_1',
      objType: 'MonitoringAlert',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new ApiLabel(body), null, trimDefaults)
    }
    return this.invokeAJAXPostCall(url, body, opts) as Observable<{body: IMonitoringAlert | IApiStatus | Error, statusCode: number}>;
  }
  
  /** List ArchiveRequest objects */
  public ListArchiveRequest_1(queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringArchiveRequestList | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/archive-requests';
    const opts = {
      eventID: 'ListArchiveRequest_1',
      objType: 'MonitoringArchiveRequestList',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringArchiveRequestList | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Create ArchiveRequest object */
  public AddArchiveRequest_1(body: IMonitoringArchiveRequest, stagingID: string = "", trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringArchiveRequest | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/archive-requests';
    const opts = {
      eventID: 'AddArchiveRequest_1',
      objType: 'MonitoringArchiveRequest',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new MonitoringArchiveRequest(body), null, trimDefaults)
    }
    return this.invokeAJAXPostCall(url, body, opts) as Observable<{body: IMonitoringArchiveRequest | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Get ArchiveRequest object */
  public GetArchiveRequest_1(O_Name, queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringArchiveRequest | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/archive-requests/{O.Name}';
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'GetArchiveRequest_1',
      objType: 'MonitoringArchiveRequest',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringArchiveRequest | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Delete ArchiveRequest object */
  public DeleteArchiveRequest_1(O_Name, stagingID: string = ""):Observable<{body: IMonitoringArchiveRequest | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/archive-requests/{O.Name}';
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'DeleteArchiveRequest_1',
      objType: 'MonitoringArchiveRequest',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXDeleteCall(url, opts) as Observable<{body: IMonitoringArchiveRequest | IApiStatus | Error, statusCode: number}>;
  }
  
  public Cancel_1(O_Name, body: IMonitoringCancelArchiveRequest, stagingID: string = "", trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringArchiveRequest | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/archive-requests/{O.Name}/Cancel';
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'Cancel_1',
      objType: 'MonitoringArchiveRequest',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new MonitoringCancelArchiveRequest(body), null, trimDefaults)
    }
    return this.invokeAJAXPostCall(url, body, opts) as Observable<{body: IMonitoringArchiveRequest | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Get AuditPolicy object */
  public GetAuditPolicy_1(queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringAuditPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/audit-policy';
    const opts = {
      eventID: 'GetAuditPolicy_1',
      objType: 'MonitoringAuditPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringAuditPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Delete AuditPolicy object */
  public DeleteAuditPolicy_1(stagingID: string = ""):Observable<{body: IMonitoringAuditPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/audit-policy';
    const opts = {
      eventID: 'DeleteAuditPolicy_1',
      objType: 'MonitoringAuditPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXDeleteCall(url, opts) as Observable<{body: IMonitoringAuditPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Create AuditPolicy object */
  public AddAuditPolicy_1(body: IMonitoringAuditPolicy, stagingID: string = "", trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringAuditPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/audit-policy';
    const opts = {
      eventID: 'AddAuditPolicy_1',
      objType: 'MonitoringAuditPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new MonitoringAuditPolicy(body), null, trimDefaults)
    }
    return this.invokeAJAXPostCall(url, body, opts) as Observable<{body: IMonitoringAuditPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Update AuditPolicy object */
  public UpdateAuditPolicy_1(body: IMonitoringAuditPolicy, stagingID: string = "", previousVal: IMonitoringAuditPolicy = null, trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringAuditPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/audit-policy';
    const opts = {
      eventID: 'UpdateAuditPolicy_1',
      objType: 'MonitoringAuditPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new MonitoringAuditPolicy(body), previousVal, trimDefaults)
    }
    return this.invokeAJAXPutCall(url, body, opts) as Observable<{body: IMonitoringAuditPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** List EventPolicy objects */
  public ListEventPolicy_1(queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringEventPolicyList | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/event-policy';
    const opts = {
      eventID: 'ListEventPolicy_1',
      objType: 'MonitoringEventPolicyList',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringEventPolicyList | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Create EventPolicy object */
  public AddEventPolicy_1(body: IMonitoringEventPolicy, stagingID: string = "", trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringEventPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/event-policy';
    const opts = {
      eventID: 'AddEventPolicy_1',
      objType: 'MonitoringEventPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new MonitoringEventPolicy(body), null, trimDefaults)
    }
    return this.invokeAJAXPostCall(url, body, opts) as Observable<{body: IMonitoringEventPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Get EventPolicy object */
  public GetEventPolicy_1(O_Name, queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringEventPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/event-policy/{O.Name}';
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'GetEventPolicy_1',
      objType: 'MonitoringEventPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringEventPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Delete EventPolicy object */
  public DeleteEventPolicy_1(O_Name, stagingID: string = ""):Observable<{body: IMonitoringEventPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/event-policy/{O.Name}';
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'DeleteEventPolicy_1',
      objType: 'MonitoringEventPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXDeleteCall(url, opts) as Observable<{body: IMonitoringEventPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Update EventPolicy object */
  public UpdateEventPolicy_1(O_Name, body: IMonitoringEventPolicy, stagingID: string = "", previousVal: IMonitoringEventPolicy = null, trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringEventPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/event-policy/{O.Name}';
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'UpdateEventPolicy_1',
      objType: 'MonitoringEventPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new MonitoringEventPolicy(body), previousVal, trimDefaults)
    }
    return this.invokeAJAXPutCall(url, body, opts) as Observable<{body: IMonitoringEventPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Label EventPolicy object */
  public LabelEventPolicy_1(O_Name, body: IApiLabel, stagingID: string = "", trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringEventPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/event-policy/{O.Name}/label';
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'LabelEventPolicy_1',
      objType: 'MonitoringEventPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new ApiLabel(body), null, trimDefaults)
    }
    return this.invokeAJAXPostCall(url, body, opts) as Observable<{body: IMonitoringEventPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** List FlowExportPolicy objects */
  public ListFlowExportPolicy_1(queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringFlowExportPolicyList | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/flowExportPolicy';
    const opts = {
      eventID: 'ListFlowExportPolicy_1',
      objType: 'MonitoringFlowExportPolicyList',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringFlowExportPolicyList | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Create FlowExportPolicy object */
  public AddFlowExportPolicy_1(body: IMonitoringFlowExportPolicy, stagingID: string = "", trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringFlowExportPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/flowExportPolicy';
    const opts = {
      eventID: 'AddFlowExportPolicy_1',
      objType: 'MonitoringFlowExportPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new MonitoringFlowExportPolicy(body), null, trimDefaults)
    }
    return this.invokeAJAXPostCall(url, body, opts) as Observable<{body: IMonitoringFlowExportPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Get FlowExportPolicy object */
  public GetFlowExportPolicy_1(O_Name, queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringFlowExportPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/flowExportPolicy/{O.Name}';
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'GetFlowExportPolicy_1',
      objType: 'MonitoringFlowExportPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringFlowExportPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Delete FlowExportPolicy object */
  public DeleteFlowExportPolicy_1(O_Name, stagingID: string = ""):Observable<{body: IMonitoringFlowExportPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/flowExportPolicy/{O.Name}';
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'DeleteFlowExportPolicy_1',
      objType: 'MonitoringFlowExportPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXDeleteCall(url, opts) as Observable<{body: IMonitoringFlowExportPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Update FlowExportPolicy object */
  public UpdateFlowExportPolicy_1(O_Name, body: IMonitoringFlowExportPolicy, stagingID: string = "", previousVal: IMonitoringFlowExportPolicy = null, trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringFlowExportPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/flowExportPolicy/{O.Name}';
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'UpdateFlowExportPolicy_1',
      objType: 'MonitoringFlowExportPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new MonitoringFlowExportPolicy(body), previousVal, trimDefaults)
    }
    return this.invokeAJAXPutCall(url, body, opts) as Observable<{body: IMonitoringFlowExportPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Label FlowExportPolicy object */
  public LabelFlowExportPolicy_1(O_Name, body: IApiLabel, stagingID: string = "", trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringFlowExportPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/flowExportPolicy/{O.Name}/label';
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'LabelFlowExportPolicy_1',
      objType: 'MonitoringFlowExportPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new ApiLabel(body), null, trimDefaults)
    }
    return this.invokeAJAXPostCall(url, body, opts) as Observable<{body: IMonitoringFlowExportPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** List FwlogPolicy objects */
  public ListFwlogPolicy_1(queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringFwlogPolicyList | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/fwlogPolicy';
    const opts = {
      eventID: 'ListFwlogPolicy_1',
      objType: 'MonitoringFwlogPolicyList',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringFwlogPolicyList | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Create FwlogPolicy object */
  public AddFwlogPolicy_1(body: IMonitoringFwlogPolicy, stagingID: string = "", trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringFwlogPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/fwlogPolicy';
    const opts = {
      eventID: 'AddFwlogPolicy_1',
      objType: 'MonitoringFwlogPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new MonitoringFwlogPolicy(body), null, trimDefaults)
    }
    return this.invokeAJAXPostCall(url, body, opts) as Observable<{body: IMonitoringFwlogPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Get FwlogPolicy object */
  public GetFwlogPolicy_1(O_Name, queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringFwlogPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/fwlogPolicy/{O.Name}';
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'GetFwlogPolicy_1',
      objType: 'MonitoringFwlogPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringFwlogPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Delete FwlogPolicy object */
  public DeleteFwlogPolicy_1(O_Name, stagingID: string = ""):Observable<{body: IMonitoringFwlogPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/fwlogPolicy/{O.Name}';
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'DeleteFwlogPolicy_1',
      objType: 'MonitoringFwlogPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXDeleteCall(url, opts) as Observable<{body: IMonitoringFwlogPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Update FwlogPolicy object */
  public UpdateFwlogPolicy_1(O_Name, body: IMonitoringFwlogPolicy, stagingID: string = "", previousVal: IMonitoringFwlogPolicy = null, trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringFwlogPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/fwlogPolicy/{O.Name}';
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'UpdateFwlogPolicy_1',
      objType: 'MonitoringFwlogPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new MonitoringFwlogPolicy(body), previousVal, trimDefaults)
    }
    return this.invokeAJAXPutCall(url, body, opts) as Observable<{body: IMonitoringFwlogPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Label FwlogPolicy object */
  public LabelFwlogPolicy_1(O_Name, body: IApiLabel, stagingID: string = "", trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringFwlogPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/fwlogPolicy/{O.Name}/label';
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'LabelFwlogPolicy_1',
      objType: 'MonitoringFwlogPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new ApiLabel(body), null, trimDefaults)
    }
    return this.invokeAJAXPostCall(url, body, opts) as Observable<{body: IMonitoringFwlogPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** List StatsAlertPolicy objects */
  public ListStatsAlertPolicy_1(queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringStatsAlertPolicyList | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/statsAlertPolicies';
    const opts = {
      eventID: 'ListStatsAlertPolicy_1',
      objType: 'MonitoringStatsAlertPolicyList',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringStatsAlertPolicyList | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Create StatsAlertPolicy object */
  public AddStatsAlertPolicy_1(body: IMonitoringStatsAlertPolicy, stagingID: string = "", trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringStatsAlertPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/statsAlertPolicies';
    const opts = {
      eventID: 'AddStatsAlertPolicy_1',
      objType: 'MonitoringStatsAlertPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new MonitoringStatsAlertPolicy(body), null, trimDefaults)
    }
    return this.invokeAJAXPostCall(url, body, opts) as Observable<{body: IMonitoringStatsAlertPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Get StatsAlertPolicy object */
  public GetStatsAlertPolicy_1(O_Name, queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringStatsAlertPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/statsAlertPolicies/{O.Name}';
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'GetStatsAlertPolicy_1',
      objType: 'MonitoringStatsAlertPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringStatsAlertPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Delete StatsAlertPolicy object */
  public DeleteStatsAlertPolicy_1(O_Name, stagingID: string = ""):Observable<{body: IMonitoringStatsAlertPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/statsAlertPolicies/{O.Name}';
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'DeleteStatsAlertPolicy_1',
      objType: 'MonitoringStatsAlertPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXDeleteCall(url, opts) as Observable<{body: IMonitoringStatsAlertPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** List TechSupportRequest objects */
  public ListTechSupportRequest(queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringTechSupportRequestList | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/techsupport';
    const opts = {
      eventID: 'ListTechSupportRequest',
      objType: 'MonitoringTechSupportRequestList',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringTechSupportRequestList | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Create TechSupportRequest object */
  public AddTechSupportRequest(body: IMonitoringTechSupportRequest, stagingID: string = "", trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringTechSupportRequest | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/techsupport';
    const opts = {
      eventID: 'AddTechSupportRequest',
      objType: 'MonitoringTechSupportRequest',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new MonitoringTechSupportRequest(body), null, trimDefaults)
    }
    return this.invokeAJAXPostCall(url, body, opts) as Observable<{body: IMonitoringTechSupportRequest | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Get TechSupportRequest object */
  public GetTechSupportRequest(O_Name, queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringTechSupportRequest | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/techsupport/{O.Name}';
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'GetTechSupportRequest',
      objType: 'MonitoringTechSupportRequest',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringTechSupportRequest | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Delete TechSupportRequest object */
  public DeleteTechSupportRequest(O_Name, stagingID: string = ""):Observable<{body: IMonitoringTechSupportRequest | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/techsupport/{O.Name}';
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'DeleteTechSupportRequest',
      objType: 'MonitoringTechSupportRequest',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXDeleteCall(url, opts) as Observable<{body: IMonitoringTechSupportRequest | IApiStatus | Error, statusCode: number}>;
  }
  
  /** List MirrorSession objects */
  public ListMirrorSession(queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringMirrorSessionList | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/MirrorSession';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    const opts = {
      eventID: 'ListMirrorSession',
      objType: 'MonitoringMirrorSessionList',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringMirrorSessionList | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Create MirrorSession object */
  public AddMirrorSession(body: IMonitoringMirrorSession, stagingID: string = "", trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringMirrorSession | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/MirrorSession';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    const opts = {
      eventID: 'AddMirrorSession',
      objType: 'MonitoringMirrorSession',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new MonitoringMirrorSession(body), null, trimDefaults)
    }
    return this.invokeAJAXPostCall(url, body, opts) as Observable<{body: IMonitoringMirrorSession | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Get MirrorSession object */
  public GetMirrorSession(O_Name, queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringMirrorSession | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/MirrorSession/{O.Name}';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'GetMirrorSession',
      objType: 'MonitoringMirrorSession',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringMirrorSession | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Delete MirrorSession object */
  public DeleteMirrorSession(O_Name, stagingID: string = ""):Observable<{body: IMonitoringMirrorSession | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/MirrorSession/{O.Name}';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'DeleteMirrorSession',
      objType: 'MonitoringMirrorSession',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXDeleteCall(url, opts) as Observable<{body: IMonitoringMirrorSession | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Update MirrorSession object */
  public UpdateMirrorSession(O_Name, body: IMonitoringMirrorSession, stagingID: string = "", previousVal: IMonitoringMirrorSession = null, trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringMirrorSession | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/MirrorSession/{O.Name}';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'UpdateMirrorSession',
      objType: 'MonitoringMirrorSession',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new MonitoringMirrorSession(body), previousVal, trimDefaults)
    }
    return this.invokeAJAXPutCall(url, body, opts) as Observable<{body: IMonitoringMirrorSession | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Label MirrorSession object */
  public LabelMirrorSession(O_Name, body: IApiLabel, stagingID: string = "", trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringMirrorSession | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/MirrorSession/{O.Name}/label';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'LabelMirrorSession',
      objType: 'MonitoringMirrorSession',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new ApiLabel(body), null, trimDefaults)
    }
    return this.invokeAJAXPostCall(url, body, opts) as Observable<{body: IMonitoringMirrorSession | IApiStatus | Error, statusCode: number}>;
  }
  
  /** List TroubleshootingSession objects */
  public ListTroubleshootingSession(queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringTroubleshootingSessionList | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/TroubleshootingSession';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    const opts = {
      eventID: 'ListTroubleshootingSession',
      objType: 'MonitoringTroubleshootingSessionList',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringTroubleshootingSessionList | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Create TroubleshootingSession object */
  public AddTroubleshootingSession(body: IMonitoringTroubleshootingSession, stagingID: string = "", trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringTroubleshootingSession | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/TroubleshootingSession';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    const opts = {
      eventID: 'AddTroubleshootingSession',
      objType: 'MonitoringTroubleshootingSession',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new MonitoringTroubleshootingSession(body), null, trimDefaults)
    }
    return this.invokeAJAXPostCall(url, body, opts) as Observable<{body: IMonitoringTroubleshootingSession | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Get TroubleshootingSession object */
  public GetTroubleshootingSession(O_Name, queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringTroubleshootingSession | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/TroubleshootingSession/{O.Name}';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'GetTroubleshootingSession',
      objType: 'MonitoringTroubleshootingSession',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringTroubleshootingSession | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Delete TroubleshootingSession object */
  public DeleteTroubleshootingSession(O_Name, stagingID: string = ""):Observable<{body: IMonitoringTroubleshootingSession | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/TroubleshootingSession/{O.Name}';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'DeleteTroubleshootingSession',
      objType: 'MonitoringTroubleshootingSession',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXDeleteCall(url, opts) as Observable<{body: IMonitoringTroubleshootingSession | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Update TroubleshootingSession object */
  public UpdateTroubleshootingSession(O_Name, body: IMonitoringTroubleshootingSession, stagingID: string = "", previousVal: IMonitoringTroubleshootingSession = null, trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringTroubleshootingSession | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/TroubleshootingSession/{O.Name}';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'UpdateTroubleshootingSession',
      objType: 'MonitoringTroubleshootingSession',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new MonitoringTroubleshootingSession(body), previousVal, trimDefaults)
    }
    return this.invokeAJAXPutCall(url, body, opts) as Observable<{body: IMonitoringTroubleshootingSession | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Label TroubleshootingSession object */
  public LabelTroubleshootingSession(O_Name, body: IApiLabel, stagingID: string = "", trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringTroubleshootingSession | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/TroubleshootingSession/{O.Name}/label';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'LabelTroubleshootingSession',
      objType: 'MonitoringTroubleshootingSession',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new ApiLabel(body), null, trimDefaults)
    }
    return this.invokeAJAXPostCall(url, body, opts) as Observable<{body: IMonitoringTroubleshootingSession | IApiStatus | Error, statusCode: number}>;
  }
  
  /** List AlertDestination objects */
  public ListAlertDestination(queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringAlertDestinationList | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/alertDestinations';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    const opts = {
      eventID: 'ListAlertDestination',
      objType: 'MonitoringAlertDestinationList',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringAlertDestinationList | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Create AlertDestination object */
  public AddAlertDestination(body: IMonitoringAlertDestination, stagingID: string = "", trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringAlertDestination | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/alertDestinations';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    const opts = {
      eventID: 'AddAlertDestination',
      objType: 'MonitoringAlertDestination',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new MonitoringAlertDestination(body), null, trimDefaults)
    }
    return this.invokeAJAXPostCall(url, body, opts) as Observable<{body: IMonitoringAlertDestination | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Get AlertDestination object */
  public GetAlertDestination(O_Name, queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringAlertDestination | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/alertDestinations/{O.Name}';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'GetAlertDestination',
      objType: 'MonitoringAlertDestination',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringAlertDestination | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Delete AlertDestination object */
  public DeleteAlertDestination(O_Name, stagingID: string = ""):Observable<{body: IMonitoringAlertDestination | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/alertDestinations/{O.Name}';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'DeleteAlertDestination',
      objType: 'MonitoringAlertDestination',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXDeleteCall(url, opts) as Observable<{body: IMonitoringAlertDestination | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Update AlertDestination object */
  public UpdateAlertDestination(O_Name, body: IMonitoringAlertDestination, stagingID: string = "", previousVal: IMonitoringAlertDestination = null, trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringAlertDestination | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/alertDestinations/{O.Name}';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'UpdateAlertDestination',
      objType: 'MonitoringAlertDestination',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new MonitoringAlertDestination(body), previousVal, trimDefaults)
    }
    return this.invokeAJAXPutCall(url, body, opts) as Observable<{body: IMonitoringAlertDestination | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Label AlertDestination object */
  public LabelAlertDestination(O_Name, body: IApiLabel, stagingID: string = "", trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringAlertDestination | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/alertDestinations/{O.Name}/label';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'LabelAlertDestination',
      objType: 'MonitoringAlertDestination',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new ApiLabel(body), null, trimDefaults)
    }
    return this.invokeAJAXPostCall(url, body, opts) as Observable<{body: IMonitoringAlertDestination | IApiStatus | Error, statusCode: number}>;
  }
  
  /** List AlertPolicy objects */
  public ListAlertPolicy(queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringAlertPolicyList | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/alertPolicies';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    const opts = {
      eventID: 'ListAlertPolicy',
      objType: 'MonitoringAlertPolicyList',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringAlertPolicyList | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Create AlertPolicy object */
  public AddAlertPolicy(body: IMonitoringAlertPolicy, stagingID: string = "", trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringAlertPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/alertPolicies';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    const opts = {
      eventID: 'AddAlertPolicy',
      objType: 'MonitoringAlertPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new MonitoringAlertPolicy(body), null, trimDefaults)
    }
    return this.invokeAJAXPostCall(url, body, opts) as Observable<{body: IMonitoringAlertPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Get AlertPolicy object */
  public GetAlertPolicy(O_Name, queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringAlertPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/alertPolicies/{O.Name}';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'GetAlertPolicy',
      objType: 'MonitoringAlertPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringAlertPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Delete AlertPolicy object */
  public DeleteAlertPolicy(O_Name, stagingID: string = ""):Observable<{body: IMonitoringAlertPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/alertPolicies/{O.Name}';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'DeleteAlertPolicy',
      objType: 'MonitoringAlertPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXDeleteCall(url, opts) as Observable<{body: IMonitoringAlertPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Update AlertPolicy object */
  public UpdateAlertPolicy(O_Name, body: IMonitoringAlertPolicy, stagingID: string = "", previousVal: IMonitoringAlertPolicy = null, trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringAlertPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/alertPolicies/{O.Name}';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'UpdateAlertPolicy',
      objType: 'MonitoringAlertPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new MonitoringAlertPolicy(body), previousVal, trimDefaults)
    }
    return this.invokeAJAXPutCall(url, body, opts) as Observable<{body: IMonitoringAlertPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Label AlertPolicy object */
  public LabelAlertPolicy(O_Name, body: IApiLabel, stagingID: string = "", trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringAlertPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/alertPolicies/{O.Name}/label';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'LabelAlertPolicy',
      objType: 'MonitoringAlertPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new ApiLabel(body), null, trimDefaults)
    }
    return this.invokeAJAXPostCall(url, body, opts) as Observable<{body: IMonitoringAlertPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** List Alert objects */
  public ListAlert(queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringAlertList | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/alerts';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    const opts = {
      eventID: 'ListAlert',
      objType: 'MonitoringAlertList',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringAlertList | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Get Alert object */
  public GetAlert(O_Name, queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringAlert | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/alerts/{O.Name}';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'GetAlert',
      objType: 'MonitoringAlert',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringAlert | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Update Alert object */
  public UpdateAlert(O_Name, body: IMonitoringAlert, stagingID: string = "", previousVal: IMonitoringAlert = null, trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringAlert | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/alerts/{O.Name}';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'UpdateAlert',
      objType: 'MonitoringAlert',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new MonitoringAlert(body), previousVal, trimDefaults)
    }
    return this.invokeAJAXPutCall(url, body, opts) as Observable<{body: IMonitoringAlert | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Label Alert object */
  public LabelAlert(O_Name, body: IApiLabel, stagingID: string = "", trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringAlert | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/alerts/{O.Name}/label';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'LabelAlert',
      objType: 'MonitoringAlert',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new ApiLabel(body), null, trimDefaults)
    }
    return this.invokeAJAXPostCall(url, body, opts) as Observable<{body: IMonitoringAlert | IApiStatus | Error, statusCode: number}>;
  }
  
  /** List ArchiveRequest objects */
  public ListArchiveRequest(queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringArchiveRequestList | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/archive-requests';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    const opts = {
      eventID: 'ListArchiveRequest',
      objType: 'MonitoringArchiveRequestList',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringArchiveRequestList | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Create ArchiveRequest object */
  public AddArchiveRequest(body: IMonitoringArchiveRequest, stagingID: string = "", trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringArchiveRequest | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/archive-requests';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    const opts = {
      eventID: 'AddArchiveRequest',
      objType: 'MonitoringArchiveRequest',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new MonitoringArchiveRequest(body), null, trimDefaults)
    }
    return this.invokeAJAXPostCall(url, body, opts) as Observable<{body: IMonitoringArchiveRequest | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Get ArchiveRequest object */
  public GetArchiveRequest(O_Name, queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringArchiveRequest | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/archive-requests/{O.Name}';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'GetArchiveRequest',
      objType: 'MonitoringArchiveRequest',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringArchiveRequest | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Delete ArchiveRequest object */
  public DeleteArchiveRequest(O_Name, stagingID: string = ""):Observable<{body: IMonitoringArchiveRequest | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/archive-requests/{O.Name}';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'DeleteArchiveRequest',
      objType: 'MonitoringArchiveRequest',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXDeleteCall(url, opts) as Observable<{body: IMonitoringArchiveRequest | IApiStatus | Error, statusCode: number}>;
  }
  
  public Cancel(O_Name, body: IMonitoringCancelArchiveRequest, stagingID: string = "", trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringArchiveRequest | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/archive-requests/{O.Name}/Cancel';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'Cancel',
      objType: 'MonitoringArchiveRequest',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new MonitoringCancelArchiveRequest(body), null, trimDefaults)
    }
    return this.invokeAJAXPostCall(url, body, opts) as Observable<{body: IMonitoringArchiveRequest | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Get AuditPolicy object */
  public GetAuditPolicy(queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringAuditPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/audit-policy';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    const opts = {
      eventID: 'GetAuditPolicy',
      objType: 'MonitoringAuditPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringAuditPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Delete AuditPolicy object */
  public DeleteAuditPolicy(stagingID: string = ""):Observable<{body: IMonitoringAuditPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/audit-policy';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    const opts = {
      eventID: 'DeleteAuditPolicy',
      objType: 'MonitoringAuditPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXDeleteCall(url, opts) as Observable<{body: IMonitoringAuditPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Create AuditPolicy object */
  public AddAuditPolicy(body: IMonitoringAuditPolicy, stagingID: string = "", trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringAuditPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/audit-policy';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    const opts = {
      eventID: 'AddAuditPolicy',
      objType: 'MonitoringAuditPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new MonitoringAuditPolicy(body), null, trimDefaults)
    }
    return this.invokeAJAXPostCall(url, body, opts) as Observable<{body: IMonitoringAuditPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Update AuditPolicy object */
  public UpdateAuditPolicy(body: IMonitoringAuditPolicy, stagingID: string = "", previousVal: IMonitoringAuditPolicy = null, trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringAuditPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/audit-policy';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    const opts = {
      eventID: 'UpdateAuditPolicy',
      objType: 'MonitoringAuditPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new MonitoringAuditPolicy(body), previousVal, trimDefaults)
    }
    return this.invokeAJAXPutCall(url, body, opts) as Observable<{body: IMonitoringAuditPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** List EventPolicy objects */
  public ListEventPolicy(queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringEventPolicyList | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/event-policy';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    const opts = {
      eventID: 'ListEventPolicy',
      objType: 'MonitoringEventPolicyList',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringEventPolicyList | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Create EventPolicy object */
  public AddEventPolicy(body: IMonitoringEventPolicy, stagingID: string = "", trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringEventPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/event-policy';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    const opts = {
      eventID: 'AddEventPolicy',
      objType: 'MonitoringEventPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new MonitoringEventPolicy(body), null, trimDefaults)
    }
    return this.invokeAJAXPostCall(url, body, opts) as Observable<{body: IMonitoringEventPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Get EventPolicy object */
  public GetEventPolicy(O_Name, queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringEventPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/event-policy/{O.Name}';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'GetEventPolicy',
      objType: 'MonitoringEventPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringEventPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Delete EventPolicy object */
  public DeleteEventPolicy(O_Name, stagingID: string = ""):Observable<{body: IMonitoringEventPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/event-policy/{O.Name}';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'DeleteEventPolicy',
      objType: 'MonitoringEventPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXDeleteCall(url, opts) as Observable<{body: IMonitoringEventPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Update EventPolicy object */
  public UpdateEventPolicy(O_Name, body: IMonitoringEventPolicy, stagingID: string = "", previousVal: IMonitoringEventPolicy = null, trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringEventPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/event-policy/{O.Name}';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'UpdateEventPolicy',
      objType: 'MonitoringEventPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new MonitoringEventPolicy(body), previousVal, trimDefaults)
    }
    return this.invokeAJAXPutCall(url, body, opts) as Observable<{body: IMonitoringEventPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Label EventPolicy object */
  public LabelEventPolicy(O_Name, body: IApiLabel, stagingID: string = "", trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringEventPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/event-policy/{O.Name}/label';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'LabelEventPolicy',
      objType: 'MonitoringEventPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new ApiLabel(body), null, trimDefaults)
    }
    return this.invokeAJAXPostCall(url, body, opts) as Observable<{body: IMonitoringEventPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** List FlowExportPolicy objects */
  public ListFlowExportPolicy(queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringFlowExportPolicyList | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/flowExportPolicy';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    const opts = {
      eventID: 'ListFlowExportPolicy',
      objType: 'MonitoringFlowExportPolicyList',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringFlowExportPolicyList | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Create FlowExportPolicy object */
  public AddFlowExportPolicy(body: IMonitoringFlowExportPolicy, stagingID: string = "", trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringFlowExportPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/flowExportPolicy';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    const opts = {
      eventID: 'AddFlowExportPolicy',
      objType: 'MonitoringFlowExportPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new MonitoringFlowExportPolicy(body), null, trimDefaults)
    }
    return this.invokeAJAXPostCall(url, body, opts) as Observable<{body: IMonitoringFlowExportPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Get FlowExportPolicy object */
  public GetFlowExportPolicy(O_Name, queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringFlowExportPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/flowExportPolicy/{O.Name}';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'GetFlowExportPolicy',
      objType: 'MonitoringFlowExportPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringFlowExportPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Delete FlowExportPolicy object */
  public DeleteFlowExportPolicy(O_Name, stagingID: string = ""):Observable<{body: IMonitoringFlowExportPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/flowExportPolicy/{O.Name}';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'DeleteFlowExportPolicy',
      objType: 'MonitoringFlowExportPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXDeleteCall(url, opts) as Observable<{body: IMonitoringFlowExportPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Update FlowExportPolicy object */
  public UpdateFlowExportPolicy(O_Name, body: IMonitoringFlowExportPolicy, stagingID: string = "", previousVal: IMonitoringFlowExportPolicy = null, trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringFlowExportPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/flowExportPolicy/{O.Name}';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'UpdateFlowExportPolicy',
      objType: 'MonitoringFlowExportPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new MonitoringFlowExportPolicy(body), previousVal, trimDefaults)
    }
    return this.invokeAJAXPutCall(url, body, opts) as Observable<{body: IMonitoringFlowExportPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Label FlowExportPolicy object */
  public LabelFlowExportPolicy(O_Name, body: IApiLabel, stagingID: string = "", trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringFlowExportPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/flowExportPolicy/{O.Name}/label';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'LabelFlowExportPolicy',
      objType: 'MonitoringFlowExportPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new ApiLabel(body), null, trimDefaults)
    }
    return this.invokeAJAXPostCall(url, body, opts) as Observable<{body: IMonitoringFlowExportPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** List FwlogPolicy objects */
  public ListFwlogPolicy(queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringFwlogPolicyList | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/fwlogPolicy';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    const opts = {
      eventID: 'ListFwlogPolicy',
      objType: 'MonitoringFwlogPolicyList',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringFwlogPolicyList | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Create FwlogPolicy object */
  public AddFwlogPolicy(body: IMonitoringFwlogPolicy, stagingID: string = "", trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringFwlogPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/fwlogPolicy';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    const opts = {
      eventID: 'AddFwlogPolicy',
      objType: 'MonitoringFwlogPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new MonitoringFwlogPolicy(body), null, trimDefaults)
    }
    return this.invokeAJAXPostCall(url, body, opts) as Observable<{body: IMonitoringFwlogPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Get FwlogPolicy object */
  public GetFwlogPolicy(O_Name, queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringFwlogPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/fwlogPolicy/{O.Name}';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'GetFwlogPolicy',
      objType: 'MonitoringFwlogPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringFwlogPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Delete FwlogPolicy object */
  public DeleteFwlogPolicy(O_Name, stagingID: string = ""):Observable<{body: IMonitoringFwlogPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/fwlogPolicy/{O.Name}';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'DeleteFwlogPolicy',
      objType: 'MonitoringFwlogPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXDeleteCall(url, opts) as Observable<{body: IMonitoringFwlogPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Update FwlogPolicy object */
  public UpdateFwlogPolicy(O_Name, body: IMonitoringFwlogPolicy, stagingID: string = "", previousVal: IMonitoringFwlogPolicy = null, trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringFwlogPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/fwlogPolicy/{O.Name}';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'UpdateFwlogPolicy',
      objType: 'MonitoringFwlogPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new MonitoringFwlogPolicy(body), previousVal, trimDefaults)
    }
    return this.invokeAJAXPutCall(url, body, opts) as Observable<{body: IMonitoringFwlogPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Label FwlogPolicy object */
  public LabelFwlogPolicy(O_Name, body: IApiLabel, stagingID: string = "", trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringFwlogPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/fwlogPolicy/{O.Name}/label';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'LabelFwlogPolicy',
      objType: 'MonitoringFwlogPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new ApiLabel(body), null, trimDefaults)
    }
    return this.invokeAJAXPostCall(url, body, opts) as Observable<{body: IMonitoringFwlogPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** List StatsAlertPolicy objects */
  public ListStatsAlertPolicy(queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringStatsAlertPolicyList | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/statsAlertPolicies';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    const opts = {
      eventID: 'ListStatsAlertPolicy',
      objType: 'MonitoringStatsAlertPolicyList',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringStatsAlertPolicyList | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Create StatsAlertPolicy object */
  public AddStatsAlertPolicy(body: IMonitoringStatsAlertPolicy, stagingID: string = "", trimObject: boolean = true, trimDefaults: boolean = true):Observable<{body: IMonitoringStatsAlertPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/statsAlertPolicies';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    const opts = {
      eventID: 'AddStatsAlertPolicy',
      objType: 'MonitoringStatsAlertPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    body = TrimUIFields(body)
    if (trimObject) {
      body = TrimDefaultsAndEmptyFields(body, new MonitoringStatsAlertPolicy(body), null, trimDefaults)
    }
    return this.invokeAJAXPostCall(url, body, opts) as Observable<{body: IMonitoringStatsAlertPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Get StatsAlertPolicy object */
  public GetStatsAlertPolicy(O_Name, queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringStatsAlertPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/statsAlertPolicies/{O.Name}';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'GetStatsAlertPolicy',
      objType: 'MonitoringStatsAlertPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringStatsAlertPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Delete StatsAlertPolicy object */
  public DeleteStatsAlertPolicy(O_Name, stagingID: string = ""):Observable<{body: IMonitoringStatsAlertPolicy | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/tenant/{O.Tenant}/statsAlertPolicies/{O.Name}';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    url = url.replace('{O.Name}', O_Name);
    const opts = {
      eventID: 'DeleteStatsAlertPolicy',
      objType: 'MonitoringStatsAlertPolicy',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXDeleteCall(url, opts) as Observable<{body: IMonitoringStatsAlertPolicy | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Watch MirrorSession objects. Supports WebSockets or HTTP long poll */
  public WatchMirrorSession_1(queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringAutoMsgMirrorSessionWatchHelper | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/watch/MirrorSession';
    const opts = {
      eventID: 'WatchMirrorSession_1',
      objType: 'MonitoringAutoMsgMirrorSessionWatchHelper',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringAutoMsgMirrorSessionWatchHelper | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Watch AlertDestination objects. Supports WebSockets or HTTP long poll */
  public WatchAlertDestination_1(queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringAutoMsgAlertDestinationWatchHelper | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/watch/alertDestinations';
    const opts = {
      eventID: 'WatchAlertDestination_1',
      objType: 'MonitoringAutoMsgAlertDestinationWatchHelper',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringAutoMsgAlertDestinationWatchHelper | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Watch AlertPolicy objects. Supports WebSockets or HTTP long poll */
  public WatchAlertPolicy_1(queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringAutoMsgAlertPolicyWatchHelper | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/watch/alertPolicies';
    const opts = {
      eventID: 'WatchAlertPolicy_1',
      objType: 'MonitoringAutoMsgAlertPolicyWatchHelper',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringAutoMsgAlertPolicyWatchHelper | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Watch Alert objects. Supports WebSockets or HTTP long poll */
  public WatchAlert_1(queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringAutoMsgAlertWatchHelper | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/watch/alerts';
    const opts = {
      eventID: 'WatchAlert_1',
      objType: 'MonitoringAutoMsgAlertWatchHelper',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringAutoMsgAlertWatchHelper | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Watch ArchiveRequest objects. Supports WebSockets or HTTP long poll */
  public WatchArchiveRequest_1(queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringAutoMsgArchiveRequestWatchHelper | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/watch/archive-requests';
    const opts = {
      eventID: 'WatchArchiveRequest_1',
      objType: 'MonitoringAutoMsgArchiveRequestWatchHelper',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringAutoMsgArchiveRequestWatchHelper | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Watch AuditPolicy objects. Supports WebSockets or HTTP long poll */
  public WatchAuditPolicy_1(queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringAutoMsgAuditPolicyWatchHelper | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/watch/audit-policy';
    const opts = {
      eventID: 'WatchAuditPolicy_1',
      objType: 'MonitoringAutoMsgAuditPolicyWatchHelper',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringAutoMsgAuditPolicyWatchHelper | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Watch EventPolicy objects. Supports WebSockets or HTTP long poll */
  public WatchEventPolicy_1(queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringAutoMsgEventPolicyWatchHelper | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/watch/event-policy';
    const opts = {
      eventID: 'WatchEventPolicy_1',
      objType: 'MonitoringAutoMsgEventPolicyWatchHelper',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringAutoMsgEventPolicyWatchHelper | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Watch FlowExportPolicy objects. Supports WebSockets or HTTP long poll */
  public WatchFlowExportPolicy_1(queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringAutoMsgFlowExportPolicyWatchHelper | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/watch/flowExportPolicy';
    const opts = {
      eventID: 'WatchFlowExportPolicy_1',
      objType: 'MonitoringAutoMsgFlowExportPolicyWatchHelper',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringAutoMsgFlowExportPolicyWatchHelper | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Watch FwlogPolicy objects. Supports WebSockets or HTTP long poll */
  public WatchFwlogPolicy_1(queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringAutoMsgFwlogPolicyWatchHelper | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/watch/fwlogPolicy';
    const opts = {
      eventID: 'WatchFwlogPolicy_1',
      objType: 'MonitoringAutoMsgFwlogPolicyWatchHelper',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringAutoMsgFwlogPolicyWatchHelper | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Watch StatsAlertPolicy objects. Supports WebSockets or HTTP long poll */
  public WatchStatsAlertPolicy_1(queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringAutoMsgStatsAlertPolicyWatchHelper | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/watch/statsAlertPolicies';
    const opts = {
      eventID: 'WatchStatsAlertPolicy_1',
      objType: 'MonitoringAutoMsgStatsAlertPolicyWatchHelper',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringAutoMsgStatsAlertPolicyWatchHelper | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Watch TechSupportRequest objects. Supports WebSockets or HTTP long poll */
  public WatchTechSupportRequest(queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringAutoMsgTechSupportRequestWatchHelper | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/watch/techsupport';
    const opts = {
      eventID: 'WatchTechSupportRequest',
      objType: 'MonitoringAutoMsgTechSupportRequestWatchHelper',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringAutoMsgTechSupportRequestWatchHelper | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Watch MirrorSession objects. Supports WebSockets or HTTP long poll */
  public WatchMirrorSession(queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringAutoMsgMirrorSessionWatchHelper | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/watch/tenant/{O.Tenant}/MirrorSession';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    const opts = {
      eventID: 'WatchMirrorSession',
      objType: 'MonitoringAutoMsgMirrorSessionWatchHelper',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringAutoMsgMirrorSessionWatchHelper | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Watch AlertDestination objects. Supports WebSockets or HTTP long poll */
  public WatchAlertDestination(queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringAutoMsgAlertDestinationWatchHelper | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/watch/tenant/{O.Tenant}/alertDestinations';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    const opts = {
      eventID: 'WatchAlertDestination',
      objType: 'MonitoringAutoMsgAlertDestinationWatchHelper',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringAutoMsgAlertDestinationWatchHelper | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Watch AlertPolicy objects. Supports WebSockets or HTTP long poll */
  public WatchAlertPolicy(queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringAutoMsgAlertPolicyWatchHelper | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/watch/tenant/{O.Tenant}/alertPolicies';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    const opts = {
      eventID: 'WatchAlertPolicy',
      objType: 'MonitoringAutoMsgAlertPolicyWatchHelper',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringAutoMsgAlertPolicyWatchHelper | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Watch Alert objects. Supports WebSockets or HTTP long poll */
  public WatchAlert(queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringAutoMsgAlertWatchHelper | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/watch/tenant/{O.Tenant}/alerts';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    const opts = {
      eventID: 'WatchAlert',
      objType: 'MonitoringAutoMsgAlertWatchHelper',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringAutoMsgAlertWatchHelper | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Watch ArchiveRequest objects. Supports WebSockets or HTTP long poll */
  public WatchArchiveRequest(queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringAutoMsgArchiveRequestWatchHelper | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/watch/tenant/{O.Tenant}/archive-requests';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    const opts = {
      eventID: 'WatchArchiveRequest',
      objType: 'MonitoringAutoMsgArchiveRequestWatchHelper',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringAutoMsgArchiveRequestWatchHelper | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Watch AuditPolicy objects. Supports WebSockets or HTTP long poll */
  public WatchAuditPolicy(queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringAutoMsgAuditPolicyWatchHelper | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/watch/tenant/{O.Tenant}/audit-policy';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    const opts = {
      eventID: 'WatchAuditPolicy',
      objType: 'MonitoringAutoMsgAuditPolicyWatchHelper',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringAutoMsgAuditPolicyWatchHelper | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Watch EventPolicy objects. Supports WebSockets or HTTP long poll */
  public WatchEventPolicy(queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringAutoMsgEventPolicyWatchHelper | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/watch/tenant/{O.Tenant}/event-policy';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    const opts = {
      eventID: 'WatchEventPolicy',
      objType: 'MonitoringAutoMsgEventPolicyWatchHelper',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringAutoMsgEventPolicyWatchHelper | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Watch FlowExportPolicy objects. Supports WebSockets or HTTP long poll */
  public WatchFlowExportPolicy(queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringAutoMsgFlowExportPolicyWatchHelper | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/watch/tenant/{O.Tenant}/flowExportPolicy';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    const opts = {
      eventID: 'WatchFlowExportPolicy',
      objType: 'MonitoringAutoMsgFlowExportPolicyWatchHelper',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringAutoMsgFlowExportPolicyWatchHelper | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Watch FwlogPolicy objects. Supports WebSockets or HTTP long poll */
  public WatchFwlogPolicy(queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringAutoMsgFwlogPolicyWatchHelper | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/watch/tenant/{O.Tenant}/fwlogPolicy';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    const opts = {
      eventID: 'WatchFwlogPolicy',
      objType: 'MonitoringAutoMsgFwlogPolicyWatchHelper',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringAutoMsgFwlogPolicyWatchHelper | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Watch StatsAlertPolicy objects. Supports WebSockets or HTTP long poll */
  public WatchStatsAlertPolicy(queryParam: any = null, stagingID: string = ""):Observable<{body: IMonitoringAutoMsgStatsAlertPolicyWatchHelper | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/monitoring/v1/watch/tenant/{O.Tenant}/statsAlertPolicies';
    url = url.replace('{O.Tenant}', this['O_Tenant']);
    const opts = {
      eventID: 'WatchStatsAlertPolicy',
      objType: 'MonitoringAutoMsgStatsAlertPolicyWatchHelper',
      isStaging: false,
    }
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
      opts.isStaging = true;
    }
    return this.invokeAJAXGetCall(url, queryParam, opts) as Observable<{body: IMonitoringAutoMsgStatsAlertPolicyWatchHelper | IApiStatus | Error, statusCode: number}>;
  }
  
  protected createListTechSupportRequestCache(): Observable<ServerEvent<MonitoringTechSupportRequest>> {
    return this.createDataCache<MonitoringTechSupportRequest>(MonitoringTechSupportRequest, `MonitoringTechSupportRequest`, () => this.ListTechSupportRequest(), (body: any) => this.WatchTechSupportRequest(body));
  }

  public ListTechSupportRequestCache(): Observable<ServerEvent<MonitoringTechSupportRequest>> {
    return this.getFromDataCache(`MonitoringTechSupportRequest`, () => { return this.createListTechSupportRequestCache() });
  }
  
  protected createListMirrorSessionCache(): Observable<ServerEvent<MonitoringMirrorSession>> {
    return this.createDataCache<MonitoringMirrorSession>(MonitoringMirrorSession, `MonitoringMirrorSession`, () => this.ListMirrorSession(), (body: any) => this.WatchMirrorSession(body));
  }

  public ListMirrorSessionCache(): Observable<ServerEvent<MonitoringMirrorSession>> {
    return this.getFromDataCache(`MonitoringMirrorSession`, () => { return this.createListMirrorSessionCache() });
  }
  
  protected createListAlertDestinationCache(): Observable<ServerEvent<MonitoringAlertDestination>> {
    return this.createDataCache<MonitoringAlertDestination>(MonitoringAlertDestination, `MonitoringAlertDestination`, () => this.ListAlertDestination(), (body: any) => this.WatchAlertDestination(body));
  }

  public ListAlertDestinationCache(): Observable<ServerEvent<MonitoringAlertDestination>> {
    return this.getFromDataCache(`MonitoringAlertDestination`, () => { return this.createListAlertDestinationCache() });
  }
  
  protected createListAlertPolicyCache(): Observable<ServerEvent<MonitoringAlertPolicy>> {
    return this.createDataCache<MonitoringAlertPolicy>(MonitoringAlertPolicy, `MonitoringAlertPolicy`, () => this.ListAlertPolicy(), (body: any) => this.WatchAlertPolicy(body));
  }

  public ListAlertPolicyCache(): Observable<ServerEvent<MonitoringAlertPolicy>> {
    return this.getFromDataCache(`MonitoringAlertPolicy`, () => { return this.createListAlertPolicyCache() });
  }
  
  protected createListAlertCache(): Observable<ServerEvent<MonitoringAlert>> {
    return this.createDataCache<MonitoringAlert>(MonitoringAlert, `MonitoringAlert`, () => this.ListAlert(), (body: any) => this.WatchAlert(body));
  }

  public ListAlertCache(): Observable<ServerEvent<MonitoringAlert>> {
    return this.getFromDataCache(`MonitoringAlert`, () => { return this.createListAlertCache() });
  }
  
  protected createListArchiveRequestCache(): Observable<ServerEvent<MonitoringArchiveRequest>> {
    return this.createDataCache<MonitoringArchiveRequest>(MonitoringArchiveRequest, `MonitoringArchiveRequest`, () => this.ListArchiveRequest(), (body: any) => this.WatchArchiveRequest(body));
  }

  public ListArchiveRequestCache(): Observable<ServerEvent<MonitoringArchiveRequest>> {
    return this.getFromDataCache(`MonitoringArchiveRequest`, () => { return this.createListArchiveRequestCache() });
  }
  
  protected createListEventPolicyCache(): Observable<ServerEvent<MonitoringEventPolicy>> {
    return this.createDataCache<MonitoringEventPolicy>(MonitoringEventPolicy, `MonitoringEventPolicy`, () => this.ListEventPolicy(), (body: any) => this.WatchEventPolicy(body));
  }

  public ListEventPolicyCache(): Observable<ServerEvent<MonitoringEventPolicy>> {
    return this.getFromDataCache(`MonitoringEventPolicy`, () => { return this.createListEventPolicyCache() });
  }
  
  protected createListFlowExportPolicyCache(): Observable<ServerEvent<MonitoringFlowExportPolicy>> {
    return this.createDataCache<MonitoringFlowExportPolicy>(MonitoringFlowExportPolicy, `MonitoringFlowExportPolicy`, () => this.ListFlowExportPolicy(), (body: any) => this.WatchFlowExportPolicy(body));
  }

  public ListFlowExportPolicyCache(): Observable<ServerEvent<MonitoringFlowExportPolicy>> {
    return this.getFromDataCache(`MonitoringFlowExportPolicy`, () => { return this.createListFlowExportPolicyCache() });
  }
  
  protected createListFwlogPolicyCache(): Observable<ServerEvent<MonitoringFwlogPolicy>> {
    return this.createDataCache<MonitoringFwlogPolicy>(MonitoringFwlogPolicy, `MonitoringFwlogPolicy`, () => this.ListFwlogPolicy(), (body: any) => this.WatchFwlogPolicy(body));
  }

  public ListFwlogPolicyCache(): Observable<ServerEvent<MonitoringFwlogPolicy>> {
    return this.getFromDataCache(`MonitoringFwlogPolicy`, () => { return this.createListFwlogPolicyCache() });
  }
  
  protected createListStatsAlertPolicyCache(): Observable<ServerEvent<MonitoringStatsAlertPolicy>> {
    return this.createDataCache<MonitoringStatsAlertPolicy>(MonitoringStatsAlertPolicy, `MonitoringStatsAlertPolicy`, () => this.ListStatsAlertPolicy(), (body: any) => this.WatchStatsAlertPolicy(body));
  }

  public ListStatsAlertPolicyCache(): Observable<ServerEvent<MonitoringStatsAlertPolicy>> {
    return this.getFromDataCache(`MonitoringStatsAlertPolicy`, () => { return this.createListStatsAlertPolicyCache() });
  }
  
}