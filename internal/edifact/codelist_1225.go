package edifact

// UN/EDIFACT D.21A code list 1225 (Message function, coded) -- the real
// values BGM's "Message function code" component (see
// segment_elements_data.go) can take.
//
// Source: https://service.unece.org/trade/untdid/d21a/tred/tred1225.htm
// section "Code Values". That URL 403s via Cloudflare when fetched
// directly; this data was extracted from the Wayback Machine's archived
// copy instead:
// http://web.archive.org/web/20221123154239/https://service.unece.org/trade/untdid/d21a/tred/tred1225.htm
// -- re-check the direct URL first if cross-checking later.
//
// Extracted mechanically (not hand-transcribed) from the page's "Code
// Values:" section by a small script that parses each entry's code,
// name, and indented description, stopping before the page's own
// "Data Element Cross Reference" footer (an early version of the script
// didn't, and picked up two footer tokens as a bogus 70th "code" --
// caught by inspecting the raw extracted output before this was
// committed, not by any automated check).
func init() {
	RegisterCodeList("1225", codeList1225)
}

var codeList1225 = map[string]CodedValue{
	"1":  {Name: "Cancellation", Description: "Message cancelling a previous transmission for a given transaction."},
	"2":  {Name: "Addition", Description: "Message containing items to be added."},
	"3":  {Name: "Deletion", Description: "Message containing items to be deleted."},
	"4":  {Name: "Change", Description: "Message containing items to be changed."},
	"5":  {Name: "Replace", Description: "Message replacing a previous message."},
	"6":  {Name: "Confirmation", Description: "Message confirming the details of a previous transmission where such confirmation is required or recommended under the terms of a trading partner agreement."},
	"7":  {Name: "Duplicate", Description: "The message is a duplicate of a previously generated message."},
	"8":  {Name: "Status", Description: "Code indicating that the referenced message is a status."},
	"9":  {Name: "Original", Description: "Initial transmission related to a given transaction."},
	"10": {Name: "Not found", Description: "Message whose reference number is not filed."},
	"11": {Name: "Response", Description: "Message responding to a previous message or document."},
	"12": {Name: "Not processed", Description: "Message indicating that the referenced message was received but not yet processed."},
	"13": {Name: "Request", Description: "Code indicating that the referenced message is a request."},
	"14": {Name: "Advance notification", Description: "Code indicating that the information contained in the message is an advance notification of information to follow."},
	"15": {Name: "Reminder", Description: "Repeated message transmission for reminding purposes."},
	"16": {Name: "Proposal", Description: "Message content is a proposal."},
	"17": {Name: "Cancel, to be reissued", Description: "Referenced transaction cancelled, reissued message will follow."},
	"18": {Name: "Reissue", Description: "New issue of a previous message (maybe cancelled)."},
	"19": {Name: "Seller initiated change", Description: "Change information submitted by buyer but initiated by seller."},
	"20": {Name: "Replace heading section only", Description: "Message to replace the heading of a previous message."},
	"21": {Name: "Replace item detail and summary only", Description: "Message to replace item detail and summary of a previous message."},
	"22": {Name: "Final transmission", Description: "Final message in a related series of messages together making up a commercial, administrative or transport transaction."},
	"23": {Name: "Transaction on hold", Description: "Message not to be processed until further release information."},
	"24": {Name: "Delivery instruction", Description: "Delivery schedule message only used to transmit short- term delivery instructions."},
	"25": {Name: "Forecast", Description: "Delivery schedule message only used to transmit long- term schedule information."},
	"26": {Name: "Delivery instruction and forecast", Description: "Combination of codes '24' and '25'."},
	"27": {Name: "Not accepted", Description: "Message to inform that the referenced message is not accepted by the recipient."},
	"28": {Name: "Accepted, with amendment in heading section", Description: "Message accepted but amended in heading section."},
	"29": {Name: "Accepted without amendment", Description: "Referenced message is entirely accepted."},
	"30": {Name: "Accepted, with amendment in detail section", Description: "Referenced message is accepted but amended in detail section."},
	"31": {Name: "Copy", Description: "Indicates that the message is a copy of an original message that has been sent, e.g. for action or information."},
	"32": {Name: "Approval", Description: "A message releasing an existing referenced message for action to the receiver."},
	"33": {Name: "Change in heading section", Description: "Message changing the referenced message heading section."},
	"34": {Name: "Accepted with amendment", Description: "The referenced message is accepted but amended."},
	"35": {Name: "Retransmission", Description: "Change-free transmission of a message previously sent."},
	"36": {Name: "Change in detail section", Description: "Message changing referenced detail section."},
	"37": {Name: "Reversal of a debit", Description: "Reversal of a previously posted debit."},
	"38": {Name: "Reversal of a credit", Description: "Reversal of a previously posted credit."},
	"39": {Name: "Reversal for cancellation", Description: "Code indicating that the referenced message is reversing a cancellation of a previous transmission for a given transaction."},
	"40": {Name: "Request for deletion", Description: "The message is given to inform the recipient to delete the referenced transaction."},
	"41": {Name: "Finishing/closing order", Description: "Last of series of call-offs."},
	"42": {Name: "Confirmation via specific means", Description: "Message confirming a transaction previously agreed via other means (e.g. phone)."},
	"43": {Name: "Additional transmission", Description: "Message already transmitted via another communication channel. This transmission is to provide electronically processable data only."},
	"44": {Name: "Accepted without reserves", Description: "Message accepted without reserves."},
	"45": {Name: "Accepted with reserves", Description: "Message accepted with reserves."},
	"46": {Name: "Provisional", Description: "Message content is provisional."},
	"47": {Name: "Definitive", Description: "Message content is definitive."},
	"48": {Name: "Accepted, contents rejected", Description: "Message to inform that the previous message is received, but it cannot be processed due to regulations, laws, etc."},
	"49": {Name: "Settled dispute", Description: "The reported dispute is settled."},
	"50": {Name: "Withdraw", Description: "Message withdrawing a previously approved message."},
	"51": {Name: "Authorisation", Description: "Message authorising a message or transaction(s)."},
	"52": {Name: "Proposed amendment", Description: "A code used to indicate an amendment suggested by the sender."},
	"53": {Name: "Test", Description: "Code indicating the message is to be considered as a test."},
	"54": {Name: "Extract", Description: "A subset of the original."},
	"55": {Name: "Notification only", Description: "The receiver may use the notification information for analysis only."},
	"56": {Name: "Advice of ledger booked items", Description: "An advice that items have been booked in the ledger."},
	"57": {Name: "Advice of items pending to be booked in the ledger", Description: "An advice that items are pending to be booked in the ledger."},
	"58": {Name: "Pre-advice of items requiring further information", Description: "A pre-advice that items require further information."},
	"59": {Name: "Pre-adviced items", Description: "A pre-advice of items."},
	"60": {Name: "No action since last message", Description: "Code indicating the fact that no action has taken place since the last message."},
	"61": {Name: "Complete schedule", Description: "The message function is a complete schedule."},
	"62": {Name: "Update schedule", Description: "The message function is an update to a schedule."},
	"63": {Name: "Not accepted, provisional", Description: "Not accepted, subject to confirmation."},
	"64": {Name: "Verification", Description: "The message is transmitted to verify information."},
	"65": {Name: "Unsettled dispute", Description: "To report an unsettled dispute."},
	"66": {Name: "Discharge of operation guarantee", Description: "A message related to a guarantee containing information about the discharge of an operation."},
	"67": {Name: "Termination of operation guarantee", Description: "A message related to a guarantee containing information about the termination of an operation."},
	"68": {Name: "Start of operation guarantee", Description: "A message related to a guarantee containing information about the start of an operation."},
	"69": {Name: "Advanced cargo information", Description: "A message related to a guarantee containing advanced cargo information."},
}
