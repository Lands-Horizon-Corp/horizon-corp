import { defineConfig } from "vitepress";

import { enConfig } from "./en.mts";
import { filConfig } from "./fil.mts";
import { shared } from "./shared.mts";

export default defineConfig({
    ...shared,
    locales: {
        root: { label: "English", ...enConfig },
        fil: { label: "Filipino", ...filConfig},
    },
});
