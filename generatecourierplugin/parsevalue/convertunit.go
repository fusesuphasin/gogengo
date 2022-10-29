package parsevalue

type ConvertUnit struct{}

func (ct *ConvertUnit) ConvertUnitCreateLabelUnit(label *domain.Label) {
	label.ShippingService.Package.Weight.Unit, label.ShippingService.Package.Weight.Value = ct.ConvertWeight(label.ShippingService.Package.Weight.Unit, label.ShippingService.Package.Weight.Value)
}

func (ct *ConvertUnit) ConvertUnitGetRateUnit(rate *domain.CheckRate) {
	rate.Order.Weight.Unit, rate.Order.Weight.Value = ct.ConvertWeight(rate.Order.Weight.Unit, rate.Order.Weight.Value)
	rate.Order.Dimension.Unit, rate.Order.Dimension.Sum = ct.ConvertDimension(rate.Order.Dimension.Unit, rate.Order.Dimension.Sum)
}

func (ct *ConvertUnit) ConvertWeight(unit string, value float64) (string, float64) {
	defualtUnit := "kg"

	switch unit {
	case defualtUnit:
		return unit, value
	case "g":
		return defualtUnit, value * 0.001
	case "lb":
		return defualtUnit, value * 0.45359237
	case "oz":
		return defualtUnit, value * 0.0283495231
	}

	return unit, value
}

func (ct *ConvertUnit) ConvertDimension(unit string, value float64) (string, float64) {
	defualtUnit := "cm"

	switch unit {
	case defualtUnit:
		return unit, value
	case "in":
		return defualtUnit, value * 2.54
	case "m":
		return defualtUnit, value * 100
	}

	return unit, value
}

func (ct *ConvertUnit) ConvertCurrency(unit string, value float64) (string, float64) {
	//defualtUnit := "THB"

	return unit, value
}