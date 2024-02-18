package construction

import (
	"errors"
	"gorm.io/gorm"
	"time"
	"work_diary/httpResponse"
)

var (
	FieldRequiredErr = errors.New("this field is required")
)

type Construction struct {
	gorm.Model
	Work                 string `json:"work" gorm:"type:text"`
	Contractor           string `json:"contractor"`
	ContractualAgreement string `json:"contractual_agreement"`
	contractualTerm
}

type contractualTerm struct {
	InitialDate time.Time `json:"initial_date"`
	FinalDate   time.Time `json:"final_date"`
}

func NewConstruction(work, contractor, contractualAgreement string, initialDate, finalDate time.Time) (*Construction, []httpResponse.Cause) {
	err := validate(work, contractor, contractualAgreement, initialDate, finalDate)
	if err != nil {
		return nil, err
	}
	return &Construction{
		Work:                 work,
		Contractor:           contractor,
		ContractualAgreement: contractualAgreement,
		contractualTerm: contractualTerm{
			InitialDate: initialDate,
			FinalDate:   finalDate,
		},
	}, nil
}

func validate(work, contractor, contractualAgreement string, initialDate, finalDate time.Time) []httpResponse.Cause {
	var causes []httpResponse.Cause
	if work == "" {
		causes = append(causes, httpResponse.Cause{
			Field:   "work",
			Message: FieldRequiredErr.Error(),
		})
	}

	if contractor == "" {
		causes = append(causes, httpResponse.Cause{
			Field:   "contractor",
			Message: FieldRequiredErr.Error(),
		})
	}

	if contractualAgreement == "" {
		causes = append(causes, httpResponse.Cause{
			Field:   "contractual_agreement",
			Message: FieldRequiredErr.Error(),
		})
	}

	if initialDate.IsZero() {
		causes = append(causes, httpResponse.Cause{
			Field:   "initial_date",
			Message: FieldRequiredErr.Error(),
		})
	}

	if finalDate.IsZero() {
		causes = append(causes, httpResponse.Cause{
			Field:   "final_date",
			Message: FieldRequiredErr.Error(),
		})
	}

	if len(causes) > 0 {
		return causes
	}

	return nil
}
