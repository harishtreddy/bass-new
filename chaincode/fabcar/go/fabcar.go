

 package main

 import (
	 "encoding/json"
	 "fmt"

	 "github.com/hyperledger/fabric-contract-api-go/contractapi"
	 "github.com/google/uuid"
 )
 
 // SmartContract provides functions for managing a car
 type SmartContract struct {
	 contractapi.Contract
 }
type QueryResult struct {
		Key    string `json:"Key"`
		Record *Institute
	
}

type QueryCourse struct {
	Key    string `json:"Key"`
	Record *Course

}
type QueryStudent struct {
	Key    string `json:"Key"`
	Record *Student

}
type QueryCerReq struct {
	Key    string `json:"Key"`
	Record *RequestCertificate

}
type QueryCerReqChange struct {
	Key    string `json:"Key"`
	Record *RequestCertificateChange

}
type QueryCertificate struct {
	Key    string `json:"Key"`
	Record *Certificate

}
type InstituteCourse struct{
	Record *Course

}
type InstituteStudent struct{
	Key    string `json:"Key"`
	Record *Student

}
type CourseStudent struct{
	Key    string `json:"Key"`
	Record *Student

}
type Institute struct {
	InstituteID  		string   	`json:"id"`
	InstituteName   	string  	`json:"Institute_name"`
	Address 			string    	`json:"address"`
	ContactNumber   	string		`json:"Contact_Number"`
	Website				string		`json:"website"`
	Email				string      `json:"email"`
	ShortDescription	string      `json:"Short_Description"`
	InstituteOwner 		string  	`json:"Institute_Owner"`
	Status  		string      `json:"status"`
	CourseId   			[]string 	`json:"Course_Id"`
	StudentId			[]string    `json:"Student_Id"`
	ClassNo 			[]string  	`json:"class_No"`
	BatchNo 			[]string  	`json:"batch_No"` 

}

type Course struct {
		Id				string		   `json:"id"`
		CourseName       string   `json:"Course_name"`
		InstituteID  	string   `json:"Institute_ID"`
		InstituteName  	string   `json:"Institute_Name"`
		ClassNo 		[]string  `json:"class_No"`
		BatchNo 		[]string  `json:"batch_No"` 
		StudentId		[]string   `json:"Student_Id"`
}
type Student struct{
	Id				string		   `json:"id"`
	StudentName 	string 			`json:"Student_Name"`
	StudentDOB 		string 			`json:"Student_DOB"`
	Email 			string			`json:"email"`
	ContactNumber   string		`json:"Contact_Number"`
	AlterNo 		string		`json:"Alter_No"`
	Address			string		`json:"address"`
	InstituteID 	string   		`json:"Institute_ID"`
    CourseId 		string 			`json:"Course_Id"`
	BatchNo 		string 			`json:"Batch_No"`
	ClassNo 		string 			`json:"class_No"`
	Status  	string      `json:"status"`
	CertificateId	string     `json:"Certificate_Id"`
}
type Certificate struct{
	Id				string		   `json:"id"`
	InstituteID 	string   		`json:"Institute_ID"`
	InstituteName 	string   		`json:"Institute_Name"`
	CourseId 		string 			`json:"Course_Id"`
	CourseName 		string 			`json:"Course_Name"`
	StudentId		string			`json:"Student_Id"`
	StudentName 	string 			`json:"Student_Name"`
    StudentDOB 		string 			`json:"Student_DOB"`
	BatchNo 		string 			`json:"Batch_No"`
	ClassNo 		string 			`json:"class_No"`
	Status 			string 			`json:"status"`

}
type QueryAllClass struct {
	ClassNo []string	   `json:"Class_No"`
}
type QueryAllBatch struct {
	BatchNo 		[]string  `json:"batch_No"` 
	}
type RequestCertificate struct{
	Id				string		   `json:"id"`
	StudentId 		string          `json:"Student_Id"`
	Status 			string			`json:"status"`
}
type RequestCertificateChange struct{
	Id				string		   `json:"id"`
	CertificateId 		string      `json:"Certificate_Id"`
	StudentName 	string 			`json:"Student_Name"`
    StudentDOB 		string 			`json:"Student_DOB"`
	Status 			string			`json:"status"`
}
func GetUId() (string, error) {
	id, err := uuid.NewUUID()
    if err != nil {
        return "", err
    }
    return id.String(), err
 }
 // InitLedger adds a base set of cars to the ledger
func (s *SmartContract) Init(ctx contractapi.TransactionContextInterface) error {	
	fmt.Printf("Hello\n")
	return nil
 }
func (s *SmartContract) RequestAffiliation(ctx contractapi.TransactionContextInterface,instituteName string,address string,contactNumber string,website string,email string,shortDescription string,instituteOwner string) error {


	fmt.Printf("Adding Institute to the ledger ...\n")
	// if len(args) != 8 {
	// 	return fmt.Errorf("InvalidArgumentError: Incorrect number of arguments. Expecting 8")
    // }

    //Prepare key for the new Org
	uid, err := GetUId()
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	id := "Institute-" + uid
	fmt.Printf("Validating Institute data\n")
	//Validate the Org data
	var institute = Institute{InstituteID: id,			   
				InstituteName: instituteName,
				Address:address,
				ContactNumber:contactNumber,
				Website:website,
				Email:email,
				ShortDescription:shortDescription,
				InstituteOwner:instituteOwner,
					Status: "PENDING",
					CourseId: make([]string, 0),
					StudentId:make([]string, 0),
					ClassNo:make([]string, 0),
					BatchNo:make([]string, 0),

					}

	//Encrypt and Marshal Org data in order to put in world state
	instituteAsBytes, _ := json.Marshal(institute)

	return ctx.GetStub().PutState(id, instituteAsBytes)
}
func (s *SmartContract) ApproveAffiliation(ctx contractapi.TransactionContextInterface,instituteID string) error {
	fmt.Printf("Changing Status in the ledger ...\n")

	instituteAsBytes, err := ctx.GetStub().GetState(instituteID)
	var institute = Institute{};
	json.Unmarshal(instituteAsBytes, &institute);
	institute.Status = "APPROVED";

	//Encrypt and Marshal Org data in order to put in world state
	fmt.Printf("Marshalling doctor data\n")
	instituteAsBytes, err = json.Marshal(institute)
	if err != nil {
		return fmt.Errorf("MarshallingError: %s", err)
	}
	
	//Add the Org to the ledger world state
	err = ctx.GetStub().PutState(instituteID, instituteAsBytes)
	if err != nil {
		return fmt.Errorf("LegderCommitError: %s", err)
	}

	  return nil
}
func (s *SmartContract) GetInstituteList(ctx contractapi.TransactionContextInterface) ([]QueryResult, error) {
	query := "{\"selector\": {\"_id\": {\"$regex\": \"^Institute-\"} } }"
	resultsIterator, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	results := []QueryResult{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		institute := new(Institute)
		_ = json.Unmarshal(queryResponse.Value, institute)
		if institute.Status == "PENDING"{

		queryResult := QueryResult{Key: queryResponse.Key, Record: institute}
		results = append(results, queryResult)
	}
	}
	return results, nil
}
func (s *SmartContract) GetApprovedInstituteList(ctx contractapi.TransactionContextInterface) ([]QueryResult, error) {
	query := "{\"selector\": {\"_id\": {\"$regex\": \"^Institute-\"} } }"
	resultsIterator, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	results := []QueryResult{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		institute := new(Institute)
		_ = json.Unmarshal(queryResponse.Value, institute)
		if institute.Status == "APPROVED"{

		queryResult := QueryResult{Key: queryResponse.Key, Record: institute}
		results = append(results, queryResult)
	}
	}
	return results, nil
}
func (s *SmartContract) GetAllInstituteList(ctx contractapi.TransactionContextInterface) ([]QueryResult, error) {
	query := "{\"selector\": {\"_id\": {\"$regex\": \"^Institute-\"} } }"
	resultsIterator, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	results := []QueryResult{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		institute := new(Institute)
		_ = json.Unmarshal(queryResponse.Value, institute)

		queryResult := QueryResult{Key: queryResponse.Key, Record: institute}
		results = append(results, queryResult)
	
	}
	return results, nil
}
func (s *SmartContract) QueryInstitute(ctx contractapi.TransactionContextInterface,instituteID string) (*Institute, error) {
	instituteAsBytes, err := ctx.GetStub().GetState(instituteID)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if instituteAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", instituteID)
	}

	institute := new(Institute)
	_ = json.Unmarshal(instituteAsBytes, institute)

	return institute, nil
}
func (s *SmartContract) CreateCourse(ctx contractapi.TransactionContextInterface,instituteID string,instituteName string,coursename string,) error {

		uid, err := GetUId()
		if err != nil {
			return fmt.Errorf("%s", err)
		}
		id := "Course-" + uid
	
		var course = Course{ Id: id,	
							CourseName: coursename,
							InstituteID: instituteID,
							InstituteName:instituteName,
							 ClassNo: make([]string, 0),
							 BatchNo: make([]string, 0),
							 StudentId: make([]string, 0),
							 }
		
		//add Prescription id in the doctor's patient ids list
		instituteAsBytes, _ := ctx.GetStub().GetState(instituteID)
	//	ptPatientAsBytes, _ := Decrypt(patientAsBytes, key, IV)
		institute := Institute{}
		err = json.Unmarshal(instituteAsBytes, &institute)
		if err != nil {
				return fmt.Errorf("%s", err)
		}
		institute.CourseId = append(institute.CourseId, id)
		instituteJSONAsBytes, _ := json.Marshal(institute)
		//ctPatientAsBytes, _ := Encrypt(patientJSONAsBytes, key, IV)
		ctx.GetStub().PutState(instituteID, instituteJSONAsBytes)
	
		//Marshal and Encrypt Prescription data in order to put in world state
		fmt.Printf("Marshalling Prescription data\n")
		courseAsBytes, err := json.Marshal(course)
		if err != nil {
			return fmt.Errorf("MarshallingError: %s", err)
		}
		
	
		//Add Patient to the ledger world state
		err = ctx.GetStub().PutState(id, courseAsBytes)
		if err != nil {
			return fmt.Errorf("LegderCommitError: %s", err)
		}
	
		fmt.Printf("Added Prescription Successfully\n")
		return nil
}

func (s *SmartContract) AddBatchclassNo(ctx contractapi.TransactionContextInterface,courseID string,classNo string,batchNo string) error {
	fmt.Printf("Changing Status in the ledger ...\n")

	courseAsBytes, err := ctx.GetStub().GetState(courseID)
	var course = Course{};
	json.Unmarshal(courseAsBytes, &course);
	course.ClassNo = append(course.ClassNo, classNo)
	course.BatchNo = append(course.BatchNo, batchNo)
	//################################################################
	instituteAsBytes, _ := ctx.GetStub().GetState(course.InstituteID)
	institute := Institute{}
	err = json.Unmarshal(instituteAsBytes, &institute)
	if err != nil {
			return fmt.Errorf("%s", err)
	}
	institute.ClassNo = append(course.ClassNo, classNo)
	institute.BatchNo = append(course.BatchNo, batchNo)

	instituteJSONAsBytes, _ := json.Marshal(institute)
	ctx.GetStub().PutState(course.InstituteID, instituteJSONAsBytes)
	//################################################################
	courseAsBytes, err = json.Marshal(course)
	if err != nil {
		return fmt.Errorf("MarshallingError: %s", err)
	}
	
	err = ctx.GetStub().PutState(courseID, courseAsBytes)
	if err != nil {
		return fmt.Errorf("LegderCommitError: %s", err)
	}

	  return nil
}

func (s *SmartContract) GetCourse(ctx contractapi.TransactionContextInterface) ([]QueryCourse, error) {
	query := "{\"selector\": {\"_id\": {\"$regex\": \"^Course-\"} } }"
	resultsIterator, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	result := []QueryCourse{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		course := new(Course)
		_ = json.Unmarshal(queryResponse.Value, course)
		queryResult := QueryCourse{Key: queryResponse.Key, Record: course}
		result = append(result, queryResult)
	}
	return result, nil
}
func (s *SmartContract) TakeAdmission(ctx contractapi.TransactionContextInterface,studentName string,studentDOB string,email string,contactNumber string,alterNo string,address string,instituteID string,courseId string) error {

	uid, err := GetUId()
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	instituteAsBytes, _ := ctx.GetStub().GetState(instituteID)
	institute := Institute{}
		err = json.Unmarshal(instituteAsBytes, &institute)
		if err != nil {
				return fmt.Errorf("%s", err)
		}
		//Institute.CertificateId = id
		// var batchNo=institute.BatchNo 
		// var classNo=institute.ClassNo 

	id := "Student-" + uid
	//Validate the Org data
	var student = Student{Id: id,
		StudentName: studentName,
		StudentDOB:studentDOB,
		Email:email,
		ContactNumber:contactNumber,
		AlterNo:alterNo,
		Address:address,
		InstituteID: instituteID,
		CourseId: courseId,
		BatchNo:"Nill",
		ClassNo:"Nill",
		Status:"PENDING",
		CertificateId:"Not Issued",
	}

	//Encrypt and Marshal Org data in order to put in world state
	studentAsBytes, _ := json.Marshal(student)

	return ctx.GetStub().PutState(id, studentAsBytes)
}
func (s *SmartContract) GetStudent(ctx contractapi.TransactionContextInterface) ([]QueryStudent, error) {
	query := "{\"selector\": {\"_id\": {\"$regex\": \"^Student-\"} } }"
	resultsIterator, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	result := []QueryStudent{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		student := new(Student)
		_ = json.Unmarshal(queryResponse.Value, student)
		queryResult := QueryStudent{Key: queryResponse.Key, Record: student}
		result = append(result, queryResult)
	}
	return result, nil
}
func (s *SmartContract) GetStforAppr(ctx contractapi.TransactionContextInterface) ([]QueryStudent, error) {
	query := "{\"selector\": {\"_id\": {\"$regex\": \"^Student-\"} } }"
	resultsIterator, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	result := []QueryStudent{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		student := new(Student)
		_ = json.Unmarshal(queryResponse.Value, student)
		if student.Status == "PENDING"{

		queryResult := QueryStudent{Key: queryResponse.Key, Record: student}
		result = append(result, queryResult)
	}
}
	return result, nil
}
func (s *SmartContract) GetApprovedStudents(ctx contractapi.TransactionContextInterface) ([]QueryStudent, error) {
	query := "{\"selector\": {\"_id\": {\"$regex\": \"^Student-\"} } }"
	resultsIterator, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	result := []QueryStudent{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		student := new(Student)
		_ = json.Unmarshal(queryResponse.Value, student)
		if student.Status == "APPROVED"{

		queryResult := QueryStudent{Key: queryResponse.Key, Record: student}
		result = append(result, queryResult)
	}
}
	return result, nil
}
func (s *SmartContract) EnrollStudent(ctx contractapi.TransactionContextInterface,studentID string,batchNo string,classNo string) error {
	//
	uid, err := GetUId()
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	id := "Certificate-" + uid

	studentAsBytes, err := ctx.GetStub().GetState(studentID)
	var student = Student{};
	json.Unmarshal(studentAsBytes, &student);
	var instituteID =student.InstituteID 
	var studentName=student.StudentName
	var studentDOB=student.StudentDOB
	var courseId=student.CourseId
	student.CertificateId =id
	student.BatchNo =batchNo
	student.ClassNo = classNo
	student.Status ="APPROVED"

	//##################institute########################################
	instituteAsBytes, _ := ctx.GetStub().GetState(instituteID)
		institute := Institute{}
		err = json.Unmarshal(instituteAsBytes, &institute)
		if err != nil {
				return fmt.Errorf("%s", err)
		}
		institute.StudentId = append(institute.StudentId, studentID)
		institute.ClassNo = append(institute.ClassNo, classNo)
		institute.BatchNo = append(institute.BatchNo, batchNo)
		var instituteName= institute.InstituteName
		instituteJSONAsBytes, _ := json.Marshal(institute)
		ctx.GetStub().PutState(instituteID, instituteJSONAsBytes)
	//##################institute########################################
	
	//##################course########################################
	courseAsBytes, _ := ctx.GetStub().GetState(courseId)
		course := Course{}
		err = json.Unmarshal(courseAsBytes, &course)
		if err != nil {
				return fmt.Errorf("%s", err)
		}
		course.StudentId = append(course.StudentId, studentID)
		course.ClassNo = append(course.ClassNo, classNo)
		course.BatchNo = append(course.BatchNo, batchNo)
		var courseName=course.CourseName 
		courseJSONAsBytes, _ := json.Marshal(course)
		ctx.GetStub().PutState(courseId, courseJSONAsBytes)
		//##################course########################################

	//##################certificate########################################
	var certificate = Certificate{Id: id,			   
		InstituteID: instituteID,
		InstituteName:instituteName,
		CourseId: courseId,
		CourseName:courseName,
		StudentId:studentID,
		StudentName:studentName,
		StudentDOB:studentDOB,
		BatchNo:batchNo,
		ClassNo:classNo,
		Status:"Not applied",
		}
	certificateAsBytes, _ := json.Marshal(certificate)

	ctx.GetStub().PutState(id, certificateAsBytes)
	//##################certificate########################################
	
	studentAsBytes, err = json.Marshal(student)
	if err != nil {
		return fmt.Errorf("MarshallingError: %s", err)
	}
	
	//Add the Org to the ledger world state
	err = ctx.GetStub().PutState(studentID, studentAsBytes)
	if err != nil {
		return fmt.Errorf("LegderCommitError: %s", err)
	}

	  return nil
}
func (s *SmartContract) RequestCertificates(ctx contractapi.TransactionContextInterface,studentId string) error {



	uid, err := GetUId()
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	id := "RequestCertificate-" + uid
	var requestCertificate = RequestCertificate{Id: id,			   
		StudentId: studentId,
		Status:"APPLYED",
					}

	//Encrypt and Marshal Org data in order to put in world state
	requestCertificateAsBytes, _ := json.Marshal(requestCertificate)
	//#################certificate#########################################################################
	//##################institute########################################
	studentAsBytes, _ := ctx.GetStub().GetState(studentId)
		student := Student{}
		err = json.Unmarshal(studentAsBytes, &student)
		if err != nil {
				return fmt.Errorf("%s", err)
		}
		var certificateId=student.CertificateId 	
		
	//##################institute########################################
	certificateAsBytes, _ := ctx.GetStub().GetState(certificateId)
	certificate := Certificate{}
		err = json.Unmarshal(certificateAsBytes, &certificate)
		if err != nil {
				return fmt.Errorf("%s", err)
		}
		certificate.Status = "PENDING"

		certificateJSONAsBytes, _ := json.Marshal(certificate)
		ctx.GetStub().PutState(certificateId, certificateJSONAsBytes)
	//##################institute########################################
	 //	studentJSONAsBytes, _ := json.Marshal(student)
	//	ctx.GetStub().PutState(studentId, studentJSONAsBytes)
	//##################institute########################################
	//################certificate######################################################################
	return ctx.GetStub().PutState(id, requestCertificateAsBytes)
}
func (s *SmartContract) GetRequestCertificates(ctx contractapi.TransactionContextInterface) ([]QueryCerReq, error) {
	query := "{\"selector\": {\"_id\": {\"$regex\": \"^RequestCertificate-\"} } }"
	resultsIterator, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	result := []QueryCerReq{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		requestCertificate := new(RequestCertificate)
		_ = json.Unmarshal(queryResponse.Value, requestCertificate)
		queryResult := QueryCerReq{Key: queryResponse.Key, Record: requestCertificate}
		result = append(result, queryResult)
	}
	return result, nil
}
func (s *SmartContract) IssueCertificate(ctx contractapi.TransactionContextInterface,requestId string) error {
	
	requestAsBytes, err := ctx.GetStub().GetState(requestId)
	var requestCertificate = RequestCertificate{};
	json.Unmarshal(requestAsBytes, &requestCertificate);
	requestCertificate.Status = "ISSUED"
	var studentId=requestCertificate.StudentId
	//##################institute########################################
	studentAsBytes, _ := ctx.GetStub().GetState(studentId)
		student := Student{}
		err = json.Unmarshal(studentAsBytes, &student)
		if err != nil {
				return fmt.Errorf("%s", err)
		}
	//	student.CertificateId = id
		var certificateId=student.CertificateId 	
		// var courseId=student.CourseId 		
		// var studentName=student.StudentName 	
		// var studentDOB=student.StudentDOB 		
		// var batchNo=student.BatchNo 		
		// var classNo=student.ClassNo 		
		// studentJSONAsBytes, _ := json.Marshal(student)
		// ctx.GetStub().PutState(studentId, studentJSONAsBytes)
	//##################institute########################################
	certificateAsBytes, _ := ctx.GetStub().GetState(certificateId)
	certificate := Certificate{}
		err = json.Unmarshal(certificateAsBytes, &certificate)
		if err != nil {
				return fmt.Errorf("%s", err)
		}
		certificate.Status = "APPROVED"

		certificateJSONAsBytes, _ := json.Marshal(certificate)
		ctx.GetStub().PutState(certificateId, certificateJSONAsBytes)

	//##################################################

	requestAsBytes, err = json.Marshal(requestCertificate)
	if err != nil {
		return fmt.Errorf("MarshallingError: %s", err)
	}
	
	//Add the Org to the ledger world state
	err = ctx.GetStub().PutState(requestId, requestAsBytes)
	if err != nil {
		return fmt.Errorf("LegderCommitError: %s", err)
	}

	  return nil
}
func (s *SmartContract) ReceiveCertificate(ctx contractapi.TransactionContextInterface,certificateId string) error {
	fmt.Printf("Changing Status in the ledger ...\n")

	certificateAsBytes, err := ctx.GetStub().GetState(certificateId)
	var certificate = Certificate{};
	json.Unmarshal(certificateAsBytes, &certificate);
	certificate.Status = "RECEIVED"

	certificateAsBytes, err = json.Marshal(certificate)
	if err != nil {
		return fmt.Errorf("MarshallingError: %s", err)
	}
	
	err = ctx.GetStub().PutState(certificateId, certificateAsBytes)
	if err != nil {
		return fmt.Errorf("LegderCommitError: %s", err)
	}

	  return nil
}
func (s *SmartContract) GetCertificates(ctx contractapi.TransactionContextInterface) ([]QueryCertificate, error) {
	query := "{\"selector\": {\"_id\": {\"$regex\": \"^Certificate-\"} } }"
	resultsIterator, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	result := []QueryCertificate{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		certificate := new(Certificate)
		_ = json.Unmarshal(queryResponse.Value, certificate)
		queryResult := QueryCertificate{Key: queryResponse.Key, Record: certificate}
		result = append(result, queryResult)
	}
	return result, nil
}
func (s *SmartContract) GetCourseFromInstitute(ctx contractapi.TransactionContextInterface,InstituteId string) ([]InstituteCourse, error) {
	query := "{\"selector\": {\"_id\": {\"$regex\": \"^Course-\"} } }"
	resultsIterator, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	result := []InstituteCourse{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		course := new(Course)
		_ = json.Unmarshal(queryResponse.Value,course)

		if course.InstituteID == InstituteId{
		queryResult := InstituteCourse{Record: course}
		result = append(result, queryResult)
	}
	}
	return result, nil
}

func (s *SmartContract) GetStudentFromInstitute(ctx contractapi.TransactionContextInterface,instituteID string) ([]InstituteStudent, error) {
	query := "{\"selector\": {\"_id\": {\"$regex\": \"^Student-\"} } }"
	resultsIterator, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	result := []InstituteStudent{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		student := new(Student)
		_ = json.Unmarshal(queryResponse.Value, student)
		if student.InstituteID == instituteID{

		queryResult := InstituteStudent{Key: queryResponse.Key, Record: student}
		result = append(result, queryResult)
	}
	}
	return result, nil
}
func (s *SmartContract) GetStudentFromCourse(ctx contractapi.TransactionContextInterface,courseId string) ([]CourseStudent, error) {
	query := "{\"selector\": {\"_id\": {\"$regex\": \"^Student-\"} } }"
	resultsIterator, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	result := []CourseStudent{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		student := new(Student)
		_ = json.Unmarshal(queryResponse.Value, student)
		if student.CourseId == courseId{

		queryResult := CourseStudent{Key: queryResponse.Key, Record: student}
		result = append(result, queryResult)
	}
	}
	return result, nil
}
func (s *SmartContract) EditStudent(ctx contractapi.TransactionContextInterface,studentID string,studentName string,studentDOB string,email string,contactNumber string,alterNo string,address string) error {

	studentAsBytes, err := ctx.GetStub().GetState(studentID)
	var student = Student{};
	json.Unmarshal(studentAsBytes, &student);
	student.StudentName = studentName
	student.StudentDOB = studentDOB
	student.Email=email
	student.ContactNumber=contactNumber
	student.AlterNo=alterNo
	student.Address=address
	studentAsBytes, err = json.Marshal(student)
	if err != nil {
		return fmt.Errorf("MarshallingError: %s", err)
	}
	
	err = ctx.GetStub().PutState(studentID, studentAsBytes)
	if err != nil {
		return fmt.Errorf("LegderCommitError: %s", err)
	}

	  return nil
}
func (s *SmartContract) IssueCertificateCourse(ctx contractapi.TransactionContextInterface, courseId string) error {
	uid, err := GetUId()
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	id := "Certificate-" + uid
	courseIdAsBytes, err := ctx.GetStub().GetState(courseId)
	var course = Course{};
	json.Unmarshal(courseIdAsBytes, &course);
	for i:=0;i<len(course.StudentId)+1;i++{
	//##################institute########################################
	studentAsBytes, _ := ctx.GetStub().GetState(course.StudentId[i])
	student := Student{}
	err = json.Unmarshal(studentAsBytes, &student)
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	student.CertificateId = id
	var instituteID=student.InstituteID 	
	var courseId=student.CourseId 		
	var studentName=student.StudentName 	
	var studentDOB=student.StudentDOB 		
	var batchNo=student.BatchNo 		
	var classNo=student.ClassNo 		
	studentJSONAsBytes, _ := json.Marshal(student)
	ctx.GetStub().PutState(course.StudentId[i], studentJSONAsBytes)
	//##################institute########################################
	var certificate = Certificate{Id: id,			   
		InstituteID: instituteID,
		CourseId: courseId,
		StudentName:studentName,
		StudentDOB:studentDOB,
		BatchNo:batchNo,
		ClassNo:classNo,
		Status:"ISSUED",
		}
	certificateAsBytes, _ := json.Marshal(certificate)

	ctx.GetStub().PutState(id, certificateAsBytes)

	}
 return nil
}
func (s *SmartContract) IssueCertificateForStudent(ctx contractapi.TransactionContextInterface, studentId string) error {
	// uid, err := GetUId()
	// if err != nil {
	// 	return fmt.Errorf("%s", err)
	// }
	// id := "Certificate-" + uid
	studentAsBytes, _ := ctx.GetStub().GetState(studentId)
	student := Student{}
	err := json.Unmarshal(studentAsBytes, &student)
	if err != nil {
			return fmt.Errorf("%s", err)
	}
//	student.CertificateId = id
	var certificateId=student.CertificateId 
	//######################certiicate####################################
	certificateAsBytes, _ := ctx.GetStub().GetState(certificateId)
	certificate := Certificate{}
		err = json.Unmarshal(certificateAsBytes, &certificate)
		if err != nil {
				return fmt.Errorf("%s", err)
		}
		certificate.Status = "APPROVED"

		certificateJSONAsBytes, _ := json.Marshal(certificate)
		return ctx.GetStub().PutState(certificateId, certificateJSONAsBytes)
	//#######################certificate################################	
}
	//##################institute########################################
func (s *SmartContract) GetStudentFromCourseBatchno(ctx contractapi.TransactionContextInterface,courseId string,batchNo string) ([]CourseStudent, error) {
	query := "{\"selector\": {\"_id\": {\"$regex\": \"^Student-\"} } }"
	resultsIterator, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	result := []CourseStudent{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		student := new(Student)
		_ = json.Unmarshal(queryResponse.Value, student)
		if student.CourseId == courseId && student.BatchNo==batchNo && student.CertificateId=="Not Issued"{

		queryResult := CourseStudent{Key: queryResponse.Key, Record: student}
		result = append(result, queryResult)
	}
	}
	return result, nil
}
func (s *SmartContract) QueryStudent(ctx contractapi.TransactionContextInterface,studentId string) (*Student, error) {
	studentAsBytes, err := ctx.GetStub().GetState(studentId)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if studentAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", studentId)
	}

	student := new(Student)
	_ = json.Unmarshal(studentAsBytes, student)

	return student, nil
}
func (s *SmartContract) Querycertstu(ctx contractapi.TransactionContextInterface,studentId string) (*Certificate, error) {
	studentAsBytes, err := ctx.GetStub().GetState(studentId)
		student := Student{}
		err = json.Unmarshal(studentAsBytes, &student)
		
		var certificateId=student.CertificateId 	
			
		certificateAsBytes, err := ctx.GetStub().GetState(certificateId)

		if err != nil {
			return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
		}
	
		if certificateAsBytes == nil {
			return nil, fmt.Errorf("%s does not exist", certificateId)
		}
	
		certificate := new(Certificate)
		_ = json.Unmarshal(certificateAsBytes, certificate)
	
		return certificate, nil
	}
func (s *SmartContract) QueryCourse(ctx contractapi.TransactionContextInterface,courseId string) (*Course, error) {
	courseAsBytes, err := ctx.GetStub().GetState(courseId)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if courseAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", courseId)
	}

	course := new(Course)
	_ = json.Unmarshal(courseAsBytes, course)

	return course, nil
}
func (s *SmartContract) QueryCertificate(ctx contractapi.TransactionContextInterface,certificateId string) (*Certificate, error) {
	certificateAsBytes, err := ctx.GetStub().GetState(certificateId)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if certificateAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", certificateId)
	}

	certificate := new(Certificate)
	_ = json.Unmarshal(certificateAsBytes, certificate)

	return certificate, nil
}
// func (s *SmartContract) GetClassList(ctx contractapi.TransactionContextInterface,instituteID string) ([]QueryAllClass, error) {
// 	query := "{\"selector\": {\"_id\": {\"$regex\": \"^Institute-\"} } }"
// 	resultsIterator, err := ctx.GetStub().GetQueryResult(query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resultsIterator.Close()
// 	results := []QueryAllClass{}
// 	for resultsIterator.HasNext() {
// 		queryResponse, err := resultsIterator.Next()
// 		if err != nil {
// 			return nil, err
// 		}
// 		institute := new(Institute)
// 		_ = json.Unmarshal(queryResponse.Value, institute)
// 		if institute.InstituteID == instituteID{

// 		queryResult := QueryAllClass{ClassNo:institute.ClassNo}
// 		results = append(results, queryResult)
// 	}
// 	}

// 	return results, nil
// }
// func (s *SmartContract) GetBatchList(ctx contractapi.TransactionContextInterface,instituteID string) ([]QueryAllBatch, error) {
// 	query := "{\"selector\": {\"_id\": {\"$regex\": \"^Institute-\"} } }"
// 	resultsIterator, err := ctx.GetStub().GetQueryResult(query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resultsIterator.Close()
// 	results := []QueryAllBatch{}
// 	for resultsIterator.HasNext() {
// 		queryResponse, err := resultsIterator.Next()
// 		if err != nil {
// 			return nil, err
// 		}
// 		institute := new(Institute)
// 		_ = json.Unmarshal(queryResponse.Value, institute)
// 		if institute.InstituteID == instituteID{

// 		queryResult := QueryAllBatch{BatchNo:institute.BatchNo}
// 		results = append(results, queryResult)
// 	}
// 	}

// 	return results, nil
// }
func (s *SmartContract) RequestCertificateChange(ctx contractapi.TransactionContextInterface,certificateId string,studentName string,studentDOB string) error {



	uid, err := GetUId()
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	id := "RequestCertificateChange-" + uid
	var requestCertificate = RequestCertificateChange{Id: id,			   
		CertificateId: certificateId,
		StudentName:studentName,
		StudentDOB:studentDOB,
		Status:"APPLYED",
					}

	//Encrypt and Marshal Org data in order to put in world state
	requestCertificateAsBytes, _ := json.Marshal(requestCertificate)

	return ctx.GetStub().PutState(id, requestCertificateAsBytes)
}
func (s *SmartContract) GetRequestforCertiChange(ctx contractapi.TransactionContextInterface) ([]QueryCerReqChange, error) {
	query := "{\"selector\": {\"_id\": {\"$regex\": \"^RequestCertificateChange-\"} } }"
	resultsIterator, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	result := []QueryCerReqChange{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		requestCertificateChange := new(RequestCertificateChange)
		_ = json.Unmarshal(queryResponse.Value, requestCertificateChange)
		queryResult := QueryCerReqChange{Key: queryResponse.Key, Record: requestCertificateChange}
		result = append(result, queryResult)
	}
	return result, nil
}
func (s *SmartContract) ApproveCertificateChange(ctx contractapi.TransactionContextInterface,requestId string) error {
	requestAsBytes, err := ctx.GetStub().GetState(requestId)
	var requestCertificateChange = RequestCertificateChange{};
	json.Unmarshal(requestAsBytes, &requestCertificateChange);
	requestCertificateChange.Status = "ISSUED"
	var certificateId=requestCertificateChange.CertificateId
	var studentName=requestCertificateChange.StudentName
	var studentDOB=requestCertificateChange.StudentDOB

	//##################institute########################################
	certificateAsBytes, _ := ctx.GetStub().GetState(certificateId)
	certificate := Certificate{}
		err = json.Unmarshal(certificateAsBytes, &certificate)
		if err != nil {
				return fmt.Errorf("%s", err)
		}
		certificate.StudentName = studentName
		certificate.StudentDOB = studentDOB

		certificateJSONAsBytes, _ := json.Marshal(certificate)
		ctx.GetStub().PutState(certificateId, certificateJSONAsBytes)
	//##################institute########################################

	requestAsBytes, err = json.Marshal(requestCertificateChange)
	if err != nil {
		return fmt.Errorf("MarshallingError: %s", err)
	}
	
	//Add the Org to the ledger world state
	err = ctx.GetStub().PutState(requestId, requestAsBytes)
	if err != nil {
		return fmt.Errorf("LegderCommitError: %s", err)
	}

	  return nil
}

func (s *SmartContract) EditInstitute(ctx contractapi.TransactionContextInterface,instituteID string,instituteName string,address string,contactNumber string,website string,email string,shortDescription string) error {

	instituteAsBytes, err := ctx.GetStub().GetState(instituteID)
	var institute = Institute{};
	json.Unmarshal(instituteAsBytes, &institute);
	institute.InstituteName = instituteName
	institute.Address = address
	institute.ContactNumber = contactNumber
	institute.Website = website
	institute.Email = email

	institute.ShortDescription = shortDescription

	
	instituteAsBytes, err = json.Marshal(institute)
	if err != nil {
		return fmt.Errorf("MarshallingError: %s", err)
	}
	
	err = ctx.GetStub().PutState(instituteID, instituteAsBytes)
	if err != nil {
		return fmt.Errorf("LegderCommitError: %s", err)
	}

	  return nil
}

func (s *SmartContract) Delete(ctx contractapi.TransactionContextInterface, Id string) error {
	err := ctx.GetStub().DelState(Id)
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	return nil
}
func (s *SmartContract) ChangeInstituteOwner(ctx contractapi.TransactionContextInterface,instituteID string,instituteOwner string) error {
	fmt.Printf("Changing Status in the ledger ...\n")

	instituteAsBytes, err := ctx.GetStub().GetState(instituteID)
	var institute = Institute{};
	json.Unmarshal(instituteAsBytes, &institute);
	institute.InstituteOwner = instituteOwner;

	//Encrypt and Marshal Org data in order to put in world state
	fmt.Printf("Marshalling doctor data\n")
	instituteAsBytes, err = json.Marshal(institute)
	if err != nil {
		return fmt.Errorf("MarshallingError: %s", err)
	}
	
	//Add the Org to the ledger world state
	err = ctx.GetStub().PutState(instituteID, instituteAsBytes)
	if err != nil {
		return fmt.Errorf("LegderCommitError: %s", err)
	}

	  return nil
}
func (s *SmartContract) GetStudentIdFromName(ctx contractapi.TransactionContextInterface, studentName string) (string, error) {
	query := "{\"selector\": {\"_id\": {\"$regex\": \"^Student-\"} } }"
	resultsIterator, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return "", err
	}
	defer resultsIterator.Close()
	result:=""
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return "", err
		}
		student := new(Student)
		_ = json.Unmarshal(queryResponse.Value, student)
		if(student.StudentName==studentName){
			result=student.Id
			break;
		}		
	}
	if(result==""){
		return "student name not found",nil
	}
	return result, nil
}

func (s *SmartContract) GetInstituteIdFromName(ctx contractapi.TransactionContextInterface, instituteName string) (string, error) {
	query := "{\"selector\": {\"_id\": {\"$regex\": \"^Institute-\"} } }"
	resultsIterator, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return "", err
	}
	defer resultsIterator.Close()
	
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return "", err
		}
		institute := new(Institute)
		_ = json.Unmarshal(queryResponse.Value, institute)
		if(institute.InstituteName==instituteName){
			return institute.InstituteID,nil
		}		
	}
	
	return "institute not found", nil
}


func main() {
 
	 chaincode, err := contractapi.NewChaincode(new(SmartContract))
 
	 if err != nil {
		 fmt.Printf("Error create  chaincode: %s", err.Error())
		 return
	 }
 
	 if err := chaincode.Start(); err != nil {
		 fmt.Printf("Error starting  chaincode: %s", err.Error())
	 }
 }
 