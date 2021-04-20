package helper

type MetaData struct {
	Accession string `json:"accession"`
	Name      string `json:"name"`
	Value     string `json:"value"`
}

// http://wwwdev.ebi.ac.uk/pride/proxi/archive/v0.1/spectra?resultType=full&usi=mzspec%3APXD000966%3ACPTAC_CompRef_00_iTRAQ_05_2Feb12_Cougar_11-10-09.mzML%3Ascan%3A12298%3A%5BiTRAQ4plex%5D-LHFFM%5BOxidation%5DPGFAPLTSR%2F2
type PRIDE []struct {
	Status      string     `json:"status"`
	Usi         string     `json:"usi"`
	Mzs         []float64  `json:"mzs"`
	Intensities []float64  `json:"intensities"`
	Attributes  []MetaData `json:"attributes"`
}

// http://proteomecentral.proteomexchange.org/api/proxi/v0.1/spectra?resultType=full&usi=mzspec:PXD000561:Adult_Frontalcortex_bRP_Elite_85_f09:scan:17555:VLHPLEGAVVIIFK/2
type ProteomeCentral []struct {
	Attributes  []MetaData `json:"attributes"`
	Intensities []float64  `json:"intensities"`
	Mzs         []float64  `json:"mzs"`
	PrivateMetaData
}
type PrivateMetaData struct {
	PrecursorCharge int
}

// http://www.peptideatlas.org/api/proxi/v0.1/spectra?resultType=full&usi=mzspec%3APXD000966%3ACPTAC_CompRef_00_iTRAQ_05_2Feb12_Cougar_11-10-09.mzML%3Ascan%3A12298%3A%5BiTRAQ4plex%5D-LHFFM%5BOxidation%5DPGFAPLTSR%2F2
type PeptideAtlas []struct {
	Attributes  []MetaData `json:"attributes"`
	Intensities []float64  `json:"intensities"`
	Mzs         []float64  `json:"mzs"`
}

// http://massive.ucsd.edu/ProteoSAFe/proxi/v0.1/spectra?resultType=full&usi=mzspec:PXD000561:Adult_Frontalcortex_bRP_Elite_85_f09:scan:17555:VLHPLEGAVVIIFK/2
type MassIVE []struct {
	Status      string   `json:"status"`
	Usi         string   `json:"usi"`
	Intensities []string `json:"intensities"`
	Mzs         []string `json:"mzs"`
}
