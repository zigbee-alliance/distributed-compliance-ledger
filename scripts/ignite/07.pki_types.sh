# PKI types

#    messages
ignite scaffold --module pki message RemoveNocX509IcaCert subject subjectKeyId serialNumber --signer signer
ignite scaffold --module pki message RemoveNocX509RootCert subject subjectKeyId serialNumber --signer signer

# CRUD data types
ignite scaffold --module pki map NocCertificatesByVidAndSkid certs:strings tq:uint schemaVersion:uint --index vid:uint,subjectKeyId --no-message
