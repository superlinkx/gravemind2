package main

import "github.com/superlinkx/gravemind2/api"

func initializeRoutes() {
	router.OPTIONS("/gravemind", api.GravemindOptions)
	router.OPTIONS("/gravemind/v1", api.V1Options)

	dashboardRoutes := router.Group("/gravemind/v1/dashboard")
	{
		dashboardRoutes.OPTIONS("/", api.DashboardOptions)
		dashboardRoutes.OPTIONS("/current", api.CurrentDashOptions)
		dashboardRoutes.OPTIONS("/daily", api.DailyDashOptions)
		dashboardRoutes.OPTIONS("/weekly", api.WeeklyDashOptions)
		dashboardRoutes.OPTIONS("/monthly", api.MonthlyDashOptions)
		dashboardRoutes.OPTIONS("/yearly", api.YearlyDashOptions)
		dashboardRoutes.OPTIONS("/custom", api.CustomDashOptions)
		dashboardRoutes.GET("/current", api.GetCurrentDash)
		dashboardRoutes.GET("/daily", api.GetDailyDash)
		dashboardRoutes.GET("/weekly", api.GetWeeklyDash)
		dashboardRoutes.GET("/monthly", api.GetMonthlyDash)
		dashboardRoutes.GET("/yearly", api.GetYearlyDash)
		dashboardRoutes.GET("/custom", api.GetCustomDash)
		dashboardRoutes.POST("/current", api.PostCurrentDash)
		dashboardRoutes.POST("/daily", api.PostDailyDash)
	}
}
