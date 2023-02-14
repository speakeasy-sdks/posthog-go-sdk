package shared

import (
	"time"
)

type ExportedAssetExportFormatEnum string

const (
	ExportedAssetExportFormatEnumImagePng       ExportedAssetExportFormatEnum = "image/png"
	ExportedAssetExportFormatEnumApplicationPdf ExportedAssetExportFormatEnum = "application/pdf"
	ExportedAssetExportFormatEnumTextCsv        ExportedAssetExportFormatEnum = "text/csv"
)

// ExportedAsset
// Standard ExportedAsset serializer that doesn't return content.
type ExportedAsset struct {
	CreatedAt     time.Time                     `json:"created_at"`
	Dashboard     *int64                        `json:"dashboard,omitempty"`
	ExportContext map[string]interface{}        `json:"export_context,omitempty"`
	ExportFormat  ExportedAssetExportFormatEnum `json:"export_format"`
	Filename      string                        `json:"filename"`
	HasContent    string                        `json:"has_content"`
	ID            int64                         `json:"id"`
	Insight       *int64                        `json:"insight,omitempty"`
}

// ExportedAssetInput
// Standard ExportedAsset serializer that doesn't return content.
type ExportedAssetInput struct {
	Dashboard     *int64                        `json:"dashboard,omitempty" form:"name=dashboard" multipartForm:"name=dashboard"`
	ExportContext map[string]interface{}        `json:"export_context,omitempty" form:"name=export_context,json" multipartForm:"name=export_context,json"`
	ExportFormat  ExportedAssetExportFormatEnum `json:"export_format" form:"name=export_format" multipartForm:"name=export_format"`
	Insight       *int64                        `json:"insight,omitempty" form:"name=insight" multipartForm:"name=insight"`
}
