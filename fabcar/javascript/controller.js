/*
 * SPDX-License-Identifier: Apache-2.0
 */

'use strict';
const { Gateway, Wallets } = require('fabric-network');
const FabricCAServices = require('fabric-ca-client');
const fs = require('fs');
const path = require('path');

const ccpPath = path.resolve(__dirname, '..', '..', 'first-network', 'connection-org1.json');

exports.requestAffiliation = async function (req, res, next) {
    var instituteName = req.body.institute_Name;
    var address = req.body.address
    var contactNumber = req.body.contactNumber
    var website = req.body.website
    var email = req.body.email
    var shortDescription = req.body.short_Description
    var instituteOwner = req.body.institute_Owner

    var args = [instituteName,address,contactNumber,website,email,shortDescription,instituteOwner];

    Invoke("RequestAffiliation", args, res);
}
exports.getInstituteList = async function (req, res, next) {
    QueryAll("GetInstituteList",res)
}
exports.getApprovedInstituteList = async function (req, res, next) {
    QueryAll("GetApprovedInstituteList",res)
}
exports.getAllInstituteList = async function (req, res, next) {
    QueryAll("GetAllInstituteList",res)
}
exports.approveAffiliation = async function (req, res, next) {
    var instituteID = req.body.institute_ID;
    
    var args = [instituteID];

    Invoke("ApproveAffiliation", args, res);
}
exports.createCourse = async function (req, res, next) {
    var InstituteID = req.body.institute_ID;
    var InstituteName = req.body.institute_Name;
    var Coursename = req.body.course_name;

    
    var args = [InstituteID,InstituteName,Coursename];

    Invoke("CreateCourse", args, res);
}
exports.addBatchclassNo = async function (req, res, next) {
    var CourseID = req.body.course_ID;
    var ClassNo = req.body.class_No;
    var BatchNo = req.body.batch_No;

    
    var args = [CourseID,ClassNo,BatchNo];

    Invoke("AddBatchclassNo", args, res);
}
exports.getCourse = async function (req, res, next) {
    QueryAll("GetCourse",res)
}
exports.takeAdmission = async function (req, res, next) {
    var studentName = req.body.student_Name;
    var studentDOB = req.body.student_DOB;
    var email = req.body.email;
    var contactNumber = req.body.contact_Number;
    var alterNo = req.body.alter_No;
    var address = req.body.address;
     var instituteID = req.body.institute_ID;
    var courseId = req.body.course_Id;

    
    var args = [studentName,studentDOB,email,contactNumber,alterNo,address,instituteID,courseId];

    Invoke("TakeAdmission", args, res);
}
exports.editStudent = async function (req, res, next) {
    var studentID = req.body.student_ID;
    var studentName = req.body.student_Name;
    var studentDOB = req.body.student_DOB;
    var email = req.body.email;
    var contactNumber = req.body.contact_Number;
    var alterNo = req.body.alter_No;
     var address = req.body.address;

    
    var args = [studentID,studentName,studentDOB,email,contactNumber,alterNo,address];

    Invoke("EditStudent", args, res);
}

exports.queryInstitute = async function (req, res, next) {
    var instituteID = req.query.institute_ID;
    var args = [instituteID];

    Query("QueryInstitute",args,res)
} 

exports.getStudent = async function (req, res, next) {
    QueryAll("GetStudent",res)
}
exports.getStforAppr = async function (req, res, next) {
    QueryAll("GetStforAppr",res)
}
exports.getApprovedStudents = async function (req, res, next) {
    QueryAll("GetApprovedStudents",res)
}
exports.enrollStudent = async function (req, res, next) {
    var studentID = req.body.student_ID;
    var batchNo = req.body.batch_No;
    var classNo = req.body.class_No;

    
    var args = [studentID,batchNo,classNo];

    Invoke("EnrollStudent", args, res);
}
exports.requestCertificates = async function (req, res, next) {
    var studentID = req.body.student_Id;

    
    var args = [studentID];

    Invoke("RequestCertificates", args, res);
}
exports.getRequestCertificates = async function (req, res, next) {
    QueryAll("GetRequestCertificates",res)
}
exports.issueCertificate = async function (req, res, next) {
    var requestId = req.body.request_Id;

    
    var args = [requestId];

    Invoke("IssueCertificate", args, res);
}
exports.getCertificates = async function (req, res, next) {
    QueryAll("GetCertificates",res)
}

exports.getStudentFromInstitute = async function (req, res, next) {
    var instituteID = req.query.institute_ID;
    var args = [instituteID];

    Query("GetStudentFromInstitute",args,res)
}
exports.getStudentFromCourse = async function (req, res, next) {
    var courseId = req.query.course_Id;
    var args = [courseId];

    Query("GetStudentFromCourse",args,res)
}  
exports.getCourseFromInstitute = async function (req, res, next) {
    var InstituteID = req.query.institute_ID;

    
    var args = [InstituteID];

    Query("GetCourseFromInstitute", args, res);
}

exports.receiveCertificate = async function (req, res, next) {
    var certificateId = req.body.certificate_Id;

    
    var args = [certificateId];

    Invoke("ReceiveCertificate", args, res);
}
exports.issueCertificateCourse = async function (req, res, next) {
    var courseId = req.body.course_Id;

    
    var args = [courseId];

    Invoke("IssueCertificateCourse", args, res);
}
exports.issueCertificateForStudent = async function (req, res, next) {
    var studentId = req.body.student_Id;

    
    var args = [studentId];

    Invoke("IssueCertificateForStudent", args, res);
}
exports.getStudentFromCourseBatchno = async function (req, res, next) {
    var courseId = req.query.course_Id;
    var batchNo = req.query.batch_No;

    
    var args = [courseId,batchNo];

    Query("GetStudentFromCourseBatchno", args, res);
}
exports.viewStudent = async function (req, res, next) {
    var studentId = req.query.student_Id;
    var args = [studentId];

    Query("QueryStudent",args,res)
} 
exports.Querycertstu = async function (req, res, next) {
    var studentId = req.query.student_Id;
    var args = [studentId];

    Query("Querycertstu",args,res)
} 
exports.viewCourse = async function (req, res, next) {
    var courseId = req.query.course_Id;
    var args = [courseId];

    Query("QueryCourse",args,res)
} 
exports.viewCertificate = async function (req, res, next) {
    var certificateId = req.query.certificate_Id;
    var args = [certificateId];

    Query("QueryCertificate",args,res)
}  
exports.getClassList = async function (req, res, next) {
    var instituteID = req.query.institute_ID;
    var args = [instituteID];

    Query("GetClassList",args,res)
}  
exports.getBatchList = async function (req, res, next) {
    var instituteID = req.query.institute_ID;
    var args = [instituteID];

    Query("GetBatchList",args,res)
}  
exports.requestCertificateChange = async function (req, res, next) {
    var certificateId = req.body.certificate_Id;
    var studentName = req.body.student_Name;
    var studentDOB = req.body.student_DOB;

    
    var args = [certificateId,studentName,studentDOB];

    Invoke("RequestCertificateChange", args, res);
}
exports.getRequestforCertiChange = async function (req, res, next) {
    QueryAll("GetRequestforCertiChange",res)
}
exports.approveCertificateChange = async function (req, res, next) {
    var requestId = req.body.request_Id;

    
    var args = [requestId];

    Invoke("ApproveCertificateChange", args, res);
}
exports.changeInstituteOwner = async function (req, res, next) {
    var instituteID = req.body.institute_ID;
    var instituteOwner = req.body.institute_Owner;

    
    var args = [instituteID,instituteOwner];

    Invoke("ChangeInstituteOwner", args, res);
}


exports.editInstitute = async function (req, res, next) {
    var instituteID = req.body.institute_ID;
    var instituteName = req.body.institute_Name;
    var address = req.body.address;
    var contactNumber = req.body.contact_Number;
    var website = req.body.website;
    var email = req.body.email;
    var shortDescription = req.body.short_Description;

    
    var args = [instituteID,instituteName,address,contactNumber,website,email,shortDescription];

    Invoke("EditInstitute", args, res);
}
exports.delete = async function (req, res, next) {
    var ID = req.body.id;
    
    var args = [ID];

    Invoke("Delete", args, res);
}


exports.getInstituteIdFromName = async function (req,res,next) {
    var instituteName =req.query.institute_Name
    var args =[instituteName]
    Query("GetInstituteIdFromName",args,res)
}
exports.getStudentIdFromName = async function (req, res, next) {
    var studentName = req.query.student_Name;
    var args = [studentName];

    Query("GetStudentIdFromName",args,res)
} 

async function Invoke(funcName,args,res){
    try {
        // load the network configuration
        const ccpPath = path.resolve(__dirname, '..', '..', 'first-network', 'connection-org1.json');
        let ccp = JSON.parse(fs.readFileSync(ccpPath, 'utf8'));

        // Create a new file system based wallet for managing identities.
        const walletPath = path.join(process.cwd(), 'wallet');
        const wallet = await Wallets.newFileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);

        // Check to see if we've already enrolled the user.
        const identity = await wallet.get('user1');
        if (!identity) {
            console.log(`An identity for the user user1 does not exist in the wallet`);
            console.log('Run the registerUser.js application before retrying');
            return;
        }

        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();
        await gateway.connect(ccp, { wallet, identity: 'user1', discovery: { enabled: true, asLocalhost: true } });

        // Get the network (channel) our contract is deployed to.
        const network = await gateway.getNetwork('mychannel');

        // Get the contract from the network.
        const contract = network.getContract('fabcar');
        if(args.length == 2 ){
       
        await contract.submitTransaction(funcName,args[0],args[1]);
    }else if(args.length == 1){
     
        await contract.submitTransaction(funcName,args[0]);   
    }else if(args.length == 4){
     
        await contract.submitTransaction(funcName,args[0],args[1],args[2],args[3]);   
    }else if(args.length == 3){
     
        await contract.submitTransaction(funcName,args[0],args[1],args[2]);   
    }else if(args.length == 6){
     
        await contract.submitTransaction(funcName,args[0],args[1],args[2],args[3],args[4],args[5]);   
    }else if(args.length == 7){
     
        await contract.submitTransaction(funcName,args[0],args[1],args[2],args[3],args[4],args[5],args[6]);   
    }else if(args.length == 8){
     
        await contract.submitTransaction(funcName,args[0],args[1],args[2],args[3],args[4],args[5],args[6],args[7]);   
    }
        console.log({Message:'Transaction has been submitted'});
        res.send({Message:'Transaction has been submitted'});

        // Disconnect from the gateway.
        await gateway.disconnect();

    } catch (error) {
        console.error(`Failed to submit transaction: ${error}`);
        res.send(`Failed to submit transaction: ${error}`);

    }
}
async function QueryAll(funcName,res){
    try {
        // load the network configuration
        const ccpPath = path.resolve(__dirname, '..', '..', 'first-network', 'connection-org1.json');
        const ccp = JSON.parse(fs.readFileSync(ccpPath, 'utf8'));

        // Create a new file system based wallet for managing identities.
        const walletPath = path.join(process.cwd(), 'wallet');
        const wallet = await Wallets.newFileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);

        // Check to see if we've already enrolled the user.
        const identity = await wallet.get('user1');
        if (!identity) {
            console.log(`An identity for the user user1does not exist in the wallet`);
            console.log('Run the registerUser.js application before retrying');
            return;
        }

        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();
        await gateway.connect(ccp, { wallet, identity: 'user1', discovery: { enabled: true, asLocalhost: true } });

        // Get the network (channel) our contract is deployed to.
        const network = await gateway.getNetwork('mychannel');

        // Get the contract from the network.
        const contract = network.getContract('fabcar');

        // Evaluate the specified transaction.
        // queryCar transaction - requires 1 argument, ex: ('queryCar', 'CAR4')
        // queryAllCars transaction - requires no arguments, ex: ('queryAllCars')
    
        const result = await contract.evaluateTransaction(funcName);
let p = JSON.parse(result)
        console.log({Result:p});
        res.send({Result:p});

    } catch (error) {
        console.error(`Failed to evaluate transaction: ${error}`);
        res.send(`Failed to evaluate transaction: ${error}`);

    }
}
async function Query(funcName,args,res){
    try {
        // load the network configuration
        const ccpPath = path.resolve(__dirname, '..', '..', 'first-network', 'connection-org1.json');
        const ccp = JSON.parse(fs.readFileSync(ccpPath, 'utf8'));

        // Create a new file system based wallet for managing identities.
        const walletPath = path.join(process.cwd(), 'wallet');
        const wallet = await Wallets.newFileSystemWallet(walletPath);
        //console.log(`Wallet path: ${walletPath}`);

        // Check to see if we've already enrolled the user.
        const identity = await wallet.get('user1');
        if (!identity) {
            console.log(`An identity for the user user1does not exist in the wallet`);
            console.log('Run the registerUser.js application before retrying');
            return;
        }

        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();
        await gateway.connect(ccp, { wallet, identity: 'user1', discovery: { enabled: true, asLocalhost: true } });

        // Get the network (channel) our contract is deployed to.
        const network = await gateway.getNetwork('mychannel');

        // Get the contract from the network.
        const contract = network.getContract('fabcar');

        // Evaluate the specified transaction.
        // queryCar transaction - requires 1 argument, ex: ('queryCar', 'CAR4')
        // queryAllCars transaction - requires no arguments, ex: ('queryAllCars')
        if(args.length == 1 ){
       
        const result = await contract.evaluateTransaction(funcName,args[0]);
        let p;
        
        if(funcName !='GetStudentIdFromName' && funcName != "GetInstituteIdFromName"){
            p = JSON.parse(result)
        }else if(funcName !='GetStudentIdFromName' || funcName != "GetInstituteIdFromName"){
            p = result.toString('utf8')
        }

        console.log({Result:p});
        res.send({Result:p});
        }else if(args.length == 2 ){
       
            const result = await contract.evaluateTransaction(funcName,args[0],args[1]);
            let p = JSON.parse(result)
            console.log({Result:p});
            res.send({Result:p});
            }
    } catch (error) {
        console.log("Failed to evaluate transaction: "+error);
        res.send("Failed to evaluate transaction: "+error);

    }
}




