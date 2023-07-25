package requests

type ItemOperation struct {
	ProductionOrder										int			`json:"ProductionOrder"`
	ProductionOrderItem									int			`json:"ProductionOrderItem"`
	Operations											int			`json:"Operations"`
	OperationsItem										int			`json:"OperationsItem"`
	OperationID											int			`json:"OperationID"`
	SupplyChainRelationshipID							int			`json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID					int			`json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID				int			`json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipProductionPlantID			int			`json:"SupplyChainRelationshipProductionPlantID"`
	Product												string		`json:"Product"`
	Buyer												int			`json:"Buyer"`
	Seller												int			`json:"Seller"`
	DeliverToParty										int			`json:"DeliverToParty"`
	DeliverToPlant										string		`json:"DeliverToPlant"`
	DeliverFromParty									int			`json:"DeliverFromParty"`
	DeliverFromPlant									string		`json:"DeliverFromPlant"`
	ProductionPlantBusinessPartner						int			`json:"ProductionPlantBusinessPartner"`
	ProductionPlant										string		`json:"ProductionPlant"`
	MRPArea												*string		`json:"MRPArea"`
	MRPController										*string		`json:"MRPController"`
	ProductionVersion									*int		`json:"ProductionVersion"`
	ProductionVersionItem								*int		`json:"ProductionVersionItem"`
	Sequence											int			`json:"Sequence"`
	SequenceText										*string		`json:"SequenceText"`
	OperationText										string		`json:"OperationText"`
	ProductBaseUnit										string		`json:"ProductBaseUnit"`
	ProductProductionUnit								string		`json:"ProductProductionUnit"`
	ProductOperationUnit								string		`json:"ProductOperationUnit"`
	ProductDeliveryUnit									string		`json:"ProductDeliveryUnit"`
	StandardLotSizeQuantity								float32		`json:"StandardLotSizeQuantity"`
	MinimumLotSizeQuantity								float32		`json:"MinimumLotSizeQuantity"`
	MaximumLotSizeQuantity								float32		`json:"MaximumLotSizeQuantity"`
	OperationPlannedQuantityInBaseUnit					float32		`json:"OperationPlannedQuantityInBaseUnit"`
	OperationPlannedQuantityInProductionUnit			float32		`json:"OperationPlannedQuantityInProductionUnit"`
	OperationPlannedQuantityInOperationUnit				float32		`json:"OperationPlannedQuantityInOperationUnit"`
	OperationPlannedQuantityInDeliveryUnit				float32		`json:"OperationPlannedQuantityInDeliveryUnit"`
	OperationPlannedScrapInPercent						*float32	`json:"OperationPlannedScrapInPercent"`
	ResponsiblePlannerGroup								*string		`json:"ResponsiblePlannerGroup"`
	PlainLongText										*string		`json:"PlainLongText"`
	WorkCenter											int			`json:"WorkCenter"`
	CapacityCategory									*string		`json:"CapacityCategory"`
	OperationCostingRelevancyType						*string		`json:"OperationCostingRelevancyType"`
	OperationSetupType									*string		`json:"OperationSetupType"`
	OperationSetupGroupCategory							*string		`json:"OperationSetupGroupCategory"`
	OperationSetupGroup									*string		`json:"OperationSetupGroup"`
	MaximumWaitDuration									*float32	`json:"MaximumWaitDuration"`
	StandardWaitDuration								*float32	`json:"StandardWaitDuration"`
	MinimumWaitDuration									*float32	`json:"MinimumWaitDuration"`
	WaitDurationUnit									*string		`json:"WaitDurationUnit"`
	MaximumQueDuration									*float32	`json:"MaximumQueDuration"`
	StandardQueueDuration								*float32	`json:"StandardQueueDuration"`
	MinimumQueueDuration								*float32	`json:"MinimumQueueDuration"`
	QueDurationUnit										*string		`json:"QueDurationUnit"`
	MaximumMoveDuration									*float32	`json:"MaximumMoveDuration"`
	StandardMoveDuration								*float32	`json:"StandardMoveDuration"`
	MinimumMoveDuration									*float32	`json:"MinimumMoveDuration"`
	MoveDurationUnit									*string		`json:"MoveDurationUnit"`
	StandardDeliveryDuration							*float32	`json:"StandardDeliveryDuration"`
	StandardDeliveryDurationUnit						*string		`json:"StandardDeliveryDurationUnit"`
	CostElement											*string		`json:"CostElement"`
	OperationErlstSchedldExecStrtDte					*string		`json:"OperationErlstSchedldExecStrtDte"`
	OperationErlstSchedldExecStrtTme					*string		`json:"OperationErlstSchedldExecStrtTme"`
	OperationErlstSchedldExecEndDte						*string		`json:"OperationErlstSchedldExecEndDte"`
	OperationErlstSchedldExecEndTme						*string		`json:"OperationErlstSchedldExecEndTme"`
	OperationActualExecutionStartDate					*string		`json:"OperationActualExecutionStartDate"`
	OperationActualExecutionStartTime					*string		`json:"OperationActualExecutionStartTime"`
	OperationActualExecutionEndDate						*string		`json:"OperationActualExecutionEndDate"`
	OperationActualExecutionEndTime						*string		`json:"OperationActualExecutionEndTime"`
	OperationConfirmedYieldQuantityInBaseUnit			*float32	`json:"OperationConfirmedYieldQuantityInBaseUnit"`
	OperationConfirmedYieldQuantityInProductionUnit		*float32	`json:"OperationConfirmedYieldQuantityInProductionUnit"`
	OperationConfirmedYieldQuantityInOperationUnit		*float32	`json:"OperationConfirmedYieldQuantityInOperationUnit"`
	CreationDate										string		`json:"CreationDate"`
	CreationTime										string		`json:"CreationTime"`
	LastChangeDate										string		`json:"LastChangeDate"`
	LastChangeTime										string		`json:"LastChangeTime"`
	IsReleased											*bool		`json:"IsReleased"`
	IsPartiallyConfirmed								*bool		`json:"IsPartiallyConfirmed"`
	IsConfirmed											*bool		`json:"IsConfirmed"`
	IsLocked											*bool		`json:"IsLocked"`
	IsCancelled											*bool		`json:"IsCancelled"`
	IsMarkedForDeletion									*bool		`json:"IsMarkedForDeletion"`
}
