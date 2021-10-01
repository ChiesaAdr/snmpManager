package main

// resources "github.com/ChiesaAdr/snmpManager/resourcesSnmp"

// const HOST = "snmpmocker"
// const PORT = 161

// func testDoGetErrors(conn *gosnmp.GoSNMP) func(*testing.T) {
// 	noSuchInstanceOid := "1.3.6.1.4.1.26138.1.2.1.1.1.6.2000"
// 	noSuchObjectOid := "1.3.6.1.4.1.26138.1.2.1.1.1.61.1"
// 	oids := []string{noSuchInstanceOid, noSuchObjectOid}
// 	return func(t *testing.T) {
// 		jsonResults := resources.DoSnmpGet(oids, conn)
// 		var results resources.SnmpResponse
// 		err := json.Unmarshal(jsonResults, &results)
// 		if err != nil {
// 			t.Fatalf("Error in Unmarshal json results")
// 		}
// 		if results[noSuchObjectOid].Value != nil || results[noSuchObjectOid].Type != gosnmp.NoSuchObject {
// 			t.Fatalf("Unexpected response for not implemented object OID %v %d", results[noSuchObjectOid].Value, results[noSuchObjectOid].Type)
// 		}
// 		if results[noSuchInstanceOid].Value != nil || results[noSuchInstanceOid].Type != gosnmp.NoSuchInstance {
// 			t.Fatalf("Unexpected response for not implemented instance OID")
// 		}
// 	}
// }
// func testDoGet(conn *gosnmp.GoSNMP) func(*testing.T) {
// 	oids := []string{"1.3.6.1.4.1.26138.1.2.1.1.1.6.1", "1.3.6.1.4.1.26138.1.2.1.1.1.6.2"}
// 	baseSn := "ITBS-00000"
// 	return func(t *testing.T) {
// 		jsonResults := resources.DoSnmpGet(oids, conn)
// 		var results resources.SnmpResponse
// 		err := json.Unmarshal(jsonResults, &results)
// 		if err != nil {
// 			t.Fatalf("Error in Unmarshal json results")
// 		}
// 		for i, oid := range oids {
// 			sn := fmt.Sprintf("%s%03x", baseSn, i+1)
// 			if results[oid].Value.(string) != sn {
// 				t.Fatalf("Value its not the expected. Value returned: " + results[oid].Value.(string) + " Value expected: " + sn)
// 			}
// 		}
// 	}
// }

// func testDoBulkWalk(conn *gosnmp.GoSNMP) func(t *testing.T) {
// 	oidRoot := "1.3.6.1.4.1.26138.1.2.1.1.1.6"
// 	baseSn := "ITBS-00000"
// 	return func(t *testing.T) {
// 		jsonResults := resources.DoSnmpBulkWalk(oidRoot, conn)
// 		results := make(resources.SnmpResponse)
// 		err := json.Unmarshal(jsonResults, &results)
// 		if err != nil {
// 			t.Fatalf("Error in Unmarshal json results")
// 		}
// 		for i := 0; i < len(results); i++ {
// 			oid := oidRoot + "." + strconv.Itoa(i+1)
// 			sn := fmt.Sprintf("%s%03x", baseSn, i+1)
// 			if results[oid].Value.(string) != sn {
// 				t.Fatalf("Value its not the expected. Value returned: " + results[oid].Value.(string) + " Value expected: " + sn)
// 			}
// 		}
// 	}
// }

// func connectAndTestGets(conn *gosnmp.GoSNMP, t *testing.T) {
// 	err := conn.Connect()
// 	if err != nil {
// 		t.Fatalf("Connect() err: %v", err)
// 	}
// 	t.Run("testDoGet", testDoGet(conn))
// 	t.Run("testDoGetErrors", testDoGetErrors(conn))
// 	t.Run("testDoBulkWalk", testDoBulkWalk(conn))
// 	_ = conn.Conn.Close()
// }

// func TestV2c(t *testing.T) {
// 	community := "public"
// 	conn := resources.ConnectionV2cFactory(HOST, uint16(PORT), community)
// 	connectAndTestGets(conn, t)
// }

// func TestV3NoAuthNoPriv(t *testing.T) {
// 	user := "testNoAuthNoPriv"
// 	conn := resources.ConnectionV3Factory(HOST, uint16(PORT), user)
// 	connectAndTestGets(conn, t)
// }

// func TestV3AuthNoPriv(t *testing.T) {
// 	user := "testAuth"
// 	authpass := "authpass"
// 	authproto := gosnmp.MD5
// 	conn := resources.ConnectionV3AuthFactory(HOST, uint16(PORT), user, authpass, authproto)
// 	connectAndTestGets(conn, t)
// }

// func TestV3AuthPriv(t *testing.T) {
// 	user := "testAuthPriv"
// 	authpass := "authpass"
// 	authproto := gosnmp.MD5
// 	privpass := "privpass"
// 	privproto := gosnmp.DES
// 	conn := resources.ConnectionV3AuthPrivFactory(HOST, uint16(PORT), user, authpass, authproto, privpass, privproto)
// 	connectAndTestGets(conn, t)
// }
