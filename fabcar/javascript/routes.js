let express = require('express');
let router = express.Router();
let main = require('./controller.js');
let format = require('date-format');

module.exports = router;

router.use(function(req, res, next) {

  console.log(format.asString('hh:mm:ss.SSS', new Date())+'::............ '+req.url+' .............');
  next(); // make sure we go to the next routes and don't stop here

  function afterResponse() {
      res.removeListener('finish', afterResponse);          
  }    
  res.on('finish', afterResponse);

});

router.post('/requestAffiliation*', main.requestAffiliation);
router.get('/getInstituteList',main.getInstituteList)
router.post('/approveAffiliation*', main.approveAffiliation);
router.post('/create_Course*', main.createCourse);
router.post('/add_Batch_class__No*', main.addBatchclassNo);
router.get('/getInstituteList',main.getInstituteList)
router.get('/getCourse',main.getCourse)
router.post('/take_Admission*', main.takeAdmission);
router.get('/get_Student_List',main.getStudent)
router.post('/enroll_Student*', main.enrollStudent);
router.post('/request_Certificates*', main.requestCertificates);
router.get('/get_Request_Certificates',main.getRequestCertificates)
router.post('/issue_Certificate*', main.issueCertificate);
router.get('/get_Certificates',main.getCertificates)
//router.post('/issue_Certificate_Batch*', main.issueCertificateBatch);
router.get('/get_Course_From_Institute',main.getCourseFromInstitute)
router.get('/get_Student_From_Institute',main.getStudentFromInstitute)
router.get('/get_Student_From_Course',main.getStudentFromCourse)
router.post('/edit_Student*', main.editStudent);
router.post('/receive_Certificate*', main.receiveCertificate);
router.post('/issueCertificateCourse*', main.issueCertificateCourse);
router.post('/issueCertificateForStudent*', main.issueCertificateForStudent);
router.get('/getStudentFromCourseBatchno',main.getStudentFromCourseBatchno)
router.get('/view_Student',main.viewStudent)
router.get('/view_Course',main.viewCourse)
router.get('/view_Certificate',main.viewCertificate)
router.get('/get_ClassList',main.getClassList)
router.get('/get_BatchList',main.getBatchList)
router.post('/request_Certificate_Change*', main.requestCertificateChange);
router.get('/get_RequestforCerti_Change',main.getRequestforCertiChange)
router.post('/approve_Certificat_Change*', main.approveCertificateChange);
router.post('/change_Institute_Owner*', main.changeInstituteOwner);
router.get('/getApprovedInstituteList',main.getApprovedInstituteList)
router.post('/editInstitute*', main.editInstitute);
router.delete('/delete',main.delete)
router.get('/queryInstitute',main.queryInstitute)
router.get('/getAllInstituteList',main.getAllInstituteList)
router.get('/getStforAppr',main.getStforAppr)
router.get('/getApprovedStudents',main.getApprovedStudents)
router.get('/Query_certificate_create_details',main.Querycertstu)
router.get('/getStudentIdFromName',main.getStudentIdFromName)
router.get('/getInstituteIdFromName',main.getInstituteIdFromName)




