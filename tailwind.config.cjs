/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./internal/presenter/html/template/templates/**/*.html"],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
};
