const colors = require('tailwindcss/colors');

const config = {
	content: ['./src/**/*.{html,js,svelte,ts}'],

	theme: {
		extend: {
			colors: {
				bg1: {
					DEFAULT: colors.neutral['200'],
					dark: colors.neutral['800']
				},
				bg2: {
					DEFAULT: '#dddddd',
					dark: '#1f1f1f'
				},
				bg3: {
					DEFAULT: colors.neutral['50'],
					dark: '#323232'
				},
				shadow: {
					DEFAULT: colors.neutral['400'],
					dark: colors.black
				},
				text1: {
					DEFAULT: colors.black,
					dark: colors.neutral['200']
				},
				text2: {
					DEFAULT: colors.neutral['800'],
					dark: colors.neutral['400']
				},
				nav: {
					DEFAULT: '#7f1d1d',
					dark: '#7f1d1d'
				}
			}
		}
	},

	plugins: [require('tw-elements/dist/plugin')]
};

module.exports = config;
