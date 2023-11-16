import { createDashboardFromModel } from "../utils/dashboardHelpers";
import { createAdBidsModel } from "../utils/dataSpecifcHelpers";
import { test, expect } from "@playwright/test";
import { startRuntimeForEachTest } from "../utils/startRuntimeForEachTest";

test.describe("~~~~~~~~~~~~~~~~~~~~FIXME RENAME THIS~~~~~~~~~~~~~~~~~~~~~~~", () => {
  startRuntimeForEachTest();

  test("~~~~~~~~~~~~~~~~~~~~FIXME RENAME THIS~~~~~~~~~~~~~~~~~~~~~~~", async ({
    page,
  }) => {
    test.setTimeout(60000);
    await page.goto("/");
    // disable animations
    await page.addStyleTag({
      content: `
        *, *::before, *::after {
          animation-duration: 0s !important;
          transition-duration: 0s !important;
        }
      `,
    });
    await createAdBidsModel(page);
    await createDashboardFromModel(page, "AdBids_model");

    // Delete this when your flow is ready.
    await page.pause();

    await expect(page.getByText("example expect - will fail")).toBeVisible();
  });
});
