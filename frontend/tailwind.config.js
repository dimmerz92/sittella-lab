/** @type {import('tailwindcss').Config} */
export default {
	content: ['./src/**/*.{html,js,svelte,ts}'],
	theme: {
		extend: {
			colors: {
				'sit-orange': 'rgb(255, 102, 1)',
				'sit-yellow': 'rgb(243, 222, 138)',
				'sit-lblue': 'rgb(80, 178, 192)',
				'sit-dblue': 'rgb(0, 38, 66)'
			},
			fontSize: {
				'clamp-sm': 'clamp(1.75rem, 5vw, 6vw)',
				'clamp-xl': 'clamp(1.75rem, 4vw, 5vw)'
			},
			screens: {
				'3xl': '1920px',
				'4xl': '2560px'
			}
		}
	},
	plugins: []
};
