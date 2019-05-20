import { UIRolePermissions } from '@sdk/v1/models/generated/UI-permissions-enum';
import { RoleGuardInput } from './components/shared/directives/roleGuard.directive';

export interface SideNavItem {
  label: string;
  icon: {
    cssClass: string,
    matIconName?: string,
  };
  children?: SideNavItem[];
  link?: string[];
  roleGuard?: RoleGuardInput | UIRolePermissions;
}

export const sideNavMenu: SideNavItem[] = [
  {
    label: 'Dashboard',
    icon: {
      cssClass: 'material-icons',
      matIconName: 'dashboard',
    },
    roleGuard: UIRolePermissions.metricsquery_read,
    link: ['/dashboard']
  },
  {
    label: 'Cluster',
    icon: {
      cssClass: 'app-l-side-nav-cluster',
    },
    roleGuard: {
      opt: [
        UIRolePermissions.clustercluster_read, UIRolePermissions.clustersmartnic_read
      ]
    },
    children: [
      {
        label: 'Cluster',
        icon: {
          cssClass: 'app-l-side-nav-cluster-detail'
        },
        roleGuard: UIRolePermissions.clustercluster_read,
        link: ['/cluster', 'cluster']
      },
      {
        label: 'NAPLES',
        icon: {
          cssClass: 'app-l-side-nav-cluster-naples'
        },
        roleGuard: UIRolePermissions.clustersmartnic_read,
        link: ['/cluster/', 'naples']
      },
      {
        label: 'Hosts',
        icon: {
          cssClass: 'app-l-side-nav-cluster-hosts'
        },
        roleGuard: UIRolePermissions.clusterhost_read,
        link: ['/cluster/', 'hosts']
      }
    ]
  },
  {
    label: 'Workload',
    icon: {
      cssClass: 'app-l-side-nav-workload',
    },
    roleGuard: UIRolePermissions.workloadworkload_read,
    link: ['/workload']
  },
  {
    label: 'Security',
    icon: {
      cssClass: 'app-l-side-nav-security',
    },
    roleGuard: {
      opt: [UIRolePermissions.securitysgpolicy_read, UIRolePermissions.securityapp_read]
    },
    children: [
      {
        label: 'SG Policies',
        icon: {
          cssClass: 'app-l-side-nav-security-securitypolicy'
        },
        roleGuard: UIRolePermissions.securitysgpolicy_read,
        link: ['/security', 'sgpolicies']
      },
      {
        label: 'Security Apps',
        icon: {
          cssClass: 'app-l-side-nav-security-apps'
        },
        roleGuard: UIRolePermissions.securityapp_read,
        link: ['/security', 'securityapps']
      }
    ]
  },
  {
    label: 'Monitoring',
    icon: {
      cssClass: 'app-l-side-nav-monitoring',
    },
    roleGuard: {
      opt: [
        UIRolePermissions.monitoringalert_read,
        UIRolePermissions.eventsevent_read,
        UIRolePermissions.auditevent_read,
        UIRolePermissions.fwlogsquery_read,
        UIRolePermissions.monitoringflowexportpolicy_read,
        UIRolePermissions.monitoringtechsupportrequest_read,
      ]
    },
    children: [
      {
        label: 'Alerts & Events',
        icon: {
          cssClass: 'material-icons',
          matIconName: 'event_available'
        },
        roleGuard: {
          opt: [
            UIRolePermissions.monitoringalert_read,
            UIRolePermissions.eventsevent_read,
          ]
        },
        link: ['/monitoring', 'alertsevents']
      },
      {
        label: 'Telemetry',
        icon: {
          cssClass: 'material-icons',
          matIconName: 'insert_chart_outlined'
        },
        roleGuard: UIRolePermissions.metricsquery_read,
        link: ['/monitoring', 'telemetry']
      },
      {
        label: 'Audit Events',
        icon: {
          cssClass: 'app-l-side-nav-monitoring-auditevents'
        },
        roleGuard: UIRolePermissions.auditevent_read,
        link: ['/monitoring', 'auditevents']
      },
      {
        label: 'Firewall Logs',
        icon: {
          cssClass: 'app-l-side-nav-monitoring-fwlogs'
        },
        roleGuard: UIRolePermissions.fwlogsquery_read,
        link: ['/monitoring', 'fwlogs']
      },
      {
        label: 'Flow Export',
        icon: {
          cssClass: 'app-l-side-nav-monitoring-flowexport'
        },
        roleGuard: UIRolePermissions.monitoringflowexportpolicy_read,
        link: ['/monitoring', 'flowexport']
      },
      {
        label: 'Tech Support',
        icon: {
          cssClass: 'app-l-side-nav-monitoring-techsupport'
        },
        roleGuard: UIRolePermissions.monitoringtechsupportrequest_read,
        link: ['/monitoring', 'techsupport']
      }
    ]
  },
];
