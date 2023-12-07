# migrationAtlas



ATLAS COMMANDS
-----------DOWNLOADS COMMANDS :----------------------------------------------------------------------
curl -sSf https://atlasgo.sh | sh
-go get -u ariga.io/atlas-provider-beego
---------------------------EXECUTING COMMANDS-------------------------------------------------------------------------
first:- atlas migrate diff --env beego	
second :-  atlas schema apply --env beego -u "postgres://postgres:Dev@123@localhost:5432/atlas_migration"
-----------------------------------------------------------------------------------------------------------------
