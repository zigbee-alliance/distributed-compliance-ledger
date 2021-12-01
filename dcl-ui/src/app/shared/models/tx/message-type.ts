export enum MessageType {
  StdTx = 'cosmos-sdk/StdTx',
  AddModelInfo = 'modelinfo/AddModelInfo',
  UpdateModelInfo = 'modelinfo/UpdateModelInfo',
  DeleteModelInfo = 'modelinfo/DeleteModelInfo',
  AddTestingResult = 'compliancetest/AddTestingResult',
  CertifyModel = 'compliance/CertifyModel',
  RevokeModel = 'compliance/RevokeModel',
  ProposeAddX509RootCert = 'pki/ProposeAddX509RootCert',
  ApproveAddX509RootCert = 'pki/ApproveAddX509RootCert',
  AddX509Cert = 'pki/AddX509Cert'
}
