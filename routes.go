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
		dashboardRoutes.GET("/current", api.RetrieveCurrentDash)
		dashboardRoutes.GET("/daily", api.RetrieveDailyDash)
		dashboardRoutes.GET("/weekly", api.RetrieveWeeklyDash)
		dashboardRoutes.GET("/monthly", api.RetrieveMonthlyDash)
		dashboardRoutes.GET("/yearly", api.RetrieveYearlyDash)
		dashboardRoutes.GET("/custom", api.RetrieveCustomDash)
		dashboardRoutes.POST("/current", api.PostCurrentDash)
		dashboardRoutes.POST("/daily", api.PostDailyDash)
	}
}
