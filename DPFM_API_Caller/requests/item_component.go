package requests

type ItemComponent struct {
	ProductionOrder									int			`json:"ProductionOrder"`
	ProductionOrderItem								int			`json:"ProductionOrderItem"`
	BillOfMaterial									int			`json:"BillOfMaterial"`
	BillOfMaterialItem								int			`json:"BillOfMaterialItem"`
	SupplyChainRelationshipID						*int		`json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID				*int		`json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID			*int		`json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipStockConfPlantID			*int		`json:"SupplyChainRelationshipStockConfPlantID"`
	ProductionPlantBusinessPartner					*int		`json:"ProductionPlantBusinessPartner"`
	ProductionPlant									*string		`json:"ProductionPlant"`
	ComponentProduct								*string		`json:"ComponentProduct"`
	ComponentProductBuyer							*int		`json:"ComponentProductBuyer"`
	ComponentProductSeller							*int		`json:"ComponentProductSeller"`
	ComponentProductDeliverToParty					*int		`json:"ComponentProductDeliverToParty"`
	ComponentProductDeliverToPlant					*string		`json:"ComponentProductDeliverToPlant"`
	ComponentProductDeliverFromParty				*int		`json:"ComponentProductDeliverFromParty"`
	ComponentProductDeliverFromPlant				*string		`json:"ComponentProductDeliverFromPlant"`
	ComponentProductBaseUnit						*string		`json:"ComponentProductBaseUnit"`
	ComponentProductDeliveryUnit					*string		`json:"ComponentProductDeliveryUnit"`
	ComponentProductRequirementDate					*string		`json:"ComponentProductRequirementDate"`
	ComponentProductRequirementTime					*string		`json:"ComponentProductRequirementTime"`
	ComponentProductRequiredQuantityInBaseUnit		*float32	`json:"ComponentProductRequiredQuantityInBaseUnit"`
	ComponentProductRequiredQuantityInDeliveryUnit	*float32	`json:"ComponentProductRequiredQuantityInDeliveryUnit"`
	ComponentProductPlannedScrapInPercent			*float32	`json:"ComponentProductPlannedScrapInPercent"`
	ComponentProductIsMarkedForBackflush			*bool		`json:"ComponentProductIsMarkedForBackflush"`
	BillOfMaterialItemText							*string		`json:"BillOfMaterialItemText"`
	StockConfirmationBusinessPartner				*int		`json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant							*string		`json:"StockConfirmationPlant"`
	StockConfirmationPlantStorageLocation			*string		`json:"StockConfirmationPlantStorageLocation"`
	MRPArea											*string		`json:"MRPArea"`
	MRPController									*string		`json:"MRPController"`
	ProductionVersion								*int		`json:"ProductionVersion"`
	ProductionVersionItem							*int		`json:"ProductionVersionItem"`
	PlannedOrder									*int		`json:"PlannedOrder"`
	PlannedOrderItem								*int		`json:"PlannedOrderItem"`
	ComponentProductBatch							*string		`json:"ComponentProductBatch"`
	ComponentProductBatchValidityStartDate			*string		`json:"ComponentProductBatchValidityStartDate"`
	ComponentProductBatchValidityStartTime			*string		`json:"ComponentProductBatchValidityStartTime"`
	ComponentProductBatchValidityEndDate			*string		`json:"ComponentProductBatchValidityEndDate"`
	ComponentProductBatchValidityEndTime			*string		`json:"ComponentProductBatchValidityEndTime"`
	ComponentProductCostingPolicy					*string		`json:"ComponentProductCostingPolicy"`
	ComponentProductPriceUnitQty					*int		`json:"ComponentProductPriceUnitQty"`
	ComponentProductStandardPrice					*float32	`json:"ComponentProductStandardPrice"`
	ComponentProductMovingAveragePrice				*float32	`json:"ComponentProductMovingAveragePrice"`
	ComponentProductWithdrawnQuantity				*float32	`json:"ComponentProductWithdrawnQuantity"`
	CreationDate									*string		`json:"CreationDate"`
	CreationTime									*string		`json:"CreationTime"`
	LastChangeDate									*string		`json:"LastChangeDate"`
	LastChangeTime									*string		`json:"LastChangeTime"`
	ComponentProductAvailabilityIsNotChecked		*bool		`json:"ComponentProductAvailabilityIsNotChecked"`
	IsReleased										*bool		`json:"IsReleased"`
	IsLocked										*bool		`json:"IsLocked"`
	IsCancelled										*bool		`json:"IsCancelled"`
	IsMarkedForDeletion								*bool		`json:"IsMarkedForDeletion"`
}
