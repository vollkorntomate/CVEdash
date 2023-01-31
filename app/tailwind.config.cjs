const colors = require('tailwindcss/colors')

const config = {
	content: ['./src/**/*.{html,js,svelte,ts}'],

	theme: {
		extend: {
			colors: {
				'bg1': {
					DEFAULT: colors.neutral['200'],
					dark: colors.neutral['800']
				},
				'bg2': {
					DEFAULT: '#dddddd',
					dark: '#1f1f1f'
				},
				'bg3': {
					DEFAULT: colors.neutral['50'],
					dark: colors.neutral['700']
				},
				'shadow': {
					DEFAULT: colors.neutral['400'],
					dark: colors.black
				}
			}
		}
	},

	plugins: []
};

module.exports = config;
