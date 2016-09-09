package api

import (
	"net/http"
	"testing"
)

func TestDashboardOptions(t *testing.T) {
	w, _ := testHandler("/gravemind/v1/dashboard", DashboardOptions, "OPTIONS", nil)

	allowedMethods := w.Header().Get("Allow")
	wantedMethods := []string{"OPTIONS"}
	contentType := w.Header().Get("Content-Type")

	allowed := testAllowedMethods(allowedMethods, wantedMethods)

	t.Log("Returned: ", w.Body)

	if (w.Code != http.StatusOK) || !allowed || (contentType != defaultContentType) {
		t.Log("Code: ", w.Code)
		t.Log("Expected Methods: ", wantedMethods, " / Got Methods: ", allowedMethods)
		t.Log("Expected Content-Type: ", defaultContentType, " / Got Content-Type: ", contentType)
		t.Fail()
	}
}

func TestCurrentDashOptions(t *testing.T) {
	w, _ := testHandler("/gravemind/v1/dashboard/current", CurrentDashOptions, "OPTIONS", nil)

	allowedMethods := w.Header().Get("Allow")
	wantedMethods := []string{"OPTIONS", "GET", "POST"}
	contentType := w.Header().Get("Content-Type")

	allowed := testAllowedMethods(allowedMethods, wantedMethods)

	t.Log("Returned: ", w.Body)

	if (w.Code != http.StatusOK) || !allowed || (w.Header().Get("Content-Type") != defaultContentType) {
		t.Log("Code: ", w.Code)
		t.Log("Expected Methods: ", wantedMethods, " / Got Methods: ", allowedMethods)
		t.Log("Expected Content-Type: ", defaultContentType, " / Got Content-Type: ", contentType)
		t.Fail()
	}
}

func TestDailyDashOptions(t *testing.T) {
	w, _ := testHandler("/gravemind/v1/dashboard/daily", DailyDashOptions, "OPTIONS", nil)

	allowedMethods := w.Header().Get("Allow")
	wantedMethods := []string{"OPTIONS", "GET", "POST"}
	contentType := w.Header().Get("Content-Type")

	allowed := testAllowedMethods(allowedMethods, wantedMethods)

	t.Log("Returned: ", w.Body)

	if (w.Code != http.StatusOK) || !allowed || (w.Header().Get("Content-Type") != defaultContentType) {
		t.Log("Code: ", w.Code)
		t.Log("Expected Methods: ", wantedMethods, " / Got Methods: ", allowedMethods)
		t.Log("Expected Content-Type: ", defaultContentType, " / Got Content-Type: ", contentType)
		t.Fail()
	}
}

func TestWeeklyDashOptions(t *testing.T) {
	w, _ := testHandler("/gravemind/v1/dashboard/weekly", WeeklyDashOptions, "OPTIONS", nil)

	allowedMethods := w.Header().Get("Allow")
	wantedMethods := []string{"OPTIONS", "GET"}
	contentType := w.Header().Get("Content-Type")

	allowed := testAllowedMethods(allowedMethods, wantedMethods)

	t.Log("Returned: ", w.Body)

	if (w.Code != http.StatusOK) || !allowed || (w.Header().Get("Content-Type") != defaultContentType) {
		t.Log("Code: ", w.Code)
		t.Log("Expected Methods: ", wantedMethods, " / Got Methods: ", allowedMethods)
		t.Log("Expected Content-Type: ", defaultContentType, " / Got Content-Type: ", contentType)
		t.Fail()
	}
}

func TestMonthlyDashOptions(t *testing.T) {
	w, _ := testHandler("/gravemind/v1/dashboard/monthly", MonthlyDashOptions, "OPTIONS", nil)

	allowedMethods := w.Header().Get("Allow")
	wantedMethods := []string{"OPTIONS", "GET"}
	contentType := w.Header().Get("Content-Type")

	allowed := testAllowedMethods(allowedMethods, wantedMethods)

	t.Log("Returned: ", w.Body)

	if (w.Code != http.StatusOK) || !allowed || (w.Header().Get("Content-Type") != defaultContentType) {
		t.Log("Code: ", w.Code)
		t.Log("Expected Methods: ", wantedMethods, " / Got Methods: ", allowedMethods)
		t.Log("Expected Content-Type: ", defaultContentType, " / Got Content-Type: ", contentType)
		t.Fail()
	}
}

func TestYearlyDashOptions(t *testing.T) {
	w, _ := testHandler("/gravemind/v1/dashboard/yearly", YearlyDashOptions, "OPTIONS", nil)

	allowedMethods := w.Header().Get("Allow")
	wantedMethods := []string{"OPTIONS", "GET"}
	contentType := w.Header().Get("Content-Type")

	allowed := testAllowedMethods(allowedMethods, wantedMethods)

	t.Log("Returned: ", w.Body)

	if (w.Code != http.StatusOK) || !allowed || (w.Header().Get("Content-Type") != defaultContentType) {
		t.Log("Code: ", w.Code)
		t.Log("Expected Methods: ", wantedMethods, " / Got Methods: ", allowedMethods)
		t.Log("Expected Content-Type: ", defaultContentType, " / Got Content-Type: ", contentType)
		t.Fail()
	}
}

func TestCustomDashOptions(t *testing.T) {
	w, _ := testHandler("/gravemind/v1/dashboard/custom", CustomDashOptions, "OPTIONS", nil)

	allowedMethods := w.Header().Get("Allow")
	wantedMethods := []string{"OPTIONS", "GET"}
	contentType := w.Header().Get("Content-Type")

	allowed := testAllowedMethods(allowedMethods, wantedMethods)

	t.Log("Returned: ", w.Body)

	if (w.Code != http.StatusOK) || !allowed || (w.Header().Get("Content-Type") != defaultContentType) {
		t.Log("Code: ", w.Code)
		t.Log("Expected Methods: ", wantedMethods, " / Got Methods: ", allowedMethods)
		t.Log("Expected Content-Type: ", defaultContentType, " / Got Content-Type: ", contentType)
		t.Fail()
	}
}

func TestGetCurrentDash(t *testing.T) {
	w, _ := testHandler("/gravemind/v1/dashboard/current", GetCurrentDash, "GET", nil)

	contentType := w.Header().Get("Content-Type")

	if (w.Code != http.StatusOK) || w.Header().Get("Content-Type") != defaultContentType {
		t.Log("Code: ", w.Code)
		t.Log("Expected Content-Type: ", defaultContentType, " / Got Content-Type: ", contentType)
		t.Fail()
	}
}

func TestGetDailyDash(t *testing.T) {
	w, _ := testHandler("/gravemind/v1/dashboard/daily", GetDailyDash, "GET", nil)

	contentType := w.Header().Get("Content-Type")

	if (w.Code != http.StatusOK) || w.Header().Get("Content-Type") != defaultContentType {
		t.Log("Code: ", w.Code)
		t.Log("Expected Content-Type: ", defaultContentType, " / Got Content-Type: ", contentType)
		t.Fail()
	}
}

func TestGetWeeklyDash(t *testing.T) {
	w, _ := testHandler("/gravemind/v1/dashboard/weekly", GetWeeklyDash, "GET", nil)

	contentType := w.Header().Get("Content-Type")

	if (w.Code != http.StatusOK) || w.Header().Get("Content-Type") != defaultContentType {
		t.Log("Code: ", w.Code)
		t.Log("Expected Content-Type: ", defaultContentType, " / Got Content-Type: ", contentType)
		t.Fail()
	}
}

func TestGetMonthlyDash(t *testing.T) {
	w, _ := testHandler("/gravemind/v1/dashboard/monthly", GetMonthlyDash, "GET", nil)

	contentType := w.Header().Get("Content-Type")

	if (w.Code != http.StatusOK) || w.Header().Get("Content-Type") != defaultContentType {
		t.Log("Code: ", w.Code)
		t.Log("Expected Content-Type: ", defaultContentType, " / Got Content-Type: ", contentType)
		t.Fail()
	}
}

func TestGetYearlyDash(t *testing.T) {
	w, _ := testHandler("/gravemind/v1/dashboard/yearly", GetYearlyDash, "GET", nil)

	contentType := w.Header().Get("Content-Type")

	if (w.Code != http.StatusOK) || w.Header().Get("Content-Type") != defaultContentType {
		t.Log("Code: ", w.Code)
		t.Log("Expected Content-Type: ", defaultContentType, " / Got Content-Type: ", contentType)
		t.Fail()
	}
}

func TestGetCustomDash(t *testing.T) {
	w, _ := testHandler("/gravemind/v1/dashboard/custom", GetCustomDash, "GET", nil)

	contentType := w.Header().Get("Content-Type")

	if (w.Code != http.StatusOK) || w.Header().Get("Content-Type") != defaultContentType {
		t.Log("Code: ", w.Code)
		t.Log("Expected Content-Type: ", defaultContentType, " / Got Content-Type: ", contentType)
		t.Fail()
	}
}

func TestPostCurrentDash(t *testing.T) {
	w, _ := testHandler("/gravemind/v1/dashboard/current", PostCurrentDash, "GET", nil)

	contentType := w.Header().Get("Content-Type")

	if (w.Code != http.StatusOK) || w.Header().Get("Content-Type") != defaultContentType {
		t.Log("Code: ", w.Code)
		t.Log("Expected Content-Type: ", defaultContentType, " / Got Content-Type: ", contentType)
		t.Fail()
	}
}

func TestPostDailyDash(t *testing.T) {
	w, _ := testHandler("/gravemind/v1/dashboard/daily", GetDailyDash, "GET", nil)

	contentType := w.Header().Get("Content-Type")

	if (w.Code != http.StatusOK) || w.Header().Get("Content-Type") != defaultContentType {
		t.Log("Code: ", w.Code)
		t.Log("Expected Content-Type: ", defaultContentType, " / Got Content-Type: ", contentType)
		t.Fail()
	}
}
