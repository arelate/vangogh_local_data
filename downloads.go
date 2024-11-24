package vangogh_local_data

import (
	"fmt"
	"github.com/arelate/southern_light/gog_integration"
	"github.com/boggydigital/kevlar"
	"github.com/boggydigital/nod"
	"log"
	"math"
	"path"
	"strconv"
	"strings"
)

const (
	mbSuffix  = "MB"
	gbSuffix  = "GB"
	bytesInGB = 1024 * 1024 * 1024
	bytesInMB = 1024 * 1024
	patchStr  = "patch"
)

type Download struct {
	ManualUrl      string
	ProductTitle   string
	Name           string
	Version        string
	Date           string
	OS             OperatingSystem
	LanguageCode   string
	Type           DownloadType
	EstimatedBytes int
}

func convertManualDownload(
	productTitle string,
	mdl *gog_integration.ManualDownload,
	dt DownloadType,
	os OperatingSystem,
	langCode string) Download {
	return Download{
		ManualUrl:      mdl.ManualUrl,
		ProductTitle:   productTitle,
		Name:           mdl.Name,
		Version:        mdl.Version,
		Date:           mdl.Date,
		OS:             os,
		LanguageCode:   langCode,
		Type:           dt,
		EstimatedBytes: SizeToEstimatedBytes(mdl.Size),
	}
}

func convertToBytes(size string, suffix string, bytesInUnit int) int {
	sizeStr := strings.TrimSuffix(size, " "+suffix)
	sz, err := strconv.ParseFloat(sizeStr, 0)
	if err != nil {
		log.Printf("error parsing size: %s", size)
		return 0
	}
	return int(sz * float64(bytesInUnit))
}

func SizeToEstimatedBytes(size string) int {
	if strings.HasSuffix(size, gbSuffix) {
		return convertToBytes(size, gbSuffix, bytesInGB)
	} else if strings.HasSuffix(size, mbSuffix) {
		return convertToBytes(size, mbSuffix, bytesInMB)
	} else {
		log.Printf("unknown size format: %s", size)
		return 0
	}
}

func (dl *Download) String() string {
	switch dl.Type {
	case Installer:
		fallthrough
	case Movie:
		fallthrough
	case DLC:
		name := dl.Name
		if !strings.HasPrefix(dl.Name, dl.ProductTitle) {
			name = fmt.Sprintf("%s %s", dl.ProductTitle, dl.Name)
		}
		return fmt.Sprintf("%s %s (%s, %s)", name, dl.Version, dl.OS, dl.LanguageCode)
	case Extra:
		return strings.Title(dl.Name)
	default:
		return ""
	}
}

type DownloadsList []Download

func FromDetails(det *gog_integration.Details, rdx kevlar.ReadableRedux) (DownloadsList, error) {
	return fromGameDetails(det, rdx)
}

func fromGameDetails(det *gog_integration.Details, rdx kevlar.ReadableRedux) (DownloadsList, error) {
	dlList := make(DownloadsList, 0)

	if det == nil {
		return dlList, fmt.Errorf("details are nil")
	}

	installerDls, err := convertGameDetails(det, rdx, Installer)
	if err != nil {
		return dlList, err
	}
	dlList = append(dlList, installerDls...)

	for _, dlc := range det.DLCs {
		dlcDls, err := convertGameDetails(&dlc, rdx, DLC)
		if err != nil {
			return dlList, err
		}
		dlList = append(dlList, dlcDls...)
	}

	return dlList, nil
}

func convertGameDetails(det *gog_integration.Details, rdx kevlar.ReadableRedux, dt DownloadType) (DownloadsList, error) {

	dlList := make(DownloadsList, 0)

	if err := rdx.MustHave(NativeLanguageNameProperty); err != nil {
		return dlList, err
	}

	downloads, err := det.GetGameDownloads()
	if err != nil {
		return dlList, err
	}

	for _, dl := range downloads {

		langCodes := rdx.Match(
			map[string][]string{NativeLanguageNameProperty: {dl.Language}},
			kevlar.FullMatch)
		if len(langCodes) != 1 {
			return dlList, fmt.Errorf("invalid native language %s", dl.Language)
		}

		langCode := ""
		for _, lc := range langCodes {
			langCode = lc
		}

		for _, winDl := range dl.Windows {
			dlList = append(dlList, convertManualDownload(det.Title, &winDl, dt, Windows, langCode))
		}
		for _, macDl := range dl.Mac {
			dlList = append(dlList, convertManualDownload(det.Title, &macDl, dt, MacOS, langCode))
		}
		for _, linDl := range dl.Linux {
			dlList = append(dlList, convertManualDownload(det.Title, &linDl, dt, Linux, langCode))
		}
	}

	for _, extraDl := range det.Extras {
		dlList = append(dlList, convertManualDownload(det.Title, &extraDl, Extra, AnyOperatingSystem, ""))
	}

	return dlList, nil
}

func (list DownloadsList) Only(
	operatingSystems []OperatingSystem,
	downloadTypes []DownloadType,
	langCodes []string,
	excludePatches bool) DownloadsList {
	osSet := make(map[OperatingSystem]bool)
	for _, os := range operatingSystems {
		if os == AnyOperatingSystem {
			for _, aos := range AllOperatingSystems() {
				osSet[aos] = true
			}
			break
		}
		osSet[os] = true
	}
	dtSet := make(map[DownloadType]bool)
	for _, dt := range downloadTypes {
		if dt == AnyDownloadType {
			for _, adt := range AllDownloadTypes() {
				dtSet[adt] = true
			}
			break
		}
		dtSet[dt] = true
	}
	langSet := make(map[string]bool)
	for _, lc := range langCodes {
		langSet[lc] = true
	}
	matchingList := make(DownloadsList, 0)
	for _, dl := range list {
		if dl.OS != AnyOperatingSystem &&
			!osSet[dl.OS] {
			continue
		}

		if dl.Type != AnyDownloadType &&
			!dtSet[dl.Type] {
			continue
		}

		if dl.LanguageCode != "" &&
			len(langSet) > 0 &&
			!langSet[dl.LanguageCode] {
			continue
		}

		if excludePatches {
			if base := path.Base(dl.ManualUrl); strings.Contains(base, patchStr) {
				continue
			}
		}

		matchingList = append(matchingList, dl)
	}
	return matchingList
}

func (list DownloadsList) TotalBytesEstimate() int {
	totalBytes := 0
	for _, dl := range list {
		totalBytes += dl.EstimatedBytes
	}
	return totalBytes
}

func (list DownloadsList) TotalGBsEstimate() float64 {
	return float64(list.TotalBytesEstimate()) / math.Pow(1000, 3)
}

type DownloadsListProcessor interface {
	Process(id string, slug string, downloadsList DownloadsList) error
}

func MapDownloads(
	ids []string,
	rdx kevlar.ReadableRedux,
	operatingSystems []OperatingSystem,
	downloadTypes []DownloadType,
	langCodes []string,
	excludePatches bool,
	dlProcessor DownloadsListProcessor,
	tpw nod.TotalProgressWriter) error {

	if dlProcessor == nil {
		return fmt.Errorf("vangogh_downloads: map downloads list processor is nil")
	}

	if err := rdx.MustHave(
		SlugProperty,
		NativeLanguageNameProperty); err != nil {
		return err
	}

	vrDetails, err := NewProductReader(Details)
	if err != nil {
		return err
	}

	tpw.TotalInt(len(ids))

	for _, id := range ids {

		detSlug, ok := rdx.GetLastVal(SlugProperty, id)

		has, err := vrDetails.Has(id)
		if err != nil {
			return err
		}

		if !has || !ok {
			tpw.Increment()
			continue
		}

		det, err := vrDetails.Details(id)
		if err != nil {
			return err
		}

		downloads, err := FromDetails(det, rdx)
		if err != nil {
			return err
		}

		filteredDownloads := make([]Download, 0)

		for _, dl := range downloads.Only(operatingSystems, downloadTypes, langCodes, excludePatches) {
			//some manualUrls have "0 MB" specified as size and don't seem to be used to create user clickable links.
			//resolving such manualUrls leads to an empty filename
			//given they don't contribute anything to download, size or validate commands - we're filtering them
			if dl.EstimatedBytes == 0 {
				continue
			}
			filteredDownloads = append(filteredDownloads, dl)
		}

		if det.IsPreOrder && len(filteredDownloads) == 0 {
			fmt.Printf("(%s is a pre-order and has no downloads)\n", id)
			return nil
		}

		// already checked for nil earlier in the function
		if err := dlProcessor.Process(
			id,
			detSlug,
			filteredDownloads); err != nil {
			return err
		}

		tpw.Increment()
	}

	return nil
}
