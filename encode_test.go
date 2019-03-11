package gosoap

import (
	"testing"
)

var (
	tests = []struct {
		Params Params
		XMLParams string
		Err    string
	}{
		{
			Params: Params{"": ""},
			Err:    "error expected: xml: start tag with no name",
		},
		{
			XMLParams: `
<foo fooAttr="foo">
  <bar barAttr="bar">
    bar
  </bar>
  <baz>1</baz>
</foo>`,
		},
	}
)

func TestClient_MarshalXML(t *testing.T) {
	for _, test := range tests {
		soap, err := SoapClient("http://ec.europa.eu/taxation_customs/vies/checkVatService.wsdl")
		if err != nil {
			t.Errorf("error not expected: %s", err)
		}

		if test.Params != nil {
			err = soap.Call("checkVat", test.Params)
		} else if test.XMLParams != "" {
			err = soap.Call("checkVat", test.XMLParams)
		}

		if err == nil && test.Err != "" {
			t.Errorf(test.Err)
		} else if err != nil && test.Err == "" {
			t.Errorf("error not expected: %s", err)
		}
	}
}
