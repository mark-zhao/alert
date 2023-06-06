package alert

import "testing"

func TestCephAlert(t *testing.T){
	alertHandler := NewWeChatAlert()
    alertMsg := "this is test"

	if err := alertHandler.getToken(); err != nil{
		t.Error("unittest alerHandler getToken failed; msg:", err)
	}
    t.Logf("unittest alertHandler getToken pass; token= %s,\n expire=%d\n", alertHandler.token, alertHandler.expire)

	if err := alertHandler.Alert(alertMsg); err != nil{
		t.Error("unittest alertHandler alert failed; msg:", err)
	}
	t.Log("unittest alertHandler pass")
}
