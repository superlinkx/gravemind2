package main

func initializeRoutes() {
	router.OPTIONS("/gravemind", GravemindOptions)
	router.OPTIONS("/gravemind/v1", V1Options)

	dashboardRoutes := router.Group("/gravemind/v1/dashboard")
	{
		dashboardRoutes.OPTIONS("/", DashboardOptions)
		dashboardRoutes.OPTIONS("/current", CurrentDashOptions)
		dashboardRoutes.OPTIONS("/daily", DailyDashOptions)
		dashboardRoutes.OPTIONS("/weekly", WeeklyDashOptions)
		dashboardRoutes.OPTIONS("/monthly", MonthlyDashOptions)
		dashboardRoutes.OPTIONS("/yearly", YearlyDashOptions)
		dashboardRoutes.OPTIONS("/custom", CustomDashOptions)
		dashboardRoutes.GET("/current", RetrieveCurrentDash)
		dashboardRoutes.GET("/daily", RetrieveDailyDash)
		dashboardRoutes.GET("/weekly", RetrieveWeeklyDash)
		dashboardRoutes.GET("/monthly", RetrieveMonthlyDash)
		dashboardRoutes.GET("/yearly", RetrieveYearlyDash)
		dashboardRoutes.GET("/custom", RetrieveCustomDash)
		dashboardRoutes.POST("/current", PostCurrentDash)
		dashboardRoutes.POST("/daily", PostDailyDash)
	}
}
