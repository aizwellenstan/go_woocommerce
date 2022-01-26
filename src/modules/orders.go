package modules

type Orders []struct {
	ID               int    `json:"id"`
	ParentID         int    `json:"parent_id"`
	Status           string `json:"status"`
	Currency         string `json:"currency"`
	Version          string `json:"version"`
	PricesIncludeTax bool   `json:"prices_include_tax"`
	DateCreated      string `json:"date_created"`
	DateModified     string `json:"date_modified"`
	DiscountTotal    string `json:"discount_total"`
	DiscountTax      string `json:"discount_tax"`
	ShippingTotal    string `json:"shipping_total"`
	ShippingTax      string `json:"shipping_tax"`
	CartTax          string `json:"cart_tax"`
	Total            string `json:"total"`
	TotalTax         string `json:"total_tax"`
	CustomerID       int    `json:"customer_id"`
	OrderKey         string `json:"order_key"`
	Billing          struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Company   string `json:"company"`
		Address1  string `json:"address_1"`
		Address2  string `json:"address_2"`
		City      string `json:"city"`
		State     string `json:"state"`
		Postcode  string `json:"postcode"`
		Country   string `json:"country"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
	} `json:"billing"`
	Shipping struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Company   string `json:"company"`
		Address1  string `json:"address_1"`
		Address2  string `json:"address_2"`
		City      string `json:"city"`
		State     string `json:"state"`
		Postcode  string `json:"postcode"`
		Country   string `json:"country"`
		Phone     string `json:"phone"`
	} `json:"shipping"`
	PaymentMethod      string      `json:"payment_method"`
	PaymentMethodTitle string      `json:"payment_method_title"`
	TransactionID      string      `json:"transaction_id"`
	CustomerIPAddress  string      `json:"customer_ip_address"`
	CustomerUserAgent  string      `json:"customer_user_agent"`
	CreatedVia         string      `json:"created_via"`
	CustomerNote       string      `json:"customer_note"`
	DateCompleted      interface{} `json:"date_completed"`
	DatePaid           interface{} `json:"date_paid"`
	CartHash           string      `json:"cart_hash"`
	Number             string      `json:"number"`
	MetaData           []struct {
		ID    int    `json:"id"`
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"meta_data"`
	LineItems []struct {
		ID          int           `json:"id"`
		Name        string        `json:"name"`
		ProductID   int           `json:"product_id"`
		VariationID int           `json:"variation_id"`
		Quantity    int           `json:"quantity"`
		TaxClass    string        `json:"tax_class"`
		Subtotal    string        `json:"subtotal"`
		SubtotalTax string        `json:"subtotal_tax"`
		Total       string        `json:"total"`
		TotalTax    string        `json:"total_tax"`
		Taxes       []interface{} `json:"taxes"`
		MetaData    []interface{} `json:"meta_data"`
		Sku         string        `json:"sku"`
		Price       int           `json:"price"`
		ParentName  interface{}   `json:"parent_name"`
	} `json:"line_items"`
	TaxLines      []interface{} `json:"tax_lines"`
	ShippingLines []interface{} `json:"shipping_lines"`
	FeeLines      []interface{} `json:"fee_lines"`
	CouponLines   []struct {
		ID          int    `json:"id"`
		Code        string `json:"code"`
		Discount    string `json:"discount"`
		DiscountTax string `json:"discount_tax"`
		MetaData    []struct {
			ID    int    `json:"id"`
			Key   string `json:"key"`
			Value struct {
				ID          int    `json:"id"`
				Code        string `json:"code"`
				Amount      string `json:"amount"`
				DateCreated struct {
					Date         string `json:"date"`
					TimezoneType int    `json:"timezone_type"`
					Timezone     string `json:"timezone"`
				} `json:"date_created"`
				DateModified struct {
					Date         string `json:"date"`
					TimezoneType int    `json:"timezone_type"`
					Timezone     string `json:"timezone"`
				} `json:"date_modified"`
				DateExpires struct {
					Date         string `json:"date"`
					TimezoneType int    `json:"timezone_type"`
					Timezone     string `json:"timezone"`
				} `json:"date_expires"`
				DiscountType              string        `json:"discount_type"`
				Description               string        `json:"description"`
				UsageCount                int           `json:"usage_count"`
				IndividualUse             bool          `json:"individual_use"`
				ProductIds                []interface{} `json:"product_ids"`
				ExcludedProductIds        []interface{} `json:"excluded_product_ids"`
				UsageLimit                int           `json:"usage_limit"`
				UsageLimitPerUser         int           `json:"usage_limit_per_user"`
				LimitUsageToXItems        interface{}   `json:"limit_usage_to_x_items"`
				FreeShipping              bool          `json:"free_shipping"`
				ProductCategories         []interface{} `json:"product_categories"`
				ExcludedProductCategories []interface{} `json:"excluded_product_categories"`
				ExcludeSaleItems          bool          `json:"exclude_sale_items"`
				MinimumAmount             string        `json:"minimum_amount"`
				MaximumAmount             string        `json:"maximum_amount"`
				EmailRestrictions         []interface{} `json:"email_restrictions"`
				Virtual                   bool          `json:"virtual"`
				MetaData                  []interface{} `json:"meta_data"`
			} `json:"value"`
			DisplayKey   string `json:"display_key"`
			DisplayValue struct {
				ID          int    `json:"id"`
				Code        string `json:"code"`
				Amount      string `json:"amount"`
				DateCreated struct {
					Date         string `json:"date"`
					TimezoneType int    `json:"timezone_type"`
					Timezone     string `json:"timezone"`
				} `json:"date_created"`
				DateModified struct {
					Date         string `json:"date"`
					TimezoneType int    `json:"timezone_type"`
					Timezone     string `json:"timezone"`
				} `json:"date_modified"`
				DateExpires struct {
					Date         string `json:"date"`
					TimezoneType int    `json:"timezone_type"`
					Timezone     string `json:"timezone"`
				} `json:"date_expires"`
				DiscountType              string        `json:"discount_type"`
				Description               string        `json:"description"`
				UsageCount                int           `json:"usage_count"`
				IndividualUse             bool          `json:"individual_use"`
				ProductIds                []interface{} `json:"product_ids"`
				ExcludedProductIds        []interface{} `json:"excluded_product_ids"`
				UsageLimit                int           `json:"usage_limit"`
				UsageLimitPerUser         int           `json:"usage_limit_per_user"`
				LimitUsageToXItems        interface{}   `json:"limit_usage_to_x_items"`
				FreeShipping              bool          `json:"free_shipping"`
				ProductCategories         []interface{} `json:"product_categories"`
				ExcludedProductCategories []interface{} `json:"excluded_product_categories"`
				ExcludeSaleItems          bool          `json:"exclude_sale_items"`
				MinimumAmount             string        `json:"minimum_amount"`
				MaximumAmount             string        `json:"maximum_amount"`
				EmailRestrictions         []interface{} `json:"email_restrictions"`
				Virtual                   bool          `json:"virtual"`
				MetaData                  []interface{} `json:"meta_data"`
			} `json:"display_value"`
		} `json:"meta_data"`
	} `json:"coupon_lines"`
	Refunds          []interface{} `json:"refunds"`
	DateCreatedGmt   string        `json:"date_created_gmt"`
	DateModifiedGmt  string        `json:"date_modified_gmt"`
	DateCompletedGmt interface{}   `json:"date_completed_gmt"`
	DatePaidGmt      interface{}   `json:"date_paid_gmt"`
	CurrencySymbol   string        `json:"currency_symbol"`
	Links            struct {
		Self []struct {
			Href string `json:"href"`
		} `json:"self"`
		Collection []struct {
			Href string `json:"href"`
		} `json:"collection"`
		Customer []struct {
			Href string `json:"href"`
		} `json:"customer"`
	} `json:"_links,omitempty"`
	Links struct {
		Self []struct {
			Href string `json:"href"`
		} `json:"self"`
		Collection []struct {
			Href string `json:"href"`
		} `json:"collection"`
	} `json:"_links,omitempty"`
}