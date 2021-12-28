# Compliance types

#    plain ones
starport scaffold --module compliance type ComplianceHistoryItem softwareVersionCertificationStatus:uint date reason 

#    messages
starport scaffold --module compliance message CertifyModel vid:int pid:int softwareVersion:uint softwareVersionString certificationDate certificationType reason --signer signer
starport scaffold --module compliance message RevokeModel vid:int pid:int softwareVersion:uint softwareVersionString revocationDate certificationType reason --signer signer

# CRUD data types
starport scaffold --module compliance map ComplianceInfo softwareVersionString cDVersionNumber:uint softwareVersionCertificationStatus:uint date reason owner history:strings --index vid:int,pid:int,softwareVersion:uint,certificationType --no-message
starport scaffold --module compliance map CertifiedModel value:bool --index vid:int,pid:int,softwareVersion:uint,certificationType --no-message
starport scaffold --module compliance map RevokedModel value:bool --index vid:int,pid:int,softwareVersion:uint,certificationType --no-message
