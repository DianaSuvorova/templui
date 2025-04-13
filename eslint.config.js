const htmlPlugin = require("@html-eslint/eslint-plugin");
const eslintHTMLParser =require( "@html-eslint/parser");

module.exports = [
{
    "files": [ "out/**/*.html" ],
    "languageOptions": {
      parser: eslintHTMLParser,
    },
    "plugins": {
      "@html-eslint": htmlPlugin,
    },
    "rules": {
      ...htmlPlugin.configs.recommended.rules,
    }
  },
]