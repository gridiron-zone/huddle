---
id: tools-to-build
title: Tools to build DApps
sidebar_label: Tools to build DApps
slug: tools-to-build
---

# Tools to build DApps
In order to let developers easily interact with the Huddle blockchain, we have developed a series of tools that will ease the development of DApps.

## DesmJS

To make it easy to develop dApps on the top of the Huddle chain we have built a `typescript` framework called **DesmJS**. 
DesmJS is based on to the analog framework for Cosmos based chains [CosmJS](https://github.com/cosmos/cosmjs) but focused on Huddle. 
DesmJS contains 3 different packages, each one with a specific scope:
1. The [@huddlelabs/desmjs](https://www.npmjs.com/package/@huddlelabs/desmjs) package contains the client that allows to easily create, sign and broadcast transactions containing Huddle (and Cosmos) messages;
2. The [@huddlelabs/desmjs-types](https://www.npmjs.com/package/@huddlelabs/desmjs-types) package contains the Huddle types protobuf definitions;
3. The [@huddlelabs/desmjs-walletconnect](https://www.npmjs.com/package/@huddlelabs/desmjs-walletconnect) package contains a `WalletConnectSigner` which implements `@huddlelabs/desmjs` `Signer` interface by signing transactions through a WalletConnect client.

The above packages can be found on the official [GitHub Repository](https://github.com/desmos-labs/desmjs) of DesmJS.

### Install DesmJS

```shell
npm install @huddlelabs/desmjs @huddlelabs/desmjs-types
# or
yarn add @huddlelabs/desmjs @huddlelabs/desmjs-types
```

### Send your first transaction

```javascript
import {
    HuddleClient,
    SigningMode,
    OfflineSignerAdapter,
    assertIsDeliverTxSuccess
} from "@huddlelabs/desmjs";


const mnemonic = "math track fish reopen project latin radio spoon please table between install cheap smile deer glide desk license bench vapor chef sock list case";
const rpcEndpoint = "https://rpc.my_tendermint_rpc_node";
const signer = await OfflineSignerAdapter.fromMnemonic(SigningMode.DIRECT, mnemonic);
const [firstAccount] = await signer.getAccounts();

const client = await HuddleClient.connectWithSigner(rpcEndpoint, signer);


const recipient = "huddle1ptvq7l4jt7n9sc3fky22mfvc6waf2jd8nuc0jv";
const amount = {
    denom: "udsm",
    amount: "1337420",
};
const result = await client.sendTokens(firstAccount.address, recipient, [amount], "Have fun with your coins");
assertIsDeliverTxSuccess(result);
```

This is a very easy example but maybe you want to interact with some of the Huddle modules like the **profiles** or **posts**. 
Let's see an example of profile creation and post creation.

### Example 1: Create a profile

```javascript
import { 
    HuddleClient, 
    SigningMode, 
    OfflineSignerAdapter,
    assertIsDeliverTxSuccess,
    MsgSaveProfileEncodeObject
} from "@huddlelabs/desmjs";
import {
    MsgSaveProfile,
} from "@huddlelabs/desmjs-types/huddle/profiles/v3/msgs_profile";


const mnemonic = "math track fish reopen project latin radio spoon please table between install cheap smile deer glide desk license bench vapor chef sock list case";
const rpcEndpoint = "https://rpc.my_tendermint_rpc_node";
const signer = await OfflineSignerAdapter.fromMnemonic(SigningMode.DIRECT, mnemonic);
const [firstAccount] = await signer.getAccounts();

const client = await HuddleClient.connectWithSigner(rpcEndpoint, signer);

const msg: MsgSaveProfileEncodeObject = {
    typeUrl: "/huddle.profiles.v3.MsgSaveProfile",
    value: MsgSaveProfile.fromPartial({
        dtag: "frieza",
        nickname: "Frieza",
        bio: "A weakling like you could never overcome me!",
        profilePicture: "https://link_to_profile_picture",
        coverPicture: "https://link_to_cover_picture",
        creator: firstAccount.address,
    })
}

const result = await client.signAndBroadcast(firstAccount.address, [msg], "auto");
assertIsDeliverTxSuccess(result);
```

### Example 2: Create a post

```javascript
import { 
    HuddleClient, 
    SigningMode, 
    OfflineSignerAdapter, 
    assertIsDeliverTxSuccess, 
    MsgCreatePostEncodeObject 
} from "@huddlelabs/desmjs";
import {
    MsgCreatePost,
} from "@huddlelabs/desmjs-types/huddle/posts/v3/msgs";
import { ReplySetting } from "@huddlelabs/desmjs-types/huddle/posts/v3/models";
import Long from "long";


const mnemonic = "math track fish reopen project latin radio spoon please table between install cheap smile deer glide desk license bench vapor chef sock list case";
const rpcEndpoint = "https://rpc.my_tendermint_rpc_node";
const signer = await OfflineSignerAdapter.fromMnemonic(SigningMode.DIRECT, mnemonic);
const [firstAccount] = await signer.getAccounts();

const client = await HuddleClient.connectWithSigner(rpcEndpoint, signer);

const msg: MsgCreatePostEncodeObject = {
    typeUrl: "/huddle.posts.v3.MsgCreatePost",
    value: MsgCreatePost.fromPartial({
        subspaceId: Long.fromNumber(1),
        sectionId: 0,
        text: "Test post",
        author: firstAccount.address,
        replySettings: ReplySettings.REPLY_SETTING_EVERYONE,
    })
}

const result = await client.signAndBroadcast(firstAccount.address, [msg], "auto");
assertIsDeliverTxSuccess(result);
```

### Example 3: Sign with wallet connect

make sure to install the wallet connect package with the following command:

```shell
npm install @huddlelabs/desmjs-walletconnect
# or
yarn add @huddlelabs/desmjs-walletconnect
```

```javascript
import {
  WalletConnect,
  QRCodeModal,
  WalletConnectSigner,
} from "@huddlelabs/desmjs-walletconnect";
import { HuddleClient, SigningMode, assertIsDeliverTxSuccess, MsgSaveProfileEncodeObject } from "@huddlelabs/desmjs";
import {
    MsgSaveProfile,
} from "@huddlelabs/desmjs-types/huddle/profiles/v3/msgs_profile";


const connector = new WalletConnect({
    bridge: "https://bridge.walletconnect.org",
    qrcodeModal: QRCodeModal,
});
const signer = new WalletConnectSigner(this.connector, {
    signingMode: SigningMode.AMINO,
});

// Show the walletconnect modal to allow connection from a client
await signer.connect();

const rpcEndpoint = "https://rpc.my_tendermint_rpc_node";
const [firstAccount] = await signer.getAccounts();

const client = await HuddleClient.connectWithSigner(rpcEndpoint, signer);

const msg: MsgSaveProfileEncodeObject = {
    typeUrl: "/huddle.profiles.v3.MsgSaveProfile",
    value: MsgSaveProfile.fromPartial({
        dtag: "frieza",
        nickname: "Frieza",
        bio: "A weakling like you could never overcome me!",
        profilePicture: "https://link_to_profile_picture",
        coverPicture: "https://link_to_cover_picture",
        creator: firstAccount.address,
    })
}

const result = await client.signAndBroadcast(firstAccount.address, [msg], "auto");
assertIsDeliverTxSuccess(result);
```

## Huddle Bindings
The [Huddle Bindings](https://github.com/gridiron-zone/huddle-bindings) are a set of packages that make possible to interact with the Huddle chain directly from smart contracts. With them, you can build your own dApp smart contracts taking full advantage of the Huddle chain modules to create even more personalisation to your app.   
You can find the bindings generated documentation here: [Huddle Bindings docs](https://docs.rs/huddle-bindings/1.0.0/huddle_bindings/index.html). 

### Example 1: Post from a contract
The below example shows you how to send a [`MsgCreatePost`](02-modules/posts/04-messages.md#msgcreatepost) from
inside a smart contract.
```rust
pub fn post_example_from_contract(deps: DepsMut, env: Env, info: MessageInfo, message: String) -> Result<Response<HuddleMsg>, ContractError> {
    let post_msg = PostsMsg::CreatePost {
        subspace_id: Uint64::new(1),
        section_id: 1,
        external_id: None,
        text: Some(message),
        entities: None,
        attachments: None,
        author: env.contract.address,
        conversation_id: None,
        reply_settings: ReplySetting::Unspecified,
        referenced_posts: vec![]
    };

    let response = Response::new()
        .add_attribute("action", "post")
        .add_attribute("author", env.contract.to_string())
        .add_messages(post_msg);

    Ok(response)
}
```

### Example 2: Query from a contract the Huddle chain state you need
The below example shows you how to query a Subspace's posts from inside a smart contract.
```rust
fn query_posts_from_contract(deps: Deps<HuddleQuery>, subspace_id: Uint64, pagination: Option<PageRequest>) -> StdResult<Binary> {
    let request = HuddleQuery::Posts(PostsQuery::SubspacePosts {
        subspace_id,
        pagination,
    });
    let response: StdResult<QuerySubspacePostsResponse> = deps.querier.query(&request.into());
}
```

#### Notes
The above examples are really simple examples of what you can achieve with the Huddle Bindings packages inside smart contracts. More examples will be added here in the future, and you can find some other inside the [Huddle Bindings Github repository](https://github.com/gridiron-zone/huddle-bindings).

## GraphQL APIs
The [GraphQL APIs](../07-graphql/01-overview.md) simplifies the way with which clients interact with the Huddle blockchain to obtain the data their applications need. These queries can be customised with a wide range of parameters giving developers a great flexibility while building their apps.

### Example: Query a user posts
#### Query
```graphql
query QueryUserPosts {
  post(where: {author_address: {_eq: "huddle1dx6h75tkj0cuvyqf6cwn6usc9qynu39v0245m4"}}) {
    id
  	text
    creation_date
    last_edited_date
    external_id
    author_address
    attachments {
      content
      id
    }
  }
}
```

#### Response
```graphql
{
  "data": {
    "post": [
      {
        "id": 1,
        "text": "This is a test post",
        "creation_date": "2022-06-30T17:04:54.57816",
        "last_edited_date": "2022-07-01T07:35:48.149871",
        "external_id": null,
        "author_address": "huddle1dx6h75tkj0cuvyqf6cwn6usc9qynu39v0245m4",
        "attachments": [
          {
            "content": {
              "uri": "https://images.app.goo.gl/g7VHpLGJYjndRfWL6",
              "@type": "/huddle.posts.v1.Media",
              "mime_type": "image/png"
            },
            "id": 1
          }
        ]
      }
    ]
  }
}
```