const esbuild = require("esbuild");
const postcss = require("esbuild-postcss");

esbuild
  .build({
    entryPoints: ["assets/styles.css", "assets/index.js"],
    bundle: true,
    loader: {
      ".woff2": "file",
    },
    outdir: "dist",
    plugins: [postcss()],
    minify: true,
    sourcemap: true,
  })
  .catch(() => process.exit(1));
