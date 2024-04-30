package d_taiwan

import (
	"encoding/hex"
	"fmt"
	"gateway/global"
	dtaiwanmodel "gateway/model/dnf/d_taiwan"
	taiwanbilling "gateway/model/dnf/taiwan_billing"
	taiwanlogin "gateway/model/dnf/taiwan_login"
	"sync"
	"time"
)

type AccountService struct {
}

type CreateAccountRequest struct {
	CashCera      int32 `json:"cashCera"`      // 0
	CashCeraPoint int32 `json:"cashCeraPoint"` // 0
	dtaiwanmodel.Account
}

// GetAccount 根据账号名获取账号信息
func (a *AccountService) GetAccount(body dtaiwanmodel.Account) (data dtaiwanmodel.Account, err error) {
	err = global.GatewayDBs["d_taiwan"].Where("accountname = ?", body.Accountname).First(&data).Error
	return data, err
}

// CreateAccount 创建账号
//func (a *AccountService) CreateAccount(body CreateAccountRequest) (data dtaiwanmodel.Account, err error) {
//	err = global.GatewayDBs["d_taiwan"].Create(&body.Account).Error
//	if err != nil {
//		return data, err
//	}
//
//	fmt.Println("注册的用户UID", body.UID)
//
//	// 创建角色限制 (LimitCreateCharacter) 默认0
//	var limitCreateCharacter dtaiwanmodel.LimitCreateCharacter
//	limitCreateCharacter = dtaiwanmodel.LimitCreateCharacter{MID: body.UID}
//	err = global.GatewayDBs["d_taiwan"].Create(&limitCreateCharacter).Error
//	if err != nil {
//		return data, err
//	}
//
//	// 创建角色信息 (CharacterInfo) 默认0
//	var memberInfo dtaiwanmodel.MemberInfo
//	memberInfo = dtaiwanmodel.MemberInfo{MID: body.UID, UserID: fmt.Sprintf("%d", body.UID)}
//	err = global.GatewayDBs["d_taiwan"].Create(&memberInfo).Error
//	if err != nil {
//		return data, err
//	}
//
//	// 创建角色加入信息 (MemberJoinInfo) 默认0
//	var memberJoinInfo dtaiwanmodel.MemberJoinInfo
//	memberJoinInfo = dtaiwanmodel.MemberJoinInfo{MID: body.UID}
//	err = global.GatewayDBs["d_taiwan"].Create(&memberJoinInfo).Error
//	if err != nil {
//		return data, err
//	}
//
//	var memberMiles dtaiwanmodel.MemberMile
//	memberMiles = dtaiwanmodel.MemberMile{MID: body.UID}
//	err = global.GatewayDBs["d_taiwan"].Create(&memberMiles).Error
//	if err != nil {
//		return data, err
//	}
//
//	// 创建角色白名单 (MemberWhiteAccount) 默认0
//	var memberWhiteAccount dtaiwanmodel.MemberWhiteAccount
//	memberWhiteAccount = dtaiwanmodel.MemberWhiteAccount{MID: body.UID}
//	err = global.GatewayDBs["d_taiwan"].Create(&memberWhiteAccount).Error
//	if err != nil {
//		return data, err
//	}
//
//	// 创建角色登录信息 (MemberLogin) 默认0
//	var memberLogin taiwanlogin.MemberLogin
//	memberLogin = taiwanlogin.MemberLogin{MID: body.UID}
//	err = global.GatewayDBs["taiwan_login"].Create(&memberLogin).Error
//	if err != nil {
//		return data, err
//	}
//
//	// 创建角色游戏选项 (MemberGameOption) 默认0
//	var memberGameOption taiwanlogin.MemberGameOption
//	hexStr := "48000000789C63646064F85FCFCC90028408F0BF9E9181112C0382CC50B117CC20F114A038023042210009AC0C9B"
//	hexStr2 := "0x10020000789C636018058319686115D5C62AAA83555417ABA81E56517D06003C02010C"
//	nullStr := ""
//	bytes, err := hex.DecodeString(hexStr[2:]) // remove the "0x" prefix
//	if err != nil {
//		return data, err
//	}
//	bytes2, err := hex.DecodeString(hexStr2[2:]) // remove the "0x" prefix
//	if err != nil {
//		return data, err
//	}
//
//	nullBytes, err := hex.DecodeString(nullStr)
//	if err != nil {
//		return data, err
//	}
//
//	memberGameOption = taiwanlogin.MemberGameOption{
//		MID:              body.UID,
//		Option1:          bytes,
//		Option2:          nullBytes,
//		Option3:          nullBytes,
//		ShortcutEmoticon: bytes2,
//	}
//	err = global.GatewayDBs["taiwan_login"].Create(&memberGameOption).Error
//	if err != nil {
//		return data, err
//	}
//
//	// 创建角色游戏信息 (MemberGameInfo) 默认0
//	var memberPlayInfo taiwanlogin.MemberPlayInfo
//	memberPlayInfo = taiwanlogin.MemberPlayInfo{
//		OccDate:  time.Now(),
//		MID:      body.UID,
//		ServerID: 1,
//	}
//	err = global.GatewayDBs["taiwan_login"].Create(&memberPlayInfo).Error
//	if err != nil {
//		return data, err
//	}
//
//	// 如果有充值 则创建充值记录 点券
//	if body.CashCera > 0 {
//		var memberCash taiwanbilling.CashCera
//		memberCash = taiwanbilling.CashCera{
//			Account: fmt.Sprintf("%d", body.UID),
//			Cera:    body.CashCera,
//			ModTran: 0,
//			ModDate: time.Now(),
//			RegDate: time.Now(),
//		}
//		err = global.GatewayDBs["taiwan_billing"].Create(&memberCash).Error
//		if err != nil {
//			return data, err
//		}
//	}
//
//	// 如果有充值 则创建充值记录 代币
//	if body.CashCeraPoint > 0 {
//		var memberCashPoint taiwanbilling.CashCeraPoint
//		memberCashPoint = taiwanbilling.CashCeraPoint{
//			Account:   fmt.Sprintf("%d", body.UID),
//			CeraPoint: body.CashCera,
//			ModDate:   time.Now(),
//			RegDate:   time.Now(),
//		}
//		err = global.GatewayDBs["taiwan_billing"].Create(&memberCashPoint).Error
//		if err != nil {
//			return data, err
//		}
//	}
//
//	return data, err
//}

// CreateAccount 创建账号
func (a *AccountService) CreateAccount(body CreateAccountRequest) (data dtaiwanmodel.Account, err error) {
	err = global.GatewayDBs["d_taiwan"].Create(&body.Account).Error
	if err != nil {
		return data, err
	}

	fmt.Println("注册的用户UID", body.UID)

	var wg sync.WaitGroup
	errChan := make(chan error, 10) // buffer size to be number of goroutines

	// 创建角色限制 (LimitCreateCharacter) 默认0
	wg.Add(1)
	go func() {
		defer wg.Done()
		var limitCreateCharacter dtaiwanmodel.LimitCreateCharacter
		limitCreateCharacter = dtaiwanmodel.LimitCreateCharacter{MID: body.UID}
		err = global.GatewayDBs["d_taiwan"].Create(&limitCreateCharacter).Error
		if err != nil {
			errChan <- err
		}
	}()

	// 创建角色信息 (CharacterInfo) 默认0
	wg.Add(1)
	go func() {
		defer wg.Done()
		var memberInfo dtaiwanmodel.MemberInfo
		memberInfo = dtaiwanmodel.MemberInfo{MID: body.UID, UserID: fmt.Sprintf("%d", body.UID)}
		err = global.GatewayDBs["d_taiwan"].Create(&memberInfo).Error
		if err != nil {
			errChan <- err
		}
	}()

	// 创建角色加入信息 (MemberJoinInfo) 默认0
	wg.Add(1)
	go func() {
		defer wg.Done()
		var memberJoinInfo dtaiwanmodel.MemberJoinInfo
		memberJoinInfo = dtaiwanmodel.MemberJoinInfo{MID: body.UID}
		err = global.GatewayDBs["d_taiwan"].Create(&memberJoinInfo).Error
		if err != nil {
			errChan <- err
		}
	}()

	// 创建角色白名单 (MemberWhiteAccount) 默认0
	wg.Add(1)
	go func() {
		defer wg.Done()
		var memberWhiteAccount dtaiwanmodel.MemberWhiteAccount
		memberWhiteAccount = dtaiwanmodel.MemberWhiteAccount{MID: body.UID}
		err = global.GatewayDBs["d_taiwan"].Create(&memberWhiteAccount).Error
		if err != nil {
			errChan <- err
		}
	}()

	// 创建角色登录信息 (MemberLogin) 默认0
	wg.Add(1)
	go func() {
		defer wg.Done()
		var memberLogin taiwanlogin.MemberLogin
		memberLogin = taiwanlogin.MemberLogin{MID: body.UID}
		err = global.GatewayDBs["taiwan_login"].Create(&memberLogin).Error
		if err != nil {
			errChan <- err
		}
	}()

	// 创建角色游戏选项 (MemberGameOption) 默认0
	wg.Add(1)
	go func() {
		defer wg.Done()
		var memberGameOption taiwanlogin.MemberGameOption
		hexStr := "48000000789C63646064F85FCFCC90028408F0BF9E9181112C0382CC50B117CC20F114A038023042210009AC0C9B"
		hexStr2 := "0x10020000789C636018058319686115D5C62AAA83555417ABA81E56517D06003C02010C"
		nullStr := ""
		bytes, err := hex.DecodeString(hexStr[2:]) // remove the "0x" prefix
		if err != nil {
			errChan <- err
			return
		}
		bytes2, err := hex.DecodeString(hexStr2[2:]) // remove the "0x" prefix
		if err != nil {
			errChan <- err
			return
		}

		nullBytes, err := hex.DecodeString(nullStr)
		if err != nil {
			errChan <- err
			return
		}

		memberGameOption = taiwanlogin.MemberGameOption{
			MID:              body.UID,
			Option1:          bytes,
			Option2:          nullBytes,
			Option3:          nullBytes,
			ShortcutEmoticon: bytes2,
		}
		err = global.GatewayDBs["taiwan_login"].Create(&memberGameOption).Error
		if err != nil {
			errChan <- err
		}
	}()

	// 创建角色游戏信息 (MemberGameInfo) 默认0
	wg.Add(1)
	go func() {
		defer wg.Done()
		var memberPlayInfo taiwanlogin.MemberPlayInfo
		memberPlayInfo = taiwanlogin.MemberPlayInfo{
			OccDate:  time.Now(),
			MID:      body.UID,
			ServerID: 1,
		}
		err = global.GatewayDBs["taiwan_login"].Create(&memberPlayInfo).Error
		if err != nil {
			errChan <- err
		}
	}()

	// 如果有充值 则创建充值记录 点券
	if body.CashCera > 0 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var memberCash taiwanbilling.CashCera
			memberCash = taiwanbilling.CashCera{
				Account: fmt.Sprintf("%d", body.UID),
				Cera:    body.CashCera,
				ModTran: 0,
				ModDate: time.Now(),
				RegDate: time.Now(),
			}
			err = global.GatewayDBs["taiwan_billing"].Create(&memberCash).Error
			if err != nil {
				errChan <- err
			}
		}()
	}

	// 如果有充值 则创建充值记录 代币
	if body.CashCeraPoint > 0 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var memberCashPoint taiwanbilling.CashCeraPoint
			memberCashPoint = taiwanbilling.CashCeraPoint{
				Account:   fmt.Sprintf("%d", body.UID),
				CeraPoint: body.CashCeraPoint,
				ModDate:   time.Now(),
				RegDate:   time.Now(),
			}
			err = global.GatewayDBs["taiwan_billing"].Create(&memberCashPoint).Error
			if err != nil {
				errChan <- err
			}
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Check if any goroutine returned an error
	close(errChan)
	if len(errChan) > 0 {
		return data, <-errChan
	}

	return data, nil
}
