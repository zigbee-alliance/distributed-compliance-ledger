# model types
# plain ones
starport scaffold --module model type Product pid:int name partNumber

# CRUD data types
#   VendorProduct
#   Change `products` field type to array of Product after scaffolding
starport scaffold --module model map VendorProducts products:Product --index vid:int --no-message
#   Model
starport scaffold --module model map Model deviceTypeId:int productName productLabel partNumber commissioningCustomFlow:int commissioningCustomFlowUrl commissioningModeInitialStepsHint:uint commissioningModeInitialStepsInstruction commissioningModeSecondaryStepsHint:uint commissioningModeSecondaryStepsInstruction userManualUrl supportUrl productUrl lsfUrl lsfRevision --index vid:int,pid:int
#   ModelVersion
starport scaffold --module model map ModelVersion softwareVersionString cdVersionNumber:int firmwareDigests softwareVersionValid:bool otaUrl otaFileSize:uint otaChecksum otaChecksumType:int minApplicableSoftwareVersion:uint maxApplicableSoftwareVersion:uint releaseNotesUrl --index vid:int,pid:int,softwareVersion:uint
#   ModelVersions
starport scaffold --module model map ModelVersions softwareVersions:array.uint --index vid:int,pid:int

# messages
