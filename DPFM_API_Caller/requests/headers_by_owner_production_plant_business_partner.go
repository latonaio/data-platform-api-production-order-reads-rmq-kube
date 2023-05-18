package requests

type HeadersByOwnerProductionPlantBusinessPartner struct {
	ProductionOrder                         int     `json:"ProductionOrder"`
	MRPArea                                 *string `json:"MRPArea"`
	OwnerProductionPlantBusinessPartnerName *string `json:"OwnerProductionPlantBusinessPartnerName"`
}
