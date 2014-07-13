package main

import (
	"github.com/bitly/go-simplejson"
	"github.com/glendc/cgreader"
	"strings"
)

func GetEmbedFiles(jsonInfo *simplejson.Json, files []string) []string {
	if i, err := jsonInfo.String(); err != nil {
		if files, err = jsonInfo.StringArray(); err != nil {
			if jsonPath := jsonInfo.Get(INFO_PATH); jsonPath != nil {
				if path, errPath := jsonPath.String(); errPath == nil {
					if jsonAmount := jsonInfo.Get(INFO_AMOUNT); jsonAmount != nil {
						if amount, errAmount := jsonAmount.Int(); errAmount == nil {
							if strings.Contains(path, "*") {
								path = strings.Replace(path, "*", "%d", 1)
								files = cgreader.GetFileList(path, amount)
							} else {
								ErrorIllegalEmbbedFormatSmartPath()
							}
						} else {
							ErrorMessage(errAmount.Error())
						}
					} else {
						ErrorMissingInEmbedFormat(INFO_AMOUNT)
					}
				} else {
					ErrorMessage(errPath.Error())
				}
			} else {
				ErrorMissingInEmbedFormat(INFO_PATH)
			}
		}
	} else {
		files = append(files, i)
	}

	return files
}
