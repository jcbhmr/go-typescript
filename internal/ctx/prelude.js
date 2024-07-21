import * as std from "std";

function format(...args) {
  if (args.length >= 1) {
    if (typeof args[0] === "string") {
      let i = 0;
      const formatted = args[0].replace(/(%%)|(%s)|(%d)|(%i)|(%f)|(%o)|(%O)/g, (...p) => {
        // match, p1, p2, /* …, */ pN, offset, string, groups
        const match = p.shift();
        const groups = p.pop();
        const string = p.pop();
        const offset = p.pop();

        if (p[0] != null) {
          return "%%";
        } else if (p[1] != null) {
          return `${args[++i]}`;
        } else if (p[2] != null) {
          return Math.floor(args[++i]).toString();
        } else if (p[3] != null) {
          return Math.floor(args[++i]).toString();
        } else if (p[4] != null) {
          return (+args[++i]).toString();
        } else if (p[5] != null) {
          return JSON.stringify(args[++i]);
        } else if (p[6] != null) {
          return JSON.stringify(args[++i]);
        } else {
          throw new Error("unreachable");
        }
      });
      if (i + 1 < args.length) {
        return `${formatted} ${args.slice(i + 1).map((a) => JSON.stringify(a)).join(" ")}`;
      } else {
        return formatted;
      }
    } else {
      return args.map((a) => JSON.stringify(a)).join(" ");
    }
  } else {
    return "";
  }
}

globalThis.console = {
  debug: (...a) => std.out.puts(format(...a) + "\n"),
  info: (...a) => std.out.puts(format(...a) + "\n"),
  log: (...a) => std.out.puts(format(...a) + "\n"),
  warn: (...a) => std.err.puts(format(...a) + "\n"),
  error: (...a) => std.err.puts(format(...a) + "\n"),
};
