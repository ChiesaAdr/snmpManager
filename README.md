# SNMP MANAGER

A GoLang module to manage SNMP agents. This project provides functions that receives as parameters root OIDs or unique OIDs  and returns responses in JSON format.
It also provides handlers for SNMP traps and SNMP informs that receives a MQTT channel to publish the SNMP inputs in the same JSON format.

JSON format adopted:
    
    {"<OID>":{"Value":<value>,"Type":<int>}}

Where the value for "Value" can be anything, like a string or an integer or even null, 
and the value for "Type" follows the SNMP RFCs:

    // Asn1BER's - http://www.ietf.org/rfc/rfc1442.txt
    const (
    EndOfContents     Asn1BER = 0x00
    UnknownType       Asn1BER = 0x00
    Boolean           Asn1BER = 0x01
    Integer           Asn1BER = 0x02
    BitString         Asn1BER = 0x03
    OctetString       Asn1BER = 0x04
    Null              Asn1BER = 0x05
    ObjectIdentifier  Asn1BER = 0x06
    ObjectDescription Asn1BER = 0x07
    IPAddress         Asn1BER = 0x40
    Counter32         Asn1BER = 0x41
    Gauge32           Asn1BER = 0x42
    TimeTicks         Asn1BER = 0x43
    Opaque            Asn1BER = 0x44
    NsapAddress       Asn1BER = 0x45
    Counter64         Asn1BER = 0x46
    Uinteger32        Asn1BER = 0x47
    OpaqueFloat       Asn1BER = 0x78
    OpaqueDouble      Asn1BER = 0x79
    NoSuchObject      Asn1BER = 0x80
    NoSuchInstance    Asn1BER = 0x81
    EndOfMibView      Asn1BER = 0x82
    )



## Examples


Get uniques OID by SNMPv2 :

    community := "public"
    host := "localhost"
    port := "161"
    conn := resources.ConnectionV2cFactory(host, uint16(port),community)
    oids := []string{"1.3.6.1.4.1.26138.1.2.1.1.1.6.1", "1.3.6.1.4.1.26138.1.2.1.1.1.6.2"}
    jsonResults := resources.DoSnmpGet(oids, conn)
    fmt.Println(string(jsonResults))

Output: 

     {"1.3.6.1.4.1.26138.1.2.1.1.1.6.1":{"Value":"ITBS-00000001","Type":4},"1.3.6.1.4.1.26138.1.2.1.1.1.6.2":{"Value":"ITBS-00000002","Type":4}}
