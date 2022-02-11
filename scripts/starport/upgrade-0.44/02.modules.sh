starport scaffold module dclauth
starport scaffold module validator --dep dclauth
starport scaffold module dclgenutil --dep dclauth,validator  
# svn checkout https://github.com/cosmos/cosmos-sdk/tags/v0.44.3/x/genutil x/genutil && rm -rf x/genutil/.svn
starport scaffold module pki --dep dclauth  
starport scaffold module vendorinfo --dep dclauth  
starport scaffold module model --dep dclauth  
starport scaffold module compliance --dep dclauth,model
