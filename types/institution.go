package types

type Institution struct {
	ISPBCode       string `json:"ispb_code"`
	NumberCode     string `json:"number_code"`
	Name           string `json:"name"`
	ShortName      string `json:"short_name"`
	SPIParticipant bool   `json:"spi_participant"`
}
