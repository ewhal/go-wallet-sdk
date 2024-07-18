// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package nft_candy_machine_v2

import (
	"errors"

	"github.com/ewhal/go-wallet-sdk/coins/solana/base"
)

// UpdateCandyMachine is the `updateCandyMachine` instruction.
type UpdateCandyMachine struct {
	Data *CandyMachineData

	// [0] = [WRITE] candyMachine
	//
	// [1] = [SIGNER] authority
	//
	// [2] = [] wallet
	base.AccountMetaSlice `bin:"-"`
}

// NewUpdateCandyMachineInstructionBuilder creates a new `UpdateCandyMachine` instruction builder.
func NewUpdateCandyMachineInstructionBuilder() *UpdateCandyMachine {
	nd := &UpdateCandyMachine{
		AccountMetaSlice: make(base.AccountMetaSlice, 3),
	}
	return nd
}

// SetData sets the "data" parameter.
func (inst *UpdateCandyMachine) SetData(data CandyMachineData) *UpdateCandyMachine {
	inst.Data = &data
	return inst
}

// SetCandyMachineAccount sets the "candyMachine" account.
func (inst *UpdateCandyMachine) SetCandyMachineAccount(candyMachine base.PublicKey) *UpdateCandyMachine {
	inst.AccountMetaSlice[0] = base.Meta(candyMachine).WRITE()
	return inst
}

// GetCandyMachineAccount gets the "candyMachine" account.
func (inst *UpdateCandyMachine) GetCandyMachineAccount() *base.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *UpdateCandyMachine) SetAuthorityAccount(authority base.PublicKey) *UpdateCandyMachine {
	inst.AccountMetaSlice[1] = base.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *UpdateCandyMachine) GetAuthorityAccount() *base.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetWalletAccount sets the "wallet" account.
func (inst *UpdateCandyMachine) SetWalletAccount(wallet base.PublicKey) *UpdateCandyMachine {
	inst.AccountMetaSlice[2] = base.Meta(wallet)
	return inst
}

// GetWalletAccount gets the "wallet" account.
func (inst *UpdateCandyMachine) GetWalletAccount() *base.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst UpdateCandyMachine) Build() *Instruction {
	return &Instruction{BaseVariant: base.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UpdateCandyMachine,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UpdateCandyMachine) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdateCandyMachine) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Data == nil {
			return errors.New("Data parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.CandyMachine is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Wallet is not set")
		}
	}
	return nil
}

func (obj UpdateCandyMachine) MarshalWithEncoder(encoder *base.Encoder) (err error) {
	// Serialize `Data` param:
	err = encoder.Encode(obj.Data)
	if err != nil {
		return err
	}
	return nil
}

// NewUpdateCandyMachineInstruction declares a new UpdateCandyMachine instruction with the provided parameters and accounts.
func NewUpdateCandyMachineInstruction(
	// Parameters:
	data CandyMachineData,
	// Accounts:
	candyMachine base.PublicKey,
	authority base.PublicKey,
	wallet base.PublicKey) *UpdateCandyMachine {
	return NewUpdateCandyMachineInstructionBuilder().
		SetData(data).
		SetCandyMachineAccount(candyMachine).
		SetAuthorityAccount(authority).
		SetWalletAccount(wallet)
}
