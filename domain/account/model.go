package domain

type AccountStatement struct {
	Account         Account       `json:"account"`
	Transactions    []Transaction `json:"transactions"`
	IBankBic1       string        `json:"i_bank_bic_1"`
	IBankBic2       string        `json:"i_bank_bic_2"`
	IBankBic3       string        `json:"i_bank_bic_3"`
	StatementPeriod string        `json:"statement_period"`
	Address         string        `json:"address"`
	Name            string        `json:"name"`
	BIC1            string        `json:"bic_1"`
	BIC2            string        `json:"bic_2"`
	BIC3            string        `json:"bic_3"`
	Currency1       string        `json:"currency_1"`
}

type Transaction struct {
	Date        string  `json:"date"`
	Description string  `json:"description"`
	MoneyOut    float64 `json:"money_out"`
	MoneyIn     float64 `json:"money_in"`
	Balance     float64 `json:"balance"`
	Currency    string  `json:"currency"`
}

type Account struct {
	OpeningBalance float64 `json:"opening_balance"`
	ClosingBalance float64 `json:"closing_balance"`
	AccountName    string  `json:"account_name"`
	StatementDate  string  `json:"statement_date"`
	MoneyOut       float64 `json:"money_out"`
	MoneyIn        float64 `json:"money_in"`
	Product        string  `json:"product"`
	Balance        float64 `json:"balance"`
}

type QueryParams struct {
	Currency string `form:"currency"`
}
