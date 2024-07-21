import * as std from "std";
import * as os from "os"

export function cwd() {
    const [str, err] = os.getcwd()
    if (err != null) {
        throw err
    }
    return str
}

export const __esModule = false;
export default {
    cwd,
    
};