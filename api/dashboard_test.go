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

	if (w.Code != http.StatusOK) || !allowed || (contentType != defaultOptionsContent) {
		t.Log("Code: ", w.Code)
		t.Log("Expected Methods: ", wantedMethods, " / Got Methods: ", allowedMethods)
		t.Log("Expected Content-Type: ", defaultOptionsContent, " / Got Content-Type: ", contentType)
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

	if (w.Code != http.StatusOK) || !allowed || (w.Header().Get("Content-Type") != defaultOptionsContent) {
		t.Log("Code: ", w.Code)
		t.Log("Expected Methods: ", wantedMethods, " / Got Methods: ", allowedMethods)
		t.Log("Expected Content-Type: ", defaultOptionsContent, " / Got Content-Type: ", contentType)
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

	if (w.Code != http.StatusOK) || !allowed || (w.Header().Get("Content-Type") != defaultOptionsContent) {
		t.Log("Code: ", w.Code)
		t.Log("Expected Methods: ", wantedMethods, " / Got Methods: ", allowedMethods)
		t.Log("Expected Content-Type: ", defaultOptionsContent, " / Got Content-Type: ", contentType)
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

	if (w.Code != http.StatusOK) || !allowed || (w.Header().Get("Content-Type") != defaultOptionsContent) {
		t.Log("Code: ", w.Code)
		t.Log("Expected Methods: ", wantedMethods, " / Got Methods: ", allowedMethods)
		t.Log("Expected Content-Type: ", defaultOptionsContent, " / Got Content-Type: ", contentType)
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

	if (w.Code != http.StatusOK) || !allowed || (w.Header().Get("Content-Type") != defaultOptionsContent) {
		t.Log("Code: ", w.Code)
		t.Log("Expected Methods: ", wantedMethods, " / Got Methods: ", allowedMethods)
		t.Log("Expected Content-Type: ", defaultOptionsContent, " / Got Content-Type: ", contentType)
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

	if (w.Code != http.StatusOK) || !allowed || (w.Header().Get("Content-Type") != defaultOptionsContent) {
		t.Log("Code: ", w.Code)
		t.Log("Expected Methods: ", wantedMethods, " / Got Methods: ", allowedMethods)
		t.Log("Expected Content-Type: ", defaultOptionsContent, " / Got Content-Type: ", contentType)
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

	if (w.Code != http.StatusOK) || !allowed || (w.Header().Get("Content-Type") != defaultOptionsContent) {
		t.Log("Code: ", w.Code)
		t.Log("Expected Methods: ", wantedMethods, " / Got Methods: ", allowedMethods)
		t.Log("Expected Content-Type: ", defaultOptionsContent, " / Got Content-Type: ", contentType)
		t.Fail()
	}
}
