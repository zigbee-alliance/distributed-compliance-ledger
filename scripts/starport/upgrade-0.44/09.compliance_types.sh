# Compliance types

#    plain ones
starport scaffold --module compliance type ComplianceHistoryItem softwareVersionCertificationStatus:uint date reason 

#    messages
starport scaffold --module compliance message CertifyModel vid:int pid:int softwareVersion:uint softwareVersionString cDVersionNumber:uint certificationDate certificationType reason --signer signer
starport scaffold --module compliance message RevokeModel vid:int pid:int softwareVersion:uint softwareVersionString cDVersionNumber:uint revocationDate certificationType reason --signer signer
starport scaffold --module compliance message ProvisionModel vid:int pid:int softwareVersion:uint softwareVersionString cDVersionNumber:uint provisionalDate certificationType reason --signer signer

# CRUD data types
starport scaffold --module compliance map ComplianceInfo softwareVersionString cDVersionNumber:uint softwareVersionCertificationStatus:uint date reason owner history:strings --index vid:int,pid:int,softwareVersion:uint,certificationType --no-message
starport scaffold --module compliance map CertifiedModel value:bool --index vid:int,pid:int,softwareVersion:uint,certificationType --no-message
starport scaffold --module compliance map RevokedModel value:bool --index vid:int,pid:int,softwareVersion:uint,certificationType --no-message
starport scaffold --module compliance map ProvisionalModel value:bool --index vid:int,pid:int,softwareVersion:uint,certificationType --no-message
starport scaffold --module compliance map DeviceSoftwareCompliance ComplianceInfo:strings --index cdCertificateId:string --no-message
