import { MetricTransform, TransformQuery, TransformDataset } from './types';
import { SelectItem } from 'primeng/api';
import { MetricsMetadata } from '@sdk/metrics/generated/metadata';

/**
 * Populates the group by field tag in metric query and
 * transforms the dataset label to include node name
 */
export class GroupByTransform extends MetricTransform {
  transformName = 'GroupBy';
  _groupBy: string;
  groupByOptions: SelectItem[] = [];

  onFieldChange() {
    this.updateOptions();
  }

  get groupBy() {
    return this._groupBy;
  }

  set groupBy(value) {
    this._groupBy = value;
    this.requestMetrics();
  }

  updateOptions() {
    const options = MetricsMetadata[this.measurement].fields.filter(x => x.baseType === 'string').map( f => {
      return {label: f.displayName, value: f.name};
    });
    options.unshift({label: 'None', value: null});
    this.groupByOptions = options;
  }

  transformQuery(opts: TransformQuery) {
    opts.query['group-by-field'] = this.groupBy;
  }

  transformDataset(opts: TransformDataset) {
    if (opts.series.tags != null) {
      const groupByValue = opts.series.tags[this.groupBy];
      opts.dataset.label += ' - ' + groupByValue;
    }
  }
}
