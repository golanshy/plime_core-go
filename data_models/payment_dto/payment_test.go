package payment_dto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPaymentsRequest_Validate_invalid_reference(t *testing.T) {
	var request = PublicPaymentsRequest{
		Reference: "",
	}
	err := request.Validate()
	assert.NotNil(t, err)
	assert.EqualValues(t, "invalid reference data", err.Message)
}

func TestPaymentsRequest_Validate_invalid_payments(t *testing.T) {
	var request = PublicPaymentsRequest{
		Reference: "123",
	}
	err := request.Validate()
	assert.NotNil(t, err)
	assert.EqualValues(t, "invalid payments data", err.Message)
}

func TestPaymentsRequest_Validate_zero_payments(t *testing.T) {
	var payments = make([]PublicPaymentRequest, 0)
	var request = PublicPaymentsRequest{
		Reference: "123",
		Payments:  payments,
	}
	err := request.Validate()
	assert.NotNil(t, err)
	assert.EqualValues(t, "invalid payments data", err.Message)
}

func TestPaymentsRequest_Validate_invalid_payment_email(t *testing.T) {
	var payments = make([]PublicPaymentRequest, 0)
	payments = append(payments, PublicPaymentRequest{
		Payee: Payee{
			Type:                0,
			Details:             "",
			CompanyName:         "",
			CompanyRegisteredId: "",
			FirstName:           "",
			LastName:            "",
			Email:               "",
			Mobile:              "",
			CountryCode:         "",
		},
		Reference:    "",
		Details:      "",
		Amount:       0,
		CurrencyCode: "",
		ArriveBy:     "",
		SendOn:       "",
	})
	var request = PublicPaymentsRequest{
		Reference: "123",
		Payments:  payments,
	}
	err := request.Validate()
	assert.NotNil(t, err)
	assert.EqualValues(t, "invalid payee email address for payment 0", err.Message)
}

func TestPaymentsRequest_Validate_invalid_payment_amount(t *testing.T) {
	var payments = make([]PublicPaymentRequest, 0)
	payments = append(payments, PublicPaymentRequest{
		Payee: Payee{
			Type:                0,
			Details:             "",
			CompanyName:         "",
			CompanyRegisteredId: "",
			FirstName:           "",
			LastName:            "",
			Email:               "abc@domain.com",
			Mobile:              "",
			CountryCode:         "",
		},
		Reference:    "456",
		Details:      "",
		Amount:       0,
		CurrencyCode: "",
		ArriveBy:     "",
		SendOn:       "",
	})
	var request = PublicPaymentsRequest{
		Reference: "123",
		Payments:  payments,
	}
	err := request.Validate()
	assert.NotNil(t, err)
	assert.EqualValues(t, "invalid amount for payment 0", err.Message)
}

func TestPaymentsRequest_Validate_invalid_payment_currency_code(t *testing.T) {
	var payments = make([]PublicPaymentRequest, 0)
	payments = append(payments, PublicPaymentRequest{
		Payee: Payee{
			Type:                0,
			Details:             "",
			CompanyName:         "",
			CompanyRegisteredId: "",
			FirstName:           "",
			LastName:            "",
			Email:               "abc@domain.com",
			Mobile:              "",
			CountryCode:         "",
		},
		Reference:    "456",
		Details:      "",
		Amount:       0.01,
		CurrencyCode: "",
		ArriveBy:     "",
		SendOn:       "",
	})
	var request = PublicPaymentsRequest{
		Reference: "123",
		Payments:  payments,
	}
	err := request.Validate()
	assert.NotNil(t, err)
	assert.EqualValues(t, "invalid currency code for payment 0", err.Message)
}

func TestPaymentsRequest_Validate_no_payment_arrive_by(t *testing.T) {
	var payments = make([]PublicPaymentRequest, 0)
	payments = append(payments, PublicPaymentRequest{
		Payee: Payee{
			Type:                0,
			Details:             "",
			CompanyName:         "",
			CompanyRegisteredId: "",
			FirstName:           "",
			LastName:            "",
			Email:               "abc@domain.com",
			Mobile:              "",
			CountryCode:         "",
		},
		Reference:    "456",
		Details:      "",
		Amount:       0.01,
		CurrencyCode: "GBP",
		ArriveBy:     "",
		SendOn:       "",
	})
	var request = PublicPaymentsRequest{
		Reference: "123",
		Payments:  payments,
	}
	err := request.Validate()
	assert.Nil(t, err)
}

func TestPaymentsRequest_Validate_invalid_payment_arrive_by_format(t *testing.T) {
	var payments = make([]PublicPaymentRequest, 0)
	payments = append(payments, PublicPaymentRequest{
		Payee: Payee{
			Type:                0,
			Details:             "",
			CompanyName:         "",
			CompanyRegisteredId: "",
			FirstName:           "",
			LastName:            "",
			Email:               "abc@domain.com",
			Mobile:              "",
			CountryCode:         "",
		},
		Reference:    "456",
		Details:      "",
		Amount:       0.01,
		CurrencyCode: "GBP",
		ArriveBy:     "111",
		SendOn:       "",
	})
	var request = PublicPaymentsRequest{
		Reference: "123",
		Payments:  payments,
	}
	err := request.Validate()
	assert.NotNil(t, err)
	assert.EqualValues(t, "invalid date format, invalid arrive by time for payment 0", err.Message)
}

func TestPaymentsRequest_Validate_payment_arrive_by_date_in_past(t *testing.T) {
	var payments = make([]PublicPaymentRequest, 0)
	payments = append(payments, PublicPaymentRequest{
		Payee: Payee{
			Type:                0,
			Details:             "",
			CompanyName:         "",
			CompanyRegisteredId: "",
			FirstName:           "",
			LastName:            "",
			Email:               "abc@domain.com",
			Mobile:              "",
			CountryCode:         "",
		},
		Reference:    "456",
		Details:      "",
		Amount:       0.01,
		CurrencyCode: "GBP",
		ArriveBy:     "2006-01-01T00:00:00.000Z",
		SendOn:       "",
	})
	var request = PublicPaymentsRequest{
		Reference: "123",
		Payments:  payments,
	}
	err := request.Validate()
	assert.NotNil(t, err)
	assert.EqualValues(t, "arrive by should be in the future for payment 0", err.Message)
}

func TestPaymentsRequest_Validate_valid_payment_arrive_by_date(t *testing.T) {
	var payments = make([]PublicPaymentRequest, 0)
	payments = append(payments, PublicPaymentRequest{
		Payee: Payee{
			Type:                0,
			Details:             "",
			CompanyName:         "",
			CompanyRegisteredId: "",
			FirstName:           "",
			LastName:            "",
			Email:               "abc@domain.com",
			Mobile:              "",
			CountryCode:         "",
		},
		Reference:    "456",
		Details:      "",
		Amount:       0.01,
		CurrencyCode: "GBP",
		ArriveBy:     "2106-01-01T00:00:00.000Z",
		SendOn:       "",
	})
	var request = PublicPaymentsRequest{
		Reference: "123",
		Payments:  payments,
	}
	err := request.Validate()
	assert.Nil(t, err)
}

// --------------------------------------------------------------------------------

func TestPaymentsRequest_Validate_no_payment_send_on(t *testing.T) {
	var payments = make([]PublicPaymentRequest, 0)
	payments = append(payments, PublicPaymentRequest{
		Payee: Payee{
			Type:                0,
			Details:             "",
			CompanyName:         "",
			CompanyRegisteredId: "",
			FirstName:           "",
			LastName:            "",
			Email:               "abc@domain.com",
			Mobile:              "",
			CountryCode:         "",
		},
		Reference:    "456",
		Details:      "",
		Amount:       0.01,
		CurrencyCode: "GBP",
		ArriveBy:     "2106-01-01T00:00:00.000Z",
		SendOn:       "",
	})
	var request = PublicPaymentsRequest{
		Reference: "123",
		Payments:  payments,
	}
	err := request.Validate()
	assert.Nil(t, err)
}

func TestPaymentsRequest_Validate_invalid_payment_send_on_format(t *testing.T) {
	var payments = make([]PublicPaymentRequest, 0)
	payments = append(payments, PublicPaymentRequest{
		Payee: Payee{
			Type:                0,
			Details:             "",
			CompanyName:         "",
			CompanyRegisteredId: "",
			FirstName:           "",
			LastName:            "",
			Email:               "abc@domain.com",
			Mobile:              "",
			CountryCode:         "",
		},
		Reference:    "456",
		Details:      "",
		Amount:       0.01,
		CurrencyCode: "GBP",
		ArriveBy:     "2106-01-01T00:00:00.000Z",
		SendOn:       "111",
	})
	var request = PublicPaymentsRequest{
		Reference: "123",
		Payments:  payments,
	}
	err := request.Validate()
	assert.NotNil(t, err)
	assert.EqualValues(t, "invalid date format, invalid send on for payment 0", err.Message)
}

func TestPaymentsRequest_Validate_payment_arrive_by_send_on_past(t *testing.T) {
	var payments = make([]PublicPaymentRequest, 0)
	payments = append(payments, PublicPaymentRequest{
		Payee: Payee{
			Type:                0,
			Details:             "",
			CompanyName:         "",
			CompanyRegisteredId: "",
			FirstName:           "",
			LastName:            "",
			Email:               "abc@domain.com",
			Mobile:              "",
			CountryCode:         "",
		},
		Reference:    "456",
		Details:      "",
		Amount:       0.01,
		CurrencyCode: "GBP",
		SendOn:       "2006-01-01T00:00:00.000Z",
		ArriveBy:     "2106-01-01T00:00:00.000Z",
	})
	var request = PublicPaymentsRequest{
		Reference: "123",
		Payments:  payments,
	}
	err := request.Validate()
	assert.NotNil(t, err)
	assert.EqualValues(t, "send on should be in the future for payment 0", err.Message)
}

func TestPaymentsRequest_Validate_valid_payment_send_on_date(t *testing.T) {
	var payments = make([]PublicPaymentRequest, 0)
	payments = append(payments, PublicPaymentRequest{
		Payee: Payee{
			Type:                0,
			Details:             "",
			CompanyName:         "",
			CompanyRegisteredId: "",
			FirstName:           "",
			LastName:            "",
			Email:               "abc@domain.com",
			Mobile:              "",
			CountryCode:         "",
		},
		Reference:    "456",
		Details:      "",
		Amount:       0.01,
		CurrencyCode: "GBP",
		SendOn:       "2106-01-01T00:00:00.000Z",
		ArriveBy:     "2106-01-01T00:00:00.001Z",
	})
	var request = PublicPaymentsRequest{
		Reference: "123",
		Payments:  payments,
	}
	err := request.Validate()
	assert.Nil(t, err)
}

func TestPaymentsRequest_Validate_payment_send_on_date_after_arrive_by(t *testing.T) {
	var payments = make([]PublicPaymentRequest, 0)
	payments = append(payments, PublicPaymentRequest{
		Payee: Payee{
			Type:                0,
			Details:             "",
			CompanyName:         "",
			CompanyRegisteredId: "",
			FirstName:           "",
			LastName:            "",
			Email:               "abc@domain.com",
			Mobile:              "",
			CountryCode:         "",
		},
		Reference:    "456",
		Details:      "",
		Amount:       0.01,
		CurrencyCode: "GBP",
		SendOn:       "2206-01-01T00:00:00.000Z",
		ArriveBy:     "2106-01-01T00:00:00.000Z",
	})
	var request = PublicPaymentsRequest{
		Reference: "123",
		Payments:  payments,
	}
	err := request.Validate()
	assert.NotNil(t, err)
}

func TestPaymentsRequest_Validate_payment_send_on_date_before_arrive_by(t *testing.T) {
	var payments = make([]PublicPaymentRequest, 0)
	payments = append(payments, PublicPaymentRequest{
		Payee: Payee{
			Type:                0,
			Details:             "",
			CompanyName:         "",
			CompanyRegisteredId: "",
			FirstName:           "",
			LastName:            "",
			Email:               "abc@domain.com",
			Mobile:              "",
			CountryCode:         "",
		},
		Reference:    "456",
		Details:      "",
		Amount:       0.01,
		CurrencyCode: "GBP",
		SendOn:       "2106-01-01T00:00:00.000Z",
		ArriveBy:     "2206-01-01T00:00:00.000Z",
	})
	var request = PublicPaymentsRequest{
		Reference: "123",
		Payments:  payments,
	}
	err := request.Validate()
	assert.Nil(t, err)
}
