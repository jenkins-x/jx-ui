import { test, expect } from '@playwright/test'

test('test', async ({ page }) => {
  // Go to http://localhost:3000/
  await page.goto('http://localhost:3000/')
  const dashboardTitle = await page.locator('h2:has-text("Dashboard")')
  await expect(dashboardTitle).toContainText('Dashboard')

})
