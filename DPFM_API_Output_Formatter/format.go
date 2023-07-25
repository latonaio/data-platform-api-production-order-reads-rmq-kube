package dpfm_api_output_formatter

import (
	"database/sql"
	"fmt"
)

func ConvertToHeader(rows *sql.Rows) (*[]Header, error) {
	defer rows.Close()
	header := make([]Header, 0)
	i := 0
	for rows.Next() {
		i++
		pm := Header{}
		err := rows.Scan(
			&pm.ProductionOrder,
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipProductionPlantID,
			&pm.SupplyChainRelationshipDeliveryID,
			&pm.SupplyChainRelationshipDeliveryPlantID,
			&pm.ProductionOrderType,
			&pm.Product,
			&pm.Buyer,
			&pm.Seller,
			&pm.OwnerProductionPlantBusinessPartner,
			&pm.OwnerProductionPlant,
			&pm.OwnerProductionPlantStorageLocation,
			&pm.DepartureDeliverFromParty,
			&pm.DepartureDeliverFromPlant,
			&pm.DepartureDeliverFromPlantStorageLocation,
			&pm.DestinationDeliverToParty,
			&pm.DestinationDeliverToPlant,
			&pm.DestinationDeliverToPlantStorageLocation,
			&pm.ProductBaseUnit,
			&pm.MRPArea,
			&pm.MRPController,
			&pm.ProductionVersion,
			&pm.BillOfMaterial,
			&pm.Operations,
			&pm.ProductionOrderQuantityInBaseUnit,
			&pm.ProductionOrderQuantityInDepartureProductionUnit,
			&pm.ProductionOrderQuantityInDestinationProductionUnit,
			&pm.ProductionOrderQuantityInDepartureDeliveryUnit,
			&pm.ProductionOrderQuantityInDestinationDeliveryUnit,
			&pm.ProductionOrderDepartureProductionUnit,
			&pm.ProductionOrderDestinationProductionUnit,
			&pm.ProductionOrderDepartureDeliveryUnit,
			&pm.ProductionOrderDestinationDeliveryUnit,
			&pm.ProductionOrderPlannedScrapQtyInBaseUnit,
			&pm.ProductionOrderPlannedStartDate,
			&pm.ProductionOrderPlannedStartTime,
			&pm.ProductionOrderPlannedEndDate,
			&pm.ProductionOrderPlannedEndTime,
			&pm.ProductionOrderActualReleaseDate,
			&pm.ProductionOrderActualReleaseTime,
			&pm.ProductionOrderActualStartDate,
			&pm.ProductionOrderActualStartTime,
			&pm.ProductionOrderActualEndDate,
			&pm.ProductionOrderActualEndTime,
			&pm.PlannedOrder,
			&pm.OrderID,
			&pm.OrderItem,
			&pm.ProductionOrderHeaderText,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.IsReleased,
			&pm.IsPartiallyConfirmed,
			&pm.IsConfirmed,
			&pm.IsLocked,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &header, err
		}
		header = append(header, pm)
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &header, nil
	}

	return &header, nil
}

func ConvertToItem(rows *sql.Rows) (*[]Item, error) {
	defer rows.Close()
	item := make([]Item, 0)
	i := 0
	for rows.Next() {
		i++
		pm := Item{}
		err := rows.Scan(
			&pm.ProductionOrder,
			&pm.ProductionOrderItem,
			&pm.PrecedingProductionOrderItem,
			&pm.FollowingProductionOrderItem,
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipProductionPlantID,
			&pm.SupplyChainRelationshipDeliveryID,
			&pm.SupplyChainRelationshipDeliveryPlantID,
			&pm.ProductionOrderType,
			&pm.Product,
			&pm.ProductProductionUnit,
			&pm.Seller,
			&pm.ProductionPlantBusinessPartner,
			&pm.ProductionPlant,
			&pm.ProductionPlantStorageLocation,
			&pm.DeliverToParty,
			&pm.DeliverToPlant,
			&pm.DeliverToPlantStorageLocation,
			&pm.DeliverFromParty,
			&pm.DeliverFromPlant,
			&pm.DeliverFromPlantStorageLocation,
			&pm.MRPArea,
			&pm.MRPController,
			&pm.ProductBaseUnit,
			&pm.ProductProductionUnit,
			&pm.ProductDeliveryUnit,
			&pm.ProductionVersion,
			&pm.ProductionVersionItem,
			&pm.BillOfMaterial,
			&pm.Operations,
			&pm.ProductionOrderQuantityInBaseUnit,
			&pm.ProductionOrderQuantityInProductionUnit,
			&pm.ProductionOrderQuantityInDeliveryUnit,
			&pm.ProductionOrderPlannedScrapQtyInBaseUnit,
			&pm.ProductionOrderMinimumLotSizeQuantity,
			&pm.ProductionOrderStandardLotSizeQuantity,
			&pm.ProductionOrderMaximumLotSizeQuantity,
			&pm.ProductionOrderLotSizeRoundingQuantity,
			&pm.ProductionOrderLotSizeIsFixed,
			&pm.ProductionOrderPlannedStartDate,
			&pm.ProductionOrderPlannedStartTime,
			&pm.ProductionOrderPlannedEndDate,
			&pm.ProductionOrderPlannedEndTime,
			&pm.ProductionOrderActualReleaseDate,
			&pm.ProductionOrderActualReleaseTime,
			&pm.ProductionOrderActualStartDate,
			&pm.ProductionOrderActualStartTime,
			&pm.ProductionOrderActualEndDate,
			&pm.ProductionOrderActualEndTime,
			&pm.ConfirmedYieldQuantityInBaseUnit,
			&pm.ConfirmedYieldQuantityInProductionUnit,
			&pm.ScrappedQuantityInBaseUnit,
			&pm.PlannedOrder,
			&pm.PlannedOrderItem,
			&pm.OrderID,
			&pm.OrderItem,
			&pm.ProductIsBatchManagedInProductionPlant,
			&pm.BatchMgmtPolicyInProductionOrder,
			&pm.ProductionOrderTargetedBatch,
			&pm.ProductionOrderTargetedBatchValidityStartDate,
			&pm.ProductionOrderTargetedBatchValidityStartTime,
			&pm.ProductionOrderTargetedBatchValidityEndDate,
			&pm.ProductionOrderTargetedBatchValidityEndTime,
			&pm.ProductionOrderItemText,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.ComponentProductAvailabilityIsNotChecked,
			&pm.IsReleased,
			&pm.IsPartiallyConfirmed,
			&pm.IsConfirmed,
			&pm.IsLocked,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &item, err
		}
		item = append(item, pm)
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &item, nil
	}

	return &item, nil
}

func ConvertToItemComponent(rows *sql.Rows) (*[]ItemComponent, error) {
	defer rows.Close()
	itemComponent := make([]ItemComponent, 0)
	i := 0
	for rows.Next() {
		i++
		pm := ItemComponent{}
		err := rows.Scan(
			&pm.ProductionOrder,
			&pm.ProductionOrderItem,
			&pm.BillOfMaterial,
			&pm.BillOfMaterialItem,
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipDeliveryID,
			&pm.SupplyChainRelationshipDeliveryPlantID,
			&pm.SupplyChainRelationshipStockConfPlantID,
			&pm.ProductionPlantBusinessPartner,
			&pm.ProductionPlant,
			&pm.MRPArea,
			&pm.MRPController,
			&pm.ProductionVersion,
			&pm.ProductionVersionItem,
			&pm.ComponentProduct,
			&pm.ComponentProductBuyer,
			&pm.ComponentProductSeller,
			&pm.ComponentProductDeliverToParty,
			&pm.ComponentProductDeliverToPlant,
			&pm.ComponentProductDeliverFromParty,
			&pm.ComponentProductDeliverFromPlant,
			&pm.ComponentProductBaseUnit,
			&pm.ComponentProductDeliveryUnit,
			&pm.ComponentProductRequirementDate,
			&pm.ComponentProductRequirementTime,
			&pm.ComponentProductRequiredQuantityInBaseUnit,
			&pm.ComponentProductRequiredQuantityInDeliveryUnit,
			&pm.ComponentProductPlannedScrapInPercent,
			&pm.ComponentProductIsMarkedForBackflush,
			&pm.StockConfirmationBusinessPartner,
			&pm.StockConfirmationPlant,
			&pm.StockConfirmationPlantStorageLocation,
			&pm.PlannedOrder,
			&pm.PlannedOrderItem,
			&pm.BillOfMaterialItemText,
			&pm.ComponentProductBatch,
			&pm.ComponentProductBatchValidityStartDate,
			&pm.ComponentProductBatchValidityStartTime,
			&pm.ComponentProductBatchValidityEndDate,
			&pm.ComponentProductBatchValidityEndTime,
			&pm.ComponentProductCostingPolicy,
			&pm.ComponentProductPriceUnitQty,
			&pm.ComponentProductStandardPrice,
			&pm.ComponentProductMovingAveragePrice,
			&pm.ComponentProductWithdrawnQuantity,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.ComponentProductAvailabilityIsNotChecked,
			&pm.IsReleased,
			&pm.IsLocked,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &itemComponent, err
		}
		itemComponent = append(itemComponent, pm)
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &itemComponent, nil
	}

	return &itemComponent, nil
}

func ConvertToItemComponentDeliveryScheduleLine(rows *sql.Rows) (*[]ItemComponentDeliveryScheduleLine, error) {
	defer rows.Close()
	itemComponentDeliveryScheduleLine := make([]ItemComponentDeliveryScheduleLine, 0)
	i := 0
	for rows.Next() {
		i++
		pm := ItemComponentDeliveryScheduleLine{}
		err := rows.Scan(
			&pm.ProductionOrder,
			&pm.ProductionOrderItem,
			&pm.BillOfMaterial,
			&pm.BillOfMaterialItem,
			&pm.ScheduleLine,
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipStockConfPlantID,
			&pm.ComponentProduct,
			&pm.StockConfirmationBusinessPartner,
			&pm.StockConfirmationPlant,
			&pm.StockConfirmationPlantTimeZone,
			&pm.ComponentProductBatch,
			&pm.ComponentProductBatchValidityStartDate,
			&pm.ComponentProductBatchValidityStartTime,
			&pm.ComponentProductBatchValidityEndDate,
			&pm.ComponentProductBatchValidityEndTime,
			&pm.RequestedDeliveryDate,
			&pm.RequestedDeliveryTime,
			&pm.ConfirmedDeliveryDate,
			&pm.ConfirmedDeliveryTime,
			&pm.OriginalRequiredQuantityInBaseUnit,
			&pm.ConfirmedQuantityByPDTAvailCheckInBaseUnit,
			&pm.OpenConfirmedQuantityInBaseUnit,
			&pm.DeliveredQuantityInBaseUnit,
			&pm.UndeliveredQuantityInBaseUnit,
			&pm.StockIsFullyConfirmed,
			&pm.PlusMinusFlag,
			&pm.ItemScheduleLineDeliveryBlockStatus,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.IsReleased,
			&pm.IsLocked,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &itemComponentDeliveryScheduleLine, err
		}
		itemComponentDeliveryScheduleLine = append(itemComponentDeliveryScheduleLine, pm)
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &itemComponentDeliveryScheduleLine, nil
	}

	return &itemComponentDeliveryScheduleLine, nil
}

func ConvertToItemComponentCosting(rows *sql.Rows) (*[]ItemComponentCosting, error) {
	defer rows.Close()
	itemComponentCosting := make([]ItemComponentCosting, 0)
	i := 0
	for rows.Next() {
		i++
		pm := ItemComponentCosting{}
		err := rows.Scan(
			&pm.ProductionOrder,
			&pm.ProductionOrderItem,
			&pm.BillOfMaterial,
			&pm.BillOfMaterialItem,
			&pm.Currency,
			&pm.CostingAmount,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.IsReleased,
			&pm.IsLocked,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &itemComponentCosting, err
		}
		itemComponentCosting = append(itemComponentCosting, pm)
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &itemComponentCosting, nil
	}

	return &itemComponentCosting, nil
}

func ConvertToItemOperationCosting(rows *sql.Rows) (*[]ItemOperationCosting, error) {
	defer rows.Close()
	itemOperationCosting := make([]ItemOperationCosting, 0)
	i := 0
	for rows.Next() {
		i++
		pm := ItemOperationCosting{}
		err := rows.Scan(
			&pm.ProductionOrder,
			&pm.ProductionOrderItem,
			&pm.Operations,
			&pm.OperationsItem,
			&pm.OperationID,
			&pm.Currency,
			&pm.CostingAmount,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.IsReleased,
			&pm.IsLocked,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &itemOperationCosting, err
		}
		itemOperationCosting = append(itemOperationCosting, pm)
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &itemOperationCosting, nil
	}

	return &itemOperationCosting, nil
}

func ConvertToItemOperation(rows *sql.Rows) (*[]ItemOperation, error) {
	defer rows.Close()
	itemOperation := make([]ItemOperation, 0)
	i := 0
	for rows.Next() {
		i++
		pm := ItemOperation{}
		err := rows.Scan(
			&pm.ProductionOrder,
			&pm.ProductionOrderItem,
			&pm.Operations,
			&pm.OperationsItem,
			&pm.OperationID,
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipDeliveryID,
			&pm.SupplyChainRelationshipDeliveryPlantID,
			&pm.SupplyChainRelationshipProductionPlantID,
			&pm.Product,
			&pm.Buyer,
			&pm.Seller,
			&pm.DeliverToParty,
			&pm.DeliverToPlant,
			&pm.DeliverFromParty,
			&pm.DeliverFromPlant,
			&pm.ProductionPlantBusinessPartner,
			&pm.ProductionPlant,
			&pm.MRPArea,
			&pm.MRPController,
			&pm.ProductionVersion,
			&pm.ProductionVersionItem,
			&pm.Sequence,
			&pm.SequenceText,
			&pm.OperationText,
			&pm.ProductBaseUnit,
			&pm.ProductProductionUnit,
			&pm.ProductOperationUnit,
			&pm.ProductDeliveryUnit,
			&pm.StandardLotSizeQuantity,
			&pm.MinimumLotSizeQuantity,
			&pm.MaximumLotSizeQuantity,
			&pm.OperationPlannedQuantityInBaseUnit,
			&pm.OperationPlannedQuantityInProductionUnit,
			&pm.OperationPlannedQuantityInOperationUnit,
			&pm.OperationPlannedQuantityInDeliveryUnit,
			&pm.OperationPlannedScrapInPercent,
			&pm.ResponsiblePlannerGroup,
			&pm.PlainLongText,
			&pm.WorkCenter,
			&pm.CapacityCategory,
			&pm.OperationCostingRelevancyType,
			&pm.OperationSetupType,
			&pm.OperationSetupGroupCategory,
			&pm.OperationSetupGroup,
			&pm.MaximumWaitDuration,
			&pm.StandardWaitDuration,
			&pm.MinimumWaitDuration,
			&pm.WaitDurationUnit,
			&pm.MaximumQueDuration,
			&pm.StandardQueueDuration,
			&pm.MinimumQueueDuration,
			&pm.QueDurationUnit,
			&pm.MaximumMoveDuration,
			&pm.StandardMoveDuration,
			&pm.MinimumMoveDuration,
			&pm.MoveDurationUnit,
			&pm.StandardDeliveryDuration,
			&pm.StandardDeliveryDurationUnit,
			&pm.CostElement,
			&pm.OperationErlstSchedldExecStrtDte,
			&pm.OperationErlstSchedldExecStrtTme,
			&pm.OperationErlstSchedldExecEndDte,
			&pm.OperationErlstSchedldExecEndTme,
			&pm.OperationActualExecutionStartDate,
			&pm.OperationActualExecutionStartTime,
			&pm.OperationActualExecutionEndDate,
			&pm.OperationActualExecutionEndTime,
			&pm.OperationConfirmedYieldQuantityInBaseUnit,
			&pm.OperationConfirmedYieldQuantityInProductionUnit,
			&pm.OperationConfirmedYieldQuantityInOperationUnit,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.IsReleased,
			&pm.IsPartiallyConfirmed,
			&pm.IsConfirmed,
			&pm.IsLocked,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &itemOperation, err
		}
		itemOperation = append(itemOperation, pm)
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &itemOperation, nil
	}

	return &itemOperation, nil
}

func ConvertToItemOperationComponent(rows *sql.Rows) (*[]ItemOperationComponent, error) {
	defer rows.Close()
	itemOperationComponent := make([]ItemOperationComponent, 0)
	i := 0
	for rows.Next() {
		i++
		pm := ItemOperationComponent{}
		err := rows.Scan(
			&pm.ProductionOrder,
			&pm.ProductionOrderItem,
			&pm.Operations,
			&pm.OperationsItem,
			&pm.OperationID,
			&pm.BillOfMaterial,
			&pm.BillOfMaterialItem,
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipDeliveryID,
			&pm.SupplyChainRelationshipDeliveryPlantID,
			&pm.ProductionPlantBusinessPartner,
			&pm.ProductionPlant,
			&pm.ComponentProduct,
			&pm.ComponentProductBuyer,
			&pm.ComponentProductSeller,
			&pm.ComponentProductDeliverToParty,
			&pm.ComponentProductDeliverToPlant,
			&pm.ComponentProductDeliverFromParty,
			&pm.ComponentProductDeliverFromPlant,
			&pm.ComponentProductDeliverToPartyInOperation,
			&pm.ComponentProductDeliverToPlantInOperation,
			&pm.ComponentProductDeliverFromPartyInOperation,
			&pm.ComponentProductDeliverFromPlantInOperation,
			&pm.ComponentProductRequirementDateInOperation,
			&pm.ComponentProductRequirementTimeInOperation,
			&pm.ComponentProductPlannedQuantityInBaseUnitInOperation,
			&pm.ComponentProductPlannedQuantityInDeliveryUnitInOperation,
			&pm.ComponentProductPlannedScrapInPercentInOperation,
			&pm.ComponentProductBaseUnit,
			&pm.ComponentProductDeliveryUnit,
			&pm.ComponentProductIsMarkedForBackflush,
			&pm.MRPArea,
			&pm.MRPController,
			&pm.ProductionVersion,
			&pm.ProductionVersionItem,
			&pm.ComponentProductReservation,
			&pm.ComponentProductReservationItem,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.IsReleased,
			&pm.IsLocked,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &itemOperationComponent, err
		}
		itemOperationComponent = append(itemOperationComponent, pm)
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &itemOperationComponent, nil
	}

	return &itemOperationComponent, nil
}

func ConvertToHeadersByOwnerProductionPlantBusinessPartner(rows *sql.Rows) (*[]HeadersByOwnerProductionPlantBusinessPartner, error) {
	defer rows.Close()
	headersByOwnerProductionPlantBusinessPartner := make([]HeadersByOwnerProductionPlantBusinessPartner, 0)
	i := 0
	for rows.Next() {
		i++
		pm := HeadersByOwnerProductionPlantBusinessPartner{}
		err := rows.Scan(
			&pm.ProductionOrder,
			&pm.MRPArea,
			&pm.OwnerProductionPlantBusinessPartnerName,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &headersByOwnerProductionPlantBusinessPartner, err
		}
		headersByOwnerProductionPlantBusinessPartner = append(headersByOwnerProductionPlantBusinessPartner, pm)
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &headersByOwnerProductionPlantBusinessPartner, nil
	}

	return &headersByOwnerProductionPlantBusinessPartner, nil
}
