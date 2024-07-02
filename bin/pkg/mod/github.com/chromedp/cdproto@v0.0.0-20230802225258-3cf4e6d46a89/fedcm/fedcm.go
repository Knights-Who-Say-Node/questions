// Package fedcm provides the Chrome DevTools Protocol
// commands, types, and events for the FedCm domain.
//
// This domain allows interacting with the FedCM dialog.
//
// Generated by the cdproto-gen command.
package fedcm

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"

	"github.com/chromedp/cdproto/cdp"
)

// EnableParams [no description].
type EnableParams struct {
	DisableRejectionDelay bool `json:"disableRejectionDelay,omitempty"` // Allows callers to disable the promise rejection delay that would normally happen, if this is unimportant to what's being tested. (step 4 of https://fedidcg.github.io/FedCM/#browser-api-rp-sign-in)
}

// Enable [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/FedCm#method-enable
//
// parameters:
func Enable() *EnableParams {
	return &EnableParams{}
}

// WithDisableRejectionDelay allows callers to disable the promise rejection
// delay that would normally happen, if this is unimportant to what's being
// tested. (step 4 of https://fedidcg.github.io/FedCM/#browser-api-rp-sign-in).
func (p EnableParams) WithDisableRejectionDelay(disableRejectionDelay bool) *EnableParams {
	p.DisableRejectionDelay = disableRejectionDelay
	return &p
}

// Do executes FedCm.enable against the provided context.
func (p *EnableParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandEnable, p, nil)
}

// DisableParams [no description].
type DisableParams struct{}

// Disable [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/FedCm#method-disable
func Disable() *DisableParams {
	return &DisableParams{}
}

// Do executes FedCm.disable against the provided context.
func (p *DisableParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandDisable, nil, nil)
}

// SelectAccountParams [no description].
type SelectAccountParams struct {
	DialogID     string `json:"dialogId"`
	AccountIndex int64  `json:"accountIndex"`
}

// SelectAccount [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/FedCm#method-selectAccount
//
// parameters:
//
//	dialogID
//	accountIndex
func SelectAccount(dialogID string, accountIndex int64) *SelectAccountParams {
	return &SelectAccountParams{
		DialogID:     dialogID,
		AccountIndex: accountIndex,
	}
}

// Do executes FedCm.selectAccount against the provided context.
func (p *SelectAccountParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandSelectAccount, p, nil)
}

// DismissDialogParams [no description].
type DismissDialogParams struct {
	DialogID        string `json:"dialogId"`
	TriggerCooldown bool   `json:"triggerCooldown,omitempty"`
}

// DismissDialog [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/FedCm#method-dismissDialog
//
// parameters:
//
//	dialogID
func DismissDialog(dialogID string) *DismissDialogParams {
	return &DismissDialogParams{
		DialogID: dialogID,
	}
}

// WithTriggerCooldown [no description].
func (p DismissDialogParams) WithTriggerCooldown(triggerCooldown bool) *DismissDialogParams {
	p.TriggerCooldown = triggerCooldown
	return &p
}

// Do executes FedCm.dismissDialog against the provided context.
func (p *DismissDialogParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandDismissDialog, p, nil)
}

// ResetCooldownParams resets the cooldown time, if any, to allow the next
// FedCM call to show a dialog even if one was recently dismissed by the user.
type ResetCooldownParams struct{}

// ResetCooldown resets the cooldown time, if any, to allow the next FedCM
// call to show a dialog even if one was recently dismissed by the user.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/FedCm#method-resetCooldown
func ResetCooldown() *ResetCooldownParams {
	return &ResetCooldownParams{}
}

// Do executes FedCm.resetCooldown against the provided context.
func (p *ResetCooldownParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandResetCooldown, nil, nil)
}

// Command names.
const (
	CommandEnable        = "FedCm.enable"
	CommandDisable       = "FedCm.disable"
	CommandSelectAccount = "FedCm.selectAccount"
	CommandDismissDialog = "FedCm.dismissDialog"
	CommandResetCooldown = "FedCm.resetCooldown"
)
