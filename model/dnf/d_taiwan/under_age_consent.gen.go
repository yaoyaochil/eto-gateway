// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUnderAgeConsent = "under_age_consent"

// UnderAgeConsent mapped from table <under_age_consent>
type UnderAgeConsent struct {
	MID               int32  `gorm:"column:m_id;not null" json:"m_id"`
	ConsentType       int32  `gorm:"column:consent_type;not null" json:"consent_type"`
	LimitMoney        int32  `gorm:"column:limit_money;not null" json:"limit_money"`
	ParentName        string `gorm:"column:parent_name;not null" json:"parent_name"`
	ParentJumin       int64  `gorm:"column:parent_jumin;not null" json:"parent_jumin"`
	ParentPhone1      int32  `gorm:"column:parent_phone1;not null" json:"parent_phone1"`
	ParentPhone2      int32  `gorm:"column:parent_phone2;not null" json:"parent_phone2"`
	ParentPhone3      int32  `gorm:"column:parent_phone3;not null" json:"parent_phone3"`
	ParentEmail       string `gorm:"column:parent_email;not null" json:"parent_email"`
	ParentConsentType int32  `gorm:"column:parent_consent_type;not null" json:"parent_consent_type"`
	NoticeType        int32  `gorm:"column:notice_type;not null" json:"notice_type"`
	NoticeAddr        string `gorm:"column:notice_addr;not null" json:"notice_addr"`
	CreateDate        int32  `gorm:"column:create_date;not null" json:"create_date"`
	ConsentDate       int32  `gorm:"column:consent_date;not null" json:"consent_date"`
	ConsentYn         int32  `gorm:"column:consent_yn;not null" json:"consent_yn"`
	HistoryYn         int32  `gorm:"column:history_yn;not null" json:"history_yn"`
}

// TableName UnderAgeConsent's table name
func (*UnderAgeConsent) TableName() string {
	return TableNameUnderAgeConsent
}
