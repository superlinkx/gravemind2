package main

import (
	"github.com/gin-gonic/gin"
)

//DashboardOptions gives OPTIONS for dashboard routes
func DashboardOptions(c *gin.Context) {
	setAllowHeader("OPTIONS", c)

	paths := map[string]string{
		"current": "/gravemind/v1/dashboard/current",
		"daily":   "/gravemind/v1/dashboard/daily",
		"weekly":  "/gravemind/v1/dashboard/weekly",
		"monthly": "/gravemind/v1/dashboard/monthly",
		"yearly":  "/gravemind/v1/dashboard/yearly",
		"custom":  "/gravemind/v1/dashboard/custom",
	}

	c.JSON(200, gin.H{
		"message": "Welcome to the gravemind dashboard. Endpoints are in the paths.",
		"paths":   paths,
	})
}

// CurrentDashOptions gives OPTIONS for `current` route
func CurrentDashOptions(c *gin.Context) {
	setAllowHeader("OPTIONS, GET, POST", c)

	actions := map[string]action{
		"GET": action{
			Description: "Returns JSON object of current dashboard results",
			Params:      "No parameters",
		},
		"POST": action{
			Description: "Accepts a special JSON payload for submitting current dashboard raw data",
			Params:      "Post with data formatted according to the docs: https://github.com/superlinkx/gravemind2/wiki/POSTCurrentData",
		},
	}

	c.JSON(200, gin.H{
		"message": "Welcome to the current endpoint. See actions.",
		"actions": actions,
	})
}

// DailyDashOptions gives OPTIONS for `daily` route
func DailyDashOptions(c *gin.Context) {
	setAllowHeader("OPTIONS, GET, POST", c)

	actions := map[string]action{
		"GET": action{
			Description: "Returns JSON object of daily dashboard results",
			Params:      "/daily/{date}, where {date} is formatted as '01-12-2009'",
		},
		"POST": action{
			Description: "Accepts a special JSON payload for submitting daily dashboard raw data",
			Params:      "Post with data formatted according to the docs: https://github.com/superlinkx/gravemind2/wiki/POSTDailyData",
		},
	}

	c.JSON(200, gin.H{
		"message": "Welcome to the daily endpoint. See actions.",
		"actions": actions,
	})
}

// WeeklyDashOptions gives OPTIONS for `weekly` route
func WeeklyDashOptions(c *gin.Context) {
	setAllowHeader("OPTIONS, GET", c)

	actions := map[string]action{
		"GET": action{
			Description: "Returns JSON object of weekly dashboard results",
			Params:      "/weekly/{date}, where {date} is formatted as '01-12-2009'",
		},
	}

	c.JSON(200, gin.H{
		"message": "Welcome to the weekly endpoint. See actions.",
		"actions": actions,
	})
}

// MonthlyDashOptions gives OPTIONS for `monthly` route
func MonthlyDashOptions(c *gin.Context) {
	setAllowHeader("OPTIONS, GET", c)

	actions := map[string]action{
		"GET": action{
			Description: "Returns JSON object of monthly dashboard results",
			Params:      "/weekly/{month}, where {month} is formatted as '01-2009'",
		},
	}

	c.JSON(200, gin.H{
		"message": "Welcome to the monthly endpoint. See actions.",
		"actions": actions,
	})
}

// YearlyDashOptions gives OPTIONS for `yearly` route
func YearlyDashOptions(c *gin.Context) {
	setAllowHeader("OPTIONS, GET", c)

	actions := map[string]action{
		"GET": action{
			Description: "Returns JSON object of yearly dashboard results",
			Params:      "/yearly/{year}, where {year} is formatted as '2009'",
		},
	}

	c.JSON(200, gin.H{
		"message": "Welcome to the yearly endpoint. See actions.",
		"actions": actions,
	})
}

// CustomDashOptions gives OPTIONS for `custom` route
func CustomDashOptions(c *gin.Context) {
	setAllowHeader("OPTIONS, GET", c)

	actions := map[string]action{
		"GET": action{
			Description: "Returns JSON object of custom date range dashboard results",
			Params:      "/custom/{startdate}/{enddate}, where {startdate} and {enddate} are formatted as '01-12-2009'",
		},
	}

	c.JSON(200, gin.H{
		"message": "Welcome to the custom endpoint. See actions.",
		"actions": actions,
	})
}

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
