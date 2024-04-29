import { COMPARIONS_COLORS } from "@rilldata/web-common/features/dashboards/config";
import { MainLineColor } from "@rilldata/web-common/features/dashboards/time-series/chart-colors";
import { TIME_GRAIN } from "@rilldata/web-common/lib/time/config";
import { V1TimeGrain } from "@rilldata/web-common/runtime-client";
import { TDDAlternateCharts, TDDChart } from "../types";

export function escapeHTML(value: any): string {
  return String(value).replace(/&/g, "&amp;").replace(/</g, "&lt;");
}

export function tddTooltipFormatter(
  chartType: TDDAlternateCharts,
  measureLabel: string,
  dimensionLabel: string | undefined,
  selectedDimensionValues: (string | null)[],
  interval: V1TimeGrain | undefined,
) {
  const colorMap: Record<string, string> = {};
  selectedDimensionValues.forEach((dimValue, i) => {
    colorMap[String(dimValue)] = COMPARIONS_COLORS[i];
  });
  colorMap[measureLabel] = MainLineColor;

  return (value: Record<string, any>) => {
    let content = "";
    const { Time, ...rest } = value;

    if (Time) {
      const formattedTime = interval
        ? new Date(Time).toLocaleDateString(
            undefined,
            TIME_GRAIN[interval].formatDate,
          )
        : Time.toString();
      content += `<h2>${formattedTime}</h2>`;
    }

    const keys = Object.keys(rest);

    if (keys.length > 0) {
      content += "<table>";

      const selectedValuesLength = selectedDimensionValues.length;
      if (chartType === TDDChart.STACKED_AREA && selectedValuesLength) {
        for (const key of keys) {
          const val = rest[key];

          if (val === null || val === "NaN") continue;

          content += `
            <tr>
              <td class="color">
                <svg width="16" height="16">
                  <circle cx="8" cy="8" r="6" style="fill:${colorMap[String(key)] || "#000"};">
                  </reccit>
                </svg>
              </td>
              <td class="key">${escapeHTML(key)}</td>
              <td class="value">${escapeHTML(String(val))}</td>
            </tr>`;
        }
      } else {
        const key =
          selectedValuesLength && dimensionLabel
            ? rest[dimensionLabel]
            : measureLabel;
        const val = rest[measureLabel];

        content += `
          <tr>
          <td class="color">
            <svg width="16" height="16">
              <circle cx="8" cy="8" r="6" style="fill:${colorMap[String(key)] || "#000"};">
              </reccit>
            </svg>
          </td>
            <td class="key">${escapeHTML(key)}</td>
            <td class="value">${escapeHTML(String(val))}</td>
          </tr>`;
      }
      content += `</table>`;
    }

    return content;
  };
}
