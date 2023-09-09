package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-production-order-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-production-order-reads-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) readSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var header *[]dpfm_api_output_formatter.Header
	var item *[]dpfm_api_output_formatter.Item
	var itemComponent *[]dpfm_api_output_formatter.ItemComponent
	var itemComponentDeliveryScheduleLine *[]dpfm_api_output_formatter.ItemComponentDeliveryScheduleLine
	var itemComponentCosting *[]dpfm_api_output_formatter.ItemComponentCosting
	var itemOperation *[]dpfm_api_output_formatter.ItemOperation
	var itemOperationComponent *[]dpfm_api_output_formatter.ItemOperationComponent
	var itemOperationCosting *[]dpfm_api_output_formatter.ItemOperationCosting

	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				header = c.Header(mtx, input, output, errs, log)
			}()
		case "HeadersByOwnerProductionPlantBP":
			func() {
				header = c.HeadersByOwnerProductionPlantBP(mtx, input, output, errs, log)
			}()
		case "Item":
			func() {
				item = c.Item(mtx, input, output, errs, log)
			}()
		case "Items":
			func() {
				item = c.Items(mtx, input, output, errs, log)
			}()
		case "ItemComponent":
			func() {
				itemComponent = c.ItemComponent(mtx, input, output, errs, log)
			}()
		case "ItemComponents":
			func() {
				itemComponent = c.ItemComponents(mtx, input, output, errs, log)
			}()
		case "ItemComponentDeliveryScheduleLine":
			func() {
				itemComponentDeliveryScheduleLine = c.ItemComponentDeliveryScheduleLine(mtx, input, output, errs, log)
			}()
		case "ItemComponentCosting":
			func() {
				itemComponentCosting = c.ItemComponentCosting(mtx, input, output, errs, log)
			}()
		case "ItemOperation":
			func() {
				itemOperation = c.ItemOperation(mtx, input, output, errs, log)
			}()
		case "ItemOperations":
			func() {
				itemOperation = c.ItemOperations(mtx, input, output, errs, log)
			}()
		case "ItemOperationComponent":
			func() {
				itemOperationComponent = c.ItemOperationComponent(mtx, input, output, errs, log)
			}()
		case "ItemOperationComponents":
			func() {
				itemOperationComponent = c.ItemOperationComponents(mtx, input, output, errs, log)
			}()
		case "ItemOperationCosting":
			func() {
				itemOperationCosting = c.ItemOperationCosting(mtx, input, output, errs, log)
			}()
		case "HeadersByOwnerProductionPlantBusinessPartner":
			func() {
				header = c.HeadersByOwnerProductionPlantBusinessPartner(mtx, input, output, errs, log)
			}()
		default:
		}
		if len(*errs) != 0 {
			break
		}
	}

	data := &dpfm_api_output_formatter.Message{
		Header:                            header,
		Item:                              item,
		ItemComponent:                     itemComponent,
		ItemComponentDeliveryScheduleLine: itemComponentDeliveryScheduleLine,
		ItemComponentCosting:              itemComponentCosting,
		ItemOperation:                     itemOperation,
		ItemOperationComponent:            itemOperationComponent,
		ItemOperationCosting:              itemOperationCosting,
	}
	return data
}

func (c *DPFMAPICaller) Header(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {
	where := "WHERE 1 = 1"

	where = fmt.Sprintf("%s\nAND ProductionOrder = %d ", where, input.Header.ProductionOrder)

	if input.Header.IsReleased != nil {
		where = fmt.Sprintf("%s\nAND IsReleased = %v", where, *input.Header.IsReleased)
	}
	if input.Header.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %v", where, *input.Header.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_production_order_header_data
		` + where + ` ORDER BY IsMarkedForDeletion DESC, IsLocked DESC, ProductionVersion DESC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) HeadersByOwnerProductionPlantBP(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {
	where := "WHERE 1 = 1"
	if input.Header.IsReleased != nil {
		where = fmt.Sprintf("%s\nAND IsReleased = %v", where, *input.Header.IsReleased)
	}
	if input.Header.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %v", where, *input.Header.IsMarkedForDeletion)
	}
	if input.Header.OwnerProductionPlantBusinessPartner != nil {
		where = fmt.Sprintf("%s\nAND OwnerProductionPlantBusinessPartner = %d", where, *input.Header.OwnerProductionPlantBusinessPartner)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_production_order_header_data
		` + where + ` ORDER BY IsMarkedForDeletion DESC, IsLocked DESC, ProductionVersion DESC;`)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) HeadersByOwnerProductionPlantBusinessPartner(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {
	ownerProductionPlantBusinessPartner := input.Header.OwnerProductionPlantBusinessPartner

	rows, err := c.db.Query(
		`SELECT
		ProductionOrderHeader.ProductionOrder, 
		ProductionOrderHeader.MRPArea, 
		BusinessPartnerGeneralOwnerProductionPlant.BusinessPartnerName as OwnerProductionPlantBusinessPartnerName
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_production_order_header_data as ProductionOrderHeader
		LEFT JOIN DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_general_data as BusinessPartnerGeneralOwnerProductionPlant
		ON ProductionOrderHeader.OwnerProductionPlantBusinessPartner = BusinessPartnerGeneralOwnerProductionPlant.BusinessPartner
		WHERE (ProductionOrderHeader.OwnerProductionPlantBusinessPartner) = (?) 
		ORDER BY ProductionOrderHeader.IsMarkedForDeletion DESC, ProductionOrderHeader.IsLocked DESC, ProductionOrderHeader.ProductionVersion DESC;`, ownerProductionPlantBusinessPartner,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Item(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Item {
	var args []interface{}
	where := fmt.Sprintf("WHERE ProductionOrder = %d ", input.Header.ProductionOrder)

	itemIDs := ""
	for _, v := range input.Header.Item {
		itemIDs = fmt.Sprintf("%s, %d", itemIDs, v.ProductionOrderItem)
	}

	if len(itemIDs) != 0 {
		where = fmt.Sprintf("%s\nAND ProductionOrderItem IN ( %s ) ", where, itemIDs[1:])
	}
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_production_order_item_data
		`+where+` ORDER BY ProductionOrder DESC;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItem(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Items(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Item {
	where := fmt.Sprintf("WHERE ProductionOrder = %d ", input.Header.ProductionOrder)
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_production_order_item_data
		` + where + ` ORDER BY ProductionOrder DESC;`)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItem(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ItemComponent(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemComponent {
	var args []interface{}
	productionOrder := input.Header.ProductionOrder
	item := input.Header.Item

	cnt := 0
	for _, v := range item {
		itemOperation := v.ItemOperation
		for _, w := range itemOperation {
			itemOperationComponent := w.ItemOperationComponent
			for _, x := range itemOperationComponent {
				args = append(args, productionOrder, v.ProductionOrderItem, x.Operations, x.OperationsItem, x.BillOfMaterial, x.BillOfMaterialItem)
			}
		}
		cnt++
	}
	repeat := strings.Repeat("(?,?,?,?,?,?),", cnt-1) + "(?,?,?,?,?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_production_order_item_component_data
		WHERE (ProductionOrder, ProductionOrderItem, Operations, OperationsItem, BillOfMaterial, BillOfMaterialItem) IN ( `+repeat+` ) 
		ORDER BY IsMarkedForDeletion ASC, ProductionOrder DESC, ProductionOrderItem ASC;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItemComponent(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ItemComponents(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemComponent {
	item := &dpfm_api_input_reader.Item{}
	itemComponent := &dpfm_api_input_reader.ItemComponent{}
	if len(input.Header.Item) > 0 {
		item = &input.Header.Item[0]
	}

	if len(item.ItemComponent) > 0 {
		itemComponent = &item.ItemComponent[0]
	}

	where := fmt.Sprintf("WHERE ProductionOrder = %d", input.Header.ProductionOrder)
	if item != nil {
		if &item.ProductionOrderItem != nil {
			where = fmt.Sprintf("%s\nAND ProductionOrderItem = %d", where, item.ProductionOrderItem)
		}
	}
	if itemComponent != nil {
		if itemComponent.IsMarkedForDeletion != nil {
			where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %v", where, *itemComponent.IsMarkedForDeletion)
		}
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_production_order_item_component_data
		` + where + ` ORDER BY IsMarkedForDeletion ASC, ProductionOrder DESC, ProductionOrderItem ASC;`)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItemComponent(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ItemOperations(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemOperation {
	item := &dpfm_api_input_reader.Item{}
	itemOperation := &dpfm_api_input_reader.ItemOperation{}
	if len(input.Header.Item) > 0 {
		item = &input.Header.Item[0]
	}

	if len(item.ItemOperation) > 0 {
		itemOperation = &item.ItemOperation[0]
	}

	where := fmt.Sprintf("WHERE ProductionOrder = %d", input.Header.ProductionOrder)

	//if item != nil {
	//	if &item.ProductionOrderItem != nil {
	//		where = fmt.Sprintf("%s\nAND ProductionOrderItem = %d", where, item.ProductionOrderItem)
	//	}
	//}

	if itemOperation != nil {
		if itemOperation.IsMarkedForDeletion != nil {
			where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %v", where, *itemOperation.IsMarkedForDeletion)
		}
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_production_order_item_operation_data
		` + where + ` ORDER BY ProductionOrder DESC, OperationsItem ASC;`)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItemOperation(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ItemComponentDeliveryScheduleLine(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemComponentDeliveryScheduleLine {
	var args []interface{}
	productionOrder := input.Header.ProductionOrder
	item := input.Header.Item

	cnt := 0
	for _, v := range item {
		itemOperation := v.ItemOperation
		for _, w := range itemOperation {
			itemOperationComponent := w.ItemOperationComponent
			for _, x := range itemOperationComponent {
				args = append(args, productionOrder, v.ProductionOrderItem, x.Operations, x.OperationsItem, x.BillOfMaterial, x.BillOfMaterialItem)
			}
		}
		cnt++
	}
	repeat := strings.Repeat("(?,?,?,?,?,?),", cnt-1) + "(?,?,?,?,?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_production_order_item_component_stock_conf_data
		WHERE (ProductionOrder, ProductionOrderItem, Operations, OperationsItem, BillOfMaterial, BillOfMaterialItem) IN ( `+repeat+` ) 
		ORDER BY IsMarkedForDeletion ASC, ProductionOrder DESC, ProductionOrderItem ASC;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItemComponentDeliveryScheduleLine(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ItemComponentCosting(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemComponentCosting {
	var args []interface{}
	productionOrder := input.Header.ProductionOrder
	item := input.Header.Item

	cnt := 0
	for _, v := range item {
		itemOperation := v.ItemOperation
		for _, w := range itemOperation {
			itemOperationComponent := w.ItemOperationComponent
			for _, x := range itemOperationComponent {
				args = append(args, productionOrder, v.ProductionOrderItem, x.Operations, x.OperationsItem)
			}
		}
		cnt++
	}

	repeat := strings.Repeat("(?,?,?,?),", cnt-1) + "(?,?,?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_production_order_item_component_costing_data
		WHERE (ProductionOrder, ProductionOrderItem, Operations, OperationsItem) IN ( `+repeat+` ) 
		ORDER BY IsMarkedForDeletion ASC, ProductionOrder DESC, ProductionOrderItem ASC;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItemComponentCosting(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ItemOperation(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemOperation {
	var args []interface{}
	productionOrder := input.Header.ProductionOrder
	item := input.Header.Item

	cnt := 0
	for _, v := range item {
		itemOperation := v.ItemOperation
		for _, w := range itemOperation {
			args = append(args, productionOrder, v.ProductionOrderItem, w.Operations, w.OperationsItem)
		}
		cnt++
	}
	repeat := strings.Repeat("(?,?,?,?),", cnt-1) + "(?,?,?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_production_order_item_operation_data
		WHERE (ProductionOrder, ProductionOrderItem, Operations, OperationsItem) IN ( `+repeat+` ) 
		ORDER BY ProductionOrder DESC, OperationsItem ASC;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItemOperation(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ItemOperationComponent(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemOperationComponent {
	var args []interface{}
	productionOrder := input.Header.ProductionOrder
	item := input.Header.Item

	cnt := 0
	for _, v := range item {
		itemOperation := v.ItemOperation
		for _, w := range itemOperation {
			itemOperationComponent := w.ItemOperationComponent
			for _, x := range itemOperationComponent {
				args = append(args, productionOrder, v.ProductionOrderItem, x.Operations, x.OperationsItem)
			}
		}
		cnt++
	}

	repeat := strings.Repeat("(?,?,?,?),", cnt-1) + "(?,?,?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_production_order_item_operation_component_data
		WHERE (ProductionOrder, ProductionOrderItem, Operations, OperationsItem) IN ( `+repeat+` ) 
		ORDER BY ProductionOrder DESC, OperationsItem ASC;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItemOperationComponent(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	return data
}

func (c *DPFMAPICaller) ItemOperationComponents(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemOperationComponent {
	item := &dpfm_api_input_reader.Item{}
	itemOperation := &dpfm_api_input_reader.ItemOperation{}
	if len(input.Header.Item) > 0 {
		item = &input.Header.Item[0]
	}

	where := fmt.Sprintf("WHERE ProductionOrder = %d", input.Header.ProductionOrder)

	if item != nil {
		if &item.ProductionOrderItem != nil {
			where = fmt.Sprintf("%s\nAND ProductionOrderItem = %d", where, item.ProductionOrderItem)
		}
	}

	if itemOperation != nil {
		if itemOperation.IsMarkedForDeletion != nil {
			where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %v", where, *itemOperation.IsMarkedForDeletion)
		}
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_production_order_item_operation_component_data
		` + where + ` ORDER BY ProductionOrder DESC;`)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItemOperationComponent(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ItemOperationCosting(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemOperationCosting {
	var args []interface{}
	productionOrder := input.Header.ProductionOrder
	item := input.Header.Item

	cnt := 0
	for _, v := range item {
		itemOperation := v.ItemOperation
		for _, w := range itemOperation {
			itemOperationCosting := w.ItemOperationCosting
			for _, x := range itemOperationCosting {
				args = append(args, productionOrder, v.ProductionOrderItem, x.Operations, x.OperationsItem)
			}
		}
		cnt++
	}

	repeat := strings.Repeat("(?,?,?,?),", cnt-1) + "(?,?,?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_production_order_item_operation_costing_data
		WHERE (ProductionOrder, ProductionOrderItem, Operations, OperationsItem) IN ( `+repeat+` ) 
		ORDER BY ProductionOrder DESC, OperationsItem ASC;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItemOperationCosting(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	return data
}
