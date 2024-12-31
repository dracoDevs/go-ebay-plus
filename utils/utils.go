package utils

import (
	"bytes"
	"strings"
)

func RemoveEndItemXML(xmlStr string) string {
	startTag := "<EndItem>"
	endTag := "</EndItem>"

	startIndex := bytes.Index([]byte(xmlStr), []byte(startTag))
	endIndex := bytes.Index([]byte(xmlStr), []byte(endTag))

	if startIndex != -1 && endIndex != -1 && endIndex > startIndex {
		innerContent := xmlStr[startIndex+len(startTag) : endIndex]
		xmlStr = xmlStr[:startIndex] + innerContent + xmlStr[endIndex+len(endTag):]
		xmlStr = strings.ReplaceAll(xmlStr, startTag, "")
		xmlStr = strings.ReplaceAll(xmlStr, endTag, "")
	}

	return xmlStr
}
