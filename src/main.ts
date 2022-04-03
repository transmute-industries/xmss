import compressedWasm from './wasm.module';
import { getGo } from './wasm_exec';
const pako = require('pako');

declare var global: any;

var Buffer = require('buffer/').Buffer; // note: the trailing slash is important!
import { isBrowser } from 'browser-or-node';

export const init: any = () => {
  const go = getGo();
  const decompressedWasm = pako.inflate(Buffer.from(compressedWasm, 'base64'));
  return WebAssembly.instantiate(decompressedWasm, go.importObject).then(
    (result) => {
      if (!isBrowser) {
        process.on('exit', (code) => {
          // Node.js exits if no event handler is pending
          if (code === 0 && !go.exited) {
            // deadlock, make Go print error and stack traces
            go._pendingEvent = { id: 0 };
            go._resume();
          }
        });
      }
      go.run(result.instance);
      return global;
    }
  );
};
