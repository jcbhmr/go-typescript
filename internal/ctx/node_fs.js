import * as std from "std";
import * as os from "os";

export function readFileSync(path, options = undefined) {
  let encoding;
  if (typeof options === "object") {
    encoding = options.encoding;
  } else {
    encoding = `${options}`;
  }
  if (encoding === "utf8" || encoding === "utf-8") {
    return std.loadFile(path);
  } else {
    const errorObj = { errno: undefined };
    const file = std.open(path, "r", errorObj);
    if (errorObj.errno) {
      throw new Error(`${path}: ${std.strerror(errorObj.errno)}`);
    }
    file.seek(0, os.SEEK_END);
    const size = file.tell();
    file.seek(0, os.SEEK_SET);
    const bytes = new Uint8Array(size);
    std.read(file.fd, bytes);
    std.close(file.fd);
    return bytes;
  }
}

export function realpathSync(path) {
    const [str, err] = os.realpath(path);
    if (err != null) {
        r
    }
}


