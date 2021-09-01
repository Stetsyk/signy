curl -d '{"username":"Oleksii1", "password":"Oleksii"}' -H "Content-Type: application/json" -X POST http://localhost:8000/auth/sign-up
echo
curl -d '{"username":"Oleksii2", "password":"Oleksii"}' -H "Content-Type: application/json" -X POST http://localhost:8000/auth/sign-up
echo
curl -d '{"username":"Oleksii3", "password":"Oleksii"}' -H "Content-Type: application/json" -X POST http://localhost:8000/auth/sign-up
echo
curl -d '{"username":"Oleksii1", "password":"Oleksii"}' -H "Content-Type: application/json" -X POST http://localhost:8000/auth/sign-in
echo 
curl -d '{"username":"Oleksii2", "password":"Oleksii"}' -H "Content-Type: application/json" -X POST http://localhost:8000/auth/sign-in
echo 
curl -d '{"username":"Oleksii3", "password":"Oleksii"}' -H "Content-Type: application/json" -X POST http://localhost:8000/auth/sign-in
echo 
curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzA1NDEzNTcsImlhdCI6MTYzMDQ5ODE1NywidXNlcl9pZCI6MX0.uaa-ZaMK0Y3JNheksT2j1XQIA-hZoWwwVNFdqE7uYro" -X GET http://localhost:8000/api/show-owned-documents
echo
curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzA1NDEzNTcsImlhdCI6MTYzMDQ5ODE1NywidXNlcl9pZCI6MX0.uaa-ZaMK0Y3JNheksT2j1XQIA-hZoWwwVNFdqE7uYro" -d '{"document":{"title":"first", "status":"open", "owner_id":1}, "UsersNeedToSign" : [2, 3] }' -H "Content-Type: application/json" -X POST  http://localhost:8000/api/add-document
echo
curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzA1NDEzNTcsImlhdCI6MTYzMDQ5ODE1NywidXNlcl9pZCI6MX0.uaa-ZaMK0Y3JNheksT2j1XQIA-hZoWwwVNFdqE7uYro" -X GET http://localhost:8000/api/show-owned-documents
echo
echo 
echo
curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzA1NDI4NjYsImlhdCI6MTYzMDQ5OTY2NiwidXNlcl9pZCI6Mn0.cSmPxv4B4GLOtXE7Nimmd_rRZYSKns1XEV8kEoXzH7A"  -X GET http://localhost:8000/api/show-need-to-sign
echo
curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzA1NDI4NjYsImlhdCI6MTYzMDQ5OTY2NiwidXNlcl9pZCI6Mn0.cSmPxv4B4GLOtXE7Nimmd_rRZYSKns1XEV8kEoXzH7A"  -d '{"signature":"approved" }' -H "Content-Type: application/json" -X POST http://localhost:8000/api/document/1/sign-document
echo
curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzA1NDI4NjYsImlhdCI6MTYzMDQ5OTY2NiwidXNlcl9pZCI6Mn0.cSmPxv4B4GLOtXE7Nimmd_rRZYSKns1XEV8kEoXzH7A" -X GET http://localhost:8000/api/document/1/status
echo
curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzA1NDI4NjYsImlhdCI6MTYzMDQ5OTY2NiwidXNlcl9pZCI6M30.FqIybkKeZPmvXF3Ki9UfgAVFe8n7pirsCZOR_EwK-0I"  -d '{"signature":"approved" }' -H "Content-Type: application/json" -X POST http://localhost:8000/api/document/1/sign-document
echo
curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzA1NDI4NjYsImlhdCI6MTYzMDQ5OTY2NiwidXNlcl9pZCI6M30.FqIybkKeZPmvXF3Ki9UfgAVFe8n7pirsCZOR_EwK-0I" -X GET http://localhost:8000/api/document/1/status
echo
