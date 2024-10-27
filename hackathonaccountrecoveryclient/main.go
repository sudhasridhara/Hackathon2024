package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"google.golang.org/appengine"
	"google.golang.org/appengine/mail"

	// Importing the general purpose Cosmos blockchain client
	"github.com/ignite/cli/v28/ignite/pkg/cosmosclient"

	// Importing the types package of your blog blockchain
	"hackathonaccountrecovery/x/hackathonaccountrecovery/types"
)

type AccountRecoveryShareKey struct {
	OwnerAddress  string `json:"owneraddress"`
	Sharekey      string `json:"sharekey"`
	SignInAddress string `json:"signInAddress"`
	Id            string `json:"Id"`
}

type KeyValue struct {
	Address  string `json:"address"`
	Sharekey string `json:"sharekey"`
}

type Acccount struct {
	AccountAdd string `json:"address"`
}

type AcccountRecovery struct {
	AccountAdd string `json:"address"`
	DeathDate  string `json:"deathDate"`
}

type RecoveryData struct {
	Owner            string     `json:"Owner"`
	NumberOfAccounts string     `json:"numberOfAccounts"`
	MinimumAccounts  string     `json:"minimumAccounts"`
	Keys             []KeyValue `json:"keys"`
	EmailId          string     `json:"emailId"`
	PrivateKey       string     `json:"privateKey"`
}

// Handler function to process the POST request
func handleRecovery(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the incoming JSON data
	var recoveryData RecoveryData
	err := json.NewDecoder(r.Body).Decode(&recoveryData)
	if err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	// Log or process the parsed data
	fmt.Printf("Received recovery data: %+v\n", recoveryData)

	ctx := context.Background()
	addressPrefix := "cosmos"

	// Create a Cosmos client instance
	client, err := cosmosclient.New(ctx, cosmosclient.WithAddressPrefix(addressPrefix))
	if err != nil {
		log.Fatal(err)
	}

	// Account `alice` was initialized during `ignite chain serve`
	accountName := "alice"

	// Get account from the keyring
	account, err := client.Account(accountName)
	if err != nil {
		log.Fatal(err)
	}

	addr, err := account.Address(addressPrefix)
	if err != nil {
		log.Fatal(err)
	}
	min, _ := strconv.ParseUint(recoveryData.MinimumAccounts, 10, 64)
	total, _ := strconv.ParseUint(recoveryData.NumberOfAccounts, 10, 64)

	queryClient := types.NewQueryClient(client.Context())
	//	queryResp, err := queryClient.ListAccount(ctx, &types.QueryListAccountRequest{})

	// Define a message to create a post
	msg := &types.MsgAddAccountInfo{
		Creator:     addr,
		Owner:       recoveryData.Owner,
		Totalshares: total,
		Minshares:   min,
	}

	// Broadcast a transaction from account `alice` with the message
	// to create a post store response in txResp
	txResp, err := client.BroadcastTx(ctx, account, msg)
	if err != nil {
		log.Fatal(err)
	}

	emailmsg := &types.MsgUpdateEmailId{
		Creator: addr,
		Account: recoveryData.Owner,
		Emailid: recoveryData.EmailId,
	}
	// Broadcast a transaction from account `alice` with the message
	// to create a post store response in txResp
	txmailResp, err := client.BroadcastTx(ctx, account, emailmsg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(txmailResp)

	//Id, _ := strconv.ParseUint(txResp, 10, 64)
	jsonData, err := json.Marshal(recoveryData.Keys)

	//queryClient1 := types.NewQueryClient(client.Context())
	queryResp, err := queryClient.ShowAccount(ctx, &types.QueryShowAccountRequest{Owner: recoveryData.Owner})
	if err != nil {
		log.Fatal(err)
	}
	//log.Printf("Event registered: %+v\n", queryResp)

	/*accountId, _ := strconv.ParseUint("0", 10, 64)
	if queryResp.Account.Id != 0 {
		accountId = queryResp.Account.Id
	}*/
	multisigMsg := &types.MsgCreateMultisite{
		Creator:   addr,
		Account:   recoveryData.Owner,
		Id:        queryResp.Account.Id,
		Multisign: string(jsonData),
	}

	// Broadcast a transaction from account `alice` with the message
	// to create a post store response in txResp
	txResponse, err := client.BroadcastTx(ctx, account, multisigMsg)
	if err != nil {
		log.Fatal(err)
	}

	// Print response from broadcasting a transaction
	fmt.Print("MsgCreateAccount:\n\n")
	fmt.Println(txResponse)
	fmt.Println(txResp)

	shareKeymsg := &types.QueryGetShareKeysRequest{
		Key:         recoveryData.PrivateKey,
		Totalshares: int64(total),
		Minshares:   int64(min),
	}

	queryRes, err := queryClient.GetShareKeys(ctx, shareKeymsg)
	if err != nil {
		log.Fatal(err)
	}
	// Send a response back to the client
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(queryRes)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/submit-recovery", handleRecovery)
	http.HandleFunc("/get-account", getAccountDetailsHandler)
	http.HandleFunc("/death-initiation", initiateDeathHandler)
	http.HandleFunc("/get-map-account", getMappedAccountDetailsHandler)
	http.HandleFunc("/recover-account", restoreSharekeyHandler)
	http.ListenAndServe(":8080", nil)
}

type PageData struct {
	Title   string
	Content string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title:   "Account Recovery Application",
		Content: "This is a simple app developed for Hackathon.",
	}

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getAccountDetailsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the incoming JSON data
	var acccount Acccount
	err := json.NewDecoder(r.Body).Decode(&acccount)
	if err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	// Log or process the parsed data
	fmt.Printf("Received recovery data: %+v\n", acccount)

	ctx := context.Background()
	addressPrefix := "cosmos"

	// Create a Cosmos client instance
	client, err := cosmosclient.New(ctx, cosmosclient.WithAddressPrefix(addressPrefix))
	if err != nil {
		log.Fatal(err)
	}

	queryClient := types.NewQueryClient(client.Context())
	queryResp, err := queryClient.ShowAccount(ctx, &types.QueryShowAccountRequest{Owner: acccount.AccountAdd})
	if err != nil {
		log.Fatal(err)
	}

	// Send a response back to the client
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(queryResp)
}

func initiateDeathHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the incoming JSON data
	var acccount AcccountRecovery
	err := json.NewDecoder(r.Body).Decode(&acccount)
	if err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	// Log or process the parsed data
	fmt.Printf("Received recovery data: %+v\n", acccount)

	ctx := context.Background()
	addressPrefix := "cosmos"

	// Create a Cosmos client instance
	client, err := cosmosclient.New(ctx, cosmosclient.WithAddressPrefix(addressPrefix))
	if err != nil {
		log.Fatal(err)
	}

	// Account `alice` was initialized during `ignite chain serve`
	accountName := "alice"

	// Get account from the keyring
	account, err := client.Account(accountName)
	if err != nil {
		log.Fatal(err)
	}

	addr, err := account.Address(addressPrefix)
	if err != nil {
		log.Fatal(err)
	}

	queryClient := types.NewQueryClient(client.Context())
	queryResp, err := queryClient.ShowAccount(ctx, &types.QueryShowAccountRequest{Owner: acccount.AccountAdd})
	if err != nil {
		log.Fatal(err)
	}

	deathMsg := &types.MsgInitiateDeathValidation{
		Creator:   addr,
		Owner:     acccount.AccountAdd,
		Id:        queryResp.Account.Id,
		DeathDate: acccount.DeathDate,
	}

	txResponse, err := client.BroadcastTx(ctx, account, deathMsg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(txResponse)

	ValidateMsg := &types.MsgValidateDeath{
		Creator:     addr,
		Owner:       acccount.AccountAdd,
		Id:          queryResp.Account.Id,
		Isvalidated: true,
	}

	txResp, err := client.BroadcastTx(ctx, account, ValidateMsg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(txResp)

	// Send a response back to the client
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})

}

func getMappedAccountDetailsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the incoming JSON data
	var acccount Acccount
	err := json.NewDecoder(r.Body).Decode(&acccount)
	if err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	// Log or process the parsed data
	fmt.Printf("Received recovery data: %+v\n", acccount)

	ctx := context.Background()
	addressPrefix := "cosmos"

	// Create a Cosmos client instance
	client, err := cosmosclient.New(ctx, cosmosclient.WithAddressPrefix(addressPrefix))
	if err != nil {
		log.Fatal(err)
	}

	queryClient := types.NewQueryClient(client.Context())
	queryResp, err := queryClient.GetMappedAccounts(ctx, &types.QueryGetMappedAccountsRequest{Account: acccount.AccountAdd})
	if err != nil {
		log.Fatal(err)
	}

	// Send a response back to the client
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(queryResp)
}

func restoreSharekeyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the incoming JSON data
	var recoveryAccount AccountRecoveryShareKey
	err := json.NewDecoder(r.Body).Decode(&recoveryAccount)
	if err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	// Log or process the parsed data
	fmt.Printf("Received recovery data: %+v\n", recoveryAccount)

	ctx := context.Background()
	addressPrefix := "cosmos"

	// Create a Cosmos client instance
	client, err := cosmosclient.New(ctx, cosmosclient.WithAddressPrefix(addressPrefix))
	if err != nil {
		log.Fatal(err)
	}

	// Account `alice` was initialized during `ignite chain serve`
	accountName := "alice"

	// Get account from the keyring
	account, err := client.Account(accountName)
	if err != nil {
		log.Fatal(err)
	}

	addr, err := account.Address(addressPrefix)
	if err != nil {
		log.Fatal(err)
	}
	print(addr)
	Id, _ := strconv.ParseUint(recoveryAccount.Id, 10, 64)

	ValidateMsg := &types.MsgUpdateMultisite{
		Creator:           addr,
		Account:           recoveryAccount.OwnerAddress,
		Id:                Id,
		MultisignAddress:  recoveryAccount.SignInAddress,
		MultisignShareKey: recoveryAccount.Sharekey,
	}

	txResp, err := client.BroadcastTx(ctx, account, ValidateMsg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(txResp)

	queryClient := types.NewQueryClient(client.Context())
	queryResp, err := queryClient.GetSecretKey(ctx, &types.QueryGetSecretKeyRequest{AccountOwner: recoveryAccount.OwnerAddress})
	if err != nil {
		log.Fatal(err)
	}

	// Send a response back to the client
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(queryResp)
}

func sendEmail(to, subject, body string, r *http.Request) error {
	// Sender data
	ctx := appengine.NewContext(r)
	msg := &mail.Message{
		Sender:   "harrysravs4@gmail.com",
		To:       []string{"sudhas.sridhara@gmail.com"},
		Subject:  subject,
		Body:     body,
		HTMLBody: "<html><body><p>" + body + "</p></body></html>",
	}

	// Send the email
	if err := mail.Send(ctx, msg); err != nil {
		fmt.Print(err)
	}
	return nil
}
