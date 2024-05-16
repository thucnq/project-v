package mappingcdn

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"
	"path"
	"strings"

	"project-v/pkg/l"
)

type CDNMappingFile struct {
	ThumbMapping  map[string]string `json:"ThumbMapping"`
	UploadMapping map[string]string `json:"UploadMapping"`
}

var cdnMappingData *CDNMappingFile

func LoadCdnFromFile(cdnConfigFile string, ll l.Logger) {
	if cdnConfigFile == "" {
		ll.Warn("cdn config file is empty")
		return
	}
	jsonFile, err := os.Open(cdnConfigFile)
	// if we os.Open returns an error then handle it
	if err != nil {
		ll.Warn(
			"cannot read cdn config file",
			l.String("cdnConfigFile", cdnConfigFile), l.Error(err),
		)
		return
	}
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)
	var data CDNMappingFile
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		ll.Error("failed to parse cdn config file", l.Error(err))
		return
	}
	cdnMappingData = &data
	ll.Info("load cdn config file mapping success")
}

func GetImageStoreToCdnMapping() map[string]string {
	if cdnMappingData != nil && cdnMappingData.UploadMapping != nil {
		return cdnMappingData.UploadMapping
	}
	return map[string]string{}
}

func GetImageOriginalToThumbCdnMapping() map[string]string {
	if cdnMappingData != nil && cdnMappingData.ThumbMapping != nil {
		return cdnMappingData.ThumbMapping
	}
	return map[string]string{}
}

func encodeLastSegment(rawUrl string) string {
	urlArr := strings.Split(rawUrl, "/")
	lastIdx := len(urlArr) - 1

	urlArr[lastIdx] = strings.ReplaceAll(urlArr[lastIdx], " ", "+space+")

	tmp, err := url.QueryUnescape(urlArr[lastIdx])
	if err != nil {
		return strings.Join(urlArr, "/")
	}

	urlArr[lastIdx] = url.QueryEscape(tmp)
	urlArr[lastIdx] = strings.ReplaceAll(urlArr[lastIdx], "+space+", " ")
	return strings.Join(urlArr, "/")
}

func ConvertRawImageToCdn(rawUrl string) string {
	_url, err := url.Parse(rawUrl)
	if err != nil {
		return rawUrl
	}

	host := _url.Host
	mapping := GetImageStoreToCdnMapping()
	if val, ok := mapping[host]; ok {
		tmpURL := strings.ReplaceAll(rawUrl, host, val)
		return encodeLastSegment(tmpURL)
	}

	return encodeLastSegment(rawUrl)
}

func ConvertRawFileToCdn(rawUrl string) string {
	_url, err := url.Parse(rawUrl)
	if err != nil {
		return rawUrl
	}

	host := _url.Host
	mapping := GetImageStoreToCdnMapping()
	if val, ok := mapping[host]; ok {
		tmpURL := strings.ReplaceAll(rawUrl, host, val)
		return encodeLastSegment(tmpURL)
	}

	return rawUrl
}

func ConvertRawImageToThumbnailCdn(cover string) string {
	cdnUrl := ConvertRawImageToCdn(cover)
	if cdnUrl == "" {
		return cdnUrl
	}

	_url, err := url.Parse(cdnUrl)
	if err != nil {
		return cdnUrl
	}

	host := _url.Host
	mapping := GetImageOriginalToThumbCdnMapping()
	thumbCdn, ok := mapping[host]
	if !ok {
		return cover
	}

	imageID := path.Base(_url.Path)
	if len(imageID) <= 0 {
		return cdnUrl
	}

	if strings.Index(_url.Path, "/images/") != 0 {
		return cdnUrl
	}

	_url.Path = strings.Replace(_url.Path, "/images/", "", 1)
	return fmt.Sprintf("https://%s/$size$/smart/%s", thumbCdn, _url.Path)
}

func ConvertRawVideoToCdn(rawUrl string) string {
	_url, err := url.Parse(rawUrl)
	if err != nil {
		return rawUrl
	}

	host := _url.Host
	mapping := GetImageStoreToCdnMapping()
	if val, ok := mapping[host]; ok {
		// rawUrl = strings.ReplaceAll(rawUrl, "videos/", "videos/results/")
		return strings.ReplaceAll(rawUrl, host, val)
	}

	return rawUrl
}
