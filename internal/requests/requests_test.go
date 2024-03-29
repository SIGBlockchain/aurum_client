package requests

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SIGBlockchain/aurum_client/internal/contracts"
	"github.com/SIGBlockchain/aurum_client/internal/publickey"
)

func TestAccountInfoRequest(t *testing.T) {
	req, err := NewAccountInfoRequest("", "xyz")
	if err != nil {
		t.Errorf("failed to create new account info request")
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"received": "`+r.URL.Query().Get("w")+`"}`)
	})
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	expected := `{"received": "xyz"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestNewContractRequest(t *testing.T) {
	senderPrivateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	testContract, err := contracts.MakeContract(1, senderPrivateKey, []byte{1}, 25, 20)
	if err != nil {
		t.Errorf("failed to make contract : %v", err)
	}
	testContract.SignContract(senderPrivateKey)
	req, err := NewContractRequest("", *testContract)
	if err != nil {
		t.Errorf("failed to create test contract: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		io.WriteString(w, buf.String())
	})
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		t.Logf("%s", rr.Body.String())
	}
	var responseBody JSONContract
	if err := json.Unmarshal(rr.Body.Bytes(), &responseBody); err != nil {
		t.Errorf("failed to unmarshall response body: %v", err)
	}
	unhexedResponsePublicKey, err := hex.DecodeString(responseBody.SenderPublicKey)
	if err != nil {
		t.Errorf("failed to hex decode public key: %v", err)
	}
	unhexedResponseSignature, err := hex.DecodeString(responseBody.Signature)
	if err != nil {
		t.Errorf("failed to hex decode signature: %v", err)
	}
	unhexedResponseRecipientHash, err := hex.DecodeString(responseBody.RecipientWalletAddress)
	if err != nil {
		t.Errorf("failed to hex decode recipient hash: %v", err)
	}

	decodedPublicKey, _ := publickey.Decode(unhexedResponsePublicKey)
	// TODO JSONContract to accounts.Contract Unmarshall?
	var responseContract = contracts.Contract{
		responseBody.Version,
		decodedPublicKey,
		responseBody.SignatureLength,
		unhexedResponseSignature,
		unhexedResponseRecipientHash,
		responseBody.Value,
		responseBody.StateNonce,
	}
	if !responseContract.Equals(*testContract) {
		t.Errorf("contracts do not match:\n got %+v want %+v", responseContract, *testContract)
	}
}
