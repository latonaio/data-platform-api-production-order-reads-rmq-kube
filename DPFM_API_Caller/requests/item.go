package requests

type Item struct {
	ProductionOrder                               int     `json:"ProductionOrder"`
	ProductionOrderItem                           int     `json:"ProductionOrderItem"`
	PrecedingProductionOrderItem                  int     `json:"PrecedingProductionOrderItem"`
	FollowingProductionOrderItem                  int     `json:"FollowingProductionOrderItem"`
	SupplyChainRelationshipID                     int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipProductionPlantID      int     `json:"SupplyChainRelationshipProductionPlantID"`
	SupplyChainRelationshipDeliveryID             int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID        int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	ProductionOrderType                           string  `json:"ProductionOrderType"`
	Product                                       string  `json:"Product"`
	Buyer                                         int     `json:"Buyer"`
	Seller                                        int     `json:"Seller"`
	ProductionPlantBusinessPartner                int     `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                               string  `json:"ProductionPlant"`
	ProductionPlantStorageLocation                string  `json:"ProductionPlantStorageLocation"`
	DeliverToParty                                int     `json:"DeliverToParty"`
	DeliverToPlant                                string  `json:"DeliverToPlant"`
	DeliverToPlantStorageLocation                 string  `json:"DeliverToPlantStorageLocation"`
	DeliverFromParty                              int     `json:"DeliverFromParty"`
	DeliverFromPlant                              string  `json:"DeliverFromPlant"`
	DeliverFromPlantStorageLocation               string  `json:"DeliverFromPlantStorageLocation"`
	ProductBaseUnit                               string  `json:"ProductBaseUnit"`
	ProductProductionUnit                         string  `json:"ProductProductionUnit"`
	ProductDeliveryUnit                           string  `json:"ProductDeliveryUnit"`
	MRPArea                                       string  `json:"MRPArea"`
	MRPController                                 string  `json:"MRPController"`
	ProductionVersion                             int     `json:"ProductionVersion"`
	ProductionVersionItem                         int     `json:"ProductionVersionItem"`
	BillOfMaterial                                int     `json:"BillOfMaterial"`
	Operations                                    int     `json:"Operations"`
	ProductionOrderQuantityInBaseUnit             float32 `json:"ProductionOrderQuantityInBaseUnit"`
	ProductionOrderQuantityInProductionUnit       float32 `json:"ProductionOrderQuantityInProductionUnit"`
	ProductionOrderQuantityInDeliveryUnit         float32 `json:"ProductionOrderQuantityInDeliveryUnit"`
	ProductionOrderPlannedScrapQtyInBaseUnit      float32 `json:"ProductionOrderPlannedScrapQtyInBaseUnit"`
	ProductionOrderMinimumLotSizeQuantity         float32 `json:"ProductionOrderMinimumLotSizeQuantity"`
	ProductionOrderStandardLotSizeQuantity        float32 `json:"ProductionOrderStandardLotSizeQuantity"`
	ProductionOrderMaximumLotSizeQuantity         float32 `json:"ProductionOrderMaximumLotSizeQuantity"`
	ProductionOrderLotSizeRoundingQuantity        float32 `json:"ProductionOrderLotSizeRoundingQuantity"`
	ProductionOrderLotSizeIsFixed                 bool    `json:"ProductionOrderLotSizeIsFixed"`
	ProductionOrderPlannedStartDate               string  `json:"ProductionOrderPlannedStartDate"`
	ProductionOrderPlannedStartTime               string  `json:"ProductionOrderPlannedStartTime"`
	ProductionOrderPlannedEndDate                 string  `json:"ProductionOrderPlannedEndDate"`
	ProductionOrderPlannedEndTime                 string  `json:"ProductionOrderPlannedEndTime"`
	ProductionOrderActualReleaseDate              string  `json:"ProductionOrderActualReleaseDate"`
	ProductionOrderActualReleaseTime              string  `json:"ProductionOrderActualReleaseTime"`
	ProductionOrderActualStartDate                string  `json:"ProductionOrderActualStartDate"`
	ProductionOrderActualStartTime                string  `json:"ProductionOrderActualStartTime"`
	ProductionOrderActualEndDate                  string  `json:"ProductionOrderActualEndDate"`
	ProductionOrderActualEndTime                  string  `json:"ProductionOrderActualEndTime"`
	ConfirmedYieldQuantityInBaseUnit              float32 `json:"ConfirmedYieldQuantityInBaseUnit"`
	ConfirmedYieldQuantityInProductionUnit        float32 `json:"ConfirmedYieldQuantityInProductionUnit"`
	ScrappedQuantityInBaseUnit                    float32 `json:"ScrappedQuantityInBaseUnit"`
	PlannedOrder                                  int     `json:"PlannedOrder"`
	PlannedOrderItem                              int     `json:"PlannedOrderItem"`
	OrderID                                       int     `json:"OrderID"`
	OrderItem                                     int     `json:"OrderItem"`
	ProductIsBatchManagedInProductionPlant        bool    `json:"ProductIsBatchManagedInProductionPlant"`
	BatchMgmtPolicyInProductionOrder              string  `json:"BatchMgmtPolicyInProductionOrder"`
	ProductionOrderTargetedBatch                  string  `json:"ProductionOrderTargetedBatch"`
	ProductionOrderTargetedBatchValidityStartDate string  `json:"ProductionOrderTargetedBatchValidityStartDate"`
	ProductionOrderTargetedBatchValidityStartTime string  `json:"ProductionOrderTargetedBatchValidityStartTime"`
	ProductionOrderTargetedBatchValidityEndDate   string  `json:"ProductionOrderTargetedBatchValidityEndDate"`
	ProductionOrderTargetedBatchValidityEndTime   string  `json:"ProductionOrderTargetedBatchValidityEndTime"`
	ProductionOrderItemText                       string  `json:"ProductionOrderItemText"`
	CreationDate                                  string  `json:"CreationDate"`
	CreationTime                                  string  `json:"CreationTime"`
	LastChangeDate                                string  `json:"LastChangeDate"`
	LastChangeTime                                string  `json:"LastChangeTime"`
	IsReleased                                    bool    `json:"IsReleased"`
	IsPartiallyConfirmed                          bool    `json:"IsPartiallyConfirmed"`
	IsConfirmed                                   bool    `json:"IsConfirmed"`
	IsLocked                                      bool    `json:"IsLocked"`
	IsCancelled                                   bool    `json:"IsCancelled"`
	IsMarkedForDeletion                           bool    `json:"IsMarkedForDeletion"`
}
