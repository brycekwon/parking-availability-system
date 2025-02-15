import globals from "globals";
import pluginJS from "@eslint/js";


/** @type {import('eslint').Linter.Config[]} */
export default [
    {
        languageOptions: {
            globals: globals.browser
        }
    },
    pluginJS.configs.recommended,
];