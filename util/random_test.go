package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomInt(t *testing.T) {
	var min, max int64
	min = 2
	max = 5

	randomInt := RandomInt(min, max)

	require.Positive(t, randomInt)
	require.GreaterOrEqual(t, randomInt, min)
	require.LessOrEqual(t, randomInt, max)
}

func TestRandomString(t *testing.T) {
	var len int = 5

	randomStr := RandomString(len)

	require.Len(t, randomStr, len)
	require.NotEmpty(t, randomStr)
}

func TestRandomOwner(t *testing.T) {
	ownerLen := 6
	owner := RandomOwner()
	require.Len(t, owner, ownerLen)
	require.NotEmpty(t, owner)
}

func TestRandomMoney(t *testing.T) {

	var min, max int64
	min = 0
	max = 1000

	money := RandomMoney()

	require.NotEmpty(t, money)
	require.Positive(t, money)
	require.GreaterOrEqual(t, money, min)
	require.LessOrEqual(t, money, max)
}

func TestRandomCurrency(t *testing.T) {
	currencyLen := 3

	currency := RandomCurrency()

	require.NotEmpty(t, currency)
	require.Len(t, currency, currencyLen)
}
