import { init } from './main';
import base64url from 'base64url';

export const generate = async () => {
  const wasm = await init();
  const request = JSON.stringify({ command: 'generate' });
  const response = wasm.handleRequest(request);
  return JSON.parse(response);
};

export const sign = async (message: Uint8Array, jwk: any) => {
  const wasm = await init();
  const request = JSON.stringify({
    command: 'sign',
    message: base64url.encode(Buffer.from(message)),
    jwk: JSON.stringify(jwk),
  });
  const response = wasm.handleRequest(request);
  const parsed = JSON.parse(response);
  return parsed.signature;
};

export const verify = async (
  message: Uint8Array,
  signature: string,
  jwk: any
) => {
  const wasm = await init();
  const request = JSON.stringify({
    command: 'verify',
    message: base64url.encode(Buffer.from(message)),
    signature,
    jwk: JSON.stringify(jwk),
  });
  const response = wasm.handleRequest(request);
  const parsed = JSON.parse(response);
  return parsed.verified;
};
