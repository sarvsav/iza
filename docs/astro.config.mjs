// @ts-check
import { defineConfig } from 'astro/config';
import starlight from '@astrojs/starlight';

// https://astro.build/config
export default defineConfig({
	integrations: [
		starlight({
			title: 'Iza Docs',
			social: {
				github: 'https://github.com/sarvsav/iza',
				twitter: 'https://twitter.com/sarvsav',
			},
			description: 'Iza is a tool to do operations on your databases, cicd, artfactory with linux commands',
			sidebar: [
				{
					label: 'Guides',
					items: [
						// Each item here is one entry in the navigation menu.
						{ label: 'Getting started', slug: 'guides/getting-started' },
					],
				},
				{
					label: 'Reference',
					autogenerate: { directory: 'reference' },
				},
			],
			logo: {
				src: './src/assets/logo-iza.png',
				replacesTitle: false, // Hide the site title
			},
			customCss: [
				// Path to your custom CSS file
				'./src/styles/style.css',
		  ],
		}),
	],
});
