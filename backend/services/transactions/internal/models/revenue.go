package models

import "time"

type Revenue struct {
	ID          int64
	UserId      int64
	Description string
	Origin      string
	Type        string
	ReceiveDate time.Time
	IsRecieved  bool
}


func (re Revenue) Isvalid() (bool,string) {

	if (re.UserId ==0){
		return false, "repense:validate:UserId required"
	}
	if (re.Origin ==""){
		return false, "repense:validate:Target required"
	}
	
	if(re.Type == "" ){
		return false, "repense:validate:Category required"
	}
	if(re.ReceiveDate.IsZero()){
		return false, "repense:validate:PaymentDate required"
	}

	return true,""
}