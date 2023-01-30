package data

import (
	"strings"
	"time"
)

type MinimalCVEData struct {
	ID               string  `json:"id"`
	Published        string  `json:"published" gorm:"index"`
	LastModified     string  `json:"lastModified"`
	Status           string  `json:"status"`
	Description      string  `json:"description"`
	CVSSVector       string  `json:"cvssVector"`
	CVSSBaseScore    float32 `json:"cvssBaseScore"`
	CVSSBaseSeverity string  `json:"cvssBaseSeverity"`
	CVSSSource       string  `json:"cvssSource"`
	// CWEs             []string
	// CPEs             []string
}

type CVEStats struct {
	NumLow             int
	NumMedium          int
	NumHigh            int
	NumCritical        int
	NumUnknownSeverity int
}

type LastNVDUpdate struct {
	LastUpdate time.Time
	Version    int
}

type NVDResponse struct {
	ResultsPerPage  int       `json:"resultsPerPage"`
	StartIndex      int       `json:"startIndex"`
	TotalResults    int       `json:"totalResults"`
	Vulnerabilities []NVDItem `json:"vulnerabilities"`
}

func (response *NVDResponse) ConvertToMinimal() []MinimalCVEData {
	minimal := make([]MinimalCVEData, response.ResultsPerPage)
	for i := 0; i < response.ResultsPerPage; i++ {
		minimal[i] = response.Vulnerabilities[i].CVE.ToMinimalCVEData()
	}

	return minimal
}

type NVDItem struct {
	CVE NVDCVE `json:"cve"`
}

type NVDCVE struct {
	ID           string `json:"id"`
	Published    string `json:"published"`
	LastModified string `json:"lastModified"`
	VulnStatus   string `json:"vulnStatus"`
	Descriptions []struct {
		Lang  string `json:"lang"`
		Value string `json:"value"`
	} `json:"descriptions"`
	Metrics struct {
		CvssMetricV31 []struct {
			Source   string `json:"source"`
			Type     string `json:"type"`
			CvssData struct {
				Version      string  `json:"version"`
				VectorString string  `json:"vectorString"`
				BaseScore    float32 `json:"baseScore"`
				BaseSeverity string  `json:"baseSeverity"`
			} `json:"cvssData"`
		} `json:"cvssMetricV31"`
		CvssMetricV30 []struct {
			Source   string `json:"source"`
			Type     string `json:"type"`
			CvssData struct {
				Version      string  `json:"version"`
				VectorString string  `json:"vectorString"`
				BaseScore    float32 `json:"baseScore"`
				BaseSeverity string  `json:"baseSeverity"`
			} `json:"cvssData"`
		} `json:"cvssMetricV30"`
		CvssMetricV2 []struct {
			Source   string `json:"source"`
			Type     string `json:"type"`
			CvssData struct {
				Version      string  `json:"version"`
				VectorString string  `json:"vectorString"`
				BaseScore    float32 `json:"baseScore"`
			} `json:"cvssData"`
			BaseSeverity string `json:"baseSeverity"`
		} `json:"cvssMetricV2"`
	} `json:"metrics"`
	Weaknesses []struct {
		Source      string `json:"source"`
		Type        string `json:"type"`
		Description []struct {
			Lang  string `json:"lang"`
			Value string `json:"value"`
		} `json:"description"`
	} `json:"weaknesses"`
	Configurations []struct {
		Nodes []struct {
			Operator string `json:"operator"`
			Negate   bool   `json:"negate"`
			CpeMatch []struct {
				Vulnerable      bool   `json:"vulnerable"`
				Criteria        string `json:"criteria"`
				MatchCriteriaID string `json:"matchCriteriaId"`
			} `json:"cpeMatch"`
		} `json:"nodes"`
	} `json:"configurations"`
}

func (nvdCVE NVDCVE) ToMinimalCVEData() MinimalCVEData {
	var description string
	for _, v := range nvdCVE.Descriptions {
		if strings.ToLower(v.Lang) == "en" {
			description = v.Value
			break
		}
	}

	var cvssVector string
	var baseScore float32
	var baseSeverity string
	var cvssSource string
	if len(nvdCVE.Metrics.CvssMetricV31) > 0 {
		cvssVector = nvdCVE.Metrics.CvssMetricV31[0].CvssData.VectorString
		baseScore = nvdCVE.Metrics.CvssMetricV31[0].CvssData.BaseScore
		baseSeverity = nvdCVE.Metrics.CvssMetricV31[0].CvssData.BaseSeverity
		cvssSource = nvdCVE.Metrics.CvssMetricV31[0].Source
	} else if len(nvdCVE.Metrics.CvssMetricV30) > 0 {
		cvssVector = nvdCVE.Metrics.CvssMetricV30[0].CvssData.VectorString
		baseScore = nvdCVE.Metrics.CvssMetricV30[0].CvssData.BaseScore
		baseSeverity = nvdCVE.Metrics.CvssMetricV30[0].CvssData.BaseSeverity
		cvssSource = nvdCVE.Metrics.CvssMetricV30[0].Source
	} else if len(nvdCVE.Metrics.CvssMetricV2) > 0 {
		cvssVector = nvdCVE.Metrics.CvssMetricV2[0].CvssData.VectorString
		baseScore = nvdCVE.Metrics.CvssMetricV2[0].CvssData.BaseScore
		baseSeverity = nvdCVE.Metrics.CvssMetricV2[0].BaseSeverity
		cvssSource = nvdCVE.Metrics.CvssMetricV2[0].Source
	}

	return MinimalCVEData{
		ID:               nvdCVE.ID,
		Published:        nvdCVE.Published,
		LastModified:     nvdCVE.LastModified,
		Status:           nvdCVE.VulnStatus,
		Description:      description,
		CVSSVector:       cvssVector,
		CVSSBaseScore:    baseScore,
		CVSSBaseSeverity: baseSeverity,
		CVSSSource:       cvssSource,
	}
}
