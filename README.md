# this project provides rest api to ceate/list/delete image from harbor docker registry

bay-create:
1. magnum uploads images(like nginx dashboard) into harbor by 'magnum image-upload'.
2. magnum creates admin role user for tenant, then writes username and password into cluster
3. magnum creates a private project(named tenant name) for tenant.
4. magnum uses heat params to create hubimage service into cluster.
5. user list/create/delete image by hubimage srvice on openstack dashboard.
6. terminal user users heat params to login harbor, then do somethings.

bay-update:
do nothing

bay-delete:
1. delete images
2. delete project
3. delete user


hubimage User Case:
image:
curl -X POST  -H "Content-Type: application/json" -d '{"file": "/root/busybox_test", "projectName": "library","name": "busybox","tag": "test","enableWait": true }'  http://x.x.x.x:30909/api/v1/image
curl -X GET  -H "Content-Type: application/json" -d '{ "projectName": "library", "page": 1, "pageSize": 10}'  http://x.x.x.x:30909/api/v1/image
curl -X GET  -H "Content-Type: application/json" -d '{ "projectName":"library", "tag": ""}'  http://x.x.x.x:30909/api/v1/image/busybox
curl -X DELETE  -H "Content-Type: application/json" -d '{"projectName": "library", "tag": "test" }'  http://x.x.x.x:30909/api/v1/image/busybox

project:
curl -X GET  -H "Content-Type: application/json" -d '{"name": "", "public": true ,"owner": "", "page": 1, "pageSize": 10}'  http://x.x.x.x:30909/api/v1/project
curl -X GET  http://x.x.x.x:30909/api/v1/project/pp
curl -X DELETE  http://x.x.x.x:30909/api/v1/project/pp
curl -X POST  -H "Content-Type: application/json" -d '{"name": "", "public": true }'  http://x.x.x.x:30909/api/v1/project

user:
curl -X POST  -H "Content-Type: application/json" -d '{"username": "test3", "password": "Passw0rd","email": "","realname": "","comment": "","role_name": "","has_admin_role": 0 }'  http://x.x.x.x:30909/api/v1/user
curl -X GET  -H "Content-Type: application/json" -d '{"username": "", "email":"","page": 1, "pageSize": 10}'  http://x.x.x.x:30909/api/v1/user
curl -X GET http://x.x.x.x:30909/api/v1/user/test3
curl -X DELETE http://x.x.x.x:30909/api/v1/user/test3

Note: x.x.x.x is a ip of kubernetes cluster node
