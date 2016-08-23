package api

import (
	"github.com/gin-gonic/gin"
)

//DashboardOptions gives OPTIONS for dashboard routes
func DashboardOptions(c *gin.Context) {}

// CurrentDashOptions gives OPTIONS for `current` route
func CurrentDashOptions(c *gin.Context) {}

// DailyDashOptions gives OPTIONS for `daily` route
func DailyDashOptions(c *gin.Context) {}

// WeeklyDashOptions gives OPTIONS for `weekly` route
func WeeklyDashOptions(c *gin.Context) {}

// MonthlyDashOptions gives OPTIONS for `monthly` route
func MonthlyDashOptions(c *gin.Context) {}

// YearlyDashOptions gives OPTIONS for `yearly` route
func YearlyDashOptions(c *gin.Context) {}

// CustomDashOptions gives OPTIONS for `custom` route
func CustomDashOptions(c *gin.Context) {}

// RetrieveCurrentDash gets dashboard for `current` route
func RetrieveCurrentDash(c *gin.Context) {}

// RetrieveDailyDash gets dashboard for `daily` route
func RetrieveDailyDash(c *gin.Context) {}

// RetrieveWeeklyDash gets dashboard for `weekly` route
func RetrieveWeeklyDash(c *gin.Context) {}

// RetrieveMonthlyDash gets dashboard for `monthly` route
func RetrieveMonthlyDash(c *gin.Context) {}

// RetrieveYearlyDash gets dashboard for `yearly` route
func RetrieveYearlyDash(c *gin.Context) {}

// RetrieveCustomDash gets dashboard for `custom` route
func RetrieveCustomDash(c *gin.Context) {}

// PostCurrentDash commits new dashboard data for `current` route
func PostCurrentDash(c *gin.Context) {}

// PostDailyDash commits new dashboard data for `daily` route
func PostDailyDash(c *gin.Context) {}
