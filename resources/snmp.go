package resources

import (
	"encoding/json"
	"log"
	"net"
	"os"
	"time"

	"github.com/gosnmp/gosnmp"
)

const RETRIES = 3
const TIMEOUT = 2
const EXPONTIALTIMEOUT = true

var Asn1BERErrors = []gosnmp.Asn1BER{gosnmp.Null, gosnmp.NoSuchObject, gosnmp.NoSuchInstance}

type SnmpData struct {
	Value interface{}
	Type  gosnmp.Asn1BER
}
type SnmpResponse map[string]SnmpData

func ParseSnmpPDU(dataUnit gosnmp.SnmpPDU, results SnmpResponse) {
	name := dataUnit.Name[1:] //remove first '.'
	for _, berError := range Asn1BERErrors {
		if dataUnit.Type == berError {
			results[name] = SnmpData{Value: nil, Type: dataUnit.Type}
			return
		}
	}
	switch dataUnit.Type {
	case gosnmp.OctetString:
		results[name] = SnmpData{Value: string(dataUnit.Value.([]byte)), Type: dataUnit.Type}
	default:
		results[name] = SnmpData{Value: gosnmp.ToBigInt(dataUnit.Value.([]byte)), Type: dataUnit.Type}
	}
}

func DoSnmpGet(oids []string, conn *gosnmp.GoSNMP) []byte {
	result, err2 := conn.Get(oids) // Get() accepts up to g.MAX_OIDS
	if err2 != nil {
		log.Printf("%v\n", err2)
		return nil
	}
	results := make(SnmpResponse)
	for _, variable := range result.Variables {
		ParseSnmpPDU(variable, results)
	}
	jsonResults, err := json.Marshal(results)
	if err != nil {
		log.Printf("%v\n", err)
	}
	return jsonResults
}

func DoSnmpBulkWalk(oidRoot string, conn *gosnmp.GoSNMP) []byte {
	results := make(SnmpResponse)

	var handleSnmpPDU = gosnmp.WalkFunc(func(dataUnit gosnmp.SnmpPDU) error {
		ParseSnmpPDU(dataUnit, results)
		return nil
	})

	err := conn.BulkWalk(oidRoot, handleSnmpPDU)
	if err != nil {
		log.Printf("%v\n", err)
		return nil
	}

	jsonResults, err := json.Marshal(results)
	if err != nil {
		log.Printf("%v\n", err)
		return nil
	}
	return jsonResults

}

//// MyTrapHandler TODO: Replace 'file' to proper output, like a WebSocket or a message queue publisher channel
func MyTrapHandler(file *os.File) gosnmp.TrapHandlerFunc {
	return func(packet *gosnmp.SnmpPacket, addr *net.UDPAddr) {
		var results SnmpResponse

		for _, v := range packet.Variables {
			ParseSnmpPDU(v, results)
		}

		jsonResults, err := json.Marshal(results)
		if err != nil {
			log.Printf("Trap() err: %v", err)
		}
		_, _ = file.Write(jsonResults)
	}
}

// ConnectionV2cFactory TODO: I'm using const values for connection, like RETRIES and TIMEOUT. It's fine this way??
func ConnectionV2cFactory(host string, port uint16, community string) *gosnmp.GoSNMP {
	conn := &gosnmp.GoSNMP{
		Target:             host,
		Port:               port,
		Community:          community,
		Version:            gosnmp.Version2c,
		Timeout:            time.Duration(TIMEOUT) + time.Second,
		Retries:            RETRIES,
		ExponentialTimeout: EXPONTIALTIMEOUT,
	}

	return conn
}

// ConnectionV3Factory TODO: I'm using const values for connection, like RETRIES and TIMEOUT. It's fine this way??
func ConnectionV3Factory(host string, port uint16, user string) *gosnmp.GoSNMP {
	conn := &gosnmp.GoSNMP{
		Target:             host,
		Port:               port,
		Version:            gosnmp.Version3,
		Timeout:            time.Duration(TIMEOUT) + time.Second,
		Retries:            RETRIES,
		ExponentialTimeout: EXPONTIALTIMEOUT,
		SecurityModel:      gosnmp.UserSecurityModel,
		MsgFlags:           gosnmp.NoAuthNoPriv,
		SecurityParameters: &gosnmp.UsmSecurityParameters{
			UserName: user,
		},
	}

	return conn
}

func ConnectionV3AuthFactory(host string, port uint16, user string,
	authPass string, auth gosnmp.SnmpV3AuthProtocol) *gosnmp.GoSNMP {

	conn := ConnectionV3Factory(host, port, user)
	conn.MsgFlags = gosnmp.AuthNoPriv
	secParams := conn.SecurityParameters.(*gosnmp.UsmSecurityParameters)
	secParams.AuthenticationProtocol = auth
	secParams.AuthenticationPassphrase = authPass
	return conn
}
func ConnectionV3AuthPrivFactory(host string, port uint16, user string, authPass string, auth gosnmp.SnmpV3AuthProtocol,
	privPass string, privacy gosnmp.SnmpV3PrivProtocol) *gosnmp.GoSNMP {

	conn := ConnectionV3Factory(host, port, user)
	conn.MsgFlags = gosnmp.AuthPriv
	secParams := conn.SecurityParameters.(*gosnmp.UsmSecurityParameters)
	secParams.PrivacyProtocol = privacy
	secParams.PrivacyPassphrase = privPass
	secParams.AuthenticationProtocol = auth
	secParams.AuthenticationPassphrase = authPass
	return conn

}
