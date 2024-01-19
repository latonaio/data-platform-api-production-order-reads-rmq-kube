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
			&pm.ProductionOrderDate,
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
			&pm.CertificateAuthorityChain,
			&pm.UsageControlChain,
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

		data := pm
		header = append(header, Header{
			ProductionOrder:                                  data.ProductionOrder,
			ProductionOrderDate:                              data.ProductionOrderDate,
			SupplyChainRelationshipID:                        data.SupplyChainRelationshipID,
			SupplyChainRelationshipProductionPlantID:         data.SupplyChainRelationshipProductionPlantID,
			SupplyChainRelationshipDeliveryID:                data.SupplyChainRelationshipDeliveryID,
			SupplyChainRelationshipDeliveryPlantID:           data.SupplyChainRelationshipDeliveryPlantID,
			ProductionOrderType:                              data.ProductionOrderType,
			Product:                                          data.Product,
			Buyer:                                            data.Buyer,
			Seller:                                           data.Seller,
			OwnerProductionPlantBusinessPartner:              data.OwnerProductionPlantBusinessPartner,
			OwnerProductionPlant:                             data.OwnerProductionPlant,
			OwnerProductionPlantStorageLocation:              data.OwnerProductionPlantStorageLocation,
			DepartureDeliverFromParty:                        data.DepartureDeliverFromParty,
			DepartureDeliverFromPlant:                        data.DepartureDeliverFromPlant,
			DepartureDeliverFromPlantStorageLocation:         data.DepartureDeliverFromPlantStorageLocation,
			DestinationDeliverToParty:                        data.DestinationDeliverToParty,
			DestinationDeliverToPlant:                        data.DestinationDeliverToPlant,
			DestinationDeliverToPlantStorageLocation:         data.DestinationDeliverToPlantStorageLocation,
			ProductBaseUnit:                                  data.ProductBaseUnit,
			MRPArea:                                          data.MRPArea,
			MRPController:                                    data.MRPController,
			ProductionVersion:                                data.ProductionVersion,
			BillOfMaterial:                                   data.BillOfMaterial,
			Operations:                                       data.Operations,
			ProductionOrderQuantityInBaseUnit:                data.ProductionOrderQuantityInBaseUnit,
			ProductionOrderQuantityInDepartureProductionUnit: data.ProductionOrderQuantityInDepartureProductionUnit,
			ProductionOrderQuantityInDestinationProductionUnit: data.ProductionOrderQuantityInDestinationProductionUnit,
			ProductionOrderQuantityInDepartureDeliveryUnit:     data.ProductionOrderQuantityInDepartureDeliveryUnit,
			ProductionOrderQuantityInDestinationDeliveryUnit:   data.ProductionOrderQuantityInDestinationDeliveryUnit,
			ProductionOrderDepartureProductionUnit:             data.ProductionOrderDepartureProductionUnit,
			ProductionOrderDestinationProductionUnit:           data.ProductionOrderDestinationProductionUnit,
			ProductionOrderDepartureDeliveryUnit:               data.ProductionOrderDepartureDeliveryUnit,
			ProductionOrderDestinationDeliveryUnit:             data.ProductionOrderDestinationDeliveryUnit,
			ProductionOrderPlannedScrapQtyInBaseUnit:           data.ProductionOrderPlannedScrapQtyInBaseUnit,
			ProductionOrderPlannedStartDate:                    data.ProductionOrderPlannedStartDate,
			ProductionOrderPlannedStartTime:                    data.ProductionOrderPlannedStartTime,
			ProductionOrderPlannedEndDate:                      data.ProductionOrderPlannedEndDate,
			ProductionOrderPlannedEndTime:                      data.ProductionOrderPlannedEndTime,
			ProductionOrderActualReleaseDate:                   data.ProductionOrderActualReleaseDate,
			ProductionOrderActualReleaseTime:                   data.ProductionOrderActualReleaseTime,
			ProductionOrderActualStartDate:                     data.ProductionOrderActualStartDate,
			ProductionOrderActualStartTime:                     data.ProductionOrderActualStartTime,
			ProductionOrderActualEndDate:                       data.ProductionOrderActualEndDate,
			ProductionOrderActualEndTime:                       data.ProductionOrderActualEndTime,
			PlannedOrder:                                       data.PlannedOrder,
			OrderID:                                            data.OrderID,
			OrderItem:                                          data.OrderItem,
			ProductionOrderHeaderText:                          data.ProductionOrderHeaderText,
			CertificateAuthorityChain:		  					data.CertificateAuthorityChain,
			UsageControlChain:		  		  					data.UsageControlChain,
			CreationDate:                                       data.CreationDate,
			CreationTime:                                       data.CreationTime,
			LastChangeDate:                                     data.LastChangeDate,
			LastChangeTime:                                     data.LastChangeTime,
			IsReleased:                                         data.IsReleased,
			IsPartiallyConfirmed:                               data.IsPartiallyConfirmed,
			IsConfirmed:                                        data.IsConfirmed,
			IsLocked:                                           data.IsLocked,
			IsCancelled:                                        data.IsCancelled,
			IsMarkedForDeletion:                                data.IsMarkedForDeletion,
		})
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
			&pm.ProductionOrderItemDate,
			&pm.PrecedingProductionOrderItem,
			&pm.FollowingProductionOrderItem,
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipProductionPlantID,
			&pm.SupplyChainRelationshipDeliveryID,
			&pm.SupplyChainRelationshipDeliveryPlantID,
			&pm.ProductionOrderType,
			&pm.Product,
			&pm.Buyer,
			&pm.Seller,
			&pm.ProductionPlantBusinessPartner,
			&pm.ProductionPlant,
			&pm.ProductionPlantStorageLocation,
			&pm.DeliverFromParty,
			&pm.DeliverFromPlant,
			&pm.DeliverFromPlantStorageLocation,
			&pm.DeliverToParty,
			&pm.DeliverToPlant,
			&pm.DeliverToPlantStorageLocation,
			&pm.ProductBaseUnit,
			&pm.ProductProductionUnit,
			&pm.ProductDeliveryUnit,
			&pm.MRPArea,
			&pm.MRPController,
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

		data := pm
		item = append(item, Item{
			ProductionOrder:                               data.ProductionOrder,
			ProductionOrderItem:                           data.ProductionOrderItem,
			ProductionOrderItemDate:                       data.ProductionOrderItemDate,
			PrecedingProductionOrderItem:                  data.PrecedingProductionOrderItem,
			FollowingProductionOrderItem:                  data.FollowingProductionOrderItem,
			SupplyChainRelationshipID:                     data.SupplyChainRelationshipID,
			SupplyChainRelationshipProductionPlantID:      data.SupplyChainRelationshipProductionPlantID,
			SupplyChainRelationshipDeliveryID:             data.SupplyChainRelationshipDeliveryID,
			SupplyChainRelationshipDeliveryPlantID:        data.SupplyChainRelationshipDeliveryPlantID,
			ProductionOrderType:                           data.ProductionOrderType,
			Product:                                       data.Product,
			Buyer:                                         data.Buyer,
			Seller:                                        data.Seller,
			ProductionPlantBusinessPartner:                data.ProductionPlantBusinessPartner,
			ProductionPlant:                               data.ProductionPlant,
			ProductionPlantStorageLocation:                data.ProductionPlantStorageLocation,
			DeliverFromParty:                              data.DeliverFromParty,
			DeliverFromPlant:                              data.DeliverFromPlant,
			DeliverFromPlantStorageLocation:               data.DeliverFromPlantStorageLocation,
			DeliverToParty:                                data.DeliverToParty,
			DeliverToPlant:                                data.DeliverToPlant,
			DeliverToPlantStorageLocation:                 data.DeliverToPlantStorageLocation,
			ProductBaseUnit:                               data.ProductBaseUnit,
			ProductProductionUnit:                         data.ProductProductionUnit,
			ProductDeliveryUnit:                           data.ProductDeliveryUnit,
			MRPArea:                                       data.MRPArea,
			MRPController:                                 data.MRPController,
			ProductionVersion:                             data.ProductionVersion,
			ProductionVersionItem:                         data.ProductionVersionItem,
			BillOfMaterial:                                data.BillOfMaterial,
			Operations:                                    data.Operations,
			ProductionOrderQuantityInBaseUnit:             data.ProductionOrderQuantityInBaseUnit,
			ProductionOrderQuantityInProductionUnit:       data.ProductionOrderQuantityInProductionUnit,
			ProductionOrderQuantityInDeliveryUnit:         data.ProductionOrderQuantityInDeliveryUnit,
			ProductionOrderPlannedScrapQtyInBaseUnit:      data.ProductionOrderPlannedScrapQtyInBaseUnit,
			ProductionOrderMinimumLotSizeQuantity:         data.ProductionOrderMinimumLotSizeQuantity,
			ProductionOrderStandardLotSizeQuantity:        data.ProductionOrderStandardLotSizeQuantity,
			ProductionOrderMaximumLotSizeQuantity:         data.ProductionOrderMaximumLotSizeQuantity,
			ProductionOrderLotSizeRoundingQuantity:        data.ProductionOrderLotSizeRoundingQuantity,
			ProductionOrderLotSizeIsFixed:                 data.ProductionOrderLotSizeIsFixed,
			ProductionOrderPlannedStartDate:               data.ProductionOrderPlannedStartDate,
			ProductionOrderPlannedStartTime:               data.ProductionOrderPlannedStartTime,
			ProductionOrderPlannedEndDate:                 data.ProductionOrderPlannedEndDate,
			ProductionOrderPlannedEndTime:                 data.ProductionOrderPlannedEndTime,
			ProductionOrderActualReleaseDate:              data.ProductionOrderActualReleaseDate,
			ProductionOrderActualReleaseTime:              data.ProductionOrderActualReleaseTime,
			ProductionOrderActualStartDate:                data.ProductionOrderActualStartDate,
			ProductionOrderActualStartTime:                data.ProductionOrderActualStartTime,
			ProductionOrderActualEndDate:                  data.ProductionOrderActualEndDate,
			ProductionOrderActualEndTime:                  data.ProductionOrderActualEndTime,
			ConfirmedYieldQuantityInBaseUnit:              data.ConfirmedYieldQuantityInBaseUnit,
			ConfirmedYieldQuantityInProductionUnit:        data.ConfirmedYieldQuantityInProductionUnit,
			ScrappedQuantityInBaseUnit:                    data.ScrappedQuantityInBaseUnit,
			PlannedOrder:                                  data.PlannedOrder,
			PlannedOrderItem:                              data.PlannedOrderItem,
			OrderID:                                       data.OrderID,
			OrderItem:                                     data.OrderItem,
			ProductIsBatchManagedInProductionPlant:        data.ProductIsBatchManagedInProductionPlant,
			BatchMgmtPolicyInProductionOrder:              data.BatchMgmtPolicyInProductionOrder,
			ProductionOrderTargetedBatch:                  data.ProductionOrderTargetedBatch,
			ProductionOrderTargetedBatchValidityStartDate: data.ProductionOrderTargetedBatchValidityStartDate,
			ProductionOrderTargetedBatchValidityStartTime: data.ProductionOrderTargetedBatchValidityStartTime,
			ProductionOrderTargetedBatchValidityEndDate:   data.ProductionOrderTargetedBatchValidityEndDate,
			ProductionOrderTargetedBatchValidityEndTime:   data.ProductionOrderTargetedBatchValidityEndTime,
			ProductionOrderItemText:                       data.ProductionOrderItemText,
			CreationDate:                                  data.CreationDate,
			CreationTime:                                  data.CreationTime,
			LastChangeDate:                                data.LastChangeDate,
			LastChangeTime:                                data.LastChangeTime,
			IsReleased:                                    data.IsReleased,
			IsPartiallyConfirmed:                          data.IsPartiallyConfirmed,
			IsConfirmed:                                   data.IsConfirmed,
			IsLocked:                                      data.IsLocked,
			IsCancelled:                                   data.IsCancelled,
			IsMarkedForDeletion:                           data.IsMarkedForDeletion,
		})
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

func ConvertToItemComponentPricingElement(rows *sql.Rows) (*[]ItemComponentPricingElement, error) {
	defer rows.Close()
	itemComponentPricingElement := make([]ItemComponentPricingElement, 0)
	i := 0
	for rows.Next() {
		i++
		pm := ItemComponentPricingElement{}
		err := rows.Scan(
			&pm.ProductionOrder,
			&pm.ProductionOrderItem,
			&pm.BillOfMaterial,
			&pm.BillOfMaterialItem,
			&pm.PricingProcedureCounter,
			&pm.SupplyChainRelationshipID,
			&pm.ComponentProductBuyer,
			&pm.ComponentProductSeller,
			&pm.ConditionRecord,
			&pm.ConditionSequentialNumber,
			&pm.ConditionType,
			&pm.PricingDate,
			&pm.ConditionRateValue,
			&pm.ConditionRateValueUnit,
			&pm.ConditionScaleQuantity,
			&pm.ConditionCurrency,
			&pm.ConditionQuantity,
			&pm.TaxCode,
			&pm.ConditionAmount,
			&pm.TransactionCurrency,
			&pm.ConditionIsManuallyChanged,
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
			return &itemComponentPricingElement, err
		}
		itemComponentPricingElement = append(itemComponentPricingElement, pm)
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &itemComponentPricingElement, nil
	}

	return &itemComponentPricingElement, nil
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
			&pm.OperationType,
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
			&pm.MaximumQueueDuration,
			&pm.StandardQueueDuration,
			&pm.MinimumQueueDuration,
			&pm.QueueDurationUnit,
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
