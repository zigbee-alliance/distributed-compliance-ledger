# validator types
# messages
#   ProposeAddAccount
starport scaffold --module dclauth message ProposeAddAccount address pubKey roles:strings vendorID:uint --signer signer
#   ApproveAddAccount
starport scaffold --module dclauth message ApproveAddAccount address --signer signer
#   ProposeRevokeAccount
starport scaffold --module dclauth message ProposeRevokeAccount address --signer signer
#   ApproveRevokeAccount
starport scaffold --module dclauth message ApproveRevokeAccount address --signer signer
# CRUD data types
#    Account
starport scaffold --module dclauth map Account roles:strings vendorID:uint --index address --no-message
#    PendingAccount
starport scaffold --module dclauth map PendingAccount approvals:strings --index address --no-message
#   PendingAccountRevocation
starport scaffold --module dclauth map PendingAccountRevocation approvals:strings --index address --no-message
#   AccountStat (for legacy AccountNumberCounter)
starport scaffold --module dclauth single AccountStat number:uint --no-message
