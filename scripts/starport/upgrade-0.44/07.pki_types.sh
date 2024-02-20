# PKI types

#    plain ones
starport scaffold --module pki type CertificateIdentifier subject subjectKeyId 
starport scaffold --module pki type Certificate pemCert serialNumber issuer authorityKeyId rootSubject rootSubjectKeyId isRoot:bool owner subject subjectKeyId
#starport scaffold --module pki type CertificateInfo subject subjectKeyId serialNumber issuer authorityKeyId rootSubject rootSubjectKeyId isRoot:bool owner 

#    messages
starport scaffold --module pki message ProposeAddX509RootCert cert --signer signer
starport scaffold --module pki message ApproveAddX509RootCert subject subjectKeyId --signer signer
starport scaffold --module pki message AddX509Cert cert --signer signer
starport scaffold --module pki message ProposeRevokeX509RootCert subject subjectKeyId --signer signer
starport scaffold --module pki message ApproveRevokeX509RootCert subject subjectKeyId --signer signer
starport scaffold --module pki message RevokeX509Cert subject subjectKeyId --signer signer
starport scaffold --module pki message RemoveX509Cert subject subjectKeyId serialNumber --signer signer
starport scaffold --module pki message RejectAddX509RootCert cert --signer signer
starport scaffold --module pki message add-pki-revocation-distribution-point vid:uint pid:uint isPAA:bool label crlSignerCertificate issuerSubjectKeyID dataURL dataFileSize:uint dataDigest dataDigestType:uint revocationType:uint --signer signer
starport scaffold --module pki message update-pki-revocation-distribution-point vid:uint label crlSignerCertificate issuerSubjectKeyID dataURL dataFileSize:uint dataDigest dataDigestType:uint --signer signer
starport scaffold --module pki message delete-pki-revocation-distribution-point vid:uint label issuerSubjectKeyID --signer signer
starport scaffold --module pki message AddNocX509RootCert  cert --signer signer

# CRUD data types
starport scaffold --module pki map ApprovedCertificates certs:strings --index subject,subjectKeyId --no-message
starport scaffold --module pki map ProposedCertificate pemCert serialNumber owner approvals:strings --index subject,subjectKeyId --no-message
starport scaffold --module pki map ChildCertificates certIds:strings --index issuer,authorityKeyId --no-message
starport scaffold --module pki map ProposedCertificateRevocation  approvals:strings --index subject,subjectKeyId --no-message
starport scaffold --module pki map RevokedCertificates certs:strings --index subject,subjectKeyId --no-message
starport scaffold --module pki map UniqueCertificate present:bool --index issuer,serialNumber --no-message
starport scaffold --module pki map PKIRevocationDistributionPoint --index vid:uint,label,issuerSubjectKeyID pid:uint isPAA:bool crlSignerCertificate dataURL dataFileSize:uint dataDigest dataDigestType:uint revocationType:uint --signer signer --no-message
starport scaffold --module pki map PKIRevocationDistributionPointByIssuerSubjectKeyId points:strings --index issuerSubjectKeyID --no-message
starport scaffold --module pki single ApprovedRootCertificates certs:strings --no-message
starport scaffold --module pki single RevokedRootCertificates certs:strings --no-message
starport scaffold --module pki map ApprovedCertificatesBySubject subjectKeyIds:strings --index subject --no-message
starport scaffold --module pki map ApprovedCertificatesBySubjectKeyId certs:strings --index subjectKeyId --no-message
starport scaffold --module pki map RejectedCertificate pemCert serialNumber owner approvals:strings --index subject,subjectKeyId --no-message
#starport scaffold --module pki map AllProposedCertificates --index subject,subjectKeyId --no-message
starport scaffold --module pki map NocRootCertificates certs:strings --index vid:uint --no-message

# Allow colons (:) in subject ID part in REST URLs
# TODO: need to copy the generated query.pb.gw.go into the correct folder
protoc -I "proto" -I "third_party/proto" --grpc-gateway_out=logtostderr=true,allow_colon_final_segments=true:. "--gocosmos_out=plugins=interfacetype+grpc,Mgoogle/protobuf/any.proto=github.com/cosmos/cosmos-sdk/codec/types:." proto/pki/query.proto 
