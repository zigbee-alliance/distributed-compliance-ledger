# PKI types

#    messages
ignite scaffold --module pki message RemoveNocX509IcaCert subject subjectKeyId serialNumber --signer signer
ignite scaffold --module pki message RemoveNocX509RootCert subject subjectKeyId serialNumber --signer signer
