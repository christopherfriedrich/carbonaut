module.exports = {
  content: [
    // @carbonaut-cloud/web and @carbonaut-cloud/docs
    'components/**/*.{js,ts,jsx,tsx}',
    'layouts/**/*.{js,ts,jsx,tsx}',
    'pages/**/*.{js,ts,jsx,tsx}',

    'src/**/*.{js,ts,jsx,tsx}',
    // include packages if not transpiling
    // "../../packages/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {},
  },
  plugins: [],
};
