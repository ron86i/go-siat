package utils

import "encoding/xml"

// MarshalIndentedXML serializes an XML document with two-space indentation.
// It does not add an XML declaration so the result can be signed without
// altering the document afterwards.
func MarshalIndentedXML(documento any) ([]byte, error) {
	return xml.MarshalIndent(documento, "", "  ")
}
