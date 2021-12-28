# PKI types

#    plain ones
starport scaffold --module compliancetest type TestingResult vid:int pid:int softwareVersion:uint softwareVersionString owner testResult testDate

#    messages
starport scaffold --module compliancetest message AddTestingResult vid:int pid:int softwareVersion:uint softwareVersionString testResult testDate --signer signer

# CRUD data types
starport scaffold --module compliancetest map TestingResults results:strings softwareVersionString --index vid:int,pid:int,softwareVersion:uint --no-message


