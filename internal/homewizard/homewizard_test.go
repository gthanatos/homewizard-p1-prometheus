package homewizard

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSuccess(t *testing.T) {
	want := &Data{
		SmrVersion: 50,
		MeterModel:            "Landis + Gyr LGBBLA4415511423",
		WifiSSID:              "chrisdoc_wifi_24",
		WifiStrength:          40,
		TotalPowerImportKwh: 2100.197,
		TotalPowerImportT1Kwh: 2000.197,
		TotalPowerImportT2Kwh: 1776.22,
		TotalPowerExportT1Kwh: 1112.12,
		TotalPowerExportT1Kwh: 1312.12,
		TotalPowerExportT2Kwh: 1243.12,
		ActivePowerW:          368,
		ActivePowerL1W:        231,
		ActivePowerL2W:        129,
		ActivePowerL3W:        8,
		TotalGasM3:            23.3,
		"external": [
			{
			  "unique_id": "37464C4F32313233303531343532",
			  "type": "gas_meter",
			  "timestamp": 230503134621,
			  "value": 16.28,
			  "unit": "m3"
			},
			{
			  "unique_id": "3853455430303030343930343531",
			  "type": "water_meter",
			  "timestamp": 230503134234,
			  "value": 1.554,
			  "unit": "m3"
			}
		  ]		
	}

	p1Response := `{
		"smr_version":50,
		"meter_model":"Landis + Gyr LGBBLA4415511423",
		"wifi_ssid":"chrisdoc_wifi_24",
		"wifi_strength":40,
		"total_power_import_kwh":2100.197,
		"total_power_import_t1_kwh":2000.197,
		"total_power_import_t2_kwh":1776.22,
		"total_power_export_kwh":1112.12,
		"total_power_export_t1_kwh":1312.12,
		"total_power_export_t2_kwh":1243.12,
		"active_power_w":368,
		"active_power_l1_w":231,
		"active_power_l2_w":129,
		"active_power_l3_w":8,
		"total_gas_m3":23.3,
		"total_water_m3": 1.554
		"gas_timestamp":null
		}`

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(p1Response))
	}))

	defer srv.Close()
	// srv URL contains procotol and we are just interested in the host
	host := strings.Replace(srv.URL, "http://", "", 1)

	sut := NewP1ClientWithHTTPClient(host, srv.Client())
	got, err := sut.Retrieve()

	if err != nil {
		t.Errorf("Unexpected error on request: %s", err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Homewizard data mismatch (-want +got):\n%s", diff)
	}
}
