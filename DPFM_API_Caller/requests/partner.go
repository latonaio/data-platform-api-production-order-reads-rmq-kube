package requests

type Partner struct {
	ProductionOrder         int     `json:"ProductionOrder"`
	PartnerFunction         string  `json:"SupplyChainRelationshipID"`
	BusinessPartner         int     `json:"SupplyChainRelationshipProductionPlantID"`
	BusinessPartnerFullName *string `json:"SupplyChainRelationshipDeliveryID"`
	BusinessPartnerName     *string `json:"SupplyChainRelationshipDeliveryPlantID"`
	Organization            *string `json:"ProductionOrderType"`
	Country                 *string `json:"Product"`
	Language                *string `json:"Buyer"`
	Currency                *string `json:"Seller"`
	ExternalDocumentID      *string `json:"OwnerProductionPlantBusinessPartner"`
	AddressID               *int    `json:"OwnerProductionPlant"`
	EmailAddress            *string `json:"OwnerProductionPlantStorageLocation"`
}
