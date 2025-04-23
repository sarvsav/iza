// @ts-check
import { defineConfig } from 'astro/config';
import starlight from '@astrojs/starlight';

// https://astro.build/config
export default defineConfig({
	integrations: [
		starlight({
			title: 'Iza Docs',
			social: [
				{ icon: 'github', label: 'GitHub', href: 'https://github.com/sarvsav/iza'},
				// { twitter: 'https://twitter.com/sarvsav' },
				// codeberg: 'https://codeberg.org/knut/examples',
				// discord: 'https://astro.build/chat',
				// gitlab: 'https://gitlab.com/delucis',
				// linkedin: 'https://www.linkedin.com/company/astroinc',
				// mastodon: 'https://m.webtoo.ls/@astro',
				// threads: 'https://www.threads.net/@nmoodev',
				// twitch: 'https://www.twitch.tv/bholmesdev',
				// 'x.com': 'https://x.com/astrodotbuild',
				// youtube: 'https://youtube.com/@astrodotbuild',
				],
			description: 'Iza is a tool to do operations on your databases, cicd, artfactory with linux commands',
			sidebar: [
				{
					label: 'Getting started',
					items: [
						{ label: 'Installation', slug: 'getting-started/installation' },
						{ label: 'Configuration', slug: 'getting-started/configuration' },
					],
				},
				{
					label: 'Knowledge Base',
					items: [
						{ label: 'Commands', slug: 'knowledge-base/commands' },
						{ label: 'Variables', slug: 'knowledge-base/variables' },
					],
				},
				{
					label: 'Commands',
					items: [
						{ label: 'cat', slug: 'commands/cat' },
						{ label: 'du', slug: 'commands/du' },
						{ label: 'ls', slug: 'commands/ls' },
						{ label: 'touch', slug: 'commands/touch' },
						{ label: 'whoami', slug: 'commands/whoami' },
					],
				},
				{
					label: 'Contributing',
					autogenerate: { directory: 'contribute'},
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
