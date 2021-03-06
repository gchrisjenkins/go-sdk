package keys

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	TestMnenomic     = "equip town gesture square tomorrow volume nephew minute witness beef rich gadget actress egg sing secret pole winter alarm law today check violin uncover"
	TestExpectedAddr = "kava15qdefkmwswysgg4qxgqpqr35k3m49pkx2jdfnw"
	TestKavaCoinID   = 459
)

func TestNewMnemonicKeyManager(t *testing.T) {

	tests := []struct {
		name       string
		mnenomic   string
		coinID     uint32
		expectpass bool
	}{
		{"normal", TestMnenomic, TestKavaCoinID, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			keyManager, err := NewMnemonicKeyManager(tc.mnenomic, tc.coinID)

			if tc.expectpass {
				require.Nil(t, err)

				// Confirm correct address
				addr := keyManager.GetAddr()
				require.Equal(t, TestExpectedAddr, addr.String())
			} else {
				require.NotNil(t, err)
			}
		})
	}
}
