import type { MetricsViewSpecMeasureV2 } from "@rilldata/web-common/runtime-client";
import type { DashboardDataSources } from "./types";

export const allMeasures = ({
  metricsSpecQueryResult,
}: DashboardDataSources): MetricsViewSpecMeasureV2[] => {
  const measures = metricsSpecQueryResult.data?.measures;
  return measures === undefined ? [] : measures;
};

export const visibleMeasures = ({
  metricsSpecQueryResult,
  dashboard,
}: DashboardDataSources): MetricsViewSpecMeasureV2[] => {
  const measures = metricsSpecQueryResult.data?.measures?.filter(
    (d) => d.name && dashboard.visibleMeasureKeys.has(d.name)
  );
  return measures === undefined ? [] : measures;
};

export const measureLabel = ({
  metricsSpecQueryResult,
}: DashboardDataSources): ((m: string) => string) => {
  return (measureName) => {
    const measure = metricsSpecQueryResult.data?.measures?.find(
      (d) => d.name === measureName
    );
    return measure?.label ?? measureName;
  };
};

export const measureSelectors = {
  /**
   * Gets all visible measures in the dashboard.
   */
  visibleMeasures,

  /**
   * Get label for a measure by name
   */
  measureLabel,
};
