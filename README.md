# @transmute/xmss

[![CI](https://github.com/transmute-industries/xmss/actions/workflows/ci.yml/badge.svg)](https://github.com/transmute-industries/xmss/actions/workflows/ci.yml)

```
npm i @transmute/xmss --save
```

Based on [danielhavir/go-xmss](https://github.com/danielhavir/go-xmss)

## Usage

```ts
import { generate, sign, verify } from "@transmute/xmss";

const message = Buffer.from("hello");
const signature = "AAAAACf...qRWhlbGxv";
const jwk = await generate();
// const jwk = {
//   kty: 'PQK',
//   alg: 'xmss.SHA2_10_256',
//   x: 'xZppyqOqWrydjvrxgOdpg-ORa8Y1IIsni32luKKn-fNP8oe1hznHgVU9UE5_KE1F9s4qDbbsqHGMxJfmW4Ca_g',
//   d: 'AA..._nz',
// };
const signature = await sign(message, jwk);
const verified = await verify(message, signature, jwk);
```

### Related Work

- [verifiable-data](https://github.com/transmute-industries/verifiable-data)
- [did-key.js](https://github.com/transmute-industries/did-key.js)
- [sidetree.js](https://github.com/transmute-industries/sidetree.js)
- [universal-wallet](https://github.com/transmute-industries/universal-wallet)
- [did actor api](https://github.com/transmute-industries/api.did.actor)

#### Standards

- [W3C Decentralized Identifiers](https://www.w3.org/TR/did-core/)
- [W3C Verifiable Credentials](https://www.w3.org/TR/vc-data-model/)
- [JSON Web Token (JWT)](https://datatracker.ietf.org/doc/html/rfc7519)
- [JSON Web Key (JWK)](https://datatracker.ietf.org/doc/html/rfc7517)
- [Bitcoin Improvement Protocol 39](https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki)
- [Bitcoin Improvement Protocol 44](https://github.com/bitcoin/bips/blob/master/bip-0044.mediawiki)

#### Community Drafts

- [W3C CCG did:key method spec](https://github.com/w3c-ccg/did-method-key)
- [W3C CCG did:web method spec](https://github.com/w3c-ccg/did-method-web)
- [W3C CCG Verifiable Credentials API](https://github.com/w3c-ccg/vc-api)
- [W3C CCG Traceability Vocabulary](https://w3id.org/traceability)
- [W3C CCG Traceability Interoperability Profile](https://w3id.org/traceability/interoperability)

#### Powered By

- [Transmute](https://transmute.industries/)
