# validator types
#    plain ones
starport scaffold --module validator type description name identity website details
#    messages
starport scaffold --module validator message CreateValidator pubKey description:Description --signer signer
# CRUD data types
#    Validator
starport scaffold --module validator map Validator description:Description pubKey power:int jailed:bool jailedReason --index owner --no-message
#    LastValidatorPower
starport scaffold --module validator map LastValidatorPower power:int --index owner --no-message
#    ValidatorSigningInfo
# starport scaffold --module validator map ValidatorSigningInfo startHeight:uint indexOffset:uint missedBlocksCounter:uint --index owner --no-message
#    ValidatorMissedBlockBitArray
# starport scaffold --module validator map ValidatorMissedBlockBitArray --index owner,index:uint --no-message
