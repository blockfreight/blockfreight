// File: ./blockfreight/bf_tx/bf_tx.go
// Summary: Application code for Blockfreight™ | The blockchain of global freight.
// License: MIT License
// Company: Blockfreight, Inc.
// Author: Julian Nunez, Neil Tran, Julian Smith & contributors
// Site: https://blockfreight.com
// Support: <support@blockfreight.com>

// Copyright 2017 Blockfreight, Inc.

// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS
// OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// =================================================================================================================================================
// =================================================================================================================================================
//
// BBBBBBBBBBBb     lll                                kkk             ffff                         iii                  hhh            ttt
// BBBB``````BBBB   lll                                kkk            fff                           ```                  hhh            ttt
// BBBB      BBBB   lll      oooooo        ccccccc     kkk    kkkk  fffffff  rrr  rrr    eeeee      iii     gggggg ggg   hhh  hhhhh   tttttttt
// BBBBBBBBBBBB     lll    ooo    oooo    ccc    ccc   kkk   kkk    fffffff  rrrrrrrr eee    eeee   iii   gggg   ggggg   hhhh   hhhh  tttttttt
// BBBBBBBBBBBBBB   lll   ooo      ooo   ccc           kkkkkkk        fff    rrrr    eeeeeeeeeeeee  iii  gggg      ggg   hhh     hhh    ttt
// BBBB       BBB   lll   ooo      ooo   ccc           kkkk kkkk      fff    rrr     eeeeeeeeeeeee  iii   ggg      ggg   hhh     hhh    ttt
// BBBB      BBBB   lll   oooo    oooo   cccc    ccc   kkk   kkkk     fff    rrr      eee      eee  iii    ggg    gggg   hhh     hhh    tttt    ....
// BBBBBBBBBBBBB    lll     oooooooo       ccccccc     kkk     kkkk   fff    rrr       eeeeeeeee    iii     gggggg ggg   hhh     hhh     ttttt  ....
//                                                                                                        ggg      ggg
//   Blockfreight™ | The blockchain of global freight.                                                      ggggggggg
//
// =================================================================================================================================================
// =================================================================================================================================================

package bf_tx

import (

	// ** COMMENTED OUT - [Tue May 7] - By J.Smith
	// NOTE: Shouldn't this be impoted with the ..blockfreight/crypto/ - PACKAGE ??
	// "crypto/ecdsa"

	"encoding/json"

	"github.com/blockfreight/blockfreight-alpha/blockfreight/bft/common"
)

// Define Blockfreight™ Transaction (BF_TX) transaction standard

func SetBF_TX(jsonpath string) BF_TX {
	var bf_tx BF_TX
	json.Unmarshal(common.ReadJSON(jsonpath), &bf_tx)
	return bf_tx
}

func BF_TXContent(bf_tx BF_TX) string {
	jsonContent, _ := json.Marshal(bf_tx)
	return string(jsonContent)
}

type BF_TX struct {
	Type       string
	Properties Properties
	PrivateKey ecdsa.PrivateKey
	Signhash   []uint8
	Signature  string
	Signed     bool
}

type Properties struct {
	Shipper              Shipper
	Bol_Num              BolNum
	Ref_Num              RefNum
	Consignee            Consignee
	Vessel               Vessel
	Port_of_Loading      PortLoading
	Port_of_Discharge    PortDischarge
	Notify_Address       NotifyAddress
	Desc_of_Goods        DescGoods
	Gross_Weight         GrossWeight
	Freight_Payable_Amt  FreightPayableAmt
	Freight_Adv_Amt      FreightAdvAmt
	General_Instructions GeneralInstructions
	Date_Shipped         Date
	Issue_Details        IssueDetails
	Num_Bol              NumBol // Is it the same Bol_Num?
	Master_Info          MasterInfo
	Agent_for_Master     AgentMaster
	Agent_for_Owner      AgentOwner
}

type Shipper struct {
	Type string
}

type BolNum struct {
	Type int
}

type RefNum struct {
	Type int
}

type Consignee struct {
	Type string //Null
}

type Vessel struct {
	Type int
}

type PortLoading struct {
	Type int
}

type PortDischarge struct {
	Type int
}

type NotifyAddress struct {
	Type string
}

type DescGoods struct {
	Type string
}

type GrossWeight struct {
	Type int //Should it be float?
}

type FreightPayableAmt struct {
	Type int
}

type FreightAdvAmt struct {
	Type int
}

type GeneralInstructions struct {
	Type string
}

type Date struct {
	Type   int
	Format string
}

type IssueDetails struct {
	Type       string
	Properties IssueDetailsProperties
}

type IssueDetailsProperties struct {
	Place_of_Issue PlaceIssue
	Date_of_Issue  Date
}

type PlaceIssue struct {
	Type string
}

type NumBol struct {
	Type int
}

type MasterInfo struct {
	Type       string
	Properties MasterInfoProperties
}

type MasterInfoProperties struct {
	First_Name FirstName
	Last_Name  LastName
	Sig        Sig
}

type AgentMaster struct {
	Type       string
	Properties AgentMasterProperties
}

type AgentMasterProperties struct {
	First_Name FirstName
	Last_Name  LastName
	Sig        Sig
}

type AgentOwner struct {
	Type       string
	Properties AgentOwnerProperties
}

type AgentOwnerProperties struct {
	First_Name              FirstName
	Last_Name               LastName
	Sig                     Sig
	Conditions_for_Carriage ConditionsCarriage
}

type FirstName struct {
	Type string
}

type LastName struct {
	Type string
}

type Sig struct {
	Type string
}

type ConditionsCarriage struct {
	Type string
}

// =================================================
// Blockfreight™ | The blockchain of global freight.
// =================================================

// BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
// BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
// BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
// BBBBBBB                    BBBBBBBBBBBBBBBBBBB
// BBBBBBB                       BBBBBBBBBBBBBBBB
// BBBBBBB                        BBBBBBBBBBBBBBB
// BBBBBBB       BBBBBBBBB        BBBBBBBBBBBBBBB
// BBBBBBB       BBBBBBBBB        BBBBBBBBBBBBBBB
// BBBBBBB       BBBBBBB         BBBBBBBBBBBBBBBB
// BBBBBBB                     BBBBBBBBBBBBBBBBBB
// BBBBBBB                        BBBBBBBBBBBBBBB
// BBBBBBB       BBBBBBBBBB        BBBBBBBBBBBBBB
// BBBBBBB       BBBBBBBBBBB       BBBBBBBBBBBBBB
// BBBBBBB       BBBBBBBBBB        BBBBBBBBBBBBBB
// BBBBBBB       BBBBBBBBB        BBB       BBBBB
// BBBBBBB                       BBBB       BBBBB
// BBBBBBB                    BBBBBBB       BBBBB
// BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
// BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
// BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB

// ==================================================
// Blockfreight™ | The blockchain for global freight.
// ==================================================