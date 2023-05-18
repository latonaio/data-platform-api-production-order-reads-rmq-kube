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
			&pm.ProductionOrderType,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.HeaderIsReleased,
			&pm.HeaderIsPartiallyConfirmed,
			&pm.HeaderIsConfirmed,
			&pm.HeaderIsLocked,
			&pm.HeaderIsMarkedForDeletion,
			&pm.Product,
			&pm.OwnerProductionPlant,
			&pm.OwnerProductionPlantBusinessPartner,
			&pm.OwnerProductionPlantStorageLocation,
			&pm.MRPArea,
			&pm.MRPController,
			&pm.ProductionSupervisor,
			&pm.ProductionVersion,
			&pm.PlannedOrder,
			&pm.OrderID,
			&pm.OrderItem,
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
			&pm.ProductionUnit,
			&pm.TotalQuantity,
			&pm.PlannedScrapQuantity,
			&pm.ConfirmedYieldQuantity,
			&pm.ProductionOrderHeaderText,
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
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.ItemIsReleased,
			&pm.ItemIsPartiallyConfirmed,
			&pm.ItemIsConfirmed,
			&pm.ItemIsLocked,
			&pm.ItemIsMarkedForDeletion,
			&pm.ProductionOrderHasGeneratedOperations,
			&pm.ProductAvailabilityIsNotChecked,
			&pm.PrecedingItem,
			&pm.FollowingItem,
			&pm.Product,
			&pm.ProductionPlant,
			&pm.ProductionPlantBusinessPartner,
			&pm.ProductionPlantStorageLocation,
			&pm.MRPArea,
			&pm.MRPController,
			&pm.ProductionSupervisor,
			&pm.ProductionVersion,
			&pm.PlannedOrder,
			&pm.OrderID,
			&pm.OrderItem,
			&pm.MinimumLotSizeQuantity,
			&pm.StandardLotSizeQuantity,
			&pm.LotSizeRoundingQuantity,
			&pm.MaximumLotSizeQuantity,
			&pm.LotSizeIsFixed,
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
			&pm.ProductionUnit,
			&pm.TotalQuantity,
			&pm.PlannedScrapQuantity,
			&pm.ConfirmedYieldQuantity,
			&pm.ProductionOrderItemText,
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
			&pm.Operations,
			&pm.OperationsItem,
			&pm.BillOfMaterial,
			&pm.BillOfMaterialItem,
			&pm.Reservation,
			&pm.ReservationItem,
			&pm.ComponentProduct,
			&pm.ComponentProductRequirementDate,
			&pm.ComponentProductRequirementTime,
			&pm.ComponentProductIsMarkedForBackflush,
			&pm.ComponentProductBusinessPartner,
			&pm.StockConfirmationPlant,
			&pm.PlannedOrder,
			&pm.OrderID,
			&pm.OrderItem,
			&pm.SortField,
			&pm.BOMItemDescription,
			&pm.StorageLocation,
			&pm.Batch,
			&pm.GoodsRecipientName,
			&pm.UnloadingPointName,
			&pm.ProductCompIsAlternativeItem,
			&pm.CostingPolicy,
			&pm.PriceUnitQty,
			&pm.StandardPrice,
			&pm.MovingAveragePrice,
			&pm.ComponentScrapInPercent,
			&pm.OperationScrapInPercent,
			&pm.BaseUnit,
			&pm.RequiredQuantity,
			&pm.WithdrawnQuantity,
			&pm.ConfirmedAvailableQuantity,
			&pm.ProductCompOriginalQuantity,
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

func ConvertToItemComponentStockConfirmation(rows *sql.Rows) (*[]ItemComponentStockConfirmation, error) {
	defer rows.Close()
	itemComponentStockConfirmation := make([]ItemComponentStockConfirmation, 0)
	i := 0
	for rows.Next() {
		i++
		pm := ItemComponentStockConfirmation{}
		err := rows.Scan(
			&pm.ProductionOrder,
			&pm.ProductionOrderItem,
			&pm.Operations,
			&pm.OperationsItem,
			&pm.BillOfMaterial,
			&pm.BillOfMaterialItem,
			&pm.InventoryStockType,
			&pm.InventorySpecialStockType,
			&pm.AvailableProductStock,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &itemComponentStockConfirmation, err
		}
		itemComponentStockConfirmation = append(itemComponentStockConfirmation, pm)
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &itemComponentStockConfirmation, nil
	}

	return &itemComponentStockConfirmation, nil
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
			&pm.Operations,
			&pm.OperationsItem,
			&pm.BillOfMaterial,
			&pm.BillOfMaterialItem,
			&pm.ComponentProduct,
			&pm.Currency,
			&pm.CostingAmount,
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
			&pm.Sequence,
			&pm.OperationsText,
			&pm.SequenceText,
			&pm.OperationIsReleased,
			&pm.OperationIsPartiallyConfirmed,
			&pm.OperationIsConfirmed,
			&pm.OperationIsClosed,
			&pm.OperationIsMarkedForDeletion,
			&pm.ProductionPlant,
			&pm.WorkCenter,
			&pm.OperationErlstSchedldExecStrtDte,
			&pm.OperationErlstSchedldExecStrtTme,
			&pm.OperationErlstSchedldExecEndDate,
			&pm.OperationErlstSchedldExecEndTme,
			&pm.OperationActualExecutionStartDate,
			&pm.OperationActualExecutionStartTime,
			&pm.OperationActualExecutionEndDate,
			&pm.OperationActualExecutionEndTime,
			&pm.ErlstSchedldExecDurnInWorkdays,
			&pm.OperationActualExecutionDays,
			&pm.OperationUnit,
			&pm.OperationPlannedTotalQuantity,
			&pm.OperationTotalConfirmedYieldQuantity,
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
