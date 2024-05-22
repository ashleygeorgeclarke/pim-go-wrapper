package models

type (
	Product struct {
		ID int `json:"id,omitempty"`
		ProductRequest
	}

	BulkUpdateProductRequestItem struct {
		ResourceID uint `json:"resourceId"`
		ProductRequest
	}

	ProductRequest struct {
		// Product type, possible types are 'PRODUCT', 'BUNDLE', 'MATRIX', 'ASSEMBLY'. By default 'PRODUCT'.
		Type string `json:"type,omitempty"`
		// ID of product group. To get the list of product groups, use getProductGroups.
		GroupID int `json:"group_id,omitempty"`
		// ID of product unit. To get the list of units, use getProductUnits.
		UnitID int `json:"unit_id,omitempty"`

		TranslatableNameJSON

		// //PlainTextDescription is translatable plain text description
		// PlainTextDescription map[string]string `json:"plain_text_description,omitempty"`
		// //HTMLDescription is translatable html description
		// HTMLDescription map[string]string `json:"html_description,omitempty"`
		// Description is translatable and of 2 types: plain text, HTML. Languages should be in ISO 639-1 Code
		TranslatableDescriptionJSON

		// Product's code. Must be UNIQUE, unless the account is configured otherwise.
		Code string `json:"code,omitempty"`
		// Product's second code (by convention, EAN barcode). Must be UNIQUE, unless the account is configured otherwise.
		Code2 string `json:"code2,omitempty"`
		// Third code of the item (note that this field may not be visible on product card by default).
		Code3             string `json:"code3,omitempty"`
		Code5             string `json:"code5,omitempty"`
		Code6             string `json:"code6,omitempty"`
		Code7             string `json:"code7,omitempty"`
		Code8             string `json:"code8,omitempty"`
		ExtraFieldOneID   int    `json:"extra_field1_id,omitempty"`
		ExtraFieldTwoID   int    `json:"extra_field2_id,omitempty"`
		ExtraFieldThreeID int    `json:"extra_field3_id,omitempty"`
		ExtraFieldFourID  int    `json:"extra_field4_id,omitempty"`
		// Supplier's product code
		SupplierCode string `json:"supplier_code,omitempty"`
		//TaxRateID is just the default tax rate of a product and the actual tax applied in a particular location depends on multiple rules: https://learn-api.erply.com/concepts/taxes.
		TaxRateID int `json:"tax_rate_id,omitempty"`
		//Price is just the default price of a product and the actual price applied in a particular location, to a particular customer, depends on price lists and promotions: https://learn-api.erply.com/concepts/pricing
		Price float64 `json:"price,omitempty"`

		PriceWithTax float64 `json:"price_with_tax,omitempty"`

		Physical
		//0 or 1
		IsGiftCard int `json:"is_gift_card,omitempty"`
		//0 or 1
		NonDiscountable int `json:"non_discountable,omitempty"`
		//0 or 1
		NonRefundable int `json:"non_refundable,omitempty"`

		// Volume is Item's fluid volume, eg. for beverages or perfumery. Unit depends on locale, check your Erply account (typically mL or fl oz).
		Volume     int `json:"volume,omitempty"`
		CategoryID int `json:"category_id,omitempty"`
		// ID of product brand. To get the list of brands, use getBrands.
		BrandID           int    `json:"brand_id,omitempty"`
		SupplierID        int    `json:"supplier_id,omitempty"`
		PriorityGroupID   int    `json:"priority_group_id,omitempty"`
		CountryOfOriginID int    `json:"country_of_origin_id,omitempty"`
		ManufacturerName  string `json:"manufacturer_name,omitempty"`
		// Cost is Product cost
		Cost float64 `json:"cost,omitempty"`
		Status
		//0 or 1
		DisplayedInWebshop *int `json:"displayed_in_webshop,omitempty"`
		// LocationInWarehouseID is ID of selected location in warehouse.
		LocationInWarehouseID int `json:"location_in_warehouse_id,omitempty"`
		// LocationInWarehouseText is Product's specific text added to location in warehouse.
		LocationInWarehouseText string `json:"location_in_warehouse_text,omitempty"`
		// Parent product ID. Only for matrix variations (specific colors/sizes of a matrix item). See guidelines below.
		ParentProductID int `json:"parent_product_id,omitempty"`
		// ContainerID is ID of another product, a beverage container that is always sold together with this item.
		DepositFeeID int `json:"deposit_fee_id,omitempty"`

		FamilyID int64 `json:"family_id,omitempty"`

		AgeRestriction   int `json:"age_restriction"`
		BackupID         int `json:"backup_id"`
		HasSerialNumbers int `json:"has_serial_numbers"`
		SoldInPackages   int `json:"sold_in_packages"`

		//These fields are not editable
		AddedByChangedBy

		*ProductAttributes
	}

	AutoCodes struct {
		NextCode  int `json:"nextCode"`
		NextCode2 int `json:"nextCode2"`
	}
)

type LinkedProductCreateRequest struct {
	ProductID       int    `json:"product_id"`
	LinkedProductID int    `json:"linked_product_id"`
	Type            string `json:"type"`
}
