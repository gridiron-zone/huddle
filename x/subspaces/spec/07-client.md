---
id: client
title: Client
sidebar_label: Client
slug: client
---

# Client

## CLI

A user can query and interact with the `subspaces` module using the CLI.

### Query

The `query` commands allow users to query the `subspaces` state.

```bash
huddle query subspaces --help
```

#### subspace
The `subspace` query command allows users to query a subspace with the given id.

```bash
huddle query subspaces subspace [id] [flags]
```

Example:
```bash
huddle query subspaces subspace 1
```

Example output:
```yaml
subspace:
  creation_time: "2022-06-20T14:52:23.049305Z"
  creator: huddle1nwp8gxrnmrsrzjdhvk47vvmthzxjtphgxp5ftc
  description: "this is a test subspace"
  id: "1"
  name: test
  owner: huddle1rfv0f7mx7w9d3jv3h803u38vqym9ygg344asm3
  treasury: huddle1rfv0f7mx7w9d3jv3h803u38vqym9ygg344asm3
```

#### subspaces
The `subspaces` query command allows users to query all the subspaces. Optional pagination is available.

```bash
huddle query subspaces subspaces [flags]
```

Example:
```bash
huddle query subspaces subspaces --page=2 --limit=100
```

Example output:
```yaml
pagination:
  next_key: null
  total: "0"
subspaces:
- creation_time: "2022-06-20T14:52:23.049305Z"
  creator: huddle1nwp8gxrnmrsrzjdhvk47vvmthzxjtphgxp5ftc
  description: ""
  id: "1"
  name: test
  owner: huddle1rfv0f7mx7w9d3jv3h803u38vqym9ygg344asm3
  treasury: huddle1rfv0f7mx7w9d3jv3h803u38vqym9ygg344asm3
```

#### sections
The `sections` query command allows users to query the sections state.

```bash
huddle query subspaces sections [command] [flags] --help
```

##### list
The `list` query command allows users to query all the sections of a subspace. Optional pagination is available.

```bash
huddle query subspaces sections list [subspace-id] [flags]
```

Example:
```bash
huddle query subspaces sections list 1 --page=2 --limit=100
```

Example output:
```yaml
pagination:
  next_key: null
  total: "0"
sections:
- description: This is the default subspace section
  id: 0
  name: Default section
  parent_id: 0
  subspace_id: "1"
- description: ""
  id: 1
  name: "1"
  parent_id: 0
  subspace_id: "1"
- description: this is a test section
  id: 2
  name: "1"
  parent_id: 0
  subspace_id: "1"
- description: this is a test section 2
  id: 3
  name: "1"
  parent_id: 0
  subspace_id: "1"
```

##### section
The `section` query command allows users to query a specific section inside a subspace.

```bash
huddle query subspaces sections section [subspace-id] [section-id] [flags]
```

Example:
```bash
huddle query subspaces sections section 1 2
```

Example output:
```yaml
section:
  description: this is a test section
  id: 2
  name: "1"
  parent_id: 0
  subspace_id: "1"
```

#### groups
The `groups` query command allows users to query the groups state.

```bash
huddle query subspaces groups [command] --help
```

##### list
The `list` query command allows users to query all the groups of a subspace. Optional pagination is available.

```bash
huddle query subspaces groups list [subspace-id] [flags] --help
```

Example:
```bash
huddle query subspaces groups list 1 
```

Example output:
```yaml
groups:
- description: This is a default user group which all users are automatically part
    of
  id: 0
  name: Default
  permissions: []
  section_id: 0
  subspace_id: "1"
- description: A test group
  id: 1
  name: TestGroup
  permissions: []
  section_id: 0
  subspace_id: "1"
pagination:
  next_key: null
  total: "0"
```

#### group
The `group` query command allows users to query a specific group of a subspace.

```bash
huddle query subspaces groups group [subspace-id] [group-id] [flags]
```

Example:
```bash
huddle query subspaces groups group 1 1
```

Example output:
```yaml
group:
  description: A test group
  id: 1
  name: TestGroup
  permissions: []
  section_id: 0
  subspace_id: "1"
```

#### members
The `members` query command allows users to query the members of a group.

```bash
huddle query subspaces members [subspace-id] [group-id] [flags]
```

Example:
```bash
huddle query subspaces members 1 1
```

Example output:
```yaml
members:
- huddle1nwp8gxrnmrsrzjdhvk47vvmthzxjtphgxp5ftc
- huddle1rfv0f7mx7w9d3jv3h803u38vqym9ygg344asm3
pagination:
  next_key: null
  total: "0"
```

#### permissions
The `permissions` query command allows users to query user's permissions of a specific subspace or section.

```bash
huddle query subspaces permissions [subspace-id] [section-id] [user] [flags]
```

Example:
```bash
huddle query subspaces permissions 1 0 huddle1nwp8gxrnmrsrzjdhvk47vvmthzxjtphgxp5ftc  
```

Example output:
```yaml
details:
- section_id: 0
  subspace_id: "1"
  user:
    permission:
    - EVERYTHING
    user: huddle1nwp8gxrnmrsrzjdhvk47vvmthzxjtphgxp5ftc
- group:
    group_id: 0
    permission: []
  section_id: 0
  subspace_id: "1"
permissions:
- EVERYTHING
```

## gRPC
Users can query the `subspaces` module gRPC endpoints.

### Subspaces
The `Subspaces` endpoint allows users to query all the subspaces inside Huddle.

```bash
huddle.subspaces.v2.Query/Subspaces
```

Example:
```bash
grpcurl -plaintext localhost:9090 huddle.subspaces.v2.Query/Subspaces
```

Example output:
```json
{
  "subspaces": [
    {
      "id": "1",
      "name": "test",
      "treasury": "huddle1rfv0f7mx7w9d3jv3h803u38vqym9ygg344asm3",
      "owner": "huddle1rfv0f7mx7w9d3jv3h803u38vqym9ygg344asm3",
      "creator": "huddle1nwp8gxrnmrsrzjdhvk47vvmthzxjtphgxp5ftc",
      "creationTime": "2022-06-20T14:52:23.049305Z"
    }
  ],
  "pagination": {
    "total": "1"
  }
}
```

### Subspace
The `Subspace` endpoint allows users to query a subspace associated with the given ID.

```bash
huddle.subspaces.v2.Query/Subspace
```

Example:
```bash
grpcurl -plaintext -d '{"subspace_id":1}' localhost:9090 huddle.subspaces.v2.Query/Subspace
```

Example output:
```json
{
  "subspace": {
    "id": "1",
    "name": "test",
    "treasury": "huddle1rfv0f7mx7w9d3jv3h803u38vqym9ygg344asm3",
    "owner": "huddle1rfv0f7mx7w9d3jv3h803u38vqym9ygg344asm3",
    "creator": "huddle1nwp8gxrnmrsrzjdhvk47vvmthzxjtphgxp5ftc",
    "creationTime": "2022-06-20T14:52:23.049305Z"
  }
}
```

### Sections
The `Sections` endpoint allows users to query the sections associated with the given subspace ID. 

```bash
huddle.subspaces.v2.Query/Sections
```

Example:
```bash
grpcurl -plaintext -d '{"subspace_id":1}' localhost:9090 huddle.subspaces.v2.Query/Sections 
```

Example output:
```json
{
  "sections": [
    {
      "subspaceId": "1",
      "name": "Default section",
      "description": "This is the default subspace section"
    },
    {
      "subspaceId": "1",
      "id": 1,
      "name": "1"
    },
    {
      "subspaceId": "1",
      "id": 2,
      "name": "1",
      "description": "this is a test section"
    },
    {
      "subspaceId": "1",
      "id": 3,
      "name": "1",
      "description": "this is a test section 2"
    }
  ],
  "pagination": {
    "total": "4"
  }
}
```

### Section
The `Section` endpoint allows users to query a section with the given ID inside the subspace with the given ID.

```bash
huddle.subspaces.v2.Query/Section
```

Example:
```bash
grpcurl -plaintext -d '{"subspace_id":1, "section_id":1}' localhost:9090 huddle.subspaces.v2.Query/Section
```

Example output:
```json
{
  "section": {
    "subspaceId": "1",
    "id": 1,
    "name": "1"
  }
}
```

### UserGroups
The `UserGroups` endpoint allows users to query all the user groups associated with a given subspace ID and section ID.

```bash
huddle.subspaces.v2.Query/UserGroups
```

Example:
```bash
grpcurl -plaintext -d '{"subspace_id":1, "section_id": 0}' localhost:9090 huddle.subspaces.v2.Query/UserGroups
```

Example output:
```json
{
  "groups": [
    {
      "subspaceId": "1",
      "name": "Default",
      "description": "This is a default user group which all users are automatically part of"
    },
    {
      "subspaceId": "1",
      "id": 1,
      "name": "TestGroup",
      "description": "A test group"
    }
  ],
  "pagination": {
    "total": "2"
  }
}
```

### UserGroup
The `UserGroup` endpoint allows users to query a specific user group with the given ID associated with a given subspace ID.

```bash
huddle.subspaces.v2.Query/UserGroup
```

Example:
```bash
grpcurl -plaintext -d '{"subspace_id":1, "group_id":1}' localhost:9090 huddle.subspaces.v2.Query/UserGroup 
```

Example output:
```json
{
  "group": {
    "subspaceId": "1",
    "id": 1,
    "name": "TestGroup",
    "description": "A test group"
  }
}
```

### UserGroupMembers
The `UserGroupMembers` endpoint allows users to query all the members of the user group with the given ID inside the
subspace with the given ID.

```bash
huddle.subspaces.v2.Query/UserGroupMembers
```

Example:
```bash
grpcurl -plaintext -d '{"subspace_id":1, "group_id":1}' localhost:9090 huddle.subspaces.v2.Query/UserGroupMembers
```

Example output:
```json
{
  "members": [
    "huddle1nwp8gxrnmrsrzjdhvk47vvmthzxjtphgxp5ftc",
    "huddle1rfv0f7mx7w9d3jv3h803u38vqym9ygg344asm3"
  ],
  "pagination": {
    "total": "2"
  }
}
```

### UserPermissions
The `UserPermissions` endpoint allows users to query all the user's permissions inside the subspace with the given ID
and the section with the given ID.

```bash
huddle.subspaces.v2.Query/UserPermissions
```

Example:
```bash
grpcurl -plaintext -d '{"subspace_id":1, "section_id":0, "user": "huddle1nwp8gxrnmrsrzjdhvk47vvmthzxjtphgxp5ftc"}' localhost:9090 huddle.subspaces.v2.Query/UserPermissions
```

Example output:
```json
{
  "permissions": [
    "EVERYTHING"
  ],
  "details": [
    {
      "subspaceId": "1",
      "user": {
        "user": "huddle1nwp8gxrnmrsrzjdhvk47vvmthzxjtphgxp5ftc",
        "permission": [
          "EVERYTHING"
        ]
      }
    },
    {
      "subspaceId": "1",
      "group": {
        
      }
    },
    {
      "subspaceId": "1",
      "group": {
        "groupId": 1
      }
    }
  ]
}
```

## REST
A user can query the `subspaces` module using REST endpoints.

### Subspaces
The `Subspaces` endpoint allows users to query all the subspaces inside Huddle.

````
/huddle/subspaces/v3/subspaces
````

### Subspace
The `Subspace` endpoint allows users to query a subspace associated with the given ID.

````
/huddle/subspaces/v3/subspaces/{subspace_id}
````

### Sections
The `Sections` endpoint allows users to query the sections associated with the given subspace ID.

````
/huddle/subspaces/v3/{subspace_id}/sections
````

### Section
The `Section` endpoint allows users to query a section with the given ID associated with the subspace with the given ID.

````
/huddle/subspaces/v3/{subspace_id}/sections/{section_id}
````

### UserGroups
The `UserGroups` endpoint allows users to query all the user groups associated with a given subspace ID.

````
/huddle/subspaces/v3/subspaces/{subspace_id}/groups
````

### UserGroup
The `UserGroup` endpoint allows users to query a specific user group with the given ID associated with a given subspace ID.

````
/huddle/subspaces/v3/subspaces/{subspace_id}/groups/{group_id}
````

### UserGroupMembers
The `UserGroupMembers` endpoint allows users to query all the members of the user group with the given ID inside the
subspace with the given ID.

````
/huddle/subspaces/v3/subspaces/{subspace_id}/groups/{group_id}/members
````

### UserPermissions
The `UserPermissions` endpoint allows users to query all the user's permissions inside the subspace with the given ID
and the section with the given ID.

````
/huddle/subspaces/v3/subspaces/{subspace_id}/permissions/{user}
````