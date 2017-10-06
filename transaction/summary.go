package transaction

type Summary []Transaction

func (s Summary) Amount() int {
	balance := 0

	for _, tx := range s {
		balance += int(tx.Amount())
	}

	return balance
}
