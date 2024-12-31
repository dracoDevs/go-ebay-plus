package utils

import "bytes"

func RemoveEndItemXML(xmlStr string) string {
	startTag := "<EndItem>"
	endTag := "</EndItem>"

	startIndex := bytes.Index([]byte(xmlStr), []byte(startTag))
	endIndex := bytes.Index([]byte(xmlStr), []byte(endTag))

	if startIndex != -1 && endIndex != -1 && endIndex > startIndex {
		xmlStr = xmlStr[:startIndex] + xmlStr[endIndex+len(endTag):]
	}

	return xmlStr
}
