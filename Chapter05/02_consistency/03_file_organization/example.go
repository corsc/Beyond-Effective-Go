package _3_file_organization

type OrderManager struct{}

func (o *OrderManager) Process(order Order) error {
	err := o.validateOrder(order)
	if err != nil {
		return err
	}

	receiptNo := o.chargeCustomer(order)

	o.sendReceipt(order, receiptNo)

	o.informWarehouse(order)

	return nil
}

func (o *OrderManager) validateOrder(order Order) error {
	// implementation removed
	return nil
}

func (o *OrderManager) chargeCustomer(order Order) string {
	// implementation removed
	return ""
}

func (o *OrderManager) sendReceipt(order Order, receiptNo string) {
	// implementation removed
}

func (o *OrderManager) informWarehouse(order Order) {
	// implementation removed
}

type Order struct {
	// implementation removed
}
