const fs = require('fs');
const path = require('path');
const pako = require('pako');

const wasmFile = fs.readFileSync(path.resolve(__dirname, '../src/main.wasm'));

const compressedWasmFile = Buffer.from(pako.deflate(wasmFile));

fs.writeFileSync(
  path.resolve(__dirname, '../src/wasm.module.ts'),
  `

export default "${compressedWasmFile.toString('base64')}";

`
);
